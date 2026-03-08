package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/course/internal/config"
	"example.com/course/internal/course"
	"example.com/course/internal/logging"
	"example.com/course/internal/metrics"
	"example.com/course/internal/tracing"
	"gorm.io/gorm"

	"github.com/agussyahrilmubarok/gox/pkg/xconfig/xviper"
	"github.com/agussyahrilmubarok/gox/pkg/xdiscovery"
	"github.com/agussyahrilmubarok/gox/pkg/xdiscovery/xconsul"
	"github.com/agussyahrilmubarok/gox/pkg/xgorm"
	"github.com/agussyahrilmubarok/gox/pkg/xlogger/xzerolog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "example.com/course/cmd/server/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
	gormLogger "gorm.io/gorm/logger"
)

// @title Course Service API
// @version 1.0
// @description This is the Course Service API.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8082
// @BasePath /api/v1
func main() {
	configFlag := flag.String("config", "configs/config.json", "Path to config file")
	flag.Parse()

	vCfg, err := xviper.NewConfig(*configFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	var cfg *config.Config
	if err := vCfg.Unmarshal(&cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger, err := xzerolog.NewLogger(cfg.Logger.Filepath, cfg.Logger.Level)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %v\n", err)
		os.Exit(1)
	}

	db, err := xgorm.NewGorm("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.SslMode,
	), &xgorm.Options{
		Config: &gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Silent),
		},
		MaxOpenConns:    cfg.Postgres.MaxOpenConns,
		MaxIdleConns:    cfg.Postgres.MaxIdleConns,
		ConnMaxLifetime: cfg.Postgres.ConnMaxLifetime,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
		os.Exit(1)
	}

	if err := db.AutoMigrate(&course.Course{}); err != nil {
		logger.Fatal().Err(err).Msg("auto migrate failed")
		os.Exit(1)
	}

	instanceID := xdiscovery.GenerateInstanceID(cfg.App.Name)
	consulRegistry, err := xconsul.NewRegistry(cfg.Consul.Address)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to register consul discovery")
		os.Exit(1)
	}

	ctx := context.Background()
	if err := consulRegistry.Register(ctx, instanceID, cfg.App.Name, fmt.Sprintf("%v:%d", cfg.App.Host, cfg.App.Port)); err != nil {
		logger.Fatal().Err(err).Msg("failed to register consul discovery")
		os.Exit(1)
	}

	go func() {
		for {
			if err := consulRegistry.ReportHealthyState(instanceID, cfg.App.Name); err != nil {
				logger.Info().Msg("failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer consulRegistry.Deregister(ctx, instanceID, cfg.App.Name)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		tickerChannel := ticker.C

		for range tickerChannel {
			metrics.UpdateCPUUsage()
			metrics.UpdateMemoryUsage()
			metrics.UpdateDiskUsage()
		}
	}()

	traceExporter := tracing.NewOTLPExporter(ctx, fmt.Sprintf("%s:%v", cfg.OTEL.Host, cfg.OTEL.Port))
	shutdownTrace := tracing.InitTraceProvider(ctx, cfg.App.Name, traceExporter)
	tracing.NewTracer(cfg.App.Name)
	defer shutdownTrace(ctx)

	store := course.NewStore(db)
	service := course.NewService(store)
	handler := course.NewHandler(service)

	e := echo.New()
	e.HideBanner = true
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(logging.RequestIDMiddleware(logger))
	e.Use(metrics.MetricAppMiddleware)
	e.Use(tracing.Middleware(cfg.App.Name))
	e.Use(logging.TraceIDMiddleware(tracing.Tracer))

	api := e.Group("/api/v1/courses")
	api.GET("", handler.FindAll)
	api.GET("/find", handler.Find)
	api.POST("/reserve", handler.ReserveCourse)
	api.POST("/release", handler.ReleaseCourse)
	api.POST("/init-dummy", handler.InitDummy)
	api.DELETE("/clean-dummy", handler.CleanDummy)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	go func() {
		logger.Info().Msgf("server running at http://%s:%d", cfg.App.Host, cfg.App.Port)
		logger.Info().Msgf("swagger available at http://%s:%d/swagger/index.html", cfg.App.Host, cfg.App.Port)
		if err := e.Start(fmt.Sprintf(":%d", cfg.App.Port)); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	go func() {
		metricsServer := metrics.NewMetricServer(cfg)
		logger.Info().Msgf("prometheus metrics server running at http://%s:%d/metrics/prometheus", cfg.App.Metric.Host, cfg.App.Metric.Port)
		if err := metricsServer.Start(fmt.Sprintf(":%d", cfg.App.Metric.Port)); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("failed to start prometheus metrics server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	logger.Info().Msg("shutting down server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctxShutdown); err != nil {
		logger.Fatal().Err(err).Msg("graceful shutdown failed")
	}

	logger.Info().Msg("server stopped gracefully")
}

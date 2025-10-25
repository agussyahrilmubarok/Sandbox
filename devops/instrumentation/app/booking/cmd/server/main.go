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

	"example.com/booking/internal/booking"
	"example.com/booking/internal/logging"
	"example.com/booking/pkg/config"
	"example.com/booking/pkg/discovery"
	"example.com/booking/pkg/discovery/consul"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "example.com/booking/cmd/server/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Booking Service API
// @version 1.0
// @description This is the Booking Service API.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
func main() {
	configFlag := flag.String("config", "configs/config.json", "Path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*configFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger, err := config.NewZerolog(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to setup logger: %v\n", err)
		os.Exit(1)
	}

	db, err := config.NewPostgres(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
		os.Exit(1)
	}

	if err := db.AutoMigrate(&booking.Booking{}); err != nil {
		logger.Fatal().Err(err).Msg("AutoMigrate failed")
		os.Exit(1)
	}

	instanceID := discovery.GenerateInstanceID(cfg.App.Name)
	consulRegistry, err := consul.NewRegistry(cfg.Consul.Address)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to register consul discovery")
		os.Exit(1)
	}

	ctx := context.Background()
	if err := consulRegistry.Register(ctx, instanceID, cfg.App.Name, fmt.Sprintf("%v:%d", cfg.App.Host, cfg.App.Port)); err != nil {
		logger.Fatal().Err(err).Msg("Failed to register consul discovery")
		os.Exit(1)
	}

	go func() {
		for {
			if err := consulRegistry.ReportHealthyState(instanceID, cfg.App.Name); err != nil {
				logger.Info().Msg("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer consulRegistry.Deregister(ctx, instanceID, cfg.App.Name)

	store := booking.NewStore(db)
	client := booking.NewClient(consulRegistry)
	service := booking.NewService(store, client)
	handler := booking.NewHandler(service)

	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(logging.RequestIDMiddleware(logger))

	// v1 routes
	apiV1 := e.Group("/api/v1")
	apiV1.POST("/booking/course", handler.Booking)

	// v2 routes
	apiV2 := e.Group("/api/v2")
	apiV2.POST("/booking/course", handler.BookingV2)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":  "ok",
			"service": "booking-service",
		})
	})

	go func() {
		logger.Info().Msgf("Server running at http://%s:%d", cfg.App.Host, cfg.App.Port)
		logger.Info().Msgf("Swagger available at http://%s:%d/swagger/index.html", cfg.App.Host, cfg.App.Port)
		if err := e.Start(fmt.Sprintf(":%d", cfg.App.Port)); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	logger.Info().Msg("Shutting down server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctxShutdown); err != nil {
		logger.Fatal().Err(err).Msg("Graceful shutdown failed")
	}

	logger.Info().Msg("Server stopped gracefully")
}

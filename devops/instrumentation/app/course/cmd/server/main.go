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

	"example.com/course/internal/course"
	"example.com/course/pkg/config"

	_ "example.com/course/cmd/server/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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

	if err := db.AutoMigrate(&course.Course{}); err != nil {
		logger.Fatal().Err(err).Msg("AutoMigrate failed")
		os.Exit(1)
	}

	store := course.NewStore(db, logger)
	service := course.NewService(store, logger)
	handler := course.NewHandler(service, logger)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/courses/init-dummy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.InitDummy(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/api/v1/courses/clean-dummy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handler.CleanDummy(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","service":"course-service"}`)
	})

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		logger.Info().Msgf("Server running at http://%s%s", cfg.App.Host, addr)
		logger.Info().Msgf("Swagger available at http://%s%s/swagger/index.html", cfg.App.Host, addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	logger.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Graceful shutdown failed")
	}

	logger.Info().Msg("Server stopped gracefully")
}

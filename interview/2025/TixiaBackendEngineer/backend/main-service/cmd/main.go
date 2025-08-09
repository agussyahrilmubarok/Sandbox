package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/main-service/internal/handler"
	"example.com/main-service/internal/repository"
	"example.com/main-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	serviceName  = "main-service"
	serviceLevel string
	servicePort  string
)

func main() {
	flag.StringVar(&serviceLevel, "serviceLevel", "development", "Service level: development or production")
	flag.StringVar(&servicePort, "servicePort", "8080", "Port to run the service on")
	flag.Parse()

	baseLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer baseLogger.Sync()

	logger := baseLogger.With(zap.String("service", serviceName))

	redisAddr := "localhost:6379"
	if serviceLevel == "production" {
		redisAddr = "redis:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       serviceName,
	})

	flightRepo := repository.NewFlightRepository(rdb, logger)
	flightUc := usecase.NewFlightUseCase(flightRepo, logger)
	flightHandler := handler.NewFlightHandler(flightUc, logger)

	apiV1 := app.Group("api/v1")
	apiV1.Post("/flights/search", flightHandler.SearchFlights)
	apiV1.Get("/flights/search/:search_id/stream", flightHandler.StreamFlightResults)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		addr := ":" + servicePort
		logger.Info("Starting server", zap.String("port", servicePort))
		if err := app.Listen(addr); err != nil {
			logger.Error("Fiber server stopped", zap.Error(err))
		}
	}()

	sig := <-sigCh
	logger.Info("Received shutdown signal", zap.String("signal", sig.String()))
	time.Sleep(2 * time.Second)
	logger.Info("Shutdown complete")
}

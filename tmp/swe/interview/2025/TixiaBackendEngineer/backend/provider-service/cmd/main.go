package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/provider-service/internal/repository"
	"example.com/provider-service/internal/worker"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	serviceName  = "provider-service"
	serviceLevel string
)

func main() {
	flag.StringVar(&serviceLevel, "serviceLevel", "development", "Service level: development or production")
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

	repo := repository.NewFlightRepository("tmp/sample.json", logger)
	consumer := worker.NewFlightSearchConsumer(repo, rdb, logger)

	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := consumer.Start(ctx); err != nil {
			logger.Error("Consumer stopped with error", zap.Error(err))
		} else {
			logger.Info("Consumer stopped gracefully")
		}
	}()

	sig := <-sigCh
	logger.Info("Received shutdown signal", zap.String("signal", sig.String()))
	cancel()
	time.Sleep(2 * time.Second)
	consumer.Stop(context.Background())
	logger.Info("Shutdown complete")
}

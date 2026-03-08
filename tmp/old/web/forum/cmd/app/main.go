package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"forum.id/internal/app"
	"forum.id/internal/postgres"
	"forum.id/internal/rest"
	"forum.id/pkg/config"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type schema struct {
	App      *fiber.App
	Config   *viper.Viper
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
}

func main() {
	fmt.Println("Forum application starting...")

	// Initialize the necessary components
	cfg := config.LoadViper()
	engine := app.NewFiber(cfg)
	logger := app.NewLogger(cfg)
	validate := app.NewValidator(cfg)
	database := app.OpenPostgres(cfg, logger)

	// Boot the application with all components
	if err := boot(&schema{
		App:      engine,
		Config:   cfg,
		DB:       database,
		Log:      logger,
		Validate: validate,
	}); err != nil {
		logger.Fatalf("BOOT: initialization failed, cause: %v", err)
	}

	// Set up graceful shutdown channel
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	serverPort := cfg.GetInt("APPLICATION.PORT")
	go func() {
		logger.Infof("SERVER: listening on port %d", serverPort)
		if err := engine.Listen(fmt.Sprintf(":%d", serverPort)); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("SERVER: starting failed, cause: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-shutdownChan
	logger.Info("SERVER: shutting down gracefully...")

	// Gracefully shutdown the server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := engine.ShutdownWithContext(ctx); err != nil {
		logger.Fatalf("SERVER: graceful shutdown failed, cause: %v", err)
	}

	logger.Info("SERVER: exited gracefully")
}

// Boot initializes the application, setting up middleware and routes
func boot(s *schema) error {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // Allow all origins
		AllowMethods: "GET,POST,PUT,DELETE",                         // Allow specific methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // Allow specific headers
	}))

	// Dependency Injection
	userRepo := postgres.NewUserRepository(s.DB, s.Log)
	threadRepo := postgres.NewThreadRepository(s.DB, s.Log)
	commentRepo := postgres.NewCommentRepository(s.DB, s.Log)
	voteRepo := postgres.NewVoteRepository(s.DB, s.Log)

	// Additional boot logic, such as route setup
	router := s.App.Group("/")
	rest.NewAuthHandler(router, userRepo, s.Log, s.Validate)
	rest.NewUserHandler(router, userRepo, s.Log, s.Validate)
	rest.NewThreadHandler(router, threadRepo, commentRepo, voteRepo, s.Log, s.Validate)
	rest.NewLeaderBoardHandler(router, userRepo, threadRepo, commentRepo, s.Log)

	return nil
}

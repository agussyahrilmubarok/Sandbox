package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// serve initializes and starts the HTTP server, handling graceful shutdown on interrupt signals.
func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port), // Set server address based on configured port
		Handler:      app.routes(),                        // Use the application's routes
		IdleTimeout:  time.Minute,                         // Timeout for idle connections
		ReadTimeout:  10 * time.Second,                    // Timeout for reading request
		WriteTimeout: 30 * time.Second,                    // Timeout for writing response
	}

	shutdownError := make(chan error) // Channel to signal shutdown errors

	// Goroutine to handle graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1) // Channel to receive OS signals

		// Notify on interrupt or terminate signals
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit // Wait for a signal

		// Log the caught signal
		app.logger.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Context with timeout for shutdown
		defer cancel()

		err := srv.Shutdown(ctx) // Attempt to gracefully shutdown the server
		if err != nil {
			shutdownError <- err // Send error to shutdown error channel
		}

		// Log message about completing background tasks
		app.logger.PrintInfo("completing background tasks", map[string]string{
			"addr": srv.Addr,
		})

		app.wg.Wait()        // Wait for background tasks to finish
		shutdownError <- nil // Signal successful shutdown
	}()

	// Log that the server is starting
	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})

	err := srv.ListenAndServe()                // Start the server
	if !errors.Is(err, http.ErrServerClosed) { // Check if the error is not due to a closed server
		return err // Return the error if it's something else
	}

	err = <-shutdownError // Wait for shutdown error or completion
	if err != nil {
		return err // Return any shutdown error
	}

	// Log that the server has stopped
	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil // Return nil indicating a successful shutdown
}

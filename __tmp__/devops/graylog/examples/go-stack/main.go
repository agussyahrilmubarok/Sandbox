package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/sirupsen/logrus"
)

// Global logger
var log = logrus.New()

func main() {
	// Setup GELF UDP writer for Graylog
	writer, err := gelf.NewWriter("127.0.0.1:12201")
	if err != nil {
		fmt.Println("Failed to create GELF writer:", err)
		os.Exit(1)
	}

	// Configure logger
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(writer)
	log.SetLevel(logrus.DebugLevel)

	// Business log examples
	generateInfo()
	generateWarn()
	generateError()

	log.Println("Finished sending logs to Graylog.")
}

// Example: normal system operation
func generateInfo() {
	log.WithFields(logrus.Fields{
		"order_id":   "ORD-202501",
		"user_id":    42,
		"event_type": "order_created",
	}).Info("Order successfully created")
}

// Example: something unexpected but not fatal
func generateWarn() {
	log.WithFields(logrus.Fields{
		"user_id":     42,
		"retry_after": "2s",
		"reason":      "cache_miss",
	}).Warn("Cache miss occurred, using fallback data")
}

// Example: real error when calling an external API
func generateError() {
	log.WithFields(logrus.Fields{
		"service":     "payment_gateway",
		"endpoint":    "/charge",
		"retry_count": 3,
		"timeout":     "5s",
		"timestamp":   time.Now().Format(time.RFC3339),
	}).Error("Payment service request failed")
}

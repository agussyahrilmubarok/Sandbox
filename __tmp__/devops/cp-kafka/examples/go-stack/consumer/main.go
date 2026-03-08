package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaReader *kafka.Reader

	bookStore = struct {
		sync.RWMutex
		data []Book
	}{data: []Book{}}
)

func connectKafka() {
	brokers := []string{
		"localhost:9092",
		"localhost:9094",
		"localhost:9096",
	}

	log.Printf("connecting to kafka cluster: %v\n", brokers)

	kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    "books",
		GroupID:  "books-consumer-group",
		MinBytes: 1e3, // 1KB
		MaxBytes: 1e6, // 1MB
	})

	log.Println("kafka cluster connected")
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func startConsumer() {
	go func() {
		log.Println("starting Kafka consumer...")

		for {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			msg, err := kafkaReader.ReadMessage(ctx)
			cancel()

			if err != nil {
				log.Printf("error reading kafka message: %v\n", err)
				continue
			}

			var book Book
			if err := json.Unmarshal(msg.Value, &book); err != nil {
				log.Printf("failed parsing message: %v\n", err)
				continue
			}

			bookStore.Lock()
			bookStore.data = append(bookStore.data, book)
			bookStore.Unlock()

			log.Printf("received book => %s : %s\n", book.ID, book.Title)
		}
	}()
}

// GET /api/v1/books/search?q=keyword
func searchBook(c echo.Context) error {
	query := c.QueryParam("q")

	bookStore.RLock()
	defer bookStore.RUnlock()

	var result []Book

	if query == "" {
		result = bookStore.data
	} else {
		query = strings.ToLower(query)
		for _, b := range bookStore.data {
			if strings.Contains(strings.ToLower(b.Title), query) ||
				strings.Contains(strings.ToLower(b.Author), query) {
				result = append(result, b)
			}
		}
	}

	return c.JSONPretty(http.StatusOK, result, "  ")
}

func main() {
	connectKafka()
	
	startConsumer()

	e := echo.New()
	e.GET("/api/v1/books/search", searchBook)

	log.Println("server running at :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

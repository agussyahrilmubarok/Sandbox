package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	rabbitConn *amqp.Connection
	rabbitCh   *amqp.Channel
	queueName  = "books"

	bookStore = struct {
		sync.RWMutex
		data []Book
	}{data: []Book{}}
)

func connectRabbitMQ() {
	var err error
	rabbitConn, err = amqp.Dial("amqp://sandbox:sandboxpass@localhost:5672/")
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq: %v", err)
	}

	rabbitCh, err = rabbitConn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}

	_, err = rabbitCh.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	log.Println("connected to rabbitmq and queue declared")
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func startConsumer() {
	msgs, err := rabbitCh.Consume(
		queueName,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("failed to start consumer: %v", err)
	}

	go func() {
		log.Println("starting RabbitMQ consumer...")
		for msg := range msgs {
			var book Book
			if err := json.Unmarshal(msg.Body, &book); err != nil {
				log.Printf("failed parsing message: %v", err)
				continue
			}

			bookStore.Lock()
			bookStore.data = append(bookStore.data, book)
			bookStore.Unlock()

			log.Printf("received book => %s : %s", book.ID, book.Title)
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
	connectRabbitMQ()
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	startConsumer()

	e := echo.New()
	e.GET("/api/v1/books/search", searchBook)

	log.Println("server running at :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

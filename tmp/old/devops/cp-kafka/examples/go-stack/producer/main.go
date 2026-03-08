package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func connectKafka() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	brokers := []string{
		"localhost:9092",
		"localhost:9094",
		"localhost:9096",
	}

	log.Printf("connecting to kafka brokers: %v\n", brokers)

	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    "books",
		Balancer: &kafka.LeastBytes{},
	}

	conn, err := kafka.DialContext(ctx, "tcp", brokers[0])
	if err != nil {
		log.Fatalf("kafka connection failed: %v\n", err)
	}
	defer conn.Close()

	log.Println("connected to kafka cluster")
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func initDummyBooks() []Book {
	return []Book{
		{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan", Year: 2015},
		{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008},
		{ID: "3", Title: "Domain Driven Design", Author: "Eric Evans", Year: 2003},
		{ID: "4", Title: "Distributed Systems", Author: "Andrew Tanenbaum", Year: 2007},
	}
}

func sendBooksToKafka(books []Book) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, book := range books {
		value, _ := json.Marshal(book)

		msg := kafka.Message{
			Key:   []byte(book.ID),
			Value: value,
		}

		if err := kafkaWriter.WriteMessages(ctx, msg); err != nil {
			log.Printf("failed to send message (ID=%s): %v\n", book.ID, err)
		} else {
			log.Printf("sent book ID=%s to kafka\n", book.ID)
		}
	}
}

func main() {
	connectKafka()

	books := initDummyBooks()
	sendBooksToKafka(books)

	log.Println("all books sent to kafka topic 'books'")
}

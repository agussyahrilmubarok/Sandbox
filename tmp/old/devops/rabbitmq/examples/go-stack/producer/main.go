package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitConn *amqp.Connection
var rabbitCh *amqp.Channel
var queueName = "books"

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

func initDummyBooks() []Book {
	return []Book{
		{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan", Year: 2015},
		{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008},
		{ID: "3", Title: "Domain Driven Design", Author: "Eric Evans", Year: 2003},
		{ID: "4", Title: "Distributed Systems", Author: "Andrew Tanenbaum", Year: 2007},
	}
}

func sendBooksToRabbitMQ(books []Book) {
	for _, book := range books {
		data, _ := json.Marshal(book)
		err := rabbitCh.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
		)
		if err != nil {
			log.Printf("failed to send book %s: %v", book.ID, err)
		} else {
			log.Printf("sent book ID=%s to rabbitmq", book.ID)
		}
	}
}

func main() {
	connectRabbitMQ()
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	books := initDummyBooks()
	sendBooksToRabbitMQ(books)

	log.Println("all books sent to rabbitmq")
}

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bookCollection *mongo.Collection

func connectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := "mongodb://sandbox:sandboxpass@localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("failed to ping mongodb: %v", err)
	}

	bookCollection = client.Database("books_db").Collection("books")
	log.Println("connected to mongodb")
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func initDummyBooks() {
	books := []interface{}{
		Book{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan", Year: 2015},
		Book{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008},
		Book{ID: "3", Title: "Domain Driven Design", Author: "Eric Evans", Year: 2003},
		Book{ID: "4", Title: "Distributed Systems", Author: "Andrew Tanenbaum", Year: 2007},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := bookCollection.InsertMany(ctx, books)
	if err != nil {
		log.Printf("warning: could not insert dummy books (maybe already exist): %v", err)
	} else {
		log.Println("dummy books inserted to mongodb")
	}
}

// GET /api/v1/books/search?q=keyword
func searchBook(c echo.Context) error {
	query := c.QueryParam("q")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}
	if query != "" {
		// case-insensitive search on title or author
		filter = bson.M{
			"$or": []bson.M{
				{"title": bson.M{"$regex": query, "$options": "i"}},
				{"author": bson.M{"$regex": query, "$options": "i"}},
			},
		}
	}

	cur, err := bookCollection.Find(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer cur.Close(ctx)

	var books []Book
	if err := cur.All(ctx, &books); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSONPretty(http.StatusOK, books, "  ")
}

func main() {
	connectMongoDB()

	initDummyBooks()

	e := echo.New()
	e.GET("/api/v1/books/search", searchBook)

	log.Println("server running at :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

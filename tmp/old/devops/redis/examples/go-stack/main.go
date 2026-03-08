package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func connectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // default DB
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	log.Println("connected to redis")
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func initDummyBooks() {
	books := []Book{
		{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan", Year: 2015},
		{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008},
		{ID: "3", Title: "Domain Driven Design", Author: "Eric Evans", Year: 2003},
		{ID: "4", Title: "Distributed Systems", Author: "Andrew Tanenbaum", Year: 2007},
	}

	for _, b := range books {
		data, _ := json.Marshal(b)
		err := rdb.Set(ctx, "book:"+b.ID, data, 0).Err() // 0 = no expiration
		if err != nil {
			log.Printf("warning: could not insert book %s: %v", b.ID, err)
		}
	}

	log.Println("dummy books inserted to redis")
}

// GET /api/v1/books/search?q=keyword
func searchBook(c echo.Context) error {
	query := strings.ToLower(c.QueryParam("q"))

	keys, err := rdb.Keys(ctx, "book:*").Result()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var books []Book
	for _, key := range keys {
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			continue
		}
		var b Book
		if err := json.Unmarshal([]byte(val), &b); err != nil {
			continue
		}

		if query == "" || strings.Contains(strings.ToLower(b.Title), query) || strings.Contains(strings.ToLower(b.Author), query) {
			books = append(books, b)
		}
	}

	return c.JSONPretty(http.StatusOK, books, "  ")
}

func main() {
	connectRedis()

	initDummyBooks()

	e := echo.New()
	e.GET("/api/v1/books/search", searchBook)

	log.Println("server running at :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

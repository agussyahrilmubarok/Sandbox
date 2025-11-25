package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var db *sql.DB

func connectPostgres() {
	// DSN format: postgres://username:password@host:port/dbname?sslmode=disable
	dsn := "postgres://postgres:sandboxpass@localhost:5432/sandbox?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	log.Println("connected to postgres")
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

	// create table if not exists
	createTable := `
	CREATE TABLE IF NOT EXISTS books (
		id VARCHAR(20) PRIMARY KEY,
		title VARCHAR(255),
		author VARCHAR(255),
		year INT
	);
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	// insert dummy data (ignore duplicates)
	for _, b := range books {
		_, err := db.Exec(
			"INSERT INTO books (id, title, author, year) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING",
			b.ID, b.Title, b.Author, b.Year,
		)
		if err != nil {
			log.Printf("warning: could not insert book %s: %v", b.ID, err)
		}
	}

	log.Println("dummy books inserted to postgres")
}

// GET /api/v1/books/search?q=keyword
func searchBook(c echo.Context) error {
	query := c.QueryParam("q")
	var rows *sql.Rows
	var err error

	if query == "" {
		rows, err = db.Query("SELECT id, title, author, year FROM books")
	} else {
		q := "%" + strings.ToLower(query) + "%"
		rows, err = db.Query(
			`SELECT id, title, author, year FROM books 
			WHERE LOWER(title) LIKE $1 OR LOWER(author) LIKE $2`,
			q, q,
		)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		books = append(books, b)
	}

	return c.JSONPretty(http.StatusOK, books, "  ")
}

func main() {
	connectPostgres()

	initDummyBooks()

	e := echo.New()
	e.GET("/api/v1/books/search", searchBook)

	log.Println("server running at :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

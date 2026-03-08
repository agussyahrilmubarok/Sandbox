package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/labstack/echo/v4"
)

var es *elasticsearch.Client

func connectES() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("error creating elasticsearch client: %v", err)
	}

	res, err := client.Ping(client.Ping.WithContext(ctx))
	if err != nil {
		log.Fatalf("elasticsearch connection failed: %v", err)
	}
	defer res.Body.Close()

	es = client
	log.Println("connected to elasticsearch")
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

		res, err := es.Index(
			"books",
			bytes.NewReader(data),
			es.Index.WithDocumentID(b.ID),
			es.Index.WithContext(context.Background()),
		)
		if err != nil {
			log.Printf("error indexing book %s: %v", b.ID, err)
			continue
		}
		res.Body.Close()
	}

	log.Println("dummy books inserted into elasticsearch")
}

// GET /api/v1/books/search?q=clean
func searchBook(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "query param 'q' is required",
		})
	}

	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"title", "author"},
			},
		},
	}
	json.NewEncoder(&buf).Encode(searchQuery)

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("books"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		log.Printf("elasticsearch search error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	defer res.Body.Close()

	var r map[string]interface{}
	json.NewDecoder(res.Body).Decode(&r)

	return c.JSONPretty(http.StatusOK, r, "  ")
}

func main() {
	connectES()
	initDummyBooks()

	e := echo.New()

	e.GET("/api/v1/books/search", searchBook)

	log.Println("Server running at :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

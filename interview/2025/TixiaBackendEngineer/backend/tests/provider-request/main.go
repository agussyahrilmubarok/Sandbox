package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	_, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "flight.search.requested",
		Values: map[string]interface{}{
			"search_id": "test123",
			"from":      "C",
			"to":        "DPS",
			"date":      "2025-07-10",
		},
	}).Result()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message sent!")
}

package repository

import (
	"context"
	"encoding/json"

	"example.com/main-service/internal/domain"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

//go:generate mockery --name=IFlightRepository
type IFlightRepository interface {
	PublishSearchRequest(ctx context.Context, req domain.FlightSearchRequest) error
	ConsumeSearchResults(ctx context.Context, searchID string) <-chan domain.FlightSearchResult
}

type flightRepository struct {
	rdb *redis.Client
	log *zap.Logger
}

func NewFlightRepository(rdb *redis.Client, log *zap.Logger) IFlightRepository {
	return &flightRepository{
		rdb: rdb,
		log: log,
	}
}

func (r *flightRepository) PublishSearchRequest(ctx context.Context, req domain.FlightSearchRequest) error {
	err := r.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: domain.StreamFlightSearchRequested,
		Values: map[string]interface{}{
			"search_id": req.SearchID,
			"from":      req.From,
			"to":        req.To,
			"date":      req.Date,
		},
	}).Err()

	if err != nil {
		r.log.Error("Failed to publish search request", zap.Error(err))
		return err
	}

	r.log.Info("Published search request to Redis stream",
		zap.String("stream", domain.StreamFlightSearchRequested),
		zap.String("search_id", req.SearchID),
	)
	return nil
}

func (r *flightRepository) ConsumeSearchResults(ctx context.Context, searchID string) <-chan domain.FlightSearchResult {
	results := make(chan domain.FlightSearchResult)

	go func() {
		defer close(results)

		lastID := "0"

		for {
			streams, err := r.rdb.XRead(ctx, &redis.XReadArgs{
				Streams: []string{domain.StreamFlightSearchResults, lastID},
				Block:   0,
			}).Result()
			if err != nil {
				if err == redis.Nil {
					continue
				}
				r.log.Error("Error reading from Redis stream", zap.Error(err))
				return
			}

			for _, stream := range streams {
				for _, msg := range stream.Messages {
					var res domain.FlightSearchResult
					if raw, ok := msg.Values["data"].(string); ok {
						if err := json.Unmarshal([]byte(raw), &res); err != nil {
							r.log.Error("Failed to unmarshal search result", zap.Error(err))
							continue
						}
						if res.SearchID == searchID {
							results <- res
						}
					}
					lastID = msg.ID
				}
			}
		}
	}()

	return results
}

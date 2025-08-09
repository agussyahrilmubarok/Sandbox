package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"example.com/provider-service/internal/domain"
	"example.com/provider-service/internal/repository"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	StreamFlightGroup = "flight_group"
	StreamFlightApp   = "flight_app"
)

type FlightSearchConsumer struct {
	repo repository.IFlightRepository
	rdb  *redis.Client
	log  *zap.Logger
}

func NewFlightSearchConsumer(repo repository.IFlightRepository, rdb *redis.Client, logger *zap.Logger) *FlightSearchConsumer {
	return &FlightSearchConsumer{
		repo: repo,
		rdb:  rdb,
		log:  logger,
	}
}

func (c *FlightSearchConsumer) Start(ctx context.Context) error {
	c.log.Info("Starting FlightSearchConsumer...")

	if err := c.rdb.XGroupCreateMkStream(ctx, domain.StreamFlightSearchRequested, StreamFlightGroup, "$").Err(); err != nil {
		if !strings.Contains(err.Error(), "BUSYGROUP") {
			c.log.Error("Failed to create consumer group", zap.Error(err))
			return fmt.Errorf("failed to create consumer group: %w", err)
		}
	}

	for {
		select {
		case <-ctx.Done():
			c.log.Info("FlightSearchConsumer received shutdown signal, exiting loop...")
			return nil
		default:
			messages, err := c.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    StreamFlightGroup,
				Consumer: StreamFlightApp,
				Streams:  []string{domain.StreamFlightSearchRequested, ">"},
				Count:    1,
				Block:    5 * time.Second,
			}).Result()

			if err != nil && err != redis.Nil {
				c.log.Error("Error reading from stream", zap.Error(err))
				continue
			}

			if len(messages) == 0 || len(messages[0].Messages) == 0 {
				continue
			}

			for _, m := range messages[0].Messages {
				c.log.Info("Processing message", zap.String("id", m.ID))
				c.ProcessMessage(ctx, m.ID, m.Values)

				if err := c.rdb.XAck(ctx, domain.StreamFlightSearchRequested, StreamFlightGroup, m.ID).Err(); err != nil {
					c.log.Error("Failed to ack message", zap.String("id", m.ID), zap.Error(err))
				} else {
					c.log.Info("Message acknowledged", zap.String("id", m.ID))
				}
			}
		}
	}
}

func (c *FlightSearchConsumer) Stop(ctx context.Context) {
	c.log.Info("Closing Redis connection for FlightSearchConsumer...")
	if err := c.rdb.Close(); err != nil {
		c.log.Error("Error closing Redis connection", zap.Error(err))
	} else {
		c.log.Info("Redis connection closed")
	}
}

func (c *FlightSearchConsumer) ProcessMessage(ctx context.Context, msgID string, values map[string]interface{}) {
	var req domain.FlightSearchRequest
	data, err := json.Marshal(values)
	if err != nil {
		c.log.Error("Failed to marshal message values", zap.Error(err))
		return
	}

	if err := json.Unmarshal(data, &req); err != nil {
		c.log.Error("Failed to unmarshal message to FlightSearchRequest", zap.Error(err))
		return
	}

	c.log.Info("Searching flights",
		zap.String("search_id", req.SearchID),
		zap.String("from", req.From),
		zap.String("to", req.To),
		zap.String("date", req.Date),
	)

	flights, err := c.repo.GetAllFlights()
	if err != nil {
		c.log.Error("Failed to get flights", zap.Error(err))
		return
	}

	var results []domain.Flight
	for _, f := range flights {
		if strings.EqualFold(f.From, req.From) &&
			strings.EqualFold(f.To, req.To) &&
			strings.HasPrefix(f.DepartureTime, req.Date) {
			results = append(results, f)
		}
	}

	c.log.Info("Flights found", zap.Int("count", len(results)), zap.String("search_id", req.SearchID))

	var resultMsg map[string]interface{}
	if len(results) == 0 {
		resultMsg = map[string]interface{}{
			"search_id": req.SearchID,
			"status":    "not_found",
			"results":   []interface{}{},
		}
	} else {
		resultMsg = map[string]interface{}{
			"search_id": req.SearchID,
			"status":    "completed",
			"results":   results,
		}
	}

	resultData, _ := json.Marshal(resultMsg)
	if err := c.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: domain.StreamFlightSearchResults,
		Values: map[string]interface{}{
			"data": string(resultData),
		},
	}).Err(); err != nil {
		c.log.Error("Failed to publish results", zap.String("search_id", req.SearchID), zap.Error(err))
	} else {
		c.log.Info("Published flight search results", zap.String("search_id", req.SearchID))
	}
}

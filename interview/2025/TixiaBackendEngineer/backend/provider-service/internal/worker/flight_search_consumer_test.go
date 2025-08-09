package worker_test

import (
	"context"
	"encoding/json"
	"testing"

	"example.com/provider-service/internal/domain"
	"example.com/provider-service/internal/worker"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockFlightRepo struct {
	mock.Mock
}

func (m *MockFlightRepo) GetAllFlights() ([]domain.Flight, error) {
	args := m.Called()
	return args.Get(0).([]domain.Flight), args.Error(1)
}

func TestProcessMessage_FindsFlights(t *testing.T) {
	logger := zap.NewNop()
	db, mockRedis := redismock.NewClientMock()

	mockRepo := new(MockFlightRepo)
	mockFlights := []domain.Flight{
		{ID: "1", From: "JKT", To: "DPS", DepartureTime: "2025-08-15T08:00", Available: true},
		{ID: "2", From: "SUB", To: "DPS", DepartureTime: "2025-08-15T10:00", Available: true},
	}
	mockRepo.On("GetAllFlights").Return(mockFlights, nil)

	consumer := worker.NewFlightSearchConsumer(mockRepo, db, logger)

	req := domain.FlightSearchRequest{
		SearchID: "abc123",
		From:     "JKT",
		To:       "DPS",
		Date:     "2025-08-15",
	}
	values := make(map[string]interface{})
	b, _ := json.Marshal(req)
	_ = json.Unmarshal(b, &values)

	// Expect Redis XAdd to be called
	resultMsg := map[string]interface{}{
		"search_id": req.SearchID,
		"status":    "completed",
		"results":   []domain.Flight{mockFlights[0]},
	}
	resultData, _ := json.Marshal(resultMsg)
	mockRedis.ExpectXAdd(&redis.XAddArgs{
		Stream: "flight.search.results",
		Values: map[string]interface{}{
			"data": string(resultData),
		},
	}).SetVal("123-0")

	consumerTest := consumer
	consumerTest.ProcessMessage(context.Background(), "1-0", values)

	assert.NoError(t, mockRedis.ExpectationsWereMet())
	mockRepo.AssertExpectations(t)
}

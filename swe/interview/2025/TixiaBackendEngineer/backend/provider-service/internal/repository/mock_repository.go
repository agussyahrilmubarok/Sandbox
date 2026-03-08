package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"example.com/provider-service/internal/domain"
	"go.uber.org/zap"
)

//go:generate mockery --name=IFlightRepository
type IFlightRepository interface {
	GetAllFlights() ([]domain.Flight, error)
}

type flightRepository struct {
	filePath string
	log      *zap.Logger
}

func NewFlightRepository(filePath string, logger *zap.Logger) IFlightRepository {
	return &flightRepository{
		filePath: filePath,
		log:      logger,
	}
}

func (r *flightRepository) GetAllFlights() ([]domain.Flight, error) {
	r.log.Debug("Opening JSON file", zap.String("path", r.filePath))

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		r.log.Error("Failed to read file", zap.Error(err))
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var flights []domain.Flight

	if err := json.Unmarshal(data, &flights); err == nil {
		r.log.Info("Loaded flights from array JSON", zap.Int("count", len(flights)))
		return flights, nil
	}

	var flightMap map[string]domain.Flight
	if err := json.Unmarshal(data, &flightMap); err != nil {
		r.log.Error("Failed to parse JSON as map", zap.Error(err))
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	for _, f := range flightMap {
		flights = append(flights, f)
	}

	r.log.Info("Loaded flights from map JSON", zap.Int("count", len(flights)))
	return flights, nil
}

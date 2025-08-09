package usecase

import (
	"context"

	"example.com/main-service/internal/domain"
	"example.com/main-service/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

//go:generate mockery --name=IFlightUseCase
type IFlightUseCase interface {
	SearchFlights(ctx context.Context, req domain.CreateSearchBody) (string, error)
	StreamResults(ctx context.Context, searchID string) <-chan domain.FlightSearchResult
}

type flightUseCase struct {
	repo repository.IFlightRepository
	log  *zap.Logger
}

func NewFlightUseCase(repo repository.IFlightRepository, log *zap.Logger) IFlightUseCase {
	return &flightUseCase{
		repo: repo,
		log:  log,
	}
}

func (uc *flightUseCase) SearchFlights(ctx context.Context, req domain.CreateSearchBody) (string, error) {
	searchID := uuid.New().String()

	searchReq := domain.FlightSearchRequest{
		SearchID:   searchID,
		From:       req.From,
		To:         req.To,
		Date:       req.Date,
	}

	if err := uc.repo.PublishSearchRequest(ctx, searchReq); err != nil {
		uc.log.Error("Failed to publish flight search request", zap.Error(err))
		return "", err
	}

	uc.log.Info("Flight search request submitted",
		zap.String("search_id", searchID),
		zap.String("from", req.From),
		zap.String("to", req.To),
		zap.String("date", req.Date),
		zap.Int("passengers", req.Passengers),
	)

	return searchID, nil
}

func (uc *flightUseCase) StreamResults(ctx context.Context, searchID string) <-chan domain.FlightSearchResult {
	return uc.repo.ConsumeSearchResults(ctx, searchID)
}

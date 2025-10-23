package course

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Handler struct {
	service IService
	log     zerolog.Logger
}

func NewHandler(service IService, log zerolog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log.With().Str("component", "course_handler").Logger(),
	}
}

// InitDummy godoc
// @Summary Initialize dummy course data
// @Description Create dummy data for testing
// @Tags Dummy
// @Produce json
// @Success 200 {array} Course
// @Router /courses/init-dummy [post]
func (h *Handler) InitDummy(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	dummies := []*Course{
		{
			ID:            uuid.New().String(),
			Code:          "C-001",
			Name:          "Go Programming Basics",
			Price:         199.99,
			StartDate:     time.Now(),                         // Starts now
			EndDate:       time.Now().Add(7 * 24 * time.Hour), // Ends in 7 days
			SeatAvailable: 30,
		},
		{
			ID:            uuid.New().String(),
			Code:          "C-002",
			Name:          "Advanced Golang",
			Price:         249.99,
			StartDate:     time.Now().Add(1 * time.Hour),      // Starts now + 1 hour
			EndDate:       time.Now().Add(8 * 24 * time.Hour), // Ends in 8 days
			SeatAvailable: 25,
		},
		{
			ID:            uuid.New().String(),
			Code:          "C-003",
			Name:          "Docker for Beginners",
			Price:         149.99,
			StartDate:     time.Now().Add(-24 * time.Hour),    // Started yesterday
			EndDate:       time.Now().Add(3 * 24 * time.Hour), // Ends in 3 days
			SeatAvailable: 20,
		},
		{
			ID:            uuid.New().String(),
			Code:          "C-004",
			Name:          "Kubernetes Mastery",
			Price:         299.99,
			StartDate:     time.Now().Add(-7 * 24 * time.Hour), // Started 7 days ago
			EndDate:       time.Now().Add(-1 * 24 * time.Hour), // Ended yesterday
			SeatAvailable: 15,
		},
		{
			ID:            uuid.New().String(),
			Code:          "C-005",
			Name:          "Machine Learning Fundamentals",
			Price:         399.99,
			StartDate:     time.Now().Add(-2 * 24 * time.Hour), // Started 2 days ago
			EndDate:       time.Now().Add(5 * 24 * time.Hour),  // Ends in 5 days
			SeatAvailable: 50,
		},
	}

	for _, course := range dummies {
		if err := h.service.Save(ctx, course); err != nil {
			h.log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to insert dummy course")
		}
	}

	h.log.Info().Int("courses_count", len(dummies)).Msg("Dummy data initialized")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummies)
}

// CleanDummy godoc
// @Summary Remove all dummy courses
// @Description Delete all dummy course data
// @Tags Dummy
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/clean-dummy [delete]
func (h *Handler) CleanDummy(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	courses, err := h.service.FindAll(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}

	for _, course := range courses {
		if err := h.service.DeleteByID(ctx, course.ID); err != nil {
			h.log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to delete dummy course")
		}
	}

	h.log.Info().Msg("Dummy courses cleaned")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Dummy courses cleaned"})
}

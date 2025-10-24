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

// FindAll godoc
// @Summary Get all courses
// @Description Retrieves all course data
// @Tags Courses
// @Produce json
// @Success 200 {array} Course
// @Failure 500 {object} map[string]string
// @Router /courses [get]
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	courses, err := h.service.FindAll(ctx)
	if err != nil {
		h.log.Error().Err(err).Msg("Failed to fetch courses")
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}

	h.log.Info().Int("courses_count", len(courses)).Msg("Fetched all courses successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// FindByCode godoc
// @Summary Get course by Code
// @Description Retrieves course data by code
// @Tags Courses
// @Produce json
// @Param code query string true "Course Code"
// @Success 200 {object} Course
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/find [get]
func (h *Handler) FindByCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	courseCode := r.URL.Query().Get("code")
	if courseCode == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}

	course, err := h.service.FindByCode(ctx, courseCode)
	if err != nil {
		h.log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to fetch course")
		http.Error(w, "Failed to fetch course", http.StatusInternalServerError)
		return
	}

	if course == nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	h.log.Info().Str("course_code", courseCode).Msg("Fetched course successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

// ReserveCourse godoc
// @Summary Reserve a seat for a course by course code
// @Description Reserve a seat for a course by specifying its course code
// @Tags Courses
// @Accept json
// @Produce json
// @Param course_code body CourseCodeRequest true "Course Code"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/reserve [post]
func (h *Handler) ReserveCourse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var payload CourseCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.log.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.service.ReserveByCode(ctx, payload.Code)
	if err != nil {
		h.log.Error().Err(err).Str("course_code", payload.Code).Msg("Failed to reserve course")
		switch err.Error() {
		case "course has already ended":
			http.Error(w, "Course has already ended", http.StatusBadRequest)
		case "no available seats to reserve":
			http.Error(w, "No available seats", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to reserve course", http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("course_code", payload.Code).Msg("Course seat reserved successfully")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Seat reserved successfully"})
}

// ReleaseCourse godoc
// @Summary Release a seat for a course by course code
// @Description Release a seat for a course by specifying its course code
// @Tags Courses
// @Accept json
// @Produce json
// @Param course_code body CourseCodeRequest true "Course Code"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/release [post]
func (h *Handler) ReleaseCourse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var payload CourseCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.log.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.service.ReleaseByCode(ctx, payload.Code)
	if err != nil {
		h.log.Error().Err(err).Str("course_code", payload.Code).Msg("Failed to release course")
		switch err.Error() {
		case "course has already ended":
			http.Error(w, "Course has already ended", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to release course", http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("course_code", payload.Code).Msg("Course seat released successfully")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Seat released successfully"})
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

package course

import (
	"net/http"
	"time"

	"github.com/agussyahrilmubarok/gohelp/exception"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

type Handler struct {
	service IService
	tracer  trace.Tracer
}

func NewHandler(service IService, tracer trace.Tracer) *Handler {
	return &Handler{
		service: service,
		tracer:  tracer,
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
func (h *Handler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := h.tracer.Start(ctx, "handler.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.FindAll").Logger()

	courses, err := h.service.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("Failed to fetch courses")
		if httpErr, ok := err.(*exception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch courses"})
	}

	log.Info().Int("courses_count", len(courses)).Msg("Fetched all courses successfully")
	return c.JSON(http.StatusOK, courses)
}

// Find godoc
// @Summary Get course by query param
// @Description Retrieves course data by query param
// @Tags Courses
// @Produce json
// @Param code query string true "Course Code"
// @Success 200 {object} Course
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/find [get]
func (h *Handler) Find(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := h.tracer.Start(ctx, "handler.Find")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.Find").Logger()

	courseCode := c.QueryParam("code")
	if courseCode == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing code parameter"})
	}

	course, err := h.service.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to fetch course")
		if httpErr, ok := err.(*exception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch course"})
	}

	log.Info().Str("course_code", courseCode).Msg("Fetched course successfully")
	return c.JSON(http.StatusOK, course)
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
func (h *Handler) ReserveCourse(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := h.tracer.Start(ctx, "handler.ReserveCourse")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.ReserveCourse").Logger()

	var payload CourseCodeRequest
	if err := c.Bind(&payload); err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("Failed to decode request body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	if err := h.service.ReserveByCode(ctx, payload.Code); err != nil {
		span.RecordError(err)
		log.Error().Err(err).Str("course_code", payload.Code).Msg("Failed to reserve course")
		if httpErr, ok := err.(*exception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to reserve course"})
	}

	log.Info().Str("course_code", payload.Code).Msg("Course seat reserved successfully")
	return c.JSON(http.StatusOK, echo.Map{"message": "Seat reserved successfully"})
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
func (h *Handler) ReleaseCourse(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := h.tracer.Start(ctx, "handler.ReleaseCourse")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.ReleaseCourse").Logger()

	var payload CourseCodeRequest
	if err := c.Bind(&payload); err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("Failed to decode request body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	err := h.service.ReleaseByCode(ctx, payload.Code)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Str("course_code", payload.Code).Msg("Failed to release course")
		if httpErr, ok := err.(*exception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to release course"})
	}

	log.Info().Str("course_code", payload.Code).Msg("Course seat released successfully")
	return c.JSON(http.StatusOK, echo.Map{"message": "Seat released successfully"})
}

// InitDummy godoc
// @Summary Initialize dummy course data
// @Description Create dummy data for testing
// @Tags Dummy
// @Produce json
// @Success 200 {array} Course
// @Router /courses/init-dummy [post]
func (h *Handler) InitDummy(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := h.tracer.Start(ctx, "handler.InitDummy")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.InitDummy").Logger()

	dummies := []Course{
		// available start_date, end_date and seat.
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-001",
			Name:          "Go Programming Basics",
			Price:         199.99,
			StartDate:     time.Now(),                         // Starts now
			EndDate:       time.Now().Add(7 * 24 * time.Hour), // Ends in 7 days
			SeatAvailable: 10,
		},
		// available start_date, end_date, but not seat.
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-002",
			Name:          "Advanced Golang",
			Price:         249.99,
			StartDate:     time.Now().Add(1 * time.Hour),      // Starts now + 1 hour
			EndDate:       time.Now().Add(8 * 24 * time.Hour), // Ends in 8 days
			SeatAvailable: 0,
		},
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-003",
			Name:          "Docker for Beginners",
			Price:         149.99,
			StartDate:     time.Now().Add(-24 * time.Hour),    // Started yesterday
			EndDate:       time.Now().Add(3 * 24 * time.Hour), // Ends in 3 days
			SeatAvailable: 20,
		},
		// not available start_date, end_date, but available seat.
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-004",
			Name:          "Kubernetes Mastery",
			Price:         299.99,
			StartDate:     time.Now().Add(-7 * 24 * time.Hour), // Started 7 days ago
			EndDate:       time.Now().Add(-1 * 24 * time.Hour), // Ended yesterday
			SeatAvailable: 15,
		},
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-005",
			Name:          "Machine Learning Fundamentals",
			Price:         399.99,
			StartDate:     time.Now().Add(-2 * 24 * time.Hour), // Started 2 days ago
			EndDate:       time.Now().Add(5 * 24 * time.Hour),  // Ends in 5 days
			SeatAvailable: 50,
		},
	}

	for _, course := range dummies {
		if err := h.service.Save(ctx, &course); err != nil {
			span.RecordError(err)
			log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to insert dummy course")
			if httpErr, ok := err.(*exception.Http); ok {
				return c.JSON(httpErr.Code, map[string]string{
					"error": httpErr.Message,
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to insert dummy course"})
		}
	}

	log.Info().Int("courses_count", len(dummies)).Msg("Dummy data initialized")
	return c.JSON(http.StatusOK, dummies)
}

// CleanDummy godoc
// @Summary Remove all dummy courses
// @Description Delete all dummy course data
// @Tags Dummy
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/clean-dummy [delete]
func (h *Handler) CleanDummy(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := h.tracer.Start(ctx, "handler.CleanDummy")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.CleanDummy").Logger()

	courses, err := h.service.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("Failed to fetch courses")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch courses"})
	}

	for _, course := range courses {
		if err := h.service.DeleteByID(ctx, course.ID); err != nil {
			span.RecordError(err)
			log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to delete dummy course")
			if httpErr, ok := err.(*exception.Http); ok {
				return c.JSON(httpErr.Code, map[string]string{
					"error": httpErr.Message,
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete dummy course"})
		}
	}

	log.Info().Msg("Dummy courses cleaned")
	return c.JSON(http.StatusOK, echo.Map{"message": "Dummy courses cleaned"})
}

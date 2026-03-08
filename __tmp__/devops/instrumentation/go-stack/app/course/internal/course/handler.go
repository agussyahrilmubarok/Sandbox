package course

import (
	"net/http"
	"time"

	"example.com/course/internal/tracing"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
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
	ctx, span := tracing.StartSpan(ctx, "handler.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.FindAll").Logger()

	courses, err := h.service.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("Failed to fetch courses")
		if httpErr, ok := err.(*xexception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch courses"})
	}

	log.Info().Int("courses_count", len(courses)).Msg("fetched all courses successfully")
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
	ctx, span := tracing.StartSpan(ctx, "handler.Find")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.Find").Logger()

	courseCode := c.QueryParam("code")
	if courseCode == "" {
		log.Warn().Msg("missing code parameter")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "missing code parameter"})
	}

	course, err := h.service.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to fetch course")
		if httpErr, ok := err.(*xexception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch course"})
	}

	log.Info().Str("course_code", courseCode).Msg("fetched course successfully")
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
	ctx, span := tracing.StartSpan(ctx, "handler.ReserveCourse")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.ReserveCourse").Logger()

	var payload CourseCodeRequest
	if err := c.Bind(&payload); err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("failed to decode request body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request payload"})
	}

	if err := h.service.ReserveByCode(ctx, payload.Code); err != nil {
		span.RecordError(err)
		log.Error().Err(err).Str("course_code", payload.Code).Msg("failed to reserve course")
		if httpErr, ok := err.(*xexception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to reserve course"})
	}

	log.Info().Str("course_code", payload.Code).Msg("course seat reserved successfully")
	return c.JSON(http.StatusOK, echo.Map{"message": "seat reserved successfully"})
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
	ctx, span := tracing.StartSpan(ctx, "handler.ReleaseCourse")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.ReleaseCourse").Logger()

	var payload CourseCodeRequest
	if err := c.Bind(&payload); err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("failed to decode request body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request payload"})
	}

	err := h.service.ReleaseByCode(ctx, payload.Code)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Str("course_code", payload.Code).Msg("failed to release course")
		if httpErr, ok := err.(*xexception.Http); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"error": httpErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to release course"})
	}

	log.Info().Str("course_code", payload.Code).Msg("course seat released successfully")
	return c.JSON(http.StatusOK, echo.Map{"message": "seat released successfully"})
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
	ctx, span := tracing.StartSpan(ctx, "handler.InitDummy")
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
			Seat:          10,
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
			Seat:          0,
			SeatAvailable: 0,
		},
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-003",
			Name:          "Docker for Beginners",
			Price:         149.99,
			StartDate:     time.Now().Add(-24 * time.Hour),    // Started yesterday
			EndDate:       time.Now().Add(3 * 24 * time.Hour), // Ends in 3 days
			Seat:          20,
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
			Seat:          15,
			SeatAvailable: 15,
		},
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-005",
			Name:          "Machine Learning Fundamentals",
			Price:         399.99,
			StartDate:     time.Now().Add(-2 * 24 * time.Hour), // Started 2 days ago
			EndDate:       time.Now().Add(5 * 24 * time.Hour),  // Ends in 5 days
			Seat:          50,
			SeatAvailable: 50,
		},
	}

	for _, course := range dummies {
		if _, err := h.service.Save(ctx, &course); err != nil {
			span.RecordError(err)
			log.Error().Err(err).Str("course_id", course.ID).Msg("failed to insert dummy course")
			if httpErr, ok := err.(*xexception.Http); ok {
				return c.JSON(httpErr.Code, map[string]string{
					"error": httpErr.Message,
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to insert dummy course"})
		}
	}

	log.Info().Int("courses_count", len(dummies)).Msg("dummy data initialized")
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
	ctx, span := tracing.StartSpan(ctx, "handler.CleanDummy")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "handler.CleanDummy").Logger()

	courses, err := h.service.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		log.Error().Err(err).Msg("failed to fetch courses")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch courses"})
	}

	for _, course := range courses {
		if err := h.service.DeleteByID(ctx, course.ID); err != nil {
			span.RecordError(err)
			log.Error().Err(err).Str("course_id", course.ID).Msg("failed to delete dummy course")
			if httpErr, ok := err.(*xexception.Http); ok {
				return c.JSON(httpErr.Code, map[string]string{
					"error": httpErr.Message,
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete dummy course"})
		}
	}

	log.Info().Msg("dummy courses cleaned")
	return c.JSON(http.StatusOK, echo.Map{"message": "dummy courses cleaned"})
}

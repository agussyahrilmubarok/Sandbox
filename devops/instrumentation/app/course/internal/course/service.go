package course

import (
	"context"
	"time"

	"example.com/course/internal/metrics"
	"github.com/agussyahrilmubarok/gohelp/exception"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

//go:generate mockery --name=IService
type IService interface {
	FindAll(ctx context.Context) ([]Course, error)
	FindByID(ctx context.Context, courseID string) (*Course, error)
	FindByCode(ctx context.Context, courseCode string) (*Course, error)
	Save(ctx context.Context, course *Course) error
	DeleteByID(ctx context.Context, courseID string) error
	ReserveByCode(ctx context.Context, courseCode string) error
	ReleaseByCode(ctx context.Context, courseCode string) error
}

type service struct {
	store  IStore
	tracer trace.Tracer
}

func NewService(store IStore, tracer trace.Tracer) IService {
	return &service{
		store:  store,
		tracer: tracer,
	}
}

func (s *service) FindAll(ctx context.Context) ([]Course, error) {
	ctx, span := s.tracer.Start(ctx, "service.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindAll").Logger()

	courses, err := s.store.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("Failed to fetch all courses from store")
		return nil, exception.NewHTTPNotFound("Failed to fetch all courses", err)
	}

	metrics.TotalCourses.Set(float64(len(courses)))

	availableCount := 0
	for _, course := range courses {
		if course.SeatAvailable > 0 {
			availableCount++
		}
	}

	metrics.AvailableCourses.Set(float64(availableCount))
	metrics.ReservedCourses.Set(float64(len(courses) - availableCount))

	log.Info().Int("courses_count", len(courses)).Msg("Successfully fetched all courses")
	return courses, nil
}

func (s *service) FindByID(ctx context.Context, courseID string) (*Course, error) {
	ctx, span := s.tracer.Start(ctx, "service.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByID").Logger()

	course, err := s.store.FindByID(ctx, courseID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("Failed to find course by ID")
		return nil, exception.NewHTTPNotFound("Course not found", err)
	}

	if course == nil {
		log.Warn().Str("course_id", courseID).Msg("Course not found")
		return nil, exception.NewHTTPNotFound("Course not found", nil)
	}

	log.Info().Str("course_id", courseID).Msg("Successfully found course by ID")
	return course, nil
}

func (s *service) FindByCode(ctx context.Context, courseCode string) (*Course, error) {
	ctx, span := s.tracer.Start(ctx, "service.FindByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByCode").Logger()

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to find course by code")
		return nil, exception.NewHTTPNotFound("Course not found", err)
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("Course not found")
		return nil, exception.NewHTTPNotFound("Course not found", nil)
	}

	log.Info().Str("course_code", courseCode).Msg("Successfully found course by code")
	return course, nil
}

func (s *service) Save(ctx context.Context, course *Course) error {
	ctx, span := s.tracer.Start(ctx, "service.Save")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", course.ID), attribute.String("course.code", course.Code))

	log := zerolog.Ctx(ctx).With().Str("component", "service.Save").Logger()

	err := s.store.Save(ctx, course)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to save course")
		return err
	}

	metrics.TotalCourses.Inc()

	log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Msg("Successfully saved course")
	return nil
}

func (s *service) DeleteByID(ctx context.Context, courseID string) error {
	ctx, span := s.tracer.Start(ctx, "service.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.DeleteByID").Logger()

	err := s.store.DeleteByID(ctx, courseID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("Failed to delete course")
		return err
	}

	metrics.TotalCourses.Dec()

	log.Info().Str("course_id", courseID).Msg("Successfully deleted course")
	return nil
}

func (s *service) ReserveByCode(ctx context.Context, courseCode string) error {
	ctx, span := s.tracer.Start(ctx, "service.ReserveByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.ReserveByCode").Logger()

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to find course by code")
		return exception.NewHTTPNotFound("Course not found", err)
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("Course not found")
		return exception.NewHTTPNotFound("Course not found", nil)
	}

	currentTime := time.Now()
	if currentTime.After(course.EndDate) {
		err := exception.NewHTTPBadRequest("Course has already ended", nil)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Warn().Str("course_code", courseCode).Msg("Course has already ended, cannot reserve seat")
		return err
	}

	if course.SeatAvailable <= 0 {
		err := exception.NewHTTPBadRequest("No available seats to reserve", nil)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Warn().Str("course_code", courseCode).Msg("No available seats to reserve")
		return err
	}

	course.SeatAvailable--
	if err := s.store.Save(ctx, course); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to update course availability")
		return err
	}

	metrics.ReservedCourses.Inc()

	log.Info().Str("course_code", courseCode).Int("remaining_seats", course.SeatAvailable).Msg("Course reserved successfully")
	return nil
}

func (s *service) ReleaseByCode(ctx context.Context, courseCode string) error {
	ctx, span := s.tracer.Start(ctx, "service.ReleaseByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.ReleaseByCode").Logger()

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to find course by code")
		return exception.NewHTTPNotFound("Course not found", err)
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("Course not found")
		return exception.NewHTTPNotFound("Course not found", nil)
	}

	currentTime := time.Now()
	if currentTime.After(course.EndDate) {
		err := exception.NewHTTPBadRequest("Course has already ended", nil)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Warn().Str("course_code", courseCode).Msg("Course has already ended, cannot release seat")
		return err
	}

	course.SeatAvailable++
	if err := s.store.Save(ctx, course); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to update course availability")
		return err
	}

	metrics.ReservedCourses.Dec()

	log.Info().Str("course_code", courseCode).Int("remaining_seats", course.SeatAvailable).Msg("Seat released successfully")
	return nil
}

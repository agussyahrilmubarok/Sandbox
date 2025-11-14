package course

import (
	"context"
	"time"

	"example.com/course/internal/metrics"
	"example.com/course/internal/tracing"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

//go:generate mockery --name=IService
type IService interface {
	FindAll(ctx context.Context) ([]Course, error)
	FindByID(ctx context.Context, courseID string) (*Course, error)
	FindByCode(ctx context.Context, courseCode string) (*Course, error)
	Save(ctx context.Context, course *Course) (*Course, error)
	DeleteByID(ctx context.Context, courseID string) error
	ReserveByCode(ctx context.Context, courseCode string) error
	ReleaseByCode(ctx context.Context, courseCode string) error
}

type service struct {
	store IStore
}

func NewService(store IStore) IService {
	return &service{
		store: store,
	}
}

func (s *service) FindAll(ctx context.Context) ([]Course, error) {
	ctx, span := tracing.StartSpan(ctx, "service.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindAll").Logger()

	courses, err := s.store.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("failed to fetch all courses from store")
		return nil, xexception.NewHTTPNotFound("failed to fetch all courses", err)
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

	log.Info().Int("courses_count", len(courses)).Msg("successfully fetched all courses")
	return courses, nil
}

func (s *service) FindByID(ctx context.Context, courseID string) (*Course, error) {
	ctx, span := tracing.StartSpan(ctx, "service.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByID").Logger()

	course, err := s.store.FindByID(ctx, courseID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("failed to find course by ID")
		return nil, xexception.NewHTTPNotFound("course not found by id", err)
	}

	if course == nil {
		log.Warn().Str("course_id", courseID).Msg("Course not found")
		return nil, xexception.NewHTTPNotFound("course not found", nil)
	}

	log.Info().Str("course_id", courseID).Msg("successfully found course by ID")
	return course, nil
}

func (s *service) FindByCode(ctx context.Context, courseCode string) (*Course, error) {
	ctx, span := tracing.StartSpan(ctx, "service.FindByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByCode").Logger()

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to find course by code")
		return nil, xexception.NewHTTPNotFound("course not found by code", err)
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("course not found")
		return nil, xexception.NewHTTPNotFound("course not found", nil)
	}

	log.Info().Str("course_code", courseCode).Msg("successfully found course by code")
	return course, nil
}

func (s *service) Save(ctx context.Context, course *Course) (*Course, error) {
	ctx, span := tracing.StartSpan(ctx, "service.Save")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", course.ID), attribute.String("course.code", course.Code))

	log := zerolog.Ctx(ctx).With().Str("component", "service.Save").Logger()

	if err := s.store.Save(ctx, course); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", course.ID).Msg("failed to save course")
		return nil, err
	}

	metrics.TotalCourses.Inc()

	log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Msg("successfully saved course")
	return course, nil
}

func (s *service) DeleteByID(ctx context.Context, courseID string) error {
	ctx, span := tracing.StartSpan(ctx, "service.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.DeleteByID").Logger()

	err := s.store.DeleteByID(ctx, courseID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("failed to delete course by id")
		return err
	}

	metrics.TotalCourses.Dec()

	log.Info().Str("course_id", courseID).Msg("successfully deleted course")
	return nil
}

func (s *service) ReserveByCode(ctx context.Context, courseCode string) error {
	ctx, span := tracing.StartSpan(ctx, "service.ReserveByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.ReserveByCode").Logger()

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to find course by code")
		return xexception.NewHTTPNotFound("course not found by code", err)
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("course not found")
		return xexception.NewHTTPNotFound("course not found", nil)
	}

	currentTime := time.Now()
	if currentTime.After(course.EndDate) {
		err := xexception.NewHTTPBadRequest("course has already ended", nil)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Warn().Str("course_code", courseCode).Msg("course has already ended, cannot reserve seat")
		return err
	}

	if course.SeatAvailable <= 0 {
		err := xexception.NewHTTPBadRequest("no available seats to reserve", nil)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Warn().Str("course_code", courseCode).Msg("no available seats to reserve")
		return err
	}

	course.SeatAvailable--
	if err := s.store.Save(ctx, course); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to update course availability")
		return err
	}

	metrics.ReservedCourses.Inc()

	log.Info().Str("course_code", courseCode).Int("remaining_seats", course.SeatAvailable).Msg("course reserved successfully")
	return nil
}

func (s *service) ReleaseByCode(ctx context.Context, courseCode string) error {
	ctx, span := tracing.StartSpan(ctx, "service.ReleaseByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.ReleaseByCode").Logger()

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to find course by code")
		return xexception.NewHTTPNotFound("course not found by code", err)
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("course not found")
		return xexception.NewHTTPNotFound("course not found", nil)
	}

	currentTime := time.Now()
	if currentTime.After(course.EndDate) {
		err := xexception.NewHTTPBadRequest("course has already ended", nil)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Warn().Str("course_code", courseCode).Msg("course has already ended, cannot release seat")
		return err
	}

	course.SeatAvailable++
	if err := s.store.Save(ctx, course); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to update course availability")
		return err
	}

	metrics.ReservedCourses.Dec()

	log.Info().Str("course_code", courseCode).Int("remaining_seats", course.SeatAvailable).Msg("Sseat released successfully")
	return nil
}

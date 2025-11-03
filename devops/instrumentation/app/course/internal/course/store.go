package course

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

//go:generate mockery --name=IStore
type IStore interface {
	FindAll(ctx context.Context) ([]Course, error)
	FindByID(ctx context.Context, courseID string) (*Course, error)
	FindByCode(ctx context.Context, courseCode string) (*Course, error)
	Save(ctx context.Context, course *Course) error
	DeleteByID(ctx context.Context, courseID string) error
}

type store struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewStore(db *gorm.DB, tracer trace.Tracer) IStore {
	return &store{
		db:     db,
		tracer: tracer,
	}
}
func (s *store) FindAll(ctx context.Context) ([]Course, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindAll").Logger()

	var courses []Course
	if err := s.db.WithContext(ctx).Find(&courses).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("Failed to fetch all courses")
		return nil, err
	}

	log.Info().Int("courses_count", len(courses)).Msg("Successfully fetched all courses")
	return courses, nil
}

func (s *store) FindByID(ctx context.Context, courseID string) (*Course, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindByID").Logger()

	var course Course
	if err := s.db.WithContext(ctx).First(&course, "id = ?", courseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Str("course_id", courseID).Msg("Course not found")
			return nil, nil
		}
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("Failed to fetch course by ID")
		return nil, err
	}

	log.Info().Str("course_id", course.ID).Msg("Successfully fetched course by ID")
	return &course, nil
}

func (s *store) FindByCode(ctx context.Context, courseCode string) (*Course, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindByCode").Logger()

	var course Course
	if err := s.db.WithContext(ctx).First(&course, "code = ?", courseCode).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Str("course_code", courseCode).Msg("Course not found")
			return nil, nil
		}
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to fetch course by code")
		return nil, err
	}

	log.Info().Str("course_code", course.Code).Msg("Successfully fetched course by code")
	return &course, nil
}

func (s *store) Save(ctx context.Context, course *Course) error {
	ctx, span := s.tracer.Start(ctx, "store.Save")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", course.ID), attribute.String("course.code", course.Code))

	log := zerolog.Ctx(ctx).With().Str("component", "store.Save").Logger()

	log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Int("seat_available_before", course.SeatAvailable).Msg("Saving course")

	if err := s.db.WithContext(ctx).Save(course).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to save course")
		return err
	}

	log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Int("seat_available_after", course.SeatAvailable).Msg("Course saved successfully")
	return nil
}

func (s *store) DeleteByID(ctx context.Context, courseID string) error {
	ctx, span := s.tracer.Start(ctx, "store.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "store.DeleteByID").Logger()

	if err := s.db.WithContext(ctx).Delete(&Course{}, "id = ?", courseID).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("Failed to delete course")
		return err
	}

	log.Info().Str("course_id", courseID).Msg("Course deleted successfully")
	return nil
}

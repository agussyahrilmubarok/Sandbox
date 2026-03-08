package course

import (
	"context"
	"errors"

	"example.com/course/internal/tracing"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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
	db *gorm.DB
}

func NewStore(db *gorm.DB) IStore {
	return &store{
		db: db,
	}
}
func (s *store) FindAll(ctx context.Context) ([]Course, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindAll").Logger()

	var courses []Course
	if err := s.db.WithContext(ctx).Find(&courses).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("failed to fetch all courses")
		return nil, err
	}

	log.Info().Int("courses_count", len(courses)).Msg("successfully fetched all courses")
	return courses, nil
}

func (s *store) FindByID(ctx context.Context, courseID string) (*Course, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindByID").Logger()

	var course Course
	if err := s.db.WithContext(ctx).First(&course, "id = ?", courseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Str("course_id", courseID).Msg("course not found")
			return nil, nil
		}
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("failed to fetch course by id")
		return nil, err
	}

	log.Info().Str("course_id", course.ID).Msg("successfully fetched course by id")
	return &course, nil
}

func (s *store) FindByCode(ctx context.Context, courseCode string) (*Course, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindByCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindByCode").Logger()

	var course Course
	if err := s.db.WithContext(ctx).First(&course, "code = ?", courseCode).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Str("course_code", courseCode).Msg("course not found")
			return nil, nil
		}
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to fetch course by code")
		return nil, err
	}

	log.Info().Str("course_code", course.Code).Msg("successfully fetched course by code")
	return &course, nil
}

func (s *store) Save(ctx context.Context, course *Course) error {
	ctx, span := tracing.StartSpan(ctx, "store.Save")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", course.ID), attribute.String("course.code", course.Code))

	log := zerolog.Ctx(ctx).With().Str("component", "store.Save").Logger()

	log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Int("seat_available_before", course.SeatAvailable).Msg("saving course")

	if err := s.db.WithContext(ctx).Save(course).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", course.ID).Msg("failed to save course")
		return err
	}

	log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Int("seat_available_after", course.SeatAvailable).Msg("course saved successfully")
	return nil
}

func (s *store) DeleteByID(ctx context.Context, courseID string) error {
	ctx, span := tracing.StartSpan(ctx, "store.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("course.id", courseID))

	log := zerolog.Ctx(ctx).With().Str("component", "store.DeleteByID").Logger()

	if err := s.db.WithContext(ctx).Delete(&Course{}, "id = ?", courseID).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_id", courseID).Msg("failed to delete course by id")
		return err
	}

	log.Info().Str("course_id", courseID).Msg("course deleted successfully")
	return nil
}

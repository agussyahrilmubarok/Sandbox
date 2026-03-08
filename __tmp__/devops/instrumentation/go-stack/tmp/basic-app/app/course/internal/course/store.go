package course

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
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
	db  *gorm.DB
	log zerolog.Logger
}

func NewStore(db *gorm.DB, log zerolog.Logger) IStore {
	return &store{
		db:  db,
		log: log.With().Str("component", "course_store").Logger(),
	}
}

func (s *store) FindAll(ctx context.Context) ([]Course, error) {
	var courses []Course
	if err := s.db.WithContext(ctx).Find(&courses).Error; err != nil {
		s.log.Error().Err(err).Msg("Failed to fetch all courses")
		return nil, err
	}

	s.log.Info().Int("courses_count", len(courses)).Msg("Successfully fetched all courses")
	return courses, nil
}

func (s *store) FindByID(ctx context.Context, courseID string) (*Course, error) {
	var course Course
	if err := s.db.WithContext(ctx).First(&course, "id = ?", courseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.Warn().Str("course_id", courseID).Msg("Course not found")
			return nil, nil
		}
		s.log.Error().Err(err).Str("course_id", courseID).Msg("Failed to fetch course by ID")
		return nil, err
	}

	s.log.Info().Str("course_id", course.ID).Msg("Successfully fetched course by ID")
	return &course, nil
}

func (s *store) FindByCode(ctx context.Context, courseCode string) (*Course, error) {
	var course Course
	if err := s.db.WithContext(ctx).First(&course, "code = ?", courseCode).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.Warn().Str("course_code", courseCode).Msg("Course not found")
			return nil, nil
		}
		s.log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to fetch course by code")
		return nil, err
	}

	s.log.Info().Str("course_code", course.Code).Msg("Successfully fetched course by code")
	return &course, nil
}

func (s *store) Save(ctx context.Context, course *Course) error {
	if err := s.db.WithContext(ctx).Save(course).Error; err != nil {
		s.log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to save course")
		return err
	}

	s.log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Msg("Course saved successfully")
	return nil
}

func (s *store) DeleteByID(ctx context.Context, courseID string) error {
	if err := s.db.WithContext(ctx).Delete(&Course{}, "id = ?", courseID).Error; err != nil {
		s.log.Error().Err(err).Str("course_id", courseID).Msg("Failed to delete course")
		return err
	}

	s.log.Info().Str("course_id", courseID).Msg("Course deleted successfully")
	return nil
}

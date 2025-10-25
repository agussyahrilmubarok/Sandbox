package course

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
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
	store IStore
}

func NewService(store IStore) IService {
	return &service{
		store: store,
	}
}

func (s *service) FindAll(ctx context.Context) ([]Course, error) {
	log := zerolog.Ctx(ctx)

	courses, err := s.store.FindAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch all courses from store")
		return nil, err
	}

	log.Info().Int("courses_count", len(courses)).Msg("Successfully fetched all courses")
	return courses, nil
}

func (s *service) FindByID(ctx context.Context, courseID string) (*Course, error) {
	log := zerolog.Ctx(ctx)

	course, err := s.store.FindByID(ctx, courseID)
	if err != nil {
		log.Error().Err(err).Str("course_id", courseID).Msg("Failed to find course by ID")
		return nil, err
	}

	if course == nil {
		log.Warn().Str("course_id", courseID).Msg("Course not found")
		return nil, errors.New("course not found")
	}

	log.Info().Str("course_id", courseID).Msg("Successfully found course by ID")
	return course, nil
}

func (s *service) FindByCode(ctx context.Context, courseCode string) (*Course, error) {
	log := zerolog.Ctx(ctx)

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to find course by code")
		return nil, err
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("Course not found")
		return nil, errors.New("course not found")
	}

	log.Info().Str("course_code", courseCode).Msg("Successfully found course by code")
	return course, nil
}

func (s *service) Save(ctx context.Context, course *Course) error {
	log := zerolog.Ctx(ctx)

	err := s.store.Save(ctx, course)
	if err != nil {
		log.Error().Err(err).Str("course_id", course.ID).Msg("Failed to save course")
		return err
	}

	log.Info().Str("course_id", course.ID).Str("course_code", course.Code).Msg("Successfully saved course")
	return nil
}

func (s *service) DeleteByID(ctx context.Context, courseID string) error {
	log := zerolog.Ctx(ctx)

	err := s.store.DeleteByID(ctx, courseID)
	if err != nil {
		log.Error().Err(err).Str("course_id", courseID).Msg("Failed to delete course")
		return err
	}

	log.Info().Str("course_id", courseID).Msg("Successfully deleted course")
	return nil
}

func (s *service) ReserveByCode(ctx context.Context, courseCode string) error {
	log := zerolog.Ctx(ctx)

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to find course by code")
		return err
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("Course not found")
		return errors.New("course not found")
	}

	currentTime := time.Now()

	if currentTime.After(course.EndDate) {
		log.Warn().Str("course_code", courseCode).Msg("Course has already ended, cannot reserve seat")
		return errors.New("course has already ended")
	}

	if course.SeatAvailable <= 0 {
		log.Warn().Str("course_code", courseCode).Msg("No available seats to reserve")
		return errors.New("no available seats to reserve")
	}

	course.SeatAvailable--

	if err := s.store.Save(ctx, course); err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to update course availability")
		return err
	}

	log.Info().Str("course_code", courseCode).Int("remaining_seats", course.SeatAvailable).Msg("Course reserved successfully")
	return nil
}

func (s *service) ReleaseByCode(ctx context.Context, courseCode string) error {
	log := zerolog.Ctx(ctx)

	course, err := s.store.FindByCode(ctx, courseCode)
	if err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to find course by code")
		return err
	}

	if course == nil {
		log.Warn().Str("course_code", courseCode).Msg("Course not found")
		return errors.New("course not found")
	}

	currentTime := time.Now()

	if currentTime.After(course.EndDate) {
		log.Warn().Str("course_code", courseCode).Msg("Course has already ended, cannot release seat")
		return errors.New("course has already ended")
	}

	course.SeatAvailable++

	if err := s.store.Save(ctx, course); err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to update course availability")
		return err
	}

	log.Info().Str("course_code", courseCode).Int("remaining_seats", course.SeatAvailable).Msg("Seat released successfully")
	return nil
}

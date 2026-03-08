package booking

import (
	"context"

	"example.com/booking/internal/tracing"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"
)

//go:generate mockery --name=IStore
type IStore interface {
	FindAll(ctx context.Context) ([]Booking, error)
	FindAllByMemberCode(ctx context.Context, memberCode string) ([]Booking, error)
	FindAllByCourseCode(ctx context.Context, courseCode string) ([]Booking, error)
	FindAllByStatus(ctx context.Context, status BookingStatus) ([]Booking, error)
	FindByID(ctx context.Context, bookingID string) (*Booking, error)
	Save(ctx context.Context, booking *Booking) error
	DeleteByID(ctx context.Context, bookingID string) error
}

type store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) IStore {
	return &store{
		db: db,
	}
}

func (s *store) FindAll(ctx context.Context) ([]Booking, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("failed to fetch all bookings")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Msg("successfully fetched all bookings")
	return bookings, nil
}

func (s *store) FindAllByMemberCode(ctx context.Context, memberCode string) ([]Booking, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindAllByMemberCode")
	defer span.End()
	span.SetAttributes(attribute.String("member.code", memberCode))

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Where("member_code = ?", memberCode).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_code", memberCode).Msg("failed to fetch bookings by member code")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Str("member_code", memberCode).Msg("successfully fetched bookings for member")
	return bookings, nil
}

func (s *store) FindAllByCourseCode(ctx context.Context, courseCode string) ([]Booking, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindAllByCourseCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Where("course_code = ?", courseCode).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("failed to fetch bookings by course code")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Str("course_code", courseCode).Msg("successfully fetched bookings for course")
	return bookings, nil
}

func (s *store) FindAllByStatus(ctx context.Context, status BookingStatus) ([]Booking, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindAllByStatus")
	defer span.End()
	span.SetAttributes(attribute.String("booking.status", status.String()))

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Where("status = ?", status).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("status", status.String()).Msg("failed to fetch bookings by status")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Str("status", status.String()).Msg("successfully fetched bookings by status")
	return bookings, nil
}

func (s *store) FindByID(ctx context.Context, bookingID string) (*Booking, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("booking.id", bookingID))

	log := zerolog.Ctx(ctx)

	var booking Booking
	if err := s.db.WithContext(ctx).First(&booking, "id = ?", bookingID).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("booking_id", bookingID).Msg("failed to fetch booking by id")
		return nil, err
	}

	log.Info().Str("booking_id", booking.ID).Msg("successfully fetched booking by id")
	return &booking, nil
}

func (s *store) Save(ctx context.Context, booking *Booking) error {
	ctx, span := tracing.StartSpan(ctx, "store.Save")
	defer span.End()
	span.SetAttributes(attribute.String("booking.id", booking.ID))

	log := zerolog.Ctx(ctx)

	if err := s.db.WithContext(ctx).Save(booking).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("booking_id", booking.ID).Msg("failed to save booking")
		return err
	}

	log.Info().Str("booking_id", booking.ID).Msg("Booking saved successfully")
	return nil
}

func (s *store) DeleteByID(ctx context.Context, bookingID string) error {
	ctx, span := tracing.StartSpan(ctx, "store.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("booking.id", bookingID))

	log := zerolog.Ctx(ctx)

	if err := s.db.WithContext(ctx).Delete(&Booking{}, "id = ?", bookingID).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("booking_id", bookingID).Msg("failed to delete booking")
		return err
	}

	log.Info().Str("booking_id", bookingID).Msg("booking deleted successfully")
	return nil
}

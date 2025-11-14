package booking

import (
	"context"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

//go:generate mockery --name=IStore
type IStore interface {
	// FindAll retrieves all bookings.
	FindAll(ctx context.Context) ([]Booking, error)

	// FindAllByMemberCode retrieves all bookings for a specific member.
	FindAllByMemberCode(ctx context.Context, memberCode string) ([]Booking, error)

	// FindAllByCourseCode retrieves all bookings for a specific course.
	FindAllByCourseCode(ctx context.Context, courseCode string) ([]Booking, error)

	// FindAllByStatus retrieves all bookings filtered by status.
	FindAllByStatus(ctx context.Context, status BookingStatus) ([]Booking, error)

	// FindByID retrieves a booking by its ID.
	FindByID(ctx context.Context, bookingID string) (*Booking, error)

	// Save saves a new or updated booking in the database.
	Save(ctx context.Context, booking *Booking) error

	// DeleteByID deletes a booking by its ID.
	DeleteByID(ctx context.Context, bookingID string) error
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

func (s *store) FindAll(ctx context.Context) ([]Booking, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("Failed to fetch all bookings")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Msg("Successfully fetched all bookings")
	return bookings, nil
}

func (s *store) FindAllByMemberCode(ctx context.Context, memberCode string) ([]Booking, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindAllByMemberCode")
	defer span.End()
	span.SetAttributes(attribute.String("member.code", memberCode))

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Where("member_code = ?", memberCode).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_code", memberCode).Msg("Failed to fetch bookings by member code")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Str("member_code", memberCode).Msg("Successfully fetched bookings for member")
	return bookings, nil
}

func (s *store) FindAllByCourseCode(ctx context.Context, courseCode string) ([]Booking, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindAllByCourseCode")
	defer span.End()
	span.SetAttributes(attribute.String("course.code", courseCode))

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Where("course_code = ?", courseCode).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to fetch bookings by course code")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Str("course_code", courseCode).Msg("Successfully fetched bookings for course")
	return bookings, nil
}

func (s *store) FindAllByStatus(ctx context.Context, status BookingStatus) ([]Booking, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindAllByStatus")
	defer span.End()
	span.SetAttributes(attribute.String("booking.status", status.String()))

	log := zerolog.Ctx(ctx)

	var bookings []Booking
	if err := s.db.WithContext(ctx).Where("status = ?", status).Find(&bookings).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("status", status.String()).Msg("Failed to fetch bookings by status")
		return nil, err
	}

	log.Info().Int("booking_count", len(bookings)).Str("status", status.String()).Msg("Successfully fetched bookings by status")
	return bookings, nil
}

func (s *store) FindByID(ctx context.Context, bookingID string) (*Booking, error) {
	ctx, span := s.tracer.Start(ctx, "store.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("booking.id", bookingID))

	log := zerolog.Ctx(ctx)

	var booking Booking
	if err := s.db.WithContext(ctx).First(&booking, "id = ?", bookingID).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("booking_id", bookingID).Msg("Failed to fetch booking by ID")
		return nil, err
	}

	log.Info().Str("booking_id", booking.ID).Msg("Successfully fetched booking by ID")
	return &booking, nil
}

func (s *store) Save(ctx context.Context, booking *Booking) error {
	ctx, span := s.tracer.Start(ctx, "store.Save")
	defer span.End()
	span.SetAttributes(attribute.String("booking.id", booking.ID))

	log := zerolog.Ctx(ctx)

	if err := s.db.WithContext(ctx).Save(booking).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("booking_id", booking.ID).Msg("Failed to save booking")
		return err
	}

	log.Info().Str("booking_id", booking.ID).Msg("Booking saved successfully")
	return nil
}

func (s *store) DeleteByID(ctx context.Context, bookingID string) error {
	ctx, span := s.tracer.Start(ctx, "store.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("booking.id", bookingID))

	log := zerolog.Ctx(ctx)

	if err := s.db.WithContext(ctx).Delete(&Booking{}, "id = ?", bookingID).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("booking_id", bookingID).Msg("Failed to delete booking")
		return err
	}

	log.Info().Str("booking_id", bookingID).Msg("Booking deleted successfully")
	return nil
}

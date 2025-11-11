package booking

import (
	"context"
	"fmt"
	"time"

	"example.com/booking/internal/metrics"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

//go:generate mockery --name=IService
type IService interface {
	Booking(ctx context.Context, request BookingRequest) (*Booking, error)
	BookingV2(ctx context.Context, request BookingRequest) (*Booking, error)
}

type service struct {
	store  IStore
	client IClient
	tracer trace.Tracer
}

func NewService(store IStore, client IClient, tracer trace.Tracer) IService {
	return &service{
		store:  store,
		client: client,
		tracer: tracer,
	}
}

func (s *service) Booking(ctx context.Context, request BookingRequest) (*Booking, error) {
	ctx, span := s.tracer.Start(ctx, "service.Booking")
	defer span.End()
	span.SetAttributes(
		attribute.String("member.code", request.MemberCode),
		attribute.String("course.code", request.CourseCode),
	)

	log := zerolog.Ctx(ctx).With().Str("component", "service.Booking").Logger()

	memberResp, err := s.client.FindMemberByCode(ctx, request.MemberCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		metrics.FailedBookings.Inc()
		log.Error().Err(err).Str("member_code", request.MemberCode).Msg("Failed to find member")
		return nil, xexception.NewHTTPNotFound("Member not found", err)
	}

	courseResp, err := s.client.ReserveCourseByCode(ctx, request.CourseCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		metrics.FailedBookings.Inc()
		log.Error().Err(err).Str("course_code", request.CourseCode).Msg("Failed to reserve course")
		return nil, xexception.NewHTTPNotFound("Failed to reserve course", err)
	}
	log.Info().Str("course_code", request.CourseCode).Msg("Course reserved: " + courseResp.Message)

	booking := &Booking{
		ID:          uuid.New().String(),
		MemberCode:  memberResp.Code,
		CourseCode:  request.CourseCode,
		BookingDate: time.Now(),
		Status:      BookingStatusPending,
		Notes:       fmt.Sprintf("Booking Course %s by %s", request.CourseCode, request.MemberCode),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.store.Save(ctx, booking); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		metrics.FailedBookings.Inc()
		log.Error().Err(err).Str("booking_id", booking.ID).Msg("Failed to save booking, releasing course")
		if _, releaseErr := s.client.ReleaseCourseByCode(ctx, request.CourseCode); releaseErr != nil {
			log.Error().Err(releaseErr).Str("course_code", request.CourseCode).Msg("Failed to release course after save failure")
		}

		return nil, xexception.NewHTTPInternal("Failed to save booking", err)
	}

	metrics.SuccessfulBookings.Inc()
	metrics.BookingStatus.WithLabelValues(BookingStatusPending.String()).Inc()

	log.Info().Str("booking_id", booking.ID).Str("member_code", memberResp.Code).Str("course_code", request.CourseCode).Msg("Booking successfully created")
	return booking, nil
}

func (s *service) BookingV2(ctx context.Context, request BookingRequest) (*Booking, error) {
	ctx, span := s.tracer.Start(ctx, "service.BookingV2")
	defer span.End()
	span.SetAttributes(
		attribute.String("member.code", request.MemberCode),
		attribute.String("course.code", request.CourseCode),
	)

	log := zerolog.Ctx(ctx).With().Str("component", "service.BookingV2").Logger()

	type result struct {
		member *MemberServiceMemberResponse
		course *CourseServiceMessageResponse
		err    error
	}

	memberCh := make(chan result, 1)
	courseCh := make(chan result, 1)

	go func() {
		mCtx, mSpan := s.tracer.Start(ctx, "client.FindMemberByCode")
		defer mSpan.End()
		memberResp, err := s.client.FindMemberByCode(mCtx, request.MemberCode)
		if err != nil {
			mSpan.RecordError(err)
			mSpan.SetStatus(codes.Error, err.Error())
		}
		memberCh <- result{member: memberResp, err: err}
	}()

	go func() {
		cCtx, cSpan := s.tracer.Start(ctx, "client.ReserveCourseByCode")
		defer cSpan.End()
		courseResp, err := s.client.ReserveCourseByCode(cCtx, request.CourseCode)
		if err != nil {
			cSpan.RecordError(err)
			cSpan.SetStatus(codes.Error, err.Error())
		}
		courseCh <- result{course: courseResp, err: err}
	}()

	var memberRes *MemberServiceMemberResponse
	// var courseRes *CourseServiceMessageResponse

	for i := 0; i < 2; i++ {
		select {
		case m := <-memberCh:
			if m.err != nil {
				span.RecordError(m.err)
				span.SetStatus(codes.Error, m.err.Error())
				metrics.FailedBookings.Inc()
				return nil, xexception.NewHTTPNotFound("Member not found", m.err)
			}
			memberRes = m.member
		case c := <-courseCh:
			if c.err != nil {
				span.RecordError(c.err)
				span.SetStatus(codes.Error, c.err.Error())
				metrics.FailedBookings.Inc()
				return nil, xexception.NewHTTPNotFound("Failed to reserve course", c.err)
			}
			// courseRes = c.course
		case <-ctx.Done():
			err := xexception.NewHTTPRequestTimeout("Context cancelled or timed out", nil)
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
	}

	booking := &Booking{
		ID:          uuid.New().String(),
		MemberCode:  memberRes.Code,
		CourseCode:  request.CourseCode,
		BookingDate: time.Now(),
		Status:      BookingStatusPending,
		Notes:       fmt.Sprintf("Booking Course %s by %s", request.CourseCode, request.MemberCode),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.store.Save(ctx, booking); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		metrics.FailedBookings.Inc()
		log.Error().Err(err).Str("booking_id", booking.ID).Msg("Failed to save booking, releasing course")
		if _, releaseErr := s.client.ReleaseCourseByCode(ctx, request.CourseCode); releaseErr != nil {
			log.Error().Err(releaseErr).Str("course_code", request.CourseCode).Msg("Failed to release course after save failure")
		}
		return nil, xexception.NewHTTPInternal("Failed to save booking", err)
	}

	metrics.SuccessfulBookings.Inc()
	metrics.BookingStatus.WithLabelValues(BookingStatusPending.String()).Inc()

	log.Info().Str("booking_id", booking.ID).Str("member_code", memberRes.Code).Str("course_code", request.CourseCode).Msg("Booking successfully created (V2)")
	return booking, nil
}

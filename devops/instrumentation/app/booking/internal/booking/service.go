package booking

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type IService interface {
	Booking(ctx context.Context, request BookingRequest) (*Booking, error)
	BookingV2(ctx context.Context, request BookingRequest) (*Booking, error)
}

type service struct {
	store  IStore
	client IClient
	log    zerolog.Logger
}

func NewService(store IStore, client IClient, log zerolog.Logger) IService {
	return &service{
		store:  store,
		client: client,
		log:    log.With().Str("component", "booking_service").Logger(),
	}
}

func (s *service) Booking(ctx context.Context, request BookingRequest) (*Booking, error) {
	memberResp, err := s.client.FindMemberByCode(ctx, request.MemberCode)
	if err != nil {
		s.log.Error().Err(err).Str("member_code", request.MemberCode).Msg("Failed to find member")
		return nil, fmt.Errorf("member not found: %w", err)
	}

	courseResp, err := s.client.ReserveCourseByCode(ctx, request.CourseCode)
	if err != nil {
		s.log.Error().Err(err).Str("course_code", request.CourseCode).Msg("Failed to reserve course")
		return nil, fmt.Errorf("failed to reserve course: %w", err)
	}
	s.log.Info().Str("course_code", request.CourseCode).Msg("Course reserved: " + courseResp.Message)

	booking := &Booking{
		ID:          uuid.New().String(),
		MemberCode:  memberResp.Code,
		CourseCode:  request.CourseCode,
		BookingDate: time.Now(),
		Status:      BookingStatusConfirmed,
		Notes:       fmt.Sprintf("Booking Course %s by %s", request.CourseCode, request.MemberCode),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.store.Save(ctx, booking); err != nil {
		s.log.Error().Err(err).Str("booking_id", booking.ID).Msg("Failed to save booking, releasing course")
		if _, releaseErr := s.client.ReleaseCourseByCode(ctx, request.CourseCode); releaseErr != nil {
			s.log.Error().Err(releaseErr).Str("course_code", request.CourseCode).Msg("Failed to release course after save failure")
		}

		return nil, fmt.Errorf("failed to save booking: %w", err)
	}

	s.log.Info().Str("booking_id", booking.ID).Str("member_code", memberResp.Code).Str("course_code", request.CourseCode).Msg("Booking successfully created")
	return booking, nil
}

func (s *service) BookingV2(ctx context.Context, request BookingRequest) (*Booking, error) {
	type result struct {
		member *MemberServiceMemberResponse
		course *CourseServiceMessageResponse
		err    error
	}

	memberCh := make(chan result, 1)
	courseCh := make(chan result, 1)

	go func() {
		memberResp, err := s.client.FindMemberByCode(ctx, request.MemberCode)
		memberCh <- result{member: memberResp, err: err}
	}()

	go func() {
		courseResp, err := s.client.ReserveCourseByCode(ctx, request.CourseCode)
		courseCh <- result{course: courseResp, err: err}
	}()

	var memberRes *MemberServiceMemberResponse
	// var courseRes *CourseServiceMessageResponse

	for i := 0; i < 2; i++ {
		select {
		case m := <-memberCh:
			if m.err != nil {
				return nil, fmt.Errorf("member not found: %w", m.err)
			}
			memberRes = m.member
		case c := <-courseCh:
			if c.err != nil {
				return nil, fmt.Errorf("failed to reserve course: %w", c.err)
			}
			// courseRes = c.course
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled or timed out")
		}
	}

	booking := &Booking{
		ID:          uuid.New().String(),
		MemberCode:  memberRes.Code,
		CourseCode:  request.CourseCode,
		BookingDate: time.Now(),
		Status:      BookingStatusConfirmed,
		Notes:       fmt.Sprintf("Booking Course %s by %s", request.CourseCode, request.MemberCode),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.store.Save(ctx, booking); err != nil {
		s.log.Error().Err(err).Str("booking_id", booking.ID).Msg("Failed to save booking, releasing course")
		if _, releaseErr := s.client.ReleaseCourseByCode(ctx, request.CourseCode); releaseErr != nil {
			s.log.Error().Err(releaseErr).Str("course_code", request.CourseCode).Msg("Failed to release course after save failure")
		}
		return nil, fmt.Errorf("failed to save booking: %w", err)
	}

	s.log.Info().Str("booking_id", booking.ID).Str("member_code", memberRes.Code).Str("course_code", request.CourseCode).Msg("Booking successfully created (V2)")
	return booking, nil
}

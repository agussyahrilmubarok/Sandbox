package member

import (
	"context"

	"example.com/member/internal/metrics"
	"example.com/member/pkg/exception"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

//go:generate mockery --name=IService
type IService interface {
	FindAll(ctx context.Context) ([]Member, error)
	FindByID(ctx context.Context, memberID string) (*Member, error)
	FindByCode(ctx context.Context, memberCode string) (*Member, error)
	FindByEmail(ctx context.Context, memberEmail string) (*Member, error)
	Save(ctx context.Context, member *Member) (*Member, error)
	DeleteByID(ctx context.Context, memberID string) error
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

func (s *service) FindAll(ctx context.Context) ([]Member, error) {
	ctx, span := s.tracer.Start(ctx, "service.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindAll").Logger()

	members, err := s.store.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("Failed to fetch members")
		return nil, err
	}

	log.Info().Int("count", len(members)).Msg("Successfully fetched all members")
	return members, nil
}

func (s *service) FindByID(ctx context.Context, memberID string) (*Member, error) {
	ctx, span := s.tracer.Start(ctx, "service.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", memberID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByID").Logger()

	member, err := s.store.FindByID(ctx, memberID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", memberID).Msg("Failed to fetch member by ID")
		return nil, exception.NewNotFound("Member not found", err)
	}

	if member == nil {
		log.Warn().Str("member_id", memberID).Msg("Member not found")
		return nil, exception.NewNotFound("Member not found", nil)
	}

	log.Info().Str("member_id", member.ID).Msg("Successfully fetched member by ID")
	return member, nil
}

func (s *service) FindByCode(ctx context.Context, memberCode string) (*Member, error) {
	ctx, span := s.tracer.Start(ctx, "service.FindByCode")
	defer span.End()
	span.SetAttributes(attribute.String("member.code", memberCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByCode").Logger()

	member, err := s.store.FindByCode(ctx, memberCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_code", memberCode).Msg("Failed to fetch member by Code")
		return nil, exception.NewNotFound("Member not found", err)
	}

	if member == nil {
		log.Warn().Str("member_code", memberCode).Msg("Member not found")
		return nil, exception.NewNotFound("Member not found", nil)
	}

	log.Info().Str("member_code", member.ID).Msg("Successfully fetched member by Code")
	return member, nil
}

func (s *service) FindByEmail(ctx context.Context, memberEmail string) (*Member, error) {
	ctx, span := s.tracer.Start(ctx, "service.FindByEmail")
	defer span.End()
	span.SetAttributes(attribute.String("member.email", memberEmail))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByEmail").Logger()

	member, err := s.store.FindByEmail(ctx, memberEmail)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_email", memberEmail).Msg("Failed to fetch member by Email")
		return nil, exception.NewNotFound("Member not found", err)
	}

	if member == nil {
		log.Warn().Str("member_email", memberEmail).Msg("Member not found")
		return nil, exception.NewNotFound("Member not found", nil)
	}

	log.Info().Str("member_email", member.ID).Msg("Successfully fetched member by Email")
	return member, nil
}

func (s *service) Save(ctx context.Context, member *Member) (*Member, error) {
	ctx, span := s.tracer.Start(ctx, "service.Save")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", member.ID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.Save").Logger()

	if err := s.store.Save(ctx, member); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", member.ID).Str("member_code", member.Code).Msg("Failed to save member")
		return nil, err
	}

	metrics.TotalMembers.Inc()

	log.Info().Str("member_id", member.ID).Str("member_code", member.Code).Msg("Member saved successfully")
	return member, nil
}

func (s *service) DeleteByID(ctx context.Context, memberID string) error {
	ctx, span := s.tracer.Start(ctx, "service.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", memberID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.DeleteByID").Logger()

	if err := s.store.DeleteByID(ctx, memberID); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", memberID).Msg("Failed to delete member")
		return err
	}

	metrics.TotalMembers.Dec()

	log.Info().Str("member_id", memberID).Msg("Member deleted successfully")
	return nil
}

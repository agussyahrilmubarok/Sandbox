package member

import (
	"context"

	"example.com/member/internal/metrics"
	"example.com/member/internal/tracing"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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
	store IStore
}

func NewService(store IStore) IService {
	return &service{
		store: store,
	}
}

func (s *service) FindAll(ctx context.Context) ([]Member, error) {
	ctx, span := tracing.StartSpan(ctx, "service.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindAll").Logger()

	members, err := s.store.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("failed to fetch members")
		return nil, err
	}

	log.Info().Int("count", len(members)).Msg("successfully fetched all members")
	return members, nil
}

func (s *service) FindByID(ctx context.Context, memberID string) (*Member, error) {
	ctx, span := tracing.StartSpan(ctx, "service.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", memberID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByID").Logger()

	member, err := s.store.FindByID(ctx, memberID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", memberID).Msg("failed to fetch member by id")
		return nil, xexception.NewHTTPNotFound("member not found", err)
	}

	if member == nil {
		log.Warn().Str("member_id", memberID).Msg("member not found")
		return nil, xexception.NewHTTPNotFound("member not found", nil)
	}

	log.Info().Str("member_id", member.ID).Msg("successfully fetched member by id")
	return member, nil
}

func (s *service) FindByCode(ctx context.Context, memberCode string) (*Member, error) {
	ctx, span := tracing.StartSpan(ctx, "service.FindByCode")
	defer span.End()
	span.SetAttributes(attribute.String("member.code", memberCode))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByCode").Logger()

	member, err := s.store.FindByCode(ctx, memberCode)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_code", memberCode).Msg("failed to fetch member by code")
		return nil, xexception.NewHTTPNotFound("member not found", err)
	}

	if member == nil {
		log.Warn().Str("member_code", memberCode).Msg("member not found")
		return nil, xexception.NewHTTPNotFound("member not found", nil)
	}

	log.Info().Str("member_code", member.ID).Msg("successfully fetched member by code")
	return member, nil
}

func (s *service) FindByEmail(ctx context.Context, memberEmail string) (*Member, error) {
	ctx, span := tracing.StartSpan(ctx, "service.FindByEmail")
	defer span.End()
	span.SetAttributes(attribute.String("member.email", memberEmail))

	log := zerolog.Ctx(ctx).With().Str("component", "service.FindByEmail").Logger()

	member, err := s.store.FindByEmail(ctx, memberEmail)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_email", memberEmail).Msg("failed to fetch member by email")
		return nil, xexception.NewHTTPNotFound("member not found", err)
	}

	if member == nil {
		log.Warn().Str("member_email", memberEmail).Msg("member not found")
		return nil, xexception.NewHTTPNotFound("member not found", nil)
	}

	log.Info().Str("member_email", member.ID).Msg("successfully fetched member by email")
	return member, nil
}

func (s *service) Save(ctx context.Context, member *Member) (*Member, error) {
	ctx, span := tracing.StartSpan(ctx, "service.Save")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", member.ID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.Save").Logger()

	if err := s.store.Save(ctx, member); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", member.ID).Str("member_code", member.Code).Msg("failed to save member")
		return nil, xexception.NewHTTPInternal("failed save member", err)
	}

	metrics.TotalMembers.Inc()

	log.Info().Str("member_id", member.ID).Str("member_code", member.Code).Msg("member saved successfully")
	return member, nil
}

func (s *service) DeleteByID(ctx context.Context, memberID string) error {
	ctx, span := tracing.StartSpan(ctx, "service.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", memberID))

	log := zerolog.Ctx(ctx).With().Str("component", "service.DeleteByID").Logger()

	if err := s.store.DeleteByID(ctx, memberID); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", memberID).Msg("failed to delete member by id")
		return err
	}

	metrics.TotalMembers.Dec()

	log.Info().Str("member_id", memberID).Msg("member deleted successfully")
	return nil
}

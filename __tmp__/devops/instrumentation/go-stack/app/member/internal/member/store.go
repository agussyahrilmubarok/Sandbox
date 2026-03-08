package member

import (
	"context"

	"example.com/member/internal/tracing"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"
)

//go:generate mockery --name=IStore
type IStore interface {
	FindAll(ctx context.Context) ([]Member, error)
	FindByID(ctx context.Context, memberID string) (*Member, error)
	FindByCode(ctx context.Context, memberCode string) (*Member, error)
	FindByEmail(ctx context.Context, memberEmail string) (*Member, error)
	Save(ctx context.Context, member *Member) error
	DeleteByID(ctx context.Context, memberID string) error
}

type store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) IStore {
	return &store{
		db: db,
	}
}

func (s *store) FindAll(ctx context.Context) ([]Member, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindAll")
	defer span.End()

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindAll").Logger()

	var members []Member
	if err := s.db.WithContext(ctx).Find(&members).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Msg("failed to find all members")
		return nil, err
	}

	log.Info().Int("count", len(members)).Msg("successfully fetched all members")
	return members, nil
}

func (s *store) FindByID(ctx context.Context, memberID string) (*Member, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindByID")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", memberID))

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindByID").Logger()

	var m Member
	if err := s.db.WithContext(ctx).First(&m, "id = ?", memberID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn().Str("member_id", memberID).Msg("member not found")
			return nil, nil
		}
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", memberID).Msg("failed to find member by id")
		return nil, err
	}

	log.Info().Str("member_id", m.ID).Msg("successfully fetched member by id")
	return &m, nil
}

func (s *store) FindByCode(ctx context.Context, memberCode string) (*Member, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindByCode")
	defer span.End()
	span.SetAttributes(attribute.String("member.code", memberCode))

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindByCode").Logger()

	var m Member
	if err := s.db.WithContext(ctx).First(&m, "code = ?", memberCode).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn().Str("member_code", memberCode).Msg("member not found")
			return nil, nil
		}
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_code", memberCode).Msg("failed to find member by code")
		return nil, err
	}

	log.Info().Str("member_code", m.ID).Msg("successfully fetched member by code")
	return &m, nil
}

func (s *store) FindByEmail(ctx context.Context, memberEmail string) (*Member, error) {
	ctx, span := tracing.StartSpan(ctx, "store.FindByEmail")
	defer span.End()
	span.SetAttributes(attribute.String("member.email", memberEmail))

	log := zerolog.Ctx(ctx).With().Str("component", "store.FindByEmail").Logger()

	var m Member
	if err := s.db.WithContext(ctx).First(&m, "email = ?", memberEmail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn().Str("member_email", memberEmail).Msg("member not found")
			return nil, nil
		}
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_email", memberEmail).Msg("failed to find member by email")
		return nil, err
	}

	log.Info().Str("member_email", m.ID).Msg("successfully fetched member by email")
	return &m, nil
}

func (s *store) Save(ctx context.Context, member *Member) error {
	ctx, span := tracing.StartSpan(ctx, "store.Save")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", member.ID))

	log := zerolog.Ctx(ctx).With().Str("component", "store.Save").Logger()

	if err := s.db.WithContext(ctx).Save(member).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", member.ID).Msg("failed to save member")
		return err
	}

	log.Info().Str("member_id", member.ID).Msg("member saved successfully")
	return nil
}

func (s *store) DeleteByID(ctx context.Context, memberID string) error {
	ctx, span := tracing.StartSpan(ctx, "store.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("member.id", memberID))

	log := zerolog.Ctx(ctx).With().Str("component", "store.DeleteByID").Logger()

	if err := s.db.WithContext(ctx).Delete(&Member{}, "id = ?", memberID).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error().Err(err).Str("member_id", memberID).Msg("failed to delete member by id")
		return err
	}

	log.Info().Str("member_id", memberID).Msg("member deleted successfully")
	return nil
}

package member

import (
	"context"

	"example.com/member/pkg/exception"
	"github.com/rs/zerolog"
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
	log := zerolog.Ctx(ctx)

	members, err := s.store.FindAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch members")
		return nil, err
	}

	log.Info().Int("count", len(members)).Msg("Successfully fetched all members")
	return members, nil
}

func (s *service) FindByID(ctx context.Context, memberID string) (*Member, error) {
	log := zerolog.Ctx(ctx)

	member, err := s.store.FindByID(ctx, memberID)
	if err != nil {
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
	log := zerolog.Ctx(ctx)

	member, err := s.store.FindByCode(ctx, memberCode)
	if err != nil {
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
	log := zerolog.Ctx(ctx)

	member, err := s.store.FindByCode(ctx, memberEmail)
	if err != nil {
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
	log := zerolog.Ctx(ctx)

	if err := s.store.Save(ctx, member); err != nil {
		log.Error().Err(err).Str("member_id", member.ID).Msg("Failed to save member")
		return nil, err
	}

	log.Info().Str("member_id", member.ID).Msg("Member saved successfully")
	return member, nil
}

func (s *service) DeleteByID(ctx context.Context, memberID string) error {
	log := zerolog.Ctx(ctx)

	if err := s.store.DeleteByID(ctx, memberID); err != nil {
		log.Error().Err(err).Str("member_id", memberID).Msg("Failed to delete member")
		return err
	}

	log.Info().Str("member_id", memberID).Msg("Member deleted successfully")
	return nil
}

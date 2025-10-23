package member

import (
	"context"

	"github.com/rs/zerolog"
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
	db  *gorm.DB
	log zerolog.Logger
}

func NewStore(db *gorm.DB, log zerolog.Logger) IStore {
	return &store{
		db:  db,
		log: log.With().Str("component", "member_store").Logger(),
	}
}

func (s *store) FindAll(ctx context.Context) ([]Member, error) {
	var members []Member
	if err := s.db.WithContext(ctx).Find(&members).Error; err != nil {
		s.log.Error().Err(err).Msg("Failed to find all members")
		return nil, err
	}

	s.log.Info().Int("count", len(members)).Msg("Successfully fetched all members")
	return members, nil
}

func (s *store) FindByID(ctx context.Context, memberID string) (*Member, error) {
	var m Member
	if err := s.db.WithContext(ctx).First(&m, "id = ?", memberID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.log.Warn().Str("member_id", memberID).Msg("Member not found")
			return nil, nil
		}
		s.log.Error().Err(err).Str("member_id", memberID).Msg("Failed to find member by ID")
		return nil, err
	}

	s.log.Info().Str("member_id", m.ID).Msg("Successfully fetched member by ID")
	return &m, nil
}

func (s *store) FindByCode(ctx context.Context, memberCode string) (*Member, error) {
	var m Member
	if err := s.db.WithContext(ctx).First(&m, "code = ?", memberCode).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.log.Warn().Str("member_code", memberCode).Msg("Member not found")
			return nil, nil
		}
		s.log.Error().Err(err).Str("member_code", memberCode).Msg("Failed to find member by Code")
		return nil, err
	}

	s.log.Info().Str("member_code", m.ID).Msg("Successfully fetched member by Code")
	return &m, nil
}

func (s *store) FindByEmail(ctx context.Context, memberEmail string) (*Member, error) {
	var m Member
	if err := s.db.WithContext(ctx).First(&m, "email = ?", memberEmail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.log.Warn().Str("member_email", memberEmail).Msg("Member not found")
			return nil, nil
		}
		s.log.Error().Err(err).Str("member_email", memberEmail).Msg("Failed to find member by Email")
		return nil, err
	}

	s.log.Info().Str("member_email", m.ID).Msg("Successfully fetched member by Email")
	return &m, nil
}

func (s *store) Save(ctx context.Context, member *Member) error {
	if err := s.db.WithContext(ctx).Save(member).Error; err != nil {
		s.log.Error().Err(err).Str("member_id", member.ID).Msg("Failed to save member")
		return err
	}

	s.log.Info().Str("member_id", member.ID).Msg("Member saved successfully")
	return nil
}

func (s *store) DeleteByID(ctx context.Context, memberID string) error {
	if err := s.db.WithContext(ctx).Delete(&Member{}, "id = ?", memberID).Error; err != nil {
		s.log.Error().Err(err).Str("member_id", memberID).Msg("Failed to delete member")
		return err
	}
	
	s.log.Info().Str("member_id", memberID).Msg("Member deleted successfully")
	return nil
}

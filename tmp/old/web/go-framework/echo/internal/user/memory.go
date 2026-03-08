package user

import (
	"context"
	"errors"
	"sync"

	"github.com/rs/zerolog"
)

//go:generate mockery --name=IMemory
type IMemory interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, userID string) (*User, error)
	FindByEmail(ctx context.Context, userEmail string) (*User, error)
	Save(ctx context.Context, user *User) error
	DeleteByID(ctx context.Context, userID string) error
}

type memory struct {
	mu     sync.RWMutex
	store  map[string]User
	logger zerolog.Logger
}

func NewMemoryRepository(logger zerolog.Logger) IMemory {
	return &memory{
		store:  make(map[string]User),
		logger: logger.With().Str("component", "user_memory").Logger(),
	}
}

func (m *memory) FindAll(ctx context.Context) ([]User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	users := make([]User, 0, len(m.store))
	for _, u := range m.store {
		users = append(users, u)
	}

	m.logger.Debug().
		Int("count", len(users)).
		Msg("FindAll: users")

	return users, nil
}

func (m *memory) FindByID(ctx context.Context, userID string) (*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	user, ok := m.store[userID]
	if !ok {
		m.logger.Warn().
			Str("user_id", userID).
			Msg("FindByID: user not found")
		return nil, errors.New("user not found")
	}

	m.logger.Debug().
		Str("user_id", userID).
		Str("email", user.Email).
		Msg("FindByID: user found")

	return &user, nil
}

func (m *memory) FindByEmail(ctx context.Context, userEmail string) (*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, u := range m.store {
		if u.Email == userEmail {
			m.logger.Debug().
				Str("user_id", u.ID).
				Str("email", u.Email).
				Msg("FindByEmail: user found")
			return &u, nil
		}
	}

	m.logger.Warn().
		Str("email", userEmail).
		Msg("FindByEmail: user not found")

	return nil, errors.New("user not found")
}

func (m *memory) Save(ctx context.Context, user *User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[user.ID] = *user

	m.logger.Info().
		Str("user_id", user.ID).
		Str("email", user.Email).
		Msg("user saved")

	return nil
}

func (m *memory) DeleteByID(ctx context.Context, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.store[userID]; !ok {
		m.logger.Warn().
			Str("user_id", userID).
			Msg("DeleteByID: user not found")
		return errors.New("user not found")
	}

	delete(m.store, userID)

	m.logger.Info().
		Str("user_id", userID).
		Msg("user deleted")

	return nil
}

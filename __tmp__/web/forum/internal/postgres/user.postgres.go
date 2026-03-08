package postgres

import (
	"context"
	"errors"

	"forum.id/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

// NewUserRepository creates a new userRepository instance
func NewUserRepository(db *gorm.DB, log *logrus.Logger) domain.UserRepository {
	return &userRepository{
		DB:  db,
		Log: log,
	}
}

// FetchAll retrieves all users from the database
func (r *userRepository) FetchAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		r.Log.WithError(err).Error("UserRepository.FetchAll: failed to retrieve users")
		return nil, err
	}
	return users, nil
}

// FetchByID retrieves a user by its ID
func (r *userRepository) FetchByID(ctx context.Context, ID int64) (*domain.User, error) {
	var user domain.User
	if err := r.DB.WithContext(ctx).First(&user, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.Log.WithError(err).Warnf("UserRepository.FetchByID: user with ID %d not found", ID)
			return nil, nil
		}
		r.Log.WithError(err).Errorf("UserRepository.FetchByID: failed to retrieve user with ID %d", ID)
		return nil, err
	}
	return &user, nil
}

// FetchByEmail retrieves a user by its email
func (r *userRepository) FetchByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.WithContext(ctx).
		Where("LOWER(email) = ?", email).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.Log.WithError(err).Warnf("UserRepository.FetchByEmail: user with email %s not found", email)
			return nil, nil
		}
		r.Log.WithError(err).Errorf("UserRepository.FetchByEmail: failed to retrieve user with email %s", email)
		return nil, err
	}
	return &user, nil
}

// Save creates a new user in the database
func (r *userRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.DB.WithContext(ctx).Create(user).Error; err != nil {
		r.Log.WithError(err).Error("UserRepository.Save: failed to create user")
		return nil, err
	}
	r.Log.Infof("UserRepository.Save: user %s created successfully", user.Username)
	return user, nil
}

// Update updates an existing user in the database
func (r *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.DB.WithContext(ctx).Save(user).Error; err != nil {
		r.Log.WithError(err).Errorf("UserRepository.Update: failed to update user %s", user.Username)
		return nil, err
	}
	r.Log.Infof("UserRepository.Update: user %s updated successfully", user.Username)
	return user, nil
}

// Delete removes a user from the database
func (r *userRepository) Delete(ctx context.Context, user *domain.User) error {
	if err := r.DB.WithContext(ctx).Delete(user).Error; err != nil {
		r.Log.WithError(err).Errorf("UserRepository.Delete: failed to delete user %s", user.Username)
		return err
	}
	r.Log.Infof("UserRepository.Delete: user %s deleted successfully", user.Username)
	return nil
}

// ExistUsernameIgnoreCase checks if a username already exists (case-insensitive)
func (r *userRepository) ExistUsernameIgnoreCase(ctx context.Context, username string) (bool, error) {
	var count int64
	if err := r.DB.WithContext(ctx).
		Model(&domain.User{}).
		Where("LOWER(username) = ?", username).
		Count(&count).Error; err != nil {
		r.Log.WithError(err).Errorf("UserRepository.ExistUsernameIgnoreCase: failed to check username %s", username)
		return false, err
	}
	if count > 0 {
		r.Log.Warnf("UserRepository.ExistUsernameIgnoreCase: username %s already exists", username)
	}
	return count > 0, nil
}

// ExistEmailIgnoreCase checks if an email already exists (case-insensitive)
func (r *userRepository) ExistEmailIgnoreCase(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.DB.WithContext(ctx).
		Model(&domain.User{}).
		Where("LOWER(email) = ?", email).
		Count(&count).Error; err != nil {
		r.Log.WithError(err).Errorf("UserRepository.ExistEmailIgnoreCase: failed to check email %s", email)
		return false, err
	}
	if count > 0 {
		r.Log.Warnf("UserRepository.ExistEmailIgnoreCase: email %s already exists", email)
	}
	return count > 0, nil
}

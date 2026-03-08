package postgres

import (
	"context"
	"errors"

	"forum.id/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type threadRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

// NewThreadRepository creates a new threadRepository instance
func NewThreadRepository(db *gorm.DB, log *logrus.Logger) domain.ThreadRepository {
	return &threadRepository{
		DB:  db,
		Log: log,
	}
}

// FetchAll retrieves all threads from the database with their related user, comments, and votes
func (r *threadRepository) FetchAll(ctx context.Context) ([]*domain.Thread, error) {
	var threads []*domain.Thread

	// No need for transaction here as it's a simple query
	if err := r.DB.WithContext(ctx).
		Preload("User").
		Preload("Comments").
		Preload("Votes.User").
		Find(&threads).Error; err != nil {
		r.Log.WithError(err).Error("ThreadRepository.FetchAll: failed to retrieve threads")
		return nil, err
	}

	return threads, nil
}

// FetchAllByUserID retrieves all threads created by a specific user with related comments and votes
func (r *threadRepository) FetchAllByUserID(ctx context.Context, userID int64) ([]*domain.Thread, error) {
	var threads []*domain.Thread

	// Simple query, no transaction needed
	if err := r.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("User").
		Preload("Comments").
		Preload("Votes.User").
		Find(&threads).Error; err != nil {
		r.Log.WithError(err).Errorf("ThreadRepository.FetchAllByUserID: failed to retrieve threads for user ID %d", userID)
		return nil, err
	}

	return threads, nil
}

// FetchByID retrieves a thread by its ID with related user, comments, and votes
func (r *threadRepository) FetchByID(ctx context.Context, ID int64) (*domain.Thread, error) {
	var thread domain.Thread

	// Simple query, no transaction needed
	if err := r.DB.WithContext(ctx).
		Preload("User").
		Preload("Comments.User").
		Preload("Votes.User").
		Preload("Comments.Votes.User").
		First(&thread, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.Log.WithError(err).Warnf("ThreadRepository.FetchByID: thread with ID %d not found", ID)
			return nil, nil
		}
		r.Log.WithError(err).Errorf("ThreadRepository.FetchByID: failed to retrieve thread with ID %d", ID)
		return nil, err
	}

	return &thread, nil
}

// Save creates a new thread in the database
func (r *threadRepository) Save(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("ThreadRepository.Save: failed to start transaction")
		return nil, tx.Error
	}

	// Insert the new thread
	if err := tx.Create(thread).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Error("ThreadRepository.Save: failed to create thread")
		return nil, err
	}

	// Preload associated fields
	if err := tx.Preload("User").
		Preload("Comments").
		Preload("Votes.User").
		Find(thread).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Error("ThreadRepository.Save: failed to preload associated fields")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("ThreadRepository.Save: failed to commit transaction")
		return nil, err
	}

	r.Log.Infof("ThreadRepository.Save: thread %s created successfully", thread.Title)
	return thread, nil
}

// Update updates an existing thread in the database
func (r *threadRepository) Update(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("ThreadRepository.Update: failed to start transaction")
		return nil, tx.Error
	}

	// Update the thread
	if err := tx.Save(thread).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Errorf("ThreadRepository.Update: failed to update thread %s", thread.Title)
		return nil, err
	}

	// Preload associated fields after update
	if err := tx.Preload("User").
		Preload("Comments").
		Preload("Votes.User").
		Find(thread).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Errorf("ThreadRepository.Update: failed to preload associated fields for thread %s", thread.Title)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("ThreadRepository.Update: failed to commit transaction")
		return nil, err
	}

	r.Log.Infof("ThreadRepository.Update: thread %s updated successfully", thread.Title)
	return thread, nil
}

// Delete removes a thread from the database
func (r *threadRepository) Delete(ctx context.Context, thread *domain.Thread) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("ThreadRepository.Delete: failed to start transaction")
		return tx.Error
	}

	// Delete the thread
	if err := tx.Delete(thread).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Errorf("ThreadRepository.Delete: failed to delete thread %s", thread.Title)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("ThreadRepository.Delete: failed to commit transaction")
		return err
	}

	r.Log.Infof("ThreadRepository.Delete: thread %s deleted successfully", thread.Title)
	return nil
}

// ExistByID checks if a thread exists by its ID
func (r *threadRepository) ExistByID(ctx context.Context, ID int64) (bool, error) {
	var count int64

	// Simple query, no transaction needed
	if err := r.DB.WithContext(ctx).
		Model(&domain.Thread{}).
		Where("id = ?", ID).
		Count(&count).Error; err != nil {
		r.Log.WithError(err).Errorf("ThreadRepository.ExistByID: failed to check existence of thread with ID %d", ID)
		return false, err
	}

	return count > 0, nil
}

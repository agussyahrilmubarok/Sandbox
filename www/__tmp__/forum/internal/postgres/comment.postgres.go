package postgres

import (
	"context"
	"errors"

	"forum.id/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type commentRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

// NewCommentRepository creates a new commentRepository instance
func NewCommentRepository(db *gorm.DB, log *logrus.Logger) domain.CommentRepository {
	return &commentRepository{
		DB:  db,
		Log: log,
	}
}

// FetchAll retrieves all comments with related user, thread, and votes
func (r *commentRepository) FetchAll(ctx context.Context) ([]*domain.Comment, error) {
	var comments []*domain.Comment

	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("CommentRepository.FetchAll: failed to start transaction")
		return nil, tx.Error
	}

	if err := tx.Preload("User").
		Preload("Thread").
		Preload("Votes.User").
		Find(&comments).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Error("CommentRepository.FetchAll: failed to retrieve comments")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("CommentRepository.FetchAll: failed to commit transaction")
		return nil, err
	}

	return comments, nil
}

// FetchAllByUserID retrieves all comments for a specific thread with related user and votes
func (r *commentRepository) FetchAllByUserID(ctx context.Context, userID int64) ([]*domain.Comment, error) {
	var comments []*domain.Comment

	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("CommentRepository.FetchAllByUserID: failed to start transaction")
		return nil, tx.Error
	}

	if err := tx.Where("user_id = ?", userID).
		Preload("User").
		Preload("Votes.User").
		Find(&comments).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Errorf("CommentRepository.FetchAllByUserID: failed to retrieve comments for thread ID %d", userID)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("CommentRepository.FetchAllByUserID: failed to commit transaction")
		return nil, err
	}

	return comments, nil
}

// FetchAllByThreadID retrieves all comments for a specific thread with related user and votes
func (r *commentRepository) FetchAllByThreadID(ctx context.Context, threadID int64) ([]*domain.Comment, error) {
	var comments []*domain.Comment

	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("CommentRepository.FetchAllByThreadID: failed to start transaction")
		return nil, tx.Error
	}

	if err := tx.Where("thread_id = ?", threadID).
		Preload("User").
		Preload("Votes.User").
		Find(&comments).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Errorf("CommentRepository.FetchAllByThreadID: failed to retrieve comments for thread ID %d", threadID)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("CommentRepository.FetchAllByThreadID: failed to commit transaction")
		return nil, err
	}

	return comments, nil
}

// FetchByID retrieves a comment by its ID with related user, thread, and votes
func (r *commentRepository) FetchByID(ctx context.Context, ID int64) (*domain.Comment, error) {
	var comment domain.Comment

	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("CommentRepository.FetchByID: failed to start transaction")
		return nil, tx.Error
	}

	if err := tx.Preload("User").
		Preload("Thread").
		Preload("Votes.User").
		First(&comment, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			r.Log.WithError(err).Warnf("CommentRepository.FetchByID: comment with ID %d not found", ID)
			return nil, nil
		}
		tx.Rollback()
		r.Log.WithError(err).Errorf("CommentRepository.FetchByID: failed to retrieve comment with ID %d", ID)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("CommentRepository.FetchByID: failed to commit transaction")
		return nil, err
	}

	return &comment, nil
}

// Save creates a new comment in the database
func (r *commentRepository) Save(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("CommentRepository.Save: failed to start transaction")
		return nil, tx.Error
	}

	if err := tx.Create(comment).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Error("CommentRepository.Save: failed to create comment")
		return nil, err
	}

	if err := tx.Preload("User").
		Preload("Thread").
		Preload("Votes.User").
		Find(comment).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Error("CommentRepository.Save: failed to preload related data")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("CommentRepository.Save: failed to commit transaction")
		return nil, err
	}

	r.Log.Infof("CommentRepository.Save: comment created successfully by user %d", comment.UserID)
	return comment, nil
}

// Update updates an existing comment in the database
func (r *commentRepository) Update(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("CommentRepository.Update: failed to start transaction")
		return nil, tx.Error
	}

	if err := tx.Save(comment).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Errorf("CommentRepository.Update: failed to update comment %d", comment.ID)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("CommentRepository.Update: failed to commit transaction")
		return nil, err
	}

	r.Log.Infof("CommentRepository.Update: comment %d updated successfully", comment.ID)
	return comment, nil
}

// Delete removes a comment from the database
func (r *commentRepository) Delete(ctx context.Context, comment *domain.Comment) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("CommentRepository.Delete: failed to start transaction")
		return tx.Error
	}

	if err := tx.Delete(comment).Error; err != nil {
		tx.Rollback()
		r.Log.WithError(err).Errorf("CommentRepository.Delete: failed to delete comment %d", comment.ID)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("CommentRepository.Delete: failed to commit transaction")
		return err
	}

	r.Log.Infof("CommentRepository.Delete: comment %d deleted successfully", comment.ID)
	return nil
}

// ExistByID checks if a comment exists by its ID
func (r *commentRepository) ExistByID(ctx context.Context, ID int64) (bool, error) {
	var count int64

	// Simple query, no transaction needed
	if err := r.DB.WithContext(ctx).
		Model(&domain.Comment{}).
		Where("id = ?", ID).
		Count(&count).Error; err != nil {
		r.Log.WithError(err).Errorf("CommentRepository.ExistByID: failed to check existence of comment with ID %d", ID)
		return false, err
	}

	return count > 0, nil
}

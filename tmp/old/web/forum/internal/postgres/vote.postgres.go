package postgres

import (
	"context"
	"errors"

	"forum.id/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type voteRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

// NewVoteRepository creates a new voteRepository instance
func NewVoteRepository(db *gorm.DB, log *logrus.Logger) domain.VoteRepository {
	return &voteRepository{
		DB:  db,
		Log: log,
	}
}

// SetVoteThread sets a vote for a thread (creates or updates)
func (r *voteRepository) SetVoteThread(ctx context.Context, voteThread *domain.VoteThread) (*domain.VoteThread, error) {
	tx := r.DB.WithContext(ctx).Begin()

	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("VoteRepository.SetVoteThread: failed to start transaction")
		return nil, tx.Error
	}

	// Check if a vote by the user on this thread already exists
	var existingVote domain.VoteThread
	err := tx.Where("user_id = ? AND thread_id = ?", voteThread.UserID, voteThread.ThreadID).First(&existingVote).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		r.Log.WithError(err).Error("VoteRepository.SetVoteThread: failed to check existing vote")
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// If not found, create a new vote
		if err := tx.Create(voteThread).Error; err != nil {
			tx.Rollback()
			r.Log.WithError(err).Error("VoteRepository.SetVoteThread: failed to create new vote")
			return nil, err
		}
	} else {
		// If vote exists, update the vote type
		existingVote.VoteType = voteThread.VoteType
		if err := tx.Save(&existingVote).Error; err != nil {
			tx.Rollback()
			r.Log.WithError(err).Error("VoteRepository.SetVoteThread: failed to update existing vote")
			return nil, err
		}
		voteThread = &existingVote // Return the updated vote
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("VoteRepository.SetVoteThread: failed to commit transaction")
		return nil, err
	}

	r.Log.Infof("VoteRepository.SetVoteThread: vote for thread %d by user %d successfully set", voteThread.ThreadID, voteThread.UserID)
	return voteThread, nil
}

// SetVoteComment sets a vote for a comment (creates or updates)
func (r *voteRepository) SetVoteComment(ctx context.Context, voteComment *domain.VoteComment) (*domain.VoteComment, error) {
	tx := r.DB.WithContext(ctx).Begin()

	if tx.Error != nil {
		r.Log.WithError(tx.Error).Error("VoteRepository.SetVoteComment: failed to start transaction")
		return nil, tx.Error
	}

	// Check if a vote by the user on this comment already exists
	var existingVote domain.VoteComment
	err := tx.Where("user_id = ? AND comment_id = ?", voteComment.UserID, voteComment.CommentID).First(&existingVote).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		r.Log.WithError(err).Error("VoteRepository.SetVoteComment: failed to check existing vote")
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// If not found, create a new vote
		if err := tx.Create(voteComment).Error; err != nil {
			tx.Rollback()
			r.Log.WithError(err).Error("VoteRepository.SetVoteComment: failed to create new vote")
			return nil, err
		}
	} else {
		// If vote exists, update the vote type
		existingVote.VoteType = voteComment.VoteType
		if err := tx.Save(&existingVote).Error; err != nil {
			tx.Rollback()
			r.Log.WithError(err).Error("VoteRepository.SetVoteComment: failed to update existing vote")
			return nil, err
		}
		voteComment = &existingVote // Return the updated vote
	}

	if err := tx.Commit().Error; err != nil {
		r.Log.WithError(err).Error("VoteRepository.SetVoteComment: failed to commit transaction")
		return nil, err
	}

	r.Log.Infof("VoteRepository.SetVoteComment: vote for comment %d by user %d successfully set", voteComment.CommentID, voteComment.UserID)
	return voteComment, nil
}

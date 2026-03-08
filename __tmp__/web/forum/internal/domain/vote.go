package domain

import "time"

type VoteType string

const (
	UpVote      VoteType = "UP_VOTE"
	DownVote    VoteType = "DOWN_VOTE"
	NeutralVote VoteType = "NEUTRAL_VOTE"
)

// Vote thread entity
type VoteThread struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	VoteType  string    `json:"voteType" gorm:"not null"`
	UserID    int64     `json:"userId" gorm:"not null"`
	ThreadID  int64     `json:"threadId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:current_timestamp"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Thread    Thread    `json:"thread" gorm:"foreignKey:ThreadID;references:ID"`
}

// Vote comment entity
type VoteComment struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	VoteType  string    `json:"voteType" gorm:"not null"`
	UserID    int64     `json:"userId" gorm:"not null"`
	CommentID int64     `json:"commentId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:current_timestamp"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Comment   Comment   `json:"comment" gorm:"foreignKey:CommentID;references:ID"`
}

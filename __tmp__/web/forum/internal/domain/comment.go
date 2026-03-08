package domain

import "time"

// Comment entity
type Comment struct {
	ID        int64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   string        `json:"content" gorm:"not null"`
	UserID    int64         `json:"userId" gorm:"not null"`
	ThreadID  int64         `json:"threadId" gorm:"not null"`
	CreatedAt time.Time     `json:"createdAt" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time     `json:"updatedAt" gorm:"not null;default:current_timestamp"`
	User      User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Thread    Thread        `json:"thread" gorm:"foreignKey:ThreadID;references:ID"`
	Votes     []VoteComment `json:"votes" gorm:"foreignKey:CommentID"`
}

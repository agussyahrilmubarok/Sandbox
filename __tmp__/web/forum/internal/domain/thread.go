package domain

import "time"

// Thread entity
type Thread struct {
	ID        int64        `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string       `json:"title" gorm:"not null"`
	Body      string       `json:"body" gorm:"not null"`
	Category  string       `json:"category" gorm:"not null"`
	UserID    int64        `json:"userId" gorm:"not null"`
	CreatedAt time.Time    `json:"createdAt" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"not null;default:current_timestamp"`
	User      User         `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Comments  []Comment    `json:"comments" gorm:"foreignKey:ThreadID"`
	Votes     []VoteThread `json:"votes" gorm:"foreignKey:ThreadID"`
}

package domain

import (
	"time"

	"forum.id/pkg/password"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleUser      UserRole = "ROLE_USER"
	RoleModerator UserRole = "ROLE_MODERATOR"
)

// User entity
type User struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:current_timestamp"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Password = password.Hash(user.Password)
	if user.Role == "" {
		user.Role = string(RoleUser)
	}
	return nil
}

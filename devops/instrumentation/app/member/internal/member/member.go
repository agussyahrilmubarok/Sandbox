package member

import "time"

type Member struct {
	ID        string    `json:"id" gorm:"type:char(36);primaryKey"`
	Code      string    `json:"code" gorm:"type:varchar(8);unique;not null"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

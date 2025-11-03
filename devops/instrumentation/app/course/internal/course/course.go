package course

import "time"

type Course struct {
	ID            string    `json:"id" gorm:"type:char(36);primaryKey"`
	Code          string    `json:"code" gorm:"type:varchar(20);unique;not null"`
	Name          string    `json:"name" gorm:"type:varchar(100);not null"`
	Price         float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	StartDate     time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate       time.Time `json:"end_date" gorm:"type:date;not null"`
	SeatAvailable int       `json:"seat_available" gorm:"type:int;not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CourseCodeRequest struct {
	Code string `json:"code"`
}

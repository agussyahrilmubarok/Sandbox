package booking

import "time"

type Booking struct {
	ID          string        `json:"id" gorm:"type:char(36);primaryKey"`
	MemberCode  string        `json:"member_code" gorm:"type:varchar(8);not null"`
	CourseCode  string        `json:"course_code" gorm:"type:varchar(8);not null"`
	BookingDate time.Time     `json:"booking_date" gorm:"type:date;not null"`
	Status      BookingStatus `json:"status" gorm:"type:int;not null;default:0"`
	Notes       string        `json:"notes" gorm:"type:text"`
	CreatedAt   time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

type BookingStatus int

const (
	BookingStatusPending BookingStatus = iota
	BookingStatusConfirmed
	BookingStatusCancelled
)

func (s BookingStatus) String() string {
	return [...]string{"pending", "confirmed", "cancelled"}[s]
}

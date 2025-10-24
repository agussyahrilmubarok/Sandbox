package booking

import "context"

//go:generate mockery --name=IService
type IService interface {
	Booking(ctx context.Context, request BookingRequest) (*Booking, error)
}

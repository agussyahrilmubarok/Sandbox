package domain

import "time"

// Shipping represents the shipping details for an order
type Shipping struct {
	ShippingID       string    `json:"shipping_id"`
	OrderID          string    `json:"order_id"`
	ShippingDate     time.Time `json:"shipping_date"`
	ShippingMethod   string    `json:"shipping_method"`
	TrackingNumber   string    `json:"tracking_number"`
	ShippingStatus   string    `json:"shipping_status"`
	EstimatedArrival time.Time `json:"estimated_arrival"`
}

package domain

import "time"

// Payment represents a payment for an order
type Payment struct {
	PaymentID     string    `json:"payment_id"`
	OrderID       string    `json:"order_id"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	AmountPaid    float64   `json:"amount_paid"`
	PaymentStatus string    `json:"payment_status"`
	TransactionID string    `json:"transaction_id"`
}

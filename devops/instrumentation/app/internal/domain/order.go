package domain

import "time"

// Order represents a customer order
type Order struct {
	OrderID         string      `json:"order_id"`
	CustomerID      string      `json:"customer_id"`
	OrderDate       time.Time   `json:"order_date"`
	Status          string      `json:"status"`
	TotalAmount     float64     `json:"total_amount"`
	ShippingAddress string      `json:"shipping_address"`
	Items           []OrderItem `json:"items"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}

package domain

// Inventory represents the details of a product in the inventory
type Inventory struct {
	ProductID       string  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	Quantity        int     `json:"quantity"`
	UnitPrice       float64 `json:"unit_price"`
	TotalStockValue float64 `json:"total_stock_value"`
}

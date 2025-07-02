package dto

type TransactionResponse struct {
	ID           uint    `json:"id"`
	ProductID    uint    `json:"product_id"`
	CustomerID   uint    `json:"customer_id"`
	Quantity     int     `json:"quantity"`
	TotalPrice   float64 `json:"total_price"`
	FreeShipping bool    `json:"free_shipping"`
	Discount     float64 `json:"discount"`
}

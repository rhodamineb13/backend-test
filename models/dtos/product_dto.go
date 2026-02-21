package dtos

import "time"

type Product struct {
	Id            uint       `json:"id,omitempty"`
	Name          string     `json:"name"`
	Description   string     `json:"description,omitempty"`
	Price         float32    `json:"price"`
	CategoryID    uint       `json:"category_id"`
	StockQuantity uint       `json:"stock_quantity"`
	IsAvailable   bool       `json:"is_available"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Category      Category   `json:"category"`
}

type ProductResponse struct {
	NumberOfProducts uint      `json:"number_of_products"`
	TotalStock       uint      `json:"total_stock"`
	AveragePrice     uint      `json:"average_price"`
	ProductDetails   []Product `json:"product_details"`
}

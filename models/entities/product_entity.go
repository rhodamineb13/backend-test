package entities

import "time"

type Product struct {
	Id            uint      `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey;autoIncrement;index:IDX_category_id"`
	Name          string    `gorm:"column:name;type:VARCHAR;not null"`
	Description   string    `gorm:"column:description;type:VARCHAR;not null"`
	Price         float32   `gorm:"column:price;type:DECIMAL(10,2);not null"`
	CategoryId    uint      `gorm:"column:category_id;type:BIGINT;not null;index:IDX_category_id;priority:1"`
	StockQuantity uint      `gorm:"column:stock_quantity;type:INT;not null"`
	IsAvailable   bool      `gorm:"column:is_available;type:BOOLEAN;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;type:TIMESTAMP;not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:TIMESTAMP;default: NULL"`

	Category Category `gorm:"foreignKey:CategoryId"`
}

type DashboardProduct struct {
	TotalProducts int64
	TotalStock    int64
	AveragePrice  float32
	Products      []Product
}

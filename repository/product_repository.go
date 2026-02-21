package repository

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/rhodamineb13/backend-test/models/entities"
)

type IProductRepository interface {
	ListProducts(ctx context.Context, name string, category uint, price float32, stock_quantity uint) ([]entities.Product, error)
	GetProductByID(ctx context.Context, id uint) (*entities.Product, error)
	InsertNewProduct(ctx context.Context, data entities.Product) error
	UpdateProduct(ctx context.Context, id uint, data entities.Product) error
	DeleteProduct(ctx context.Context, id uint) error
}

type productRepository struct {
	db *gorm.DB
	rc *redis.Client
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{}
}

func (ps *productRepository) ListProducts(ctx context.Context, name string, category uint, price float32, stock_quantity uint) ([]entities.Product, error) {
	key := "LIST_PRODUCTS"
	var products []entities.Product
	b, err := ps.rc.Get(ctx, key).Bytes()
	if err == nil {
		if errJSON := json.Unmarshal(b, &products); errJSON != nil {
			return nil, errJSON
		}
		return products, nil
	} else {
		query := `SELECT p.name, p.description, p.price, p.category_id, p.stock_quantity
		FROM products AS p`
		if name != "" {
			query += `WHERE p.name ILIKE ?`
		}

		if category != 0 {
			query += `AND p.category_id = ?`
		}

		if price != 0.0 {
			query += `AND p.price > ?`
		}

		if stock_quantity != 0 {
			query += `AND p.stock_quantity > ?`
		}

		query += `INNER JOIN
		SELECT c.id, c.category_name
		FROM categories as c
		`
		errSQL := ps.db.Raw(query, name, category, price, stock_quantity).Find(&products).Error
		if errSQL != nil {
			return nil, errSQL
		}

		return products, nil

	}
}

func (ps *productRepository) GetProductByID(ctx context.Context, id uint) (*entities.Product, error) {
	var product entities.Product
	if err := ps.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *productRepository) InsertNewProduct(ctx context.Context, data entities.Product) error {
	return ps.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&data).Error
	})
}

func (ps *productRepository) UpdateProduct(ctx context.Context, id uint, data entities.Product) error {
	return ps.db.Transaction(func(tx *gorm.DB) error {
		return tx.Save(&data).Error
	})
}

func (ps *productRepository) DeleteProduct(ctx context.Context, id uint) error {
	return ps.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entities.Product{}).Error
	})
}

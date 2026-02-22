package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/rhodamineb13/backend-test/models/entities"
	"github.com/rhodamineb13/backend-test/utils"
)

type IProductRepository interface {
	ListProducts(ctx context.Context, name string, category uint, price float32, stockQuantity uint, limit, page uint) (*entities.DashboardProduct, error)
	GetProductByID(ctx context.Context, id uint) (*entities.Product, error)
	InsertNewProduct(ctx context.Context, data entities.Product) error
	UpdateProduct(ctx context.Context, id uint, data entities.Product) error
	DeleteProduct(ctx context.Context, id uint) error
}

type productRepository struct {
	db *gorm.DB
	rc *redis.Client
}

func NewProductRepository(db *gorm.DB, rc *redis.Client) IProductRepository {
	return &productRepository{db: db, rc: rc}
}

func (pr *productRepository) ListProducts(ctx context.Context, name string, category uint, price float32, stockQuantity uint, limit uint, page uint) (*entities.DashboardProduct, error) {
	key := fmt.Sprintf("LIST_PRODUCTS:%s:%d:%f:%d:%d:%d", name, category, price, stockQuantity, limit, page)
	var dashboard entities.DashboardProduct

	// Check from cache first, then get from the database
	b, err := pr.rc.Get(ctx, key).Bytes()
	if err == nil {
		if errJSON := json.Unmarshal(b, &dashboard); errJSON != nil {
			return nil, errJSON
		}
		return &dashboard, nil
	}

	baseQuery := pr.db.Model(&entities.Product{})
	if name != "" {
		baseQuery = baseQuery.Where("LOWER(name) LIKE LOWER(?)", "'%"+name+"%'")
	}
	if category != 0 {
		baseQuery = baseQuery.Where("category_id = ?", category)
	}
	if price != 0.0 {
		baseQuery = baseQuery.Where("price > ?", price)
	}
	if stockQuantity != 0 {
		baseQuery = baseQuery.Where("stock_quantity > ?", stockQuantity)
	}

	// Session for aggregations, to prevent original query from being affected
	err = baseQuery.Session(&gorm.Session{}).
		Select(`
			COUNT(1) as total_products,
			COALESCE(SUM(stock_quantity), 0) as total_stock,
			COALESCE(AVG(price), 0.0) as average_price
		`).
		Scan(&dashboard).Error
	if err != nil {
		return nil, err
	}

	preloadedQuery := baseQuery.Preload("Category")

	// Coerce limit into the maximum value (50) if it's greater
	if limit != 0 {
		if limit > 100 {
			limit = 100
		}
		preloadedQuery.Limit(int(limit))
	}
	// Coerce page into the maximum value (2000) if it's greater
	if page != 0 {
		if page > 2000 {
			page = 2000
		}
		offset := (page - 1) * limit
		preloadedQuery.Offset(int(offset))
	}

	err = preloadedQuery.Find(&dashboard.Products).Error
	if err != nil {
		return nil, err
	}

	pr.rc.Set(ctx, key, dashboard, 0)

	return &dashboard, nil
}

func (pr *productRepository) GetProductByID(ctx context.Context, id uint) (*entities.Product, error) {
	var product entities.Product
	if err := pr.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *productRepository) InsertNewProduct(ctx context.Context, data entities.Product) error {
	return pr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		// Deletes cache since there's change in the database
		if err := utils.DeleteRedisKeyMatchingPattern(ctx, "LIST_PRODUCTS", pr.rc); err != nil {
			return err
		}
		return nil
	})
}

func (pr *productRepository) UpdateProduct(ctx context.Context, id uint, data entities.Product) error {
	return pr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&data).Error; err != nil {
			return err
		}

		// Deletes cache since there's change in the database
		if err := utils.DeleteRedisKeyMatchingPattern(ctx, "LIST_PRODUCTS", pr.rc); err != nil {
			return err
		}
		return nil
	})
}

func (pr *productRepository) DeleteProduct(ctx context.Context, id uint) error {
	return pr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&entities.Product{}).Error; err != nil {
			return err
		}

		// Deletes cache since there's change in the database
		if err := utils.DeleteRedisKeyMatchingPattern(ctx, "LIST_PRODUCTS", pr.rc); err != nil {
			return err
		}
		return nil
	})
}

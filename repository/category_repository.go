package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/rhodamineb13/backend-test/models/entities"
)

type ICategoryRepository interface {
	ListCategories(ctx context.Context) ([]entities.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (*entities.Category, error)
	InsertNewCategory(ctx context.Context, data entities.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{db: db}
}

func (cs *categoryRepository) ListCategories(ctx context.Context) ([]entities.Category, error) {
	var categories []entities.Category
	err := cs.db.Find(&categories).Error
	return categories, err
}

func (cs *categoryRepository) GetCategoryByID(ctx context.Context, id uint) (*entities.Category, error) {
	var category entities.Category
	err := cs.db.Where("id = ?").First(&category).Error
	return &category, err
}

func (cs *categoryRepository) InsertNewCategory(ctx context.Context, data entities.Category) error {
	return cs.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&data).Error
	})
}

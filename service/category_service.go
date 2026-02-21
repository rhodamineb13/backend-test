package service

import (
	"context"

	"github.com/rhodamineb13/backend-test/models/dtos"
	"github.com/rhodamineb13/backend-test/models/entities"
	"github.com/rhodamineb13/backend-test/repository"
)

type ICategoryService interface {
	ListCategories(ctx context.Context) ([]dtos.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (*dtos.Category, error)
	InsertNewCategory(ctx context.Context, data dtos.Category) error
}

type categoryService struct {
	categoryRepo repository.ICategoryRepository
}

func NewCategoryService(categoryRepo repository.ICategoryRepository) ICategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (cs *categoryService) ListCategories(ctx context.Context) ([]dtos.Category, error) {
	categories, err := cs.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	categoriesRes := make([]dtos.Category, len(categories))
	for i := range categories {
		categoriesRes[i] = dtos.Category{
			Id:   categories[i].Id,
			Name: categories[i].Name,
		}
	}
	return categoriesRes, nil
}

func (cs *categoryService) GetCategoryByID(ctx context.Context, id uint) (*dtos.Category, error) {
	category, err := cs.categoryRepo.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dtos.Category{
		Id:   category.Id,
		Name: category.Name,
	}, nil
}

func (cs *categoryService) InsertNewCategory(ctx context.Context, data dtos.Category) error {
	input := entities.Category{
		Name:        data.Name,
		Description: data.Description,
	}

	if err := cs.categoryRepo.InsertNewCategory(ctx, input); err != nil {
		return err
	}
	return nil
}

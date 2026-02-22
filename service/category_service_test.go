package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	customerrors "github.com/rhodamineb13/backend-test/errors"
	"github.com/rhodamineb13/backend-test/mocks"
	"github.com/rhodamineb13/backend-test/models/dtos"
	"github.com/rhodamineb13/backend-test/models/entities"
	"github.com/rhodamineb13/backend-test/service"
)

func TestGetCategories(t *testing.T) {
	mockCategory := []entities.Category{
		{
			Id:   uint(1),
			Name: "Tools",
		},
		{
			Id:   uint(2),
			Name: "Beverages",
		},
	}
	ctx := context.Background()
	t.Run("Should get all products if there's no error", func(t *testing.T) {
		mockCategoryRepo := new(mocks.ICategoryRepository)
		mockCategoryRepo.On("ListCategories", ctx).Return(mockCategory, nil)

		service := service.NewCategoryService(mockCategoryRepo)

		res, err := service.ListCategories(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.IsType(t, res, []dtos.Category{})
	})

	t.Run("should return nil data and non-nil error if there's error from database", func(t *testing.T) {
		mockCategoryRepo := new(mocks.ICategoryRepository)
		mockCategoryRepo.On("ListCategories", ctx).Return(nil, customerrors.ErrBadRequest(fmt.Errorf("Internal server error mock error")))

		service := service.NewCategoryService(mockCategoryRepo)

		res, err := service.ListCategories(ctx)

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorAs(t, err, &customerrors.Errors{})
	})
}

func TestGetCategoriesByID(t *testing.T) {
	mockRepo := new(mocks.ICategoryRepository)
	ps := service.NewCategoryService(mockRepo)
	ctx := context.Background()

	t.Run("should return product dto when found", func(t *testing.T) {
		id := uint(1)
		mockEntity := &entities.Category{
			Id: id, Name: "Footwear",
		}

		mockRepo.On("GetCategoryByID", ctx, id).Return(mockEntity, nil).Once()

		res, err := ps.GetCategoryByID(ctx, id)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, uint(1), res.Id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return custom not found error", func(t *testing.T) {
		id := uint(99)
		mockRepo.On("GetCategoryByID", ctx, id).Return(nil, gorm.ErrRecordNotFound).Once()

		res, err := ps.GetCategoryByID(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, res)

		mockRepo.AssertExpectations(t)
	})
}

func TestInsertNewCategory(t *testing.T) {
	mockRepo := new(mocks.ICategoryRepository)
	ps := service.NewCategoryService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		input := dtos.Category{Name: "Home appliances", Description: "Home appliances includes TV, washing machines, and fridges"}

		mockRepo.On("InsertNewCategory", ctx, mock.MatchedBy(func(p entities.Product) bool {
			return p.Name == "Laptop" && p.CategoryId == 5
		})).Return(nil).Once()

		err := ps.InsertNewCategory(ctx, input)
		assert.NoError(t, err)
	})
}

package service_test

import (
	"context"
	"errors"
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

func TestGetProducts(t *testing.T) {
	mockProducts := &entities.DashboardProduct{
		TotalStock:    10,
		TotalProducts: 1,
		AveragePrice:  3.3,
		Products: []entities.Product{
			{
				Id:            3,
				Name:          "White Fleet Shoes",
				Description:   "Man footwear, size 38 (EU)",
				StockQuantity: 10,
				CategoryId:    5,
				Category: entities.Category{
					Id:          7,
					Name:        "Footwears",
					Description: "Sandals, shoes, etc.",
				},
			},
		},
	}

	ctx := context.Background()

	t.Run("Should get all products if there's no error", func(t *testing.T) {
		mockProductRepo := new(mocks.IProductRepository)
		mockProductRepo.On("ListProducts", ctx, "", uint(0), float32(0.0), uint(0), uint(0), uint(0)).Return(mockProducts, nil)

		service := service.NewProductService(mockProductRepo)

		res, err := service.ListProducts(ctx, "", uint(0), float32(0.0), uint(0), uint(0), uint(0))

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.IsType(t, res, &dtos.ProductResponse{})
	})

	t.Run("should return nil data and non-nil error if there's error from database", func(t *testing.T) {
		mockProductRepo := new(mocks.IProductRepository)
		mockProductRepo.On("ListProducts", ctx, "", uint(0), float32(0.0), uint(0), uint(0), uint(0)).Return(nil, customerrors.ErrBadRequest(fmt.Errorf("Internal server error mock error")))

		service := service.NewProductService(mockProductRepo)

		res, err := service.ListProducts(ctx, "", uint(0), float32(0.0), uint(0), uint(0), uint(0))

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorAs(t, err, &customerrors.Errors{})
	})
}

func TestGetProductByID(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	ps := service.NewProductService(mockRepo)
	ctx := context.Background()

	t.Run("should return product dto when found", func(t *testing.T) {
		id := uint(1)
		mockEntity := &entities.Product{
			Id: id, Name: "Gaming Mouse", Price: 50.0,
			Category: entities.Category{Id: 10, Name: "Peripherals"},
		}

		mockRepo.On("GetProductByID", ctx, id).Return(mockEntity, nil).Once()

		res, err := ps.GetProductByID(ctx, id)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "Gaming Mouse", res.Name)
		assert.Equal(t, uint(10), res.Category.Id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return custom not found error", func(t *testing.T) {
		id := uint(99)
		mockRepo.On("GetProductByID", ctx, id).Return(nil, gorm.ErrRecordNotFound).Once()

		res, err := ps.GetProductByID(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, res)

		mockRepo.AssertExpectations(t)
	})
}

func TestInsertNewProduct(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	ps := service.NewProductService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		input := dtos.Product{Name: "Laptop", Price: 1000.0, CategoryID: 5}

		mockRepo.On("InsertNewProduct", ctx, mock.MatchedBy(func(p entities.Product) bool {
			return p.Name == "Laptop" && p.CategoryId == 5
		})).Return(nil).Once()

		err := ps.InsertNewProduct(ctx, input)
		assert.NoError(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	ps := service.NewProductService(mockRepo)
	ctx := context.Background()

	t.Run("should only update provided fields", func(t *testing.T) {
		id := uint(1)
		existing := &entities.Product{Id: id, Name: "Old Name", Description: "Old Desc", Price: 10.0}

		updateData := dtos.Product{Name: "New Name"}

		mockRepo.On("GetProductByID", ctx, id).Return(existing, nil).Once()

		// Verify: Name changes, but Description and Price remain from 'existing'
		mockRepo.On("UpdateProduct", ctx, id, mock.MatchedBy(func(p entities.Product) bool {
			return p.Name == "New Name" && p.Description == "Old Desc" && p.Price == 10.0
		})).Return(nil).Once()

		err := ps.UpdateProduct(ctx, id, updateData)
		assert.NoError(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	ps := service.NewProductService(mockRepo)
	ctx := context.Background()

	t.Run("success case", func(t *testing.T) {
		id := uint(5)
		mockRepo.On("DeleteProduct", ctx, id).Return(nil).Once()

		err := ps.DeleteProduct(ctx, id)
		assert.NoError(t, err)
	})

	t.Run("handle unexpected db error", func(t *testing.T) {
		mockRepo.On("DeleteProduct", ctx, uint(1)).Return(errors.New("db crash")).Once()

		err := ps.DeleteProduct(ctx, 1)
		assert.Error(t, err)
	})
}

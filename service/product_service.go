package service

import (
	"context"
	"errors"

	"gorm.io/gorm"

	customerrors "github.com/rhodamineb13/backend-test/errors"
	"github.com/rhodamineb13/backend-test/models/dtos"
	"github.com/rhodamineb13/backend-test/models/entities"
	"github.com/rhodamineb13/backend-test/repository"
)

type IProductService interface {
	ListProducts(ctx context.Context, name string, category uint, price float32, stock_quantity, limit, offset uint) (*dtos.ProductResponse, error)
	GetProductByID(ctx context.Context, id uint) (*dtos.Product, error)
	InsertNewProduct(ctx context.Context, data dtos.Product) error
	UpdateProduct(ctx context.Context, id uint, data dtos.Product) error
	DeleteProduct(ctx context.Context, id uint) error
}

type productService struct {
	productRepo repository.IProductRepository
}

func NewProductService(productRepo repository.IProductRepository) IProductService {
	return &productService{productRepo: productRepo}
}

func (ps *productService) ListProducts(ctx context.Context, name string, category uint, price float32, stock_quantity uint, limit, offset uint) (*dtos.ProductResponse, error) {
	productsRes, err := ps.productRepo.ListProducts(ctx, name, category, price, stock_quantity, limit, offset)
	if err != nil {
		return nil, customerrors.ErrUnexpected(err)
	}

	dashboardProduct := &dtos.ProductResponse{
		TotalStock:       int(productsRes.TotalStock),
		NumberOfProducts: int(productsRes.TotalProducts),
		AveragePrice:     productsRes.AveragePrice,
	}

	dashboardProduct.ProductDetails = make([]dtos.Product, len(productsRes.Products))
	for i, prod := range productsRes.Products {
		dashboardProduct.ProductDetails[i] = dtos.Product{
			Id:            prod.Id,
			Name:          prod.Name,
			Description:   prod.Description,
			StockQuantity: prod.StockQuantity,
			Price:         prod.Price,
			Category: dtos.Category{
				Id:   prod.CategoryId,
				Name: prod.Category.Name,
			},
		}
	}

	return dashboardProduct, nil
}

func (ps *productService) GetProductByID(ctx context.Context, id uint) (*dtos.Product, error) {
	product, err := ps.productRepo.GetProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrNotFound(err)
		} else {
			return nil, customerrors.ErrUnexpected(err)
		}
	}

	return &dtos.Product{
		Id:            product.Id,
		Name:          product.Name,
		Price:         product.Price,
		StockQuantity: product.StockQuantity,
		Category: dtos.Category{
			Id:   product.Category.Id,
			Name: product.Category.Name,
		},
	}, nil
}

func (ps *productService) InsertNewProduct(ctx context.Context, data dtos.Product) error {
	input := entities.Product{
		Name:          data.Name,
		Description:   data.Description,
		Price:         data.Price,
		StockQuantity: data.StockQuantity,
		CategoryId:    data.CategoryID,
	}

	if err := ps.productRepo.InsertNewProduct(ctx, input); err != nil {
		return customerrors.ErrUnexpected(err)
	}

	return nil
}

func (ps *productService) UpdateProduct(ctx context.Context, id uint, data dtos.Product) error {
	product, err := ps.productRepo.GetProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customerrors.ErrNotFound(err)
		}
		return customerrors.ErrUnexpected(err)
	}
	if data.Name != "" {
		product.Name = data.Name
	}
	if data.Description != "" {
		product.Description = data.Description
	}
	if data.Price != 0.0 {
		product.Price = data.Price
	}
	if err := ps.productRepo.UpdateProduct(ctx, id, *product); err != nil {
		return customerrors.ErrUnexpected(err)
	}
	return nil
}

func (ps *productService) DeleteProduct(ctx context.Context, id uint) error {
	if err := ps.productRepo.DeleteProduct(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customerrors.ErrNotFound(err)
		}
		return customerrors.ErrUnexpected(err)
	}
	return nil
}

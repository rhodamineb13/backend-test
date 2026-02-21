package service

import (
	"context"

	"github.com/rhodamineb13/backend-test/models/dtos"
	"github.com/rhodamineb13/backend-test/models/entities"
	"github.com/rhodamineb13/backend-test/repository"
)

type IProductService interface {
	ListProducts(ctx context.Context, name string, category uint, price float32, stock_quantity uint) ([]dtos.Product, error)
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

func (ps *productService) ListProducts(ctx context.Context, name string, category uint, price float32, stock_quantity uint) ([]dtos.Product, error) {
	productsEntity, err := ps.productRepo.ListProducts(ctx, name, category, price, stock_quantity)
	if err != nil {
		return nil, err
	}

	productsDto := make([]dtos.Product, len(productsEntity))
	for i := range productsEntity {
		productsDto[i] = dtos.Product{
			Id:            productsEntity[i].Id,
			Name:          productsEntity[i].Name,
			Price:         productsEntity[i].Price,
			StockQuantity: productsEntity[i].StockQuantity,
			CreatedAt:     &productsEntity[i].CreatedAt,
			Category: dtos.Category{
				Id:   productsEntity[i].Category.Id,
				Name: productsEntity[i].Category.Name,
			},
		}
	}

	return productsDto, nil
}

func (ps *productService) GetProductByID(ctx context.Context, id uint) (*dtos.Product, error) {
	product, err := ps.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
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
		return err
	}

	return nil
}

func (ps *productService) UpdateProduct(ctx context.Context, id uint, data dtos.Product) error {
	product, err := ps.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return err
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
		return err
	}
	return nil
}

func (ps *productService) DeleteProduct(ctx context.Context, id uint) error {
	if err := ps.productRepo.DeleteProduct(ctx, id); err != nil {
		return err
	}
	return nil
}

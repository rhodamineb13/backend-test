package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/rhodamineb13/backend-test/service"
)

type ProductHandler struct {
	productService service.IProductService
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (ph *ProductHandler) ListProducts(c *gin.Context) {}

func (ph *ProductHandler) GetProductByID(c *gin.Context) {}

func (ph *ProductHandler) InsertNewProduct(c *gin.Context) {}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {}

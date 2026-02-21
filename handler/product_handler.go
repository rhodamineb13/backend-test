package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rhodamineb13/backend-test/models/dtos"
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

func (ph *ProductHandler) ListProducts(c *gin.Context) {
	var query struct {
		Name          string  `json:"name"`
		Category      uint    `json:"category"`
		Price         float32 `json:"price"`
		StockQuantity uint    `json:"stock_quantity"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		return
	}

	products, err := ph.productService.ListProducts(c, query.Name, query.Category, query.Price, query.StockQuantity)
	if err != nil {
		return
	}

	c.JSON(200, products)
}

func (ph *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	product, err := ph.productService.GetProductByID(c, uint(id))
	if err != nil {
		return
	}

	c.JSON(200, product)
}

func (ph *ProductHandler) InsertNewProduct(c *gin.Context) {
	var product dtos.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return
	}

	if err := ph.productService.InsertNewProduct(c, product); err != nil {
		return
	}

	c.JSON(200, "")
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	var product dtos.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	if err := ph.productService.UpdateProduct(c, uint(id), product); err != nil {
		return
	}

	c.JSON(200, "")
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	if err := ph.productService.DeleteProduct(c, uint(id)); err != nil {
		return
	}
}

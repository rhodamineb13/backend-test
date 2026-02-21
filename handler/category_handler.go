package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rhodamineb13/backend-test/models/dtos"
	"github.com/rhodamineb13/backend-test/service"
)

type CategoryHandler struct {
	categoryService service.ICategoryService
}

func NewCategoryHandler(categoryService service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (ch *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := ch.categoryService.ListCategories(c)
	if err != nil {
		return
	}

	c.JSON(200, categories)
}

func (ch *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	category, err := ch.categoryService.GetCategoryByID(c, uint(id))
	if err != nil {
		return
	}

	c.JSON(200, category)
}

func (ch *CategoryHandler) InsertNewCategory(c *gin.Context) {
	var category dtos.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		return
	}

	if err := ch.categoryService.InsertNewCategory(c, category); err != nil {
		return
	}

	c.JSON(200, "")
}

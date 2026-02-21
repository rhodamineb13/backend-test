package handler

import (
	"github.com/gin-gonic/gin"

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

func (ch *CategoryHandler) ListCategories(c *gin.Context) {}

func (ch *CategoryHandler) GetCategoryByID(c *gin.Context) {}

func (ch *CategoryHandler) InsertNewCategory(c *gin.Context) {}

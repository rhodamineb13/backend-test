package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/rhodamineb13/backend-test/database"
	customerrors "github.com/rhodamineb13/backend-test/errors"
	"github.com/rhodamineb13/backend-test/handler"
	"github.com/rhodamineb13/backend-test/models/entities"
	"github.com/rhodamineb13/backend-test/repository"
	"github.com/rhodamineb13/backend-test/service"
	"github.com/rhodamineb13/backend-test/utils"
)

// @title Backend Test API
// @version 1.0
// @description This is a project that consists of Product API and Category API
func main() {
	utils.LoadEnv()
	database.ConnectRedis()
	database.ConnectSQL()


	if err:= database.DB.AutoMigrate(&entities.Category{}, &entities.Product{}); err != nil {
		panic(err)
	}

	categoryRepo := repository.NewCategoryRepository(database.DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(database.DB, database.RedisDB)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(customerrors.ErrorMiddleware)
	products := r.Group("/products")
	{
		products.GET("", productHandler.ListProducts)
		products.GET("/:id", productHandler.GetProductByID)
		products.POST("", productHandler.InsertNewProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}
	categories := r.Group("/categories")
	{
		categories.GET("", categoryHandler.ListCategories)
		categories.GET("/:id", categoryHandler.GetCategoryByID)
		categories.POST("", categoryHandler.InsertNewCategory)
	}

	r.Run()
}

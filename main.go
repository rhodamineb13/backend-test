package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	products := r.Group("/products")
	{
		products.GET("")
		products.GET("/:id")
		products.POST("")
		products.PUT("/:id")
		products.DELETE("/:id")
	}
	categories := r.Group("/categories")
	{
		categories.GET("")
		categories.GET("/:id")
		categories.POST("")
	}

	r.Run()
}

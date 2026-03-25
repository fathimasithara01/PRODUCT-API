package main

import (
	"os"

	"github.com/fathima-sithara/PRODUCT-API/internal/config"
	"github.com/fathima-sithara/PRODUCT-API/internal/handler"
	"github.com/fathima-sithara/PRODUCT-API/internal/middleware"
	"github.com/fathima-sithara/PRODUCT-API/internal/model"
	"github.com/fathima-sithara/PRODUCT-API/internal/repository"
	usecase "github.com/fathima-sithara/PRODUCT-API/internal/usercase"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDB()

	db.AutoMigrate(&model.Product{})

	productRepo := repository.NewProductRepo(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running "})
	})

	api := r.Group("/api/v1")

	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/products", productHandler.CreateProduct)
		auth.GET("/products", productHandler.GetAllProducts)
		auth.GET("/products/:id", productHandler.GetProductByID)
		auth.PUT("/products/:id", productHandler.UpdateProduct)
		auth.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

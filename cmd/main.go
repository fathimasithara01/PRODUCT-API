package main

import (
	"os"

	"github.com/fathima-sithara/PRODUCT-API/internal/auth"
	"github.com/fathima-sithara/PRODUCT-API/internal/config"
	"github.com/fathima-sithara/PRODUCT-API/internal/middleware"
	"github.com/fathima-sithara/PRODUCT-API/internal/product/handler"
	"github.com/fathima-sithara/PRODUCT-API/internal/product/model"
	"github.com/fathima-sithara/PRODUCT-API/internal/product/repository"
	"github.com/fathima-sithara/PRODUCT-API/internal/product/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDB()

	db.AutoMigrate(&model.Product{})

	authRepo := auth.NewRepository(db)
	authUsecase := auth.NewUsecase(authRepo)
	authHandler := auth.NewHandler(authUsecase)

	productRepo := repository.NewProductRepo(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	r := gin.Default()

	r.POST("/signup", authHandler.Signup)
	r.POST("/login", authHandler.Login)

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

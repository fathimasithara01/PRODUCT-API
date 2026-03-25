package main

import (
	"github.com/fathima-sithara/PRODUCT-API/internal/config"
	"github.com/fathima-sithara/PRODUCT-API/internal/handler"
	"github.com/fathima-sithara/PRODUCT-API/internal/middleware"
	"github.com/fathima-sithara/PRODUCT-API/internal/model"
	"github.com/fathima-sithara/PRODUCT-API/internal/repository"
	usecase "github.com/fathima-sithara/PRODUCT-API/internal/usercase"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	db.AutoMigrate(&model.Product{})

	repo := repository.NewProductRepo(db)
	usercase := usecase.NewProductUsecase(repo)
	handler := handler.NewProductHandler(usercase)

	r := gin.Default()

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/products", handler.CreateProduct)
		auth.GET("/products", handler.GetAllProducts)
		auth.GET("/products/:id", handler.GetProductByID)
		auth.PUT("/products/:id", handler.UpdateProduct)
		auth.DELETE("/products/:id", handler.DeleteProduct)
	}

	r.Run(":8080")

}

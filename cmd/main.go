package main

import (
	"github.com/fathima-sithara/PRODUCT-API/internal/config"
	"github.com/fathima-sithara/PRODUCT-API/internal/handler"
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

	r.POST("/products", handler.CreateProduct)
	r.GET("/products", handler.GetAllProducts)
	r.GET("/products/:id", handler.GetProductByID)
	r.PUT("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)

	r.Run(":8080")

}

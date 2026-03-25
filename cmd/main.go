package main

import (
	"os"

	"github.com/fathima-sithara/PRODUCT-API/internal/auth"
	"github.com/fathima-sithara/PRODUCT-API/internal/cart/cart_handler"
	"github.com/fathima-sithara/PRODUCT-API/internal/cart/cart_repository"
	"github.com/fathima-sithara/PRODUCT-API/internal/cart/cart_usecase"
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

	// auth := api.Group("/")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.GetAllProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	cartRepo := cart_repository.NewCartRepo(db)
	cartUsecase := cart_usecase.NewCartUsecase(cartRepo)
	cartHandler := cart_handler.NewCartHandler(cartUsecase)

	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/:user_id/:product_id", cartHandler.AddToCart)
		api.GET("/:user_id", cartHandler.GetCart)
		api.PUT("/:cart_id", cartHandler.UpdateQuantity)
		api.DELETE("/:cart_id", cartHandler.RemoveFromCart)
		api.DELETE("/clear/:user_id", cartHandler.ClearCart)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

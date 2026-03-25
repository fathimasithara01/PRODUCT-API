package handler

import (
	"net/http"
	"strconv"

	"github.com/fathima-sithara/PRODUCT-API/internal/model"
	usecase "github.com/fathima-sithara/PRODUCT-API/internal/usercase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{u}
}

// POST /products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GET /products
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, _ := h.usecase.GetAll()
	c.JSON(http.StatusOK, products)
}

// GET /products/:id
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// PUT /products/:id
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.usecase.Update(product)
	c.JSON(http.StatusOK, product)
}

// DELETE /products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	h.usecase.Delete(uint(id))
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

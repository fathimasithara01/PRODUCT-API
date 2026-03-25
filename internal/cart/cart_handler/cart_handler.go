package cart_handler

import (
	"net/http"
	"strconv"

	"github.com/fathima-sithara/PRODUCT-API/internal/cart/cart_usecase"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase cart_usecase.CartUsecase
}

func NewCartHandler(u cart_usecase.CartUsecase) *CartHandler {
	return &CartHandler{usecase: u}
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	productID, _ := strconv.Atoi(c.Param("product_id"))
	quantity, _ := strconv.Atoi(c.DefaultQuery("quantity", "1"))

	if err := h.usecase.AddToCart(uint(userID), uint(productID), quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "product added to cart"})
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	carts, total, err := h.usecase.GetCart(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": carts, "total": total})
}

func (h *CartHandler) UpdateQuantity(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))
	quantity, _ := strconv.Atoi(c.DefaultQuery("quantity", "1"))

	if err := h.usecase.UpdateQuantity(uint(cartID), quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "quantity updated"})
}

func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))

	if err := h.usecase.RemoveFromCart(uint(cartID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "item removed from cart"})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))

	if err := h.usecase.ClearCart(uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "cart cleared"})
}

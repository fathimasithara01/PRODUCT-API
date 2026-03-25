package cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(u Usecase) *Handler {
	return &Handler{u}
}

func (h *Handler) AddToCart(c *gin.Context) {

	userID := c.GetUint("user_id")

	productID, _ := strconv.Atoi(c.Param("product_id"))

	var input struct {
		Price float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.AddToCart(userID, uint(productID), input.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "added to cart"})
}

func (h *Handler) GetCart(c *gin.Context) {

	userID := c.GetUint("user_id")

	items, total, err := h.usecase.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": total,
	})
}

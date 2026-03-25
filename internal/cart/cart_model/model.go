package cart_model

import (
	"github.com/fathima-sithara/PRODUCT-API/internal/product/model"
	"gorm.io/gorm"
)

// type Cart struct {
// 	gorm.Model
// 	UserID    uint    `json:"user_id"`
// 	ProductID uint    `json:"product_id"`
// 	Quantity  int     `json:"quantity"`
// 	Price     float64 `json:"price"`
// }

type Cart struct {
	gorm.Model
	UserID    uint          `json:"user_id" gorm:"not null"`
	ProductID uint          `json:"product_id" gorm:"not null"`
	Product   model.Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int           `json:"quantity" gorm:"not null;default:1"`
}

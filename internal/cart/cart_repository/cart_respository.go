package cart_repository

import (
	"errors"

	"github.com/fathima-sithara/PRODUCT-API/internal/cart/cart_model"
	"gorm.io/gorm"
)

type CartRepository interface {
	Add(cart *cart_model.Cart) error
	GetUserCart(userID uint) ([]cart_model.Cart, error)
	GetByUserProduct(userID, productID uint) (*cart_model.Cart, error)
	Update(cart *cart_model.Cart) error
	Delete(cartID uint) error
	ClearCart(userID uint) error
}

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) CartRepository {
	return &cartRepo{db: db}
}

func (r *cartRepo) Add(cart *cart_model.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepo) GetUserCart(userID uint) ([]cart_model.Cart, error) {
	var carts []cart_model.Cart
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func (r *cartRepo) GetByUserProduct(userID, productID uint) (*cart_model.Cart, error) {
	var cart cart_model.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepo) Update(cart *cart_model.Cart) error {
	result := r.db.Model(&cart_model.Cart{}).Where("id = ?", cart.ID).Updates(cart)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *cartRepo) Delete(cartID uint) error {
	result := r.db.Delete(&cart_model.Cart{}, cartID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *cartRepo) ClearCart(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&cart_model.Cart{}).Error
}

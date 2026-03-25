package cart_usecase

import (
	"errors"

	"github.com/fathima-sithara/PRODUCT-API/internal/cart/cart_model"
	"github.com/fathima-sithara/PRODUCT-API/internal/cart/cart_repository"
	"gorm.io/gorm"
)

type CartUsecase interface {
	AddToCart(userID, productID uint, quantity int) error
	GetCart(userID uint) ([]cart_model.Cart, float64, error)
	UpdateQuantity(cartID uint, quantity int) error
	RemoveFromCart(cartID uint) error
	ClearCart(userID uint) error
}

type cartUsecase struct {
	repo cart_repository.CartRepository
}

func NewCartUsecase(r cart_repository.CartRepository) CartUsecase {
	return &cartUsecase{repo: r}
}

func (u *cartUsecase) AddToCart(userID, productID uint, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be at least 1")
	}

	cart, err := u.repo.GetByUserProduct(userID, productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newCart := &cart_model.Cart{
				UserID:    userID,
				ProductID: productID,
				Quantity:  quantity,
			}
			return u.repo.Add(newCart)
		}
		return err
	}

	cart.Quantity += quantity
	return u.repo.Update(cart)
}

func (u *cartUsecase) GetCart(userID uint) ([]cart_model.Cart, float64, error) {
	carts, err := u.repo.GetUserCart(userID)
	if err != nil {
		return nil, 0, err
	}

	var total float64
	for _, item := range carts {
		total += float64(item.Quantity) * item.Product.Price
	}
	return carts, total, nil
}

func (u *cartUsecase) UpdateQuantity(cartID uint, quantity int) error {
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	if quantity == 0 {
		return u.repo.Delete(cartID)
	}

	cart := &cart_model.Cart{
		Model:    gorm.Model{ID: cartID},
		Quantity: quantity,
	}
	return u.repo.Update(cart)
}

func (u *cartUsecase) RemoveFromCart(cartID uint) error {
	return u.repo.Delete(cartID)
}

func (u *cartUsecase) ClearCart(userID uint) error {
	return u.repo.ClearCart(userID)
}

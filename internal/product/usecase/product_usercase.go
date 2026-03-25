package usecase

import (
	"errors"

	"github.com/fathima-sithara/PRODUCT-API/internal/product/model"
	"github.com/fathima-sithara/PRODUCT-API/internal/product/repository"
	"gorm.io/gorm"
)

type ProductUsecase interface {
	Create(product *model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Delete(id uint) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) ProductUsecase {
	return &productUsecase{r}
}

func (u *productUsecase) Create(product *model.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("invalid product price")
	}
	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return u.repo.Create(product)
}

func (u *productUsecase) GetAll() ([]model.Product, error) {
	return u.repo.GetAll()
}

func (u *productUsecase) GetByID(id uint) (*model.Product, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}

	return u.repo.GetByID(id)
}

func (u *productUsecase) Update(product *model.Product) (*model.Product, error) {
	if product.ID == 0 {
		return nil, errors.New("invalid product id")
	}
	if product.Price <= 0 {
		return nil, errors.New("invalid product price")
	}
	if product.Stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}

	updatedProduct, err := u.repo.Update(product)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return updatedProduct, nil
}

func (u *productUsecase) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}

	return u.repo.Delete(id)
}

package usecase

import (
	"errors"

	"github.com/fathima-sithara/PRODUCT-API/internal/model"
	"github.com/fathima-sithara/PRODUCT-API/internal/repository"
)

type ProductUsecase interface {
	Create(product *model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
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

func (u *productUsecase) Update(product *model.Product) error {
	if product.ID == 0 {
		return errors.New("invalid product id")
	}
	if product.Price <= 0 {
		return errors.New("invalid product price")
	}

	return u.repo.Update(product)
}

func (u *productUsecase) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}

	return u.repo.Delete(id)
}

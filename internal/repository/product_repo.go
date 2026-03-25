package repository

import (
	"github.com/fathima-sithara/PRODUCT-API/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepo) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepo) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepo) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepo) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

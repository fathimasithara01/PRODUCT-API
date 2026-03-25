package product

import (
	"errors"

	"github.com/fathima-sithara/PRODUCT-API/internal/product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product Product) error
	GetAll() ([]Product, error)
	GetByID(id uint) (Product, error)
	Update(product Product) error
	Delete(id uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) Create(product Product) error {
	return r.db.Create(product).Error
}

func (r *productRepo) GetAll() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) GetByID(id uint) (*Product, error) {
	var product model.Product

	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) Update(product *model.Product) error {
	result := r.db.Model(&model.Product{}).
		Where("id = ?", product.ID).
		Updates(product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *productRepo) Delete(id uint) error {
	result := r.db.Delete(&model.Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

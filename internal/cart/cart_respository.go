package cart

import "gorm.io/gorm"

type Repository interface {
	Add(item *Cart) error
	GetByUser(userID uint) ([]Cart, error)
	GetItem(userID, productID uint) (*Cart, error)
	Update(item *Cart) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) Add(item *Cart) error {
	return r.db.Create(item).Error
}

func (r *repo) GetByUser(userID uint) ([]Cart, error) {
	var items []Cart
	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

func (r *repo) GetItem(userID, productID uint) (*Cart, error) {
	var item *Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&item).Error
	return item, err
}

func (r *repo) Update(item *Cart) error {
	return r.db.Save(item).Error
}

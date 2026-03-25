package auth

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, err
}

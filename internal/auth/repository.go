package auth

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repo) GetByEmail(email string) (*User, error) {
	var user User
	err := r.db.Find("email = ?", email).First(&user).Error
	return &user, err
}

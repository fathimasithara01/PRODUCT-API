package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	SignUp(user *User) error
	Login(email, password string) (string, error)
}
type authUsecase struct {
	repo UserRepository
}

func NewUsecase(repo UserRepository) AuthUsecase {
	return &authUsecase{repo}
}

func (u *authUsecase) SignUp(user *User) error {
	if user.Name == "" || user.Email == " " || user.Password == "" {
		return errors.New("all fields are required")
	}

	existing, _ := u.repo.GetByEmail(user.Email)
	if existing != nil && existing.ID != 0 {
		return errors.New("user already exist")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return u.repo.Create(user)
}

func (u *authUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

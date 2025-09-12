package service

import "go-shop-api/internal/model"

type IAuthService interface {
	Authenticate(email, password string) (string, string, *model.User, error)
	Register(email, password, role string) (*model.User, error)
	HashPassword(plain string) (string, error)
	ComparePassword(hash, plain string) error
	RefreshToken(refreshToken string) (string, string, error)
	Logout(userID uint) error
}

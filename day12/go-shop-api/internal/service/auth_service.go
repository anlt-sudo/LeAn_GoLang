package service

import (
	"errors"
	"time"

	"go-shop-api/internal/auth/jwt"
	"go-shop-api/internal/model"
	"go-shop-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
	AccessTokenTTL time.Duration
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo:       userRepo,
		AccessTokenTTL: time.Minute * 60,
	}
}

func (s *AuthService) HashPassword(plain string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bs), err
}

func (s *AuthService) ComparePassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

func (s *AuthService) Authenticate(email, password string) (string, *model.User, error) {
	u, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", nil, err
	}
	if err := s.ComparePassword(u.Password, password); err != nil {
		return "", nil, errors.New("invalid credentials")
	}
	token, err := jwt.CreateAccessToken(u.ID, u.Email, u.Role, s.AccessTokenTTL)
	if err != nil {
		return "", nil, err
	}
	return token, u, nil
}

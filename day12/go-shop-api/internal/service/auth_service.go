package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"go-shop-api/internal/auth/jwt"
	"go-shop-api/internal/model"
	"go-shop-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)



type AuthService struct {
	UserRepo        *repository.UserRepository
	RefreshRepo     *repository.RefreshTokenRepository
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func NewAuthService(userRepo *repository.UserRepository, refreshRepo *repository.RefreshTokenRepository) *AuthService {
	return &AuthService{
		UserRepo:        userRepo,
		RefreshRepo:     refreshRepo,
		AccessTokenTTL:  time.Minute * 60,
		RefreshTokenTTL: time.Hour * 24 * 7,
	}
}

func (s *AuthService) HashPassword(plain string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bs), err
}

func (s *AuthService) ComparePassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

func hashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *AuthService) Authenticate(email, password string) (string, string, *model.User, error) {
	u, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", "", nil, err
	}
	if err := s.ComparePassword(u.Password, password); err != nil {
		return "", "", nil, ErrInvalidCredentials
	}
	accessToken, err := jwt.CreateAccessToken(u.ID, u.Email, u.Role, s.AccessTokenTTL)
	if err != nil {
		return "", "", nil, err
	}
	refreshToken, err := jwt.GenerateRefreshToken(u.ID, s.RefreshTokenTTL)
	if err != nil {
		return "", "", nil, err
	}
	rt := &model.RefreshToken{
		UserID:    u.ID,
		TokenHash: hashToken(refreshToken),
		ExpiredAt: time.Now().Add(s.RefreshTokenTTL),
		Revoked:   false,
	}
	if err := s.RefreshRepo.Save(rt); err != nil {
		return "", "", nil, err
	}
	return accessToken, refreshToken, u, nil
}

func (s *AuthService) Register(email, password, role string) (*model.User, error) {
	_, err := s.UserRepo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already in use")
	}
	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return nil, err
	}
	u := &model.User{
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}
	if err := s.UserRepo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {
	tokenHash := hashToken(refreshToken)
	rt, err := s.RefreshRepo.FindValid(tokenHash)
	if err != nil {
		return "", "", errors.New("invalid or expired refresh token")
	}

	user, err := s.UserRepo.FindByID(rt.UserID)
	if err != nil {
		return "", "", err
	}

	var newAccess string
	newAccess, err = jwt.CreateAccessToken(user.ID, user.Email, user.Role, s.AccessTokenTTL)
	if err != nil {
		return "", "", err
	}

	if err = s.RefreshRepo.RevokeByUserID(user.ID); err != nil {
		return "", "", err
	}

	var newRefresh string
	newRefresh, err = jwt.GenerateRefreshToken(user.ID, s.RefreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	rtNew := &model.RefreshToken{
		UserID:    user.ID,
		TokenHash: hashToken(newRefresh),
		ExpiredAt: time.Now().Add(s.RefreshTokenTTL),
		Revoked:   false,
	}
	if err = s.RefreshRepo.Save(rtNew); err != nil {
		return "", "", err
	}

	return newAccess, newRefresh, nil
}



func (s *AuthService) Logout(userID uint) error {
	return s.RefreshRepo.RevokeByUserID(userID)
}

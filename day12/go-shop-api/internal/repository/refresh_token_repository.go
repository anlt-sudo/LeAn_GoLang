package repository

import (
	"go-shop-api/internal/model"

	"gorm.io/gorm"
)

type RefreshTokenRepositoryInterface interface {
	Save(token *model.RefreshToken) error
	FindValid(tokenHash string) (*model.RefreshToken, error)
	RevokeByUserID(userID uint) error
}

type RefreshTokenRepository struct {
	DB *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{DB: db}
}

func (r *RefreshTokenRepository) Save(token *model.RefreshToken) error {
	return r.DB.Create(token).Error
}

func (r *RefreshTokenRepository) FindByUserID(userID uint) ([]model.RefreshToken, error) {
	var tokens []model.RefreshToken
	err := r.DB.Where("user_id = ? AND revoked = false", userID).Find(&tokens).Error
	return tokens, err
}

func (r *RefreshTokenRepository) FindValid(tokenHash string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.DB.Where("token_hash = ? AND revoked = false AND expired_at > NOW()", tokenHash).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *RefreshTokenRepository) RevokeByUserID(userID uint) error {
	return r.DB.Model(&model.RefreshToken{}).Where("user_id = ?", userID).Update("revoked", true).Error
}

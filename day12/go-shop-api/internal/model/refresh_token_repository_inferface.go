package model

type RefreshTokenRepositoryInterface interface {
	Save(token *RefreshToken) error
	FindValid(tokenHash string) (*RefreshToken, error)
	RevokeByUserID(userID uint) error
}
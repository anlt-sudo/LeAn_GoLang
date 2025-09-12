package handler

import (
	"net/http"

	"go-shop-api/internal/dto"
	"go-shop-api/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.IAuthService
}

func NewAuthHandler(s service.IAuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access, refresh, u, err := h.Service.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	res := dto.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		User: dto.UserDTO{
			ID:    u.ID,
			Email: u.Email,
			Role:  u.Role,
		},
	}
	c.JSON(http.StatusOK, res)
}


func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.Service.Register(req.Email, req.Password, "user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.RegisterResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  u.Role,
	}
	c.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAccess,refreshToken , err := h.Service.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": newAccess, "refresh_token": refreshToken})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Logout(req.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

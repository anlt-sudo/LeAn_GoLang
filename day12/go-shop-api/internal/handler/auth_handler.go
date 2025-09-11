package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-shop-api/internal/dto"
	"go-shop-api/internal/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginResponse struct {
	AccessToken string      `json:"access_token"`
	User        dto.UserDTO `json:"user"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, u, err := h.Service.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	res := loginResponse{
		AccessToken: token,
		User: dto.UserDTO{
			ID:    u.ID,
			Email: u.Email,
			Role:  u.Role,
		},
	}

	c.JSON(http.StatusOK, res)
}

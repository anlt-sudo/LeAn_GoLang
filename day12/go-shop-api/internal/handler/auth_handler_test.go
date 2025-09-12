package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-shop-api/internal/dto"
	"go-shop-api/internal/handler"
	"go-shop-api/internal/model"

	"github.com/gin-gonic/gin"
)

type mockAuthService struct{}

func (m *mockAuthService) Authenticate(email, password string) (string, string, *model.User, error) {
	if email == "admin@example.com" && password == "password123" {
		return "access-token", "refresh-token",
			&model.User{ID: 1, Email: email, Role: "admin"}, nil
	}
	return "", "", nil, errors.New("invalid credentials")
}
func (m *mockAuthService) Register(email, password, role string) (*model.User, error) { return nil, nil }
func (m *mockAuthService) HashPassword(plain string) (string, error)                  { return "", nil }
func (m *mockAuthService) ComparePassword(hash, plain string) error                   { return nil }
func (m *mockAuthService) RefreshToken(refreshToken string) (string, string, error)   { return "", "", nil }
func (m *mockAuthService) Logout(userID uint) error                                   { return nil }

func TestLoginSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &mockAuthService{}
	h := handler.NewAuthHandler(mockService)

	router := gin.New()
	router.POST("/login", h.Login)

	body := []byte(`{"email":"admin@example.com","password":"password123"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp dto.LoginResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if resp.AccessToken == "" || resp.RefreshToken == "" {
		t.Errorf("expected access & refresh tokens, got %+v", resp)
	}
}

func TestLoginFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &mockAuthService{}
	h := handler.NewAuthHandler(mockService)

	router := gin.New()
	router.POST("/login", h.Login)

	body := []byte(`{"email":"wrong@example.com","password":"wrong"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}

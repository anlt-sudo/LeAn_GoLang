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
	"go-shop-api/internal/repository"
	"go-shop-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func BenchmarkLoginHandler(b *testing.B) {
	gin.SetMode(gin.TestMode)
	mockService := &mockAuthService{}
	h := handler.NewAuthHandler(mockService)

	router := gin.New()
	router.POST("/login", h.Login)

	body := []byte(`{"email":"admin@example.com","password":"password123"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

func setupMySQL() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/go-shop-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&model.User{}, &model.RefreshToken{})

	var count int64
	db.Model(&model.User{}).Where("email = ?", "admin@example.com").Count(&count)
	if count == 0 {
		hash, _ := service.NewAuthService(nil, nil).HashPassword("password123")
		db.Create(&model.User{Email: "admin@example.com", Password: hash, Role: "admin"})
	}

	return db
}

func BenchmarkLoginHandlerWithMySQL(b *testing.B) {
	gin.SetMode(gin.TestMode)
	db := setupMySQL()

	userRepo := repository.NewUserRepository(db)
	refreshRepo := repository.NewRefreshTokenRepository(db)
	authService := service.NewAuthService(userRepo, refreshRepo)
	authHandler := handler.NewAuthHandler(authService)

	router := gin.New()
	router.POST("/login", authHandler.Login)

	body := []byte(`{"email":"admin@example.com","password":"password123"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

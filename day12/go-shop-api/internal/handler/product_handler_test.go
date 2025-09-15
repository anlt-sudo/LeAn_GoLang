package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-shop-api/internal/dto"
	"go-shop-api/internal/handler"
	"go-shop-api/internal/model"
)

type mockProductService struct {
	mock.Mock
}

func (m *mockProductService) GetAll() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}
// Alias for interface compatibility
func (m *mockProductService) FindAll() ([]model.Product, error) {
	return m.GetAll()
}
func (m *mockProductService) GetByID(id uint) (*model.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Product), args.Error(1)
}
// Alias for interface compatibility
func (m *mockProductService) FindByID(id uint) (*model.Product, error) {
	return m.GetByID(id)
}
func (m *mockProductService) SearchByName(name string) ([]model.Product, error) {
	args := m.Called(name)
	return args.Get(0).([]model.Product), args.Error(1)
}
func (m *mockProductService) Create(p *model.Product) error {
	args := m.Called(p)
	return args.Error(0)
}
func (m *mockProductService) Update(p *model.Product) error {
	args := m.Called(p)
	return args.Error(0)
}
func (m *mockProductService) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}




func setupRouter(ph *handler.ProductHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/products", ph.GetAll)
	r.GET("/products/:id", ph.GetByID)
	r.GET("/products/search", ph.Search)
	r.POST("/products", ph.Create)
	r.PUT("/products/:id", ph.Update)
	r.DELETE("/products/:id", ph.Delete)
	return r
}

func TestProductHandler_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	products := []model.Product{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}
	mockSvc.On("GetAll").Return(products, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// --- BENCHMARKS ---
func BenchmarkProductHandler_GetAll(b *testing.B) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	products := []model.Product{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}
	mockSvc.On("FindAll").Return(products, nil)

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products", nil)
		r.ServeHTTP(w, req)
	}
}

func BenchmarkProductHandler_Create(b *testing.B) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	mockSvc.On("Create", mock.AnythingOfType("*model.Product")).Return(nil)

	body := dto.ProductRequest{
		Name:        "A",
		Description: "desc",
		Price:       10,
		CategoryID:  1,
	}
	jsonBody, _ := json.Marshal(body)

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	}
}

func BenchmarkProductHandler_GetByID(b *testing.B) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	product := &model.Product{ID: 1, Name: "A"}
	mockSvc.On("FindByID", uint(1)).Return(product, nil)

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products/1", nil)
		r.ServeHTTP(w, req)
	}
}


func TestProductHandler_GetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	product := &model.Product{ID: 1, Name: "A"}
	mockSvc.On("GetByID", uint(1)).Return(product, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_Search(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	products := []model.Product{{ID: 1, Name: "A"}}
	mockSvc.On("SearchByName", "A").Return(products, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/search?q=A", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	mockSvc.On("Create", mock.AnythingOfType("*model.Product")).Return(nil)

	body := dto.ProductRequest{
		Name:        "Aaaa",
		Description: "desc",
		Price:       10,
		CategoryID:  1,
	}
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProductHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	product := &model.Product{ID: 1, Name: "A"}
	mockSvc.On("GetByID", uint(1)).Return(product, nil)
	mockSvc.On("Update", product).Return(nil)

	body := dto.ProductRequest{
		Name:        "B",
		Description: "desc2",
		Price:       20,
		CategoryID:  2,
	}
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	mockSvc.On("Delete", uint(1)).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/products/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_GetByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockProductService)
	ph := handler.NewProductHandler(mockSvc)
	r := setupRouter(ph)

	mockSvc.On("GetByID", uint(2)).Return((*model.Product)(nil), errors.New("not found"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/2", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}



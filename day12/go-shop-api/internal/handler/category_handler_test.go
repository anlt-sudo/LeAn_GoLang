package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-shop-api/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockCategoryService struct {
    categories []model.Category
    category   *model.Category
    err        error
}

func (m *mockCategoryService) GetAll() ([]model.Category, error) {
    return m.categories, m.err
}
func (m *mockCategoryService) GetByID(id uint) (*model.Category, error) {
    return m.category, m.err
}
func (m *mockCategoryService) Create(c *model.Category) error {
    return m.err
}
func (m *mockCategoryService) Update(c *model.Category) error {
    return m.err
}
func (m *mockCategoryService) Delete(id uint) error {
    return m.err
}
func (m *mockCategoryService) SearchByName(keyword string) ([]model.Category, error) {
    return m.categories, m.err
}

func setupRouter(h *CategoryHandler) *gin.Engine {
    r := gin.Default()
    r.GET("/categories", h.GetAll)
    r.GET("/categories/:id", h.GetByID)
    r.POST("/categories", h.Create)
    r.PUT("/categories/:id", h.Update)
    r.DELETE("/categories/:id", h.Delete)
    return r
}

func TestGetAllCategoriesHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := &mockCategoryService{categories: []model.Category{{ID: 1, Name: "A"}}}
    h := NewCategoryHandler(svc)
    router := setupRouter(h)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/categories", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t,  http.StatusOK, w.Code)
    var got []model.Category
    json.Unmarshal(w.Body.Bytes(), &got)
    assert.Equal(t, 1, len(got))
    assert.Equal(t, "A", got[0].Name)
}

func TestGetCategoryByIDHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := &mockCategoryService{category: &model.Category{ID: 2, Name: "B"}}
    h := NewCategoryHandler(svc)
    router := setupRouter(h)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/categories/2", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateCategoryHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := &mockCategoryService{}
    h := NewCategoryHandler(svc)
    router := setupRouter(h)

    body := `{"name":"Test", "description":"Desc"}`
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/categories", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateCategoryHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := &mockCategoryService{category: &model.Category{ID: 4, Name: "D"}}
    h := NewCategoryHandler(svc)
    router := setupRouter(h)

    body := `{"name":"D"}`
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/categories/4", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCategoryHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := &mockCategoryService{}
    h := NewCategoryHandler(svc)
    router := setupRouter(h)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/categories/5", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllCategoriesHandler_Error(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := &mockCategoryService{err: errors.New("fail")}
    h := NewCategoryHandler(svc)
    router := setupRouter(h)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/categories", nil)
    router.ServeHTTP(w, req)

    assert.NotEqual(t, http.StatusOK, w.Code)
}
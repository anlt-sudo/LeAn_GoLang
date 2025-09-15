package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-shop-api/internal/model"
	"go-shop-api/internal/service"
)

type mockProductRepo struct {
	mock.Mock
}

func (m *mockProductRepo) GetAll() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

// Alias for interface compatibility
func (m *mockProductRepo) FindAll() ([]model.Product, error) {
	return m.GetAll()
}

func (m *mockProductRepo) GetByID(id uint) (*model.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Product), args.Error(1)
}

// Alias for interface compatibility
func (m *mockProductRepo) FindByID(id uint) (*model.Product, error) {
	return m.GetByID(id)
}

func (m *mockProductRepo) SearchByName(name string) ([]model.Product, error) {
	args := m.Called(name)
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *mockProductRepo) Create(p *model.Product) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *mockProductRepo) Update(p *model.Product) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *mockProductRepo) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestProductService_GetAll(t *testing.T) {
	repo := new(mockProductRepo)
	svc := service.NewProductService(repo)
	products := []model.Product{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}
	repo.On("GetAll").Return(products, nil)

	result, err := svc.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, products, result)
}

func TestProductService_GetByID(t *testing.T) {
	repo := new(mockProductRepo)
	svc := service.NewProductService(repo)
	product := &model.Product{ID: 1, Name: "A"}
	repo.On("GetByID", uint(1)).Return(product, nil)

	result, err := svc.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_SearchByName(t *testing.T) {
	repo := new(mockProductRepo)
	svc := service.NewProductService(repo)
	products := []model.Product{{ID: 1, Name: "A"}}
	repo.On("SearchByName", "A").Return(products, nil)

	result, err := svc.SearchByName("A")
	assert.NoError(t, err)
	assert.Equal(t, products, result)
}

func TestProductService_Create(t *testing.T) {
	repo := new(mockProductRepo)
	svc := service.NewProductService(repo)
	product := &model.Product{ID: 1, Name: "A"}
	repo.On("Create", product).Return(nil)

	err := svc.Create(product)
	assert.NoError(t, err)
}

func TestProductService_Update(t *testing.T) {
	repo := new(mockProductRepo)
	svc := service.NewProductService(repo)
	product := &model.Product{ID: 1, Name: "A"}
	repo.On("Update", product).Return(nil)

	err := svc.Update(product)
	assert.NoError(t, err)
}

func TestProductService_Delete(t *testing.T) {
	repo := new(mockProductRepo)
	svc := service.NewProductService(repo)
	repo.On("Delete", uint(1)).Return(nil)

	err := svc.Delete(1)
	assert.NoError(t, err)
}

func TestProductService_GetByID_NotFound(t *testing.T) {
	repo := new(mockProductRepo)
	svc := service.NewProductService(repo)
	repo.On("GetByID", uint(2)).Return((*model.Product)(nil), errors.New("not found"))

	result, err := svc.GetByID(2)
	assert.Error(t, err)
	assert.Nil(t, result)
}

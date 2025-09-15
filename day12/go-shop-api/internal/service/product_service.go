package service

import (
	"go-shop-api/internal/model"
)


type ProductService struct {
	Repo model.IProductRepository
}

// Implement IProductService compatibility
func (s *ProductService) FindAll() ([]model.Product, error) {
	return s.GetAll()
}

func (s *ProductService) FindByID(id uint) (*model.Product, error) {
	return s.GetByID(id)
}

func NewProductService(repo model.IProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.Repo.FindAll()
}

func (s *ProductService) GetByID(id uint) (*model.Product, error) {
	return s.Repo.FindByID(id)
}

func (s *ProductService) SearchByName(name string) ([]model.Product, error) {
	return s.Repo.SearchByName(name)
}

func (s *ProductService) Create(p *model.Product) error {
	return s.Repo.Create(p)
}

func (s *ProductService) Update(p *model.Product) error {
	return s.Repo.Update(p)
}

func (s *ProductService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

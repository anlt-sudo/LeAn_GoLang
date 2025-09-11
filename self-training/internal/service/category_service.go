package service

import (
	"errors"
	"self-training/internal/model"
	"self-training/internal/repository"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.Repo.FindAll()
}

func (s *CategoryService) GetByID(id uint) (*model.Category, error) {
	category, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (s *CategoryService) Create(c *model.Category) error {
	return s.Repo.Create(c)
}

func (s *CategoryService) Update(c *model.Category) error {
	return s.Repo.Update(c)
}

func (s *CategoryService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

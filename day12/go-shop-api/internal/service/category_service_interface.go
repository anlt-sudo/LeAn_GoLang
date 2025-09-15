package service

import "go-shop-api/internal/model"

type ICategoryService interface {
	GetAll() ([]model.Category, error)
	GetByID(id uint) (*model.Category, error)
	Create(c *model.Category) error
	Update(c *model.Category) error
	Delete(id uint) error
	SearchByName(keyword string) ([]model.Category, error)
}
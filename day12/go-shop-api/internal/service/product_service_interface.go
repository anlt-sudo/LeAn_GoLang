package service

import "go-shop-api/internal/model"

type IProductService interface {
	FindAll() ([]model.Product, error)
	FindByID(id uint) (*model.Product, error)
	Create(p *model.Product) error
	Update(p *model.Product) error
	Delete(id uint) error
	SearchByName(name string) ([]model.Product, error)
}

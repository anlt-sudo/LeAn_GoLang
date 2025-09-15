package model

type IProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(id uint) (*Product, error)
	Create(p *Product) error
	Update(p *Product) error
	Delete(id uint) error
	SearchByName(keyword string) ([]Product, error)
}
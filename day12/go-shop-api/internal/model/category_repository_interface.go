package model

type CategoryRepositoryInterface interface {
	FindAll() ([]Category, error)
	FindByID(id uint) (*Category, error)
	Create(c *Category) error
	Update(c *Category) error
	Delete(id uint) error
	SearchByName(keyword string) ([]Category, error)
}
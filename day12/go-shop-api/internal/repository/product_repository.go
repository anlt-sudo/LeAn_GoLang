package repository

import (
	"go-shop-api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByID(id uint) (*model.Product, error) {
	var p model.Product
	err := r.DB.First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) SearchByName(name string) ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

func (r *ProductRepository) Create(p *model.Product) error {
	return r.DB.Create(p).Error
}

func (r *ProductRepository) Update(p *model.Product) error {
	return r.DB.Save(p).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Product{}, id).Error
}

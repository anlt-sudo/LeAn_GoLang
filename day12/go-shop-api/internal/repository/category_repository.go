package repository

import (
	"go-shop-api/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.DB.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) Create(c *model.Category) error {
	return r.DB.Create(c).Error
}

func (r *CategoryRepository) Update(c *model.Category) error {
	return r.DB.Save(c).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Category{}, id).Error
}

func (r *CategoryRepository) SearchByName(keyword string) ([]model.Category, error) {
    var categories []model.Category
    err := r.DB.Where("name LIKE ?", "%"+keyword+"%").Find(&categories).Error
    return categories, err
}


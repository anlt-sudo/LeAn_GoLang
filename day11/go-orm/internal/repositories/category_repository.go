package repositories

import (
	"go-orm/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r *CategoryRepository) Create(name string) (string, error) {
	id := uuid.New().String()
	category := models.Category{
		ID:   id,
		Name: name,
	}
	if err := r.DB.Create(&category).Error; err != nil {
		return "", err
	}
	return id, nil
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := r.DB.Preload("Albums").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id string) (models.Category, error) {
	var cat models.Category
	if err := r.DB.Preload("Albums").First(&cat, "id = ?", id).Error; err != nil {
		return models.Category{}, err
	}
	return cat, nil
}

func (r *CategoryRepository) Update(id string, name string) error {
	return r.DB.Model(&models.Category{}).
		Where("id = ?", id).
		Update("name", name).Error
}

func (r *CategoryRepository) Delete(id string) error {
	return r.DB.Delete(&models.Category{}, "id = ?", id).Error
}

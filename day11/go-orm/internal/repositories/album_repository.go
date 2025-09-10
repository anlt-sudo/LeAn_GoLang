package repositories

import (
	"go-orm/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlbumRepository struct {
	DB *gorm.DB
}

func (r *AlbumRepository) Create(name string, categoryID string) (string, error) {
	id := uuid.New().String()
	album := models.Album{
		ID:         id,
		Name:       name,
		CategoryID: categoryID,
	}
	if err := r.DB.Create(&album).Error; err != nil {
		return "", err
	}
	return id, nil
}

func (r *AlbumRepository) GetAll() ([]models.Album, error) {
	var albums []models.Album
	if err := r.DB.Preload("Category").Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *AlbumRepository) GetByID(id string) (models.Album, error) {
	var alb models.Album
	if err := r.DB.Preload("Category").First(&alb, "id = ?", id).Error; err != nil {
		return models.Album{}, err
	}
	return alb, nil
}

func (r *AlbumRepository) Update(id string, name string, categoryID string) error {
	return r.DB.Model(&models.Album{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":        name,
			"category_id": categoryID,
		}).Error
}

func (r *AlbumRepository) Delete(id string) error {
	return r.DB.Delete(&models.Album{}, "id = ?", id).Error
}

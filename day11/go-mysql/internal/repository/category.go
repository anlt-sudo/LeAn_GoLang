package repository

import (
	"database/sql"
	"fmt"
	"go-mysql/internal/model"

	"github.com/google/uuid"
)

type CategoryRepository struct {
	DB *sql.DB
}

func (r *CategoryRepository) AddCategory(catName string) (string, error) {
	id := uuid.New().String()
	_, err := r.DB.Exec(
		"INSERT INTO category (id, name) VALUES (?, ?)",
		id, catName,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	rows, err := r.DB.Query("SELECT id, name FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) UpdateCategory(id string, newName string) error {
	_, err := r.DB.Exec(
		"UPDATE category SET name = ? WHERE id = ?",
		newName, id,
	)
	return err
}

func (r *CategoryRepository) CheckHasAlbums(categoryID string) (bool, error) {
	var count int
	err := r.DB.QueryRow(
		"SELECT COUNT(*) FROM album WHERE category_id = ?",
		categoryID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *CategoryRepository) DeleteCategory(id string) error {
	hasAlbums, err := r.CheckHasAlbums(id)
	if err != nil {
		return err
	}
	if hasAlbums {
		return fmt.Errorf("không thể xóa category %s vì vẫn còn album tham chiếu", id)
	}

	_, err = r.DB.Exec(
		"DELETE FROM category WHERE id = ?",
		id,
	)
	return err
}

package repository

import (
	"database/sql"
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

package dto

import (
	"go-shop-api/internal/model"
	"time"
)

type ProductRequest struct {
	Name        string  `json:"name" binding:"required,min=3"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToProductResponse(p model.Product) ProductResponse {
	var catID uint
	if p.CategoryID != nil {
		catID = *p.CategoryID
	}
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CategoryID:  catID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ToProductResponses(products []model.Product) []ProductResponse {
	res := make([]ProductResponse, 0, len(products))
	for _, p := range products {
		res = append(res, ToProductResponse(p))
	}
	return res
}

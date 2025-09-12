package model

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	SKU         string    `gorm:"unique" json:"sku"`
	Price       float64   `gorm:"not null" json:"price"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	CategoryID  *uint     `json:"category_id"`
	Category    Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
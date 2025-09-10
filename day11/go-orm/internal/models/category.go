package models

type Category struct {
	ID     string  `gorm:"type:char(36);primaryKey"`
	Name   string  `gorm:"type:varchar(100);not null"`
	Albums []Album `gorm:"foreignKey:CategoryID"`
}
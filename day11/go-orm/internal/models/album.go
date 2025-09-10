package models

type Album struct {
	ID         string      `gorm:"type:char(36);primaryKey"`
	Name       string   `gorm:"type:varchar(128);not null"`
	CategoryID string   `gorm:"type:char(36);index"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
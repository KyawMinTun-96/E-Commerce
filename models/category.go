package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `json:"category_name"`
	Products     []Product `gorm:"foreignKey:CategoryID"` // One-to-many relationship
}

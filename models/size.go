package models

import (
	"gorm.io/gorm"
)

type Size struct {
	gorm.Model
	Sizevalue string    `json:"size_value"`
	Products  []Product `gorm:"many2many:product_size"`
}

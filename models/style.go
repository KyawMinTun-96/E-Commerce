package models

import (
	"gorm.io/gorm"
)

type Style struct {
	gorm.Model
	StyleName string    `json:"style_name"`
	Products  []Product `gorm:"many2many:product_style"`
}

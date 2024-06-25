package models

import (
	"gorm.io/gorm"
)

type Price struct {
	gorm.Model
	Price     float64 `json:"price"`
	ProductID uint    `json:"product_id"`
}

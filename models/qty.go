package models

import (
	"gorm.io/gorm"
)

type Qty struct {
	gorm.Model
	Quantity  int  `json:"qty"`
	ProductID uint `gorm:"product_id"`
}

package models

import (
	"gorm.io/gorm"
)

// type Product struct {
// 	gorm.Model
// 	Gender string  `json:"gender"`
// 	Style  string  `json:"style"`
// 	Size   string  `json:"size"`
// 	Price  float64 `json:"price"`
// }

type Product struct {
	gorm.Model
	CategoryID uint `json:"category_id"`

	Item string `json:"item"`

	QtyID   Qty   `gorm:"foreignKey:ProductID"`
	PriceID Price `gorm:"foreignKey:ProductID"`

	StyleID  []Style  `gorm:"many2many:StyleID"`
	GenderID []Gender `gorm:"many2many:GenderID"`
	SizeID   []Size   `gorm:"many2many:SizeID"`
}

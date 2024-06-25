package models

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	GenderType string    `json:"gender_type"`
	Products   []Product `gorm:"many2many:product_genders"`
}

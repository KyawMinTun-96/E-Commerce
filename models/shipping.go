package models

import (
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	Orders []Order `json:"orders"`
	Status string  `json:"status"`
	UserID uint    `json:"user_id"`
}

package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderQty int  `json:"order_quantity"`
	OrderID  uint `json:"order_id"`
}

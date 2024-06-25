package models

import (
	"gorm.io/gorm"
)

// type Order struct {
// 	gorm.Model
// 	ShippingAddress string      `json:"shipping_address"`
// 	Status          string      `json:"status"`
// 	Items           []OrderItem `json:"items"`
// }

type Order struct {
	gorm.Model
	ProductID   []Product `gorm:"many2many:ProductID"`
	UserID      uint      `json:"user_id"`
	ShippingID  uint      `json:"shipping_id"`
	OrderItemID OrderItem `gorm:"foreignKey:OrderID"`
}

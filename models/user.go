package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string     `json:"cuser_name"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	Orders    []Order    `json:"orders"`
	Shippings []Shipping `json:"shippings"`
}

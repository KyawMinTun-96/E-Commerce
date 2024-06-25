package models

import (
	"gorm.io/gorm"
)

type AdminUser struct {
	gorm.Model
	AdminName string `json:"admin_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	StaffID   string `json:"staff_id"`
	Role      string `json:"role"`
}

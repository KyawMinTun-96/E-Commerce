package migrations

import (
	"ecommerce-app/models"
	"log"
)

func RunMigrations() {
	err := models.DB.AutoMigrate(
		&models.User{},
		&models.Gender{},
		&models.Style{},
		&models.Category{},
		&models.Product{},
		&models.Size{},
		&models.Qty{},
		&models.Shipping{},
		&models.Order{},
		&models.Price{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

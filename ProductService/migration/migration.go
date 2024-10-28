package migration

import (
	"log"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"gorm.io/gorm"
)

func Must(db *gorm.DB) {
	log.Println("Starting database migration...")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get raw DB instance: %v", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Database connection lost: %v", err)
	}

	if err := db.Exec("CREATE SCHEMA IF NOT EXISTS order_service").Error; err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	if err := db.AutoMigrate(
		&entity.Product{},
		&entity.Category{},
		&entity.Image{},
		&entity.Inventory{},
		&entity.ProductCategory{},
	); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migration completed successfully.")
}
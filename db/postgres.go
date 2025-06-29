package db

import (
	"log"
	"taas/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=taasuser password=taasuser dbname=taas port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := db.AutoMigrate(&models.Tag{}); err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	return db
}

package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"taas/models"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=taskuser password=taskpass dbname=taskdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := db.AutoMigrate(&models.Tag{}); err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	return db
}

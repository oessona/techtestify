package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"techtestify/internal/models"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=7206 dbname=techtestify port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}
	
	DB.Exec("DELETE FROM results")
	if err := DB.AutoMigrate(&models.User{}, &models.Test{}, &models.Question{}, &models.Result{}); err != nil {
		log.Fatal("❌ Failed to migrate User:", err)
	}

	log.Println("✅ Connected to database")
}

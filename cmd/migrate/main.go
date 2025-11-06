package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kaidora-labs/mitter-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("DB_URI is not set")
	}

	DB, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Failed ")
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Database Migration Failed")
	}

}

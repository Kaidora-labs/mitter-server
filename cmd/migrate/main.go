package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kaidora-labs/mitter-server/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = repositories.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v\n", err)
	}

	err = repositories.Migrate()
	if err != nil {
		log.Fatalf("Database migration failed: %v\n", err)
	}

	log.Println("Database migration completed successfully")
}

package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kaidora-labs/mitter-server/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}

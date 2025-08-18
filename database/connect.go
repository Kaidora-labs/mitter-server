package database

import (
	"fmt"
	"os"

	"github.com/kaidora-labs/mitter-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	uri := os.Getenv("DB_URL")
	if uri == "" {
		return fmt.Errorf("DB URI is empty")
	}

	DB, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	err := Connect()
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}

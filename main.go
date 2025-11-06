package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kaidora-labs/mitter-server/database"
	"github.com/kaidora-labs/mitter-server/handlers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s\n", err)
	}
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	userGroup := router.Group("/users")

	userGroup.GET("/", handlers.GetUsersHandler)
	userGroup.POST("/", handlers.PostUserHandler)
	userGroup.GET("/:id", handlers.GetUserHandler)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	router.Run()
}

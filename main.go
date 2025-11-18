package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kaidora-labs/mitter-server/handlers"
	"github.com/kaidora-labs/mitter-server/repositories"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	err = repositories.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.POST("/auth/initiate", handlers.InitiateHandler)
	router.POST("/auth/validate", handlers.ValidateHandler)
	router.POST("/auth/reset", handlers.ResetHandler)

	router.GET("/users", handlers.GetUsersHandler)
	router.GET("/users/:id", handlers.GetUserHandler)
	router.POST("/users", handlers.PostUserHandler)
	router.DELETE("/users/:id", handlers.DeleteUserHandler)

	router.GET("/businesses/:id", handlers.GetBusinessHandler)
	router.PATCH("/businesses/:id", handlers.UpdateBusinessHandler)
	router.DELETE("/businesses/:id", handlers.DeleteBusinessHandler)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Service is up",
			"data":    "OK",
		})
	})

	router.Run()
}

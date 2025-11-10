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

	authGroup := router.Group("/auth")
	authGroup.POST("/initiate", handlers.InitiateHandler)
	authGroup.POST("/validate", handlers.ValidateHandler)
	authGroup.POST("/reset", handlers.ResetHandler)

	userGroup := router.Group("/users")
	userGroup.GET("/", handlers.GetUsersHandler)
	userGroup.POST("/", handlers.PostUserHandler)
	userGroup.GET("/:id", handlers.GetUserHandler)
	userGroup.DELETE("/:id", handlers.DeleteUserHandler)

	businessGroup := router.Group("/businesses")
	businessGroup.GET("/:id", handlers.GetBusinessHandler)
	businessGroup.PATCH("/:id", handlers.UpdateBusinessHandler)
	businessGroup.DELETE("/:id", handlers.DeleteBusinessHandler)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run()
}

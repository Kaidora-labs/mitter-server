package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kaidora-labs/mitter-server/database"
	"github.com/kaidora-labs/mitter-server/handlers"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URI")), &gorm.Config{})
	if err != nil {
		log.Fatal("Database")

	}

	if database.CACHE = redis.NewClient(&redis.Options{
		Addr: os.Getenv("CACHE_ADDR"),
	}); database.CACHE == nil {
		log.Fatal("Redis Connection Failed")
	}

}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	authGroup := router.Group("/auth")
	authGroup.POST("/login", handlers.LoginHandler)
	authGroup.POST("/register", handlers.RegisterHandler)
	authGroup.POST("/validate", handlers.ValidateHandler)
	authGroup.POST("/forgot-password", handlers.ForgotPasswordHandler)

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

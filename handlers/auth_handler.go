package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidora-labs/mitter-server/repositories"
	"github.com/kaidora-labs/mitter-server/services"
)

type Credentials struct {
	EmailAddress string `json:"emailAddress" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

func InitiateHandler(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	repo := repositories.New()
	user, err := repo.FindUserByEmailAddress(credentials.EmailAddress)

	if err != nil || !services.ValidateHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	jwtToken, err := services.GenerateJWT(credentials.EmailAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func ValidateHandler(c *gin.Context) {
	panic("Unimplemented")
}

func ResetHandler(c *gin.Context) {
	panic("Unimplemented")
}

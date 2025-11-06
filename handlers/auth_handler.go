package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kaidora-labs/mitter-server/database"
	"github.com/kaidora-labs/mitter-server/models"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func LoginHandler(c *gin.Context) {
	var credentials Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userRepo := models.NewUserRepository(database.DB)
	user, err := userRepo.FindByUsername(credentials.Username)

	if err != nil || !CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	jwtToken, err := GenerateJWT(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func RegisterHandler(c *gin.Context) {
	var credentials Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	panic("Unimplemented")
}

func ValidateHandler(c *gin.Context) {
	panic("Unimplemented")
}

func ForgotPasswordHandler(c *gin.Context) {
	panic("Unimplemented")
}

func GenerateJWT(username string) (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return "", fmt.Errorf("JWT_SECRET environment variable not set")
	}

	expirationTime := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable not set")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

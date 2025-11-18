package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kaidora-labs/mitter-server/repositories"
	"github.com/kaidora-labs/mitter-server/services"
)

func GetUserHandler(c *gin.Context) {
	repo := repositories.New()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})

		return
	}

	user, err := repo.FindUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User retrieved successfully",
		"data":    user,
	})
}

func GetUsersHandler(c *gin.Context) {
	repo := repositories.New()
	users, err := repo.FindAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve users",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

func PostUserHandler(c *gin.Context) {
	repo := repositories.New()

	var params repositories.CreateUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing or invalid fields",
			"error":   err.Error(),
		})

		return
	}

	encryptedPassword, err := services.HashPassword(params.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not hash password",
			"error":   err.Error(),
		})

		return
	}
	params.Password = encryptedPassword

	user, err := repo.SaveUser(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
			"error":   err.Error(),
		})

		return
	}

	mailer, err := services.NewMailService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Mail service down",
			"error":   err.Error(),
		})

		return
	}

	otp, err := services.GenerateOTP(6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate OTP",
			"error":   err.Error(),
		})

		return
	}

	if err := mailer.SendOTP(user.EmailAddress, otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not send OTP",
			"error":   err.Error(),
		})

		return
	}

	if err := repo.StoreOTP(c, user.ID, otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cache OTP",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User verification in progress",
		"data":    user,
	})
}

func DeleteUserHandler(c *gin.Context) {
	repo := repositories.New()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})

		return
	}

	if err := repo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete user",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"data":    nil,
	})
}

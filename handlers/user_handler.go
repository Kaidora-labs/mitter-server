package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidora-labs/mitter-server/database"
	"github.com/kaidora-labs/mitter-server/models"
)

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	userRepo := models.NewUserRepository(database.DB)
	user, err := userRepo.Find(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User retrieved successfully",
		"data":    user,
	})
}

func PostUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})

		return
	}

	userRepo := models.NewUserRepository(database.DB)
	createdUser, err := userRepo.Save(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User created successfully",
		"data":    createdUser,
	})
}

func GetUsersHandler(c *gin.Context) {
	userRepo := models.NewUserRepository(database.DB)
	users, err := userRepo.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	userRepo := models.NewUserRepository(database.DB)
	err := userRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deleted successfully",
		"data":    nil,
	})
}

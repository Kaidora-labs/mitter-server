package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	// id := c.Param("id")

	// userRepo := repositories.NewUserRepo()
	// user, err := userRepo.Find(id)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"message": "User not found",
	// 		"error":   err.Error(),
	// 	})

	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "User retrieved successfully",
	// 	"data":    user,
	// })
}

func PostUserHandler(c *gin.Context) {
	// var user repositories.User

	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Missing or invalid fields",
	// 		"error":   err.Error(),
	// 	})

	// 	return
	// }

	// encryptedPassword, err := services.HashPassword(user.Password)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Could not hash password",
	// 		"error":   err.Error(),
	// 	})

	// 	return
	// }

	// user.Password = encryptedPassword

	// userRepo := repositories.NewUserRepo()
	// createdUser, err := userRepo.Save(&user)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Could not create user",
	// 		"error":   err.Error(),
	// 	})

	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{
	// 	"message": "User created successfully",
	// 	"data":    createdUser,
	// })
}

func GetUsersHandler(c *gin.Context) {
	// userRepo := repositories.NewUserRepo()
	// users, err := userRepo.FindAll()

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Could not retrieve users",
	// 		"error":   err.Error(),
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Users retrieved successfully",
	// 	"data":    users,
	// })
}

func DeleteUserHandler(c *gin.Context) {
	// id := c.Param("id")

	// userRepo := repositories.NewUserRepo()
	// err := userRepo.Delete(id)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Could not delete user",
	// 		"error":   err.Error(),
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "User deleted successfully",
	// 	"data":    nil,
	// })
}

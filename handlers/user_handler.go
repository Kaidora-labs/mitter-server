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
		c.JSON(http.StatusBadRequest, Result{
			Message: "Invalid user ID",
			Error:   err.Error(),
		})

		return
	}

	user, err := repo.FindUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, Result{
			Message: "User not found",
			Error:   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Result{
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func GetUsersHandler(c *gin.Context) {
	repo := repositories.New()
	users, err := repo.FindAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not retrieve users",
			Error:   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Result{
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func PostUserHandler(c *gin.Context) {
	repo := repositories.New()

	var params repositories.CreateUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, Result{
			Message: "Missing or invalid fields",
			Error:   err.Error(),
		})

		return
	}

	encryptedPassword, err := services.HashPassword(params.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not hash password",
			Error:   err.Error(),
		})

		return
	}
	params.Password = encryptedPassword

	user, err := repo.SaveUser(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not create user",
			Error:   err.Error(),
		})

		return
	}

	mailer, err := services.NewMailService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Mail service down",
			Error:   err.Error(),
		})

		return
	}

	otp, err := services.GenerateOTP(6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not generate OTP",
			Error:   err.Error(),
		})

		return
	}

	if err := mailer.SendOTP(user.EmailAddress, otp); err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not send OTP",
			Error:   err.Error(),
		})

		return
	}

	if err := repo.StoreOTP(c, user.ID, otp); err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not cache OTP",
			Error:   err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, Result{
		Message: "User verification in progress",
		Data:    user,
	})
}

func DeleteUserHandler(c *gin.Context) {
	repo := repositories.New()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{
			Message: "Invalid user ID",
			Error:   err.Error(),
		})

		return
	}

	if err := repo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not delete user",
			Error:   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Result{
		Message: "User deleted successfully",
		Data:    nil,
	})
}

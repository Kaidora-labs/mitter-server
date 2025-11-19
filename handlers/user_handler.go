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

	value, exists := c.Get(services.ClaimsKey{})
	if !exists {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not retrieve user claims",
			Error:   "User claims not found in context",
		})
		return
	}

	claims, ok := value.(*services.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not retrieve user claims",
			Error:   "Invalid claims format",
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
		return
	}

	if claims.ID != id {
		c.JSON(http.StatusForbidden, Result{
			Message: "Access denied",
			Error:   "You do not have permission to access this user",
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

func DeleteUserHandler(c *gin.Context) {
	repo := repositories.New()

	value, exists := c.Get(services.ClaimsKey{})
	if !exists {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not retrieve user claims",
			Error:   "User claims not found in context",
		})
		return
	}

	claims, ok := value.(*services.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not retrieve user claims",
			Error:   "Invalid claims format",
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
		return
	}

	if claims.ID != id {
		c.JSON(http.StatusForbidden, Result{
			Message: "Access denied",
			Error:   "You do not have permission to delete this user",
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

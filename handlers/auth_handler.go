package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidora-labs/mitter-server/repositories"
	"github.com/kaidora-labs/mitter-server/services"
)

func RegisterHandler(c *gin.Context) {
	repo := repositories.New()
	var params Params

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

	user := repositories.User{
		FirstName:     params.FirstName,
		LastName:      params.LastName,
		PhoneNumber:   params.PhoneNumber,
		EmailAddress:  params.EmailAddress,
		WalletAddress: params.WalletAddress,
		Password:      params.Password,
		Role:          params.Role,
		Businesses:    params.Businesses,
	}

	if err := repo.SaveUser(&user); err != nil {
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

func InitiateHandler(c *gin.Context) {
	repo := repositories.New()
	var credentials Credentials

	if err := c.Bind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, Result{
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	user, err := repo.FindUserByEmailAddress(credentials.EmailAddress)
	if err != nil || !services.ValidateHash(credentials.Password, user.Password) {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "invalid email address or password",
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
		Message: "Initialized Authentication",
		Data:    user,
	})
}

func ValidateHandler(c *gin.Context) {
	repo := repositories.New()
	var credentials Credentials

	if err := c.Bind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, Result{
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	user, err := repo.FindUserByEmailAddress(credentials.EmailAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "invalid email address",
			Error:   err.Error(),
		})
		return
	}

	otp, err := repo.RetrieveOTP(c, user.ID)
	if err != nil || otp != credentials.OTP {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "invalid OTP",
			Error:   err.Error(),
		})
		return
	}

	if otp != credentials.OTP {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "invalid OTP",
			Error:   "OTP does not match",
		})
		return
	}

	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not generate token",
			Error:   err.Error(),
		})
		return
	}

	err = repo.DeleteOTP(c, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Message: "Could not delete OTP",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Result{
		Data: Result{
			Message: "Authentication successful",
			Data: gin.H{
				"user":  user,
				"token": token,
			},
		},
	})

}

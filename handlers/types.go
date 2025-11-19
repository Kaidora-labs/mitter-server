package handlers

import "github.com/kaidora-labs/mitter-server/repositories"

type Params struct {
	FirstName     string                  `json:"firstName" binding:"required,min=2,max=50"`
	LastName      string                  `json:"lastName" binding:"required,min=2,max=50"`
	PhoneNumber   string                  `json:"phoneNumber" binding:"required,e164"`
	EmailAddress  string                  `json:"emailAddress" binding:"required,email"`
	WalletAddress string                  `json:"walletAddress" binding:"required"`
	Password      string                  `json:"password" binding:"required,min=8"`
	Role          repositories.Role       `json:"role" binding:"required,oneof=PROPRIETOR INDIVIDUAL"`
	Businesses    []repositories.Business `json:"businesses"`
}

type Credentials struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password,omitempty" binding:"omitempty,min=8"`
	OTP          string `json:"otp,omitempty" binding:"omitempty,len=6"`
}

type Result struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

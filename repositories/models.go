package repositories

import (
	"database/sql"

	"gorm.io/gorm"
)

// TODO: Implement Enum in DB
type Role = string

const (
	Proprietor Role = "PROPRIETOR"
	Individual Role = "INDIVIDUAL"
)

// TODO: Switch to UUID for ID
// TODO: Add Struct Binding Validation
type User struct {
	gorm.Model
	FirstName     string         `json:"firstName"`
	LastName      string         `json:"lastName"`
	Password      string         `json:"password"`
	PhoneNumber   string         `json:"phoneNumber" gorm:"uniqueIndex"`
	EmailAddress  string         `json:"emailAddress" gorm:"uniqueIndex"`
	WalletAddress sql.NullString `json:"walletAddress" gorm:"uniqueIndex"`
	Businesses    []Business     `json:"business_id"`
	Role          Role           `json:"role"`
	OTP           string         `gorm:"-:all"`
}

type BusinessType = string

const (
	SoleProprietor BusinessType = "SOLE_PROPRIETOR"
	PrivateLimited BusinessType = "PRIVATE_LIMITED"
	PublicLimited  BusinessType = "PUBLIC_LIMITED"
)

// TODO: Switch to UUID for ID
// TODO: Add Struct Binding Validation
type Business struct {
	gorm.Model
	Name      string       `json:"name"`
	Address   string       `json:"address"`
	CacNumber uint64       `json:"cacNumber" gorm:"uniqueIndex"`
	UserID    uint         `json:"user_id" gorm:"foreignKey"`
	Type      BusinessType `json:"type"`
}

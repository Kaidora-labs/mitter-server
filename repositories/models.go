package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO: Implement Enum in DB
type Role = string

const (
	Admin      Role = "ADMIN"
	Proprietor Role = "PROPRIETOR"
	Individual Role = "INDIVIDUAL"
)

type User struct {
	gorm.Model
	ID            uuid.UUID  `json:"id" gorm:"type:uuid"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	PhoneNumber   string     `json:"phoneNumber" gorm:"uniqueIndex"`
	EmailAddress  string     `json:"emailAddress" gorm:"uniqueIndex"`
	WalletAddress string     `json:"walletAddress" gorm:"uniqueIndex"`
	Businesses    []Business `json:"businesses"`
	Role          Role       `json:"role"`
	Password      string     `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	if u.Role == Admin {
		return fmt.Errorf("invalid role")
	}

	return nil
}

// TODO: Implement Enum in DB
type BusinessType = string

const (
	SoleProprietor BusinessType = "SOLE_PROPRIETOR"
	PrivateLimited BusinessType = "PRIVATE_LIMITED"
	PublicLimited  BusinessType = "PUBLIC_LIMITED"
)

type Business struct {
	gorm.Model
	ID        uuid.UUID    `json:"id" gorm:"type:uuid"`
	UserID    uuid.UUID    `json:"userId" gorm:"foreignKey"`
	Name      string       `json:"name"`
	Address   string       `json:"address"`
	CacNumber uint64       `json:"cacNumber" gorm:"uniqueIndex"`
	Type      BusinessType `json:"type"`
}

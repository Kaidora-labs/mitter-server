package models

import (
	"gorm.io/gorm"
)

// User model for the application
type User struct {
	gorm.Model
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PhoneNumber  string `json:"phoneNumber"`
	EmailAddress string `json:"emailAddress"`
}

// DTOs (Data Transfer Objects) for User
type CreateUserDTO struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
}

type GetUserDTO struct {
	Id string `json:"id"`
}

type UpdateUserDTO struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

type DeleteUserDTO struct {
	Id string `json:"id"`
}

// User Repository interface for database operations
type UserRepository interface {
	Find(id string) (*User, error)
	Save(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
	FindAll() (*[]User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user *User) (*User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindAll() (*[]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepository) Find(id string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *User) (*User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(id string) error {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

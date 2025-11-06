package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	UserName     string `json:"userName" binding:"required" gorm:"uniqueIndex"`
	Password     string `json:"password" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required" gorm:"uniqueIndex"`
	EmailAddress string `json:"emailAddress" binding:"required" gorm:"uniqueIndex"`
}

type UserRepository interface {
	Save(user *User) (*User, error)
	Delete(id string) error
	FindAll() (*[]User, error)
	Find(id string) (*User, error)
	FindByUsername(username string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
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

func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
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

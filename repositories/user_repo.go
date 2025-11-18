package repositories

import "github.com/google/uuid"

type CreateUserParams struct {
	FirstName     string     `json:"firstName" binding:"required,min=2,max=50"`
	LastName      string     `json:"lastName" binding:"required,min=2,max=50"`
	PhoneNumber   string     `json:"phoneNumber" binding:"required,e164"`
	EmailAddress  string     `json:"emailAddress" binding:"required,email"`
	WalletAddress string     `json:"walletAddress" binding:"required"`
	Password      string     `json:"password" binding:"required,min=8"`
	Role          Role       `json:"role" binding:"required,oneof=PROPRIETOR INDIVIDUAL"`
	Businesses    []Business `json:"businesses"`
}

func (r *Repo) SaveUser(params *CreateUserParams) (*User, error) {
	user := User{
		FirstName:     params.FirstName,
		LastName:      params.LastName,
		PhoneNumber:   params.PhoneNumber,
		EmailAddress:  params.EmailAddress,
		WalletAddress: params.WalletAddress,
		Password:      params.Password,
		Role:          params.Role,
		Businesses:    params.Businesses,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repo) DeleteUser(id uuid.UUID) error {
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

func (r *Repo) FindUser(id uuid.UUID) (*User, error) {
	var user User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repo) FindAllUsers() (*[]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *Repo) FindUserByUserName(userName string) (*User, error) {
	var user User

	err := r.db.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repo) FindUserByEmailAddress(emailAddress string) (*User, error) {
	var user User

	err := r.db.Where("email_address = ?", emailAddress).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

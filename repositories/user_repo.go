package repositories

func (r *Repo) SaveUser(user *User) (*User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) DeleteUser(id string) error {
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

func (r *Repo) FindUser(id string) (*User, error) {
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

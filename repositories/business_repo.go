package repositories

// type BusinessRepo struct {
// 	db    *gorm.DB
// 	cache *redis.Client
// }

// func NewBusinessRepo() *BusinessRepo {
// 	return &BusinessRepo{database, cache}
// }

// func (r *BusinessRepo) Save(user *Business) (*Business, error) {
// 	err := r.db.Create(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (r *BusinessRepo) Delete(id string) error {
// 	var user Business

// 	err := r.db.Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		return err
// 	}

// 	err = r.db.Delete(&user).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *BusinessRepo) Find(id string) (*Business, error) {
// 	var user Business

// 	err := r.db.Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (r *BusinessRepo) FindAll() (*[]Business, error) {
// 	var users []Business

// 	err := r.db.Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &users, nil
// }

// func (r *BusinessRepo) FindByBusinessName(userName string) (*Business, error) {
// 	var user Business

// 	err := r.db.Where("user_name = ?", userName).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (r *BusinessRepo) FindByEmailAddress(emailAddress string) (*Business, error) {
// 	var user Business

// 	err := r.db.Where("email_address = ?", emailAddress).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (r *BusinessRepo) SetOTP(ctx context.Context, emailAddress string, otp string) error {
// 	err := r.cache.Set(ctx, emailAddress, otp, 30*time.Minute).Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *BusinessRepo) GetOTP(ctx context.Context, emailAddress string) (string, error) {
// 	otp, err := r.cache.Get(ctx, emailAddress).Result()
// 	if err != nil {
// 		return "", err
// 	}

// 	return otp, nil
// }

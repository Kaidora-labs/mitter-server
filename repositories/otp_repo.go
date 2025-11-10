package repositories

import (
	"context"
	"time"
)

func (r *Repo) StoreOTP(ctx context.Context, emailAddress string, otp string) error {
	err := r.cache.Set(ctx, emailAddress, otp, 30*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) RetrieveOTP(ctx context.Context, emailAddress string) (string, error) {
	otp, err := r.cache.Get(ctx, emailAddress).Result()
	if err != nil {
		return "", err
	}

	return otp, nil
}

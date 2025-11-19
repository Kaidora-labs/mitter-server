package repositories

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cache *redis.Client
var database *gorm.DB

type Repo struct {
	db    *gorm.DB
	cache *redis.Client
}

func New() *Repo {
	return &Repo{database, cache}
}

func Connect() error {
	var err error

	databaseUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		return fmt.Errorf("DB_URL is not set")
	}

	database, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return err
	}

	cacheUrl, ok := os.LookupEnv("CACHE_URL")
	if !ok {
		return fmt.Errorf("CACHE_URL is not set")
	}

	opt, err := redis.ParseURL(cacheUrl)
	if err != nil {
		return fmt.Errorf("invalid CACHE_URL: %w", err)
	}

	// BUG: This is an issue from redis/go-redis. Update once fixed upstream
	// https://github.com/redis/go-redis/issues/3536
	opt.MaintNotificationsConfig = &maintnotifications.Config{
		Mode: maintnotifications.ModeDisabled,
	}

	cache = redis.NewClient(opt)

	return nil
}

func Migrate() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = database.AutoMigrate(User{}, Business{})

	if err != nil {
		return fmt.Errorf("database migration failed: %w", err)
	}

	return nil
}

func (r *Repo) StoreOTP(ctx context.Context, id uuid.UUID, otp string) error {
	err := r.cache.Set(ctx, id.String(), otp, 30*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) RetrieveOTP(ctx context.Context, id uuid.UUID) (string, error) {
	otp, err := r.cache.Get(ctx, id.String()).Result()
	if err != nil {
		return "", err
	}

	return otp, nil
}

func (r *Repo) DeleteOTP(ctx context.Context, id uuid.UUID) error {
	err := r.cache.Del(ctx, id.String()).Err()
	if err != nil {
		return err
	}

	return nil
}

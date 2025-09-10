package cache

import (
	"context"
	"errors"
	"sync"
	"time"
	"go-api-starter/core/logger"
	"go-api-starter/core/utils"

	"go-api-starter/core/constants"

	"github.com/redis/go-redis/v9"
)

var (
	instance *Cache
	once     sync.Once
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr, password string, db int) *Cache {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		// Test the connection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := client.Ping(ctx).Result()
		if err != nil {
			logger.Error("Failed to connect to Redis: " + err.Error())
		}

		instance = &Cache{
			client: client,
		}
	})
	return instance
}

// GetClient returns the Redis client
func (c *Cache) GetClient() *redis.Client {
	return c.client
}

// Set sets a key-value pair with expiration
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func (c *Cache) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.client.Get(ctx, key)
}

// Del removes a key
func (c *Cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Incr increments a key's value
func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

// Expire sets expiration for a key
func (c *Cache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

// Close closes the Redis connection
func (c *Cache) Close() error {
	return c.client.Close()
}

func (c *Cache) GetOTP(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) SetOTP(ctx context.Context, key string, value string) error {
	return c.client.Set(ctx, key, value, constants.DefaultOTPExpiration).Err()
}

func (c *Cache) IsLoginBlocked(ctx context.Context, key string) (bool, error) {
	count, err := c.client.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return false, err
	}
	if count >= constants.MaxLoginAttempts {

		return true, nil
	}
	return false, nil
}

func (c *Cache) IncrementLoginAttempt(ctx context.Context, key string) error {
	// Tăng +1, nếu chưa tồn tại thì set TTL luôn
	val, err := c.client.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	if val == 1 {
		// Set TTL nếu là lần đầu tiên
		err = c.client.Expire(ctx, key, constants.BlockDuration).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cache) ClearLoginAttempt(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *Cache) AddToTokenBlacklist(ctx context.Context, token string) error {
	tokenData, err := utils.ValidateAndParseToken(token)
	if err != nil {
		return err
	}

	jti := tokenData.RegisteredClaims.ID
	if jti == "" {
		return errors.New("token missing jti")
	}

	if tokenData.RegisteredClaims.ExpiresAt == nil {
		return errors.New("token missing exp")
	}

	ttl := time.Until(tokenData.RegisteredClaims.ExpiresAt.Time)
	if ttl <= 0 {
		return errors.New("token already expired")
	}

	// Lưu jti vào Redis với TTL, value chỉ cần "1"
	return c.client.SetEx(ctx, constants.TokenBlacklistKey+jti, 1, ttl).Err()
}

func (c *Cache) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	parsedToken, err := utils.ValidateAndParseToken(token)
	if err != nil {
		return false, err
	}

	jti := parsedToken.RegisteredClaims.ID
	if jti == "" {
		return false, errors.New("token missing jti")
	}

	exists, err := c.client.Exists(ctx, constants.TokenBlacklistKey+jti).Result()
	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

// ClearTokenBlacklist xóa toàn bộ blacklist
func (c *Cache) ClearTokenBlacklist(ctx context.Context) error {
	keys, err := c.client.Keys(ctx, "blacklist:*").Result()
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}
	return c.client.Del(ctx, keys...).Err()
}

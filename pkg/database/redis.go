package database

import (
	"context"
	"github.com/samber/do/v2"
	"go-api-starter/core/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RedisClient *redis.Client
}

// NewRedis creates a new Cache instance
func NewRedis(injector do.Injector) (*Redis, error) {
	appConfig := do.MustInvoke[*config.Config](injector)
	cfg := appConfig.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Redis{RedisClient: rdb}, nil
}

type IRedis interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string, value any) error
	Del(ctx context.Context, key string) error
}

func (r *Redis) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return r.RedisClient.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Get(ctx context.Context, key string, value any) error {
	return r.RedisClient.Get(ctx, key).Scan(value)
}

func (r *Redis) Del(ctx context.Context, key string) error {
	return r.RedisClient.Del(ctx, key).Err()
}

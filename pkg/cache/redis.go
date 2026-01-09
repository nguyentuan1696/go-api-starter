package cache

import (
	"context"
	"fmt"
	"go-api-starter/pkg/config"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(injector do.Injector) (*Redis, error) {

	appConfig := do.MustInvoke[*config.Config](injector)
	cfg := appConfig.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {

		return nil, fmt.Errorf("redis ping error: %w", err)
	}

	return &Redis{
		client: client,
	}, nil
}

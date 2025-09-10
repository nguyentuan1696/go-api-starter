package database

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Icache interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string, value any) error
	Del(ctx context.Context, key string) error
}

type Cache struct {
	Client *redis.Client
}

// Singleton variables
var (
	cacheInstance *Cache
	cacheOnce     sync.Once
)

// GetCacheInstance returns the singleton instance of Cache
func GetCacheInstance(client *redis.Client) *Cache {
	cacheOnce.Do(func() {
		cacheInstance = &Cache{Client: client}
	})
	return cacheInstance
}

// NewCache creates a new Cache instance (deprecated, use GetCacheInstance instead)
func NewCache(client *redis.Client) *Cache {
	return &Cache{Client: client}
}

func (c *Cache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return c.Client.Set(ctx, key, value, expiration).Err()
}

func (c *Cache) Get(ctx context.Context, key string, value any) error {
	return c.Client.Get(ctx, key).Scan(value)
}

func (c *Cache) Del(ctx context.Context, key string) error {
	return c.Client.Del(ctx, key).Err()
}

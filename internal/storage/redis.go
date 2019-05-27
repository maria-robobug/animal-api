package storage

import (
	"github.com/go-redis/redis"
)

const (
	defaultRedisPort = "6379"
	host             = "localhost"
)

// RedisCache a cache to store frequent data
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache returns an instance of RedisCache
func NewRedisCache() (*RedisCache, error) {
	rc := &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     host + ":" + defaultRedisPort,
			Password: "",
			DB:       0,
		}),
	}

	_, err := rc.client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return rc, nil
}

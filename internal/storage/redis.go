package storage

import (
	"fmt"

	"github.com/go-redis/redis"
)

// RedisCache a cache to store frequent data
type RedisCache struct {
	Host   string
	client *redis.Client
}

// NewCache returns an instance of RedisCache
func NewCache(host string) (*RedisCache, error) {
	rc := &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     host,
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

// Get retrieves key value from cache store
func (c *RedisCache) Get(key string) (interface{}, error) {
	val, err := c.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("key", val)

	return val, nil
}

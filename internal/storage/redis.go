package storage

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type redisPool struct {
	Host   string
	client *redis.Pool
}

type RedisConn struct {
	conn redis.Conn
}

// NewRedisPool returns an instance of RedisCache
func NewRedisPool(host string) (*RedisConn, error) {
	rp := &redisPool{
		client: &redis.Pool{
			// Maximum number of idle connections in the pool.
			MaxIdle: 80,
			// max number of connections
			MaxActive: 12000,
			// Dial is an application supplied function for creating and
			// configuring a connection.
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", host)
				if err != nil {
					return nil, err
				}
				return c, err
			},
		},
	}

	return &RedisConn{
		conn: rp.client.Get(),
	}, nil
}

// PingRedis tests connectivity for redis
func (c *RedisConn) PingRedis() error {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(c.conn.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}

// CloseRedisConn closes redis connection pool
func (c *RedisConn) CloseRedisConn() {
	c.conn.Close()
}

// Get closes redis connection pool
func (c *RedisConn) Get(key string) (interface{}, error) {
	return key, nil
}

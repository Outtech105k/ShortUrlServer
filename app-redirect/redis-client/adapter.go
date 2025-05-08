package redisclient

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter() (*RedisAdapter, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, fmt.Errorf("connectRedis: %w", err)
	}

	return &RedisAdapter{client: client}, nil
}

func (r *RedisAdapter) Close() error {
	return r.client.Close()
}

func (r *RedisAdapter) IsExists(key string) (bool, error) {
	exists, err := r.client.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

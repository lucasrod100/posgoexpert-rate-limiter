package strategy

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorageStrategy struct {
	client *redis.Client
}

func NewRedisStorageStrategy(addrRedis string) *RedisStorageStrategy {
	client := redis.NewClient(&redis.Options{
		Addr: addrRedis,
		DB:   0,
	})
	return &RedisStorageStrategy{client: client}
}

func (r *RedisStorageStrategy) Get(ctx context.Context, key string) (int, error) {
	val, err := r.client.Get(r.client.Context(), key).Int()
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

func (r *RedisStorageStrategy) Incr(ctx context.Context, key string) error {
	_, err := r.client.Incr(r.client.Context(), key).Result()
	return err
}

func (r *RedisStorageStrategy) Expire(ctx context.Context, key string, duration time.Duration) error {
	_, err := r.client.Expire(r.client.Context(), key, duration).Result()
	return err
}

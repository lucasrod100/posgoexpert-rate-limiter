package strategy

import (
	"context"
	"time"
)

type StorageStrategy interface {
	Get(ctx context.Context, key string) (int, error)
	Incr(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, duration time.Duration) error
}

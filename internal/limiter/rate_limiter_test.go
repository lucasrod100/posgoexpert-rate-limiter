package limiter

import (
	"testing"
	"time"

	"github.com/lucasrod100/posgoexpert/RateLimiter/internal/limiter/strategy"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestRateLimiter_UnderLimit(t *testing.T) {
	storage := strategy.NewRedisStorageStrategy("localhost:6322")
	ctx := context.Background()

	rl := NewRateLimiter(storage, 2, 2, time.Duration(10)*time.Second)

	allowed, err := rl.Allow(ctx, "ip:127.0.0.1", true)
	assert.NoError(t, err)
	assert.True(t, allowed)
}

func TestRateLimiter_OverLimit(t *testing.T) {
	storage := strategy.NewRedisStorageStrategy("localhost:6322")
	ctx := context.Background()

	rl := NewRateLimiter(storage, 2, 2, time.Duration(10)*time.Second)

	rl.Allow(ctx, "ip:127.0.0.1", true)
	rl.Allow(ctx, "ip:127.0.0.1", true)
	allowed, err := rl.Allow(ctx, "ip:127.0.0.1", true)
	assert.NoError(t, err)
	assert.False(t, allowed)
}

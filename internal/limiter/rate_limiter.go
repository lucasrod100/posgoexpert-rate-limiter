package limiter

import (
	"context"
	"log"
	"time"

	"github.com/lucasrod100/posgoexpert/RateLimiter/internal/limiter/strategy"
)

type RateLimiterInterface interface {
	Allow(ctx context.Context, key string, isIp bool) (bool, error)
}

type RateLimiter struct {
	Storage          strategy.StorageStrategy
	MaxRequestsIP    int
	MaxRequestsToken int
	BlockTime        time.Duration
}

func NewRateLimiter(storage strategy.StorageStrategy, maxRequestsIP, maxRequestsToken int, blockTime time.Duration) RateLimiterInterface {
	return &RateLimiter{
		Storage:          storage,
		MaxRequestsIP:    maxRequestsIP,
		MaxRequestsToken: maxRequestsToken,
		BlockTime:        blockTime,
	}
}

func (r *RateLimiter) Allow(ctx context.Context, key string, isIp bool) (bool, error) {
	count, _ := r.Storage.Get(ctx, key)
	log.Println(count)

	maxRequests := r.MaxRequestsToken
	if isIp {
		maxRequests = r.MaxRequestsIP
	}
	log.Println(maxRequests)

	if count >= maxRequests {
		return false, nil
	}
	log.Println(maxRequests)
	err := r.Storage.Incr(ctx, key)
	log.Println(err)
	if err != nil {
		return false, err
	}
	err = r.Storage.Expire(ctx, key, r.BlockTime)
	if err != nil {
		return false, err
	}
	return true, nil
}

package main

import (
	"context"
	"time"

	"github.com/lucasrod100/posgoexpert/RateLimiter/configs"
	"github.com/lucasrod100/posgoexpert/RateLimiter/internal/infra/server"
	"github.com/lucasrod100/posgoexpert/RateLimiter/internal/limiter"
	"github.com/lucasrod100/posgoexpert/RateLimiter/internal/limiter/strategy"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	storage := strategy.NewRedisStorageStrategy(configs.RedisADDR)

	rateLimiter := limiter.NewRateLimiter(storage, configs.MaxRequestsIP, configs.MaxRequestsToken, time.Duration(configs.BlockTime)*time.Second)

	webserver := server.NewWebServer("8080")
	webserver.Run(ctx, rateLimiter)
}

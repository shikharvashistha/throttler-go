package utils

import (
	"context"
	"os"

	"github.com/go-redis/redis/v9"
)

var (
	rdb *redis.Client = nil
	ctx               = context.Background()
)

func RedisConnect() (*redis.Client, context.Context) {
	if rdb != nil {
		return rdb, ctx
	}
	logger := NewLogger("cache")

	address := os.Getenv("REDIS_ADDRESS")

	password := os.Getenv("REDIS_PASSWORD")

	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.WithError(ADB, err).Info("Failed to initialize the redis client")
		return nil, nil
	}
	return rdb, ctx
}

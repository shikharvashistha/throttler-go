package utils

import (
	"context"

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

	// address := os.Getenv("REDIS_ADDRESS")

	// password := os.Getenv("REDIS_PASSWORD")

	rdb = redis.NewClient(&redis.Options{
		Addr:     "containers-us-west-187.railway.app:7556",
		Password: "uEP58w0FMoKs3cPutYyc",
		DB:       0,
		Username: "default",
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.WithError(ADB, err).Info("Failed to initialize the redis client")
		return nil, nil
	}
	return rdb, ctx
}

package utils

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var (
	rdb *redis.Client = nil
	ctx               = context.Background()
)

func GetRedis() (*redis.Client, context.Context) {
	if rdb == nil {
		panic("Redis not connected")
	}

	return rdb, ctx
}

func RedisConnect(addr, password, username string, db int) (*redis.Client, context.Context) {
	if rdb != nil {
		return rdb, ctx
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		Username: username,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, nil
	}
	return rdb, ctx
}

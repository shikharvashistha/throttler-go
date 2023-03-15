package keyvalue

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

type kvs struct{}

func (k *kvs) Set(key, value string, time time.Duration) error {
	rdb, ctx := utils.RedisConnect()
	err := rdb.Set(ctx, key, value, time).Err()
	return err
}

func (k *kvs) Get(key string) (string, error) {
	rdb, ctx := utils.RedisConnect()
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (k *kvs) Remove(key string) error {
	rdb, ctx := utils.RedisConnect()
	err := rdb.Del(ctx, key).Err()
	return err
}

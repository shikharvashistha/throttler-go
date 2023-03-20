package keyvalue

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

type KVS struct{}

func (k *KVS) Push(key string, values []string, time time.Duration) error {
	rdb, ctx := utils.GetRedis()
	err := rdb.RPush(ctx, key, values).Err()

	if err != nil {
		return err
	}

	err = rdb.Expire(ctx, key, time).Err()
	return err
}

func (k *KVS) Get(key string) ([]string, error) {
	rdb, ctx := utils.GetRedis()
	val, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err == redis.Nil {
		return make([]string, 0), nil
	}
	return val, err
}

func (k *KVS) Overwrite(key string, values []string, time time.Duration) error {
	rdb, ctx := utils.GetRedis()
	pipe := rdb.TxPipeline()

	pipe.Del(ctx, key)
	pipe.RPush(ctx, key, values)
	pipe.Expire(ctx, key, time)

	_, err := pipe.Exec(ctx)
	return err

}

func (k *KVS) Remove(key string) error {
	rdb, ctx := utils.GetRedis()
	err := rdb.Del(ctx, key).Err()
	return err
}

package keyvalue

import "time"

type KV interface {
	Get(key string) (string, error)
	Set(key, value string, time time.Duration) error
	Remove(key string) error
}

func NewKVStore() KV {
	return &kvs{}
}

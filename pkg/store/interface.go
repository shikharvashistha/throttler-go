package keyvalue

import "time"

type KV interface {
	Get(key string) ([]string, error)
	Push(key string, values []string, time time.Duration) error
	Remove(key string) error
	Overwrite(key string, values []string, time time.Duration) error
}

func NewKVStore() KV {
	return &KVS{}
}

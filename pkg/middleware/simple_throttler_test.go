package middleware

import (
	"net/http"
	"testing"
	"time"
)

func get_cache_key_func(r *http.Request) (string, error) { return "test_key", nil }

type test_kvs struct {
	kv map[string][]string
}

func (k *test_kvs) Push(key string, values []string, time time.Duration) error {
	k.kv[key] = values
	return nil
}

func (k *test_kvs) Get(key string) ([]string, error) {
	value, ok := k.kv[key]
	if ok {
		return value, nil
	}
	return make([]string, 0), nil
}

func (k *test_kvs) Overwrite(key string, values []string, time time.Duration) error {
	k.kv[key] = values
	return nil
}

func (k *test_kvs) Remove(key string) error {
	delete(k.kv, key)
	return nil
}

func TestAllowRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "test.com", nil)
	req.RemoteAddr = "test123"

	simple_throttle := SimpleRateThrottle{}

	simple_throttle.Allow_request(nil, req)
	kvs := &test_kvs{}
	kvs.kv = make(map[string][]string)

	simple_throttle.Init(kvs, get_cache_key_func)

	simple_throttle.Allow_request(nil, req)
}

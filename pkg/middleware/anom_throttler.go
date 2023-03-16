package middleware

import (
	"net/http"

	keyvalue "github.com/shikharvashistha/throttler-go/pkg/store"
)

type AnomThrottle struct {
	Simple_throttle SimpleRateThrottle
}

func (t *AnomThrottle) get_cache_key(r *http.Request) (string, error) {
	return t.Simple_throttle.base_throttle.GetIndent(r), nil
}

func (t *AnomThrottle) Init() {
	t.Simple_throttle.Init(keyvalue.NewKVStore(), t.get_cache_key)
}

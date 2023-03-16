package middleware

import (
	"net/http"

	keyvalue "github.com/shikharvashistha/throttler-go/pkg/store"
)

type CustomThrottle struct {
	Simple_throttle SimpleRateThrottle
}

func (t *CustomThrottle) Init(get_cache_key func(r *http.Request) (string, error)) {
	t.Simple_throttle.Init(keyvalue.NewKVStore(), get_cache_key)
}

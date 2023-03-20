package middleware

import (
	"net/http"
	"time"

	keyvalue "github.com/shikharvashistha/throttler-go/pkg/store"
)

func GetCustomThrottle(
	reqAllowed int,
	inDur time.Duration,
	scope string,
	kvs keyvalue.KV,
	getCacheKey func(r *http.Request, scope string) (string, error)) throttle {

	throttle := &SimpleRateThrottle{}
	throttle.Init(reqAllowed, inDur, kvs, scope, getCacheKey)
	return throttle
}

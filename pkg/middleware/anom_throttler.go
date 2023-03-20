package middleware

import (
	"net/http"
	"time"

	keyvalue "github.com/shikharvashistha/throttler-go/pkg/store"
	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

func GetAnonymousThrottle(
	reqAllowed int,
	inDur time.Duration,
	scope string,
	numProxies int,
	kvs keyvalue.KV) throttle {
	throttle := &SimpleRateThrottle{}

	throttle.Init(reqAllowed, inDur, kvs, scope, func(r *http.Request, scope string) (string, error) {
		return utils.GetIndent(r, numProxies) + scope, nil
	})

	return throttle
}

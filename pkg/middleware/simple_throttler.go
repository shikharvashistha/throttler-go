package middleware

import (
	"net/http"
	"strconv"
	"time"

	keyvalue "github.com/shikharvashistha/throttler-go/pkg/store"

	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

type SimpleRateThrottle struct {
	cache       keyvalue.KV
	key, scope  string
	history     []string
	now         time.Time
	getCacheKey func(r *http.Request, scope string) (string, error)
	reqAllowed  int
	inDur       time.Duration
}

func (t *SimpleRateThrottle) Init(
	reqAllowed int,
	inDur time.Duration,
	cache keyvalue.KV,
	scope string,
	getCacheKey func(r *http.Request, scope string) (string, error)) {

	t.cache = cache
	t.scope = scope
	t.getCacheKey = getCacheKey
	t.reqAllowed = reqAllowed
	t.inDur = inDur
}

func (t *SimpleRateThrottle) throttleSuccess(r *http.Request) (bool, error) {

	t.history = append(t.history, strconv.FormatInt(t.now.Unix(), 10))

	err := t.cache.Overwrite(t.key, t.history, t.inDur)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *SimpleRateThrottle) throttleFailure(r *http.Request) (bool, error) {
	return false, nil
}

func (t *SimpleRateThrottle) AllowRequest(r *http.Request) (bool, error) {

	t.now = time.Now()
	var err error

	t.key, err = t.getCacheKey(r, t.scope)
	if err != nil {
		return false, err
	}

	t.key = t.key + t.scope

	t.history, err = t.cache.Get(t.key)
	if err != nil {
		return false, err
	}

	// Drop any requests from the history which have now passed the
	// throttle duration
	for len(t.history) > 0 {

		lastTime, err := utils.ParseTimestamp(t.history[0])
		if err != nil {
			return false, err
		}

		if lastTime.Before(t.now.Add(-t.inDur)) {
			t.history = t.history[1:len(t.history)]
		} else {
			break
		}
	}

	if len(t.history) >= t.reqAllowed {
		return t.throttleFailure(r)
	}
	return t.throttleSuccess(r)
}

func (t *SimpleRateThrottle) Wait() (float64, error) {
	var remainingDuration, availableRequests int

	if t.history != nil {
		lastTime, err := utils.ParseTimestamp(t.history[0])

		if err != nil {
			return 0, err
		}

		remainingDuration = int(t.inDur.Seconds()) - int(t.now.Sub(*lastTime).Seconds())
		availableRequests = t.reqAllowed - len(t.history)
	} else {
		remainingDuration = int(t.inDur.Seconds())
		availableRequests = t.reqAllowed
	}

	return float64(remainingDuration) / float64(availableRequests), nil
}

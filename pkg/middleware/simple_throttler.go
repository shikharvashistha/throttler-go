package middleware

import (
	"net/http"
	"strconv"
	"time"

	keyvalue "github.com/shikharvashistha/throttler-go/pkg/store"

	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

var (
	req = 10
	dur = 60
)

type SimpleRateThrottle struct {
	base_throttle BaseThrottle
	cache         keyvalue.KV
	key           string
	history       []string
	now           time.Time
	get_cache_key func(r *http.Request) (string, error)
}

func (t *SimpleRateThrottle) Init(cache keyvalue.KV, get_cache_key_func func(r *http.Request) (string, error)) {
	utils.RedisConnect()
	t.cache = cache
	t.get_cache_key = get_cache_key_func
}

func (t *SimpleRateThrottle) throttle_success(w http.ResponseWriter, r *http.Request) (bool, error) {
	t.history = append(t.history, strconv.FormatInt(t.now.Unix(), 10))

	err := t.cache.Overwrite(t.key, t.history, time.Second*time.Duration(dur))

	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *SimpleRateThrottle) throttle_failure(w http.ResponseWriter, r *http.Request) (bool, error) {
	return false, nil
}

func (t *SimpleRateThrottle) Allow_request(w http.ResponseWriter, r *http.Request) (bool, error) {

	var err error

	t.key, err = t.get_cache_key(r)

	if err != nil {
		return false, err
	}

	t.history, err = t.cache.Get(t.key)

	if err != nil {
		return false, err
	}

	t.now = time.Now()

	// Drop any requests from the history which have now passed the
	// throttle duration
	for len(t.history) > 0 {

		last_time, err := utils.ParseTimestamp(t.history[0])

		if err != nil {
			return false, err
		}

		if last_time.Before(t.now.Add(-time.Second * time.Duration(dur))) {
			t.history = t.history[1:len(t.history)]
		} else {
			break
		}
	}

	if len(t.history) >= req {
		return t.throttle_failure(w, r)
	}
	return t.throttle_success(w, r)
}

func (t *SimpleRateThrottle) wait(w http.ResponseWriter, r *http.Request) (float64, error) {
	var remaining_duration, available_requests int
	if t.history != nil {
		last_time, err := utils.ParseTimestamp(t.history[len(t.history)-1])

		if err != nil {
			return 0, err
		}

		remaining_duration = int(time.Since(*last_time))
		available_requests = req - len(t.history) + 1
	} else {
		remaining_duration = dur
		available_requests = req
	}

	return float64(remaining_duration) / float64(available_requests), nil
}

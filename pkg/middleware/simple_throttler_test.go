package middleware

import (
	"net/http"
	"testing"
	"time"
)

const (
	testkey = "test_key"
)

func getCacheKey(r *http.Request) (string, error) { return testkey, nil }

type testKvs struct {
	kv map[string][]string
}

func (k *testKvs) Push(key string, values []string, time time.Duration) error {
	k.kv[key] = values
	return nil
}

func (k *testKvs) Get(key string) ([]string, error) {
	value, ok := k.kv[key]
	if ok {
		return value, nil
	}
	return make([]string, 0), nil
}

func (k *testKvs) Overwrite(key string, values []string, time time.Duration) error {
	k.kv[key] = values
	return nil
}

func (k *testKvs) Remove(key string) error {
	delete(k.kv, key)
	return nil
}

func TestAllowRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "test.com", nil)
	req.RemoteAddr = "test123"

	kvs := &testKvs{}
	kvs.kv = make(map[string][]string)

	simple_throttle := GetAnonymousThrottle(10, time.Second*10, "test", 10, kvs)

	res, err := simple_throttle.AllowRequest(req)

	if err != nil || !res {
		t.Error("Error in Allow_request")
	}

	res, err = simple_throttle.AllowRequest(req)
	if err != nil || !res {
		t.Error("Error in Allow_request")
	}

	time.Sleep(time.Second * 10)

	for i := 1; i <= 10; i++ {
		res, err := simple_throttle.AllowRequest(req)
		if err != nil || !res {
			t.Error("Allow_request fail before req_limit")
		}
	}

	res, err = simple_throttle.AllowRequest(req)
	if err != nil || res {
		t.Error("Allow_request dont fail after req_limit")
	}
}

func TestWait(t *testing.T) {
	req, _ := http.NewRequest("GET", "test.com", nil)
	req.RemoteAddr = "test123"

	kvs := &testKvs{}
	kvs.kv = make(map[string][]string)

	simple_throttle := GetCustomThrottle(10, time.Second*10, "test", kvs, getCacheKey)

	wait, err := simple_throttle.Wait()
	if err != nil || wait != 1 {
		t.Error("Error in wait")
	}

	simple_throttle.AllowRequest(req)
	wait, err = simple_throttle.Wait()
	if err != nil || wait != 1.1111111111111112 {
		t.Error("Error in wait")
	}
	time.Sleep(time.Second * 3)

	simple_throttle.AllowRequest(req)
	wait, err = simple_throttle.Wait()
	if err != nil || wait != 0.875 {
		t.Error("Error in wait")
	}
}

package middleware

import (
	"net/http"
	"testing"
)

func TestGetIndent(t *testing.T) {
	req, _ := http.NewRequest("GET", "test.com", nil)
	req.RemoteAddr = "test123"
	base_throttle := BaseThrottle{}
	indent := base_throttle.GetIndent(req)
	if indent != req.RemoteAddr {
		t.Error("Incorrect Indent")
	}
}

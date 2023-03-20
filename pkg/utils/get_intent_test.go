package utils

import (
	"net/http"
	"testing"
)

func TestGetIndent(t *testing.T) {
	req, _ := http.NewRequest("GET", "test.com", nil)
	req.RemoteAddr = "test123"
	indent := GetIndent(req, 10)
	if indent != req.RemoteAddr {
		t.Error("Incorrect Indent")
	}
}

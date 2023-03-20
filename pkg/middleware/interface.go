package middleware

import (
	"net/http"
)

type throttle interface {
	AllowRequest(r *http.Request) (bool, error)
	Wait() (float64, error)
}

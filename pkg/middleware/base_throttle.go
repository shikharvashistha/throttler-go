package middleware

import (
	"net/http"
	"strings"

	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

var (
	NUM_PROXIES = 10
)

type BaseThrottle struct {
}

func (t *BaseThrottle) GetIndent(r *http.Request) string {
	xff := r.Header.Get("HTTP_X_FORWARDED_FOR")
	remote_addr := r.RemoteAddr
	num_proxies := NUM_PROXIES

	if num_proxies != 0 {
		if num_proxies == 0 || xff == "" {
			return remote_addr
		}
		addrs := strings.Split(xff, ",")
		client_addr := addrs[-utils.Min(num_proxies, len(addrs))]
		return strings.TrimSpace(client_addr)
	}

	if xff != "" {
		return xff
	}
	return remote_addr

}

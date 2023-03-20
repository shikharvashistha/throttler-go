package utils

import (
	"net/http"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GetIndent(r *http.Request, numProxies int) string {
	xff := r.Header.Get("HTTP_X_FORWARDED_FOR")
	remoteAddr := r.RemoteAddr

	if numProxies != 0 {
		if numProxies == 0 || xff == "" {
			return remoteAddr
		}
		addrs := strings.Split(xff, ",")
		clientAddr := addrs[-min(numProxies, len(addrs))]
		return strings.TrimSpace(clientAddr)
	}

	if xff != "" {
		return xff
	}
	return remoteAddr

}

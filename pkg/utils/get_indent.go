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
	/*
		ref:
			https://gist.github.com/17twenty/c815680c9c585cd9c16e62cbee7317b6,
			https://oxpedia.org/wiki/index.php?title=AppSuite:Grizzly#Multiple_Proxies_in_front_of_the_cluster
	*/
	xff := r.Header.Get("X-Forwarded-For")
	remoteAddr := r.RemoteAddr

	if numProxies != 0 {
		if xff == "" {
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

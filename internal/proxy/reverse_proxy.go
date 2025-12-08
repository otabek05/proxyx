package proxy

import (
	"net/http"
	"net/http/httputil"
)

func reverseProxyxHandler(w http.ResponseWriter, r *http.Request, matched *routeInfo) {
	target := matched.loadBalancer.Next()
	if target == nil {
		ServeProxyHomepage(w)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

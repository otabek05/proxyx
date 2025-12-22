package proxy

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

func reverseProxyxHandler(w http.ResponseWriter, r *http.Request, matched *routeInfo) {
	target := matched.loadBalancer.Next()
	if target == nil {
		ServeProxyHomepage(w)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	if strings.EqualFold(r.Header.Get("Connection"), "Upgrade") &&
		strings.EqualFold(r.Header.Get("Upgrade"), "websocket") {
		r.Header.Set("Connection", "Upgrade")
		r.Header.Set("Upgrade", "websocket")
	}

	proxy.ServeHTTP(w, r)
}

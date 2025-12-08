package proxy

import (
	"ProxyX/internal/common"
	"net"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, r *http.Request, servers map[string][]routeInfo) {
	routes, ok := servers[r.Host]
	if !ok {
		ServeProxyHomepage(w)
		return
	}

	matched := findMatchingRoute(routes, r.URL.Path)
	if matched == nil {
		ServeProxyHomepage(w)
		return
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if !matched.rateLimiter.Allow(ip) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("429 Too Many Requests"))
		return
	}

	switch matched.routeConfig.Type {
	case common.RouteReverseProxy:
		reverseProxyxHandler(w,r, matched)
	case common.RouteStatic:
		staticRouteHandler(w, r , matched)
	default:
		ServeProxyHomepage(w)
	}
}

func findMatchingRoute(routes []routeInfo, path string) *routeInfo {
	for _, rt := range routes {
		routePrefix := rt.routeConfig.Path
		strippedPath := path

		if strings.HasSuffix(routePrefix, "/**") {
			base := strings.TrimSuffix(routePrefix, "/**")
			if strippedPath == base {
				return &rt
			}

			if strings.HasPrefix(strippedPath, base+"/") {
				return &rt
			}

			continue
		} else {
			normalizedRoute := strings.TrimSuffix(routePrefix, "/")
			normalizedPath := strings.TrimSuffix(strippedPath, "/")
			if normalizedPath == normalizedRoute {
				return &rt
			}

			if normalizedRoute == "" {
				return &rt
			}
		}
	}

	return nil
}
package proxy

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)



func websocketProxyHandler(w http.ResponseWriter, r *http.Request, matched *routeInfo) {
	if !strings.EqualFold(r.Header.Get("Connection"), "Upgrade") ||
		!strings.EqualFold(r.Header.Get("Upgrade"), "websocket") {
		http.Error(w, "Not a websocket upgrade request", http.StatusBadRequest)
		return
	}

	u, err := url.Parse(matched.routeConfig.Websocket.URL)
	if err != nil {
		http.Error(w, "Invalid backend URL", http.StatusInternalServerError)
		return
	}

	backendConn, err := net.Dial("tcp", u.Host)
	if err != nil {
		http.Error(w, "Cannot connect to backend", http.StatusBadGateway)
		return
	}
	defer backendConn.Close()

	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hj.Hijack()
	if err != nil {
		http.Error(w, "Hijack failed", http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	r.Write(backendConn)

	go io.Copy(backendConn, clientConn)
	io.Copy(clientConn, backendConn)
}
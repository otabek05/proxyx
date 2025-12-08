package proxy

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func staticRouteHandler(w http.ResponseWriter, r *http.Request, matched *routeInfo) {
	if matched.routeConfig.Static == nil || matched.routeConfig.Static.Root == "" {
		ServeProxyHomepage(w)
		return
	}

	staticDir := filepath.Join(matched.routeConfig.Static.Root)
	requestedFile := filepath.Join(staticDir, r.URL.Path)
	if info, err := os.Stat(requestedFile); err == nil && !info.IsDir() {
		http.ServeFile(w, r, requestedFile)
		return
	}

	indexFile := filepath.Join(staticDir, "index.html")
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		log.Println("index.html not found at", indexFile)
		ServeProxyHomepage(w)
		return
	}

	http.ServeFile(w, r, indexFile)
}
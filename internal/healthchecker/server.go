package healthchecker

import (
	"net/url"
	"sync"
)

type Server struct {
	URL *url.URL
	Health bool
	mu sync.RWMutex
}

func (b *Server) SetHealthy(v bool ) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Health = v
}

func (b *Server) IsHealthy() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Health
}

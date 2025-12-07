package healthchecker

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)


type Registry struct {
	servers []*Server
	mu sync.RWMutex
}

var global = &Registry{}


func RegisterServer(rawURL string) (*Server, error) {
	parsedURL , err := url.Parse(rawURL)
	if err != nil {
		return nil, err 
	}

	backend := &Server{
		URL: parsedURL,
		Health: true,
	}

	global.mu.Lock()
	global.servers = append(global.servers, backend)
	global.mu.Unlock()

	return backend,nil
}

func Start(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func ()  {
		client := &http.Client{Timeout: 500 *time.Millisecond}
		for range ticker.C {
			global.mu.RLock()
			backends := global.servers
			global.mu.RUnlock()

			var wg sync.WaitGroup
			for _, backend :=  range backends {
				wg.Add(1)
				go func (b *Server)  {
					defer wg.Done()
					checkHealth(b, client)
				}(backend)
			}

		}	
	}()
}


func checkHealth(b *Server, client *http.Client) {
	resp , err := client.Get(b.URL.String())
	health := err == nil && resp.StatusCode < 500
	b.SetHealthy(health)
}
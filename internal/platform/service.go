package platform

import (
	"ProxyX/internal/platform/darwin"
	"ProxyX/internal/platform/linux"
	"fmt"
	"runtime"
)

type Service interface {
	Start() error
	Stop() error 
	Restart() error
	Status()  error  
}


func NewService() (Service, error ) {
	switch runtime.GOOS {
	case "linux":
		return linux.New(), nil 
	case "darwin":
		return darwin.New(), nil 
	default:
		return nil, fmt.Errorf("Unsupported OS")
	}
}
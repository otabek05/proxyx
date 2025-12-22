package main

import (
	"ProxyX/internal/cli"
	"ProxyX/internal/proxy"
	"ProxyX/pkg/config"
	"fmt"
	"log"
	"os"
)

func main() {
	requireRoot()
	if len(os.Args) > 1 {
		cli.Execute()
		return
	}

	proxyConfig, err := config.LoadProxyXConfig()
	if err != nil {
		log.Fatalf("Failed to load proxy config: %v", err)
	}

	
	serverConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println(serverConfig)

	srv := proxy.NewServer(serverConfig, proxyConfig)
	srv.Start()
}


func requireRoot() {
	if os.Geteuid() != 0 {
		fmt.Println("This command must be run with sudo (root privileges required)")
		fmt.Println("Example: sudo proxyx services")
		os.Exit(1)
	}
}
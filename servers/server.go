package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Response struct {
	Time string `json:"time"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	clientIP, clientPort, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		clientIP = r.RemoteAddr
		clientPort = ""
	}

	fmt.Printf("[%s] Incoming request: %s %s from %s:%s\n",
		time.Now().Format(time.RFC3339), r.Method, r.URL.Path, clientIP, clientPort)

	resp := Response{
		Time: time.Now().Format(time.RFC3339),
		IP:   clientIP,
		Port: clientPort,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
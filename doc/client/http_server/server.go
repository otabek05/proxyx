package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type ApiResponse struct {
	Message string `json:"message"`
	Data any `json:"data"`
	StatusCode string `json:"status_code"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: ip: %s, path: %s\n ", r.RemoteAddr, r.URL.Path)
		response := &ApiResponse{
			Message: "success",
			Data: "raw data",
			StatusCode: "1002",
		}

		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
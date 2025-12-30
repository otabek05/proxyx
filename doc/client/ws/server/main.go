package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {return true},
}


func wsHandler( w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader error: ", err )
		return 
	}

	defer conn.Close()

	log.Println("Client connected:", r.RemoteAddr)

    for {
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }
        log.Printf("Received: %s\n", msg)

        response := fmt.Sprintf("Server echo: %s", msg)
        if err := conn.WriteMessage(msgType, []byte(response)); err != nil {
            log.Println("Write error:", err)
            break
        }
    }
}


func main() {
    http.HandleFunc("/ws", wsHandler)

    log.Println("WebSocket server listening on :9000")
    log.Fatal(http.ListenAndServe(":9000", nil))
}
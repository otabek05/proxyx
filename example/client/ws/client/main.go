package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost/ws/chat"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	var mu sync.Mutex

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Read error:", err)
				return
			}
			mu.Lock()
			fmt.Println("Received:", string(msg))
			mu.Unlock()
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		mu.Lock()
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')
		mu.Unlock()
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}

		time.Sleep(1 *time.Second)
	}
}

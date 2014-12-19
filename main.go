package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if err = conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}()
}

func main() {
	http.HandleFunc("/ws", webSocketHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":1234", nil))
}

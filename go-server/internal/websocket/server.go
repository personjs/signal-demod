package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Start(addr string, hub *Hub) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade error:", err)
			return
		}
		client := &Client{conn: conn, send: make(chan []byte, 256), hub: hub}
		hub.register <- client
		go client.writePump()
	})

	log.Println("🔌 WebSocket server started on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

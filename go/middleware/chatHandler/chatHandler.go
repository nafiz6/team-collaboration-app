package chatHandler

import (
	"fmt"
	"log"
	"net/http"

	"teams/middleware/accountsHandler"
	. "teams/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

// define our WebSocket endpoint
func ServeWs(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CHAT CLIENT REQ")
	fmt.Println(r.Host)
	UID := accountsHandler.GetUserId(r)
	fmt.Println("UID: " + UID)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	client := &Client{
		ID:    UID,
		Conn:  ws,
		Pools: make(map[string]*Pool),
	}

	client.Read()
}

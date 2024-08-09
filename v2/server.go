package grid_cli

// Simplified overview of the system design based on the discussions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Serve() {
	// Start the WebSocket server
	http.HandleFunc("/ws", HandleWebSocket)
	fmt.Println("WebSocket server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a WebSocket connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Create a new client
	client := NewClient(conn)
	kernel := NewKernel()
	// kernel.AddClient(client)
	_ = kernel // XXX
	_ = client // XXX
}

// NewClient creates a new client
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		send: make(chan []byte),
	}
}

// Client represents a WebSocket client
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

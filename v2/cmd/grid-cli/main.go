package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/stevegt/grid-cli/v2"
)

func main() {
	ctx := context.Background()
	kernel := grid_cli.NewKernel()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(ctx, kernel, w, r)
	})

	fmt.Println("WebSocket server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleWebSocket(ctx context.Context, k *grid_cli.Kernel, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to websocket:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		parms, err := grid_cli.DeserializeMessage(message)
		if err != nil {
			fmt.Println("Failed to deserialize message:", err)
			continue
		}

		data, err := k.ConsultModules(ctx, parms...)
		if err != nil {
			fmt.Printf("Failed to consult modules: %v\n", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			fmt.Println("Failed to write message:", err)
			break
		}
	}
}

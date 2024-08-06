package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/spf13/afero"
	"github.com/stevegt/grid-cli/v2"
)

// Kernel struct with the syscall tree root and file system abstraction
type Kernel struct {
	root    *v2.SyscallNode
	fs      afero.Fs
	modules map[string]v2.Module // Known modules
}

// NewKernel initializes a new Kernel instance with embedded modules
func NewKernel() *Kernel {
	return &Kernel{
		root: &v2.SyscallNode{
			Children: make(map[string]*v2.SyscallNode),
			Modules:  []v2.Module{},
		},
		fs:      afero.NewOsFs(),
		modules: make(map[string]v2.Module),
	}
}

func (k *Kernel) addSyscall(parms ...interface{}) {
	current := k.root
	for _, parm := range parms {
		key := fmt.Sprintf("%v", parm)
		if _, exists := current.Children[key]; !exists {
			current.Children[key] = &v2.SyscallNode{
				Children: make(map[string]*v2.SyscallNode),
				Modules:  []v2.Module{},
			}
		}
		current = current.Children[key]
	}
	// Assuming module is pre-initialized and available in context
	if module, exists := k.modules[fmt.Sprintf("%v", parms[len(parms)-1])]; exists {
		current.Modules = append(current.Modules, module)
	}
}

func (k *Kernel) findBestMatch(parms ...interface{}) *v2.SyscallNode {
	current := k.root
	for _, parm := range parms {
		key := fmt.Sprintf("%v", parm)
		if next, exists := current.Children[key]; exists {
			current = next
		} else {
			break
		}
	}
	return current
}

func (k *Kernel) consultModules(ctx context.Context, parms ...interface{}) ([]byte, error) {
	bestMatch := k.findBestMatch(parms...)
	var promisingModules []v2.Module

	for _, module := range bestMatch.Modules {
		promise, err := module.Accept(ctx, parms...)
		if err != nil {
			continue // Log and handle error
		}
		if promise.Parms[0].(bool) {
			promisingModules = append(promisingModules, module)
			k.addSyscall(parms...) // Add to syscall tree
		}
	}

	for _, module := range promisingModules {
		data, err := module.HandleMessage(ctx, parms...)
		if err == nil {
			return data, nil
		}
		// Log broken promise if HandleMessage fails after Accept returned true
	}

	return nil, fmt.Errorf("no module could handle the request")
}

func handleWebSocket(ctx context.Context, k *Kernel, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade to WebSocket: %v\n", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Failed to read message: %v\n", err)
			break
		}
		parms, err := v2.DeserializeMessage(message)
		if err != nil {
			fmt.Printf("Failed to deserialize message: %v\n", err)
			continue
		}

		// Consult modules based on the message parms
		response, err := k.consultModules(ctx, parms...)
		if err != nil {
			fmt.Printf("Failed to consult modules: %v\n", err)
			continue
		}

		// Response handling (omitted for brevity)
		fmt.Println("Message processed, response generated.", response)
	}
}

func main() {
	// Example for starting the WebSocket server
	ctx := context.Background()
	k := NewKernel()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(ctx, k, w, r)
	})
	fmt.Println("Starting WebSocket server on :8080...")
	http.ListenAndServe(":8080", nil)
}

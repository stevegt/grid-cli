package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/spf13/afero"
)

// SyscallNode represents a node in the hierarchical syscall tree
type SyscallNode struct {
	modules  []Module
	children map[string]*SyscallNode
}

// Kernel struct with the syscall tree root and file system abstraction
type Kernel struct {
	root    *SyscallNode
	fs      afero.Fs
	modules map[string]Module // Known modules
}

// NewKernel initializes a new Kernel instance with embedded modules
func NewKernel() *Kernel {
	return &Kernel{
		root: &SyscallNode{
			children: make(map[string]*SyscallNode),
		},
		fs:      afero.NewOsFs(),
		modules: make(map[string]Module), // Initialize with known modules
	}
}

func (k *Kernel) addSyscall(parms ...interface{}) {
	current := k.root
	for _, parm := range parms {
		key := fmt.Sprintf("%v", parm)
		if _, exists := current.children[key]; !exists {
			current.children[key] = &SyscallNode{
				children: make(map[string]*SyscallNode),
			}
		}
		current = current.children[key]
	}
	// Append the module to the leaf node
	// Assuming module is pre-initialized and available in context
	current.modules = append(current.modules, k.modules[fmt.Sprintf("%v", parms[len(parms)-1])])
}

func (k *Kernel) findBestMatch(parms ...interface{}) *SyscallNode {
	current := k.root
	for _, parm := range parms {
		key := fmt.Sprintf("%v", parm)
		if next, exists := current.children[key]; exists {
			current = next
		} else {
			break
		}
	}
	return current
}

func (k *Kernel) consultModules(ctx context.Context, parms ...interface{}) ([]byte, error) {
	bestMatch := k.findBestMatch(parms...)
	var promisingModules []Module

	for _, module := range bestMatch.modules {
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
		parms, err := deserializeMessage(message)
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

func deserializeMessage(data []byte) ([]interface{}, error) {
	var parms []interface{}
	if err := json.Unmarshal(data, &parms); err != nil {
		return nil, err
	}
	return parms, nil
}

func serializeMessage(parms []interface{}) ([]byte, error) {
	data, err := json.Marshal(parms)
	if err != nil {
		return nil, err
	}
	return data, nil
}

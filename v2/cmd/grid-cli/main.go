package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/spf13/afero"
)

const (
	cacheDir   = "cache"
	configFile = ".grid/config"
)

// Message defines the structure for communication messages.
// The first element in Parms is the promise of handling or not handling the syscall.
type Message struct {
	Parms   []interface{}          `json:"parms"`   // Parameters: first element is the promise
	Payload map[string]interface{} `json:"payload"` // Metadata about the promise
}

// SyscallNode represents a node in the hierarchical syscall tree
type SyscallNode struct {
	modules  []Module
	children map[string]*SyscallNode
}

// Kernel struct with the syscall tree root and file system abstraction
type Kernel struct {
	root          *SyscallNode
	fs            afero.Fs
	knownMessages map[string]Message // Known messages for acceptance
	modules       []Module           // List of modules including built-ins
}

// Module is an interface with Accept and HandleMessage functions
type Module interface {
	Accept(ctx context.Context, parms ...interface{}) (Message, error)
	HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error)
}

// NewKernel initializes a new Kernel instance with embedded modules
func NewKernel(knownMessages map[string]Message) *Kernel {
	fs := afero.NewOsFs()
	return &Kernel{
		root: &SyscallNode{
			children: make(map[string]*SyscallNode),
		},
		fs:            fs,
		knownMessages: knownMessages,
		modules:       []Module{NewLocalCacheModule(fs, cacheDir)}, // Example built-in module
	}
}

// LocalCacheModule represents a local cache as a module
type LocalCacheModule struct {
	fs       afero.Fs
	cacheDir string
}

func NewLocalCacheModule(fs afero.Fs, cacheDir string) *LocalCacheModule {
	return &LocalCacheModule{fs: fs, cacheDir: cacheDir}
}

func (m *LocalCacheModule) Accept(ctx context.Context, parms ...interface{}) (Message, error) {
	// Implement logic to accept or reject based on parms
	// For simplicity, accept all parms
	return Message{Parms: append([]interface{}{true}, parms...), Payload: map[string]interface{}{"info": "local cache"}}, nil
}

func (m *LocalCacheModule) HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error) {
	if len(parms) < 2 {
		return nil, fmt.Errorf("insufficient arguments")
	}

	promiseHash, ok1 := parms[0].([]byte)
	moduleHash, ok2 := parms[1].([]byte)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("invalid argument types")
	}

	cacheKey := constructCacheKey(promiseHash, moduleHash, parms[2:]...)
	return loadFromCache(m.fs, cacheKey)
}

func constructCacheKey(promiseHash, moduleHash []byte, args ...interface{}) string {
	var keyBuilder strings.Builder
	keyBuilder.WriteString(fmt.Sprintf("%x", promiseHash))
	keyBuilder.WriteString("/")
	keyBuilder.WriteString(fmt.Sprintf("%x", moduleHash))

	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			keyBuilder.WriteString("/")
			keyBuilder.WriteString(url.QueryEscape(v))
		case []byte:
			keyBuilder.WriteString("/")
			keyBuilder.WriteString(url.QueryEscape(string(v)))
		default:
			// Handle unsupported argument types
		}
	}

	return keyBuilder.String()
}

func loadFromCache(fs afero.Fs, cacheKey string) ([]byte, error) {
	filePath := afero.GetTempDir(fs, cacheDir) + "/" + cacheKey
	return afero.ReadFile(fs, filePath)
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
	// Append the module to the leaf node (for simplicity, assuming a default module exists)
	current.modules = append(current.modules, k.modules[0])
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
	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, err
	}
	return msg.Parms, nil
}

func main() {
	knownMessages := map[string]Message{
		"accept": {Parms: []interface{}{true}, Payload: map[string]interface{}{"info": "accepted"}},
		"reject": {Parms: []interface{}{false}, Payload: map[string]interface{}{"info": "rejected"}},
	}

	kernel := NewKernel(knownMessages)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(context.Background(), kernel, w, r)
	})
	fmt.Println("WebSocket server started on :8080")
	http.ListenAndServe(":8080", nil)
}

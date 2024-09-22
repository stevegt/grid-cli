# PromiseGrid Trie Node Structure

## Introduction

The PromiseGrid's Trie Node structure is central to efficiently storing and retrieving promises, module hashes, and arguments within the decentralized system. This document details the design and functionality of the trie node structure, including its components and the importance of reputation accounting.

## Trie Node Structure

### Definition

A Trie Node in PromiseGrid contains key elements crucial for efficient promise handling and storage. The node structure is designed to ensure quick access and modification while supporting the system's dynamic requirements.

```go
type TrieNode struct {
    KeyPart           byte                  // Part of the key corresponding to this node
    Children          map[byte]*TrieNode    // Map of child nodes
    IsHandler         bool                  // Flag indicating if this node is a handler
    MultipleHandlers  []HandlerFunc         // List of handler functions if this node has multiple handlers
    RawValue          interface{}           // Raw value if this node is not a handler
    Handler           func(prefix []byte, remainder io.Reader) interface{}  // Handler function if this node is a handler
    Reputation        map[string]ReputationStats  // Reputation accounting for handlers
}
```

### Components

1. **KeyPart**:
    - Represents a single byte of the key stored in the node.

2. **Children**:
    - A map that associates each byte part with its child node.

3. **IsHandler**:
    - A boolean flag that signifies whether the node is associated with a handler function.

4. **MultipleHandlers**:
    - Stores a list of handler functions if the node is associated with multiple handlers.

5. **RawValue**:
    - Stores the value directly within the node if it is not a handler.

6. **Handler**:
    - A function that dynamically generates or retrieves the value upon invocation.
    - The handler receives the prefix (path leading to the handler node) as a byte slice and the remainder (additional key parts beyond the handler) as an `io.Reader` for context-sensitive operations.

7. **Reputation**:
    - Keeps track of promises filled and broken and the value of each promise (which might be denominated in a personal currency or points issued by the requestor).
    - Used for making decisions about which handler(s) to route a message to or which handler's results to use.

### Reputation Structure

The reputation structure helps in tracking the performance and reliability of each handler. It includes:

```go
type ReputationStats struct {
    PromisesFilled int
    PromisesBroken int
    TotalValue     float64 // Can be denominated in personal currency/points issued by the requestor
}
```

## Handler Call Signature

### Definition

The handler enables nodes to return context-specific values when part of the trie is traversed. The handler is defined as follows:

```go
type HandlerFunc func(prefix []byte, remainder io.Reader) interface{}
```

### Parameters

- **prefix**: A byte slice representing the path leading to the handler node.
- **remainder**: An `io.Reader` that provides access to the remaining key parts beyond the handler.

### Example Usage

Consider a scenario where the trie stores various configurations, and each configuration can be statically defined or dynamically generated:

```go
// Define a handler function
func configHandler(prefix []byte, remainder io.Reader) interface{} {
    remainderBytes, _ := io.ReadAll(remainder) // Read the remainder of the path
    fullPath := append(prefix, remainderBytes...)
    return fmt.Sprintf("Dynamic config for %s", string(fullPath))
}

// Example TrieNode setup
root := &TrieNode {
    KeyPart:   0x00,
    Children:  make(map[byte]*TrieNode),
    IsHandler: false,
}

configNode := &TrieNode {
    KeyPart:   0x01,
    Children:  make(map[byte]*TrieNode),
    IsHandler: true,
    Handler:   configHandler,
}

root.Children[0x01] = configNode

// Traversing the trie and invoking the handler
path := "/config/somepath"
result := root.Children[0x01].Handler([]byte("/config"), strings.NewReader("somepath"))
fmt.Println(result)  // Output: Dynamic config for /config/somepath
```

## Pros and Cons of Having Multiple Handlers per Node

### Pros

1. **Flexibility**:
    - Having multiple handlers per node allows for more flexible and context-specific responses.
    - Enables different handlers to be executed based on varying conditions, such as user permissions, request types, or runtime states.

2. **Extensibility**:
    - New handlers can be added to handle additional scenarios or requirements.

3. **Modularity**:
    - Facilitates modular design by separating concerns within handlers.

4. **Reputation Accounting**:
    - Allows the system to track the performance of each handler.
    - Decision-making can be improved by considering the reputation of handlers, selecting or prioritizing those with better track records.



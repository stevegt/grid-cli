# PromiseGrid Trie Node Structure

## Introduction

The PromiseGrid's Trie Node structure is central to efficiently storing and retrieving promises, module hashes, and arguments within the decentralized system. This document details the design and functionalities of the trie node structure, including a flag to signify whether a node represents a raw value or a handler that must be invoked to retrieve the value.

## Trie Node Structure

### Definition

A Trie Node in PromiseGrid contains key elements crucial for efficient promise handling and storage. The node structure is designed to ensure quick access and modification while supporting the system's dynamic requirements.

```go
type TrieNode struct {
    KeyPart    byte                  // Part of the key corresponding to this node
    Children   map[byte]*TrieNode    // Map of child nodes
    IsHandler  bool                  // Flag indicating if this node is a handler
    RawValue   interface{}           // Raw value if this node is not a handler
    Handler    func(path string) interface{}  // Handler function if this node is a handler
}
```

### Components

1. **KeyPart**:
    - Represents a single byte of the key stored in the node.
    - Allows breaking down the key into a sequence of nodes for efficient retrieval.

2. **Children**:
    - A map that associates each byte part to its child node in the trie.
    - Supports quick traversal to deeper nodes based on successive key parts.

3. **IsHandler**:
    - A boolean flag that signifies whether the node is associated with a handler function.
    - `true` if the node is a handler; `false` if it represents a raw value.

4. **RawValue**:
    - Stores the value directly within the node if it is not a handler.
    - Used when `IsHandler` is `false`.

5. **Handler**:
    - A function that dynamically generates or retrieves the value upon invocation.
    - Used when `IsHandler` is `true`.
    - The handler requires the full trie path (both prefix and suffix) as an argument for context-sensitive operations.

## Handler Call Signature

### Signature Definition

The handler is a crucial component in dynamic value retrieval. It enables nodes to return context-specific values when part of the trie is traversed. The handler is defined as follows:

```go
type HandlerFunc func(path string) interface{}
```

### Path Representation

The `path` parameter provided to the handler function is a string that concatenates the prefix leading to the handler node and the suffix representing any additional key parts beyond the handler. This comprehensive path enables the handler to understand the full context it operates within.

### Example Usage

Consider a scenario where the trie stores various configurations, and each configuration can be statically defined or dynamically generated:

```go
// Define a handler function
func configHandler(fullPath string) interface{} {
    // Use the full path to generate or retrieve the value
    return fmt.Sprintf("Dynamic config for %s", fullPath)
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
result := root.Children[0x01].Handler(path)
fmt.Println(result)  // Output: Dynamic config for /config/somepath
```

### Advantages

- **Dynamic Value Generation**: Handlers allow for real-time value creation based on the full context provided by the path.
- **Storage Efficiency**: By storing handlers instead of all possible values, the system can optimize memory usage and processing time.
- **Flexibility**: Handlers enable the trie to manage diverse and dynamic configurations seamlessly.

## Conclusion

The PromiseGrid Trie Node structure is designed to efficiently handle both static and dynamic values within a decentralized system. By incorporating flags and handler functions, the trie nodes provide a flexible and powerful mechanism for promise management, facilitating quick lookups, efficient storage, and dynamic value generation.

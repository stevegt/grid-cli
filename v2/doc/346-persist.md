# Persisting the Trie to Disk

## Introduction

This document describes various options for persisting the trie data structure to disk within the PromiseGrid system. The design aims to function in both native and WebAssembly (WASM) runtimes. Several strategies are discussed, including the use of the Origin Private File System (OPFS) and the `afero` library.

## Design Considerations for Persistence

The chosen persistence strategy must meet the following requirements:
1. **Compatibility**: Works seamlessly in both native and WASM environments.
2. **Reliability**: Ensures data integrity and durability.
3. **Performance**: Provides efficient read and write operations.
4. **Flexibility**: Supports dynamic updates and retrievals.

## Potential Persistence Strategies

### 1. Origin Private File System (OPFS)

**Description**: OPFS is a file system designed for use in web environments, providing secure, performant, and private storage.

**Advantages**:
- **Security**: Provides a secure sandboxed environment.
- **Performance**: Optimized for web performance.
- **Compatibility**: Works well with WASM.

**Implementation**:
- **WASM**: Utilize OPFS directly via JavaScript to interact with the file system.
- **Native**: Implement a thin abstraction layer to handle file I/O using standard file system calls.

### 2. `afero` Library

**Description**: `afero` is a filesystem abstraction library for Go, providing a common interface for various file system operations.

**Advantages**:
- **Abstraction**: Abstracts filesystem operations, making the code agnostic to the underlying storage backend.
- **Flexibility**: Supports multiple backends like OS file systems, in-memory systems, and more.
- **Compatibility**: Can be used in native environments; minimal work is needed to extend it to WASM.

**Implementation**:
- **WASM**: Create a custom backend for `afero` that communicates with the browser’s file API (OPFS).
- **Native**: Use the default OS filesystem backend or other suitable backends provided by `afero`.

## Example Implementations

### OPFS in WASM

Here’s an example of interacting with OPFS in a WASM environment:

```javascript
// JavaScript code to interact with OPFS in a WASM environment

async function writeFile(filename, content) {
  const handle = await navigator.storage.getDirectory();
  const fileHandle = await handle.getFileHandle(filename, { create: true });
  const writable = await fileHandle.createWritable();
  await writable.write(content);
  await writable.close();
}

async function readFile(filename) {
  const handle = await navigator.storage.getDirectory();
  const fileHandle = await handle.getFileHandle(filename);
  const file = await fileHandle.getFile();
  return await file.text();
}
```

### `afero` in Native

Here’s an example using `afero` in a native Go application:

```go
package main

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
)

func main() {
	fs := afero.NewOsFs()

	// Writing to the file
	filePath := "example.txt"
	err := afero.WriteFile(fs, filePath, []byte("Hello, Afero!"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Reading from the file
	content, err := afero.ReadFile(fs, filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(content))
}
```

## Persisting the Trie to Disk in Go

Below is a Go code example, based on `trie.go`, that persists the trie to disk using the `afero` library:

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/spf13/afero"
	"io"
)

type TrieNode struct {
	Children map[byte]*TrieNode
	Handlers []Handler
}

type Trie struct {
	Root *TrieNode
	Fs   afero.Fs
}

type Handler func()

func NewTrie(fs afero.Fs) *Trie {
	return &Trie{
		Root: &TrieNode{Children: make(map[byte]*TrieNode), Handlers: []Handler{}},
		Fs:   fs,
	}
}

func (t *Trie) Insert(r io.Reader, handler Handler) error {
	node := t.Root
	buf := make([]byte, 1)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if n > 0 {
			b := buf[0]
			if _, ok := node.Children[b]; !ok {
				node.Children[b] = &TrieNode{Children: make(map[byte]*TrieNode), Handlers: []Handler{}}
			}
			node = node.Children[b]
		}
	}
	node.Handlers = append(node.Handlers, handler)
	return nil
}

func (t *Trie) Search(r io.Reader) ([]Handler, error) {
	node := t.Root
	buf := make([]byte, 1)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if n > 0 {
			b := buf[0]
			if child, ok := node.Children[b]; ok {
				node = child
				if len(node.Handlers) > 0 {
					return node.Handlers, nil
				}
			} else {
				break
			}
		}
	}
	return nil, nil
}

func (t *Trie) Save(filename string) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(t.Root); err != nil {
		return err
	}

	return afero.WriteFile(t.Fs, filename, buf.Bytes(), 0644)
}

func (t *Trie) Load(filename string) error {
	data, err := afero.ReadFile(t.Fs, filename)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(&t.Root)
}

func main() {
	fs := afero.NewOsFs()
	trie := NewTrie(fs)
	trie.Insert(bytes.NewReader([]byte{0x01, 0x02}), func() { fmt.Println("Handling Message Type A") })
	trie.Insert(bytes.NewReader([]byte{0x03, 0x04}), func() { fmt.Println("Handling Message Type B") })

	// Save trie to disk
	if err := trie.Save("trie_data.gob"); err != nil {
		fmt.Println("Error saving trie:", err)
		return
	}

	// Load trie from disk
	newTrie := NewTrie(fs)
	if err := newTrie.Load("trie_data.gob"); err != nil {
		fmt.Println("Error loading trie:", err)
		return
	}

	// Search in the loaded trie
	message := bytes.NewReader([]byte{0x01, 0x02})
	handlers, err := newTrie.Search(message)
	if err != nil {
		fmt.Println("Search error:", err)
		return
	}
	if handlers != nil {
		for _, handler := range handlers {
			handler()
		}
	} else {
		fmt.Println("No handler found")
	}
}
```

### Conclusion

By leveraging the strengths of OPFS and `afero`, the persistence of the trie data structure can be efficiently managed across both native and WASM environments. OPFS provides a secure and performant option for web environments, while `afero` offers flexibility and abstraction for native applications. This dual approach ensures that the PromiseGrid system remains robust, adaptable, and efficient in diverse runtime conditions.

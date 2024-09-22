# Persisting the Trie to Disk

## Design Considerations

Persistence strategy must address:
1. **Compatibility**: Across Native and WebAssembly (WASM) environments.
2. **Reliability**: Ensuring data integrity and durability.
3. **Performance**: Efficient read/write operations.
4. **Flexibility**: Support for dynamic updates and retrievals.

## Persistence Strategies

### 1. Origin Private File System (OPFS)

* **Security**: Sandboxed, ensuring data privacy.
* **Performance**: Optimized for web use.
* **Compatibility**: WASM-compatible.

**Implementation**:
```javascript
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

### 2. `afero` Library

* **Abstraction**: Generic interface for various storage backends.
* **Flexibility**: Supports multiple backends.
* **Compatibility**: Native environments with minor adaptations for WASM.

**Implementation**:
```go
package main

import (
	"fmt"
	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()

	err := afero.WriteFile(fs, "example.txt", []byte("Hello, Afero!"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	content, err := afero.ReadFile(fs, "example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(content))
}
```

## Persisting Trie to Disk in Go

### Lazy-Loading Nodes

Nodes are loaded from disk as needed to optimize memory usage.

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/spf13/afero"
	"io"
	"path"
)

type TrieNode struct {
	Children map[byte]*TrieNode
	Handlers []Handler
	Path     string
	Fs       afero.Fs
}

type Trie struct {
	Root    *TrieNode
	Fs      afero.Fs
	BaseDir string
}

type Handler func()

func NewTrie(fs afero.Fs, baseDir string) *Trie {
	return &Trie{
		Root: &TrieNode{Children: make(map[byte]*TrieNode), Handlers: []Handler{}, Path: baseDir, Fs: fs},
		Fs: fs,
		BaseDir: baseDir,
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
				node.Children[b] = &TrieNode{Children: make(map[byte]*TrieNode), Handlers: []Handler{}, Path: path.Join(node.Path, fmt.Sprintf("%d", b)), Fs: t.Fs}
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
				if child == nil {
					loadedNode, err := t.LoadNode(path.Join(node.Path, fmt.Sprintf("%d", b)))
					if err != nil {
						return nil, err
					}
					node.Children[b] = loadedNode
				}
				node = node.Children[b]
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

func (t *Trie) SaveNode(node *TrieNode) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(node); err != nil {
		return err
	}
	return afero.WriteFile(t.Fs, node.Path, buf.Bytes(), 0644)
}

func (t *Trie) SaveNodesRecursively(node *TrieNode) error {
	if err := t.SaveNode(node); err != nil {
		return err
	}
	for _, childNode := range node.Children {
		if childNode != nil {
			if err := t.SaveNodesRecursively(childNode); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Trie) Save() error {
	return t.SaveNodesRecursively(t.Root)
}

func (t *Trie) LoadNode(nodePath string) (*TrieNode, error) {
	data, err := afero.ReadFile(t.Fs, nodePath)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	node := &TrieNode{}
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(node); err != nil {
		return nil, err
	}
	node.Path = nodePath
	node.Fs = t.Fs
	for key := range node.Children {
		node.Children[key] = nil // Mark for lazy load
	}
	return node, nil
}

func (t *Trie) Load() error {
	rootNode, err := t.LoadNode(t.BaseDir)
	if err != nil {
		return err
	}
	t.Root = rootNode
	return nil
}

func main() {
	fs := afero.NewOsFs()
	trie := NewTrie(fs, "trie_data")

	trie.Insert(bytes.NewReader([]byte{0x01, 0x02}), func() { fmt.Println("Handling Message Type A") })
	trie.Insert(bytes.NewReader([]byte{0x03, 0x04}), func() { fmt.Println("Handling Message Type B") })

	if err := trie.Save(); err != nil {
		fmt.Println("Error saving trie:", err)
		return
	}

	newTrie := NewTrie(fs, "trie_data")
	if err := newTrie.Load(); err != nil {
		fmt.Println("Error loading trie:", err)
		return
	}

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

### Summary

Combining OPFS and `afero` provides robust trie persistence across both native and WASM environments. Lazy-loading optimizes memory usage.

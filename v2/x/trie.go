package main

import (
	"bytes"
	"fmt"
	"io"
)

type Handler func()

type TrieNode struct {
	children map[byte]*TrieNode
	handlers []Handler
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[byte]*TrieNode), handlers: []Handler{}}}
}

func (t *Trie) Insert(r io.Reader, handler Handler) error {
	node := t.root
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
			if _, ok := node.children[b]; !ok {
				node.children[b] = &TrieNode{children: make(map[byte]*TrieNode), handlers: []Handler{}}
			}
			node = node.children[b]
		}
	}
	node.handlers = append(node.handlers, handler)
	return nil
}

func (t *Trie) Search(r io.Reader) ([]Handler, error) {
	node := t.root
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
			if child, ok := node.children[b]; ok {
				node = child
				if len(node.handlers) > 0 {
					return node.handlers, nil
				}
			} else {
				break
			}
		}
	}
	return nil, nil
}

func main() {
	trie := NewTrie()
	trie.Insert(io.NopCloser(bytes.NewReader([]byte{0x01, 0x02})), func() { fmt.Println("Handling Message Type A") })
	trie.Insert(io.NopCloser(bytes.NewReader([]byte{0x03, 0x04})), func() { fmt.Println("Handling Message Type B") })

	message := io.NopCloser(bytes.NewReader([]byte{0x01, 0x02}))
	handlers, err := trie.Search(message)
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

package main

import (
	"fmt"
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

func (t *Trie) Insert(sequence []byte, handler Handler) {
	node := t.root
	for _, b := range sequence {
		if _, ok := node.children[b]; !ok {
			node.children[b] = &TrieNode{children: make(map[byte]*TrieNode), handlers: []Handler{}}
		}
		node = node.children[b]
	}
	node.handlers = append(node.handlers, handler)
}

func (t *Trie) Search(message []byte) []Handler {
	node := t.root
	for _, b := range message {
		if child, ok := node.children[b]; ok {
			node = child
			if len(node.handlers) > 0 {
				return node.handlers
			}
		} else {
			break
		}
	}
	return nil
}

func main() {
	trie := NewTrie()
	trie.Insert([]byte{0x01, 0x02}, func() { fmt.Println("Handling Message Type A") })
	trie.Insert([]byte{0x03, 0x04}, func() { fmt.Println("Handling Message Type B") })

	message := []byte{0x01, 0x02}
	handlers := trie.Search(message)
	if handlers != nil {
		for _, handler := range handlers {
			handler()
		}
	} else {
		fmt.Println("No handler found")
	}
}

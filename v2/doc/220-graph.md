# Sequence Matching Graph for PromiseGrid

## Overview

A sequence matching graph is used to match input sequences to stored sequences by traversing nodes and edges efficiently. Each node and edge in the sequence-matching graph is composed to facilitate efficient sequence matching.

## Node Composition

Each node in the sequence-matching graph includes:

- **Label**: Represents the segment of the sequence.
- **Occurrences**: List of positions in the reference sequences where this segment occurs.
- **Children**: Map from sequence characters to child nodes, representing the continuation of the sequence.

```go
type Node struct {
    Label      string
    Occurrences []int
    Children   map[rune]*Node
}
```

## Edge Composition

Each edge represents the next segment in the sequence.

- **Character**: The character from the parent node's label to this edge.
- **ChildNode**: Reference to the child node.

```go
type Edge struct {
    Character  rune
    ChildNode *Node
}
```

## Structure Example

Example with "ACGT" and "AGT" sequences:

```plaintext
       (root)
       /   \
      A     A
     /       \
    C         G
     \         \
      G         T
       \
        T
```

## Input Sequence Matching

To match an input sequence:

1. **Traversal**: Start from the root and traverse the graph using the characters of the input sequence.
2. **Edge Matching**: Follow the edge that matches the next character at each node.
3. **Match Determination**: Terminate without a match if no suitable edge or child node is found.

## Traversal Mechanism

Graph traversal can be recursive or iterative.

### Recursive Method

```go
func MatchSequence(node *Node, sequence string, index int) bool {
    if index == len(sequence) {
        return true
    }
    char := rune(sequence[index])
    if child, ok := node.Children[char]; ok {
        return MatchSequence(child, sequence, index+1)
    }
    return false
}
```

### Iterative Method

```go
func MatchSequenceIterative(root *Node, sequence string) bool {
    currentNode := root
    for _, char := range sequence {
        if child, ok := currentNode.Children[char]; ok {
            currentNode = child
        } else {
            return false
        }
    }
    return true
}
```

## Matching Examples

1. **"ACGT"**: Traverse edges labeled 'A', 'C', 'G', 'T'. Successful match if all edges are followed.
2. **"AGT"**: Traverse edges labeled 'A', 'G', 'T'. Successful match if all edges are followed.

## Adding a New Sequence

To add a new sequence "AGC":

```go
func AddSequence(root *Node, sequence string) {
    currentNode := root
    for _, char := range sequence {
        if _, ok := currentNode.Children[char]; !ok {
            newNode := &Node{
                Label:      string(char),
                Occurrences: []int{},
                Children:   make(map[rune]*Node),
            }
            currentNode.Children[char] = newNode
        }
        currentNode = currentNode.Children[char]
    }
}

// Example usage:
root := &Node{Children: make(map[rune]*Node)}
AddSequence(root, "AGC")
```

## Advanced Questions Answered

### Children Field

The 'Children' field contains only the next node(s) in the sequence, mapping sequence characters to their child nodes.

### Genome Storage

The entire genome is not stored -- only sequences of interest are. These sequences are identified based on specific criteria such as the presence of particular motifs or functional relevance.

## Optimizations and Further Research

### Matching with Unknown Leading Characters

To match a fragment with unknown leading characters:

1. **Segment the Fragment**: Divide into smaller overlapping segments.
2. **Search for Segments**: Locate nodes corresponding to the segments in the graph.
3. **Extend Matches**: Extend partial matches by traversing adjacent nodes until a complete match is found.

### Example with Unknown Leading Characters

For "CGTA":

1. **Segment**: "CGT", "GTA".
2. **Search**: Find nodes for "CGT" and "GTA".
3. **Extend**: Form the complete sequence by exploring neighboring nodes.

This approach ensures robust sequence matching capabilities.

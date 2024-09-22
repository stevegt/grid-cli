# Sequence Matching Graph for PromiseGrid

## Overview

The sequence-matching graph in PromiseGrid is a data structure used to match input sequences to stored sequences by traversing nodes and edges efficiently. Each node and edge in the sequence-matching graph has specific compositions that facilitate efficient sequence matching.

## Node Composition

Each node in the sequence-matching graph contains:

- **Label**: A string representing the segment of the sequence at this node.
- **Occurrences**: A list of positions in the reference sequences where this segment occurs.
- **Children**: A map from sequence characters to child nodes, representing the continuation of the sequence.

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

## Example Structure

Consider a simplified example with "ACGT" and "AGT" sequences:

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

To match an input sequence to stored sequences:

1. **Traversal**: Start from the root and traverse through the graph using the characters of the input sequence.
2. **Edge Matching**: At each node, follow the edge that matches the next character.
3. **Match Determination**: If a suitable edge or child node cannot be found, the matching process terminates without a match.

## Graph Traversal Mechanism

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

## Examples of Matching

1. **Matching "ACGT"**: Traverse edges labeled 'A', 'C', 'G', and 'T' from the root. Match is successful if all edges are followed.
2. **Matching "AGT"**: Traverse edges labeled 'A', 'G', and 'T' from the root. Match is successful if all edges are followed.

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

### Sequence-Matching Graph and 'Children' Field

The 'Children' field of a node contains only the next node(s) in the sequence. It maps sequence characters to their corresponding child nodes, representing the continuation of the sequence.

### Genome Storage in Graph

The entire genome is not stored -- only the sequences of interest are stored in the graph. These sequences are identified based on specific criteria such as the presence of particular motifs or functional relevance, which can be based on their positions in the genome or other biological markers.

## Optimizations and Further Research

### Sequence Matching with Unknown Leading Characters

To match a sequence fragment with unknown leading characters:

1. **Segment the Fragment**: Divide the input sequence fragment into smaller overlapping segments.
2. **Search for Segments**: Locate nodes corresponding to the segments within the graph.
3. **Extend Matches**: Extend partial matches by traversing adjacent nodes in the graph till a complete match is found.

### Practical Example with Unknown Leading Characters

Consider an input fragment "CGTA":

1. **Segment**: Break down "CGTA" into "CGT", "GTA".
2. **Search**: Find nodes corresponding to "CGT" and "GTA" in the graph.
3. **Extend**: Extend matches by exploring neighboring nodes to form the complete sequence.

In this way, PromiseGrid effectively handles sequence fragments with unknown leading characters, ensuring robust sequence matching capabilities.

### Sequence Matching Graph Documentation

#### 1. Overview
A sequence matching graph is a data structure designed to match input sequences to stored sequences by traversing nodes and edges. Each node and edge in the sequence-matching graph has specific compositions that facilitate efficient sequence matching.

#### 2. Node Composition
Each node in the sequence-matching graph contains the following components:

1. **Label**:
    - A string representing the segment of the sequence at this node.
2. **Occurrences**:
    - A list of positions in the reference sequences where this segment occurs.
3. **Children**:
    - A map from sequence characters to child nodes, representing the continuation of the sequence.
        
    ```go
    type Node struct {
        Label      string
        Occurrences []int
        Children   map[rune]*Node
    }
    ```

#### 3. Edge Composition
Each edge in the sequence-matching graph connects to a child node and represents the next segment in the sequence.

1. **Character**:
    - The character from the parent node's label to this edge.
2. **ChildNode**:
    - Reference to the child node.

    ```go
    type Edge struct {
        Character  rune
        ChildNode *Node
    }
    ```

#### 4. Example Structure

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

#### 5. Input Sequence Matching
To match an input sequence to stored sequences in a sequence-matching graph, follow these steps:

- Starting from the root, traverse through the graph using the characters of the input sequence.
- At each node, follow the edge that matches the next character of the input sequence.
- If a suitable edge or child node cannot be found at any point, the matching process terminates without a match.

#### 6. Graph Traversal Mechanism
Graph traversal utilizes a recursive or iterative technique to explore nodes and edges.

1. **Recursive Method**:
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

2. **Iterative Method**:
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

### Examples of Matching
1. **Example 1**: Matching "ACGT" in the diagram
    - Starting from the root, traverse edges labeled 'A', 'C', 'G', and 'T'. 
    - Successful match if all edges are followed correctly.

2. **Example 2**: Matching "AGT" in the diagram
    - Starting from the root, traverse edges labeled 'A', 'G', and 'T'.
    - Successful match if all edges are followed correctly.

### Advanced Questions Answered

#### Sequence-Matching Graph and 'Children' Field of a Node
In a sequence-matching graph, does the 'Children' field of a node contain the next node in the sequence? Or does it contain both previous and next nodes?
- The 'Children' field of a node contains only the next node(s) in the sequence. It maps sequence characters to their corresponding child nodes, representing the continuation of the sequence.

#### Genome Storage in Graph
Are entire genomes stored in the graph? Or are only the sequences of interest stored? If the latter, how are the sequences of interest identified? Are they identified by their position in the genome? Or are they identified by some other means?
- Only the sequences of interest are stored in the graph. These sequences are identified based on specific criteria such as the presence of particular motifs or functional relevance. The identification can be based on their positions in the genome or other biological markers.

#### Adding a New Sequence to the Graph
Show an example of how a new sequence is added to the graph.
- To add a new sequence "AGC" to the graph, follow these steps:

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

### Advanced Questions Answered

#### Sequence-Matching Graph and 'Children' Field of a Node
In a sequence-matching graph, the 'Children' field of a node contains only the next node(s) in the sequence. It maps sequence characters to their corresponding child nodes, representing the continuation of the sequence.

#### Genome Storage in Graph
Only the sequences of interest are stored in the graph. These sequences are identified based on specific criteria such as the presence of particular motifs or functional relevance. The identification can be based on their positions in the genome or other biological markers.

### Optimizations and Further Research

#### Sequence Matching with Unknown Leading Characters
To match a sequence fragment with unknown leading characters, the system can:
1. **Segment the Fragment**: Divide the input sequence fragment into smaller overlapping segments.
2. **Search for Segments**: Locate nodes corresponding to the segments within the graph.
3. **Extend Matches**: Extend partial matches by traversing adjacent nodes in the graph till a complete match is found.

### Practical Example with Unknown Leading Characters
Consider an input fragment "CGTA":
1. **Segment**: Break down "CGTA" into "CGT", "GTA".
2. **Search**: Find nodes corresponding to "CGT" and "GTA" in the graph.
3. **Extend**: Extend matches by exploring neighboring nodes to form the complete sequence.

In this way, the system effectively handles sequence fragments with unknown leading characters, ensuring robust sequence matching capabilities.


# Hosting Conventional Operating Systems or Applications as Guests on PromiseGrid

## Introduction

PromiseGrid (PG) is a decentralized computation platform that diverges from traditional Von Neumann architecture by utilizing byte sequence completion for memory management and execution. This document explores how conventional operating systems or applications, designed for Von Neumann architecture, can be hosted as guests on PromiseGrid.

## Conventional Von Neumann Architecture

### Core Components

1. **Memory**: Unified storage for both code and data.
2. **Central Processing Unit (CPU)**: Executes instructions and processes data.
3. **Input/Output (I/O) Devices**: Facilitate external communication and data exchange.
4. **Sequential Execution**: Follows a linear process of fetching, decoding, and executing instructions from memory.

## PromiseGrid Architecture

### Byte Sequence Completion

- **Memory Management**: Based on completing byte sequences rather than linear addressing.
- **Trie Data Structure**: Utilized for efficient storage and retrieval of byte sequences.
- **Decentralized Storage**: Data is distributed across multiple nodes, enhancing fault tolerance and scalability.

## Mapping Von Neumann Programs to PromiseGrid

### Executing Instructions

1. **Instruction Fetch**:
    - In a Von Neumann machine, the CPU fetches instructions from memory.
    - In PromiseGrid, byte sequences representing instructions are fetched from a distributed trie structure.

2. **Instruction Decode**:
    - The CPU decodes instructions to determine the operation.
    - PromiseGrid decodes byte sequences to identify the required completion or action.

3. **Instruction Execution**:
    - The CPU executes the decoded instruction.
    - PromiseGrid completes the byte sequence, performing the associated action, such as data retrieval, computation, or message passing.

4. **Result Storage**:
    - Von Neumann architecture stores results back in memory.
    - PromiseGrid stores results in the decentralized trie structure.

### Memory Management

#### Von Neumann Memory

- **Unified Storage**: Code and data share the same memory space.
- **Linear Addressing**: Memory is addressed sequentially using numeric addresses.
- **Finite Capacity**: Limited by physical storage.

#### PromiseGrid Memory

- **Content-Addressable Storage**: Data and code are accessed based on their content hashes, rather than numeric addresses.
- **Trie-Based Structure**: Represents memory using a trie, allowing flexible and dynamic data access.
- **Decentralized Storage**: Memory is distributed across the network, ensuring scalability and fault tolerance.

### Addressing and Data Retrieval

1. **Addresses as Content Hashes**:
    - Unlike the linear addressing of Von Neumann architecture, PromiseGrid uses content hashes derived from data or code as addresses.
    - Hashing ensures unique and collision-resistant references to data.

2. **Trie Traversal**:
    - Memory and data retrieval in PromiseGrid involve traversing a trie to locate nodes representing specific byte sequences.
    - This dynamic and decentralized approach supports efficient and scalable data access.

### Execution Layer Adaptation

1. **Virtualization**:
    - The PromiseGrid runtime leverages virtual memory management when hosted on a Von Neumann machine.
    - Byte sequences and trie nodes are stored within the host machine's memory and disk.

2. **Disk Storage**:
    - Persistent storage is managed using the Origin Private File System (OPFS) or similar mechanisms.
    - Trie nodes and content-addressable data are persisted as files, ensuring durability.

3. **Instruction Emulation**:
    - The Von Neumann CPU emulates byte sequence completion by interpreting byte sequences as instructions.
    - This involves memory and I/O operations executed natively on the host machine.

### Harmonizing Paradigms

1. **Instruction Set**:
    - Von Neumann architecture relies on fixed instruction sets, whereas PromiseGrid introduces flexible and dynamically defined operations.
    - The runtime adapts instruction handling to various computational tasks seamlessly.

2. **Data Integrity**:
    - PromiseGrid ensures data integrity through content hashes, even when distributed across the host machine's memory and disk.
    - Combining deterministic access of Von Neumann memory with the consistency of content-addressable storage.

## Container Management on PromiseGrid

PromiseGrid (PG) is a platform inspired by decentralized computation principles. This document explores how a container might be hosted as a guest on PromiseGrid. Specifically, we will investigate how container image layers could be stored and loaded as byte sequences, and how the container runtime could update the writable layer to provide persistent storage and container mobility across PromiseGrid nodes.

## Container Image Storage and Loading

### Image Layers as Byte Sequences

Containers are composed of layers, each layer representing a set of filesystem changes. In PromiseGrid, these layers can be stored and managed as byte sequences. Here's how it might work:

1. **Layer Storage**:
    - Each layer is stored as a unique byte sequence. Layers are referenced using content hashes, ensuring data integrity and easy retrieval.
    - Layers are stored in a decentralized manner, distributed across multiple nodes to enhance fault tolerance and scalability.

2. **Layer Loading**:
    - When a container is instantiated, its image layers are fetched from the network using their content hashes.
    - The layers are assembled into a complete filesystem on the local node, preparing the container for execution.

3. **Efficient Data Retrieval**:
    - Content-addressable storage ensures efficient retrieval of layers. Layers previously downloaded and stored locally can be reused, reducing network overhead.

### Example

Let's consider a container with three layers:
- Base Layer: Contains the OS.
- Layer 1: Contains application dependencies.
- Layer 2: Contains the application itself.

In PromiseGrid, these layers would be stored and retrieved as follows:

1. **Storage**:
    - Each layer is converted to a byte sequence.
    - Byte sequences are stored distributedly across different nodes using content hashes.

2. **Retrieval**:
    - When the container is needed, its layers are retrieved using their content hashes.
    - The byte sequences are reassembled to form the container's filesystem.

## Writable Layer Management

### Log-Structured Filesystem for Writable Layer

To provide a writeable layer for containers that enables persistent storage and mobility, a log-structured filesystem approach is a viable option. Here's how it can be implemented in PromiseGrid.

1. **Log-structured Filesystem (LFS)**:
    - Changes to the filesystem (writes) are appended to a log rather than modifying files in place. This approach is efficient for capturing deltas and supports quick snapshots.
    - Logs can be segmented and stored as byte sequences, allowing distributed nodes to manage parts of the writable layer. 

2. **Persistence**:
    - When a container writes data, the changes are captured in the log. This log is periodically committed and distributed to other nodes.
    - This mechanism ensures that the writable layer persists even if the container is moved or the node fails.

3. **Mobility**:
    - Containers can be moved across different nodes. Since the writable layer is managed via logs, the logs can be transferred to the new node, reassembling the filesystem state.
    - Log entries are content-addressed, allowing efficient transfer and reassembly.

### Example Implementation

Imagine a container running a database with a writable layer storing its current state:

1. **Initial Write**:
    - A write operation modifies the database.
    - The change is captured and appended to the log as a byte sequence.

2. **Log Commit**:
    - Periodically, the log is committed and its entries are distributed to other nodes for redundancy.
    - Each log entry is content-addressed, ensuring integrity.

3. **Container Migration**:
    - The container needs to be moved to another node.
    - The log entries are transferred to the target node.
    - At the target node, the log is replayed to reassemble the writable layer, maintaining the state of the database.

## Conclusion

Hosting conventional operating systems or applications as guests on PromiseGrid involves efficient handling of container image layers and managing a writeable layer. By storing image layers as byte sequences and using a log-structured filesystem for the writable layer, PromiseGrid can ensure persistent storage, efficient data retrieval, and container mobility. This decentralized approach leverages the strengths of PromiseGrid to provide robust and scalable container management.


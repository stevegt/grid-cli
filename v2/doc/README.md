## PromiseGrid grid-cli v2 Documentation

### Table of Contents

1. [Introduction](#introduction)
2. [Design Notes](#design-notes)
3. [Module Integration](#module-integration)
4. [Capability Tokens](#capability-tokens)
5. [Data Flow and Sequence Matching](#data-flow-and-sequence-matching)
6. [Inter-Process Communication (IPC)](#inter-process-communication-ipc)
7. [Sequence Graphs](#sequence-graphs)
8. [Hybrid DHT and Open Market](#hybrid-dht-and-open-market)
9. [Persisting the Trie to Disk](#persisting-the-trie-to-disk)
10. [Advanced Configuration](#advanced-configuration)

### Design Notes

Includes detailed notes on the design philosophy and architectural principles guiding PromiseGrid. 

- [Design Overview](./00.md)
- [Discussion Notes](./01.md)

### Module Integration

Discusses strategies for effective module integration, ensuring modular, flexible, and decentralized processing.

- [Cache and Module Handling](./02.md)
- [Advanced Discussions](./03.md)
- [Module Integration Concepts](./04.md)

### Capability Tokens

Details the use and implementation of self-contained capability tokens in PromiseGrid. Tokens manage permissions and identity securely.

- [Capability Tokens and Prior Art](./150-capability-tokens.md)

### Data Flow and Sequence Matching

Explores the techniques of sequence matching, data structures like tries and graphs, and their applications.

- [Introduction to Sequence Matching](./200-sequence.md)
- [Graph or Node-Based Representation](./210-sequence-graph.md)

### Inter-Process Communication (IPC)

Inter-Process Communication is critical for a microkernel architecture, facilitating communication between various modules or components within the PromiseGrid.

- [IPC in PromiseGrid](./320-ipc.md)

### Sequence Graphs

Details on implementing sequence matchers using graph-based structures within PromiseGrid.

- [Sequence Matching Implementation](./220-graph.md)
- [Advanced Sequence Matching Concepts](./221-graph.md)

### Hybrid DHT and Open Market

Combines DHT with an open market system to leverage decentralized storage and lookup efficiency.

- [Hybrid DHT and Market Design](./343-dht.md)
- [Pure Market Design](./344-market.md)

### Persisting the Trie to Disk

Describes various options for persisting the trie data structure to disk within PromiseGrid.

- [Decentralized Trie](./345-dtrie.md)
- [Persistence Strategies](./346-persist.md)

### Advanced Configuration

Additional configuration details for specialized use cases in PromiseGrid.

- [Technical Overview of GPUs and Their Architecture](./205-gpu.md)
- [FPGA Architecture and Prefix Matching](./206-fpga.md)
- [Implementing Bloom Filters](./350-bloom.md)

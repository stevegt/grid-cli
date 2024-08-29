# Fundamental Axioms in PromiseGrid

This document lists the fundamental axioms that form the basis of the PromiseGrid architecture and operation, extracted from various design and discussion documents.

## Core Concepts and Principles

1. **Decentralized Architecture**: PromiseGrid operates as a decentralized computing, communications, and governance system, designed to be owned and operated by its users rather than any single entity.
2. **Capability-as-Promise Model**: Capabilities are treated as promises. A capability token represents a promise that can either be fulfilled or revoked.
3. **Content-Addressable Code**: Code and data are addressed by their content, not by location or name. This allows the grid to store and execute code and access data from any node in the network.
4. **Promises All the Way Down**: Every interaction in the system is based on promises, and a response to a promise is another promise.
5. **Non-Sandboxed Modules**: Non-sandboxed modules handle specific external operations in the grid (e.g., network communications, file access), analogous to device drivers in a microkernel OS.

## Cache and Promise Handling

1. **Directory Tree Structure**: The cache **SHOULD BE** a directory tree rather than an in-memory map.
2. **Persistent Cache Keys**: Cache keys **MUST** use filesystem separators (`/`) and **MUST** be URL-encoded.
3. **Multiple Caches**: There **MAY BE** multiple caches, including built-in kernel cache and module-provided caches.
4. **Built-in Cache Loading**: The kernel **SHOULD** load the built-in cache from embedded resources using Go’s `embed` feature.
5. **Filesystem Abstracting**: The kernel **MAY** use OPFS for disk file access and **MAY** utilize the `afero` library to abstract filesystem interactions.
6. **Modules as Caches**: The kernel **MUST** treat modules as caches, consulting them in case of a cache miss.
7. **Unified Interface**: There **SHALL BE** no difference between a cache lookup and a function call.

### Treating Acceptance as Promises

1. **Acceptance Function**: Modules **MUST** define acceptance criteria for promises, module hashes, and arguments into a single `Accept()` function.
2. **Promises as Acceptance Response**: The `Accept` function returns a promise message.
3. **Commitment**: Acceptance is treated as a commitment to handle the message correctly. Failure is considered a broken promise.
4. **Ant Routing Mechanism**: The syscall tree acts as an "ant routing" mechanism, caching successful paths.

## Integration with Computational Theory

1. **Acceptance in Computational Theory**: The term "accept" aligns with its use in computing theory (Church, Turing, Chomsky).
2. **Recognizers Concept**: Modules act as recognizers for specific tasks, based on promise hash, module hash, and arguments.
3. **Promises Integration**: Acceptance as a promise message, ensuring robustness and trustworthiness. Each layer (modules, syscall tree, kernel) fulfills promises based on promises made by layers below.

## Kernel's Dynamic Syscall Tree

1. **Syscall Tree Structure**: The syscall tree **MUST** use hierarchical keys and store acceptance history for all modules.
2. **Acceptance History Storage**: The syscall tree **SHOULD BE** the same as acceptanceHist and **SHOULD** store positive and negative acceptance history.
3. **Routing Optimization**: The kernel **MUST** route messages to the module whose syscall tree key matches the most leading parameter components.

## Explicit Module Registration vs Hash-Based Routing

1. **Explicit Registration**:
    - **Clarity and Explicitness**
    - **Fine-Grained Control**
    - **Dynamic Adaptation**
    - **Complexity and Overhead**
    - **Dependency Management**
2. **Hash-Based Routing**:
    - **Simplicity and Efficiency**
    - **Decentralized Management**
    - **Opaque Mapping and Flexibility Limits**
    - **Security Risks**

## Handling and Execution Efficiency

1. **Combining Accept and HandleMessage**: Combining decision-making and handling in a single function can simplify logic and improve consistency and efficiency.
2. **Separate Functions for Clarity**: Maintaining separate `Accept()` and `HandleMessage()` functions for clarity and modular logic but ensuring early rejection.

## Error Handling and Broken Promises

1. **Logging and Handling**: Broken promises **MUST** be logged and managed to ensure accountability and allow dynamic rerouting of requests.
2. **Reputation Impact**: Reliability and performance affect the reputation of modules, influencing future message routing and handling.

## Unified Message Structure

1. **Message Structure**: The message structure includes the promise as the first element in the `Parms` field. 
2. **Routing Based on Promises**: Recipients route or discard messages based on the leading promise.

### Multihash, Multibase, and Multicodec

1. **Multihash**: Ensures flexibility by specifying multiple hash algorithms.
2. **Multibase**: Encodes hashes for automatic format interpretation, enhancing compatibility and extensibility.

## Conclusion

By integrating these fundamental axioms, PromiseGrid ensures a robust, flexible, and efficient system for decentralized governance and cooperation. This axiom-driven approach underpins the trust, scalability, and adaptability of PromiseGrid's architecture and operational mechanisms.

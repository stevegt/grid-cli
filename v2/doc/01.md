# Design and Discussion Notes from All Files

## Overview

This document consolidates all design and discussion notes from various files to provide a comprehensive view of the architectural principles and strategic considerations guiding the development of PromiseGrid. This includes integration of the fundamental principles outlined in previous documents and merging in the contents of `141-cache.md`.

## Core Concepts

### Decentralized Architecture

PromiseGrid operates as a decentralized computing, communications, and governance system. It is designed to be owned and operated by its users rather than any single entity, ensuring distributed control and ownership.

### Capability-as-Promise Model

Capabilities are treated as promises, similar to concepts from Promise Theory. A capability token represents a promise that can either be fulfilled or revoked, establishing trust and accountability within the system.

### Content-Addressable Code

Both code and data are addressed by their content, not by location or name. This design enables the grid to store, execute code, and access data from any node in the network efficiently, leveraging the integrity ensured by content-addressable storage.

### Promises All the Way Down

Every interaction in the system is based on promises. A response to a promise is another promise, creating a network where all entities are interlinked through a chain of commitments, fostering reliability and traceability.

### Non-Sandboxed Modules

Non-sandboxed modules in PromiseGrid are comparable to device drivers in a microkernel OS. These modules handle specific external operations such as network communications and file access. The kernel delegates these operations to non-sandboxed modules while maintaining overall control.

Non-sandboxed modules SHOULD be loaded from the embedded cache to
improve local security.

## Cache and Promise Handling in the PromiseGrid Kernel

### Cache Structures


- There **MAY BE** multiple caches, including the built-in cache in the kernel and caches provided by various modules.
- The kernel **SHOULD** load the built-in cache from embedded resources using Go’s `embed` feature.
- The kernel **MAY** use the Origin Private File System (OPFS) for disk file access and **MAY** utilize the `afero` library to abstract filesystem interactions.

### Treating Modules as Caches

#### Role of Modules

- The kernel **MUST** treat modules as caches. In the event of a cache miss, the kernel **MUST** consult one or more modules to retrieve the requested value.
- Modules **MAY** contribute to the cache or provide the requested data dynamically.

#### Unified Interface

- From the caller's perspective, there **SHALL BE** no difference between a cache lookup and a function call. Both operations **SHALL BE** treated as byte sequence completion operations.  The caller sends a message consisting of a byte sequence, and receives a response message containing the remainder of the sequence.

### Acceptance and Promises

#### Acceptance Criteria

- Modules **MUST** define acceptance criteria for promises, module hashes, and arguments. This can be implemented as an interface for modules, simplifying acceptance to a single `Accept()` function.

#### Promises as Acceptance

- By returning a promise message from `Accept` instead of a boolean, modules provide additional guarantees and meta-information. The acceptance is treated as a commitment to handle the message correctly. If `HandleMessage` fails, it is considered a broken promise.

#### Ant Routing Mechanism

- The syscall tree acts as an "ant routing" mechanism, caching successful paths to optimize future routing. In future similar calls, the same path will be followed to the same module.
- The syscall tree **MUST** use hierarchical keys and **SHOULD** filter based on whether any module accepts the leading parameters, matches the module hash, and accepts the arguments.

## Integration with Church, Turing, and Chomsky's Concept of "Accept"

1. **Computational Theory**:
    - The term "accept" aligns with Church, Turing, and Chomsky's use in computing theory and languages, where an automaton or machine accepts an input if it reaches an accepting state.
    - In PromiseGrid, modules act as recognizers for specific tasks based on the promise hash, module hash, and arguments.

2. **Promises All the Way Down**:
    - The concept of "promises all the way down" integrates acceptance as a promise message, enhancing robustness and trustworthiness.
    - Each layer (modules, syscall tree, kernel) makes and fulfills promises based on the promises made by the layers below it.

## Kernel's Dynamic Syscall Tree

### Syscall Table and Acceptance History

- The acceptanceHist and syscallTable **SHOULD BE** the same. This dynamic syscall tree **SHOULD** store acceptance history for all modules.
- The kernel **SHOULD** store positive and negative acceptance history for all modules to look up which modules accept the leading bytes, skip modules that reject those bytes, and call accepting modules with the full sequence of bytes.
- This table **SHOULD** start empty and be populated during operation as the kernel consults built-in and other modules to handle messages.
- The kernel **MUST** route messages to the module whose syscall tree key matches the most leading parameter components, optimizing routing and reducing redundant checks.

## Flexible Design for Module Registration

### Explicit Module Registration

**Pros**
1. **Clarity and Explicitness**: Modules explicitly report their capabilities, making it easier to understand the system's configuration and functionality.
2. **Fine-Grained Control**: The kernel can enforce specific rules and constraints on modules based on their declared capabilities.
3. **Dynamic Adaptation**: Modules can dynamically update their capabilities, allowing for on-the-fly changes and adaptation.

**Cons**
1. **Complexity**: The registration process adds complexity to the module initialization and management process.
2. **Performance Overhead**: The kernel must maintain and query a registry of module capabilities, which can introduce performance overhead.
3. **Dependency Management**: Changes in module capabilities may require updates to the kernel or other modules, increasing the risk of dependencies and compatibility issues.

### Hash-Based Module Routing

**Pros**
1. **Simplicity**: The kernel routes messages based on cryptographic hashes, reducing the need for an explicit registration step.
2. **Efficiency**: Hash-based routing can be highly efficient, leveraging cryptographic properties to ensure unique and consistent module addressing.
3. **Decentralized Management**: Modules are self-contained and can be managed independently without requiring kernel updates or reconfiguration.

**Cons**
1. **Opaque Mapping**: It may be less clear which module handles a specific message, as the mapping relies on hashes rather than explicit declarations.
2. **Limited Flexibility**: Modules cannot dynamically update their capabilities without changing their hash, reducing adaptability.
3. **Security Risks**: The kernel must ensure that the hash-based routing mechanism is secure against hash collisions and attacks.

## Modular Interaction and Delegation

Modules can unwrap nested messages and forward them to other modules, allowing for complex interactions and message flows. The kernel is not necessarily the final arbiter of which module handles a message.

**Implications**
1. **Delegated Control**: Modules can delegate tasks to other modules, enabling a modular and extensible system design.
2. **Dynamic Message Flows**: Nested messages allow for dynamic and context-specific message handling, improving flexibility.

### Acceptance and Handling Functions

#### Combined Function Approach

**Pros**:
1. **Simplicity**: Combines decision-making and handling into a single function.
2. **Consistency**: Ensures the decision to handle and the actual handling are tightly coupled, increasing consistency and reducing logic duplication.
3. **Efficiency**: Eliminates redundant checks by combining acceptance and handling.

#### Separate Function Approach

**Pros**:
1. **Clarity**: Maintains a clear separation between decision-making and handling logic.
2. **Early Rejection**: Allows for quick rejection of messages based on promise, module hash, or arguments without performing any handling.
3. **Modular Logic**: Facilitates modular and specialized acceptance and handling logic.

### Conclusion

Combining promises at every level and implementing a hierarchical syscall tree with caching and acceptance history ensures trust, accountability, and efficient message handling in a decentralized governance framework. The consistent handling of cache and modules contributes to a robust and flexible system, aligning with computational theory and modular architecture principles.

# Cache

## Overview

This document consolidates the design and discussion notes from various files to provide a comprehensive view of the architectural principles and strategic considerations for PromiseGrid.

## Core Concepts

### Decentralized Architecture

PromiseGrid operates as a decentralized computing, communications, and governance system owned and operated by its users, ensuring distributed control.

### Capability-as-Promise Model

Capabilities are treated as promises. A capability token represents a promise that can either be fulfilled or revoked, establishing trust and accountability within the system.

### Content-Addressable Code

Both code and data are addressed by their content. This design enables efficient storage, execution, and data access from any node in the network, leveraging the integrity ensured by content-addressable storage.

### Promises All the Way Down

Every interaction in the system is based on promises. A response to a promise is another promise, creating a network where all entities are interlinked through a chain of commitments, fostering reliability and traceability.

### Non-Sandboxed Modules

Non-sandboxed modules handle specific external operations such as network communications and file access. The kernel delegates these operations to non-sandboxed modules while maintaining overall control.

## Cache and Promise Handling in the PromiseGrid Kernel

### Cache Structures

- Multiple caches, including the built-in cache in the kernel and caches provided by various modules.
- The kernel should load the built-in cache from embedded resources using Go’s `embed` feature.
- The kernel may use the Origin Private File System (OPFS) and the `afero` library for filesystem interactions.

### Treating Modules as Caches

- The kernel must treat modules as caches, consulting them in case of cache misses.
- Modules may contribute to the cache or provide requested data dynamically.
- From the caller's perspective, there should be no difference between a cache lookup and a function call: both are byte sequence completion operations.

### Acceptance and Promises

- Modules must define acceptance criteria for promises, module hashes, and arguments using an `Accept()` function.
- Returning a promise message from `Accept` provides additional guarantees and meta-information.
- A syscall tree acts as an "ant routing" mechanism, optimizing routing by caching successful paths. It uses hierarchical keys.

## Integration with Church, Turing, and Chomsky's Concept of "Accept"

- The term "accept" aligns with Church, Turing, and Chomsky's use in computing theory, where an automaton or machine accepts an input if it reaches an accepting state.
- The concept of "promises all the way down" integrates acceptance as a promise message, enhancing robustness and trustworthiness.

## Kernel's Dynamic Syscall Tree

- The syscall table should store acceptance history for all modules, including both positive and negative acceptance, to optimize routing.
- Messages are routed to the module whose syscall tree key matches the most leading parameter components.

## Flexible Design for Module Registration

### Explicit Module Registration

**Pros:**
1. Clarity and explicitness.
2. Fine-grained control
3. Dynamic adaptation.

**Cons:**
1. Complexity.
2. Performance overhead.
3. Dependency management.

### Hash-Based Module Routing

**Pros:**
1. Simplicity.
2. Efficiency.
3. Decentralized management.

**Cons:**
1. Opaque mapping.
2. Limited flexibility.
3. Security risks.

## Modular Interaction and Delegation

Modules can unwrap nested messages and forward them to other modules, enabling complex interactions and message flows.

### Acceptance and Handling Functions

#### Combined Function Approach

**Pros:**
1. Simplicity.
2. Consistency.
3. Efficiency.

#### Separate Function Approach

**Pros:**
1. Clarity.
2. Early rejection.
3. Modular logic.


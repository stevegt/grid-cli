# PromiseGrid Design Overview

## Introduction

PromiseGrid is designed as a decentralized computing, communications, and governance framework. The design emphasizes promise-based interactions and content-addressable storage.

## Core Concepts

1. **Decentralized Architecture**: Owned and operated by users, ensuring distributed control and scalability.
2. **Promise-Based Interaction**: Interactions are based on making, fulfilling, or managing promises.
3. **Content-Addressable Code and Data**: Code and data are addressed by their content, not by location or name.
4. **Promises All the Way Down**: Every response to a promise is another promise, promoting reliability and traceability.
5. **Non-Sandboxed Modules**: Modules function similarly to device drivers, handling specific external operations.

## System Architecture

### Cache Handling

#### Cache Structures

1. **Multiple Caches**:
   - Multiple caches, including a built-in cache and module-provided caches.
   - Embedded resources loaded using Go's `embed` feature.
   - Possibility of OPFS for disk access and `afero` for filesystem abstraction.

#### Modules as Caches

1. **Kernel and Modules**:
   - Modules are treated as caches. Kernel consults modules on cache miss.
   - Unified interface for cache lookup and function call, treated as hashed function calls.

### Acceptance and Promises

#### Acceptance Criteria

- Modules must define acceptance criteria via an `Accept()` function.
- Acceptance is treated as a commitment; handling failure is a broken promise.
- "Ant routing" mechanism: syscall tree caches successful paths to optimize routing.

### Integration with Computational Theory

- **Automaton Concept**: Modules accept input similar to an automaton reaching an accepting state.
- **Promises All the Way Down**: Each layer makes and fulfills promises based on underlying promises.

### Flexible Module Registration

#### Explicit Registration

**Pros**:
1. Clarity and control.
2. Dynamic capability updates.

**Cons**:
1. Complexity and performance overhead.
2. Dependency risks.

#### Hash-Based Routing

**Pros**:
1. Simplicity and decentralization.
2. Efficiency.

**Cons**:
1. Opaque mapping.
2. Limited dynamic update capability.
3. Security considerations.

### Exchange Rate Routing

#### Concept

Hosts route messages based on exchange rates of personal currencies, incentivizing reliability and decentralized control.

#### Advantages

- Encourages reliability.
- Uses market dynamics for control.
- Self-regulating.

### Dynamic Syscall Tree

- Utilizes a dynamic syscall table with acceptance history.
- Messages routed based on leading parameter component match.

### Combining Accept and HandleMessage Functions

#### Pros and Cons

**Combined Function**:
- Simplicity and consistency.
- Efficiency without redundant checks.

**Separate Functions**:
- Clear separation and modular logic.
- Early rejection capability.

### Dynamic Syscall Table

- Hierarchical syscall tree with caching.
- Populated during operation to optimize message handling.

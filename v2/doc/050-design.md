# PromiseGrid Design Overview

## Introduction

The PromiseGrid Kernel is designed as a decentralized computing, communications, and governance framework. It leverages promise-based interactions and content-addressable storage to create a robust, modular, and secure system. The core principles emphasize decentralization, modularity, and the use of promises at every level of the system architecture.

## Core Concepts

1. **Decentralized Architecture**: PromiseGrid operates as a decentralized system owned and operated by its users, eliminating reliance on any single entity. This design ensures distributed control and ownership, enhancing resilience and scalability.

2. **Promise-Based Interaction**: The system is founded on the concept of promises, akin to those in Promise Theory. Every interaction is based on making, fulfilling, or managing promises, ensuring accountability and trust in operations.

3. **Content-Addressable Code and Data**: Both code and data are addressed by their content, not by their location or name. This content-addressable storage allows the grid to store and execute code and access data from any node in the network efficiently.

4. **Promises All the Way Down**: Every response to a promise is another promise, embedding the concept deeply into the system's operations. This approach fosters a network where components are interlinked through a chain of commitments, promoting reliability and traceability.

5. **Non-Sandboxed Modules**: Non-sandboxed modules function similarly to device drivers in a microkernel OS. They handle specific external operations such as network communications and file access, with the kernel delegating these operations to maintain control over execution.

## System Architecture

### Cache Handling in the PromiseGrid Kernel

#### Cache Structures

1. **Multiple Caches**:
   - There **MAY BE** multiple caches, including a built-in cache within the kernel and caches provided by various modules.
   - The kernel **SHOULD** load the built-in cache from embedded resources using Go's `embed` feature.
   - The kernel **MAY** use the Origin Private File System (OPFS) for disk file access and **MAY** utilize the `afero` library to abstract filesystem interactions.

#### Treating Modules as Caches

1. **Role of Modules**:
   - The kernel **MUST** treat modules as caches. In the event of a cache miss, the kernel **MUST** consult one or more modules to retrieve the requested value.
   - Modules **MAY** contribute to the cache or provide the requested data dynamically.

2. **Unified Interface**:
   - From the caller's perspective, there **SHALL BE** no difference between a cache lookup and a function call. Both operations are treated as hashed function calls, where the caller sends a message with a leading hash and any arguments and receives a message containing the requested content.

### Acceptance and Promises

#### Acceptance Criteria

- Modules **MUST** define acceptance criteria for promises, module hashes, and arguments, implemented as an `Accept()` function.
- By returning a promise message from `Accept` instead of a boolean, modules provide additional guarantees and meta-information.
- The acceptance is treated as a commitment to handle the message correctly. If `HandleMessage` fails, it's considered a broken promise.

#### Ant Routing Mechanism

- The syscall tree acts as an "ant routing" mechanism, caching successful paths to optimize future routing.
- The syscall tree **MUST** use hierarchical keys and **SHOULD** filter based on whether any module accepts the leading parameters, matches the module hash, and accepts the arguments.
- The kernel **SHOULD** store positive and negative acceptance history, starting empty and populated during operation.

### Integration with Computational Theory

#### Aligning with Church, Turing, and Chomsky's "Accept" Concept

- The term "accept" aligns with computational theory, where an automaton accepts an input if it reaches an accepting state.
- In PromiseGrid, modules act as recognizers for specific tasks based on the promise hash, module hash, and arguments.

#### Promises All the Way Down

- The concept of "promises all the way down" integrates acceptance as a promise message, enhancing robustness and trustworthiness.
- Each layer (modules, syscall tree, kernel) makes and fulfills promises based on the promises made by the layers below it.

### Flexible Design for Module Registration

#### Explicit Module Registration

**Pros**:

1. **Clarity and Explicitness**: Modules explicitly report their capabilities, simplifying system understanding and management.
2. **Fine-Grained Control**: The kernel can enforce specific rules and constraints on modules based on their declared capabilities.
3. **Dynamic Adaptation**: Modules can dynamically update their capabilities, supporting on-the-fly changes.

**Cons**:

1. **Complexity**: Adds complexity to module initialization and management.
2. **Performance Overhead**: Maintaining and querying a registry introduces overhead.
3. **Dependency Management**: Changes in capabilities may necessitate updates across the system, increasing dependency risks.

#### Hash-Based Module Routing

**Pros**:

1. **Simplicity**: The kernel routes messages based on cryptographic hashes, reducing the need for explicit registration.
2. **Efficiency**: Hash-based routing leverages cryptographic properties for unique and consistent module addressing.
3. **Decentralized Management**: Modules are self-contained and managed independently.

**Cons**:

1. **Opaque Mapping**: Mapping relies on hashes rather than explicit declarations, which may reduce clarity.
2. **Limited Flexibility**: Modules cannot dynamically update capabilities without changing their hash.
3. **Security Risks**: Ensuring hash-based routing is secure against collisions and attacks is essential.

### Exchange Rate Routing

#### Concept

Exchange rate routing is an innovative mechanism where hosts within the grid route messages based on the exchange rates of personal currencies. Each host acts like its own "central bank," issuing a form of currency represented by reputation points. These reputation points influence routing decisions, promoting reliable behavior and network stability.

#### How It Works

1. **Currency Evaluation**: Each host evaluates the currency (reputation) of other hosts.
2. **Routing Decisions**: Hosts prefer routes through hosts with higher-valued currencies to ensure stability and reliability.
3. **Dynamic Adjustments**: Exchange rates are dynamic, adjusting based on host behavior and market conditions.

#### Advantages

- **Encourages Reliability**: Hosts are incentivized to maintain strong reputations.
- **Economic Incentives**: The system uses market dynamics, encouraging decentralized control.
- **Self-Regulating**: Misbehaving hosts receive less traffic, reducing their ability to affect the network adversely.

### Kernel's Dynamic Syscall Tree

#### Syscall Table and Acceptance History

- The kernel uses a dynamic syscall table, storing positive and negative acceptance history for all modules.
- Messages are routed to the module whose syscall tree key matches the most leading parameter components.

### Combining Accept and HandleMessage Functions

#### Pros and Cons

**Combined Function**:

- **Pros**:
  - Simplicity: Combines decision-making and handling.
  - Consistency: Tight coupling increases consistency and reduces logic duplication.
  - Efficiency: Eliminates redundant checks.

**Separate Functions**:

- **Pros**:
  - Clarity: Maintains separation between decision-making and handling logic.
  - Early Rejection: Allows quick rejection without handling.
  - Modular Logic: Facilitates specialized acceptance and handling logic.

### Proposed Dynamic Syscall Table and Ant Routing

- Implementing a hierarchical syscall tree with caching and acceptance history enhances efficient message handling.
- The syscall tree starts empty and is populated as modules accept and handle messages.
- Successful paths are cached, optimizing future routing.

## Conclusion

By integrating promises at every level and implementing a hierarchical syscall tree with caching and acceptance history, PromiseGrid ensures trust, accountability, and efficient message handling in a decentralized governance framework. The consistent handling of cache and modules contributes to a robust and flexible system, aligning with computational theory and modular architecture principles.


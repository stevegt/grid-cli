# PromiseGrid Design

## Overview

PromiseGrid is designed as a decentralized computing, communications, and governance framework, leveraging promise-based interactions and content-addressable storage to create a robust and modular system.

## Core Concepts

1. **Decentralized Architecture**: PromiseGrid operates as a decentralized computing, communications, and governance system. It is designed to be owned and operated by its users rather than any single entity.

2. **Capability-as-Promise Model**: Capabilities are treated as promises, similar to the concepts from Promise Theory. A capability token represents a promise that can either be fulfilled or revoked.

3. **Content-addressable Code**: Both code and data are addressed by their content, not by location or name. This allows the grid to store and execute code and access data from any node in the network.

4. **Promises All the Way Down**: Every interaction in the system is based on promises. A response to a promise is another promise.

5. **Non-Sandboxed Modules**: Non-sandboxed modules in PromiseGrid are analogous to device drivers in a microkernel OS. Just as device drivers handle specific hardware functionality in a microkernel, non-sandboxed modules handle specific external operations in the grid (e.g., network communications, file access). The kernel delegates these operations to non-sandboxed modules while maintaining control over the overall execution.

## System Architecture

### Flexible Design for Module Registration

Modules in PromiseGrid interface with the kernel via either explicit registration or through hash-based routing mechanisms. These two primary methods offer distinct advantages and challenges:

#### Explicit Module Registration

**Pros**:
1. **Clarity and Explicitness**: Modules report their capabilities, simplifying system understanding and management.
2. **Fine-Grained Control**: The kernel can enforce rules and constraints based on module capabilities.
3. **Dynamic Adaptation**: Modules can update their capabilities dynamically, supporting on-the-fly changes.

**Cons**:
1. **Complexity**: Adds complexity to module initialization and management.
2. **Performance Overhead**: Maintaining and querying a registry introduces overhead.
3. **Dependency Management**: Changes in capabilities may necessitate updates across the system, increasing dependency risks.

#### Hash-Based Routing

**Pros**:
1. **Simplicity**: The kernel routes messages based on cryptographic hashes, reducing the need for an explicit registration step.
2. **Efficiency**: Hash-based routing can be highly efficient, leveraging cryptographic properties to ensure unique and consistent module addressing.
3. **Decentralized Management**: Modules are self-contained and can be managed independently without requiring kernel updates or reconfiguration.

**Cons**:
1. **Opaque Mapping**: It may be less clear which module handles a specific message, as the mapping relies on hashes rather than explicit declarations.
2. **Limited Flexibility**: Modules cannot dynamically update their capabilities without changing their hash, reducing adaptability.
3. **Security Risks**: The kernel must ensure that the hash-based routing mechanism is secure against hash collisions and attacks.

### Exchange Rate Routing

Exchange rate routing is an innovative mechanism where hosts within the grid route messages based on the exchange rates of personal currencies. Each host acts like its own "central bank," issuing a form of currency represented by reputation points. These reputation points influence the routing decisions, promoting reliable behavior and network stability.

#### How It Works

1. **Currency Evaluation**: Each host evaluates the currency (reputation) of other hosts.
2. **Routing Decisions**: Hosts prefer routes through other hosts with higher-valued currencies to ensure stability and reliability.
3. **Dynamic Adjustments**: Exchange rates are dynamic and adjust based on host behavior and market conditions.

#### Advantages

- **Encourages Reliability**: Hosts are incentivized to maintain strong reputations.
- **Economic Incentives**: The system uses market dynamics, encouraging decentralized control.
- **Self-Regulating**: Misbehaving hosts receive less traffic, reducing their ability to affect the network adversely.

### Nested Messages and Kernel Arbitration

Modules can unwrap nested messages, allowing dynamic and complex interactions. The kernel's arbitration is not final, supporting a modular design where tasks can be delegated further through nested messaging.

#### Implications
1. **Delegated Control**: Enables modular and extensible system design by delegating tasks.
2. **Dynamic Message Flows**: Supports flexible message handling based on context.
3. **Complex Dependency Graphs**: Requires careful management to maintain stability.

### Cache (Syscall Tree) Node Structure - The Ant-Routing Mechanism



The syscall tree acts as an "ant routing" mechanism, where successful paths from previous calls are cached, optimizing future routing. The syscall tree utilizes hierarchical keys, filtering based on module acceptance of leading parameters and arguments. 

### Messages and Promises

1. **Message Structure**: 
    - Messages are treated as byte sequences, which have no inherent meaning to the kernel and are interpreted by the appropriate modules.

2. **Multihash, Multibase, and Multicodec**: 
    - PromiseGrid leverages these standards to autodetect formats, enhancing compatibility and extensibility across various data formats.

### Conclusion

PromiseGrid leverages decentralized architecture, promise-based interactions, and flexible module registration to create a resilient, scalable, and modular system. By integrating these elements, PromiseGrid ensures efficient message handling, dynamic adaptation, and robust promise fulfillment across the network.

### Open Questions

- How can byte sequence completion be reconciled with explicit registration and hash-based routing?
- What mechanisms can be developed to handle broken promises effectively?
- How can the determinant factors of routing efficiency be further optimized with multiple handlers?

### Suggestions for Document Improvement

- Include more examples and case studies.
- Expand on the hierarchical syscall tree descriptions.
- Update sections with new questions arising during implementation.
- Refine the glossary with specialized terms and concepts.
- Incorporate visual aids to showcase system architecture and data flow.

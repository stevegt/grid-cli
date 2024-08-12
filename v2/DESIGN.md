# PromiseGrid Design

## Overview

This document outlines the design considerations and architecture for implementing PromiseGrid.

## Core Concepts

1. **Decentralized Architecture**: PromiseGrid operates as a decentralized computing, communications, and governance system. It is designed to be owned and operated by its users rather than any single entity.

2. **Capability-as-Promise Model**: Capabilities are treated as promises, similar to the concepts from Promise Theory. A capability token represents a promise that can either be fulfilled or revoked.

3. **Content-addressable Code**: Both code and data are addressed by their content, not by location or name. This allows the grid to store and execute code and access data from any node in the network.

4. **Promises All the Way Down**: Every interaction in the system is based on promises. A response to a promise is another promise.

5. **Non-Sandboxed Modules**: Non-sandboxed modules in PromiseGrid are analogous to device drivers in a microkernel OS. Just as device drivers handle specific hardware functionality in a microkernel, non-sandboxed modules handle specific external operations in the grid (e.g., network communications, file access). The kernel delegates these operations to non-sandboxed modules while maintaining control over the overall execution.

## Glossary of Terms and Concepts

- **Decentralized Architecture**: A system design where control is distributed among various actors or nodes, rather than centralized in a single entity.
- **Capability-as-Promise Model**: A model where capabilities represent promises that can either be fulfilled or revoked.
- **Content-addressable Code**: Code and data addressed based on their content, allowing decentralized storage and execution.
- **Promises All the Way Down**: A paradigm where every interaction in the system results in a promise, with responses being new promises.
- **Non-Sandboxed Modules**: Modules granted more access and responsibility, analogous to device drivers, for handling specific external operations.

## Flexible Design for Module Registration

### Overview

In PromiseGrid, modules can interface with the kernel through two primary mechanisms:
1. Modules explicitly register with the kernel, reporting what messages they can handle.
2. The kernel hashes each module and uses that hash to route messages to the appropriate module.

Both approaches have their pros and cons, and the chosen method impacts the system's flexibility, security, and maintainability.

### Explicit Module Registration

#### Pros
1. **Clarity and Explicitness**: Modules explicitly report their capabilities, making it easier to understand the system's configuration and functionality.
2. **Fine-Grained Control**: The kernel can enforce specific rules and constraints on modules based on their declared capabilities.
3. **Dynamic Adaptation**: Modules can dynamically update their capabilities, allowing for on-the-fly changes and adaptation.

#### Cons
1. **Complexity**: The registration process adds complexity to the module initialization and management process.
2. **Performance Overhead**: The kernel must maintain and query a registry of module capabilities, which can introduce performance overhead.
3. **Dependency Management**: Changes in module capabilities may require updates to the kernel or other modules, increasing the risk of dependencies and compatibility issues.

### Hash-Based Module Routing

#### Pros
1. **Simplicity**: The kernel routes messages based on cryptographic hashes, reducing the need for an explicit registration step.
2. **Efficiency**: Hash-based routing can be highly efficient, leveraging cryptographic properties to ensure unique and consistent module addressing.
3. **Decentralized Management**: Modules are self-contained and can be managed independently without requiring kernel updates or reconfiguration.

#### Cons
1. **Opaque Mapping**: It may be less clear which module handles a specific message, as the mapping relies on hashes rather than explicit declarations.
2. **Limited Flexibility**: Modules cannot dynamically update their capabilities without changing their hash, reducing adaptability.
3. **Security Risks**: The kernel must ensure that the hash-based routing mechanism is secure against hash collisions and attacks.

### Combining Both Approaches

Combining both explicit registration and hash-based routing can provide a balanced approach, leveraging the strengths of each method.

1. **Initialization Phase**: During initialization, modules register with the kernel, reporting their capabilities. The kernel maintains a registry of these capabilities.
2. **Hash-Based Execution**: At runtime, the kernel uses hash-based routing to efficiently route messages to the appropriate modules, consulting the registry as needed for additional information.

### Nested Messages and Kernel Arbitration

Modules can unwrap nested messages and forward them to other modules, allowing for complex interactions and message flows. This means that the kernel is not necessarily the final arbiter of which module handles a message.

#### Implications
1. **Delegated Control**: Modules can delegate tasks to other modules, enabling a modular and extensible system design.
2. **Dynamic Message Flows**: Nested messages allow for dynamic and context-specific message handling, improving flexibility.
3. **Complex Dependency Graphs**: The potential for complex interactions and dependencies increases, requiring careful management and monitoring to ensure system stability.

### Conclusion

The choice between explicit module registration and hash-based routing significantly impacts the design and functionality of the PromiseGrid system. Combining both approaches can provide a balanced solution, leveraging the strengths of each method while mitigating their weaknesses. The flexibility in message handling, including the delegation of nested messages, further enhances the system's modularity and adaptability.

## Cache (Syscall Tree) Node Structure - The Ant-Routing Mechanism

The cache or syscall tree node structure is integral to the efficient operation of the PromiseGrid system. It serves as a hierarchical routing mechanism, caching successful paths for optimized future lookups.

### Syscall Tree (Ant Routing Mechanism)

1. **Node Structure**:
    - Each node in the syscall tree can have multiple children, representing different components of the parameters (`parms`).
    - Vertices in the tree are typically hashes of components like promise hashes, module hashes, and URL-encoded arguments.

2. **Dynamic Population**:
    - The syscall tree starts empty and is populated during kernel operation as it consults built-in modules and other available modules to handle received messages.
    - Successful paths (where modules accept and handle messages) are cached in the syscall tree.

3. **Acceptance History**:
    - The kernel stores both positive and negative acceptance history for all modules. This acceptance history is integrated into the syscall tree, allowing the kernel to skip modules that have previously rejected similar messages.
    - This mechanism reduces redundant checks and streamlines the message-handling process.

4. **Routing Messages**:
    - When a message is received, the kernel looks up the syscall tree to find the module whose key matches the most leading components of the parameters.
    - Modules are consulted in hierarchical order to find the longest matching parameter slice.

### Messages and Promises

1. **Unified Message Structure**:
    - The `Message` structure includes the promise as the first element in the `Parms` field. Recipients route or discard messages based on the leading promise.
    - Example: A `Message` with the promise as the first element and subsequent elements as parameters to be handled by the module.

### Multihash, Multibase, and Multicodec

PromiseGrid uses multihash and multibase to specify the first byte(s) of a promise hash. This ensures parsers can autodetect the hash format, be it binary, hex, or base58, enhancing compatibility and extensibility.

- **Multihash**: Provides a consistent way to specify multiple hash algorithms, ensuring flexibility and future-proofing.
- **Multibase**: Encodes multihash hashes such that their base (binary, hex, base58) is automatically interpreted by parsers.

### Combining Accept and HandleMessage Functions

Discussion of Pros and Cons:

#### Combined Function

1. **Simplicity**: Combines decision-making and handling into a single function.
2. **Consistency**: Ensures the decision to handle and the actual handling are tightly coupled, increasing consistency and reducing logic duplication.
3. **Efficiency**: Eliminates redundant checks by combining acceptance and handling.

#### Separate Functions

1. **Clarity**: Maintains clear separation between decision-making and handling logic.
2. **Early Rejection**: Allows for quick rejection of messages based on promise, module hash, or arguments without performing any handling.
3. **Modular Logic**: Facilitates modular and specialized acceptance and handling logic.

### Acceptance as Promise

Combining the `Accept` and `HandleMessage` functions would imply that the acceptance itself is a promise, aligning with the concept that "it's promises all the way down." This structure implies that:

1. **Acceptance as a Promise**: The first element in `Parms` indicates whether the module promises to handle the message.
2. **Handling as Promise Fulfillment**: `HandleMessage` attempts to fulfill this promise. If it fails (breaking the promise), it is logged and handled appropriately.

### Proposed Dynamic Syscall Table and Ant Routing

Implementing a dynamic syscall table with ant routing involves:

1. **Hierarchical Syscall Tree**: Each node in the tree can have multiple children representing different components of the parameters (`parms`).
2. **Caching Successful Paths**: Cache successful paths to optimize future routing based on leading parameter sequences.
3. **Populating During Operation**: The syscall tree starts empty and is populated during operation as modules accept and handle messages.

### Examples of Unified and Separate Function Approaches

#### Combined Function:

```go
type Module interface {
    ProcessMessage(ctx context.Context, parms ...interface{}) (Message, error)
}

type LocalCacheModule struct {
    cacheDir string
}

func NewLocalCacheModule(cacheDir string) *LocalCacheModule {
    return &LocalCacheModule{cacheDir: cacheDir}
}

func (m *LocalCacheModule) ProcessMessage(ctx context.Context, parms ...interface{}) (Message, error) {
    // Implement logic to process message, which includes accepting and handling
    // Return a promise message with acceptance and handling results or errors
}
```

#### Separate Functions:

```go
type Module interface {
    Accept(ctx context.Context, parms ...interface{}) (Message, error)
    HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error)
}

type LocalCacheModule struct {
    cacheDir string
}

func NewLocalCacheModule(cacheDir string) *LocalCacheModule {
    return &LocalCacheModule{cacheDir: cacheDir}
}

func (m *LocalCacheModule) Accept(ctx context.Context, parms ...interface{}) (Message, error) {
    // Implement logic to accept or reject based on parms
}

func (m *LocalCacheModule) HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error) {
    // Implement logic to handle the message following acceptance
}
```

### Conclusion

Implementing a hierarchical syscall tree and using a unified `PromiseMessage` structure simplifies the routing, acceptance, and handling of messages. By embedding known accept/reject messages in the kernel, the system can recognize and validate promises effectively. This approach aligns with the philosophy of "promises all the way down," ensuring trust, accountability, and optimized routing in a decentralized governance framework.

## Open Questions

- What are the specific conditions under which the kernel should invalidate or update cached syscall paths in the hierarchical syscall tree?
- How can we dynamically adjust the acceptance criteria of modules to adapt to changing workloads and conditions without manual intervention?
- What mechanisms can be implemented to handle broken promises more effectively, ensuring minimal disruption to the system?
- Regarding the design choice of using a separate `Accept()` and `HandleMessage()` method -- does this not break promise theory's principle of not making promises on behalf of others? If there is a separate `Accept()` and `HandleMessage()` method, this means that the `Accept()` code path is making a promise on behalf of the `HandleMessage()` code path. What are the implications of this? Should this design be changed?
- How does the kernel determine the best route when multiple modules provide promises to handle a message?
- How will the system manage latency in the context of decentralized and distributed nodes?

## Suggestions for Improving this Document

- Develop comprehensive error handling and logging for broken promises to ensure accountability and system robustness.
- Add more examples and case studies to illustrate core principles and use cases.
- Expand on the description of the hierarchical syscall tree and its role in routing and filtering messages.
- Update the sections with any additional questions or areas of exploration that arise during implementation and testing.
- Refine the glossary to include more specialized terms and concepts as they arise in the documentation.
- Incorporate visual aids (e.g., diagrams, flowcharts) to enhance understanding of system architecture and data flow.

2. **Acceptance as Promise**:
    - The kernel treats the acceptance of a message as a promise. The `Accept()` function returns a promise message instead of a simple boolean, allowing for richer interactions and more detailed responses.
    - This promise message contains meta-information about the module's capabilities and the conditions under which it will handle the message.

3. **Fulfillment and Accountability**:
    - Once a module accepts a message (thereby making a promise), it is expected to handle the message correctly. If `HandleMessage()` fails after `Accept()` returns true, it is considered a broken promise.
    - The kernel can track broken promises, enabling dynamic rerouting of requests and de-prioritizing unreliable modules.

### Cache and File System

1. **Cache as Filesystem**:
    - The cache uses filesystem separators (`/`) between each key component, and arguments are URL-encoded to ensure safe and consistent storage.
    - The kernel's local cache itself is implemented as a built-in module, ensuring uniform handling of cache lookups and module consultations.

2. **Integration with OPFS**:
    - The kernel uses the Origin Private File System (OPFS) for disk file access, providing a secure and performant storage backend.
    - The `afero` library is used to abstract filesystem interactions, allowing flexibility in choosing storage backends.

## Message Structure

- The `Message` structure includes the promise as the first element in the `Parms` field. Recipients route or discard messages based on the leading promise.

## Why (or Why Not) the Message Structure Should Have Module Hash as the Second Element

### Why the Module Hash Should Be the Second Element:

1. **Efficient Routing**: Including the module hash as the second element helps in efficiently routing the message to the correct module. This makes the lookup faster as the kernel can quickly identify which module is responsible for handling the message.

2. **Clarity and Determinism**: By specifying the module hash right after the promise hash, the message structure becomes more deterministic. It ensures that the correct module is always targeted, reducing ambiguity and potential errors.

3. **Security and Trust**: Having the module hash helps in verifying that the message is being handled by the appropriate module. This strengthens the security model by ensuring that promises are fulfilled by trusted modules.

4. **Modularity**: This structure supports a highly modular system where each module can clearly advertise which promises it can handle, leading to better organization and maintainability.

### Why the Module Hash Should Not Be the Second Element:

1. **Flexibility**: Some might argue that having a fixed slot for the module hash reduces flexibility. If the system evolves to use a different routing mechanism, this fixed structure might hinder adaptability.

2. **Overhead**: Including additional hashes in the message might increase the message size slightly, adding a minor overhead. However, this trade-off is often negligible compared to the benefits in routing efficiency and security.

Overall, while there are minor trade-offs, including the module hash as the second element right after the promise hash greatly enhances the system’s efficiency, clarity, security, and modularity.


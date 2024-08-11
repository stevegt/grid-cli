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

## Cache (Syscall Tree) Node Structure

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

### Why (or Why Not) the Message Structure Should Have Module Hash as the Second Element

### Why the Module Hash Should Be the Second Element:

1. **Efficient Routing**: Including the module hash as the second element helps in efficiently routing the message to the correct module. This makes the lookup faster as the kernel can quickly identify which module is responsible for handling the message.

2. **Clarity and Determinism**: By specifying the module hash right after the promise hash, the message structure becomes more deterministic. It ensures that the correct module is always targeted, reducing ambiguity and potential errors.

3. **Security and Trust**: Having the module hash helps in verifying that the message is being handled by the appropriate module. This strengthens the security model by ensuring that promises are fulfilled by trusted modules.

4. **Modularity**: This structure supports a highly modular system where each module can clearly advertise which promises it can handle, leading to better organization and maintainability.

### Why the Module Hash Should Not Be the Second Element:

1. **Flexibility**: Some might argue that having a fixed slot for the module hash reduces flexibility. If the system evolves to use a different routing mechanism, this fixed structure might hinder adaptability.

2. **Overhead**: Including additional hashes in the message might increase the message size slightly, adding a minor overhead. However, this trade-off is often negligible compared to the benefits in routing efficiency and security.

Overall, while there are minor trade-offs, including the module hash as the second element right after the promise hash greatly enhances the system’s efficiency, clarity, security, and modularity.

## Syscall Tree

- **Hierarchical Syscall Tree**: The kernel uses a hierarchical syscall tree to store acceptance history. This tree functions as an "ant routing" mechanism, caching successful paths to optimize future routing.

- **Dynamic Acceptance History**: The syscall tree captures positive and negative acceptance history. It starts empty and is populated during operation as the kernel consults built-in and other modules to handle received messages.

## Multihash, Multibase, and Multicodec

PromiseGrid uses multihash and multibase to specify the first byte(s) of a promise hash. This ensures parsers can autodetect the hash format, be it binary, hex, or base58, enhancing compatibility and extensibility.

- **Multihash**: Provides a consistent way to specify multiple hash algorithms, ensuring flexibility and future-proofing.
- **Multibase**: Encodes multihash hashes such that their base (binary, hex, base58) is automatically interpreted by parsers.

## Routing and Filtering

- **Optimized Routing**: In the case of a cache miss, the kernel consults modules based on the hierarchical syscall tree. It routes the message to the module with the longest matching parameter slice.

```go
func (k *Kernel) consultModules(ctx context.Context, parms ...interface{}) ([]byte, error) {
    bestMatch := k.findBestMatch(parms...)
    var promisingModules []Module

    for _, module := range bestMatch.modules {
        promise, err := module.Accept(ctx, parms...)
        if err != nil {
            continue // Log and handle error
        }
        if promise.Parms[0].(bool) {
            promisingModules = append(promisingModules, module)
            k.addSyscall(parms...) // Add to syscall tree
        }
    }

    for _, module := range promisingModules {
        acceptable, data, err := module.HandleMessage(ctx, true, parms...)
        if err != nil || !acceptable {
            continue
        }
        // Now handle the message
        data, err := module.HandleMessage(ctx, parms...)
        if err == nil {
            return data, nil
        }
        // Log broken promise if HandleMessage fails after Accept returned true
    }

    return nil, fmt.Errorf("no module could handle the request")
}
```

## Acceptance as a Form of Promise

- **Promises and Accountability**: The acceptance of a message itself is a promise. Modules track which requests they accept and must fulfill these promises by successfully handling the requests.
- The use of "accept" in this context aligns with the definitions in computing theory: An automaton accepts an input if it reaches an accepting state. Similarly, PromiseGrid modules accept a message if they can handle it, making a promise to process it, akin to how a Turing machine or a language automaton accepts strings belonging to a language.

## Integration with WebSocket

- Nodes interact with peers over the network via WebSocket connections.
- WebSocket is the message transport mechanism we're using for now, although other mechanisms may be adopted in the future.
- A sandboxed module can interact with the network by sending and receiving messages through the kernel.
- The kernel communicates with the outside world (both network and local I/O) via non-sandboxed modules.

```go
func handleWebSocket(ctx context.Context, k *Kernel, w http.ResponseWriter, r *http.Request) {
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            break
        }

        parms, err := deserializeMessage(message)
        if err != nil {
            continue
        }

        data, err := k.consultModules(ctx, parms...)
        if err != nil {
            continue
        }

        if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
            break
        }
    }
}
```


## Implementation Strategy

### Hierarchical Syscall Tree (Ant Routing)

- **Cached successful paths**: Enhance routing efficiency for future calls with similar parameters.
- **Tree Structure**: Each node represents a level of parameter matching, with leaf nodes potentially corresponding to specific modules or cached results.
    ```go
    type SyscallNode struct {
        modules  []Module
        children map[string]*SyscallNode
    }

    type Kernel struct {
        root *SyscallNode
        knownMessages map[string]Message
    }
    ```

### Simplifying Accept and HandleMessage into a Promise-Based Unified Approach

- **Promise-Based Decision**: Consolidate *Accept* and *HandleMessage* into one step, where the promise to process a message is fulfilled immediately, and a promise message is generated and returned.
- **PromiseMessage as Message**: Treat the *PromiseMessage* as merely the first element of `Parms`, introducing a seamless contract between caller and module.
- **Dynamic Sys.

#### Promise-Based Acceptance Mechanism

- **Acceptance as a Promise**: The first element in `Parms` indicates whether the module promises to handle the request.
    - **Accept Function**: 
      ```go
      func (m *Module) Accept(ctx context.Context, parms ...interface{}) (Promise, error)
      ```
    - **Handling Function**: 
      ```go
      func (m *Module) HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error)
      ```

### Consulting Modules with Promise-Based Acceptance

- **Optimized Module Consultation**: When a message is received, consult the syscall tree to find the longest matching parameter slice and route the message to the respective module.
- **Caching Successful Calls**: Use the syscall tree to cache paths of successful module consultations, akin to ant routing, where a trail is left for future use.

### Promise Message and Meta-Information Handling

- **Unified Message Handling**: Integrate additional meta-information (such as trust levels, module capabilities, etc.) into the Promise message returned by the Accept or Handling process.
- **Meta as Message.Payload**: Include any metadata or additional guarantees about the promise in the `Message.Payload`.

```go
func handleWebSocket(ctx context.Context, k *Kernel, w http.ResponseWriter, r *http.Request) { ... }
```

### Discussion: Integration with "Promises All the Way Down"

- **Trust and Accountability with Promise-Based Design**: In this design, every interaction within the system is based on trust represented by promises. The acceptance and fulfillment of these promises form the foundation of module interactions, enhancing system resilience and adaptability.

## Open Questions

- What are the specific conditions under which the kernel should invalidate or update cached syscall paths in the hierarchical syscall tree?
- How can we dynamically adjust the acceptance criteria of modules to adapt to changing workloads and conditions without manual intervention?
- What mechanisms can be implemented to handle broken promises more effectively, ensuring minimal disruption to the system?
- Regarding the design choice of using a separate `Accept()` and `HandleMessage()` method -- does this not break promise theory's principle of not making promises on behalf of others? If there is a separate `Accept()` and `HandleMessage()` method, this means that the `Accept()` code path is making a promise on behalf of the `HandleMessage()` code path. What are the implications of this? Should this design be changed?
- How does the kernel determine the best route when multiple modules provide promises to handle a message?
- How will the system manage latency in the context of decentralized and distributed nodes?
- What measures can be taken to ensure data consistency across the decentralized cache system over unreliable networks?

## Suggestions for Improving this Document

- Develop comprehensive error handling and logging for broken promises to ensure accountability and system robustness.
- Add more examples and case studies to illustrate core principles and use cases.
- Expand on the description of the hierarchical syscall tree and its role in routing and filtering messages.
- Update the sections with any additional questions or areas of exploration that arise during implementation and testing.
- Refine the glossary to include more specialized terms and concepts as they arise in the documentation.
- Incorporate visual aids (e.g., diagrams, flowcharts) to enhance understanding of system architecture and data flow.

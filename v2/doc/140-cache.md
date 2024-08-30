# Module Integration and Advanced Concepts

## Advanced Discussions

This document consolidates advanced discussions from previous versions and integrates key learnings, strategies, and design considerations.

### Cache and Modules

The kernel MUST treat modules as caches. In the event of a cache miss, the kernel MUST consult one or more modules to retrieve the requested value.

2. **Consulting Modules as Caches**:
    - Modules act as part of the cache mechanism. On a cache miss, the kernel consults one or more modules to retrieve the requested value.
    - The cache lookup function takes multiple arguments: `promiseHash`, `moduleHash`, and zero or more arguments from the caller’s message.
    - If a value isn't in the cache, it's retrieved from the module(s), added to the cache, and returned.

### Cache Structures

In the context of the PromiseGrid Kernel, the cache plays a critical role in optimizing and managing the execution and storage of data. Here's a summary of the key points regarding the caching strategy:

1. **Cache Structures**:
   - The cache **SHOULD BE** a directory tree rather than an in-memory map.
   - There **MAY BE** multiple caches, including the built-in cache in the kernel and caches provided by various modules.

2. **Integration with Modules**:
   - The kernel **MUST** treat modules as caches. In the event of a cache miss, the kernel **MUST** consult one or more modules to retrieve the requested value.
   - The kernel **SHOULD** load the built-in cache from embedded resources using Go’s embed feature.
   - The kernel **MAY** use the Origin Private File System (OPFS) for disk file access and **MAY** utilize the afero library to abstract filesystem interactions.

3. **Unified Interface**:
   - From the caller's perspective, there **SHALL BE** no difference between a cache lookup and a function call. Both operations **SHALL BE** treated as hashed function calls. The caller sends a message with a leading hash and any arguments and receives a message containing the requested content.

### Kernel and Caching

#### Does the kernel have its own cache, or does it always rely on one or more modules for caching?

The kernel **SHOULD** maintain its own cache but **MUST** also treat modules as part of the caching mechanism. When the kernel encounters a cache miss, it **MUST** consult one or more modules to retrieve the requested value. This approach allows for a distributed and scalable caching solution where multiple modules can contribute to the cache.

### Network Communications

#### Are network communications always done through modules, or can the kernel do some of that itself?

Network communications in the PromiseGrid Kernel model are primarily intended to be handled by modules rather than the kernel itself. This design aligns with the microkernel architecture principles, where the kernel provides minimal core functionalities and delegates most tasks, such as network communication, to service modules. The kernel’s role includes managing module execution, handling inter-process communication (IPC), and maintaining security and resource control.

Integrating a module-based architecture for network communications ensures modularity, flexibility, and ease of updating or replacing network-related functionalities without altering the core kernel. However, the kernel **MAY** include basic communication capabilities for essential operations, but these would generally be minimal and supportive in nature, primarily to coordinate and manage module interactions.

### Ant Routing Mechanism: The Syscall Tree

The cache tree acts like an "ant routing" mechanism:

    - On a cache miss, the kernel routes the message to the module whose syscall tree key matches the most leading parameter components.

2. **Acceptance as a Promise**:
    - Modules define an `Accept()` function that encompasses the acceptance criteria for promises, module hashes, and arguments.
    - The response from an `Accept()` call is considered a promise. If `HandleMessage()` fails after acceptance, it is considered a broken promise.

### Integration with Church, Turing, and Chomsky's Concept of "Accept"

1. **Computational Theory**:
    - Aligning acceptance with the concepts used by Church, Turing, and Chomsky. In computational theory, a machine or automaton accepts an input if it transitions into an accepting state.
    - Similarly, in PromiseGrid, modules act as recognizers or acceptors based on the promise hash, module hash, and arguments.

2. **Promises All the Way Down**:
    - The entire system leverages promises for each interaction layer. Acceptance and fulfillment of promises are managed consistently across modules, the syscall tree, and the kernel.

### Dynamic Syscall Table

The dynamic syscall table maintains positive and negative acceptance history:

1. **Populating the Syscall Table**:
    - The syscall table starts empty and populates during operation.
    - As the kernel consults built-in and other modules to handle messages, it updates the table based on which modules accept or reject specific parameter sets.

2. **Optimized Routing**:
    - On receiving a message, the kernel consults the syscall table to determine which modules to consult, optimizing message handling and reducing redundant checks.

### Example Implementations

#### Combined Function Approach

Combining decision-making and handling into a single function:

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
    // Implement logic for processing message, including acceptance and handling
    // Return a promise message with acceptance and handling results or errors
}
```

#### Separate Functions Approach

Using distinct functions for acceptance and handling:

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
    // Implement logic for accepting or rejecting based on parameters
}

func (m *LocalCacheModule) HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error) {
    // Implement logic for handling the message after acceptance
}
```

### Conclusion

By integrating promises at all levels and employing a hierarchical syscall tree with caching and acceptance history, PromiseGrid ensures trust, accountability, and efficient message handling. This structure supports a decentralized governance framework, robust system interactions, and reliable module functionality.


## Cache and Module Handling in the PromiseGrid Kernel

### Introduction

In the PromiseGrid Kernel, caching and module handling are central to optimizing and managing the execution and storage of data. This document discusses the strategies and implementation details of these mechanisms, highlighting their alignment with decentralized governance and microkernel architecture principles.

### Cache Structures

2. **Multiple Caches**:
    - There **MAY BE** multiple caches, including the built-in cache in the kernel and caches provided by various modules.
    - The kernel **SHOULD** load the built-in cache from embedded resources using Go's `embed` feature.
    - The kernel **MAY** use the Origin Private File System (OPFS) for disk file access and **MAY** utilize the `afero` library to abstract filesystem interactions.

### Treating Modules as Caches

1. **Role of Modules**:
    - The kernel **MUST** treat modules as caches. In the event of a cache miss, the kernel **MUST** consult one or more modules to retrieve the requested value.
    - Modules **MAY** contribute to the cache or provide the requested data dynamically.

2. **Unified Interface**:
    - From the caller's perspective, there **SHALL BE** no difference between a cache lookup and a function call. Both operations **SHALL BE** treated as hashed function calls. The caller sends a message with a leading hash and any arguments and receives a message containing the requested content.

### Acceptance and Promises

1. **Acceptance Criteria**:
    - Modules **MUST** define acceptance criteria for promises, module hashes, and arguments into a single function, `Accept()`.
    - By returning a promise message from `Accept` instead of a boolean, modules provide additional guarantees and meta-information.

2. **Promises as Acceptance**:
    - The `Accept` function returns a promise message that includes whether the module accepts the request and any relevant metadata. If `HandleMessage` fails, it is considered a broken promise.

3. **Ant Routing Mechanism**:
    - The syscall tree acts as an "ant routing" mechanism, caching successful paths to optimize future routing. In the future, similar calls follow the same path to the same module.
    - The syscall tree **MUST** use hierarchical keys and **SHOULD** filter based on whether any module accepts the leading parameters, matches the module hash, and accepts the arguments.

### Integration with Church, Turing, and Chomsky's Concept of "Accept"

1. **Computational Theory**:
    - The term "accept" aligns with Church, Turing, and Chomsky's use in computing theory and languages, where an automaton or machine accepts an input if it reaches an accepting state.
    - In PromiseGrid, modules act as recognizers for specific tasks, based on the promise hash, module hash, and arguments.

2. **Promises All the Way Down**:
    - The concept of "promises all the way down" integrates acceptance as a promise message, enhancing robustness and trustworthiness.
    - Each layer (modules, syscall tree, kernel) makes and fulfills promises based on the promises made by the layers below it.

### Kernel's Dynamic Syscall Tree

1. **Syscall Table and Acceptance History**:
    - The kernel **SHOULD** use a dynamic syscall table to store positive and negative acceptance history for all modules.
    - This table **SHOULD** start empty and be populated during operation as the kernel consults built-in and other modules to handle received messages.
    - The kernel **MUST** route messages to the module whose syscall table key matches the most leading parameters components, optimizing routing and reducing redundant checks.

### Responsibilities of Non-Sandboxed Modules vs. Core Functions of the Kernel

1. **Kernel Responsibilities**:
    - The kernel is responsible for **message routing only**. It maintains the syscall table, handles the routing of messages to the appropriate module based on the leading parameters, and performs minimal core functions.

2. **Non-Sandboxed Module Responsibilities**:
    - **All syscall-like operations** are delegated to non-sandboxed modules. This includes operations such as security handling, file reading and writing, and process launching. 
    - Non-sandboxed modules implement the logic for these operations while ensuring security, performance, and compliance with overall system policies.

### Combining Accept and HandleMessage Functions

#### Pros and Cons of Separate `Accept()` and `HandleMessage()` Functions versus a Single Function

##### Combined Function:

1. **Simplicity**: Combines decision-making and handling into a single function.
2. **Consistency**: Ensures the decision to handle and the actual handling are tightly coupled, increasing consistency and reducing logic duplication.
3. **Efficiency**: Eliminates redundant checks by combining acceptance and handling.

##### Separate Functions:

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

**Combined Function:

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

**Separate Functions:

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


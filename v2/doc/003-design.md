## Cache and Module Handling in the PromiseGrid Kernel

### Introduction

In the PromiseGrid Kernel, caching and module handling are central to optimizing and managing the execution and storage of data. This document discusses the strategies and implementation details of these mechanisms, highlighting their alignment with decentralized governance and microkernel architecture principles.

### Cache Structures

1. **Directory Tree Structure**:
    - The cache **SHOULD BE** a directory tree rather than an in-memory map. This approach ensures persistent storage and efficient management of large datasets.
    - Cache keys **MUST** use filesystem separators (`/`) between each key component, and arguments **MUST** be URL-encoded when building the cache key to handle special characters safely.

2. **Multiple Caches**:
    - There **MAY BE** multiple caches, including the built-in cache in the kernel and caches provided by various modules.
    - The kernel **SHOULD** load the built-in cache from embedded resources using Go’s `embed` feature.
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

### Conclusion

By integrating promises at every level and implementing a hierarchical syscall tree with caching and acceptance history, PromiseGrid ensures trust, accountability, and efficient message handling in a decentralized governance framework. The simplified message structure and consistent handling of cache and modules contribute to a robust and flexible system.

---

# Advanced Discussions on Caching, Promises, and Governance in PromiseGrid

## Caching Strategy in PromiseGrid Kernel

### Key Points from ADVICE.md

1. **Cache Structures**:
   * The cache **SHOULD BE** a directory tree rather than an in-memory map.
   * Cache keys **MUST** use filesystem separators (`/`) between each key component and arguments **MUST** be URL-encoded when building the cache key, in order to safely handle characters in file or directory names.
   * There **MAY BE** multiple caches, including the built-in cache in the kernel and caches provided by various modules.

2. **Integration with Modules**:
   * The kernel **MUST** treat modules as caches. In the event of a cache miss, the kernel **MUST** consult one or more modules to retrieve the requested value.
   * The kernel **SHOULD** load the built-in cache from embedded resources using Go’s embed feature.
   * The kernel **MAY** use the Origin Private File System (OPFS) for disk file access and **MAY** utilize the afero library to abstract filesystem interactions.

3. **Unified Interface**:
   * From the caller's perspective, there **SHALL BE** no difference between a cache lookup and a function call. Both operations **SHALL BE** treated as hashed function calls. The caller sends a message with a leading hash and any arguments and receives a message containing the requested content.

### Kernel and Caching Perspectives

**Does the kernel have its own cache, or does it always rely on one or more modules for caching?**
The kernel **SHOULD** maintain its own cache but **MUST** also treat modules as part of the caching mechanism. When the kernel encounters a cache miss, it **MUST** consult one or more modules to retrieve the requested value. This approach allows for a distributed and scalable caching solution where multiple modules can contribute to the cache.

### Network Communications

**Are network communications always done through modules, or can the kernel do some of that itself?**
Network communications in the PromiseGrid Kernel model are primarily intended to be handled by modules rather than the kernel itself. This design aligns with the microkernel architecture principles, where the kernel provides minimal core functionalities and delegates most tasks, such as network communication, to service modules. The kernel’s role includes managing module execution, handling inter-process communication (IPC), and maintaining security and resource control.

### Capability Tokens and Encoding

Several systems and frameworks implement self-contained capability tokens:

1. **Macaroons**: These are flexible authorization credentials that support delegation and attenuation. A macaroon is a bearer token that can encapsulate permissions and is augmented with caveats, which are conditions that must be satisfied for the macaroon to be considered valid.
2. **JSON Web Tokens (JWT)**: JWTs are widely used in web applications to assert claims between parties. They can include custom claims to specify permissions and the intended audience. A JWT can be signed to ensure authenticity and integrity.
3. **Caveats in Capability-Based Systems**: Traditional capability-based security systems sometimes support a form of caveats or restrictions within the tokens themselves to specify the scope of permissions and the authorized user.

### Discussion on the Leading Hash in Messages

In the PromiseGrid Kernel, messages MAY start with a promise hash followed by a module hash, and then additional arguments. This structure allows receivers to filter messages based on promises they are willing to accept and route the message to the appropriate module. Here are the pros and cons of the promise hash coming first:

**Pros:**
- **Enhanced Filtering**: Placing the promise hash first allows modules and nodes to quickly filter messages based on the promises they accept.
- **Trust and Governance**: This aligns with decentralized governance. Nodes can establish trust relationships by agreeing on specific promises they will accept.
- **Modular Routing**: Early identification of the promise allows the kernel to route messages to the appropriate module, ensuring efficient distributed handling.

**Cons:**
- **Complexity in Parsing**: Parsing multiple hashes increases the complexity of message handling.
- **Overhead**: Including multiple hashes adds overhead to message size, though this is generally offset by the efficiency gains in routing and filtering.

## Cache and Module Handling in the PromiseGrid Kernel

PromiseGrid uses a directory tree structure rather than an in-memory map for cache organization. Cache keys use filesystem separators (`/`) and URL-encoded arguments to handle special characters safely.

### Treating Modules as Caches

Modules are treated as part of the caching mechanism. On a cache miss, the kernel consults one or more modules to retrieve the requested value dynamically. This unified interface ensures no difference between a cache lookup and a function call from the caller's perspective.

### The Concept of Acceptance

1. **Acceptance Criteria**:
   - Modules **MUST** define acceptance criteria for promises, module hashes, and arguments into a single function, `Accept()`.
   - By returning a promise message from `Accept` instead of a boolean, modules provide additional guarantees and meta-information.

2. **Promises as Acceptance**:
   - The `Accept` function returns a promise message. If `HandleMessage` fails, it's considered a broken promise.

### Kernel’s Dynamic Syscall Tree

Acceptance history should be integrated into a dynamic syscall tree, using hierarchical keys to optimize routing based on leading parameters. This tree stores acceptance and rejection history, reducing redundant checks by leveraging the "ant routing" mechanism.

### Combining Accept and HandleMessage Functions

**Pros of Combined Function:**
- **Simplicity**: Combines decision-making and handling.
- **Consistency**: Tightly couples acceptance and handling, reducing logic duplication.
- **Efficiency**: Eliminates redundant checks.

**Pros of Separate Functions:**
- **Clarity**: Separates decision-making and handling logic.
- **Early Rejection**: Quickly rejects messages without handling.
- **Modular Logic**: Specialized acceptance and handling logic.

### Conclusion

By integrating promises at every level and implementing a hierarchical syscall tree with caching and acceptance history, PromiseGrid ensures trust, accountability, and efficient message handling in a decentralized governance framework. The simplified message structure and consistent handling of cache and modules contribute to a robust and flexible system.

### Integrated Summary

- **Key Points from ADVICE.md**: Discussed the cache strategy including multiple caches, unified interface, kernel's caching responsibilities, capability tokens, and hashing mechanisms.
- **Key Points from DISCUSS.md**: Focused on treating modules as caches, promise-based acceptance criteria, integration with Church, Turing, and Chomsky's "accept" concept, and combining accept/handle functions.

## Discussion Integration with Church, Turing, and Chomsky's Accept Concept

### Kernel's Dynamic Handling of Accept and HandleMessage

Implementing a dynamic syscall table with ant routing involves:
1. **Hierarchical Syscall Tree**: Each node in the tree can have multiple children representing different parameter components.
2. **Caching Successful Paths**: Cache successful paths to optimize future routing based on leading parameter sequences.
3. **Populating During Operation**: The syscall tree starts empty and is populated during operation as modules accept and handle messages.

### Unified Message Example

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
    // Implement logic to process message, including accepting and handling
}
```

By integrating promises at every level and implementing a hierarchical syscall tree with caching and acceptance history, PromiseGrid ensures trust, accountability, and efficient message handling in a decentralized governance framework.

## TODO: Future Enhancements

1. **Develop Comprehensive Error Handling**: Ensure system robustness by tracking and logging broken promises.
2. **Expand Acceptance Criteria**: Enhance the acceptance criteria to adapt dynamically to changing conditions.
3. **Implement Visual Aids**: Incorporate diagrams and flowcharts to enhance understanding of system architecture and data flow.
4. **Refine Glossary**: Update and expand the glossary to include more specialized terms and concepts as they arise.

This document serves as a comprehensive guide to the advanced aspects of caching, promise handling, and governance in PromiseGrid, integrating discussions from various sources to provide a robust framework for development.


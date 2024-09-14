# Cache and Promise Handling in the PromiseGrid Kernel

## Introduction

In the PromiseGrid Kernel, caching and module handling are crucial for optimizing the retrieval and execution of data. This document outlines the strategies and implementation details for these systems, emphasizing their alignment with decentralized governance and microkernel architectural principles.

## Kernel Services

- The kernel **MUST** provide a built-in cache for efficient data retrieval and storage.
- The built-in cache **SHOULD** be loaded from embedded resources using Go's `embed` feature.
- The built-in cache **MUST** include one or more modules (microkernel services) for handling filesystem and network access, dependent on which platform the kernel is running on (WASM, Linux, Windows, etc.)
- The filesystem module(s) **MAY** use the Origin Private File System (OPFS) for disk file access when running in WASM, and **MAY** use the `afero` library to abstract filesystem interactions.
- The network module(s) **MAY** use XXX for network access when running in WASM, and **MAY** use the `net` package for network interactions when native.

## Cache Structures

### Built-In and Modular Cache

1. **Multiple Caches**:
   - There **MAY BE** multiple caches, including the built-in cache in the kernel and caches provided by various modules.
   - The kernel **SHOULD** load the built-in cache from embedded resources using Go’s `embed` feature.

2. **Filesystem Integration**:
   - Non-sandboxed modules **MAY** use the Origin Private File System (OPFS) for disk file access.
   - Non-sandboxed modules **MAY** utilize the `afero` library to abstract filesystem interactions, ensuring compatibility across different platforms.

## Treating Modules as Caches

### Unified Interface for Cache and Function Calls

1. **Role of Modules**:
   - The kernel **MUST** treat modules as caches. In the event of a cache miss, the kernel **MUST** consult one or more modules to retrieve the requested value.
   - Modules **MAY** contribute to the cache or provide the requested data dynamically.

2. **Unified Interface**:
   - From the caller's perspective, there **SHALL BE** no difference between a cache lookup and a function call. Both operations **SHALL BE** treated as byte sequence completions. The caller sends a message consisting of a byte sequence and receives a message containing the remainder of the sequence.

## internal channels, external function calls

- internally, the kernel uses Go channels 
- all modules communicate with the kernel using functions calls and
  callbacks


## Acceptance and Promises

### Acceptance as Promise

1. **Acceptance Criteria**:
   - Modules **MUST** define acceptance criteria for promises, module hashes, and arguments into a single function, `Accept()`.
   - By returning a promise message from `Accept` instead of a boolean, modules provide additional guarantees and meta-information.

2. **Promise-Based Acceptance**:
   - The `Accept` function returns a promise message including details on whether the module accepts the request and any relevant metadata. If `HandleMessage` fails, it is considered a broken promise.

### Ant Routing Mechanism

3. **Dynamic Routing and Optimization**:
   - The syscall tree acts as an "ant routing" mechanism, caching successful paths to optimize future routing. Similar future calls follow the same paths to the same module.
   - The syscall tree **MUST** use hierarchical keys and **SHOULD** filter based on whether any module accepts the leading parameters, matches the module hash, and accepts the arguments.

## Integration with Computational Theory

### Alignment with Church, Turing, and Chomsky's "Accept" Concept

1. **Computational Theory**:
   - The term "accept" aligns with the computational theory usage (Church, Turing, Chomsky), where an automaton or machine accepts an input if it reaches an accepting state.
   - In PromiseGrid, modules act as recognizers for specific tasks based on the promise hash, module hash, and arguments.

2. **Promises All the Way Down**:
   - The "promises all the way down" concept integrates acceptance as a promise message, enhancing robustness and trustworthiness.
   - Each layer (modules, syscall tree, kernel) makes and fulfills promises based on the promises made by the layers below it.

## Kernel's Dynamic Syscall Tree

### Dynamic and Optimized Message Routing

1. **Syscall Tree and Acceptance History**:
   - The kernel **SHOULD** use a dynamic syscall table to store both positive and negative acceptance histories for all modules.
   - This table **SHOULD** start empty and be populated as the kernel handles messages, consulting built-in and other modules.
   - Messages **MUST** be routed to the module whose syscall tree key matches the most leading parameters, optimizing routing efficiency and reducing redundant checks.

## Conclusion

The PromiseGrid Kernel's design for caching, module handling, and promise-based acceptance ensures a robust, flexible, and efficient system for decentralized governance and cooperation. By treating acceptance as a promise and integrating a hierarchical syscall tree with ant routing mechanisms, the kernel maintains trust, accountability, and optimized performance, consistent with principles of computational theory and modular architecture.

## Q&A

### Known Microkernel Architectures Using Capability-Based Security

- Are there any known microkernel architectures that use a capability-based security model? If so, how do they handle permissions and access control?

Yes, there are known microkernel architectures that use a capability-based security model. One notable example is the **seL4** microkernel.
- **seL4**: It's a high-assurance, high-performance microkernel developed with a strong focus on security and formal verification. In seL4, capabilities are used to control access to kernel objects, providing fine-grained access control. These capabilities are unforgeable tokens that describe what operations can be performed on objects (such as memory segments, threads, and IPC channels). The kernel maintains the mapping and distribution of these capabilities, ensuring that only authorized entities can access and manipulate kernel resources.

Handling Permissions and Access Control in **seL4**:
1. **Capabilities**: In seL4, capabilities are fundamental to its security model. Each capability specifies a set of rights to a particular object, and these rights determine what actions the holder can perform on the object.
2. **Capability Space**: Each thread/task in seL4 has a capability space, which is a structure that holds all the capabilities that the thread possesses. This ensures that threads can only operate on objects for which they have the necessary capabilities.
3. **Capability Invocation**: To perform an operation on an object, a thread must invoke the corresponding capability. The kernel then checks the rights encoded in the capability to determine whether the operation is permitted.
4. **Controlled Delegation**: Capabilities can be delegated (or propagated) to other threads, but the delegation is controlled by the kernel to prevent unauthorized access. This allows for flexible and secure sharing of resources within a system.

By leveraging a capability-based security model, microkernels like seL4 ensure that access control is both fine-grained and secure, providing a robust foundation for building secure systems.

#### Examples of Existing Microkernels

In a typical microkernel architecture, the kernel has minimal knowledge of permissions, and the responsibility for handling permissions is generally left to microservers (drivers). The kernel focuses on providing essential services such as communication and resource management, while higher-level functions, including access control and permissions, are managed by user-level services.

1. **seL4**: As previously mentioned, seL4 uses a capability-based security model where the kernel maintains capability mappings but does not deeply involve itself in access control logic. Instead, it relies on the capabilities held by the user-level services to determine access rights.

2. **Mach**: The Mach microkernel provides basic Inter-Process Communication (IPC) mechanisms and leaves most of the higher-level functionality, including access control, to user-space programs. The kernel does not have intrinsic knowledge of the permissions but enforces access controls passed to it via messages.

3. **L4**: The L4 microkernel architecture also uses a minimalistic design philosophy, where the kernel provides basic mechanisms, and policy, including permissions, is implemented in user space. Access control is managed through capabilities issued and checked by user-space services.

By delegating permission and access control management to microservers or user-level services, microkernel architectures can maintain a small and efficient core, allowing for easier verification and higher reliability.

### Microkernel interaction with userland drivers

- How do microkernels interact with userland drivers?  Does the driver register with the kernel, or does the kernel load the driver as a part of the boot process, already knowing about the driver's address space and abilities?

Microkernels interact with userland drivers through a well-defined communication interface, usually using Inter-Process Communication (IPC) mechanisms. The typical process is as follows:

1. **Registration**: Userland drivers register with the microkernel during the system's initialization phase. The registration involves specifying the driver's capabilities and the resources it will manage. During this process, the kernel may allocate specific address spaces and IPC channels for communication with the driver.

2. **Dynamic Loading**: The kernel does not necessarily need to know about the driver's address space and abilities beforehand. Drivers can be dynamically loaded as required, allowing for a more modular and flexible system. The microkernel will map the driver's address space as part of the registration or loading process.

3. **Message Passing**: Once loaded and initialized, userland drivers interact with the kernel and other system components via message passing. This interaction can include handling hardware interrupts, servicing IO requests, or managing other resources. The microkernel facilitates this communication through its IPC mechanism, ensuring that messages are delivered to the appropriate recipient.

In summary, userland drivers typically register with the kernel, which then maps their address spaces and sets up communication channels. These drivers can be dynamically loaded, allowing the kernel to extend its capabilities without needing to know about the driver specifics in advance.

### Caches, Invocation, and Logging

### How do we know when we need to load a module from the cache, execute it, and store the result back in the cache?

In the PromiseGrid Kernel, the determination of when to load a module from the cache, execute it, and store the result back in the cache follows a systematic approach:

1. **Cache Lookup**: When a message arrives, the kernel first performs a cache lookup using the leading hash (capability token, function address) and the provided arguments. If the requested function call and its arguments are already in the cache, the kernel retrieves and returns the cached result.

2. **Cache Miss**: If the cache lookup fails (cache miss), the kernel proceeds to consult the appropriate module(s) to handle the request. This involves invoking the module's function corresponding to the capability token and passing the arguments from the message.

3. **Module Execution**: The module executes the function with the given arguments. During this execution, the module can dynamically generate the response based on the input or retrieve the necessary data from other sources.

4. **Store Result in Cache**: Once the module generates the response, the kernel stores this result back in the cache. This ensures that subsequent requests with the same capability token and arguments can be served directly from the cache, optimizing future retrievals and reducing redundant computations.

By following this procedure, the PromiseGrid Kernel ensures efficient use of the cache while leveraging modular execution to handle cache misses. This approach harmonizes decentralized cache handling with the overall system's modular and promise-based architecture.

## Open Questions About Caching Mechanisms

- Does the cache even need to know anything about modules, or is it just a simple nested key-value store?

The cache in the PromiseGrid kernel is designed to be a simple trie; facilitating the efficient completion of byte sequences. It is not
necessary for the cache to know about the specifics of modules,
promises, or protocols; it primarily manages the storage and retrieval of messages.

- trie nodes should have a handlers slice

### Example of How the System Can Handle Cache and Module Execution:

- **Message Processing Flow**:
  - The system receives a message with a byte sequence.
  - The leading bytes are used for cache lookup.
  - On a cache miss, the kernel invokes the appropriate module to handle or complete the message.
  - The result is stored back in the trie for future lookups.

### Example Conversation Flow

- Every message is a promise, an assertion of truth. When we talk about the cache and the messages it stores, we are talking about these messages.  How will this work in practice?  What's an example conversation flow?

To illustrate how the cache and promise-based handling work in practice, let's consider an example conversation flow:

**Scenario**: A node requests the current weather for a specific location from another node that provides weather information.

0. **Ad**
    - Node B has a weather module that can retrieve weather data for various locations.  It broadcasts this capability to the network, promising that it can provide weather information for specific locations.

1. **Request Message**:
    - Node A sends a request message with node B's weather capability token, asking Node B to
    return the current weather for `Location X`. The request message includes:
        - XXX
        - Capability token: A unique identifier representing node B's promise to provide weather data.
        - Module Hash: The identifier of the weather module.
        - Arguments: The location `X`.

2. **Cache Lookup**:
    - The kernel of Node B receives the request and performs a cache lookup using the promise hash, module hash, and arguments.
    - If a cached response exists (e.g., weather data for `Location X`), the kernel retrieves it and sends it back to Node A.

3. **Cache Miss and Module Consultation**:
    - If the cache lookup fails (cache miss), the kernel consults the appropriate module to handle the request.
    - The kernel invokes the weather module's function corresponding to the capability token (promise hash) and passes the arguments.
    - The weather module processes the request, retrieves the current weather data for `Location X`, and generates the response.

4. **Execute and Cache Result**:
    - The weather module sends the response back to the kernel.
    - The kernel caches the result using the same keys (promise hash, module hash, and arguments) to ensure future requests can be served directly
 from the cache.
    - The kernel sends the response back to Node A.

5. **Response Message**:
    - Node A receives the response message containing the current weather data for `Location X`.
    - The promise made by Node B to provide the weather data is fulfilled.

In this flow, every message exchanged is a promise, an assertion of truth, and the system leverages the cache and modules to handle requests efficiently and reliably.


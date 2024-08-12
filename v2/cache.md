# Cache and Promise Handling in the PromiseGrid Kernel

## Introduction

In the PromiseGrid Kernel, caching and module handling are crucial for optimizing the retrieval and execution of data. This document outlines the strategies and implementation details for these systems, emphasizing their alignment with decentralized governance and microkernel architectural principles.

## Cache Structures

1. **Directory Tree Structure**:
    - The cache **SHOULD BE** a directory tree rather than an in-memory map. This structure ensures persistent storage and efficient management of large datasets.
    - Cache keys **MUST** use filesystem separators (`/`) between each key component, and arguments **MUST** be URL-encoded when building the cache key to handle special characters safely.

2. **Multiple Caches**:
    - There **MAY BE** multiple caches, including the built-in cache in the kernel and caches provided by various modules.
    - The kernel **SHOULD** load the built-in cache from embedded resources using Go’s `embed` feature.
    - The kernel **MAY** use the Origin Private File System (OPFS) for disk file access and **MAY** utilize the `afero` library to abstract filesystem interactions.

## Treating Modules as Caches

1. **Role of Modules**:
    - The kernel **MUST** treat modules as caches. In the event of a cache miss, the kernel **MUST** consult one or more modules to retrieve the requested value.
    - Modules **MAY** contribute to the cache or provide the requested data dynamically.

2. **Unified Interface**:
    - From the caller's perspective, there **SHALL BE** no difference between a cache lookup and a function call. Both operations **SHALL BE** treated as hashed function calls. The caller sends a message with a leading hash and any arguments and receives a message containing the requested content.

## Acceptance and Promises

1. **Acceptance Criteria**:
    - Modules **MUST** define acceptance criteria for promises, module hashes, and arguments into a single function, `Accept()`.
    - By returning a promise message from `Accept` instead of a boolean, modules provide additional guarantees and meta-information.

2. **Promises as Acceptance**:
    - The `Accept` function returns a promise message that includes whether the module accepts the request and any relevant metadata. If `HandleMessage` fails, it is considered a broken promise.

3. **Ant Routing Mechanism**:
    - The syscall tree acts as an "ant routing" mechanism, caching successful paths to optimize future routing. In the future, similar calls follow the same path to the same module.
    - The syscall tree **MUST** use hierarchical keys and **SHOULD** filter based on whether any module accepts the leading parameters, matches the module hash, and accepts the arguments.

## Integration with Church, Turing, and Chomsky's Concept of "Accept"

1. **Computational Theory**:
    - The term "accept" aligns with Church, Turing, and Chomsky's use in computing theory and languages, where an automaton or machine accepts an input if it reaches an accepting state.
    - In PromiseGrid, modules act as recognizers for specific tasks based on the promise hash, module hash, and arguments.

2. **Promises All the Way Down**:
    - The concept of "promises all the way down" integrates acceptance as a promise message, enhancing robustness and trustworthiness.
    - Each layer (modules, syscall tree, kernel) makes and fulfills promises based on the promises made by the layers below it.

## Kernel's Dynamic Syscall Tree

1. **Syscall Tree and Acceptance History**:
    - The acceptanceHist and syscallTable **SHOULD BE** the same. This dynamic syscall tree **SHOULD** store acceptance history for all modules.
    - The kernel **SHOULD** store positive and negative acceptance history for all modules such that when it receives a message, it can look up which modules accept the leading parms, skip the modules that reject the leading parms, and call accepting modules, providing the full set of parms.
    - This table **SHOULD** start empty and be populated during operation as the kernel consults built-in and other modules to handle received messages.
    - The kernel **MUST** route messages to the module whose syscall tree key matches the most leading parameter components, optimizing routing and reducing redundant checks.

## Conclusion

The design of caching, module handling, and promise-based acceptance in the PromiseGrid Kernel ensures a robust, flexible, and efficient system for decentralized governance and cooperation. By treating acceptance as a promise and integrating a hierarchical syscall tree with ant routing mechanisms, the kernel maintains trust, accountability, and optimized performance consistent with principles of computational theory and modular architecture.

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

### How do we know when we need to load a module from the cache, execute it, and store the result back in the cache?

In the PromiseGrid Kernel, the determination of when to load a module from the cache, execute it, and store the result back in the cache follows a systematic approach:

1. **Cache Lookup**: When a message arrives, the kernel first performs a cache lookup using the leading hash (capability token, function address) and the provided arguments. If the requested function call and its arguments are already in the cache, the kernel retrieves and returns the cached result.

2. **Cache Miss**: If the cache lookup fails (cache miss), the kernel proceeds to consult the appropriate module(s) to handle the request. This involves invoking the module's function corresponding to the capability token and passing the arguments from the message.

3. **Module Execution**: The module executes the function with the given arguments. During this execution, the module can dynamically generate the response based on the input or retrieve the necessary data from other sources.

4. **Store Result in Cache**: Once the module generates the response, the kernel stores this result back in the cache. This ensures that subsequent requests with the same capability token and arguments can be served directly from the cache, optimizing future retrievals and reducing redundant computations.

By following this procedure, the PromiseGrid Kernel ensures efficient use of the cache while leveraging modular execution to handle cache misses. This approach harmonizes decentralized cache handling with the overall system's modular and promise-based architecture.

### Cache features and capabilities 

- Does the cache even need to know anything about modules, or is it just a simple nested key-value store?

The cache in the PromiseGrid kernel is designed to be a simple nested key-value store that focuses on storing and retrieving messages based on their positional parameters. It does not need to be aware of the specifics of modules, promises, or protocols. Its primary role is to facilitate efficient retrieval of responses to hashed function calls, and it can be extended with capability-based security mechanisms if needed. The actual handling of permissions and validation of capabilities can be delegated to modules, maintaining a clean separation of concerns.

## Open Questions

- It sounds like a cache node struct might include a field that marks or flags the node as being an executable, an argument, or a result message. Is that a good design?  Or is it better to assume that the first key field is always an executable and the remaining fields are positional arguments?  

- The cache stores messages intact, so the cache index tree is built from the message's positional parameters, starting with the first parameter in position zero, which we've been calling the promise hash. It appears that the cache knows very little about protocols, promises, or anything else other than the positional parameters of the message. The cache is a simple nested key-value store. The value is the message, and the key is the message's parameters.

- As far as permissions and capabilities go, we might have a situation where a cache key or value is encrypted, and the cache node is unlocked by a capability. This would be a way to implement a capability-based security model; the capability would be a key that is used by the kernel to unlock the cache node.
- Alternatively, the kernel knows nothing about capabilities, and it is up to modules to verify that the caller has the necessary permissions to access a resource. This would be a more traditional capability model, where the capability is a token that is passed to the module, and the module verifies the token before granting access to the resource.

- When we talk about the cache and the messages it stores, are we talking about request messages or response messages?  Which is it?  Or is it both? Or are they both the same thing?  Is every message a promise, an assertion of truth?


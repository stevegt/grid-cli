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

1. **Syscall Table and Acceptance History**:
    - The kernel **SHOULD** use a dynamic syscall tree to store positive and negative acceptance histories for all modules.
    - This table **SHOULD** start empty and be populated during operation as the kernel consults built-in and other modules to handle received messages.
    - The kernel **MUST** route messages to the module whose syscall tree key matches the most leading parameters components, optimizing routing and reducing redundant checks.

## Conclusion

The design of caching, module handling, and promise-based acceptance in the PromiseGrid Kernel ensures a robust, flexible, and efficient system for decentralized governance and cooperation. By treating acceptance as a promise and integrating a hierarchical syscall tree with ant routing mechanisms, the kernel maintains trust, accountability, and optimized performance consistent with principles of computational theory and modular architecture.

## open questions

It sounds like a cache node struct might include a field that marks or
flags the node as being an executable, an argument, or a result.  Is
that correct?  Or do we simply store the node's sequence number in
that field, i.e. the node's position in the received message?  Oh
wait, the cache stores messages intact, so the cache index tree is
built from the message's positional parameters, starting with the
first parameter in position zero, which we've been calling the promise
hash.  It appears that the cache actually knows very little about
protocols, promises, or anything else other than the positional
parameters of the message.  The cache is a simple nested key-value
store.  The value is the message, and the key is the message's
parameters.  But then how do we know when we need to load a module
from the cache, execute it, and store the result back in the cache?

As far as permissions and capabilities go, we might have a situation
where a cache key or value is encrypted, and the cache node is 
unlocked by a capability.  This would be a way to implement a
capability-based security model.  The capability would be a key
that is used by the kernel to unlock the cache node.  Alternatively,
the kernel knows nothing about capabilities, and it is up to modules
to verify that the caller has the necessary permissions to access
a resource.  This would be a more traditional capability model, where
the capability is a token that is passed to the module, and the module
verifies the token before granting access to the resource.  



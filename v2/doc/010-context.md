### Background and Evolution of the Project

The project began with the idea of deploying and managing a multipoint WireGuard network using PromiseGrid, leveraging decentralized computing capabilities alongside secure network tunneling. Initial discussions led to the design of a CLI approach with a syntax facilitating promises and arguments to handle various tasks.

### Key Decisions, Conclusions, and Open Questions

1. **Use of Decentralized Content-Addressable Storage (CAS)**
   - **Decision**: Modules are fetched and stored using their hashes in a decentralized CAS.
   - **Conclusion**: This allows for a flexible, decentralized approach where modules can dynamically be fetched and used based on their hashes.
   - **Open Question**: How to ensure the security and authenticity of fetched modules.

2. **Local Cache and Embedded Built-ins**
   - **Decision**: Use a directory-based cache for storing fetched modules, and embed essential built-in modules using Go's embed feature.
   - **Conclusion**: Provides persistency and ease of management for large datasets by leveraging the filesystem.
   - **Open Question**: How to manage updates and invalidation of cached modules.

3. **Unified Message Structure**
   - **Decision**: Simplify the message structure by including the promise as the first element in the `Parms` field.
   - **Conclusion**: Streamlines the interface and handling of messages, making it easier to reason about promises and parameters.
   - **Open Question**: How to standardize and version these promise-based messages.

4. **Hierarchical Syscall Tree**
   - **Decision**: Implement a hierarchical syscall tree to cache successful paths and optimize routing.
   - **Conclusion**: Enhances efficiency by caching the paths of previously successful calls and optimizing future routing.
   - **Open Question**: The best way to handle concurrency and synchronization in the syscall tree.

5. **Promises and Governance**
   - **Decision**: Treat acceptance as a promise, ensuring that modules can reliably accept and handle messages.
   - **Conclusion**: This approach introduces a layer of trust and accountability, enabling decentralized governance.
   - **Open Question**: How to handle broken promises and ensure reliability and resilience across the network.

6. **Message Serialization and Routing**
   - **Decision**: Use URL-encoded arguments and filesystem-safe separators in cache keys; always serialize messages with a leading hash.
   - **Conclusion**: Ensures safe and consistent storage and retrieval of cache values while maintaining compatibility with file systems.
   - **Open Question**: How to balance serialization efficiency with readability and debug-ability.

### Detailed Design Choices

1. **Embedded Built-in Modules**:
   - Built-in modules are embedded into the kernel at compile time using Go's `embed` feature, ensuring that critical functions are always available and executed efficiently.

2. **Filesystem-based Cache**:
   - The cache uses directory structures for storing values based on URL-encoded arguments and safely formatted keys, with modules consulted automatically on cache misses.

3. **Hierarchical Syscall Tree**:
   - The syscall tree caches successful module calls by mapping parameter tuples to module nodes, following the most specific path available.

4. **Unified Message Structure**:
   - Messages are structured with promises as the first element in the parameters field, simplifying the interface and maintaining a clear contract between accepting and handling modules.

5. **Promises All the Way Down**:
   - Acceptance is treated as a promise, with modules making and fulfilling promises to handle requests accurately. Broken promises are logged, ensuring accountability.

### Governance and Trust Mechanisms

The project embraces decentralized governance, where trust relationships are formed based on shared promises. Modules and nodes interact by making and fulfilling promises, with acceptance acting as a binding contract. This introduces a robust layer of trust, enabling dynamic and resilient system interactions.

### Open Questions and Future Work

1. **Security and Authenticity**:
   - How can we ensure that dynamically fetched modules are secure and authentic?
2. **Concurrency and Synchronization**:
   - What are the best practices for handling concurrency within the syscall tree to ensure efficient and safe operations?
3. **Handling Broken Promises**:
   - How should the system handle broken promises, and what mechanisms can ensure resilience and reliability?
4. **Standardization and Versioning**:
   - How can we standardize promise-based messages to ensure compatibility and ease of use across different versions and implementations?
5. **Balance in Serialization**:
   - Finding the right balance between efficient serialization and maintaining readability and ease of debugging.

This document captures the evolving thought processes, decisions, and ongoing questions addressed throughout the project's development. The approach aims to create a flexible, efficient, and trust-based system, aligning with the principles of decentralized governance and modularity.


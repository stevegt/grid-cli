## Cache Discussion Summary

In the context of the PromiseGrid Kernel, the cache plays a critical role in optimizing and managing the execution and storage of data. Here's a summary of the key points regarding the caching strategy:

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

### Kernel and Caching

#### Does the kernel have its own cache, or does it always rely on one or more modules for caching?

The kernel **SHOULD** maintain its own cache but **MUST** also treat modules as part of the caching mechanism. When the kernel encounters a cache miss, it **MUST** consult one or more modules to retrieve the requested value. This approach allows for a distributed and scalable caching solution where multiple modules can contribute to the cache.

### Network Communications

#### Are network communications always done through modules, or can the kernel do some of that itself?

Network communications in the PromiseGrid Kernel model are primarily intended to be handled by modules rather than the kernel itself. This design aligns with the microkernel architecture principles, where the kernel provides minimal core functionalities and delegates most tasks, such as network communication, to service modules. The kernel’s role includes managing module execution, handling inter-process communication (IPC), and maintaining security and resource control.

Integrating a module-based architecture for network communications ensures modularity, flexibility, and ease of updating or replacing network-related functionalities without altering the core kernel. However, the kernel **MAY** include basic communication capabilities for essential operations, but these would generally be minimal and supportive in nature, primarily to coordinate and manage module interactions.

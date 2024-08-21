
### Treating Lazy-Loading from Disk and Network 

In the PromiseGrid system, lazy-loading a node from disk should be treated the same as lazy-loading a node from a remote host over the network; both are IO-bound operations. 

- In both cases, an in-memory trie miss generates a call via the kernel.
- The kernel handles disk and network I/O.
- The trie code does not import any OPFS or `afero` library; file and network I/O are the responsibility of other microkernel services.

This approach ensures that the trie code remains focused on its primary responsibility—efficiently managing in-memory trie data structures—while delegating I/O operations to dedicated microkernel services. This separation of concerns leads to cleaner, more maintainable code and leverages the existing microkernel services for handling diverse I/O operations.


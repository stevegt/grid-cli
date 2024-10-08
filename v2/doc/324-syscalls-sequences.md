# Syscalls as Sequence Completions

## Overview

In the PromiseGrid framework, syscalls (system calls) are represented as sequences of bytes that denote specific operations or requests. The kernel processes these sequences using a trie-based structure to optimize and manage their execution. This document explores the concept of treating syscalls as sequence completions, the use of a trie to hardcode these sequences, the implications of including timestamps for sequences with side effects, and employing stdout for syscall communication.

## Syscalls as Byte Sequences

### Conceptual Model

Each syscall is represented as a unique sequence of bytes. These sequences act like signatures, allowing the kernel to recognize and process them accordingly. For example:
- `file_read` could be encoded as a specific byte sequence that initiates a file read operation.
- `file_write` includes additional bytes representing the data to be written, with a timestamp for versioning and synchronization.

### Hardcoding in a Trie

The kernel hardcodes these byte sequences in an embedded trie, which acts as the root structure where all known syscall sequences are stored. This enables efficient lookup and matching of incoming syscall requests.

The kernel might handle these hardcoded byte sequences itself, or it might delegate the handling to non-sandboxed modules.

- **Trie Structure**: Each node in the trie represents a byte or a sequence of bytes, creating a path from the root to the leaf nodes, which correspond to complete syscalls.

```
Trie example:
(root)
   |
  0x01 (file_read)
   |
  0x02 (file_write)
     |
    data...
```

### Including Timestamps

Incorporating timestamps in the request message sequence allows the kernel to handle synchronization and versioning effectively. For example:
- `file_write` request: `0x02[timestamp][filename][data]`
- By including a timestamp, the kernel ensures that writes are applied in the correct order, preventing race conditions and conflicts.

XXX disk i/o must be handled by non-sandboxed modules.  The kernel is only a message router.

**Timestamp Example**:

Consider two modules attempting to write to the same file concurrently. The timestamp ensures that the writes are applied in the sequence they were issued:

- Module A sends: `0x02[20230401123000][fileA.txt][Hello, world!]`
- Module B sends: `0x02[20230401123100][fileA.txt][Welcome to PromiseGrid!]`

Here, the kernel ensures that `Module A's` write is processed before `Module B's` write, maintaining the correct sequence based on their timestamps.

### Handling Timestamps in the Trie

When the kernel receives a syscall request with a timestamp, it incorporates the timestamp into the trie traversal and processing:

1. **Reception**:
   - The kernel receives the byte sequence: `0x02[20230401123000][fileA.txt][Hello, world!]`
2. **Trie Lookup with Timestamp**:
   - The kernel locates the `0x02` node (indicating `file_write`).
   - It further traverses to match the `20230401123000` timestamp node.
   - Finally, it matches the `fileA.txt` filename and processes the `Hello, world!` data.
3. **Conflict Resolution**:
   - If another request with a later timestamp for the same file is received, the kernel ensures it processes this subsequent request accordingly.

```
Trie:
(root) - 0x02 (file_write)
           |
      20230401123000 (timestamp)
           |
        fileA.txt (filename)
           |
  Hello, world! (data)
```

By embedding timestamps within the byte sequences and trie structure, the kernel achieves a robust mechanism for maintaining data integrity and sequencing.

## Using stdout for Syscall Communication

### Mechanism

In PromiseGrid, modules send syscall requests via stdout, similar to the message-passing mechanism in Mach. Each message includes:
- **Byte Sequence**: Represents the syscall.
- **Port**: Indicates where the module expects to receive the response.

```
Example:
Module -> Kernel: [byte sequence for syscall] [port for response]
Kernel -> Module: [response data]
```

### Port as Capability

A port number acts as a capability, granting the module the right to communicate and receive responses. This port is often represented by a large hash, ensuring uniqueness and security.

## Integration with Trie Structures

### Root Trie

The embedded trie can be the root structure where all other tries are mounted, creating a unified system for handling syscalls.

### Treating stdout Messages as Tries

Stdout messages can be treated as part of the syscall trie, enabling the kernel to efficiently match and route messages based on their byte sequences.

1. **Incoming Message**:
   - The kernel receives a message on stdout.
   - It parses the byte sequence to identify the syscall.

2. **Trie Lookup**:
   - The kernel traverses the embedded trie, matching the byte sequence to a specific syscall node.
   - If found, the kernel processes the syscall and sends the response back to the module via the specified port.

### Handling Multiple Tries

The kernel can manage multiple tries, recursively traversing them until a match is found. Each trie corresponds to different contexts or modules, ensuring scalability and modularity.

## Examples

### Syscall for Reading a File

1. **Syscall Byte Sequence**: `0x01[filename]`
2. **Kernel Trie Lookup**: Matches `0x01` to `file_read`.
3. **Process**: Kernel reads the file and sends the contents back to the requesting module.

```
Trie:
(root) - 0x01 (file_read)
           |
         filename (dynamic key)
```

### Syscall for Writing to a File

1. **Syscall Byte Sequence**: `0x02[timestamp][filename][data]`
2. **Kernel Trie Lookup**: Matches `0x02` to `file_write`.
3. **Process**:
   - Kernel checks the timestamp to ensure the correct order.
   - Writes the data to the file.
   - Sends an acknowledgment back to the requesting module.

```
Trie:
(root) - 0x02 (file_write)
           |
      timestamp (dynamic key)
           |
        filename (dynamic key)
           |
         data (dynamic key)
```

## Conclusion

In PromiseGrid, treating syscalls as byte sequences and utilizing a trie-based structure for efficient lookup provides a dynamic and effective way to handle syscalls. Including timestamps for operations with side effects ensures consistency, while using stdout for communication enables a streamlined message-passing mechanism. This approach enhances modularity, scalability, and reliability in syscall processing within the PromiseGrid framework.

### Open Questions

1. **How can we further optimize trie traversal for high-throughput environments?**
2. **What additional security measures can be introduced to protect against tampering or replay attacks in syscall sequences?**
3. **How can we ensure compatibility between different versions of modules that may define syscalls differently?**

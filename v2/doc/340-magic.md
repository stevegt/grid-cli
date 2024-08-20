# Design Discussion on Sequence Matching for Message Handling

## Introduction

In the context of designing robust message handling mechanisms, an alternative approach involves sequence matching on the message data, byte by byte, to decide how to handle the message. This method diverges from traditional parsing methods and bears similarity to the UNIX/Linux magic number system, which uses the first few bytes of a file to determine its type.

This document discusses the pros and cons of implementing such a sequence matching approach in PromiseGrid.

## Design Overview

### Sequence Matching Mechanism

Rather than parsing an incoming message, the proposed design involves performing sequence matching on the message data, byte by byte. This approach uses the initial bytes of the message to decide how to handle it, potentially invoking various decoding mechanisms such as multibase/multihash decoding or other future decoding techniques.

This design is akin to the UNIX/Linux magic number system, defined by `file` and `/etc/magic`, which identifies file types based on their initial bytes.

Byte sequences might overlap ambiguously. For example, similar initial bytes could represent different message types. Ambiguous sequence patterns are resolved by the first match. The ordering of first match is node-specific; the local node's own configuration determines the order of tries.

Complexity in managing a large set of sequence patterns is mitigated by using the decentralized cache to store and retrieve patterns efficiently.  Cache keys are nested and are each one byte in size.  Each byte is "owned" (registered) by one or more modules (handlers).  As a message is processed, each successive byte of the message is looked up as the next, nested, key.  This process proceeds until a lookup fails; on failure, the entire message is routed to the module that owns the most recent byte key that succeeded.  In the case of multiple modules owning the same key, the local node SHALL choose: it MAY route to the first module in the list, or it MAY route to all, or some other heuristic.  XXX revisit this

Handling after sequence matching would include any handler's own validation and error detection.  On failure of a handler, the message SHOULD be routed to the next handler in the list.  On failure of all handlers, the message MAY be dropped, and the router (dispatcher) SHOULD generate a new message to register the apparent spam (broken promise?).  XXX revisit this

### Pros of Sequence Matching for Message Handling

1. **Simplicity and Speed**:
    - **Byte-by-Byte Matching**: By directly matching byte sequences, this approach simplifies the decision-making process, potentially speeding up the message handling by avoiding complex parsing.
    - **Efficiency**: For known sequences, matching bytes may be quicker than interpreting structured data, leading to faster dispatch of messages.

2. **Flexibility**:
    - **Dynamic Adaptability**: The system can accommodate new types of messages and decodings by simply updating the sequence match patterns without altering the core parsing logic.
    - **Compatibility**: Supports various encoding schemes (multibase, multihash, etc.) by recognizing patterns directly from the byte stream.

3. **Modular Handling**:
    - **Future-Proof**: Easier to integrate future decoding mechanisms by adding new sequence patterns.
    - **Universal Matching**: Allows for a uniform way to handle diverse message formats and structures, enhancing modularity.

4. **Deterministic Execution**:
    - **Content-Based Dispatch**: Using the initial bytes for dispatch ensures that messages are routed deterministically, reducing ambiguity in handling.

### Cons of Sequence Matching for Message Handling

2. **Security Implications**:
    - **Tampering Risk**: Matching based on bytes might expose the system to tampering if attackers craft byte sequences to deceive the match routines, potentially leading to security vulnerabilities.
    - **Insufficient Validation**: Without robust parsing, the system might skip essential validation steps, increasing the risk of processing malformed or malicious messages.

3. **Performance Overhead**:
    - **Sequence Database Management**: Maintaining and searching a large database of sequence patterns could introduce performance overhead, particularly as the number of patterns grows.
    - **Resource Intensive**: Repeated sequence matching might consume more resources compared to directly parsing structured messages with known formats.

### Conclusion

The sequence matching approach offers significant flexibility and potential performance advantages for handling diverse message types in PromiseGrid. However, the design must carefully address the complexity, ambiguity, and security implications of this method. Future implementations can leverage pattern databases and robust validation mechanisms to mitigate these challenges.

Overall, while byte-by-byte sequence matching is a promising method for dynamic message handling, it requires careful design considerations to ensure reliability, security, and efficiency in a production environment.

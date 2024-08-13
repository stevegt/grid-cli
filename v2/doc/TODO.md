# TODO List for Implementing PromiseGrid Kernel

## Message Structure DONE
- DONE Define a `Message` structure where the promise is the first element in `Parms`.
  - DONE Prerequisite: None

- Ensure that the message format allows for the leading hash to be used as a promise hash, followed by the module hash and arguments.
  - DONE Prerequisite: Define a `Message` structure

## Cache Design and Lookup
- Ensure that the cache uses filesystem separators (`/`) between each key component and URL-encodes the arguments.
  - Prerequisite: Define a `Message` structure

- Treat modules as caches and consult them on cache misses, potentially using multiple caches including the built-in kernel cache and module-provided caches.
  - Prerequisite: Define a `Message` structure, Implement Hierarchical Syscall Tree

- Implement a complex filter algorithm to route messages based on promise acceptance, module hash match, and argument acceptance.
  - Prerequisite: None

## Hierarchical Syscall Tree
- Implement a hierarchical syscall tree to cache successful paths, optimizing future routing.
  - Prerequisite: Cache Design and Lookup

- Incorporate promise-based acceptance in the syscall tree to enhance dynamic routing and module selection.
  - Prerequisite: Implement Ant Routing Mechanism, Cache Design and Lookup

## Implement Ant Routing Mechanism
- Create an ant routing mechanism to cache the path to modules that have successfully handled previous calls.
  - Prerequisite: Define a `Message` structure

- Use the syscall tree to route messages to the module whose `syscallTable` key matches the most leading `parms` components.
  - Prerequisite: Implement Hierarchical Syscall Tree

## Dynamic Acceptance History and Syscall Table
- Merge `acceptanceHist` and `syscallTable` into a single dynamic syscall table.
  - Prerequisite: Hierarchical Syscall Tree

- Populate the dynamic syscall table during kernel operation, starting with built-in modules and then consulting other modules as needed.
  - Prerequisite: Dynamic Acceptance History and Syscall Table

- Store positive and negative acceptance history for all modules to optimize future lookups.
  - Prerequisite: Dynamic Acceptance History and Syscall Table

## Promise-Based Design
- Deepen the "promises all the way down" design by treating acceptance as a promise message.
  - Prerequisite: None

- Explore the possibility of using the `Accept()` response itself as a promise message, with meta-information included in `Payload`.
  - Prerequisite: Promise-Based Design

- Ensure that the kernel includes known accept/reject messages in its embedded data for validation and recognition.
  - Prerequisite: None

## Error Handling and Promise Fulfillment
- Log and handle broken promises when `HandleMessage()` fails after `Accept()` returned true.
  - Prerequisite: Explore the possibility of using the `Accept()` response as a promise message

- Develop mechanisms to reroute requests or de-prioritize unreliable modules to maintain system resilience and trust.
  - Prerequisite: Log and handle broken promises

- Track and validate acceptance promises to ensure accountability and trustworthiness within the system.
  - Prerequisite: Ensure that the kernel includes known accept/reject messages in its embedded data for validation and recognition

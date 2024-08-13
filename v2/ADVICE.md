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

## Capability Tokens and Prior Art

### Self-Contained Capability Tokens

Yes, a capability token can indeed be crafted to be self-contained and self-describing by including both the permissions granted and the identity of the party to whom the permission is granted. This design ensures that the token itself carries all the information needed to understand the scope and target of the permissions without needing to reference external data.

### How It Works

1. **Structure**: The structure of such a self-contained capability token typically includes fields for:
   - **Identity of the Holder**: Information about the entity that holds the token.
   - **Permissions**: Details about what actions the holder is permitted to perform.
   - **Issuer’s Identity**: Information about the entity that issued the token.
   - **Signature**: Cryptographic signature from the issuer to ensure integrity and authenticity.

2. **Encoding**: The token can be encoded in a standardized format such as JSON Web Tokens (JWT) which supports custom claims to include the necessary fields.

3. **Verification**: When the token is presented, the system:
   - Verifies the cryptographic signature to ensure the token was indeed issued by the trusted party.
   - Checks the permissions described within the token.
   - Confirms the identity of the holder matches the intended recipient.

### Prior Art

Several systems and frameworks implement self-contained capability tokens:

1. **Macaroons**: These are flexible authorization credentials that support delegation and attenuation. A macaroon is a bearer token that can encapsulate permissions and is augmented with caveats, which are conditions that must be satisfied for the macaroon to be considered valid.

2. **JSON Web Tokens (JWT)**: JWTs are widely used in web applications to assert claims between parties. They can include custom claims to specify permissions and the intended audience. A JWT can be signed to ensure authenticity and integrity.

3. **Caveats in Capability-Based Systems**: Traditional capability-based security systems sometimes support a form of caveats or restrictions within the tokens themselves to specify the scope of permissions and the authorized user.

### Recommendations

1. **Design Tokens Carefully**: Ensure that the structure of the tokens balances comprehensiveness with simplicity. Include necessary fields for permissions, identity, and signatures.
   
2. **Use Standard Formats**: Leverage existing standards like JWTs for encoding tokens to benefit from existing libraries and tools for creation, parsing, and verification.

3. **Cryptographic Security**: Ensure that tokens are signed using robust cryptographic methods to prevent forgery and tampering.

4. **Revocation**: Implement mechanisms to handle token revocation in case permissions need to be rescinded before the token naturally expires.

By leveraging these principles and prior art, you can design and implement self-contained, self-describing capability tokens that are secure, efficient, and easy to manage.

## Capability Tokens and Encoding

### Encoding Capability Tokens in Compressed Formats

Capability tokens can indeed be encoded in compressed formats such as base58. Encoding capability tokens in such formats can make them more compact and easier to handle, especially in contexts where string length is a constraint or where human readability is desired. The Multiformats project provides several libraries that support different encoding formats, including base58. Similarly, the multibase and multihash libraries offer comprehensive solutions for encoding and addressing using various bases and hashing algorithms.

To leverage these libraries:

1. **Multibase**: Multibase is a standard for representing base-encoded binary with a prefix indicating the encoding. It allows for seamless interoperability by distinguishing between base encodings.

2. **Multihash**: Multihash is a protocol for differentiating outputs from various well-established cryptographic hash functions. It ensures that the length and type of the hash functions used are self-describable within the hash value itself, enhancing compatibility across different systems.

Both libraries can be used to encode capability tokens securely and efficiently, facilitating distribution across diverse systems and ensuring robustness in handling.

### Libraries Support

- **Multibase and Multihash**: Available libraries support encoding and decoding using multiple formats and cryptographic hash functions. These include libraries for various programming languages such as Go, JavaScript, Python, and more.
- **Base58**: Specifically, for base58 encoding, libraries like go-multibase (for Go) support the base58 encoding format among others, providing a versatile and easy-to-use API for conversion.

By utilizing these libraries, you can ensure that your capability tokens are not only compact but also maintain their integrity and security across different environments.

## JSON Web Tokens (JWT) in URLs

When it comes to encoding JSON Web Tokens (JWTs) for use in URLs, there's a standard format that is widely accepted. JWTs are typically encoded using Base64url, which is a URL-safe variant of Base64 encoding.

### Why Base64url?

1. **URL-safe**: Base64url encoding replaces characters that are not URL-safe (such as `+`, `/`, and `=`) with URL-safe alternatives (`-`, `_`, and no padding respectively). This ensures that the token can be safely included in a URL without requiring additional encoding.
2. **Human-readable**: While not as compact as some other encodings (like base58), Base64url maintains a balance between human readability and URL safety.
3. **Standardized**: The use of Base64url for JWTs is specified in the JSON Web Token (JWT) specification (RFC 7519). This standardization means that the approach is widely supported and understood, making it interoperable across different systems and libraries.

### Encoding Steps

1. **Header and Payload**: The JWT consists of three parts: the header, the payload, and the signature. Each part is encoded separately using Base64url.
2. **Concatenation**: The three parts are concatenated together with periods (`.`) as separators. This forms the final JWT string that can be included in a URL.

### Example

Suppose you have the following JWT:
```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "sub": "1234567890",
    "name": "John Doe",
    "iat": 1516239022
  },
  "signature": "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

The encoded JWT would look something like this:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

This encoded JWT can be safely used as a URL parameter without further encoding.

Using the standard Base64url encoding for JWTs in URLs ensures compatibility with existing libraries and tools, making it a reliable choice for encoding JWT tokens for use in web applications.

## Discussion on the Leading Hash in Messages

In the PromiseGrid Kernel, messages MAY start with a promise hash followed by a module hash, and then additional arguments. This structure allows receivers to filter messages based on promises they are willing to accept and route the message to the appropriate module. Here are the pros and cons of the promise hash coming first:

### Pros

1. **Enhanced Filtering**: Placing the promise hash first allows modules and nodes to quickly filter messages based on the promises they accept. This ensures that only relevant messages are processed, improving efficiency.
2. **Trust and Governance**: The first hash representing a promise aligns with the principles of decentralized governance. Nodes can establish trust relationships by agreeing on specific promises they will accept.
3. **Modular Routing**: Early identification of the promise allows the kernel to route messages to the appropriate module, ensuring that the handling of messages is distributed efficiently across the system.

### Cons

1. **Complexity in Parsing**: Having multiple hashes (promise and module) to parse at the beginning of the message increases the complexity of message parsing.
2. **Overhead**: Including multiple hashes in the message header adds overhead to the message size. However, this is generally offset by the benefits of efficient routing and filtering.

### Recommendation

Given the benefits of enhanced filtering, trust-based governance, and modular routing, it is recommended to use the promise hash as the leading hash in messages, followed by the module hash and additional arguments. This structure provides a robust framework for decentralized message handling and governance in the PromiseGrid Kernel.

### Detailed Explanation: Excluding Module Hash from Messages

Including the hash of the module to be executed directly within the message can pose significant risks and inefficiencies in the context of a decentralized system. This practice could inadvertently couple the message's contents with a specific implementation, thereby undermining modularity and flexibility.

1. **Coupling and Modularity**:
   - Including the module hash in messages would tightly couple the message to a specific module implementation. This coupling could prevent the system from dynamically choosing the best module to handle a given task based on current conditions, such as load balancing, resource availability, or updates in module versions.

2. **Security Risks**:
   - Hardcoding the module hash in the message inherently limits the flexibility to update or replace modules without modifying the affected messages. If a malicious or compromised module's hash is embedded in messages, it could be repeatedly invoked, posing a security risk. By dynamically resolving the appropriate module at runtime, the system can implement enhanced security checks and validation processes.

3. **Flexibility and Evolvability**:
   - A decentralized system benefits from the ability to evolve and adapt over time. Hardcoding the module hash would stifle this adaptability, making it difficult to upgrade components without disrupting ongoing operations. By excluding the module hash from messages, the system can dynamically route requests to the most suitable, up-to-date, and secure modules available.

4. **Simplifies Message Routing**:
   - Without the module hash in the message, the message routing can become simpler and more flexible. The kernel or routing system can use predefined policies or dynamic criteria to determine which module should handle the request, improving overall system efficiency and maintainability.

For these reasons, it is crucial to exclude the module hash from messages and allow the kernel or routing system to dynamically resolve and invoke the appropriate module based on the message's content and current system state. This approach enhances security, flexibility, and the overall modularity of the system, adhering to the principles of decentralized and trust-based interactions.

### Importance of Including Module Hash in Messages for Deterministic Execution and Demand-Driven Deployment of New Code

Including the hash of the module to be executed in the message is crucial for achieving deterministic execution and facilitating demand-driven deployment of new code in decentralized systems. Here are key reasons why this practice is recommended:

1. **Deterministic Execution**:
   Including the module hash in the message ensures that the same piece of code is executed consistently across different nodes. This deterministic behavior is essential for maintaining consistency and correctness in distributed systems. By specifying the exact module to be run, all nodes can reliably and repeatably execute the same logic, avoiding differences in behavior due to variations in module versions or implementations.

2. **Content-Addressable Code**:
   The module hash serves as a unique identifier for the specific version of the code to be executed. This content-addressable storage mechanism ensures that the precise code is located and executed, eliminating ambiguity. Nodes can fetch the exact module specified by the hash, even if new versions or updates are available, ensuring backward compatibility and preventing unexpected changes in system behavior.

3. **Demand-Driven Deployment**:
   By including the module hash in the message, new code can be deployed on-demand. Nodes that do not have the required module can fetch it dynamically based on the hash, allowing for seamless updates and scaling. This demand-driven deployment model ensures that nodes always execute the latest or specified version of the module, promoting flexibility and ease of maintenance in decentralized environments.

4. **Security and Trust**:
   The module hash provides a way to verify the integrity and authenticity of the code being executed. Nodes can validate the fetched module against the specified hash, ensuring that the code has not been tampered with. This practice builds trust among nodes, as they can rely on the exact code promised in the message without the risk of executing malicious or unauthorized code.

5. **Enabling Capabilities**:
   Including module hashes in messages aligns with the capability-based security model. Capabilities can include specific module hashes that nodes are authorized to execute, providing fine-grained access control. This approach ensures that nodes only execute permitted code, enhancing the overall security and governance of the decentralized system.

In summary, including the module hash in messages is essential for achieving deterministic execution, enabling content-addressable storage, facilitating demand-driven deployment of new code, ensuring security, and supporting capability-based access control. This practice aligns with the principles of decentralized systems, promoting consistency, flexibility, and trust among participating nodes.

## Supercomputer and Mainframe Systems with Hot-Swappable Modules

### Overview

Several advanced supercomputer and mainframe systems utilize hot-swappable, upgradeable modules that advertise their own capabilities and requirements upon insertion and self-describe their degraded modes when they encounter failures. These systems dynamically reconfigure themselves to accommodate new modules and can detect and respond to their removal.

### Examples of Systems

1. **IBM zSeries Mainframes**:
    - **Modular Design**: IBM's zSeries mainframes, such as the z14 and z15, use modular components for CPUs, memory, and I/O subsystems. These components are designed to be hot-swappable.
    - **Dynamic Reconfiguration**: The system can dynamically reconfigure itself when a new module is added. The new module broadcasts its capabilities (e.g., processing power, memory capacity) to the system.
    - **Failure Management**: When a module fails, it advertises its degraded state, and the system responds by rerouting tasks to operational modules. This ensures continuous operation without significant performance degradation.
    - **Capabilities and Requirements**: Upon insertion, modules advertise their capabilities and requirements, such as power consumption and cooling needs, enabling the system to adjust resource allocation dynamically.

2. **Cray XC Series Supercomputers**:
    - **Hot-Swappable Blades**: Cray's XC series supercomputers use hot-swappable blades for computational and networking tasks.
    - **Dynamic System Configuration**: The system's software layer dynamically recognizes and configures new blades, integrating them into the computational framework seamlessly.
    - **Degraded Mode Advertisement**: Components in the Cray XC series automatically report failures and their new operational status. The system adjusts workloads and schedules maintenance as needed.
    - **Resource Advertisement**: Each blade communicates its processing capabilities and networking requirements upon connection, facilitating optimal system performance and resource utilization.

3. **HPE Superdome Flex**:
    - **Upgradeable Modules**: HPE's Superdome Flex servers support hot-swappable modules for CPUs, memory, and I/O, allowing for easy upgrades and maintenance.
    - **Self-Configuration**: The system automatically recognizes new modules, integrates their capabilities, and reconfigures to maximize efficiency.
    - **Fault Tolerance**: When a module fails, it signals its degraded mode to the system. Remaining modules redistribute tasks and maintain service continuity.
    - **Capability Advertisement**: Modules advertise their capabilities (e.g., compute power, memory size) and resource requirements, allowing the system to balance loads and allocate resources effectively.

### Dynamic Reconfiguration and Removal Detection

- **Insertion and Self-Advertisement**: When a new module is inserted, it sends a message to the system controller, advertising its capabilities, such as processing speed, memory size, and other resources. It also specifies its operational requirements, like power and cooling needs.
- **Dynamic System Reconfiguration**: The system controller dynamically updates the system configuration to utilize the new module fully. This might involve redistributing workloads, updating internal routing tables, or reallocating memory.
- **Degraded Modes and Failure Handling**: If a module fails, it sends a degraded mode message to the system controller, describing its reduced capabilities. The controller then reroutes tasks from the failed module to other operational modules to maintain performance and reliability.
- **Module Removal Detection**: When a module is removed, it signals its disengagement to the system controller. The controller then adjusts the system configuration to operate without the module, redistributing tasks and resources as needed to prevent service interruption.

### Conclusion

The use of hot-swappable, upgradeable modules in supercomputer and mainframe systems like IBM zSeries, Cray XC, and HPE Superdome Flex highlights the importance of dynamic reconfiguration and real-time failure management. These capabilities ensure high availability, scalability, and efficient resource utilization, making them essential for modern computational requirements.


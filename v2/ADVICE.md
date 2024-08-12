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

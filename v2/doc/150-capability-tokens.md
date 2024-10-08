## Capability Tokens and Prior Art

### Self-Contained Capability Tokens

A capability token can be crafted to be self-contained and self-describing by including both the permissions granted and the identity of the party to whom the permission is granted. This design ensures that the token itself carries all the information needed to understand the scope and target of the permissions without needing to reference external data.

### How It Works

1. **Structure**: The structure of such a self-contained capability token typically includes fields for:
   - **Identity of the Holder**: Information about the entity that holds the token.
   - **Permissions**: Details about what actions the holder is permitted to perform.
   - **Issuer’s Identity**: Information about the entity that issued the token.
   - **Signature**: Cryptographic signature from the issuer to ensure integrity and authenticity.

2. **Encoding**: The token can be encoded in a standardized format such as JSON Web Tokens (JWT), which supports custom claims to include the necessary fields.

3. **Verification**: When the token is presented, the system:
   - Verifies the cryptographic signature to ensure the token was indeed issued by the trusted party.
   - Checks the permissions described within the token.
   - Confirms the identity of the holder matches the intended recipient.

Several systems and frameworks implement self-contained capability tokens:

1. **Macaroons**: These are flexible authorization credentials that support delegation and attenuation. A macaroon is a bearer token that can encapsulate permissions and is augmented with caveats, which are conditions that must be satisfied for the macaroon to be considered valid.
2. **JSON Web Tokens (JWT)**: JWTs are widely used in web applications to assert claims between parties. They can include custom claims to specify permissions and the intended audience. A JWT can be signed to ensure authenticity and integrity.
3. **Caveats in Capability-Based Systems**: Traditional capability-based security systems sometimes support a form of caveats or restrictions within the tokens themselves to specify the scope of permissions and the authorized user.

### Recommendations

1. **Design Tokens Carefully**: Ensure that the structure of the tokens balances comprehensiveness with simplicity. Include necessary fields for permissions, identity, and signatures.
   
2. **Use Standard Formats**: Leverage existing standards like JWTs for encoding tokens to benefit from existing libraries and tools for creation, parsing, and verification.

3. **Cryptographic Security**: Ensure that tokens are signed using robust cryptographic methods to prevent forgery and tampering.

4. **Revocation**: Implement mechanisms to handle token revocation in case permissions need to be rescinded before the token naturally expires.

### JSON Web Tokens (JWT)

JWTs are typically encoded using Base64url, which is a URL-safe variant of Base64 encoding.

#### Why Base64url?

1. **URL-safe**: Base64url encoding replaces characters that are not URL-safe (such as `+`, `/`, and `=`) with URL-safe alternatives (`-`, `_`, and no padding). This ensures that the token can be safely included in a URL without requiring additional encoding.
2. **Human-readable**: While not as compact as some other encodings (like base58), Base64url maintains a balance between human readability and URL safety.
3. **Standardized**: The use of Base64url for JWTs is specified in the JSON Web Token (JWT) specification (RFC 7519). This standardization means that the approach is widely supported and understood, making it interoperable across different systems and libraries.

#### Encoding Steps

1. **Header and Payload**: The JWT consists of three parts: the header, the payload, and the signature. Each part is encoded separately using Base64url.
2. **Concatenation**: The three parts are concatenated together with periods (`.`) as separators. This forms the final JWT string that can be included in a URL.

#### Example

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

### Discussion on the Leading Hash in Messages

In the PromiseGrid Kernel, messages MAY start with a promise hash followed by a module hash and then additional arguments. This structure allows receivers to filter messages based on promises they are willing to accept and route the message to the appropriate module. Here are the pros and cons of the promise hash coming first:

**Pros:**
- **Enhanced Filtering**: Placing the promise hash first allows modules and nodes to quickly filter messages based on the promises they accept.
- **Trust and Governance**: This aligns with decentralized governance. Nodes can establish trust relationships by agreeing on specific promises they will accept.
- **Modular Routing**: Early identification of the promise allows the kernel to route messages to the appropriate module, ensuring efficient distributed handling.

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

### Kernel's Dynamic Handling of Accept and HandleMessage

Implementing an advanced interaction model involves considering different versions of function handling of messages. There are two possible structures: combined functions or separate functions.

### Combined Function:

```go
type Module interface {
    ProcessMessage(ctx context.Context, parms ...interface{}) (Message, error)
}

type LocalCacheModule struct {
    cacheDir string
}

func NewLocalCacheModule(cacheDir string) *LocalCacheModule {
    return &LocalCacheModule{cacheDir: cacheDir}
}

func (m *LocalCacheModule) ProcessMessage(ctx context.Context, parms ...interface{}) (Message, error) {
    // Implement logic to process message, which includes accepting and handling
    // Return a promise message with acceptance and handling results or errors
}
```

### Separate Functions:

```go
type Module interface {
    Accept(ctx context.Context, parms ...interface{}) (Message, error)
    HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error)
}

type LocalCacheModule struct {
    cacheDir string
}

func NewLocalCacheModule(cacheDir string) *LocalCacheModule {
    return &LocalCacheModule{cacheDir: cacheDir}
}

func (m *LocalCacheModule) Accept(ctx context.Context, parms ...interface{}) (Message, error) {
   // Implement logic for accepting or rejecting based on parameters
}

func (m *LocalCacheModule) HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error) {
    // Implement logic for handling the message after acceptance
}
```

### Conclusion

By integrating promises at every level and implementing a hierarchical syscall tree with caching and acceptance history, PromiseGrid ensures trust, accountability, and efficient message handling in a decentralized governance framework. The simplified message structure and consistent handling of cache and modules contribute to a robust and flexible system.

### Open Questions

1. **Handling Broken Promises**: What strategies can we implement to handle broken promises effectively? How can we ensure that nodes that consistently break promises have their reputation impacted in a measurable way?
3. **Integration with Other Systems**: How can we better integrate JSON Web Tokens (JWT) and similar systems with PromiseGrid’s architecture? What role should JWT play in managing capabilities within the grid?
4. **Security Measures**: Given the reliance on module hashes for deterministic execution, what additional security measures can enhance the integrity and trustworthiness of the system? How can we mitigate potential risks associated with hash collisions and tampering? 

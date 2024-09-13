# PromiseGrid Layer Structure in an OSI-like Model

PromiseGrid's architecture can be understood through a layered
approach, similar to the OSI Model used in traditional networking.
Here, we outline the key layers and their functions within the
PromiseGrid system.

## Layer Structure

### Host Layer

This layer includes the physical or virtual hardware running
PromiseGrid. It represents the fundamental resources such as storage,
memory, network, and processing units that support the operation of
the system.

#### Provides 
  - persistent storage (disk)
  - transient storage (memory)
  - processing units (CPU, GPU)
  - data transport (network, internal buses)

#### Providers
  - underlying operating system running on a physical or virtual machine
  - container runtime like Docker
  - WASM/WASI runtime in browser or standalone

Examples of host layer APIs include the Linux syscall interface or
WASI syscalls in WebAssembly.


### Kernel Layer

This layer acts as an abstraction between the host layer and
PromiseGrid modules. It provides a standardized interface for modules
to interact with the underlying hardware and resources.

#### Provides
  - standardized message-based APIs for:
    - disk I/O
    - network I/O
    - process control 
    - module control 

#### Providers
  - PromiseGrid Kernel 

The kernel is customized for the underlying host environment and is
kept minimal to reduce porting effort and maintain compatibility
across different host environments.  It uses a microkernel-style
architecture, where any code that is not specific to the host API is
located in builtin modules that compose the module runtime.

All communication with the kernel is done through message passing.


### Completion Layer

This layer is the common, decentralized, Turing-complete computing
platform that PromiseGrid modules run on.  The only operation that 
this layer supports is the completion of a given byte sequence.  

#### Provides
  - byte sequence completion services

#### Providers
  - PromiseGrid Kernel
  - PromiseGrid Modules (via the kernel)

All communication with the completion layer is done through message
passing; a request message contains a byte sequence to be completed,
and the response message contains the completed byte sequence.

Messages are typically passed as handles or pointers to a software
object that supports read() and write() operations.  The message
sender calls write() to add bytes to a message, and the message
receiver calls read() to read bytes from the message.

XXX


### Module runtime

This layer hosts PromiseGrid modules and provides the runtime
environment for their execution. It manages the lifecycle of modules,
including loading, unloading, and communication between modules.

#### Provides
  - module lifecycle management
  - message routing between modules
  - resource allocation and management

#### Providers
  - PromiseGrid Kernel builtin modules

The module runtime is composed of a set of builtin modules that
ship with the PromiseGrid kernel.  

All communication with the runtime is done through message passing.

  
### Module Layer

This layer represents the user-space modules that run on top of the PromiseGrid runtime. These modules provide various services and applications within the PromiseGrid framework.

#### Provides
  - user-facing services
  - application-specific functionality

#### Providers
  - user-developed modules


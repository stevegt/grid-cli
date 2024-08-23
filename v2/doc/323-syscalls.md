# PromiseGrid Syscall Mechanism

## Overview

In PromiseGrid, syscalls are implemented via a message-passing mechanism, where a syscall request is sent on a specific port allocated to the module at load time. Each module is provided with standard ports equivalent to stdin, stdout, and stderr. This document outlines the design and handling of syscalls in PromiseGrid.

## Syscall Port Allocation

### Port Types

At the time of module load, PromiseGrid allocates the following standard ports to each module:

- **stdin (Standard Input Port)**: Used for receiving input messages.
- **stdout (Standard Output Port)**: Used for sending output messages.
- **stderr (Standard Error Port)**: Used for sending error messages and logs.

### Dynamic Allocation

Modules can also dynamically request additional ports as needed. These ports are granted by the kernel based on resource availability and permissions.

## Message-Passing Mechanism

### Sending a Syscall Request

To perform a syscall, a module sends a syscall request message on the appropriate port. The message includes the following:

1. **Syscall Type**: Denotes the type of syscall being requested (e.g., file access, network communication).
2. **Params**: Arguments or parameters required for the syscall.
3. **Promise Hash**: Identifies the promise being fulfilled or action being taken.

### Receiving and Handling Syscall Requests

The kernel listens on the allocated ports for syscall requests from modules. Upon receiving a request, the kernel performs the following steps:

1. **Validation**: The kernel validates the syscall request, checking the permissions and ensuring the message is correctly formatted.
2. **Execution**: The kernel executes the requested syscall, interacting with necessary resources or services.
3. **Response**: After execution, the kernel sends a response message back to the module on the appropriate port (stdout for success, stderr for errors).

### Example Syscall Flow

1. **Module Initialization**:
    - The module is loaded and standard ports (stdin, stdout, stderr) are allocated.

2. **Send Syscall Request**:
    - The module sends a syscall request message on the stdin port:
    ```json
    {
      "syscall_type": "file_read",
      "params": { "filename": "/example/path" },
      "promise_hash": "0xabc123"
    }
    ```

3. **Kernel Receives Request**:
    - The kernel receives and validates the request on the stdin port.
    - The kernel performs the file read operation as requested.

4. **Send Response**:
    - The kernel sends the read data back to the module on the stdout port:
    ```json
    {
      "promise_hash": "0xabc123",
      "status": "success",
      "data": "file content..."
    }
    ```

### Error Handling

In case of an error during syscall execution, the kernel sends an error message back to the module on the stderr port:
```json
{
  "promise_hash": "0xabc123",
  "status": "error",
  "error_message": "File not found"
}
```

## Module Ports and Capabilities

### Standard Ports

Modules are allocated standard ports (stdin, stdout, stderr) for basic input, output, and error communication.

### Capability Tokens

Additional ports can be requested dynamically by modules using capability tokens. These tokens represent the permissions and resources allocated to the module.

### Dynamic Port Management

Modules can manage their ports dynamically, requesting and releasing ports as required by their operational needs. The kernel maintains control over the allocation and deallocation of these ports to ensure optimal resource usage and security.

## Summary

Syscalls in PromiseGrid are implemented using a message-passing mechanism where modules communicate with the kernel via allocated ports. This design ensures modularity, security, and flexibility by leveraging the standard ports and capability tokens for dynamic resource management. The kernel validates, executes, and responds to syscall requests, maintaining an efficient and secure system operation.

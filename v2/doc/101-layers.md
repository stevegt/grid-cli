# Message Layers

This document describes the message layering architecture for grid-cli.

## Framing

- **Format**: CBOR is used for framing messages, enabling efficient binary encoding.

## Message Structure

1. **Protocol Hash**: The first field in each CBOR message is a hash of the protocol document, ensuring integrity and version control.
2. **Protocol-Specific Fields**: Subsequent CBOR fields contain data specific to the selected protocol, allowing for flexible and extensible communication.


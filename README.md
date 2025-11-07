# ChaosRPC

ChaosRPC is a simple tool for reproducing transient errors on RPC (Remote Procedure Call).

## Concepts

- **Call Site**: A specific location in the code where an RPC call is made. It is composed of the call stack (including function names, discarding linenoes to preserve robustness between versions) and the RPC method name.

## Features

- **gRPC Support**: Unary and streaming RPC calls are supported.
- **HTTP Support**: Coming soon.

## Examples

Please refer to the `examples` directory.

- **Toy**: A simple toy service to demonstrate ChaosRPC.
- **SeaweedFS Issues**: Reproduce transient errors on RPC related to some issues.

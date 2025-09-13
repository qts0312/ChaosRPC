# Toy Service

This is toy service written in Go, which is simple but enough to show how to use ChaosRPC.

## Step

1. Download Protobuf compiler and compile the proto file.

    ```bash
    protoc --go-grpc_out=. --go_out=. toy.proto
    ```

2. Run the server and client respectively.

3. Inject faults with setting environment variables.

    ```bash
    export CHAOS_CALL_SITE="runtime.goexit:1223;runtime.main:283;main.main:32;github.com/qts0312/ChaosRPC/examples/toy/proto.(*toyServiceClient).Handshake:42;/toy.ToyService/Handshake"
    export CHAOS_ERROR_CODE=1
    ```

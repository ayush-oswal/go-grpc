# GRPC in Go Lang Demo :
This repository demonstrates how to implement gRPC in Go.

## Prerequisites
- Go (1.16 or later): [Official Go Installation Guide](https://golang.org/doc/install)
- Protocol Buffers (`protoc`): [Installation Guide](https://grpc.io/docs/protoc-installation/)


## Steps to Run Locally
 - Clone the repo.
 - Ensure you have the necessary dependencies.
 - Delete the existing chat folder from both directories and create a new chat folder.
 - Generate gRPC Code by navigating to the directory containing your chat.proto file and use protoc to generate Go code:
 ```
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative chat.proto

 ```
 - cd grpc-server and cd grpc-client and run go run main.go
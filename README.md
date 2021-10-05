Xem thêm ở [đây](https://grpc.io/docs/languages/go/quickstart/)

## 1. Init go mod

```sh
go mod init github.com/ducnguyen96/ducnguyen96.xyz-protos
```

## 2. Install Protobuf compiler

```sh
sudo apt install -y protobuf-compiler
```

## 3. Write proto

```proto
syntax = "proto3";

package protogen;
option go_package = "protogen/v1";

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

## 4. Gen proto

```sh
protoc --experimental_allow_proto3_optional \
  --go_out=. \
  --go-grpc_out=. \
  --go-grpc_opt=paths=import \
  --proto_path=. ./protos/**/*.proto ./protos/*.proto
```

## 5. Publish
```shell
git tag v1.0.4
git push origin v1.0.4
```

## 6. Get latest
```shell
go get -x github.com/ducnguyen96/ducnguyen96.xyz-protos@latest
```
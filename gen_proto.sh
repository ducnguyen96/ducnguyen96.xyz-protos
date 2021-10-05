#!/usr/bin/bash

protoc --experimental_allow_proto3_optional \
  --go_out=. \
  --go-grpc_out=. \
  --go-grpc_opt=paths=import \
  --proto_path=. ./protos/**/*.proto ./protos/*.proto
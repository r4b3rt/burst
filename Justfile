#!/usr/bin/env just --justfile

run:
   go run .

prepare:
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/bufbuild/buf/cmd/buf@latest

gen:
    buf generate

# call go mod tidy
tidy:
    go mod tidy

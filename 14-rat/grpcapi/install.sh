#!/bin/sh

#if using modules

#export GO111MODULE=on  # Enable module mode
#go get github.com/golang/protobuf/protoc-gen-go
#go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0

# else

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc

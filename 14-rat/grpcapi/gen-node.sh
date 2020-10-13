#!/bin/sh

#protoc --proto_path=. --js_out=import_style=commonjs,binary:../client-node implant.proto
protoc-gen-grpc \
--js_out=import_style=commonjs,binary:../client-node/src \
--grpc_out=../client-node/src \
--proto_path . \
./implant.proto

protoc-gen-grpc-ts \
--ts_out=service=true:../client-node/src \
--proto_path . \
./implant.proto

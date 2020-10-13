#/bin/sh
set -e
echo 'download protoc https://github.com/protocolbuffers/protobuf/releases/tag/v3.13.0'

protoc -I . implant.proto --go_out=plugins=grpc:./


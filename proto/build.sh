#!/usr/bin/env bash

set -e

echo "Generating protobuf for gRPC golang"
protoc -I./messenger --go_out=plugins=grpc:messenger messenger/messenger.proto
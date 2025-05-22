#!/bin/bash

protoc -Iproto \
       --go_out=pkg/api/wal --go_opt=paths=source_relative \
       --go-grpc_out=pkg/api/wal --go-grpc_opt=paths=source_relative \
       proto/wal.proto

protoc -Iproto \
       --go_out=pkg/api/user --go_opt=paths=source_relative \
       --go-grpc_out=pkg/api/user --go-grpc_opt=paths=source_relative \
       proto/user.proto

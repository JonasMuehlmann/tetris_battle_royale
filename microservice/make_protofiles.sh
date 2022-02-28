#!/usr/bin/env sh

SERVICES=("user_service" "game_service")

# Build all protofiles
for service in ${SERVICES[@]}; do
    protoc \
        -I internal/core/protofiles/ \
        --go_out=./internal/core/protofiles \
        --go_opt=paths=source_relative \
        --go-grpc_out=./internal/core/protofiles \
        --go-grpc_opt=paths=source_relative \
        internal/core/protofiles/${service}/${service}.proto
done

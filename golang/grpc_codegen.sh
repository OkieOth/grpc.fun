#!/bin/bash

scriptPos=${0%/*}

cd $scriptPos

protoc --proto_path=../models/grpc -I=. --go_out=. \
    --go-grpc_out=. \
    --experimental_allow_proto3_optional ../models/grpc/service.proto
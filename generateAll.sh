scriptPos=${0%/*}

cd $scriptPos

yacgVersion=6.3.0

# generate all the things from the business model
# if ! docker run -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/project --rm -t \
#     ghcr.io/okieoth/yacg:${yacgVersion} \
#      --config /project/codeGen/yacg_config.json; then
#     echo "error while run codegen, cancel"
#     exit 1
# fi

# generate golang code
# Install go plugins
protoc --proto_path=./models/grpc --go_out=./golang/generated/pkg/proto_defs \
    --go-grpc_out=./golang/generated/pkg/proto_defs \
    --experimental_allow_proto3_optional \
        ./models/grpc/service.proto \
        ./models/grpc/generated/types.proto


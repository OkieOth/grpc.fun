Source: https://grpc.io/docs/languages/go/quickstart/


```bash
# Install go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# take care the the plugins are found by protoc
export PATH="$PATH:$(go env GOPATH)/bin"


# generate code for the example
 protoc \                                  # compiler call
   --proto_path=../models/grpc \           # specify where to find the protofile
   -I=. \                                  # ???
   --go_out=. \                            # base path for the go output
   --go_opt=Mservice.proto=grpc_test \     # link the proto file to a package
   --experimental_allow_proto3_optional \  # allows optional values
   ../models/grpc/service.proto            # proto file to use



protoc --proto_path=../models/grpc -I=. --go_out=. --go_opt=Mservice.proto=.grpc_test --experimental_allow_proto3_optional ../models/grpc/service.proto

# load dependencies in the project
go get google.golang.org/protobuf
```
package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatalf("Can't bin to port 10001: %v", err)
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("can't start server: %v", err)
	}
}

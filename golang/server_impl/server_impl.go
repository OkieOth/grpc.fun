package server_impl

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type TestServer struct {
}

// Create operation
func (s *TestServer) CreateObject(context.Context, *Object) (*Object, error) {
	return status.Errorf(codes.Unimplemented, "TODO") // TODO
}

// Read operation
func (s *TestServer) ReadObjectById(context.Context, *ReadObjectRequest) (*Object, error) {
	return status.Errorf(codes.Unimplemented, "TODO") // TODO
}

// Update operation
func (s *TestServer) UpdateObject(context.Context, *Object) (*Object, error) {
	return status.Errorf(codes.Unimplemented, "TODO") // TODO
}

// Delete operation
func (s *TestServer) DeleteObject(context.Context, *DeleteObjectRequest) (*Object, error) {
	return status.Errorf(codes.Unimplemented, "TODO") // TODO
}

// Read a list of objects in a stream
func (s *TestServer) ListObjects(*ListObjectsRequest, ObjectService_ListObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "TODO") // TODO
}

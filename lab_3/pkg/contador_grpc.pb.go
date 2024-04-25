// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pkg

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContadorClient is the client API for Contador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContadorClient interface {
}

type contadorClient struct {
	cc grpc.ClientConnInterface
}

func NewContadorClient(cc grpc.ClientConnInterface) ContadorClient {
	return &contadorClient{cc}
}

// ContadorServer is the server API for Contador service.
// All implementations must embed UnimplementedContadorServer
// for forward compatibility
type ContadorServer interface {
	mustEmbedUnimplementedContadorServer()
}

// UnimplementedContadorServer must be embedded to have forward compatible implementations.
type UnimplementedContadorServer struct {
}

func (UnimplementedContadorServer) mustEmbedUnimplementedContadorServer() {}

// UnsafeContadorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContadorServer will
// result in compilation errors.
type UnsafeContadorServer interface {
	mustEmbedUnimplementedContadorServer()
}

func RegisterContadorServer(s grpc.ServiceRegistrar, srv ContadorServer) {
	s.RegisterService(&Contador_ServiceDesc, srv)
}

// Contador_ServiceDesc is the grpc.ServiceDesc for Contador service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contador_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contador.Contador",
	HandlerType: (*ContadorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "pkg/contador.proto",
}

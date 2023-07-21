// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: outofstoremessagetypes/outofstoremessage.proto

package outofstoremessagetypes

import (
	protocoltypes "berty.tech/weshnet/pkg/protocoltypes"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OutOfStoreMessageService_OutOfStoreReceive_FullMethodName = "/weshnet.outofstoremessage.v1.OutOfStoreMessageService/OutOfStoreReceive"
)

// OutOfStoreMessageServiceClient is the client API for OutOfStoreMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutOfStoreMessageServiceClient interface {
	// OutOfStoreReceive parses a payload received outside a synchronized store
	OutOfStoreReceive(ctx context.Context, in *protocoltypes.OutOfStoreReceive_Request, opts ...grpc.CallOption) (*protocoltypes.OutOfStoreReceive_Reply, error)
}

type outOfStoreMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOutOfStoreMessageServiceClient(cc grpc.ClientConnInterface) OutOfStoreMessageServiceClient {
	return &outOfStoreMessageServiceClient{cc}
}

func (c *outOfStoreMessageServiceClient) OutOfStoreReceive(ctx context.Context, in *protocoltypes.OutOfStoreReceive_Request, opts ...grpc.CallOption) (*protocoltypes.OutOfStoreReceive_Reply, error) {
	out := new(protocoltypes.OutOfStoreReceive_Reply)
	err := c.cc.Invoke(ctx, OutOfStoreMessageService_OutOfStoreReceive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OutOfStoreMessageServiceServer is the server API for OutOfStoreMessageService service.
// All implementations must embed UnimplementedOutOfStoreMessageServiceServer
// for forward compatibility
type OutOfStoreMessageServiceServer interface {
	// OutOfStoreReceive parses a payload received outside a synchronized store
	OutOfStoreReceive(context.Context, *protocoltypes.OutOfStoreReceive_Request) (*protocoltypes.OutOfStoreReceive_Reply, error)
	mustEmbedUnimplementedOutOfStoreMessageServiceServer()
}

// UnimplementedOutOfStoreMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOutOfStoreMessageServiceServer struct {
}

func (UnimplementedOutOfStoreMessageServiceServer) OutOfStoreReceive(context.Context, *protocoltypes.OutOfStoreReceive_Request) (*protocoltypes.OutOfStoreReceive_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OutOfStoreReceive not implemented")
}
func (UnimplementedOutOfStoreMessageServiceServer) mustEmbedUnimplementedOutOfStoreMessageServiceServer() {
}

// UnsafeOutOfStoreMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutOfStoreMessageServiceServer will
// result in compilation errors.
type UnsafeOutOfStoreMessageServiceServer interface {
	mustEmbedUnimplementedOutOfStoreMessageServiceServer()
}

func RegisterOutOfStoreMessageServiceServer(s grpc.ServiceRegistrar, srv OutOfStoreMessageServiceServer) {
	s.RegisterService(&OutOfStoreMessageService_ServiceDesc, srv)
}

func _OutOfStoreMessageService_OutOfStoreReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocoltypes.OutOfStoreReceive_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutOfStoreMessageServiceServer).OutOfStoreReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutOfStoreMessageService_OutOfStoreReceive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutOfStoreMessageServiceServer).OutOfStoreReceive(ctx, req.(*protocoltypes.OutOfStoreReceive_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OutOfStoreMessageService_ServiceDesc is the grpc.ServiceDesc for OutOfStoreMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OutOfStoreMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weshnet.outofstoremessage.v1.OutOfStoreMessageService",
	HandlerType: (*OutOfStoreMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OutOfStoreReceive",
			Handler:    _OutOfStoreMessageService_OutOfStoreReceive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "outofstoremessagetypes/outofstoremessage.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/statusmanager.proto

package spec

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StatusManagerServiceClient is the client API for StatusManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatusManagerServiceClient interface {
	RegisterStatus(ctx context.Context, in *RegisterStatusRequest, opts ...grpc.CallOption) (*RegisterStatusResponse, error)
}

type statusManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatusManagerServiceClient(cc grpc.ClientConnInterface) StatusManagerServiceClient {
	return &statusManagerServiceClient{cc}
}

func (c *statusManagerServiceClient) RegisterStatus(ctx context.Context, in *RegisterStatusRequest, opts ...grpc.CallOption) (*RegisterStatusResponse, error) {
	out := new(RegisterStatusResponse)
	err := c.cc.Invoke(ctx, "/StatusManagerService/RegisterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusManagerServiceServer is the server API for StatusManagerService service.
// All implementations must embed UnimplementedStatusManagerServiceServer
// for forward compatibility
type StatusManagerServiceServer interface {
	RegisterStatus(context.Context, *RegisterStatusRequest) (*RegisterStatusResponse, error)
	mustEmbedUnimplementedStatusManagerServiceServer()
}

// UnimplementedStatusManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatusManagerServiceServer struct {
}

func (UnimplementedStatusManagerServiceServer) RegisterStatus(context.Context, *RegisterStatusRequest) (*RegisterStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterStatus not implemented")
}
func (UnimplementedStatusManagerServiceServer) mustEmbedUnimplementedStatusManagerServiceServer() {}

// UnsafeStatusManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatusManagerServiceServer will
// result in compilation errors.
type UnsafeStatusManagerServiceServer interface {
	mustEmbedUnimplementedStatusManagerServiceServer()
}

func RegisterStatusManagerServiceServer(s grpc.ServiceRegistrar, srv StatusManagerServiceServer) {
	s.RegisterService(&StatusManagerService_ServiceDesc, srv)
}

func _StatusManagerService_RegisterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusManagerServiceServer).RegisterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StatusManagerService/RegisterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusManagerServiceServer).RegisterStatus(ctx, req.(*RegisterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatusManagerService_ServiceDesc is the grpc.ServiceDesc for StatusManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatusManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StatusManagerService",
	HandlerType: (*StatusManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterStatus",
			Handler:    _StatusManagerService_RegisterStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/statusmanager.proto",
}

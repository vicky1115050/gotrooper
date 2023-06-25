// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.5
// source: pb/shellService.proto

package pb

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

// ShellServiceClient is the client API for ShellService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShellServiceClient interface {
	GetScripts(ctx context.Context, in *ShellRequest, opts ...grpc.CallOption) (*ShellResponse, error)
	SendFragment(ctx context.Context, in *ShellFragmentRquest, opts ...grpc.CallOption) (*FragmentResponse, error)
}

type shellServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShellServiceClient(cc grpc.ClientConnInterface) ShellServiceClient {
	return &shellServiceClient{cc}
}

func (c *shellServiceClient) GetScripts(ctx context.Context, in *ShellRequest, opts ...grpc.CallOption) (*ShellResponse, error) {
	out := new(ShellResponse)
	err := c.cc.Invoke(ctx, "/ShellService/GetScripts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shellServiceClient) SendFragment(ctx context.Context, in *ShellFragmentRquest, opts ...grpc.CallOption) (*FragmentResponse, error) {
	out := new(FragmentResponse)
	err := c.cc.Invoke(ctx, "/ShellService/SendFragment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShellServiceServer is the server API for ShellService service.
// All implementations must embed UnimplementedShellServiceServer
// for forward compatibility
type ShellServiceServer interface {
	GetScripts(context.Context, *ShellRequest) (*ShellResponse, error)
	SendFragment(context.Context, *ShellFragmentRquest) (*FragmentResponse, error)
	mustEmbedUnimplementedShellServiceServer()
}

// UnimplementedShellServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShellServiceServer struct {
}

func (UnimplementedShellServiceServer) GetScripts(context.Context, *ShellRequest) (*ShellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScripts not implemented")
}
func (UnimplementedShellServiceServer) SendFragment(context.Context, *ShellFragmentRquest) (*FragmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFragment not implemented")
}
func (UnimplementedShellServiceServer) mustEmbedUnimplementedShellServiceServer() {}

// UnsafeShellServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShellServiceServer will
// result in compilation errors.
type UnsafeShellServiceServer interface {
	mustEmbedUnimplementedShellServiceServer()
}

func RegisterShellServiceServer(s grpc.ServiceRegistrar, srv ShellServiceServer) {
	s.RegisterService(&ShellService_ServiceDesc, srv)
}

func _ShellService_GetScripts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShellServiceServer).GetScripts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShellService/GetScripts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShellServiceServer).GetScripts(ctx, req.(*ShellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShellService_SendFragment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellFragmentRquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShellServiceServer).SendFragment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShellService/SendFragment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShellServiceServer).SendFragment(ctx, req.(*ShellFragmentRquest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShellService_ServiceDesc is the grpc.ServiceDesc for ShellService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShellService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ShellService",
	HandlerType: (*ShellServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScripts",
			Handler:    _ShellService_GetScripts_Handler,
		},
		{
			MethodName: "SendFragment",
			Handler:    _ShellService_SendFragment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/shellService.proto",
}

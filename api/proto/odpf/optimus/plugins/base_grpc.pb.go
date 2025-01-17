// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package optimus

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

// BaseClient is the client API for Base service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseClient interface {
	// PluginInfo provides basic details for this plugin
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error)
}

type baseClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseClient(cc grpc.ClientConnInterface) BaseClient {
	return &baseClient{cc}
}

func (c *baseClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error) {
	out := new(PluginInfoResponse)
	err := c.cc.Invoke(ctx, "/odpf.optimus.plugins.Base/PluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServer is the server API for Base service.
// All implementations must embed UnimplementedBaseServer
// for forward compatibility
type BaseServer interface {
	// PluginInfo provides basic details for this plugin
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error)
	mustEmbedUnimplementedBaseServer()
}

// UnimplementedBaseServer must be embedded to have forward compatible implementations.
type UnimplementedBaseServer struct {
}

func (UnimplementedBaseServer) PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginInfo not implemented")
}
func (UnimplementedBaseServer) mustEmbedUnimplementedBaseServer() {}

// UnsafeBaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServer will
// result in compilation errors.
type UnsafeBaseServer interface {
	mustEmbedUnimplementedBaseServer()
}

func RegisterBaseServer(s grpc.ServiceRegistrar, srv BaseServer) {
	s.RegisterService(&Base_ServiceDesc, srv)
}

func _Base_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/odpf.optimus.plugins.Base/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Base_ServiceDesc is the grpc.ServiceDesc for Base service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Base_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "odpf.optimus.plugins.Base",
	HandlerType: (*BaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginInfo",
			Handler:    _Base_PluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "odpf/optimus/plugins/base.proto",
}

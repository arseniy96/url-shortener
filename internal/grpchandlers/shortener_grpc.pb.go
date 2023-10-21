// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: src/proto/shortener.proto

package grpchandlers

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

const (
	ShortenerProto_Ping_FullMethodName             = "/grpchandlers.ShortenerProto/Ping"
	ShortenerProto_CreateLink_FullMethodName       = "/grpchandlers.ShortenerProto/CreateLink"
	ShortenerProto_CreateLinksBatch_FullMethodName = "/grpchandlers.ShortenerProto/CreateLinksBatch"
	ShortenerProto_ResolveLink_FullMethodName      = "/grpchandlers.ShortenerProto/ResolveLink"
	ShortenerProto_UserUrls_FullMethodName         = "/grpchandlers.ShortenerProto/UserUrls"
)

// ShortenerProtoClient is the client API for ShortenerProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerProtoClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error)
	CreateLinksBatch(ctx context.Context, in *CreateLinksBatchRequest, opts ...grpc.CallOption) (*CreateLinksBatchResponse, error)
	ResolveLink(ctx context.Context, in *ResolveLinkRequest, opts ...grpc.CallOption) (*ResolveLinkResponse, error)
	UserUrls(ctx context.Context, in *UserUrlsRequest, opts ...grpc.CallOption) (*UserUrlsResponse, error)
}

type shortenerProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerProtoClient(cc grpc.ClientConnInterface) ShortenerProtoClient {
	return &shortenerProtoClient{cc}
}

func (c *shortenerProtoClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, ShortenerProto_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerProtoClient) CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error) {
	out := new(CreateLinkResponse)
	err := c.cc.Invoke(ctx, ShortenerProto_CreateLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerProtoClient) CreateLinksBatch(ctx context.Context, in *CreateLinksBatchRequest, opts ...grpc.CallOption) (*CreateLinksBatchResponse, error) {
	out := new(CreateLinksBatchResponse)
	err := c.cc.Invoke(ctx, ShortenerProto_CreateLinksBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerProtoClient) ResolveLink(ctx context.Context, in *ResolveLinkRequest, opts ...grpc.CallOption) (*ResolveLinkResponse, error) {
	out := new(ResolveLinkResponse)
	err := c.cc.Invoke(ctx, ShortenerProto_ResolveLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerProtoClient) UserUrls(ctx context.Context, in *UserUrlsRequest, opts ...grpc.CallOption) (*UserUrlsResponse, error) {
	out := new(UserUrlsResponse)
	err := c.cc.Invoke(ctx, ShortenerProto_UserUrls_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerProtoServer is the server API for ShortenerProto service.
// All implementations must embed UnimplementedShortenerProtoServer
// for forward compatibility
type ShortenerProtoServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error)
	CreateLinksBatch(context.Context, *CreateLinksBatchRequest) (*CreateLinksBatchResponse, error)
	ResolveLink(context.Context, *ResolveLinkRequest) (*ResolveLinkResponse, error)
	UserUrls(context.Context, *UserUrlsRequest) (*UserUrlsResponse, error)
	mustEmbedUnimplementedShortenerProtoServer()
}

// UnimplementedShortenerProtoServer must be embedded to have forward compatible implementations.
type UnimplementedShortenerProtoServer struct {
}

func (UnimplementedShortenerProtoServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedShortenerProtoServer) CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (UnimplementedShortenerProtoServer) CreateLinksBatch(context.Context, *CreateLinksBatchRequest) (*CreateLinksBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLinksBatch not implemented")
}
func (UnimplementedShortenerProtoServer) ResolveLink(context.Context, *ResolveLinkRequest) (*ResolveLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveLink not implemented")
}
func (UnimplementedShortenerProtoServer) UserUrls(context.Context, *UserUrlsRequest) (*UserUrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUrls not implemented")
}
func (UnimplementedShortenerProtoServer) mustEmbedUnimplementedShortenerProtoServer() {}

// UnsafeShortenerProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerProtoServer will
// result in compilation errors.
type UnsafeShortenerProtoServer interface {
	mustEmbedUnimplementedShortenerProtoServer()
}

func RegisterShortenerProtoServer(s grpc.ServiceRegistrar, srv ShortenerProtoServer) {
	s.RegisterService(&ShortenerProto_ServiceDesc, srv)
}

func _ShortenerProto_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerProtoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenerProto_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerProtoServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenerProto_CreateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerProtoServer).CreateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenerProto_CreateLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerProtoServer).CreateLink(ctx, req.(*CreateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenerProto_CreateLinksBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinksBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerProtoServer).CreateLinksBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenerProto_CreateLinksBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerProtoServer).CreateLinksBatch(ctx, req.(*CreateLinksBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenerProto_ResolveLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerProtoServer).ResolveLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenerProto_ResolveLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerProtoServer).ResolveLink(ctx, req.(*ResolveLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenerProto_UserUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerProtoServer).UserUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenerProto_UserUrls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerProtoServer).UserUrls(ctx, req.(*UserUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortenerProto_ServiceDesc is the grpc.ServiceDesc for ShortenerProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortenerProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpchandlers.ShortenerProto",
	HandlerType: (*ShortenerProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ShortenerProto_Ping_Handler,
		},
		{
			MethodName: "CreateLink",
			Handler:    _ShortenerProto_CreateLink_Handler,
		},
		{
			MethodName: "CreateLinksBatch",
			Handler:    _ShortenerProto_CreateLinksBatch_Handler,
		},
		{
			MethodName: "ResolveLink",
			Handler:    _ShortenerProto_ResolveLink_Handler,
		},
		{
			MethodName: "UserUrls",
			Handler:    _ShortenerProto_UserUrls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/shortener.proto",
}

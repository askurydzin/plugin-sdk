// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package internal

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

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	// Information about what a provider supports/expects
	GetProviderSchema(ctx context.Context, in *GetProviderSchema_Request, opts ...grpc.CallOption) (*GetProviderSchema_Response, error)
	// Gets a provider's configuration example
	GetProviderConfig(ctx context.Context, in *GetProviderConfig_Request, opts ...grpc.CallOption) (*GetProviderConfig_Response, error)
	// One-time initialization, called before other functions below
	ConfigureProvider(ctx context.Context, in *ConfigureProvider_Request, opts ...grpc.CallOption) (*ConfigureProvider_Response, error)
	// Fetch Provider Resources
	FetchResources(ctx context.Context, in *FetchResources_Request, opts ...grpc.CallOption) (Provider_FetchResourcesClient, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) GetProviderSchema(ctx context.Context, in *GetProviderSchema_Request, opts ...grpc.CallOption) (*GetProviderSchema_Response, error) {
	out := new(GetProviderSchema_Response)
	err := c.cc.Invoke(ctx, "/proto.Provider/GetProviderSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) GetProviderConfig(ctx context.Context, in *GetProviderConfig_Request, opts ...grpc.CallOption) (*GetProviderConfig_Response, error) {
	out := new(GetProviderConfig_Response)
	err := c.cc.Invoke(ctx, "/proto.Provider/GetProviderConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ConfigureProvider(ctx context.Context, in *ConfigureProvider_Request, opts ...grpc.CallOption) (*ConfigureProvider_Response, error) {
	out := new(ConfigureProvider_Response)
	err := c.cc.Invoke(ctx, "/proto.Provider/ConfigureProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) FetchResources(ctx context.Context, in *FetchResources_Request, opts ...grpc.CallOption) (Provider_FetchResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Provider_ServiceDesc.Streams[0], "/proto.Provider/FetchResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &providerFetchResourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Provider_FetchResourcesClient interface {
	Recv() (*FetchResources_Response, error)
	grpc.ClientStream
}

type providerFetchResourcesClient struct {
	grpc.ClientStream
}

func (x *providerFetchResourcesClient) Recv() (*FetchResources_Response, error) {
	m := new(FetchResources_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	// Information about what a provider supports/expects
	GetProviderSchema(context.Context, *GetProviderSchema_Request) (*GetProviderSchema_Response, error)
	// Gets a provider's configuration example
	GetProviderConfig(context.Context, *GetProviderConfig_Request) (*GetProviderConfig_Response, error)
	// One-time initialization, called before other functions below
	ConfigureProvider(context.Context, *ConfigureProvider_Request) (*ConfigureProvider_Response, error)
	// Fetch Provider Resources
	FetchResources(*FetchResources_Request, Provider_FetchResourcesServer) error
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) GetProviderSchema(context.Context, *GetProviderSchema_Request) (*GetProviderSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProviderSchema not implemented")
}
func (UnimplementedProviderServer) GetProviderConfig(context.Context, *GetProviderConfig_Request) (*GetProviderConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProviderConfig not implemented")
}
func (UnimplementedProviderServer) ConfigureProvider(context.Context, *ConfigureProvider_Request) (*ConfigureProvider_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureProvider not implemented")
}
func (UnimplementedProviderServer) FetchResources(*FetchResources_Request, Provider_FetchResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchResources not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_GetProviderSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderSchema_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetProviderSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Provider/GetProviderSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetProviderSchema(ctx, req.(*GetProviderSchema_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_GetProviderConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetProviderConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Provider/GetProviderConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetProviderConfig(ctx, req.(*GetProviderConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ConfigureProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureProvider_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ConfigureProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Provider/ConfigureProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ConfigureProvider(ctx, req.(*ConfigureProvider_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_FetchResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FetchResources_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProviderServer).FetchResources(m, &providerFetchResourcesServer{stream})
}

type Provider_FetchResourcesServer interface {
	Send(*FetchResources_Response) error
	grpc.ServerStream
}

type providerFetchResourcesServer struct {
	grpc.ServerStream
}

func (x *providerFetchResourcesServer) Send(m *FetchResources_Response) error {
	return x.ServerStream.SendMsg(m)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProviderSchema",
			Handler:    _Provider_GetProviderSchema_Handler,
		},
		{
			MethodName: "GetProviderConfig",
			Handler:    _Provider_GetProviderConfig_Handler,
		},
		{
			MethodName: "ConfigureProvider",
			Handler:    _Provider_ConfigureProvider_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchResources",
			Handler:       _Provider_FetchResources_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/plugin.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: service/rfq/mm.proto

package proto

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

// MMApiClient is the client API for MMApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MMApiClient interface {
	// APIs for market makers. All rpc methods require API key authentication.
	PendingOrders(ctx context.Context, in *PendingOrdersRequest, opts ...grpc.CallOption) (*PendingOrdersResponse, error)
	UpdateOrders(ctx context.Context, in *UpdateOrdersRequest, opts ...grpc.CallOption) (*UpdateOrdersResponse, error)
	UpdateConfigs(ctx context.Context, in *UpdateConfigsRequest, opts ...grpc.CallOption) (*UpdateConfigsResponse, error)
}

type mMApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMMApiClient(cc grpc.ClientConnInterface) MMApiClient {
	return &mMApiClient{cc}
}

func (c *mMApiClient) PendingOrders(ctx context.Context, in *PendingOrdersRequest, opts ...grpc.CallOption) (*PendingOrdersResponse, error) {
	out := new(PendingOrdersResponse)
	err := c.cc.Invoke(ctx, "/service.rfq.MMApi/PendingOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mMApiClient) UpdateOrders(ctx context.Context, in *UpdateOrdersRequest, opts ...grpc.CallOption) (*UpdateOrdersResponse, error) {
	out := new(UpdateOrdersResponse)
	err := c.cc.Invoke(ctx, "/service.rfq.MMApi/UpdateOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mMApiClient) UpdateConfigs(ctx context.Context, in *UpdateConfigsRequest, opts ...grpc.CallOption) (*UpdateConfigsResponse, error) {
	out := new(UpdateConfigsResponse)
	err := c.cc.Invoke(ctx, "/service.rfq.MMApi/UpdateConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MMApiServer is the server API for MMApi service.
// All implementations should embed UnimplementedMMApiServer
// for forward compatibility
type MMApiServer interface {
	// APIs for market makers. All rpc methods require API key authentication.
	PendingOrders(context.Context, *PendingOrdersRequest) (*PendingOrdersResponse, error)
	UpdateOrders(context.Context, *UpdateOrdersRequest) (*UpdateOrdersResponse, error)
	UpdateConfigs(context.Context, *UpdateConfigsRequest) (*UpdateConfigsResponse, error)
}

// UnimplementedMMApiServer should be embedded to have forward compatible implementations.
type UnimplementedMMApiServer struct {
}

func (UnimplementedMMApiServer) PendingOrders(context.Context, *PendingOrdersRequest) (*PendingOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingOrders not implemented")
}
func (UnimplementedMMApiServer) UpdateOrders(context.Context, *UpdateOrdersRequest) (*UpdateOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (UnimplementedMMApiServer) UpdateConfigs(context.Context, *UpdateConfigsRequest) (*UpdateConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfigs not implemented")
}

// UnsafeMMApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MMApiServer will
// result in compilation errors.
type UnsafeMMApiServer interface {
	mustEmbedUnimplementedMMApiServer()
}

func RegisterMMApiServer(s grpc.ServiceRegistrar, srv MMApiServer) {
	s.RegisterService(&MMApi_ServiceDesc, srv)
}

func _MMApi_PendingOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PendingOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MMApiServer).PendingOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.rfq.MMApi/PendingOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MMApiServer).PendingOrders(ctx, req.(*PendingOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MMApi_UpdateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MMApiServer).UpdateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.rfq.MMApi/UpdateOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MMApiServer).UpdateOrders(ctx, req.(*UpdateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MMApi_UpdateConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MMApiServer).UpdateConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.rfq.MMApi/UpdateConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MMApiServer).UpdateConfigs(ctx, req.(*UpdateConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MMApi_ServiceDesc is the grpc.ServiceDesc for MMApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MMApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.rfq.MMApi",
	HandlerType: (*MMApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PendingOrders",
			Handler:    _MMApi_PendingOrders_Handler,
		},
		{
			MethodName: "UpdateOrders",
			Handler:    _MMApi_UpdateOrders_Handler,
		},
		{
			MethodName: "UpdateConfigs",
			Handler:    _MMApi_UpdateConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/rfq/mm.proto",
}
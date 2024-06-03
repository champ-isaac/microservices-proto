// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shipping

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ShippingClient is the client API for Shipping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingClient interface {
	CreateShipping(ctx context.Context, in *CreateShippingRequest, opts ...grpc.CallOption) (*CreateShippingResponse, error)
}

type shippingClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingClient(cc grpc.ClientConnInterface) ShippingClient {
	return &shippingClient{cc}
}

func (c *shippingClient) CreateShipping(ctx context.Context, in *CreateShippingRequest, opts ...grpc.CallOption) (*CreateShippingResponse, error) {
	out := new(CreateShippingResponse)
	err := c.cc.Invoke(ctx, "/Shipping/CreateShipping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServer is the server API for Shipping service.
// All implementations must embed UnimplementedShippingServer
// for forward compatibility
type ShippingServer interface {
	CreateShipping(context.Context, *CreateShippingRequest) (*CreateShippingResponse, error)
	mustEmbedUnimplementedShippingServer()
}

// UnimplementedShippingServer must be embedded to have forward compatible implementations.
type UnimplementedShippingServer struct {
}

func (UnimplementedShippingServer) CreateShipping(context.Context, *CreateShippingRequest) (*CreateShippingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipping not implemented")
}
func (UnimplementedShippingServer) mustEmbedUnimplementedShippingServer() {}

// UnsafeShippingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServer will
// result in compilation errors.
type UnsafeShippingServer interface {
	mustEmbedUnimplementedShippingServer()
}

func RegisterShippingServer(s grpc.ServiceRegistrar, srv ShippingServer) {
	s.RegisterService(&Shipping_ServiceDesc, srv)
}

func _Shipping_CreateShipping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShippingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).CreateShipping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shipping/CreateShipping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).CreateShipping(ctx, req.(*CreateShippingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shipping_ServiceDesc is the grpc.ServiceDesc for Shipping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shipping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Shipping",
	HandlerType: (*ShippingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShipping",
			Handler:    _Shipping_CreateShipping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipping/shipping.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package targetservice

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

// BouncerClient is the client API for Bouncer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BouncerClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// define a gRPC method that's not implemented in the target
	UnknownMethod(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	BounceIt(ctx context.Context, in *BallIn, opts ...grpc.CallOption) (*BallOut, error)
	GrowTail(ctx context.Context, in *Body, opts ...grpc.CallOption) (*Body, error)
}

type bouncerClient struct {
	cc grpc.ClientConnInterface
}

func NewBouncerClient(cc grpc.ClientConnInterface) BouncerClient {
	return &bouncerClient{cc}
}

func (c *bouncerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/targetservice.Bouncer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bouncerClient) UnknownMethod(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/targetservice.Bouncer/UnknownMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bouncerClient) BounceIt(ctx context.Context, in *BallIn, opts ...grpc.CallOption) (*BallOut, error) {
	out := new(BallOut)
	err := c.cc.Invoke(ctx, "/targetservice.Bouncer/BounceIt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bouncerClient) GrowTail(ctx context.Context, in *Body, opts ...grpc.CallOption) (*Body, error) {
	out := new(Body)
	err := c.cc.Invoke(ctx, "/targetservice.Bouncer/GrowTail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BouncerServer is the server API for Bouncer service.
// All implementations must embed UnimplementedBouncerServer
// for forward compatibility
type BouncerServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	// define a gRPC method that's not implemented in the target
	UnknownMethod(context.Context, *HelloRequest) (*HelloResponse, error)
	BounceIt(context.Context, *BallIn) (*BallOut, error)
	GrowTail(context.Context, *Body) (*Body, error)
	mustEmbedUnimplementedBouncerServer()
}

// UnimplementedBouncerServer must be embedded to have forward compatible implementations.
type UnimplementedBouncerServer struct {
}

func (UnimplementedBouncerServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBouncerServer) UnknownMethod(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnknownMethod not implemented")
}
func (UnimplementedBouncerServer) BounceIt(context.Context, *BallIn) (*BallOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BounceIt not implemented")
}
func (UnimplementedBouncerServer) GrowTail(context.Context, *Body) (*Body, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrowTail not implemented")
}
func (UnimplementedBouncerServer) mustEmbedUnimplementedBouncerServer() {}

// UnsafeBouncerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BouncerServer will
// result in compilation errors.
type UnsafeBouncerServer interface {
	mustEmbedUnimplementedBouncerServer()
}

func RegisterBouncerServer(s grpc.ServiceRegistrar, srv BouncerServer) {
	s.RegisterService(&Bouncer_ServiceDesc, srv)
}

func _Bouncer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BouncerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/targetservice.Bouncer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BouncerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bouncer_UnknownMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BouncerServer).UnknownMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/targetservice.Bouncer/UnknownMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BouncerServer).UnknownMethod(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bouncer_BounceIt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BallIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BouncerServer).BounceIt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/targetservice.Bouncer/BounceIt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BouncerServer).BounceIt(ctx, req.(*BallIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bouncer_GrowTail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Body)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BouncerServer).GrowTail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/targetservice.Bouncer/GrowTail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BouncerServer).GrowTail(ctx, req.(*Body))
	}
	return interceptor(ctx, in, info, handler)
}

// Bouncer_ServiceDesc is the grpc.ServiceDesc for Bouncer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bouncer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "targetservice.Bouncer",
	HandlerType: (*BouncerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Bouncer_SayHello_Handler,
		},
		{
			MethodName: "UnknownMethod",
			Handler:    _Bouncer_UnknownMethod_Handler,
		},
		{
			MethodName: "BounceIt",
			Handler:    _Bouncer_BounceIt_Handler,
		},
		{
			MethodName: "GrowTail",
			Handler:    _Bouncer_GrowTail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "targetservice.proto",
}

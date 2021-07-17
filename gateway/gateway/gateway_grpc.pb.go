// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// GateWayClient is the client API for GateWay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GateWayClient interface {
	Echo(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
}

type gateWayClient struct {
	cc grpc.ClientConnInterface
}

func NewGateWayClient(cc grpc.ClientConnInterface) GateWayClient {
	return &gateWayClient{cc}
}

func (c *gateWayClient) Echo(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/gateway.GateWay/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateWayServer is the server API for GateWay service.
// All implementations must embed UnimplementedGateWayServer
// for forward compatibility
type GateWayServer interface {
	Echo(context.Context, *StringRequest) (*StringResponse, error)
	mustEmbedUnimplementedGateWayServer()
}

// UnimplementedGateWayServer must be embedded to have forward compatible implementations.
type UnimplementedGateWayServer struct {
}

func (UnimplementedGateWayServer) Echo(context.Context, *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedGateWayServer) mustEmbedUnimplementedGateWayServer() {}

// UnsafeGateWayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GateWayServer will
// result in compilation errors.
type UnsafeGateWayServer interface {
	mustEmbedUnimplementedGateWayServer()
}

func RegisterGateWayServer(s grpc.ServiceRegistrar, srv GateWayServer) {
	s.RegisterService(&GateWay_ServiceDesc, srv)
}

func _GateWay_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateWayServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GateWay/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateWayServer).Echo(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GateWay_ServiceDesc is the grpc.ServiceDesc for GateWay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GateWay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.GateWay",
	HandlerType: (*GateWayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _GateWay_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/market.proto

package tdexav1

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

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	ListMarketIDs(ctx context.Context, in *ListMarketIDsRequest, opts ...grpc.CallOption) (*ListMarketIDsReply, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) ListMarketIDs(ctx context.Context, in *ListMarketIDsRequest, opts ...grpc.CallOption) (*ListMarketIDsReply, error) {
	out := new(ListMarketIDsReply)
	err := c.cc.Invoke(ctx, "/tdexa.v1.Market/ListMarketIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations should embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	ListMarketIDs(context.Context, *ListMarketIDsRequest) (*ListMarketIDsReply, error)
}

// UnimplementedMarketServer should be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) ListMarketIDs(context.Context, *ListMarketIDsRequest) (*ListMarketIDsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMarketIDs not implemented")
}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_ListMarketIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMarketIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).ListMarketIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdexa.v1.Market/ListMarketIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).ListMarketIDs(ctx, req.(*ListMarketIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tdexa.v1.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMarketIDs",
			Handler:    _Market_ListMarketIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/market.proto",
}

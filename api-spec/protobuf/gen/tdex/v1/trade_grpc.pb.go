// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tdex/v1/trade.proto

package tdexv1

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

// TradeServiceClient is the client API for TradeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeServiceClient interface {
	// ListMarkets lists all the markets open for trading.
	ListMarkets(ctx context.Context, in *ListMarketsRequest, opts ...grpc.CallOption) (*ListMarketsResponse, error)
	// GetMarketBalance retutns the balance of the two current reserves of the
	// given market.
	GetMarketBalance(ctx context.Context, in *GetMarketBalanceRequest, opts ...grpc.CallOption) (*GetMarketBalanceResponse, error)
	// GetMarketPrice retutns the spot price for the requested market and its
	// minimum tradable amount of base asset.
	GetMarketPrice(ctx context.Context, in *GetMarketPriceRequest, opts ...grpc.CallOption) (*GetMarketPriceResponse, error)
	// PreviewTrade returns a counter amount and asset in response to the
	// provided ones and a trade type for a market.
	//
	// The trade type can assume values BUY or SELL and it always refer to the
	// fixed base asset.
	// For example:
	//  * if trade type is BUY, it means the trader wants to buy base asset funds.
	//  * if trade type is SELL, it means the trader wants to sell base asset funds.
	PreviewTrade(ctx context.Context, in *PreviewTradeRequest, opts ...grpc.CallOption) (*PreviewTradeResponse, error)
	// ProposeTrade allows a trader to present a SwapRequest. The service answers
	// with a SwapAccept, filling the request's partial transaction, + an
	// expiration time to complete the swap when accepting the swap, or,
	// otherwise, with a SwapFail containg the reason for the rejection of the
	// proposal.
	ProposeTrade(ctx context.Context, in *ProposeTradeRequest, opts ...grpc.CallOption) (*ProposeTradeResponse, error)
	// CompleteTrade can be used by the trader to let the daemon finalizing,
	// extracting, and broadcasting the swap transaction, once he's signed his
	// inputs.
	// This is not mandatory, the trader can do the steps above on his own
	// alternatively.
	CompleteTrade(ctx context.Context, in *CompleteTradeRequest, opts ...grpc.CallOption) (*CompleteTradeResponse, error)
}

type tradeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeServiceClient(cc grpc.ClientConnInterface) TradeServiceClient {
	return &tradeServiceClient{cc}
}

func (c *tradeServiceClient) ListMarkets(ctx context.Context, in *ListMarketsRequest, opts ...grpc.CallOption) (*ListMarketsResponse, error) {
	out := new(ListMarketsResponse)
	err := c.cc.Invoke(ctx, "/tdex.v1.TradeService/ListMarkets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) GetMarketBalance(ctx context.Context, in *GetMarketBalanceRequest, opts ...grpc.CallOption) (*GetMarketBalanceResponse, error) {
	out := new(GetMarketBalanceResponse)
	err := c.cc.Invoke(ctx, "/tdex.v1.TradeService/GetMarketBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) GetMarketPrice(ctx context.Context, in *GetMarketPriceRequest, opts ...grpc.CallOption) (*GetMarketPriceResponse, error) {
	out := new(GetMarketPriceResponse)
	err := c.cc.Invoke(ctx, "/tdex.v1.TradeService/GetMarketPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) PreviewTrade(ctx context.Context, in *PreviewTradeRequest, opts ...grpc.CallOption) (*PreviewTradeResponse, error) {
	out := new(PreviewTradeResponse)
	err := c.cc.Invoke(ctx, "/tdex.v1.TradeService/PreviewTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) ProposeTrade(ctx context.Context, in *ProposeTradeRequest, opts ...grpc.CallOption) (*ProposeTradeResponse, error) {
	out := new(ProposeTradeResponse)
	err := c.cc.Invoke(ctx, "/tdex.v1.TradeService/ProposeTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) CompleteTrade(ctx context.Context, in *CompleteTradeRequest, opts ...grpc.CallOption) (*CompleteTradeResponse, error) {
	out := new(CompleteTradeResponse)
	err := c.cc.Invoke(ctx, "/tdex.v1.TradeService/CompleteTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradeServiceServer is the server API for TradeService service.
// All implementations should embed UnimplementedTradeServiceServer
// for forward compatibility
type TradeServiceServer interface {
	// ListMarkets lists all the markets open for trading.
	ListMarkets(context.Context, *ListMarketsRequest) (*ListMarketsResponse, error)
	// GetMarketBalance retutns the balance of the two current reserves of the
	// given market.
	GetMarketBalance(context.Context, *GetMarketBalanceRequest) (*GetMarketBalanceResponse, error)
	// GetMarketPrice retutns the spot price for the requested market and its
	// minimum tradable amount of base asset.
	GetMarketPrice(context.Context, *GetMarketPriceRequest) (*GetMarketPriceResponse, error)
	// PreviewTrade returns a counter amount and asset in response to the
	// provided ones and a trade type for a market.
	//
	// The trade type can assume values BUY or SELL and it always refer to the
	// fixed base asset.
	// For example:
	//  * if trade type is BUY, it means the trader wants to buy base asset funds.
	//  * if trade type is SELL, it means the trader wants to sell base asset funds.
	PreviewTrade(context.Context, *PreviewTradeRequest) (*PreviewTradeResponse, error)
	// ProposeTrade allows a trader to present a SwapRequest. The service answers
	// with a SwapAccept, filling the request's partial transaction, + an
	// expiration time to complete the swap when accepting the swap, or,
	// otherwise, with a SwapFail containg the reason for the rejection of the
	// proposal.
	ProposeTrade(context.Context, *ProposeTradeRequest) (*ProposeTradeResponse, error)
	// CompleteTrade can be used by the trader to let the daemon finalizing,
	// extracting, and broadcasting the swap transaction, once he's signed his
	// inputs.
	// This is not mandatory, the trader can do the steps above on his own
	// alternatively.
	CompleteTrade(context.Context, *CompleteTradeRequest) (*CompleteTradeResponse, error)
}

// UnimplementedTradeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTradeServiceServer struct {
}

func (UnimplementedTradeServiceServer) ListMarkets(context.Context, *ListMarketsRequest) (*ListMarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMarkets not implemented")
}
func (UnimplementedTradeServiceServer) GetMarketBalance(context.Context, *GetMarketBalanceRequest) (*GetMarketBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketBalance not implemented")
}
func (UnimplementedTradeServiceServer) GetMarketPrice(context.Context, *GetMarketPriceRequest) (*GetMarketPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketPrice not implemented")
}
func (UnimplementedTradeServiceServer) PreviewTrade(context.Context, *PreviewTradeRequest) (*PreviewTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewTrade not implemented")
}
func (UnimplementedTradeServiceServer) ProposeTrade(context.Context, *ProposeTradeRequest) (*ProposeTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposeTrade not implemented")
}
func (UnimplementedTradeServiceServer) CompleteTrade(context.Context, *CompleteTradeRequest) (*CompleteTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTrade not implemented")
}

// UnsafeTradeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradeServiceServer will
// result in compilation errors.
type UnsafeTradeServiceServer interface {
	mustEmbedUnimplementedTradeServiceServer()
}

func RegisterTradeServiceServer(s grpc.ServiceRegistrar, srv TradeServiceServer) {
	s.RegisterService(&TradeService_ServiceDesc, srv)
}

func _TradeService_ListMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMarketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).ListMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdex.v1.TradeService/ListMarkets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).ListMarkets(ctx, req.(*ListMarketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_GetMarketBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).GetMarketBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdex.v1.TradeService/GetMarketBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).GetMarketBalance(ctx, req.(*GetMarketBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_GetMarketPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).GetMarketPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdex.v1.TradeService/GetMarketPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).GetMarketPrice(ctx, req.(*GetMarketPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_PreviewTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviewTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).PreviewTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdex.v1.TradeService/PreviewTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).PreviewTrade(ctx, req.(*PreviewTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_ProposeTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).ProposeTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdex.v1.TradeService/ProposeTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).ProposeTrade(ctx, req.(*ProposeTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_CompleteTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).CompleteTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdex.v1.TradeService/CompleteTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).CompleteTrade(ctx, req.(*CompleteTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TradeService_ServiceDesc is the grpc.ServiceDesc for TradeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tdex.v1.TradeService",
	HandlerType: (*TradeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMarkets",
			Handler:    _TradeService_ListMarkets_Handler,
		},
		{
			MethodName: "GetMarketBalance",
			Handler:    _TradeService_GetMarketBalance_Handler,
		},
		{
			MethodName: "GetMarketPrice",
			Handler:    _TradeService_GetMarketPrice_Handler,
		},
		{
			MethodName: "PreviewTrade",
			Handler:    _TradeService_PreviewTrade_Handler,
		},
		{
			MethodName: "ProposeTrade",
			Handler:    _TradeService_ProposeTrade_Handler,
		},
		{
			MethodName: "CompleteTrade",
			Handler:    _TradeService_CompleteTrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tdex/v1/trade.proto",
}

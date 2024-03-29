// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: balances/v1/balances.proto

package balancesv1

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

// BalancesClient is the client API for Balances service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalancesClient interface {
	ListBalances(ctx context.Context, in *ListBalancesRequest, opts ...grpc.CallOption) (*ListBalancesResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	CreateBalance(ctx context.Context, in *CreateBalanceRequest, opts ...grpc.CallOption) (*CreateBalanceResponse, error)
	UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error)
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error)
}

type balancesClient struct {
	cc grpc.ClientConnInterface
}

func NewBalancesClient(cc grpc.ClientConnInterface) BalancesClient {
	return &balancesClient{cc}
}

func (c *balancesClient) ListBalances(ctx context.Context, in *ListBalancesRequest, opts ...grpc.CallOption) (*ListBalancesResponse, error) {
	out := new(ListBalancesResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/ListBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancesClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancesClient) CreateBalance(ctx context.Context, in *CreateBalanceRequest, opts ...grpc.CallOption) (*CreateBalanceResponse, error) {
	out := new(CreateBalanceResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/CreateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancesClient) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error) {
	out := new(UpdateBalanceResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/UpdateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancesClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/ListEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancesClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancesClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error) {
	out := new(CreateEntryResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/CreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancesClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error) {
	out := new(UpdateEntryResponse)
	err := c.cc.Invoke(ctx, "/biscuit.balances.v1.Balances/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalancesServer is the server API for Balances service.
// All implementations must embed UnimplementedBalancesServer
// for forward compatibility
type BalancesServer interface {
	ListBalances(context.Context, *ListBalancesRequest) (*ListBalancesResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	CreateBalance(context.Context, *CreateBalanceRequest) (*CreateBalanceResponse, error)
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error)
	ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error)
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	CreateEntry(context.Context, *CreateEntryRequest) (*CreateEntryResponse, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error)
	mustEmbedUnimplementedBalancesServer()
}

// UnimplementedBalancesServer must be embedded to have forward compatible implementations.
type UnimplementedBalancesServer struct {
}

func (UnimplementedBalancesServer) ListBalances(context.Context, *ListBalancesRequest) (*ListBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBalances not implemented")
}
func (UnimplementedBalancesServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBalancesServer) CreateBalance(context.Context, *CreateBalanceRequest) (*CreateBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBalance not implemented")
}
func (UnimplementedBalancesServer) UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedBalancesServer) ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}
func (UnimplementedBalancesServer) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedBalancesServer) CreateEntry(context.Context, *CreateEntryRequest) (*CreateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntry not implemented")
}
func (UnimplementedBalancesServer) UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (UnimplementedBalancesServer) mustEmbedUnimplementedBalancesServer() {}

// UnsafeBalancesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalancesServer will
// result in compilation errors.
type UnsafeBalancesServer interface {
	mustEmbedUnimplementedBalancesServer()
}

func RegisterBalancesServer(s grpc.ServiceRegistrar, srv BalancesServer) {
	s.RegisterService(&Balances_ServiceDesc, srv)
}

func _Balances_ListBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).ListBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/ListBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).ListBalances(ctx, req.(*ListBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balances_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balances_CreateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).CreateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/CreateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).CreateBalance(ctx, req.(*CreateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balances_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/UpdateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).UpdateBalance(ctx, req.(*UpdateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balances_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balances_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balances_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balances_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biscuit.balances.v1.Balances/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Balances_ServiceDesc is the grpc.ServiceDesc for Balances service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Balances_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "biscuit.balances.v1.Balances",
	HandlerType: (*BalancesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBalances",
			Handler:    _Balances_ListBalances_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Balances_GetBalance_Handler,
		},
		{
			MethodName: "CreateBalance",
			Handler:    _Balances_CreateBalance_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _Balances_UpdateBalance_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _Balances_ListEntries_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _Balances_GetEntry_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _Balances_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _Balances_UpdateEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "balances/v1/balances.proto",
}

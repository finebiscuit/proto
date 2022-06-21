// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: balances/v1/balances.proto

package balancesv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/finebiscuit/proto/biscuit/balances/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// BalancesName is the fully-qualified name of the Balances service.
	BalancesName = "biscuit.balances.v1.Balances"
)

// BalancesClient is a client for the biscuit.balances.v1.Balances service.
type BalancesClient interface {
	ListBalances(context.Context, *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error)
	GetBalance(context.Context, *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error)
	CreateBalance(context.Context, *connect_go.Request[v1.CreateBalanceRequest]) (*connect_go.Response[v1.CreateBalanceResponse], error)
	UpdateBalance(context.Context, *connect_go.Request[v1.UpdateBalanceRequest]) (*connect_go.Response[v1.UpdateBalanceResponse], error)
	ListEntries(context.Context, *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error)
	GetEntry(context.Context, *connect_go.Request[v1.GetEntryRequest]) (*connect_go.Response[v1.GetEntryResponse], error)
	CreateEntry(context.Context, *connect_go.Request[v1.CreateEntryRequest]) (*connect_go.Response[v1.CreateEntryResponse], error)
	UpdateEntry(context.Context, *connect_go.Request[v1.UpdateEntryRequest]) (*connect_go.Response[v1.UpdateEntryResponse], error)
}

// NewBalancesClient constructs a client for the biscuit.balances.v1.Balances service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBalancesClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BalancesClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &balancesClient{
		listBalances: connect_go.NewClient[v1.ListBalancesRequest, v1.ListBalancesResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/ListBalances",
			opts...,
		),
		getBalance: connect_go.NewClient[v1.GetBalanceRequest, v1.GetBalanceResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/GetBalance",
			opts...,
		),
		createBalance: connect_go.NewClient[v1.CreateBalanceRequest, v1.CreateBalanceResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/CreateBalance",
			opts...,
		),
		updateBalance: connect_go.NewClient[v1.UpdateBalanceRequest, v1.UpdateBalanceResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/UpdateBalance",
			opts...,
		),
		listEntries: connect_go.NewClient[v1.ListEntriesRequest, v1.ListEntriesResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/ListEntries",
			opts...,
		),
		getEntry: connect_go.NewClient[v1.GetEntryRequest, v1.GetEntryResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/GetEntry",
			opts...,
		),
		createEntry: connect_go.NewClient[v1.CreateEntryRequest, v1.CreateEntryResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/CreateEntry",
			opts...,
		),
		updateEntry: connect_go.NewClient[v1.UpdateEntryRequest, v1.UpdateEntryResponse](
			httpClient,
			baseURL+"/biscuit.balances.v1.Balances/UpdateEntry",
			opts...,
		),
	}
}

// balancesClient implements BalancesClient.
type balancesClient struct {
	listBalances  *connect_go.Client[v1.ListBalancesRequest, v1.ListBalancesResponse]
	getBalance    *connect_go.Client[v1.GetBalanceRequest, v1.GetBalanceResponse]
	createBalance *connect_go.Client[v1.CreateBalanceRequest, v1.CreateBalanceResponse]
	updateBalance *connect_go.Client[v1.UpdateBalanceRequest, v1.UpdateBalanceResponse]
	listEntries   *connect_go.Client[v1.ListEntriesRequest, v1.ListEntriesResponse]
	getEntry      *connect_go.Client[v1.GetEntryRequest, v1.GetEntryResponse]
	createEntry   *connect_go.Client[v1.CreateEntryRequest, v1.CreateEntryResponse]
	updateEntry   *connect_go.Client[v1.UpdateEntryRequest, v1.UpdateEntryResponse]
}

// ListBalances calls biscuit.balances.v1.Balances.ListBalances.
func (c *balancesClient) ListBalances(ctx context.Context, req *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error) {
	return c.listBalances.CallUnary(ctx, req)
}

// GetBalance calls biscuit.balances.v1.Balances.GetBalance.
func (c *balancesClient) GetBalance(ctx context.Context, req *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error) {
	return c.getBalance.CallUnary(ctx, req)
}

// CreateBalance calls biscuit.balances.v1.Balances.CreateBalance.
func (c *balancesClient) CreateBalance(ctx context.Context, req *connect_go.Request[v1.CreateBalanceRequest]) (*connect_go.Response[v1.CreateBalanceResponse], error) {
	return c.createBalance.CallUnary(ctx, req)
}

// UpdateBalance calls biscuit.balances.v1.Balances.UpdateBalance.
func (c *balancesClient) UpdateBalance(ctx context.Context, req *connect_go.Request[v1.UpdateBalanceRequest]) (*connect_go.Response[v1.UpdateBalanceResponse], error) {
	return c.updateBalance.CallUnary(ctx, req)
}

// ListEntries calls biscuit.balances.v1.Balances.ListEntries.
func (c *balancesClient) ListEntries(ctx context.Context, req *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error) {
	return c.listEntries.CallUnary(ctx, req)
}

// GetEntry calls biscuit.balances.v1.Balances.GetEntry.
func (c *balancesClient) GetEntry(ctx context.Context, req *connect_go.Request[v1.GetEntryRequest]) (*connect_go.Response[v1.GetEntryResponse], error) {
	return c.getEntry.CallUnary(ctx, req)
}

// CreateEntry calls biscuit.balances.v1.Balances.CreateEntry.
func (c *balancesClient) CreateEntry(ctx context.Context, req *connect_go.Request[v1.CreateEntryRequest]) (*connect_go.Response[v1.CreateEntryResponse], error) {
	return c.createEntry.CallUnary(ctx, req)
}

// UpdateEntry calls biscuit.balances.v1.Balances.UpdateEntry.
func (c *balancesClient) UpdateEntry(ctx context.Context, req *connect_go.Request[v1.UpdateEntryRequest]) (*connect_go.Response[v1.UpdateEntryResponse], error) {
	return c.updateEntry.CallUnary(ctx, req)
}

// BalancesHandler is an implementation of the biscuit.balances.v1.Balances service.
type BalancesHandler interface {
	ListBalances(context.Context, *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error)
	GetBalance(context.Context, *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error)
	CreateBalance(context.Context, *connect_go.Request[v1.CreateBalanceRequest]) (*connect_go.Response[v1.CreateBalanceResponse], error)
	UpdateBalance(context.Context, *connect_go.Request[v1.UpdateBalanceRequest]) (*connect_go.Response[v1.UpdateBalanceResponse], error)
	ListEntries(context.Context, *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error)
	GetEntry(context.Context, *connect_go.Request[v1.GetEntryRequest]) (*connect_go.Response[v1.GetEntryResponse], error)
	CreateEntry(context.Context, *connect_go.Request[v1.CreateEntryRequest]) (*connect_go.Response[v1.CreateEntryResponse], error)
	UpdateEntry(context.Context, *connect_go.Request[v1.UpdateEntryRequest]) (*connect_go.Response[v1.UpdateEntryResponse], error)
}

// NewBalancesHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBalancesHandler(svc BalancesHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/biscuit.balances.v1.Balances/ListBalances", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/ListBalances",
		svc.ListBalances,
		opts...,
	))
	mux.Handle("/biscuit.balances.v1.Balances/GetBalance", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/GetBalance",
		svc.GetBalance,
		opts...,
	))
	mux.Handle("/biscuit.balances.v1.Balances/CreateBalance", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/CreateBalance",
		svc.CreateBalance,
		opts...,
	))
	mux.Handle("/biscuit.balances.v1.Balances/UpdateBalance", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/UpdateBalance",
		svc.UpdateBalance,
		opts...,
	))
	mux.Handle("/biscuit.balances.v1.Balances/ListEntries", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/ListEntries",
		svc.ListEntries,
		opts...,
	))
	mux.Handle("/biscuit.balances.v1.Balances/GetEntry", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/GetEntry",
		svc.GetEntry,
		opts...,
	))
	mux.Handle("/biscuit.balances.v1.Balances/CreateEntry", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/CreateEntry",
		svc.CreateEntry,
		opts...,
	))
	mux.Handle("/biscuit.balances.v1.Balances/UpdateEntry", connect_go.NewUnaryHandler(
		"/biscuit.balances.v1.Balances/UpdateEntry",
		svc.UpdateEntry,
		opts...,
	))
	return "/biscuit.balances.v1.Balances/", mux
}

// UnimplementedBalancesHandler returns CodeUnimplemented from all methods.
type UnimplementedBalancesHandler struct{}

func (UnimplementedBalancesHandler) ListBalances(context.Context, *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.ListBalances is not implemented"))
}

func (UnimplementedBalancesHandler) GetBalance(context.Context, *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.GetBalance is not implemented"))
}

func (UnimplementedBalancesHandler) CreateBalance(context.Context, *connect_go.Request[v1.CreateBalanceRequest]) (*connect_go.Response[v1.CreateBalanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.CreateBalance is not implemented"))
}

func (UnimplementedBalancesHandler) UpdateBalance(context.Context, *connect_go.Request[v1.UpdateBalanceRequest]) (*connect_go.Response[v1.UpdateBalanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.UpdateBalance is not implemented"))
}

func (UnimplementedBalancesHandler) ListEntries(context.Context, *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.ListEntries is not implemented"))
}

func (UnimplementedBalancesHandler) GetEntry(context.Context, *connect_go.Request[v1.GetEntryRequest]) (*connect_go.Response[v1.GetEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.GetEntry is not implemented"))
}

func (UnimplementedBalancesHandler) CreateEntry(context.Context, *connect_go.Request[v1.CreateEntryRequest]) (*connect_go.Response[v1.CreateEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.CreateEntry is not implemented"))
}

func (UnimplementedBalancesHandler) UpdateEntry(context.Context, *connect_go.Request[v1.UpdateEntryRequest]) (*connect_go.Response[v1.UpdateEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("biscuit.balances.v1.Balances.UpdateEntry is not implemented"))
}

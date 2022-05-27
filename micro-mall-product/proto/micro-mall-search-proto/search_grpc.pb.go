// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: mall-demo/micro-mall-product/proto/micro-mall-search-proto/search.proto

package proto_search

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

// SearchRpcClient is the client API for SearchRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchRpcClient interface {
	ProductStatusUp(ctx context.Context, in *ProductStatusUpRequest, opts ...grpc.CallOption) (*ProductStatusUpResponse, error)
}

type searchRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchRpcClient(cc grpc.ClientConnInterface) SearchRpcClient {
	return &searchRpcClient{cc}
}

func (c *searchRpcClient) ProductStatusUp(ctx context.Context, in *ProductStatusUpRequest, opts ...grpc.CallOption) (*ProductStatusUpResponse, error) {
	out := new(ProductStatusUpResponse)
	err := c.cc.Invoke(ctx, "/search.SearchRpc/ProductStatusUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRpcServer is the server API for SearchRpc service.
// All implementations must embed UnimplementedSearchRpcServer
// for forward compatibility
type SearchRpcServer interface {
	ProductStatusUp(context.Context, *ProductStatusUpRequest) (*ProductStatusUpResponse, error)
	mustEmbedUnimplementedSearchRpcServer()
}

// UnimplementedSearchRpcServer must be embedded to have forward compatible implementations.
type UnimplementedSearchRpcServer struct {
}

func (UnimplementedSearchRpcServer) ProductStatusUp(context.Context, *ProductStatusUpRequest) (*ProductStatusUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductStatusUp not implemented")
}
func (UnimplementedSearchRpcServer) mustEmbedUnimplementedSearchRpcServer() {}

// UnsafeSearchRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchRpcServer will
// result in compilation errors.
type UnsafeSearchRpcServer interface {
	mustEmbedUnimplementedSearchRpcServer()
}

func RegisterSearchRpcServer(s grpc.ServiceRegistrar, srv SearchRpcServer) {
	s.RegisterService(&SearchRpc_ServiceDesc, srv)
}

func _SearchRpc_ProductStatusUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductStatusUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchRpcServer).ProductStatusUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchRpc/ProductStatusUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchRpcServer).ProductStatusUp(ctx, req.(*ProductStatusUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchRpc_ServiceDesc is the grpc.ServiceDesc for SearchRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchRpc",
	HandlerType: (*SearchRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductStatusUp",
			Handler:    _SearchRpc_ProductStatusUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mall-demo/micro-mall-product/proto/micro-mall-search-proto/search.proto",
}

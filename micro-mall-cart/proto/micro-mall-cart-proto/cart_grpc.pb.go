// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: mall-demo/micro-mall-cart/proto/micro-mall-cart-proto/cart.proto

package proto_cart

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CartRpcClient is the client API for CartRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartRpcClient interface {
}

type cartRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCartRpcClient(cc grpc.ClientConnInterface) CartRpcClient {
	return &cartRpcClient{cc}
}

// CartRpcServer is the server API for CartRpc service.
// All implementations must embed UnimplementedCartRpcServer
// for forward compatibility
type CartRpcServer interface {
	mustEmbedUnimplementedCartRpcServer()
}

// UnimplementedCartRpcServer must be embedded to have forward compatible implementations.
type UnimplementedCartRpcServer struct {
}

func (UnimplementedCartRpcServer) mustEmbedUnimplementedCartRpcServer() {}

// UnsafeCartRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartRpcServer will
// result in compilation errors.
type UnsafeCartRpcServer interface {
	mustEmbedUnimplementedCartRpcServer()
}

func RegisterCartRpcServer(s grpc.ServiceRegistrar, srv CartRpcServer) {
	s.RegisterService(&CartRpc_ServiceDesc, srv)
}

// CartRpc_ServiceDesc is the grpc.ServiceDesc for CartRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_cart.CartRpc",
	HandlerType: (*CartRpcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "mall-demo/micro-mall-cart/proto/micro-mall-cart-proto/cart.proto",
}
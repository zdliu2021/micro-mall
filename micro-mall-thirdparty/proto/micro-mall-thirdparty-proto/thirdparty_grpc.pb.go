// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: mall-demo/micro-mall-thirdparty/proto/micro-mall-thirdparty-proto/thirdparty.proto

package proto_thirdparty

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

// ThirdPartyRpcClient is the client API for ThirdPartyRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdPartyRpcClient interface {
	GetOSSToken(ctx context.Context, in *GetOSSTokenRequest, opts ...grpc.CallOption) (*GetOSSTokenResponse, error)
}

type thirdPartyRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdPartyRpcClient(cc grpc.ClientConnInterface) ThirdPartyRpcClient {
	return &thirdPartyRpcClient{cc}
}

func (c *thirdPartyRpcClient) GetOSSToken(ctx context.Context, in *GetOSSTokenRequest, opts ...grpc.CallOption) (*GetOSSTokenResponse, error) {
	out := new(GetOSSTokenResponse)
	err := c.cc.Invoke(ctx, "/thirdparty.ThirdPartyRpc/GetOSSToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdPartyRpcServer is the server API for ThirdPartyRpc service.
// All implementations must embed UnimplementedThirdPartyRpcServer
// for forward compatibility
type ThirdPartyRpcServer interface {
	GetOSSToken(context.Context, *GetOSSTokenRequest) (*GetOSSTokenResponse, error)
	mustEmbedUnimplementedThirdPartyRpcServer()
}

// UnimplementedThirdPartyRpcServer must be embedded to have forward compatible implementations.
type UnimplementedThirdPartyRpcServer struct {
}

func (UnimplementedThirdPartyRpcServer) GetOSSToken(context.Context, *GetOSSTokenRequest) (*GetOSSTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOSSToken not implemented")
}
func (UnimplementedThirdPartyRpcServer) mustEmbedUnimplementedThirdPartyRpcServer() {}

// UnsafeThirdPartyRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdPartyRpcServer will
// result in compilation errors.
type UnsafeThirdPartyRpcServer interface {
	mustEmbedUnimplementedThirdPartyRpcServer()
}

func RegisterThirdPartyRpcServer(s grpc.ServiceRegistrar, srv ThirdPartyRpcServer) {
	s.RegisterService(&ThirdPartyRpc_ServiceDesc, srv)
}

func _ThirdPartyRpc_GetOSSToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOSSTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyRpcServer).GetOSSToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdparty.ThirdPartyRpc/GetOSSToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyRpcServer).GetOSSToken(ctx, req.(*GetOSSTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThirdPartyRpc_ServiceDesc is the grpc.ServiceDesc for ThirdPartyRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThirdPartyRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thirdparty.ThirdPartyRpc",
	HandlerType: (*ThirdPartyRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOSSToken",
			Handler:    _ThirdPartyRpc_GetOSSToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mall-demo/micro-mall-thirdparty/proto/micro-mall-thirdparty-proto/thirdparty.proto",
}

package server

import (
	"context"
	proto_thirdparty "mall-demo/micro-mall-thirdparty/proto/micro-mall-thirdparty-proto"
	"mall-demo/micro-mall-thirdparty/service"
)

type Server struct {
	proto_thirdparty.UnimplementedThirdPartyRpcServer
	ossService service.OSS
}

func (s *Server) GetOSSToken(ctx context.Context, req *proto_thirdparty.GetOSSTokenRequest) (*proto_thirdparty.GetOSSTokenResponse, error) {
	return s.ossService.GetOSSToken(ctx, req)
}

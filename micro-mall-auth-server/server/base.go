package server

import (
	"context"
	proto_auth_server "mall-demo/micro-mall-auth-server/proto/micro-mall-auth-server"
	"mall-demo/micro-mall-auth-server/service"
)

type Server struct {
	proto_auth_server.UnimplementedAuthServerRpcServer
	bs *service.BaseService
}

func (s *Server) SendCode(ctx context.Context, req *proto_auth_server.SendCodeRequest) (*proto_auth_server.SendCodeResponse, error) {
	return s.bs.SendCode(ctx, req)
}

func (s *Server) Register(ctx context.Context, req *proto_auth_server.RegisterRequest) (*proto_auth_server.RegisterResponse, error) {
	return s.bs.Register(ctx, req)
}

func (s *Server) Login(ctx context.Context, req *proto_auth_server.LoginRequest) (*proto_auth_server.LoginResponse, error) {
	return s.bs.Login(ctx, req)
}

func (s *Server) OAuthGitteSuccess(ctx context.Context, req *proto_auth_server.OAuthGitteSuccessRequest) (*proto_auth_server.OAuthGitteSuccessResponse, error) {
	return s.bs.OAuthGitteSuccess(ctx, req)
}

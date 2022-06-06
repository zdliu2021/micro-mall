package server

import (
	"context"
	proto_search "mall-demo/micro-mall-search/proto/micro-mall-search-proto"
	"mall-demo/micro-mall-search/service"
)

type Server struct {
	proto_search.UnimplementedSearchRpcServer
	baseService service.BaseService
}

func (s *Server) ProductStatusUp(ctx context.Context, req *proto_search.ProductStatusUpRequest) (*proto_search.ProductStatusUpResponse, error) {
	return s.baseService.ProductStatusUp(ctx, req)
}

func (s *Server) SearchProduct(ctx context.Context, req *proto_search.SearchProductRequest) (*proto_search.SearchProductResponse, error) {
	return s.baseService.SearchProduct(ctx, req)
}

package server

import (
	"context"
	proto_coupon "mall-demo/micro-mall-coupon/proto/micro-mall-coupon-proto"
	"mall-demo/micro-mall-coupon/service"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	proto_coupon.UnimplementedCouponRpcServer
	skuService service.SpuService
}

func (s *Server) SaveSkuLadder(ctx context.Context, req *proto_coupon.SaveSkuLadderRequest) (*proto_coupon.SaveSkuLadderResponse, error) {
	return s.skuService.SaveSkuLadder(ctx, req)
}
func (s *Server) SaveSkuFullReduction(ctx context.Context, req *proto_coupon.SaveSkuFullReductionRequest) (*proto_coupon.SaveSkuFullReductionResponse, error) {
	return s.skuService.SaveSkuFullReduction(ctx, req)
}
func (s *Server) SaveSpuBounds(ctx context.Context, req *proto_coupon.SaveSpuBoundsRequest) (*proto_coupon.SaveSpuboundsResponse, error) {
	return s.skuService.SaveSpuBounds(ctx, req)
}
func (s *Server) SaveMemberPrice(ctx context.Context, req *proto_coupon.SaveMemberPriceRequest) (*proto_coupon.SaveMemberPriceResponse, error) {
	return s.skuService.SaveMemberPrice(ctx, req)
}

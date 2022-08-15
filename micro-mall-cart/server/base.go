package server

import (
	"context"
	proto_cart "mall-demo/micro-mall-cart/proto/micro-mall-cart-proto"
	"mall-demo/micro-mall-cart/service"
)

type Server struct {
	proto_cart.UnimplementedCartRpcServer
	*service.BaseService
}

func (s *Server) AddToCart(ctx context.Context, in *proto_cart.AddToCartRequest) (*proto_cart.AddToCartResponse, error) {
	return s.BaseService.AddToCart(ctx, in)

}
func (s *Server) GetCartItem(ctx context.Context, in *proto_cart.GetCartItemRequest) (*proto_cart.GetCartItemResponse, error) {
	return s.BaseService.GetCartItem(ctx, in)
}
func (s *Server) GetCart(ctx context.Context, in *proto_cart.GetCartRequest) (*proto_cart.GetCartResponse, error) {
	return s.BaseService.GetCart(ctx, in)
}
func (s *Server) ClearCartInfo(ctx context.Context, in *proto_cart.ClearCartInfoRequest) (*proto_cart.ClearCartInfoResponse, error) {
	return s.BaseService.ClearCartInfo(ctx, in)
}
func (s *Server) CheckItem(ctx context.Context, in *proto_cart.CheckItemRequest) (*proto_cart.CheckItemResponse, error) {
	return s.BaseService.CheckItem(ctx, in)
}
func (s *Server) ChangeItemCount(ctx context.Context, in *proto_cart.ChangeItemCountRequest) (*proto_cart.ChangeItemCountResponse, error) {
	return s.BaseService.ChangeItemCount(ctx, in)
}
func (s *Server) DeleteIdCartInfo(ctx context.Context, in *proto_cart.DeleteIdCartInfoRequest) (*proto_cart.DeleteIdCartInfoResponse, error) {
	return s.BaseService.DeleteIdCartInfo(ctx, in)
}
func (s *Server) GetUserCartItems(ctx context.Context, in *proto_cart.GetUserCartItemsRequest) (*proto_cart.GetUserCartItemsResponse, error) {
	return s.BaseService.GetUserCartItems(ctx, in)
}

package server

import (
	"context"
	proto_ware "mall-demo/micro-mall-ware/proto/micro-mall-ware-proto"
	"mall-demo/micro-mall-ware/service"
)

type Server struct {
	proto_ware.UnimplementedWareRpcServer
	wareService service.WareService
}

func (s *Server) GetWareInfoList(ctx context.Context, req *proto_ware.GetWareInfoListRequest) (*proto_ware.GetWareInfoListResponse, error) {
	return s.wareService.GetWareInfoList(ctx, req)
}
func (s *Server) SaveWare(ctx context.Context, req *proto_ware.SaveWareRequest) (*proto_ware.SaveWareResponse, error) {
	return s.wareService.SaveWare(ctx, req)
}

func (s *Server) GetSkuWareInfo(ctx context.Context, req *proto_ware.GetSkuWareInfoRequest) (*proto_ware.GetSkuWareInfoResponse, error) {
	return s.wareService.GetSkuWareInfo(ctx, req)
}
func (s *Server) GetPurchaseDetailInfo(ctx context.Context, req *proto_ware.GetPurchaseDetailInfoRequest) (*proto_ware.GetPurchaseDetailInfoResponse, error) {
	return s.wareService.GetPurchaseDetailInfo(ctx, req)
}
func (s *Server) SavePurchaseDetail(ctx context.Context, req *proto_ware.SavePurchaseDetailRequest) (*proto_ware.SavePurchaseDetailResponse, error) {
	return s.wareService.SavePurchaseDetail(ctx, req)
}

// purchase
func (s *Server) GetPurchaseList(ctx context.Context, req *proto_ware.GetPurchaseListRequest) (*proto_ware.GetPurchaseListResponse, error) {
	return s.wareService.GetPurchaseList(ctx, req)
}
func (s *Server) SavePurcase(ctx context.Context, req *proto_ware.SavePurchaseRequest) (*proto_ware.SavePurchaseResponse, error) {
	return s.wareService.SavePurcase(ctx, req)
}
func (s *Server) MergePurchase(ctx context.Context, req *proto_ware.MergePurchaseRequest) (*proto_ware.MergePurchaseResponse, error) {
	return s.wareService.MergePurchase(ctx, req)
}
func (s *Server) GetUnReceivedPurchaseInfo(ctx context.Context, req *proto_ware.GetUnReceivedPurchaseInfoRequest) (*proto_ware.GetUnReceivedPurchaseInfoResponse, error) {
	return s.wareService.GetUnReceivedPurchaseInfo(ctx, req)
}
func (s *Server) ReceivePurchase(ctx context.Context, req *proto_ware.ReceivePurchaseRequest) (*proto_ware.ReceivePurchaseResponse, error) {
	return s.wareService.ReceivePurchase(ctx, req)
}
func (s *Server) UpdatePurchase(ctx context.Context, req *proto_ware.UpdatePurchaseRequest) (*proto_ware.UpdatePurchaseResponse, error) {
	return s.wareService.UpdatePurchase(ctx, req)
}

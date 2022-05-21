package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"mall-demo/micro-mall-coupon/global"
	"mall-demo/micro-mall-coupon/model"
	proto_coupon "mall-demo/micro-mall-coupon/proto/micro-mall-coupon-proto"
)

type SpuService struct {
}

func (ss *SpuService) SaveSkuLadder(ctx context.Context, req *proto_coupon.SaveSkuLadderRequest) (*proto_coupon.SaveSkuLadderResponse, error) {
	var skuLadders []model.SmsSkuLadder
	skuLadders = make([]model.SmsSkuLadder, 0, 10)
	for i := 0; i < len(req.SkuLadderEntities); i++ {
		var skuLadder model.SmsSkuLadder
		copier.Copy(&skuLadder, req.SkuLadderEntities[i])
		skuLadders = append(skuLadders, skuLadder)
	}
	global.SmsMysqlConn.Save(&skuLadders)
	return &proto_coupon.SaveSkuLadderResponse{}, nil
}

func (ss *SpuService) SaveSkuFullReduction(ctx context.Context, req *proto_coupon.SaveSkuFullReductionRequest) (*proto_coupon.SaveSkuFullReductionResponse, error) {
	var skuFullReductions []model.SmsSkuFullReduction
	skuFullReductions = make([]model.SmsSkuFullReduction, 0, 10)
	for i := 0; i < len(req.SkuFullReductionEntities); i++ {
		var skuFullReduction model.SmsSkuFullReduction
		copier.Copy(&skuFullReduction, req.SkuFullReductionEntities[i])
		skuFullReductions = append(skuFullReductions, skuFullReduction)
	}
	global.SmsMysqlConn.Save(&skuFullReductions)
	return &proto_coupon.SaveSkuFullReductionResponse{}, nil
}

func (ss *SpuService) SaveSpuBounds(ctx context.Context, req *proto_coupon.SaveSpuBoundsRequest) (*proto_coupon.SaveSpuboundsResponse, error) {
	var spuBounds model.SmsSpuBounds
	copier.Copy(&spuBounds, req.SpuBoundEntities)
	global.SmsMysqlConn.Save(&spuBounds)
	return &proto_coupon.SaveSpuboundsResponse{}, nil
}

func (ss *SpuService) SaveMemberPrice(ctx context.Context, req *proto_coupon.SaveMemberPriceRequest) (*proto_coupon.SaveMemberPriceResponse, error) {
	memberPrices := make([]model.SmsMemberPrice, 0, 10)
	for i := 0; i < len(req.MemberPriceEntities); i++ {
		var memberPrice model.SmsMemberPrice
		_ = copier.Copy(&memberPrice, req.MemberPriceEntities[i])
		memberPrices = append(memberPrices, memberPrice)
	}

	fmt.Printf("%+v", memberPrices)

	global.SmsMysqlConn.Save(&memberPrices)
	return &proto_coupon.SaveMemberPriceResponse{}, nil
}

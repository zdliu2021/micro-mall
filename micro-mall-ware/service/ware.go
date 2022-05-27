package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"mall-demo/micro-mall-ware/global"
	"mall-demo/micro-mall-ware/model"
	proto_ware "mall-demo/micro-mall-ware/proto/micro-mall-ware-proto"
	"time"
)

type WareService struct {
}

func (ws *WareService) GetWareInfoList(ctx context.Context, req *proto_ware.GetWareInfoListRequest) (*proto_ware.GetWareInfoListResponse, error) {
	var wareInfos []model.WmsWareInfo
	global.WmsMysqlConn.Model(&model.WmsWareInfo{}).Offset(int((req.PageNum - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&wareInfos)
	resp := proto_ware.GetWareInfoListResponse{
		TotalCount: int32(len(wareInfos)),
		List:       make([]*proto_ware.GetWareInfoListResponse_List, 0, 10),
	}
	for i := 0; i < len(wareInfos); i++ {
		wareInfo := proto_ware.GetWareInfoListResponse_List{}
		copier.Copy(&wareInfo, &wareInfos[i])
		resp.List = append(resp.List, &wareInfo)
	}
	return &resp, nil
}
func (ws *WareService) SaveWare(ctx context.Context, req *proto_ware.SaveWareRequest) (*proto_ware.SaveWareResponse, error) {
	var ware model.WmsWareInfo
	copier.Copy(&ware, req)
	global.WmsMysqlConn.Create(&ware)
	return &proto_ware.SaveWareResponse{}, nil
}

func (ws *WareService) GetSkuWareInfo(ctx context.Context, req *proto_ware.GetSkuWareInfoRequest) (*proto_ware.GetSkuWareInfoResponse, error) {
	var wareSkus []model.WmsWareSku
	global.WmsMysqlConn.Model(&model.WmsWareSku{}).Offset(int((req.PageNum - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&wareSkus)
	resp := proto_ware.GetSkuWareInfoResponse{
		TotalCount: int32(len(wareSkus)),
		List:       make([]*proto_ware.GetSkuWareInfoResponse_List, 0, 10),
	}
	for i := 0; i < len(wareSkus); i++ {
		wareSku := proto_ware.GetSkuWareInfoResponse_List{}
		copier.Copy(&wareSku, &wareSkus[i])
		resp.List = append(resp.List, &wareSku)
	}
	return &resp, nil
}

func (ws *WareService) GetSkuHasStock(ctx context.Context, req *proto_ware.GetSkuHasStockRequest) (*proto_ware.GetSkuHasStockResponse, error) {
	var skuStocks []model.WmsWareSku
	global.WmsMysqlConn.Model(&model.WmsWareSku{}).Where("sku_id in ?", req.SkuIds).Find(&skuStocks)
	resp := proto_ware.GetSkuHasStockResponse{
		SkuHasStock: make(map[int64]bool),
	}
	for _, skuStock := range skuStocks {
		if skuStock.Stock-skuStock.StockLocked > 0 {
			resp.SkuHasStock[skuStock.SkuId] = true
		} else {
			resp.SkuHasStock[skuStock.SkuId] = false
		}
	}
	for _, skuId := range req.SkuIds {
		if _, ok := resp.SkuHasStock[skuId]; !ok {
			resp.SkuHasStock[skuId] = false
		}
	}
	return &resp, nil
}

func (ws *WareService) GetPurchaseDetailInfo(ctx context.Context, req *proto_ware.GetPurchaseDetailInfoRequest) (*proto_ware.GetPurchaseDetailInfoResponse, error) {
	var purchaseDetaileds []model.WmsPurchaseDetail
	global.WmsMysqlConn.Model(&model.WmsPurchaseDetail{}).Offset(int((req.PageNum-1)*req.PageSize)).Limit(int(req.PageSize)).Where("ware_id = ? and status = ?", req.WareId, req.Status).Find(&purchaseDetaileds)
	resp := proto_ware.GetPurchaseDetailInfoResponse{
		TotalCount: int32(len(purchaseDetaileds)),
		List:       make([]*proto_ware.GetPurchaseDetailInfoResponse_List, 0, 10),
	}
	for i := 0; i < len(purchaseDetaileds); i++ {
		purchaseDetailed := proto_ware.GetPurchaseDetailInfoResponse_List{}
		copier.Copy(&purchaseDetailed, &purchaseDetaileds[i])
		resp.List = append(resp.List, &purchaseDetailed)
	}
	return &resp, nil
}
func (ws *WareService) SavePurchaseDetail(ctx context.Context, req *proto_ware.SavePurchaseDetailRequest) (*proto_ware.SavePurchaseDetailResponse, error) {
	var purchaseDetail model.WmsPurchaseDetail
	copier.Copy(&purchaseDetail, req)
	fmt.Printf("%+v \n", purchaseDetail)
	global.WmsMysqlConn.Create(&purchaseDetail)
	return &proto_ware.SavePurchaseDetailResponse{}, nil
}

func (ws *WareService) SavePurcase(ctx context.Context, req *proto_ware.SavePurchaseRequest) (*proto_ware.SavePurchaseResponse, error) {
	var purchase model.WmsPurchase
	copier.Copy(&purchase, req)
	purchase.UpdateTime = time.Now()
	purchase.CreateTime = time.Now()

	global.WmsMysqlConn.Create(&purchase)
	return &proto_ware.SavePurchaseResponse{}, nil
}

func (ws *WareService) GetPurchaseList(ctx context.Context, req *proto_ware.GetPurchaseListRequest) (*proto_ware.GetPurchaseListResponse, error) {
	var res []model.WmsPurchase
	global.WmsMysqlConn.Model(&model.WmsPurchase{}).Offset(int((req.PageNum - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&res)
	var resp proto_ware.GetPurchaseListResponse
	copier.CopyWithOption(&resp.PurchaseEntities, &res, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	resp.TotalCount = int32(len(res))
	fmt.Printf("%+v \n", resp)
	return &resp, nil
}

func (ws *WareService) MergePurchase(ctx context.Context, req *proto_ware.MergePurchaseRequest) (*proto_ware.MergePurchaseResponse, error) {
	var purchaseDetaileds []model.WmsPurchaseDetail
	global.WmsMysqlConn.Model(&model.WmsPurchaseDetail{}).Where("id in ?", req.Items).Find(&purchaseDetaileds)

	var purchase model.WmsPurchase
	global.WmsMysqlConn.Model(&model.WmsPurchase{}).Where("id = ?", req.PurchaseId).First(&purchase)

	for i := 0; i < len(purchaseDetaileds); i++ {
		purchase.Amount += purchaseDetaileds[i].SkuPrice * float64(purchaseDetaileds[i].SkuNum)
		purchaseDetaileds[i].PurchaseId = req.PurchaseId
	}
	purchase.UpdateTime = time.Now()
	purchase.WareId = purchaseDetaileds[0].WareId

	global.WmsMysqlConn.Save(&purchase)
	global.WmsMysqlConn.Save(&purchaseDetaileds)
	return &proto_ware.MergePurchaseResponse{}, nil
}

func (ws *WareService) GetUnReceivedPurchaseInfo(ctx context.Context, req *proto_ware.GetUnReceivedPurchaseInfoRequest) (*proto_ware.GetUnReceivedPurchaseInfoResponse, error) {
	var purchaseInfos []model.WmsPurchase
	global.WmsMysqlConn.Model(&model.WmsPurchase{}).Where("status <= 1").Find(&purchaseInfos)
	resp := proto_ware.GetUnReceivedPurchaseInfoResponse{
		TotalCount: int32(len(purchaseInfos)),
		List:       make([]*proto_ware.GetUnReceivedPurchaseInfoResponse_List, 0, 10),
	}
	for i := 0; i < len(purchaseInfos); i++ {
		var purchaseInfo proto_ware.GetUnReceivedPurchaseInfoResponse_List
		copier.Copy(&purchaseInfo, &purchaseInfos[i])
		resp.List = append(resp.List, &purchaseInfo)
	}
	return &resp, nil
}

func (ws *WareService) ReceivePurchase(ctx context.Context, req *proto_ware.ReceivePurchaseRequest) (*proto_ware.ReceivePurchaseResponse, error) {
	//return &proto_ware.ReceivePurchaseResponse{}, nil
	var purchases []model.WmsPurchase
	global.WmsMysqlConn.Model(&model.WmsPurchase{}).Where("id in ?", req.Items).Find(&purchases)
	for i := 0; i < len(purchases); i++ {
		purchases[i].Status = 2
	}
	global.WmsMysqlConn.Save(&purchases)
	return &proto_ware.ReceivePurchaseResponse{}, nil
}
func (ws *WareService) UpdatePurchase(ctx context.Context, req *proto_ware.UpdatePurchaseRequest) (*proto_ware.UpdatePurchaseResponse, error) {
	var purchase model.WmsPurchase
	global.WmsMysqlConn.Model(&model.WmsPurchase{}).Where("id = ?", req.Id).First(&purchase)
	copier.Copy(&purchase, req)
	global.WmsMysqlConn.Save(&purchase)
	return &proto_ware.UpdatePurchaseResponse{}, nil
}

package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"mall-demo/micro-mall-api/global"
	"mall-demo/micro-mall-api/model/request"
	response2 "mall-demo/micro-mall-api/model/response"
	"mall-demo/micro-mall-api/proto/micro-mall-ware-proto"
	"mall-demo/micro-mall-api/rpc-client"
	"strconv"
)

func GetWareInfoList(ctx *gin.Context) {
	rpcClient := rpc_client.GetWareRpcClient()
	page := GetPage(ctx)
	resp, err := rpcClient.GetWareInfoList(context.TODO(), &proto_ware.GetWareInfoListRequest{
		PageSize: int32(page.PageSize),
		PageNum:  int32(page.PageNum),
		Keyword:  page.Keyword,
		Order:    page.Order,
		Sidx:     page.Sidx,
	})
	fmt.Printf(" data: %+v \n", resp)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	var res response2.GetWareInfoListResponse
	copier.CopyWithOption(&res, resp, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	response2.OkWithData(res, ctx)
}

func SaveWare(ctx *gin.Context) {
	var req request.SaveWareRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetWareRpcClient()
	var rpcReq proto_ware.SaveWareRequest
	copier.Copy(&rpcReq, &req)
	_, err := rpcClient.SaveWare(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}
func GetPurchaseList(ctx *gin.Context) {
	page := GetPage(ctx)
	var req proto_ware.GetPurchaseListRequest
	copier.Copy(&req, page)
	rpcClient := rpc_client.GetWareRpcClient()
	resp, err := rpcClient.GetPurchaseList(context.TODO(), &req)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res response2.GetPurchaseListResponse
	copier.CopyWithOption(&res.List, &resp.PurchaseEntities, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	res.TotalCount = resp.TotalCount
	response2.OkWithData(res, ctx)
}

func SavePurchase(ctx *gin.Context) {
	var req request.SavePurchaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetWareRpcClient()
	var rpcReq proto_ware.SavePurchaseRequest
	copier.Copy(&rpcReq, &req)
	_, err := rpcClient.SavePurcase(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}
func SavePurchaseDetail(ctx *gin.Context) {
	var req request.SavePurchaseDetailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	var rpcReq proto_ware.SavePurchaseDetailRequest
	copier.Copy(&rpcReq, &req)
	rpcClient := rpc_client.GetWareRpcClient()
	_, err := rpcClient.SavePurchaseDetail(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func GetSkuWareInfo(ctx *gin.Context) {
	rpcClient := rpc_client.GetWareRpcClient()
	page := GetPage(ctx)
	wareId, _ := strconv.ParseInt(ctx.Query("wareId"), 10, 64)
	skuId, _ := strconv.ParseInt(ctx.Query("skuId"), 10, 64)
	resp, err := rpcClient.GetSkuWareInfo(context.TODO(), &proto_ware.GetSkuWareInfoRequest{
		PageSize: int32(page.PageSize),
		PageNum:  int32(page.PageNum),
		Keyword:  page.Keyword,
		Order:    page.Order,
		Sidx:     page.Sidx,
		WareId:   wareId,
		SkuId:    skuId,
	})

	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res response2.GetSkuWareInfoResponse
	copier.CopyWithOption(&res, resp, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	response2.OkWithData(res, ctx)
}

func GetPurchaseDetailedInfo(ctx *gin.Context) {
	rpcClient := rpc_client.GetWareRpcClient()
	page := GetPage(ctx)
	wareId, _ := strconv.ParseInt(ctx.Query("wareId"), 10, 64)
	status, _ := strconv.ParseInt(ctx.Query("status"), 10, 32)
	resp, err := rpcClient.GetPurchaseDetailInfo(context.TODO(), &proto_ware.GetPurchaseDetailInfoRequest{
		PageSize: int32(page.PageSize),
		PageNum:  int32(page.PageNum),
		Keyword:  page.Keyword,
		Order:    page.Order,
		Sidx:     page.Sidx,
		WareId:   wareId,
		Status:   int32(status),
	})

	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res response2.GetPurchaseDetailedInfo
	copier.CopyWithOption(&res, resp, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	response2.OkWithData(res, ctx)

}

func MergePurchase(ctx *gin.Context) {
	rpcClient := rpc_client.GetWareRpcClient()
	var req request.MergePurchaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	_, err := rpcClient.MergePurchase(context.TODO(), &proto_ware.MergePurchaseRequest{
		PurchaseId: req.PurchaseId,
		Items:      req.Items,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func GetUnReceivedPurchaseInfo(ctx *gin.Context) {
	rpcClient := rpc_client.GetWareRpcClient()
	resp, err := rpcClient.GetUnReceivedPurchaseInfo(context.TODO(), &proto_ware.GetUnReceivedPurchaseInfoRequest{})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res response2.GetUnReceivedPurchaseInfo
	copier.CopyWithOption(&res, resp, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	response2.OkWithData(res, ctx)
}

func ReceivePurchase(ctx *gin.Context) {
	rpcClient := rpc_client.GetWareRpcClient()
	var req request.ReceivePurchaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	_, err := rpcClient.ReceivePurchase(context.TODO(), &proto_ware.ReceivePurchaseRequest{
		Items: req.Ids,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func UpdatePurchase(ctx *gin.Context) {
	var req request.UpdatePurchaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetWareRpcClient()
	var rpcReq proto_ware.UpdatePurchaseRequest
	copier.Copy(&rpcReq, &req)
	_, err := rpcClient.UpdatePurchase(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

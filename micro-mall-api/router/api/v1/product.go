package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"mall-demo/micro-mall-api/global"
	request2 "mall-demo/micro-mall-api/model/request"
	response2 "mall-demo/micro-mall-api/model/response"
	proto_product "mall-demo/micro-mall-api/proto/micro-mall-product-proto"
	"mall-demo/micro-mall-api/rpc-client"
	"mall-demo/micro-mall-api/utils"
	"strconv"
)

func GetPage(ctx *gin.Context) *request2.Page {
	var page request2.Page
	page.PageNum, _ = strconv.Atoi(ctx.DefaultQuery("page", "0"))
	page.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page.Keyword = ctx.DefaultQuery("key", "")
	page.Order = ctx.DefaultQuery("micro-mall-micro-mall-order-proto", "asc")
	page.Sidx = ctx.DefaultQuery("sidx", "")

	return &page
}

func ListCategoryTree(ctx *gin.Context) {
	var solve func(*response2.ListCategoryTreeResponse, *proto_product.CategoryEntity)
	solve = func(r1 *response2.ListCategoryTreeResponse, r2 *proto_product.CategoryEntity) {
		r1.CatId = r2.CatId
		r1.Name = r2.Name
		r1.ShowStatus = r2.ShowStatus
		r1.CatLevel = r2.CatLevel
		r1.Icon = r2.Icon
		r1.ProductCount = r2.ProductCount
		r1.Sort = r2.Sort
		r1.ParentCid = r2.ParentCid
		r1.Children = make([]*response2.ListCategoryTreeResponse, 0, 10)
		for i := 0; i < len(r2.Children); i++ {
			var tmp response2.ListCategoryTreeResponse
			solve(&tmp, r2.Children[i])
			r1.Children = append(r1.Children, &tmp)
		}
	}

	rpcClient := rpc_client.GetProductClient()
	req := &proto_product.ListCategoryTreeRequest{}
	resp, err := rpcClient.ListCategoryTree(context.TODO(), req)
	if err != nil {
		global.GVA_LOG.Error("listcategorytree error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res []response2.ListCategoryTreeResponse
	res = make([]response2.ListCategoryTreeResponse, 0, 10)
	if len(resp.CategoryEntities) != 0 {
		for i := 0; i < len(resp.CategoryEntities); i++ {
			var tmp response2.ListCategoryTreeResponse
			solve(&tmp, resp.CategoryEntities[i])
			res = append(res, tmp)
		}
	}
	response2.OkWithData(res, ctx)
}

func DeleteCategory(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()
	var req proto_product.DeleteCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	fmt.Println(req)
	resp, err := rpcClient.DeleteCategory(context.TODO(), &req)
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	response2.OkWithData(resp, ctx)
}

func SaveCategory(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()
	var req request2.SaveCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var rpcReq proto_product.SaveCategoryRequest
	utils.Assignment(&rpcReq, &req)
	resp, err := rpcClient.SaveCategory(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	response2.OkWithData(resp, ctx)
}

func GetCategoryInfo(ctx *gin.Context) {
	cat_id := ctx.Param("cat_id")

	rpcClient := rpc_client.GetProductClient()
	var req proto_product.GetCategoryInfoRequest
	req.CatId, _ = strconv.ParseInt(cat_id, 10, 64)

	resp, err := rpcClient.GetCategoryInfo(context.TODO(), &req)
	fmt.Println(resp)
	if err != nil {
		global.GVA_LOG.Error("get category info error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	res := response2.GetCategoryInfoResponse{
		CatId:        resp.CatId,
		Name:         resp.Name,
		ParentCid:    resp.ParentCid,
		CatLevel:     resp.CatLevel,
		ShowStatus:   resp.ShowStatus,
		Sort:         resp.Sort,
		Icon:         resp.Icon,
		ProductUnit:  resp.ProductUnit,
		ProductCount: resp.ProductCount,
	}

	response2.OkWithData(res, ctx)

}

func UpdateCategory(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()
	var req request2.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcReq := proto_product.UpdateCategoryRequest{
		CatId:       req.CatID,
		Name:        req.Name,
		ParentCid:   req.ParentCid,
		CatLevel:    req.CatLevel,
		ShowStatus:  req.ShowStatus,
		Sort:        req.Sort,
		Icon:        req.Icon,
		ProductUnit: req.ProductUnit,
	}
	resp, err := rpcClient.UpdateCategory(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	response2.OkWithData(resp, ctx)
}

func GetBrandList(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()

	var page request2.Page
	page.PageNum, _ = strconv.Atoi(ctx.DefaultQuery("page", "0"))
	page.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page.Keyword = ctx.DefaultQuery("key", "")
	page.Order = ctx.DefaultQuery("micro-mall-micro-mall-order-proto", "asc")
	page.Sidx = ctx.DefaultQuery("sidx", "")

	rpcReq := proto_product.GetBrandListRequest{
		PageSize: int32(page.PageSize),
		PageNum:  int32(page.PageNum),
		Keyword:  page.Keyword,
	}
	resp, err := rpcClient.GetBrandList(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res = make([]response2.BrandEntity, 0, 10)
	for i := 0; i < len(resp.BrandEntities); i++ {
		res = append(res, response2.BrandEntity{
			BrandId:     resp.BrandEntities[i].BrandId,
			Name:        resp.BrandEntities[i].Name,
			Logo:        resp.BrandEntities[i].Logo,
			Descript:    resp.BrandEntities[i].Logo,
			ShowStatus:  int32(resp.BrandEntities[i].ShowStatus),
			FirstLetter: resp.BrandEntities[i].FirstLetter,
			Sort:        int32(resp.BrandEntities[i].Sort),
		})
	}
	response2.OkWithData(response2.GetBrandListResponse{List: res, TotalCount: int32(len(res))}, ctx)
}

func GetBrandInfo(ctx *gin.Context) {
	brandId, _ := strconv.ParseInt(ctx.Param("brand_id"), 10, 64)
	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.GetBrandInfo(context.TODO(), &proto_product.GetBrandInfoRequest{
		BrandId: brandId,
	})
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	res := response2.GetBrandInfoRequest{
		Brand: response2.BrandEntity{
			BrandId:     resp.BrandEntity.BrandId,
			Name:        resp.BrandEntity.Name,
			Logo:        resp.BrandEntity.Logo,
			Descript:    resp.BrandEntity.Logo,
			ShowStatus:  int32(resp.BrandEntity.ShowStatus),
			FirstLetter: resp.BrandEntity.FirstLetter,
			Sort:        int32(resp.BrandEntity.Sort),
		},
	}
	response2.OkWithData(res, ctx)
}

func UpdateBrand(ctx *gin.Context) {
	var brand request2.UpdateBrandRequest
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.UpdateBrand(context.TODO(), &proto_product.UpdateBrandRequest{
		BrandEntity: &proto_product.BrandEntity{
			BrandId:     brand.BrandId,
			Name:        brand.Name,
			Logo:        brand.Logo,
			Descript:    brand.Logo,
			ShowStatus:  int32(brand.ShowStatus),
			FirstLetter: brand.FirstLetter,
			Sort:        int32(brand.Sort),
		},
	})
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func SaveBrand(ctx *gin.Context) {
	var brand request2.SaveBrandRequest
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.SaveBrand(context.TODO(), &proto_product.SaveBrandRequest{
		BrandEntity: &proto_product.BrandEntity{
			BrandId:     brand.BrandId,
			Name:        brand.Name,
			Logo:        brand.Logo,
			Descript:    brand.Logo,
			ShowStatus:  int32(brand.ShowStatus),
			FirstLetter: brand.FirstLetter,
			Sort:        int32(brand.Sort),
		},
	})
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func DeleteBrand(ctx *gin.Context) {
	var req request2.DeleteBrandRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetProductClient()

	_, err := rpcClient.DeleteBrand(context.TODO(), &proto_product.DeleteBrandRequest{
		BrandIds: req.BrandIds,
	})
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func GetBrandRelationedCateLog(ctx *gin.Context) {
	brandId, _ := strconv.ParseInt(ctx.Query("brandId"), 10, 64)
	rpcClient := rpc_client.GetProductClient()

	resp, err := rpcClient.GetBrandRelatedCateLog(context.TODO(), &proto_product.GetBrandRelatedCateLogRequest{
		BrandId: brandId,
	})
	if err != nil {
		global.GVA_LOG.Error("get brand related catelog error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res []response2.GetBrandRelationedCateLogResponse
	res = make([]response2.GetBrandRelationedCateLogResponse, 0, 10)
	for i := 0; i < len(resp.BrandRelatedCateLogEntities); i++ {
		res = append(res, response2.GetBrandRelationedCateLogResponse{
			CateLogId:   resp.BrandRelatedCateLogEntities[i].CatelogId,
			CateLogName: resp.BrandRelatedCateLogEntities[i].CatlogName,
			Id:          resp.BrandRelatedCateLogEntities[i].Id,
			BrandName:   resp.BrandRelatedCateLogEntities[i].BrandName,
			BrandId:     resp.BrandRelatedCateLogEntities[i].BrandId,
		})
	}

	response2.OkWithData(res, ctx)
}

func GetBrandRelatedWithCatelog(ctx *gin.Context) {
	cat_id, _ := strconv.ParseInt(ctx.Query("cat_id"), 10, 64)

	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.GetBrandRelatedWithCatelog(context.TODO(), &proto_product.GetBrandRelatedWithCatelogRequest{
		CatId: cat_id,
	})
	if err != nil {
		global.GVA_LOG.Error("delete category error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res []response2.BrandCatelogRelation
	res = make([]response2.BrandCatelogRelation, 0, 10)
	for i := 0; i < len(resp.BrandCatelogRelations); i++ {
		res = append(res, response2.BrandCatelogRelation{
			BrandName: resp.BrandCatelogRelations[i].BrandName,
			BrandId:   resp.BrandCatelogRelations[i].BrandId,
		})
	}

	response2.OkWithData(res, ctx)
}

func SaveBrandCatelogRelation(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()
	var req request2.SaveBrandCatelogRelationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	_, err := rpcClient.SaveBrandCatelogRelation(context.TODO(), &proto_product.SaveBrandCatelogRelationRequest{
		BrandId:   req.BrandId,
		CatelogId: req.CateLogId,
	})

	if err != nil {
		global.GVA_LOG.Error("save brand catelog relation error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func DeleteBrandCatelogRelation(ctx *gin.Context) {
	var req proto_product.DeleteBrandCatelogRelationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.DeleteBrandCatelogRelation(context.TODO(), &req)
	if err != nil {
		global.GVA_LOG.Error("save brand catelog relation error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func ListCatelogAttrGroup(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()
	catelogId, _ := strconv.ParseInt(ctx.Param("catelogId"), 10, 64)

	page := GetPage(ctx)
	resp, err := rpcClient.ListCatelogAttrGroup(context.TODO(), &proto_product.ListCatelogAttrGroupRequest{
		CatelogId: catelogId,
		PageSize:  int32(page.PageSize),
		PageNum:   int32(page.PageNum),
		Key:       page.Keyword,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res []response2.ListCatelogAttrGroupResponse
	res = make([]response2.ListCatelogAttrGroupResponse, 0, 10)
	for i := 0; i < len(resp.CatelogAttrGroupEntities); i++ {
		res = append(res, response2.ListCatelogAttrGroupResponse{
			CatelogId:     resp.CatelogAttrGroupEntities[i].CatelogId,
			AttrGroupName: resp.CatelogAttrGroupEntities[i].AttrGroupName,
			AttrGroupId:   resp.CatelogAttrGroupEntities[i].AttrGroupId,
			Descript:      resp.CatelogAttrGroupEntities[i].Descript,
			Icon:          resp.CatelogAttrGroupEntities[i].Icon,
			Sort:          resp.CatelogAttrGroupEntities[i].Sort,
			CatelogPath:   resp.CatelogAttrGroupEntities[i].CatelogPath,
		})
	}
	response2.OkWithData(map[string]interface{}{
		"totalCount": resp.TotalCount,
		"paegSize":   resp.PageSize,
		"totalPage":  resp.TotalPage,
		"currPage":   resp.CurrPage,
		"list":       res,
	}, ctx)
}

func GetAttrListRelatedWithCatelog(ctx *gin.Context) {
	catelogId, _ := strconv.ParseInt(ctx.Param("catelogId"), 10, 64)
	page := GetPage(ctx)
	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.GetAttrListRelatedWithCatelog(context.TODO(), &proto_product.GetAttrListRelatedWithCatelogRequest{
		CatelogId: catelogId,
		Page:      int32(page.PageNum),
		Limit:     int32(page.PageSize),
		Sidx:      page.Sidx,
		Order:     page.Order,
		Key:       page.Keyword,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	res := make([]response2.AttrEntity, 0, 10)
	for i := 0; i < len(resp.AttrEntities); i++ {
		res = append(res, response2.AttrEntity{
			AttrId:      resp.AttrEntities[i].AttrId,
			AttrName:    resp.AttrEntities[i].AttrName,
			SearchType:  int32(resp.AttrEntities[i].SearchType),
			ValueSelect: resp.AttrEntities[i].ValueSelect,
			ValueType:   int32(resp.AttrEntities[i].ValueType),
			Icon:        resp.AttrEntities[i].Icon,
			AttrType:    int32(resp.AttrEntities[i].AttrType),
			Enable:      int32(resp.AttrEntities[i].Enable),
			CatelogName: resp.AttrEntities[i].CatelogName,
			GroupName:   resp.AttrEntities[i].GroupName,
		})
	}
	response2.OkWithData(map[string]interface{}{
		"totalCount": resp.TotalCount,
		"paegSize":   resp.PageSize,
		"totalPage":  resp.TotalPage,
		"currPage":   resp.CurrPage,
		"list":       res,
	}, ctx)
}

func GetSaleAttrListRelatedWithCatelog(ctx *gin.Context) {
	catelogId, _ := strconv.ParseInt(ctx.Param("catelogId"), 10, 64)
	page := GetPage(ctx)
	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.GetAttrListRelatedWithCatelog(context.TODO(), &proto_product.GetAttrListRelatedWithCatelogRequest{
		CatelogId: catelogId,
		Page:      int32(page.PageNum),
		Limit:     int32(page.PageSize),
		Sidx:      page.Sidx,
		Order:     page.Order,
		Key:       page.Keyword,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	res := make([]response2.AttrEntity, 0, 10)
	for i := 0; i < len(resp.AttrEntities); i++ {
		res = append(res, response2.AttrEntity{
			AttrId:      resp.AttrEntities[i].AttrId,
			AttrName:    resp.AttrEntities[i].AttrName,
			SearchType:  int32(resp.AttrEntities[i].SearchType),
			ValueSelect: resp.AttrEntities[i].ValueSelect,
			ValueType:   int32(resp.AttrEntities[i].ValueType),
			Icon:        resp.AttrEntities[i].Icon,
			AttrType:    int32(resp.AttrEntities[i].AttrType),
			Enable:      int32(resp.AttrEntities[i].Enable),
			CatelogName: resp.AttrEntities[i].CatelogName,
			GroupName:   resp.AttrEntities[i].GroupName,
		})
	}
	response2.OkWithData(map[string]interface{}{
		"totalCount": resp.TotalCount,
		"paegSize":   resp.PageSize,
		"totalPage":  resp.TotalPage,
		"currPage":   resp.CurrPage,
		"list":       res,
	}, ctx)
}

func ListSpu(ctx *gin.Context) {
	catelogId, _ := strconv.ParseInt(ctx.Query("catelogId"), 10, 64)
	brandId, _ := strconv.ParseInt(ctx.Query("brandId"), 10, 64)
	status, _ := strconv.ParseInt(ctx.Query("status"), 10, 32)
	page := GetPage(ctx)
	rpcClient := rpc_client.GetProductClient()
	rpcReq := proto_product.SearchSpuInfoRequest{
		PageNum:   int32(page.PageNum),
		PageSize:  int32(page.PageSize),
		Keyword:   page.Keyword,
		Order:     page.Order,
		Sidx:      page.Sidx,
		CatalogId: catelogId,
		BrandId:   brandId,
		Status:    int32(status),
	}
	resp, err := rpcClient.SearchSpuInfo(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res response2.GetSpuInfoResponse
	copier.Copy(&res, &resp)
	copier.CopyWithOption(&res.List, resp.SpuInfoEntities, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	response2.OkWithData(res, ctx)
}

func GetAttrGroupInfo(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()

	attrGroupID, _ := strconv.ParseInt(ctx.Param("attrGroupID"), 10, 64)
	resp, err := rpcClient.GetAttrGroupInfo(context.TODO(), &proto_product.GetAttrGroupInfoRequest{
		AttrGroupId: attrGroupID,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	res := response2.GetAttrGroupInfoResponse{
		AttrGroupId:   resp.CatelogAttrGroupEntity.AttrGroupId,
		AttrGroupName: resp.CatelogAttrGroupEntity.AttrGroupName,
		Sort:          int32(resp.CatelogAttrGroupEntity.Sort),
		Descript:      resp.CatelogAttrGroupEntity.Descript,
		Icon:          resp.CatelogAttrGroupEntity.Icon,
		CatelogId:     resp.CatelogAttrGroupEntity.CatelogId,
		CatelogPath:   resp.CatelogAttrGroupEntity.CatelogPath,
	}
	response2.OkWithData(res, ctx)

}

func GetAttrInfo(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()

	attrId, _ := strconv.ParseInt(ctx.Param("attrId"), 10, 64)
	resp, err := rpcClient.GetAttrInfo(context.TODO(), &proto_product.GetAttrInfoRequest{
		AttrId: attrId,
	})

	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	res := response2.GetAttrInfoResponse{
		AttrId:      resp.AttrId,
		AttrName:    resp.AttrName,
		SearchType:  int32(resp.SearchType),
		ValueType:   int32(resp.ValueType),
		Icon:        resp.Icon,
		ValueSelect: resp.ValueSelect,
		AttrType:    int32(resp.AttrType),
		Enable:      int32(resp.Enable),
		CatelogId:   resp.CatelogId,
		ShowDesc:    int32(resp.ShowDesc),
		AttrGroupId: int32(resp.AttrGroupId),
		CatelogPath: resp.CatelogPath,
	}
	response2.OkWithData(res, ctx)
}

func SaveAttrGroup(ctx *gin.Context) {
	var req request2.AttrGroupEntity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.SaveAttrGroup(context.TODO(), &proto_product.SaveAttrGroupRequest{
		AttrGroupId:   req.AttrGroupId,
		AttrGroupName: req.AttrGroupName,
		Sort:          int32(req.Sort),
		Descript:      req.Descript,
		Icon:          req.Icon,
		CatelogId:     req.CatelogId,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func UpdateAttrGroup(ctx *gin.Context) {
	var req request2.AttrGroupEntity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.UpdateAttrGroup(context.TODO(), &proto_product.UpdateAttrGroupRequest{
		AttrGroupId:   req.AttrGroupId,
		AttrGroupName: req.AttrGroupName,
		Sort:          int32(req.Sort),
		Descript:      req.Descript,
		Icon:          req.Icon,
		CatelogId:     req.CatelogId,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func DeleteAttrGroup(ctx *gin.Context) {
	var req request2.DeleteAttrGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.DeleteAttrGroup(context.TODO(), &proto_product.DeleteAttrGroupRequest{
		AttrGroupId: req.AttrGroupIds,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func GetAttrNotCorrelation(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()
	attrgroupId, _ := strconv.ParseInt(ctx.Param("attrgroupId"), 10, 64)

	page := GetPage(ctx)
	resp, err := rpcClient.GetAttrNotCorrelation(context.TODO(), &proto_product.GetAttrNotCorrelationRequest{
		AttrGroupId: attrgroupId,
		Page:        int32(page.PageNum),
		Limit:       int32(page.PageSize),
		Sidx:        page.Sidx,
		Order:       page.Order,
		Key:         page.Keyword,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	var res []response2.GetAttrNotCorrelationResponse
	res = make([]response2.GetAttrNotCorrelationResponse, 0, 10)
	for i := 0; i < len(resp.PmsAttrEntities); i++ {
		res = append(res, response2.GetAttrNotCorrelationResponse{
			AttrId:      resp.PmsAttrEntities[i].AttrId,
			AttrName:    resp.PmsAttrEntities[i].AttrName,
			SearchType:  int32(resp.PmsAttrEntities[i].SearchType),
			ValueType:   int32(resp.PmsAttrEntities[i].ValueType),
			Icon:        resp.PmsAttrEntities[i].Icon,
			ValueSelect: resp.PmsAttrEntities[i].ValueSelect,
			AttrType:    int32(resp.PmsAttrEntities[i].AttrType),
			Enable:      int32(resp.PmsAttrEntities[i].Enable),
			CatelogId:   resp.PmsAttrEntities[i].CatelogId,
			ShowDesc:    int32(resp.PmsAttrEntities[i].ShowDesc),
		})
	}
	response2.OkWithData(map[string]interface{}{
		"list":       res,
		"totalCount": resp.TotalCount,
		"pageSize":   resp.PageSize,
		"totalPage":  resp.TotalPage,
		"currPage":   resp.CurrPage,
	}, ctx)
}

func DeleteAttrAttrGroupRelation(ctx *gin.Context) {
	rpcClient := rpc_client.GetProductClient()

	var req []request2.DeleteAttrAttrGroupRelationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcReq := make([]*proto_product.AttrAttrGroupRelationEntity, 0, 10)
	for i := 0; i < len(req); i++ {
		rpcReq = append(rpcReq, &proto_product.AttrAttrGroupRelationEntity{
			AttrGroupId: req[i].AttrGroupID,
			AttrId:      req[i].AttrId,
		})
	}

	_, err := rpcClient.DeleteAttrAttrGroupRelation(context.TODO(), &proto_product.DeleteAttrAttrGroupRelationRequest{
		AttrAttrGroupRelationEntities: rpcReq,
	})

	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func SaveAttrAttrGroupRelation(ctx *gin.Context) {
	var req []request2.SaveAttrAttrGroupRelationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetProductClient()

	rpcReq := make([]*proto_product.AttrAttrGroupRelationEntity, 0, 10)
	for i := 0; i < len(req); i++ {
		rpcReq = append(rpcReq, &proto_product.AttrAttrGroupRelationEntity{
			AttrGroupId: req[i].AttrGroupID,
			AttrId:      req[i].AttrId,
		})
	}

	_, err := rpcClient.SaveAttrAttrGroupRelation(context.TODO(), &proto_product.SaveAttrAttrGroupRelationRequest{
		AttrAttrGroupRelationEntities: rpcReq,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}
func GetAttrRelatedAttrGroup(ctx *gin.Context) {
	attrGroupId, _ := strconv.ParseInt(ctx.Param("attrGroupId"), 10, 64)
	rpcClient := rpc_client.GetProductClient()

	resp, err := rpcClient.GetAttrRelatedAttrGroup(context.TODO(), &proto_product.GetAttrRelatedAttrGroupRequest{
		AttrGroupId: attrGroupId,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	var res []response2.GetAttrRelatedAttrGroupResponse
	for i := 0; i < len(resp.PmsAttrEntities); i++ {
		res = append(res, response2.GetAttrRelatedAttrGroupResponse{
			AttrId:      resp.PmsAttrEntities[i].AttrId,
			AttrName:    resp.PmsAttrEntities[i].AttrName,
			SearchType:  int32(resp.PmsAttrEntities[i].SearchType),
			ValueType:   int32(resp.PmsAttrEntities[i].ValueType),
			Icon:        resp.PmsAttrEntities[i].Icon,
			ValueSelect: resp.PmsAttrEntities[i].ValueSelect,
			AttrType:    int32(resp.PmsAttrEntities[i].AttrType),
			Enable:      int32(resp.PmsAttrEntities[i].Enable),
			CatelogId:   resp.PmsAttrEntities[i].CatelogId,
			ShowDesc:    int32(resp.PmsAttrEntities[i].ShowDesc),
		})
	}
	fmt.Printf("%+v \n", res)

	response2.OkWithData(res, ctx)
}

func GetAllGroupAndAttrRelatedWithCatelog(ctx *gin.Context) {
	catelogId, _ := strconv.ParseInt(ctx.Param("catelogId"), 10, 64)
	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.GetAllGroupAndAttrRelatedWithCatelog(context.TODO(), &proto_product.GetAllGroupAndAttrRelatedWithCatelogRequest{
		CatId: catelogId,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	res := make([]response2.AttrGroupAndAttrsEntity, 0, 10)

	for i := 0; i < len(resp.AttrGroupAndAtrrsEntities); i++ {
		attrs := make([]response2.PmsAttrEntity, 0, 10)
		for j := 0; j < len(resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities); j++ {
			attrs = append(attrs, response2.PmsAttrEntity{
				AttrId:      resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].AttrId,
				AttrName:    resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].AttrName,
				SearchType:  int32(resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].SearchType),
				ValueType:   int32(resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].ValueType),
				Icon:        resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].Icon,
				ValueSelect: resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].ValueSelect,
				AttrType:    int32(resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].AttrType),
				Enable:      int32(resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].Enable),
				CatelogId:   resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].CatelogId,
				ShowDesc:    int32(resp.AttrGroupAndAtrrsEntities[i].PmsAttrEntities[j].ShowDesc),
			})
		}
		res = append(res, response2.AttrGroupAndAttrsEntity{
			AttrGroupId:   resp.AttrGroupAndAtrrsEntities[i].AttrGroupId,
			AttrGroupName: resp.AttrGroupAndAtrrsEntities[i].AttrGroupName,
			Sort:          int32(resp.AttrGroupAndAtrrsEntities[i].Sort),
			Descript:      resp.AttrGroupAndAtrrsEntities[i].Descript,
			Icon:          resp.AttrGroupAndAtrrsEntities[i].Icon,
			CatelogId:     resp.AttrGroupAndAtrrsEntities[i].CatelogId,
			Attrs:         attrs,
		})
	}
	response2.OkWithData(res, ctx)
}

func UpdateAttr(ctx *gin.Context) {
	var req request2.AttrEntity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.UpdateAttr(context.TODO(), &proto_product.UpdateAttrRequest{
		AttrId:      req.AttrId,
		AttrName:    req.AttrName,
		SearchType:  int32(req.SearchType),
		ValueType:   int32(req.ValueType),
		Icon:        req.Icon,
		ValueSelect: req.ValueSelect,
		AttrType:    int32(req.AttrType),
		Enable:      int32(req.Enable),
		CatelogId:   req.CatelogId,
		ShowDesc:    int32(req.ShowDesc),
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func DeleteAttr(ctx *gin.Context) {
	var req proto_product.DeleteAttrRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.DeleteAttr(context.TODO(), &req)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func GetSpuAttrsInfo(ctx *gin.Context) {
	spuId, _ := strconv.ParseInt(ctx.Query("spuId"), 10, 64)
	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.GetSpuInfo(context.TODO(), &proto_product.GetSpuInfoRequest{
		SpuId: spuId,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	res := make([]response2.SpuAttr, 0, 10)
	copier.CopyWithOption(&res, &resp.Data, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	response2.OkWithData(res, ctx)
}

func UpdateSpuAttrs(ctx *gin.Context) {
	var req []request2.UpdateSpuAttrRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcReq := &proto_product.UpdateSpuAttrsRequest{}
	copier.CopyWithOption(rpcReq.Data, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.UpdateSpuAttrs(context.TODO(), rpcReq)
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func SaveProductAttr(ctx *gin.Context) {
	var req request2.AttrEntity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.SaveProductAttr(context.TODO(), &proto_product.SaveProductAttrRequest{
		AttrName:    req.AttrName,
		SearchType:  int32(req.SearchType),
		ValueType:   int32(req.ValueType),
		Icon:        req.Icon,
		ValueSelect: req.ValueSelect,
		AttrType:    int32(req.AttrType),
		Enable:      int32(req.Enable),
		CatelogId:   req.CatelogId,
		ShowDesc:    int32(req.ShowDesc),
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func ListSku(ctx *gin.Context) {
	page := GetPage(ctx)
	catelogId, _ := strconv.ParseInt(ctx.Query("catelogId"), 10, 64)
	brandId, _ := strconv.ParseInt(ctx.Query("brandId"), 10, 64)
	min, _ := strconv.ParseInt(ctx.Query("min"), 10, 32)
	max, _ := strconv.ParseInt(ctx.Query("max"), 10, 32)
	rpcClient := rpc_client.GetProductClient()
	resp, err := rpcClient.SearchSkuInfo(context.TODO(), &proto_product.SearchSkuInfoRequest{
		PageSize:  int32(page.PageSize),
		PageNum:   int32(page.PageNum),
		Keyword:   page.Keyword,
		Order:     page.Order,
		Sidx:      page.Sidx,
		Min:       int32(min),
		Max:       int32(max),
		CatelogId: catelogId,
		BrandId:   brandId,
	})

	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res response2.ListSkuResponse
	copier.CopyWithOption(&res, resp, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	response2.OkWithData(res, ctx)
}

func UpSpu(ctx *gin.Context) {
	spuId, _ := strconv.ParseInt(ctx.Param("spuId"), 10, 64)
	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.UpSpu(context.TODO(), &proto_product.UpSpuRequest{
		SpuId: spuId,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func SaveSpu(ctx *gin.Context) {
	var req request2.SaveSPU
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	//fmt.Printf("%+v \n\n\n", req)

	rpcReq := proto_product.SaveSpuRequest{
		SpuName:        req.SpuName,
		SpuDescription: req.SpuDescription,
		CatalogId:      req.CatalogID,
		BrandId:        req.BrandID,
		Weight:         req.Weight,
		PublishStatus:  req.PublishStatus,
		Decript:        req.Decript,
		Images:         req.Images,
		Bounds: &proto_product.SaveSpuRequest_Bounds{
			BuyBounds:  req.Bounds.BuyBounds,
			GrowBounds: req.Bounds.GrowBounds,
		},
		BaseAttrs: make([]*proto_product.SaveSpuRequest_Baseattrs, 0, 10),
		Skus:      make([]*proto_product.SaveSpuRequest_Skus, 0, 10),
	}
	for i := 0; i < len(req.BaseAttrs); i++ {
		rpcReq.BaseAttrs = append(rpcReq.BaseAttrs, &proto_product.SaveSpuRequest_Baseattrs{
			AttrId:     req.BaseAttrs[i].AttrID,
			AttrValues: req.BaseAttrs[i].AttrValues,
			ShowDesc:   req.BaseAttrs[i].ShowDesc,
		})
	}
	for i := 0; i < len(req.Skus); i++ {
		sku := proto_product.SaveSpuRequest_Skus{
			Attr:        make([]*proto_product.SaveSpuRequest_Attr, 0, 10),
			SkuName:     req.Skus[i].SkuName,
			Price:       req.Skus[i].Price,
			SkuTitle:    req.Skus[i].SkuTitle,
			SkuSubtitle: req.Skus[i].SkuSubtitle,
			Descar:      req.Skus[i].Descar,
			Images:      make([]*proto_product.SaveSpuRequest_Images, 0, 10),
			FullCount:   req.Skus[i].FullCount,
			Discount:    req.Skus[i].Discount,
			CountStatus: req.Skus[i].CountStatus,
			FullPrice:   req.Skus[i].FullPrice,
			ReducePrice: req.Skus[i].ReducePrice,
			PriceStatus: req.Skus[i].PriceStatus,
			MemberPrice: make([]*proto_product.SaveSpuRequest_Memberprice, 0, 10),
		}
		for j := 0; j < len(req.Skus[i].Attr); j++ {
			sku.Attr = append(sku.Attr, &proto_product.SaveSpuRequest_Attr{
				AttrId:    req.Skus[i].Attr[j].AttrID,
				AttrName:  req.Skus[i].Attr[j].AttrName,
				AttrValue: req.Skus[i].Attr[j].AttrValue,
			})
		}
		for j := 0; j < len(req.Skus[i].Images); j++ {
			sku.Images = append(sku.Images, &proto_product.SaveSpuRequest_Images{
				ImgUrl:     req.Skus[i].Images[j].ImgURL,
				DefaultTmg: req.Skus[i].Images[j].DefaultImg,
			})
		}
		for j := 0; j < len(req.Skus[i].MemberPrice); j++ {
			sku.MemberPrice = append(sku.MemberPrice, &proto_product.SaveSpuRequest_Memberprice{
				Id:    req.Skus[i].MemberPrice[j].ID,
				Name:  req.Skus[i].MemberPrice[j].Name,
				Price: req.Skus[i].MemberPrice[j].Price,
			})
		}
		rpcReq.Skus = append(rpcReq.Skus, &sku)
	}

	rpcClient := rpc_client.GetProductClient()
	_, err := rpcClient.SaveSpu(context.TODO(), &rpcReq)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

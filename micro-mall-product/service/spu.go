package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"mall-demo/micro-mall-product/global"
	"mall-demo/micro-mall-product/model"
	proto_coupon "mall-demo/micro-mall-product/proto/micro-mall-coupon-proto"
	proto_product "mall-demo/micro-mall-product/proto/micro-mall-product-proto"
	proto_search "mall-demo/micro-mall-product/proto/micro-mall-search-proto"
	proto_ware "mall-demo/micro-mall-product/proto/micro-mall-ware-proto"
	"mall-demo/micro-mall-product/rpc-client"
	"time"
)

type SpuService struct {
}

func (ss *SpuService) SaveSpu(ctx context.Context, req *proto_product.SaveSpuRequest) (*proto_product.SaveSpuResponse, error) {
	global.PmsMysqlConn.Transaction(func(tx *gorm.DB) error {
		// pms_spu_info
		var spuInfo model.PmsSpuInfo
		copier.Copy(&spuInfo, req)
		spuInfo.CreateTime = time.Now()
		spuInfo.UpdateTime = time.Now()
		if err := tx.Create(&spuInfo).Error; err != nil {
			return err
		}

		//pms_spu_decs
		spuDesc := make([]model.PmsSpuInfoDesc, 0, 10)
		for i := 0; i < len(req.Decript); i++ {
			spuDesc = append(spuDesc, model.PmsSpuInfoDesc{
				SpuId:   spuInfo.Id,
				Decript: req.Decript[i],
			})
		}
		if err := tx.Create(&spuDesc).Error; err != nil {
			return err
		}

		//spu_image
		var spuImages []model.PmsSpuImages
		spuImages = make([]model.PmsSpuImages, 0, 10)
		for i := 0; i < len(req.Images); i++ {
			spuImages = append(spuImages, model.PmsSpuImages{
				SpuId:  spuInfo.Id,
				ImgUrl: req.Images[i],
			})
		}
		if err := tx.Create(&spuImages).Error; err != nil {
			return err
		}

		//product_attr_value
		var productAttrValues []model.PmsProductAttrValue
		productAttrValues = make([]model.PmsProductAttrValue, 0, 10)
		for i := 0; i < len(req.BaseAttrs); i++ {
			productAttrValues = append(productAttrValues, model.PmsProductAttrValue{
				SpuId:     spuInfo.Id,
				AttrId:    int64(req.BaseAttrs[i].AttrId),
				AttrValue: req.BaseAttrs[i].AttrValues,
				QuickShow: int(req.BaseAttrs[i].ShowDesc),
			})
		}
		if err := tx.Create(&productAttrValues).Error; err != nil {
			return err
		}

		// skuInfos
		var skuInfos []model.PmsSkuInfo
		skuInfos = make([]model.PmsSkuInfo, 0, 10)
		for i := 0; i < len(req.Skus); i++ {
			skuInfos = append(skuInfos, model.PmsSkuInfo{
				SpuId:       spuInfo.Id,
				SkuName:     req.Skus[i].SkuName,
				CatalogId:   req.CatalogId,
				BrandId:     req.BrandId,
				SkuTitle:    req.Skus[i].SkuTitle,
				SkuSubtitle: req.Skus[i].SkuSubtitle,
				Price:       req.Skus[i].Price,
			})
		}
		if err := tx.Create(&skuInfos).Error; err != nil {
			return err
		}

		// sku_images
		var skuImages []model.PmsSkuImages
		skuImages = make([]model.PmsSkuImages, 0, 10)
		for i := 0; i < len(skuInfos); i++ {
			for j := 0; j < len(req.Skus[i].Images); j++ {
				skuImages = append(skuImages, model.PmsSkuImages{
					SkuId:      skuInfos[i].SkuId,
					ImgUrl:     req.Skus[i].Images[j].ImgUrl,
					DefaultImg: int(req.Skus[i].Images[j].DefaultTmg),
				})
			}
		}
		if err := tx.Create(&skuImages).Error; err != nil {
			return err
		}

		//sku_sale_attr_value
		var saleAttrValues []model.PmsSkuSaleAttrValue
		saleAttrValues = make([]model.PmsSkuSaleAttrValue, 0, 10)
		for i := 0; i < len(skuInfos); i++ {
			for j := 0; j < len(req.Skus[i].Attr); j++ {
				saleAttrValues = append(saleAttrValues, model.PmsSkuSaleAttrValue{
					SkuId:     skuInfos[i].SkuId,
					AttrId:    int64(req.Skus[i].Attr[j].AttrId),
					AttrName:  req.Skus[i].Attr[j].AttrName,
					AttrValue: req.Skus[i].Attr[j].AttrValue,
				})
			}
		}
		if err := tx.Create(&saleAttrValues).Error; err != nil {
			return err
		}

		// skuLadders
		rpcClient := rpc_client.GetCouponClient()
		var skuLadders []*proto_coupon.SkuLadderEntity
		skuLadders = make([]*proto_coupon.SkuLadderEntity, 0, 10)
		for i := 0; i < len(req.Skus); i++ {
			skuLadders = append(skuLadders, &proto_coupon.SkuLadderEntity{
				SkuId:     skuInfos[i].SkuId,
				FullCount: req.Skus[i].FullCount,
				Discount:  req.Skus[i].Discount,
				Price:     req.Skus[i].Price,
			})
		}
		if _, err := rpcClient.SaveSkuLadder(context.TODO(), &proto_coupon.SaveSkuLadderRequest{SkuLadderEntities: skuLadders}); err != nil {
			return err
		}

		// spuBounds
		if _, err := rpcClient.SaveSpuBounds(context.TODO(), &proto_coupon.SaveSpuBoundsRequest{SpuBoundEntities: &proto_coupon.SpuBoundEntity{
			SpuId:      spuInfo.Id,
			GrowBounds: float64(req.Bounds.GrowBounds),
			BuyBounds:  float64(req.Bounds.BuyBounds),
		}}); err != nil {
			return err
		}

		//sku_full_reduction
		skuFullReductions := make([]*proto_coupon.SkuFullReductionEntity, 0, 10)
		for i := 0; i < len(req.Skus); i++ {
			skuFullReductions = append(skuFullReductions, &proto_coupon.SkuFullReductionEntity{
				SkuId:       skuInfos[i].SkuId,
				FullPrice:   float64(req.Skus[i].FullPrice),
				ReducePrice: float64(req.Skus[i].ReducePrice),
			})
		}
		if _, err := rpcClient.SaveSkuFullReduction(context.TODO(), &proto_coupon.SaveSkuFullReductionRequest{
			SkuFullReductionEntities: skuFullReductions,
		}); err != nil {
			return err
		}

		//memberPrices
		memberPrices := make([]*proto_coupon.MemberPriceEntity, 0, 10)
		for i := 0; i < len(req.Skus); i++ {
			for j := 0; j < len(req.Skus[i].MemberPrice); j++ {
				memberPrices = append(memberPrices, &proto_coupon.MemberPriceEntity{
					SkuId:           skuInfos[i].SkuId,
					MemberPrice:     req.Skus[i].MemberPrice[j].Price,
					MemberLevelId:   int64(req.Skus[i].MemberPrice[j].Id),
					MemberLevelName: req.Skus[i].MemberPrice[j].Name,
				})
			}
		}
		fmt.Printf("%+v \n", memberPrices)

		if _, err := rpcClient.SaveMemberPrice(context.TODO(), &proto_coupon.SaveMemberPriceRequest{
			MemberPriceEntities: memberPrices,
		}); err != nil {
			return err
		}
		return nil
	})
	return &proto_product.SaveSpuResponse{}, nil
}

func (ss *SpuService) SearchSpuInfo(ctx context.Context, req *proto_product.SearchSpuInfoRequest) (*proto_product.SearchSpuInfoResponse, error) {
	var spuInfos []model.PmsSpuInfo
	global.PmsMysqlConn.Model(&model.PmsSpuInfo{}).Where("catalog_id = ? and brand_id = ? and publish_status = ?", req.CatalogId, req.BrandId, req.Status).Find(&spuInfos)
	fmt.Printf("%+v \n", spuInfos)

	var catelogName, brandName string
	global.PmsMysqlConn.Model(&model.PmsCategory{}).Select("name").Where("cat_id = ?", req.CatalogId).First(&catelogName)
	global.PmsMysqlConn.Model(&model.PmsBrand{}).Select("name").Where("brand_id = ?", req.BrandId).First(&brandName)
	resp := proto_product.SearchSpuInfoResponse{
		TotalCount:      int32(len(spuInfos)),
		SpuInfoEntities: make([]*proto_product.SpuInfoEntity, 0, 10),
	}
	for i := 0; i < len(spuInfos); i++ {
		var spuInfoEntity proto_product.SpuInfoEntity
		copier.Copy(&spuInfoEntity, &spuInfos[i])
		spuInfoEntity.BrandName = brandName
		spuInfoEntity.CatalogName = catelogName
		fmt.Printf("%+v \n", spuInfoEntity)
		resp.SpuInfoEntities = append(resp.SpuInfoEntities, &spuInfoEntity)
	}
	return &resp, nil
}

// 商品上架，刷新ES

func (ss *SpuService) UpSpu(ctx context.Context, req *proto_product.UpSpuRequest) (*proto_product.UpSpuResponse, error) {
	var spuInfo model.PmsSpuInfo
	global.PmsMysqlConn.Model(&model.PmsSpuInfo{}).Where("id = ?", req.SpuId).First(&spuInfo)
	spuInfo.PublishStatus = 1
	global.PmsMysqlConn.Save(spuInfo)

	// 1. 查出当前spuId对应的所有sku信息,品牌的名字
	var skusInfo []model.PmsSkuInfo
	global.PmsMysqlConn.Model(&model.PmsSkuInfo{}).Where("spu_id = ?", req.SpuId).Find(&skusInfo)
	// 2. 查出当前sku的所有可以被用来检索的规格属性
	var attrIds []int64
	global.PmsMysqlConn.Model(&model.PmsAttr{}).Select("attr_id").Where("search_type = 1").Find(&attrIds)
	var attrs []model.PmsProductAttrValue
	global.PmsMysqlConn.Model(&model.PmsProductAttrValue{}).Where("spu_id = ? and attr_id in ?", req.SpuId, attrIds).Find(&attrs)

	// 3.调用库存服务，查看库存
	var skuIds = make([]int64, 0, 10)
	for _, skuInfo := range skusInfo {
		skuIds = append(skuIds, skuInfo.SkuId)
	}
	var wareRpcReq proto_ware.GetSkuHasStockRequest
	copier.CopyWithOption(&wareRpcReq.SkuIds, &skuIds, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	wareClient := rpc_client.GetWareRpcClient()
	skuHasStock, err := wareClient.GetSkuHasStock(context.TODO(), &wareRpcReq)
	// 4. 封装sku 信息
	var esReq = make([]*proto_search.ProductStatusUpRequest_Entity, 0, 10)
	for _, skuInfo := range skusInfo {
		var esReqEntity proto_search.ProductStatusUpRequest_Entity
		copier.CopyWithOption(&esReqEntity, &skuInfo, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		esReqEntity.SkuPrice = skuInfo.Price
		esReqEntity.HasStock = skuHasStock.SkuHasStock[skuInfo.SkuId]
		var catalogName string
		var brand model.PmsBrand
		global.PmsMysqlConn.Model(&model.PmsCategory{}).Select("name").Where("cat_id = ?", skuInfo.CatalogId).First(&catalogName)
		global.PmsMysqlConn.Model(&model.PmsBrand{}).Where("brand_id = ?", skuInfo.BrandId).First(&brand)
		esReqEntity.BrandName = brand.Name
		esReqEntity.BrandImg = brand.Logo
		esReqEntity.CatalogName = catalogName
		copier.CopyWithOption(&esReqEntity.Attrs, &attrs, copier.Option{IgnoreEmpty: true, DeepCopy: true})

		// 设置热度评分
		esReqEntity.HotScore = 0
		esReq = append(esReq, &esReqEntity)
	}

	// 5. 调用es服务
	rpcClient := rpc_client.GetSearchClient()
	_, err = rpcClient.ProductStatusUp(context.TODO(), &proto_search.ProductStatusUpRequest{
		Entities: esReq,
	})
	return &proto_product.UpSpuResponse{}, err
}

func (ss *SpuService) SearchSkuInfo(ctx context.Context, req *proto_product.SearchSkuInfoRequest) (*proto_product.SearchSkuInfoResponse, error) {
	var skuInfos []model.PmsSkuInfo
	global.PmsMysqlConn.Model(&model.PmsSkuInfo{}).Offset(int((req.PageNum-1)*req.PageSize)).Limit(int(req.PageSize)).Where("catalog_id = ? and brand_id = ? and price >= ? and price <= ?", req.CatelogId, req.BrandId, req.Min, req.Max).Find(&skuInfos)
	resp := proto_product.SearchSkuInfoResponse{
		TotalCount: int32(len(skuInfos)),
		List:       make([]*proto_product.SearchSkuInfoResponse_List, 0, 10),
	}
	for i := 0; i < len(skuInfos); i++ {
		var skuInfo proto_product.SearchSkuInfoResponse_List
		copier.Copy(&skuInfo, &skuInfos[i])
		resp.List = append(resp.List, &skuInfo)
	}
	return &resp, nil
}

func (ss *SpuService) GetSpuInfo(ctx context.Context, req *proto_product.GetSpuInfoRequest) (*proto_product.GetSpuInfoResponse, error) {
	var spuAttrs []model.PmsProductAttrValue
	global.PmsMysqlConn.Model(&model.PmsProductAttrValue{}).Where("spu_id = ?", req.SpuId).Find(&spuAttrs)
	resp := proto_product.GetSpuInfoResponse{
		Data: make([]*proto_product.GetSpuInfoResponse_Data, 0, 10),
	}
	for i := 0; i < len(spuAttrs); i++ {
		var spuAttr proto_product.GetSpuInfoResponse_Data
		copier.Copy(&spuAttr, &spuAttrs[i])
		resp.Data = append(resp.Data, &spuAttr)
	}
	return &resp, nil
}

func (ss *SpuService) UpdateSpuAttrs(ctx context.Context, req *proto_product.UpdateSpuAttrsRequest) (*proto_product.UpdateSpuAttrsResponse, error) {
	for i := 0; i < len(req.Data); i++ {
		var spuAttr model.PmsProductAttrValue
		global.PmsMysqlConn.Model(&model.PmsProductAttrValue{}).Where("spu_id = ? and attr_id = ?", req.SpuId, req.Data[i].AttrId).First(&spuAttr)
		spuAttr.AttrName = req.Data[i].AttrName
		spuAttr.AttrValue = req.Data[i].AttrValue
		spuAttr.QuickShow = int(req.Data[i].QuickShow)
		global.PmsMysqlConn.Save(&spuAttr)
	}
	return &proto_product.UpdateSpuAttrsResponse{}, nil
}

func (ss *SpuService) GetSkuItem(ctx context.Context, req *proto_product.GetSkuItemRequest) (*proto_product.GetSkuItemResponse, error) {
	skuInfo := model.PmsSkuInfo{}
	global.PmsMysqlConn.Model(&model.PmsSkuInfo{}).Where("sku_id = ?", req.SkuId).First(&skuInfo)
	//images := make([]proto_product.GetSkuItemResponse_SkuImagesEntity, 0, 10)
	//global.PmsMysqlConn.Model(&model.PmsSkuImages{}).Where("sku_id = ?", req.SkuId).Find(&images)

	var skuInfoEntity proto_product.GetSkuItemResponse_SkuInfoEntity
	//
	copier.CopyWithOption(&skuInfoEntity, &skuInfo, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	resp := &proto_product.GetSkuItemResponse{
		SkuInfoEntity: &skuInfoEntity,
	}
	fmt.Println(resp)
	return resp, nil
}

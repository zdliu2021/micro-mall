package service

import (
	"context"
	"errors"
	"fmt"
	"mall-demo/micro-mall-product/global"
	"mall-demo/micro-mall-product/model"
	proto_product "mall-demo/micro-mall-product/proto/micro-mall-product-proto"
	"mall-demo/micro-mall-product/utils"
	"strconv"
)

type BrandService struct {
}

func (bds *BrandService) GetBrandList(ctx context.Context, req *proto_product.GetBrandListRequest) (*proto_product.GetBrandListResponse, error) {
	if req.Keyword != "" {
		brand_id, _ := strconv.Atoi(req.Keyword)
		var res []model.PmsBrand
		global.PmsMysqlConn.Model(&model.PmsBrand{}).Offset(int(req.PageSize*(req.PageNum-1))).Limit(int(req.PageSize)).Where("brand_id = ? or name like ?", brand_id, req.Keyword).Find(&res)
		var resp proto_product.GetBrandListResponse
		resp.BrandEntities = make([]*proto_product.BrandEntity, 0, 10)
		for i := 0; i < len(res); i++ {
			resp.BrandEntities = append(resp.BrandEntities, &proto_product.BrandEntity{
				BrandId:     res[i].BrandId,
				Name:        res[i].Name,
				Logo:        res[i].Logo,
				Descript:    res[i].Logo,
				ShowStatus:  int32(res[i].ShowStatus),
				FirstLetter: res[i].FirstLetter,
				Sort:        int32(res[i].Sort),
			})
		}
		return &resp, nil
	} else {
		var res []model.PmsBrand
		global.PmsMysqlConn.Model(&model.PmsBrand{}).Offset(int(req.PageSize * (req.PageNum - 1))).Limit(int(req.PageSize)).Find(&res)
		var resp proto_product.GetBrandListResponse
		resp.BrandEntities = make([]*proto_product.BrandEntity, 0, 10)
		for i := 0; i < len(res); i++ {
			resp.BrandEntities = append(resp.BrandEntities, &proto_product.BrandEntity{
				BrandId:     res[i].BrandId,
				Name:        res[i].Name,
				Logo:        res[i].Logo,
				Descript:    res[i].Logo,
				ShowStatus:  int32(res[i].ShowStatus),
				FirstLetter: res[i].FirstLetter,
				Sort:        int32(res[i].Sort),
			})
		}
		return &resp, nil
	}
}

func (bds *BrandService) GetBrandInfo(ctx context.Context, req *proto_product.GetBrandInfoRequest) (*proto_product.GetBrandInfoResponse, error) {
	var brand model.PmsBrand
	if err := global.PmsMysqlConn.Model(&model.PmsBrand{}).Where("brand_id = ?", req.BrandId).First(&brand).Error; err != nil {
		return nil, errors.New("brand id not exists")
	}
	resp := proto_product.GetBrandInfoResponse{
		BrandEntity: &proto_product.BrandEntity{
			BrandId:     brand.BrandId,
			Name:        brand.Name,
			Logo:        brand.Logo,
			Descript:    brand.Logo,
			ShowStatus:  int32(brand.ShowStatus),
			FirstLetter: brand.FirstLetter,
			Sort:        int32(brand.Sort),
		},
	}
	return &resp, nil
}

func (bds *BrandService) SaveBrand(ctx context.Context, req *proto_product.SaveBrandRequest) (*proto_product.SaveBrandResponse, error) {
	newBrand := &model.PmsBrand{}
	if ok := utils.Assignment(newBrand, req.BrandEntity); !ok {
		global.GVA_LOG.Error("Assigment failed.")
		return nil, errors.New("assigment failed")
	}

	if err := global.PmsMysqlConn.Model(&model.PmsBrand{}).Create(newBrand).Error; err != nil {
		global.GVA_LOG.Error("save brand failed" + err.Error())
		return nil, errors.New("save brand failed" + err.Error())
	}
	return &proto_product.SaveBrandResponse{}, nil
}

func (bds *BrandService) DeleteBrand(ctx context.Context, req *proto_product.DeleteBrandRequest) (*proto_product.DeleteBrandResponse, error) {

	if len(req.BrandIds) == 0 {
		return nil, errors.New("the list cannot be empty")
	}
	if err := global.PmsMysqlConn.Delete(&model.PmsBrand{}, req.BrandIds).Error; err != nil {
		return nil, err
	}
	return &proto_product.DeleteBrandResponse{}, nil

}

func (bds *BrandService) UpdateBrand(ctx context.Context, req *proto_product.UpdateBrandRequest) (*proto_product.UpdateBrandResponse, error) {
	var brand model.PmsBrand
	global.PmsMysqlConn.Model(&model.PmsBrand{}).Where("brand_id = ?", req.BrandEntity.BrandId).First(&brand)
	brand.Name = req.BrandEntity.Name
	brand.Sort = int(req.BrandEntity.Sort)
	brand.Logo = req.BrandEntity.Logo
	brand.ShowStatus = int(req.BrandEntity.ShowStatus)
	brand.FirstLetter = req.BrandEntity.FirstLetter
	brand.Descript = req.BrandEntity.Descript
	if err := global.PmsMysqlConn.Save(&brand).Error; err != nil {
		return nil, err
	}
	return &proto_product.UpdateBrandResponse{}, nil
}

func (bds *BrandService) UpdateBrandStatus(ctx context.Context, req *proto_product.UpdateBrandStatusRequest) (*proto_product.UpdateBrandStatusResponse, error) {
	if err := global.PmsMysqlConn.Model(&model.PmsBrand{}).Where("brand_id = ?", req.BrandId).Update("show_status = ?", req.ShowStatus).Error; err != nil {
		return nil, err
	}
	return &proto_product.UpdateBrandStatusResponse{}, nil
}

func (bds *BrandService) GetBrandRelatedCateLog(ctx context.Context, req *proto_product.GetBrandRelatedCateLogRequest) (*proto_product.GetBrandRelatedCateLogResponse, error) {
	var res []model.PmsCategoryBrandRelation
	global.PmsMysqlConn.Model(&model.PmsCategoryBrandRelation{}).Where("brand_id = ?", req.BrandId).Find(&res)
	resp := &proto_product.GetBrandRelatedCateLogResponse{
		BrandRelatedCateLogEntities: make([]*proto_product.BrandRelatedCateLogEntity, 0, 10),
	}
	for i := 0; i < len(res); i++ {
		resp.BrandRelatedCateLogEntities = append(resp.BrandRelatedCateLogEntities, &proto_product.BrandRelatedCateLogEntity{
			CatelogId:  res[i].CatelogId,
			CatlogName: res[i].CatelogName,
			Id:         res[i].Id,
			BrandId:    res[i].BrandId,
			BrandName:  res[i].BrandName,
		})
	}
	return resp, nil
}

func (bds *BrandService) SaveBrandCatelogRelation(ctx context.Context, req *proto_product.SaveBrandCatelogRelationRequest) (*proto_product.SaveBrandCatelogRelationResponse, error) {
	var brandName, catelogName string
	global.PmsMysqlConn.Model(&model.PmsCategory{}).Select("name").Where("cat_id = ?", req.CatelogId).First(&catelogName)
	global.PmsMysqlConn.Model(&model.PmsBrand{}).Select("name").Where("brand_id = ?", req.BrandId).First(&brandName)

	relation := model.PmsCategoryBrandRelation{
		BrandId:     req.BrandId,
		CatelogId:   req.CatelogId,
		BrandName:   brandName,
		CatelogName: catelogName,
	}
	global.PmsMysqlConn.Save(&relation)
	return &proto_product.SaveBrandCatelogRelationResponse{}, nil
}

func (bds *BrandService) DeleteBrandCatelogRelation(ctx context.Context, req *proto_product.DeleteBrandCatelogRelationRequest) (*proto_product.DeleteBrandCatelogRelationResponse, error) {
	if len(req.Ids) != 0 {
		fmt.Println(req.Ids)
		if err := global.PmsMysqlConn.Delete(&model.PmsCategoryBrandRelation{}, req.Ids).Error; err != nil {
			fmt.Println(err)
		}
	}
	return &proto_product.DeleteBrandCatelogRelationResponse{}, nil
}

func (bds *BrandService) GetBrandRelatedWithCatelog(ctx context.Context, req *proto_product.GetBrandRelatedWithCatelogRequest) (*proto_product.GetBrandRelatedWithCatelogResponse, error) {
	var relations []model.PmsCategoryBrandRelation
	global.PmsMysqlConn.Model(&model.PmsCategoryBrandRelation{}).Where("catelog_id = ?", req.CatId).Find(&relations)

	resp := proto_product.GetBrandRelatedWithCatelogResponse{
		BrandCatelogRelations: make([]*proto_product.BrandCatelogRelation, 0, 10),
	}

	for i := 0; i < len(relations); i++ {
		resp.BrandCatelogRelations = append(resp.BrandCatelogRelations, &proto_product.BrandCatelogRelation{
			Id:          relations[i].Id,
			BrandId:     relations[i].BrandId,
			BrandName:   relations[i].BrandName,
			CatelogId:   relations[i].CatelogId,
			CatelogName: relations[i].CatelogName,
		})
	}
	return &resp, nil
}

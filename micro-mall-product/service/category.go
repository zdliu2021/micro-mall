package service

import (
	"context"
	"encoding/json"
	"errors"
	"mall-demo/micro-mall-product/global"
	"mall-demo/micro-mall-product/model"
	proto_product "mall-demo/micro-mall-product/proto/micro-mall-product-proto"
	"mall-demo/micro-mall-product/utils"
	"mall-demo/micro-mall-product/utils/db"
)

type CategoryService struct {
}

func (cgy *CategoryService) ListCategoryTree(ctx context.Context, req *proto_product.ListCategoryTreeRequest) (*proto_product.ListCategoryTreeResponse, error) {
	redisClient := db.GetRedisInstance(db.DefaultRedisOption)
	val, err := redisClient.Get("category_tree").Result()
	if err == nil {
		global.GVA_LOG.Info("redis get.")
		var res proto_product.ListCategoryTreeResponse
		_ = json.Unmarshal([]byte(val), &res)
		return &res, nil
	}

	var solve func(entity *proto_product.CategoryEntity)
	solve = func(entity *proto_product.CategoryEntity) {
		var children []model.PmsCategory
		global.PmsMysqlConn.Where("parent_cid = ?", entity.CatId).Find(&children)
		for _, child := range children {
			entity.Children = append(entity.Children, &proto_product.CategoryEntity{
				CatId:        child.CatId,
				Name:         child.Name,
				ParentCid:    child.ParentCid,
				CatLevel:     int32(child.CatLevel),
				ShowStatus:   int32(child.ShowStatus),
				Sort:         int32(child.Sort),
				Icon:         child.Icon,
				ProductUnit:  child.ProductUnit,
				ProductCount: int32(child.ProductCount),
			})
		}
		for _, child := range entity.Children {
			solve(child)
		}
	}

	var pmsCategories []model.PmsCategory
	global.PmsMysqlConn.Model(&model.PmsCategory{}).Where("cat_level = ?", 1).Find(&pmsCategories)

	resp := make([]*proto_product.CategoryEntity, 0, len(pmsCategories))
	for _, pmsCategorry := range pmsCategories {
		entity := &proto_product.CategoryEntity{
			CatId:        pmsCategorry.CatId,
			Name:         pmsCategorry.Name,
			ParentCid:    pmsCategorry.ParentCid,
			CatLevel:     int32(pmsCategorry.CatLevel),
			ShowStatus:   int32(pmsCategorry.ShowStatus),
			Sort:         int32(pmsCategorry.Sort),
			Icon:         pmsCategorry.Icon,
			ProductUnit:  pmsCategorry.ProductUnit,
			ProductCount: int32(pmsCategorry.ProductCount),
			Children:     make([]*proto_product.CategoryEntity, 0, 100),
		}
		solve(entity)
		resp = append(resp, entity)
	}

	tmp, _ := json.Marshal(proto_product.ListCategoryTreeResponse{
		CategoryEntities: resp,
	})
	redisClient.Set("category_tree", tmp, 0)

	return &proto_product.ListCategoryTreeResponse{
		CategoryEntities: resp,
	}, nil
}

func (cgy *CategoryService) DeleteCategory(ctx context.Context, req *proto_product.DeleteCategoryRequest) (*proto_product.DeleteCategoryResponse, error) {
	// TODO 检查要删除的cat是否被引用
	if len(req.CatIds) == 0 {
		return nil, errors.New("the list cannot be empty")
	}
	if err := global.PmsMysqlConn.Delete(&model.PmsCategory{}, req.CatIds).Error; err != nil {
		return nil, err
	}
	return &proto_product.DeleteCategoryResponse{}, nil
}

func (cgy *CategoryService) SaveCategory(ctx context.Context, req *proto_product.SaveCategoryRequest) (*proto_product.SaveCategoryResponse, error) {
	newCategory := &model.PmsCategory{}
	if ok := utils.Assignment(newCategory, req); !ok {
		global.GVA_LOG.Error("Assigment failed.")
		return nil, errors.New("assigment failed")
	}

	if err := global.PmsMysqlConn.Model(&model.PmsCategory{}).Create(newCategory).Error; err != nil {
		global.GVA_LOG.Error("save category failed" + err.Error())
		return nil, errors.New("save category failed" + err.Error())
	}
	return &proto_product.SaveCategoryResponse{}, nil
}

func (cgy *CategoryService) UpdateCategory(ctx context.Context, req *proto_product.UpdateCategoryRequest) (*proto_product.UpdateCategoryResponse, error) {
	var category model.PmsCategory
	global.PmsMysqlConn.Where("cat_id = ?", req.CatId).First(&category)

	if ok := utils.NotZeroAssignment(&category, req); !ok {
		global.GVA_LOG.Error("Assigment failed.")
		return nil, errors.New("assigment failed")
	}
	global.PmsMysqlConn.Save(&category)

	return &proto_product.UpdateCategoryResponse{}, nil
}

func (cgy *CategoryService) GetCategoryInfo(ctx context.Context, req *proto_product.GetCategoryInfoRequest) (*proto_product.GetCategoryInfoResponse, error) {
	var category model.PmsCategory
	global.PmsMysqlConn.Where("cat_id = ?", req.CatId).First(&category)

	var resp proto_product.GetCategoryInfoResponse
	utils.Assignment(&resp, &category)
	return &resp, nil
}
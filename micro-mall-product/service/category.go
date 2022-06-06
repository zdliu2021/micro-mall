package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"mall-demo/micro-mall-product/global"
	"mall-demo/micro-mall-product/model"
	proto_product "mall-demo/micro-mall-product/proto/micro-mall-product-proto"
	"mall-demo/micro-mall-product/utils"
	"mall-demo/micro-mall-product/utils/idgenerator"
	"time"
)

type CategoryService struct {
}

func (cgy *CategoryService) ListCategoryTree(ctx context.Context, req *proto_product.ListCategoryTreeRequest) (*proto_product.ListCategoryTreeResponse, error) {
	val, err := global.GVA_REDIS.Get("category_tree").Result()
	if err == nil {
		global.GVA_LOG.Info("redis get.")
		var res proto_product.ListCategoryTreeResponse
		_ = json.Unmarshal([]byte(val), &res)
		return &res, nil
	} else {
		if err == redis.Nil {
			global.GVA_LOG.Warn("缓存不存在该数据")
		} else {
			global.GVA_LOG.Error("其他错误:" + err.Error())
		}
	}
	global.GVA_LOG.Info("redis 内不存在")
	// 防止缓存击穿
	uuid := idgenerator.GetSnowflakeId()
	global.GVA_LOG.Info("uuid:", zap.String("uuid", uuid))
	for {
		if ok, _ := global.GVA_REDIS.SetNX("category_lock", uuid, 30*time.Second).Result(); !ok {
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	global.GVA_LOG.Info("获取分布式锁成功")

	defer func() {
		//// 删除锁
		//// 删除分布式锁
		global.GVA_LOG.Info("删除分布式锁")
		script := "if redis.call('GET',KEYS[1])==ARGV[1] then return redis.call('DEL',KEYS[1]) else return 0 end"
		err = global.GVA_REDIS.Eval(script, []string{"category_lock"}, uuid).Err()
		if err != nil {
			global.GVA_LOG.Error("脚本错误:" + err.Error())
		}
	}()

	val, err = global.GVA_REDIS.Get("category_tree").Result()
	if err == nil {
		global.GVA_LOG.Info("redis get.")
		var res proto_product.ListCategoryTreeResponse
		_ = json.Unmarshal([]byte(val), &res)
		return &res, nil
	}
	global.GVA_LOG.Info("访问数据库获取数据")
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
	global.GVA_REDIS.Set("category_tree", tmp, 0)

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

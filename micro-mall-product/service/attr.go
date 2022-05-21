package service

import (
	"context"
	"fmt"
	"mall-demo/micro-mall-product/global"
	model "mall-demo/micro-mall-product/model"
	proto_product "mall-demo/micro-mall-product/proto/micro-mall-product-proto"
	"math"
)

type AttrService struct {
}

func (as *AttrService) GetAttrListRelatedWithCatelog(ctx context.Context, req *proto_product.GetAttrListRelatedWithCatelogRequest) (*proto_product.GetAttrListRelatedWithCatelogResponse, error) {
	var res proto_product.GetAttrListRelatedWithCatelogResponse
	res.AttrEntities = make([]*proto_product.AttrEntity, 0, 10)
	if len(req.Key) == 0 {
		var attrs []model.PmsAttr
		global.PmsMysqlConn.Model(&model.PmsAttr{}).Offset(int((req.Page-1)*req.Limit)).Limit(int(req.Limit)).Where("catelog_id = ?", req.CatelogId).Find(&attrs)
		//"catelogName": "手机/数码/手机", //所属分类名字
		//			"groupName": "主体", //所属分组名字
		for i := 0; i < len(attrs); i++ {
			var catelogName, groupName string
			global.PmsMysqlConn.Model(&model.PmsCategory{}).Select("name").Where("cat_id", attrs[i].CatelogId).First(&catelogName)
			global.PmsMysqlConn.Raw("select attr_group_name from pms_attr,pms_attr_attrgroup"+
				"_relation,pms_attr_group where pms_attr.attr_id = pms_attr_attrgroup_relation.attr_id and "+
				"pms_attr_attrgroup_relation.attr_group_id = pms_attr_group.attr_group_id and pms_attr.attr_id = ? ", attrs[i].AttrId).Scan(&groupName)
			res.AttrEntities = append(res.AttrEntities, &proto_product.AttrEntity{
				AttrId:      attrs[i].AttrId,
				AttrName:    attrs[i].AttrName,
				SearchType:  int32(attrs[i].SearchType),
				ValueSelect: attrs[i].ValueSelect,
				ValueType:   int32(attrs[i].ValueType),
				Icon:        attrs[i].Icon,
				AttrType:    int32(attrs[i].AttrType),
				Enable:      int32(attrs[i].Enable),
				CatelogName: catelogName,
				GroupName:   groupName,
			})
		}
		var count int64
		global.PmsMysqlConn.Model(&model.PmsAttr{}).Where("catelog_id = ?", req.CatelogId).Count(&count)
		res.CurrPage = req.Page
		res.TotalPage = int32(math.Ceil(float64(1.0 * count / int64(req.Limit))))
		res.PageSize = req.Limit
		res.TotalCount = int32(len(res.AttrEntities))
	}

	return &res, nil
}

func (as *AttrService) SaveProductAttr(ctx context.Context, req *proto_product.SaveProductAttrRequest) (*proto_product.SaveProductAttrResponse, error) {
	attrModel := &model.PmsAttr{
		AttrName:    req.AttrName,
		SearchType:  int(req.SearchType),
		ValueType:   int(req.ValueType),
		Icon:        req.Icon,
		ValueSelect: req.ValueSelect,
		AttrType:    int(req.AttrType),
		Enable:      int64(req.Enable),
		CatelogId:   req.CatelogId,
		ShowDesc:    int(req.ShowDesc),
	}
	global.PmsMysqlConn.Save(&attrModel)

	return &proto_product.SaveProductAttrResponse{}, nil
}

func (as *AttrService) GetAttrInfo(ctx context.Context, req *proto_product.GetAttrInfoRequest) (*proto_product.GetAttrInfoResponse, error) {
	var attrModel model.PmsAttr
	global.PmsMysqlConn.Model(&model.PmsAttr{}).Where("attr_id = ?", req.AttrId).First(&attrModel)

	var attrGroupID int64
	global.PmsMysqlConn.Model(&model.PmsAttrAttrgroupRelation{}).Select("attr_group_id").Where("attr_id = ?", req.AttrId).First(&attrGroupID)

	pars := []int64{attrModel.CatelogId}
	findParent(&pars, attrModel.CatelogId)
	return &proto_product.GetAttrInfoResponse{
		AttrId:      attrModel.AttrId,
		AttrName:    attrModel.AttrName,
		SearchType:  int32(attrModel.SearchType),
		ValueType:   int32(attrModel.ValueType),
		Icon:        attrModel.Icon,
		ValueSelect: attrModel.ValueSelect,
		AttrType:    int32(attrModel.AttrType),
		Enable:      int32(attrModel.Enable),
		CatelogId:   attrModel.CatelogId,
		ShowDesc:    int32(attrModel.ShowDesc),
		AttrGroupId: int32(attrGroupID),
		CatelogPath: pars,
	}, nil
}

func (as *AttrService) GetAttrNotCorrelation(ctx context.Context, req *proto_product.GetAttrNotCorrelationRequest) (*proto_product.GetAttrNotCorrelationResponse, error) {
	var attrModel []model.PmsAttr
	global.PmsMysqlConn.Raw("select * from pms_attr as A where not exists (select * from pms_attr as B join pms_attr_attrgroup_relation as C using(attr_id) where B.attr_id = A.attr_id and C.attr_group_id = ?)", req.AttrGroupId).Scan(&attrModel)

	var count int64
	global.PmsMysqlConn.Table("pms_attr").Joins("join pms_attr_attrgroup_relation on pms_attr.attr_id = pms_attr_attrgroup_relation.attr_id").Where(""+
		"pms_attr_attrgroup_relation.attr_group_id  <> ?", req.AttrGroupId).Count(&count)

	resp := proto_product.GetAttrNotCorrelationResponse{
		PmsAttrEntities: make([]*proto_product.PmsAttrEntity, 0, 10),
		TotalCount:      int32(len(attrModel)),
		TotalPage:       int32(math.Ceil(float64(1.0 * count / int64(req.Limit)))),
		CurrPage:        req.Page,
		PageSize:        req.Limit,
	}

	for i := 0; i < len(attrModel); i++ {
		resp.PmsAttrEntities = append(resp.PmsAttrEntities, &proto_product.PmsAttrEntity{
			AttrId:      attrModel[i].AttrId,
			AttrName:    attrModel[i].AttrName,
			SearchType:  int32(attrModel[i].SearchType),
			ValueType:   int32(attrModel[i].ValueType),
			Icon:        attrModel[i].Icon,
			ValueSelect: attrModel[i].ValueSelect,
			AttrType:    int32(attrModel[i].AttrType),
			Enable:      int32(attrModel[i].Enable),
			CatelogId:   attrModel[i].CatelogId,
			ShowDesc:    int32(attrModel[i].ShowDesc),
		})
	}
	return &resp, nil
}

func (as *AttrService) DeleteAttrAttrGroupRelation(ctx context.Context, req *proto_product.DeleteAttrAttrGroupRelationRequest) (*proto_product.DeleteAttrAttrGroupRelationResponse, error) {
	for i := 0; i < len(req.AttrAttrGroupRelationEntities); i++ {
		global.PmsMysqlConn.Where("attr_id  = ? and attr_group_id = ?", req.AttrAttrGroupRelationEntities[i].AttrId, req.AttrAttrGroupRelationEntities[i].AttrGroupId).Delete(&model.PmsAttrAttrgroupRelation{})
	}
	return &proto_product.DeleteAttrAttrGroupRelationResponse{}, nil
}

func (as *AttrService) SaveAttrAttrGroupRelation(ctx context.Context, req *proto_product.SaveAttrAttrGroupRelationRequest) (*proto_product.SaveAttrAttrGroupRelationResponse, error) {
	for i := 0; i < len(req.AttrAttrGroupRelationEntities); i++ {
		newRelation := model.PmsAttrAttrgroupRelation{
			AttrId:      req.AttrAttrGroupRelationEntities[i].AttrId,
			AttrGroupId: req.AttrAttrGroupRelationEntities[i].AttrGroupId,
		}
		global.PmsMysqlConn.Save(&newRelation)
	}

	return &proto_product.SaveAttrAttrGroupRelationResponse{}, nil
}

func (as *AttrService) GetAttrRelatedAttrGroup(ctx context.Context, req *proto_product.GetAttrRelatedAttrGroupRequest) (*proto_product.GetAttrRelatedAttrGroupResponse, error) {
	fmt.Printf("%+v \n", req)
	var attrModel []model.PmsAttr
	global.PmsMysqlConn.Table("pms_attr").Select("pms_attr.*").Joins("join pms_attr_attrgroup_relation").Where("pms_attr_attrgroup_relation.attr_id = pms_attr.attr_id and pms_attr_attrgroup_relation.attr_group_id = ?", req.AttrGroupId).Find(&attrModel)
	resp := proto_product.GetAttrRelatedAttrGroupResponse{
		PmsAttrEntities: make([]*proto_product.PmsAttrEntity, 0, 10),
	}
	for i := 0; i < len(attrModel); i++ {
		resp.PmsAttrEntities = append(resp.PmsAttrEntities, &proto_product.PmsAttrEntity{
			AttrId:      attrModel[i].AttrId,
			AttrName:    attrModel[i].AttrName,
			SearchType:  int32(attrModel[i].SearchType),
			ValueType:   int32(attrModel[i].ValueType),
			Icon:        attrModel[i].Icon,
			ValueSelect: attrModel[i].ValueSelect,
			AttrType:    int32(attrModel[i].AttrType),
			Enable:      int32(attrModel[i].Enable),
			CatelogId:   attrModel[i].CatelogId,
			ShowDesc:    int32(attrModel[i].ShowDesc),
		})
	}
	return &resp, nil
}
func (as *AttrService) UpdateAttr(ctx context.Context, req *proto_product.UpdateAttrRequest) (*proto_product.UpdateAttrResponse, error) {
	attrModel := model.PmsAttr{
		AttrId:      req.AttrId,
		AttrName:    req.AttrName,
		SearchType:  int(req.SearchType),
		ValueType:   int(req.ValueType),
		Icon:        req.Icon,
		ValueSelect: req.ValueSelect,
		AttrType:    int(req.AttrType),
		Enable:      int64(req.Enable),
		CatelogId:   req.CatelogId,
		ShowDesc:    int(req.ShowDesc),
	}
	global.PmsMysqlConn.Save(&attrModel)
	return &proto_product.UpdateAttrResponse{}, nil
}

func (as *AttrService) DeleteAttr(ctx context.Context, req *proto_product.DeleteAttrRequest) (*proto_product.DeleteAttrResponse, error) {
	global.PmsMysqlConn.Delete(&model.PmsAttr{}, req.AttrIds)
	return &proto_product.DeleteAttrResponse{}, nil
}

package service

import (
	"context"
	"fmt"
	"mall-demo/micro-mall-product/global"
	model "mall-demo/micro-mall-product/model"
	proto_product "mall-demo/micro-mall-product/proto/micro-mall-product-proto"
	"math"
)

type AttrGroupService struct {
}

func findParent(pars *[]int64, cur int64) {
	var par int64
	global.PmsMysqlConn.Model(&model.PmsCategory{}).Select("parent_cid").Where("cat_id = ?", cur).First(&par)
	if par == 0 {
		return
	}
	*pars = append(*pars, par)
	copy((*pars)[1:], (*pars)[:len(*pars)-1])
	(*pars)[0] = par
	findParent(pars, par)
}

func (ags *AttrGroupService) ListCatelogAttrGroup(ctx context.Context, req *proto_product.ListCatelogAttrGroupRequest) (*proto_product.ListCatelogAttrGroupResponse, error) {
	var res []model.PmsAttrGroup
	global.PmsMysqlConn.Model(&model.PmsAttrGroup{}).Offset(int((req.PageNum-1)*req.PageSize)).Limit(int(req.PageSize)).Where("catelog_id = ?", req.CatelogId).Find(&res)

	var count int64
	global.PmsMysqlConn.Model(&model.PmsAttrGroup{}).Where("catelog_id = ?", req.CatelogId).Count(&count)

	resp := make([]*proto_product.CatelogAttrGroupEntity, 0, 10)
	for i := 0; i < len(res); i++ {
		pars := []int64{res[i].CatelogId}
		findParent(&pars, res[i].CatelogId)
		resp = append(resp, &proto_product.CatelogAttrGroupEntity{
			AttrGroupId:   res[i].AttrGroupId,
			AttrGroupName: res[i].AttrGroupName,
			CatelogId:     res[i].CatelogId,
			Descript:      res[i].Descript,
			Icon:          res[i].Icon,
			Sort:          int32(res[i].Sort),
			CatelogPath:   pars,
		})
	}
	fmt.Printf("%+v \n", resp)
	return &proto_product.ListCatelogAttrGroupResponse{
		TotalCount:               int32(len(resp)),
		PageSize:                 req.PageSize,
		TotalPage:                int32(math.Ceil(float64(count * 1.0 / int64(req.PageSize)))),
		CurrPage:                 req.PageNum,
		CatelogAttrGroupEntities: resp,
	}, nil
}

func (ags *AttrGroupService) GetAttrGroupInfo(ctx context.Context, req *proto_product.GetAttrGroupInfoRequest) (*proto_product.GetAttrGroupInfoResponse, error) {
	var res model.PmsAttrGroup
	global.PmsMysqlConn.Model(&model.PmsAttrGroup{}).Where("attr_group_id = ?", req.AttrGroupId).First(&res)

	pars := []int64{res.CatelogId}
	findParent(&pars, res.CatelogId)

	return &proto_product.GetAttrGroupInfoResponse{
		CatelogAttrGroupEntity: &proto_product.CatelogAttrGroupEntity{
			AttrGroupId:   res.AttrGroupId,
			AttrGroupName: res.AttrGroupName,
			Sort:          int32(res.Sort),
			Descript:      res.Descript,
			Icon:          res.Icon,
			CatelogId:     res.CatelogId,
			CatelogPath:   pars,
		},
	}, nil
}

func (ags *AttrGroupService) UpdateAttrGroup(ctx context.Context, req *proto_product.UpdateAttrGroupRequest) (*proto_product.UpdateAttrGroupResponse, error) {
	attrGroup := model.PmsAttrGroup{
		AttrGroupId:   req.AttrGroupId,
		AttrGroupName: req.AttrGroupName,
		Sort:          int(req.Sort),
		Descript:      req.Descript,
		Icon:          req.Icon,
		CatelogId:     req.CatelogId,
	}
	global.PmsMysqlConn.Save(&attrGroup)
	return &proto_product.UpdateAttrGroupResponse{}, nil
}

func (ags *AttrGroupService) SaveAttrGroup(ctx context.Context, req *proto_product.SaveAttrGroupRequest) (*proto_product.SaveAttrGroupResponse, error) {
	attrGroup := model.PmsAttrGroup{
		AttrGroupId:   req.AttrGroupId,
		AttrGroupName: req.AttrGroupName,
		Sort:          int(req.Sort),
		Descript:      req.Descript,
		Icon:          req.Icon,
		CatelogId:     req.CatelogId,
	}
	global.PmsMysqlConn.Save(&attrGroup)
	return &proto_product.SaveAttrGroupResponse{}, nil
}
func (ags *AttrGroupService) DeleteAttrGroup(ctx context.Context, req *proto_product.DeleteAttrGroupRequest) (*proto_product.DeleteAttrGroupResponse, error) {
	global.PmsMysqlConn.Where("attr_group_id in ?", req.AttrGroupId).Delete(&model.PmsAttrGroup{})
	return &proto_product.DeleteAttrGroupResponse{}, nil
}

func (ags *AttrGroupService) GetAllGroupAndAttrRelatedWithCatelog(ctx context.Context, req *proto_product.GetAllGroupAndAttrRelatedWithCatelogRequest) (*proto_product.GetAllGroupAndAttrRelatedWithCatelogResponse, error) {
	var attrgroups []model.PmsAttrGroup
	global.PmsMysqlConn.Model(&model.PmsAttrGroup{}).Where("catelog_id = ?", req.CatId).Find(&attrgroups)

	resp := proto_product.GetAllGroupAndAttrRelatedWithCatelogResponse{
		AttrGroupAndAtrrsEntities: make([]*proto_product.AttrGroupAndAtrrsEntity, 0, 10),
	}

	for i := 0; i < len(attrgroups); i++ {
		var attrs []model.PmsAttr
		global.PmsMysqlConn.Raw("select pms_attr.* from pms_attr join pms_attr_attrgroup_relation using(attr_id) where attr_group_id = ?", attrgroups[i].AttrGroupId).Scan(&attrs)

		attrsEntity := make([]*proto_product.PmsAttrEntity, 0, 10)
		for j := 0; j < len(attrs); j++ {
			attrsEntity = append(attrsEntity, &proto_product.PmsAttrEntity{
				AttrId:      attrs[j].AttrId,
				AttrName:    attrs[j].AttrName,
				SearchType:  int32(attrs[j].SearchType),
				ValueType:   int32(attrs[j].ValueType),
				Icon:        attrs[j].Icon,
				ValueSelect: attrs[j].ValueSelect,
				AttrType:    int32(attrs[j].AttrType),
				Enable:      int32(attrs[j].Enable),
				CatelogId:   attrs[j].CatelogId,
				ShowDesc:    int32(attrs[j].ShowDesc),
			})
		}
		resp.AttrGroupAndAtrrsEntities = append(resp.AttrGroupAndAtrrsEntities, &proto_product.AttrGroupAndAtrrsEntity{
			AttrGroupId:     attrgroups[i].AttrGroupId,
			AttrGroupName:   attrgroups[i].AttrGroupName,
			Sort:            int32(attrgroups[i].Sort),
			Descript:        attrgroups[i].Descript,
			Icon:            attrgroups[i].Icon,
			CatelogId:       attrgroups[i].CatelogId,
			PmsAttrEntities: attrsEntity,
		})
	}
	return &resp, nil
}

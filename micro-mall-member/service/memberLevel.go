package service

import (
	"context"
	"mall-demo/micro-mall-member/global"
	"mall-demo/micro-mall-member/model"
	proto_member "mall-demo/micro-mall-member/proto/micro-mall-member-proto"
	"math"
)

type MemberLevelService struct {
}

func (obj *MemberLevelService) GetMemberLevelList(ctx context.Context, req *proto_member.GetMemberLevelListRequest) (*proto_member.GetMemberLevelListResponse, error) {
	var memberLevels []model.UmsMemberLevel
	global.UmsMysqlConn.Model(&model.UmsMemberLevel{}).Find(&memberLevels)
	var count int64
	global.UmsMysqlConn.Model(&model.UmsMemberLevel{}).Count(&count)

	resp := proto_member.GetMemberLevelListResponse{
		TotalCount:          int32(len(memberLevels)),
		PageSize:            req.PageSize,
		TotalPage:           int32(math.Ceil(float64(count * 1.0 / int64(req.PageSize)))),
		CurrPage:            req.PageNum,
		MemberLevelEntities: make([]*proto_member.MemberLevelEntity, 0, 10),
	}

	for i := 0; i < len(memberLevels); i++ {
		resp.MemberLevelEntities = append(resp.MemberLevelEntities, &proto_member.MemberLevelEntity{
			Id:                    memberLevels[i].Id,
			Name:                  memberLevels[i].Name,
			GrowthPoint:           int32(memberLevels[i].GrowthPoint),
			DefaultStatus:         int32(memberLevels[i].DefaultStatus),
			FreeFreightPoint:      float32(memberLevels[i].FreeFreightPoint),
			CommentGrowthPoint:    int32(memberLevels[i].CommentGrowthPoint),
			PriviledgeFreeFreight: int32(memberLevels[i].PriviledgeFreeFreight),
			PriviledgeMemberPrice: int32(memberLevels[i].PriviledgeMemberPrice),
			PriviledgeBirthday:    int32(memberLevels[i].PriviledgeBirthday),
			Note:                  memberLevels[i].Note,
		})
	}
	return &resp, nil
}

func (obj *MemberLevelService) SaveMemberLevel(ctx context.Context, req *proto_member.SaveMemberLevelRequest) (*proto_member.SaveMemberLevelResponse, error) {
	memberLevel := model.UmsMemberLevel{
		Id:                    req.MemberLevelEntity.Id,
		Name:                  req.MemberLevelEntity.Name,
		GrowthPoint:           int(req.MemberLevelEntity.GrowthPoint),
		DefaultStatus:         int(req.MemberLevelEntity.DefaultStatus),
		FreeFreightPoint:      float64(req.MemberLevelEntity.FreeFreightPoint),
		CommentGrowthPoint:    int(req.MemberLevelEntity.CommentGrowthPoint),
		PriviledgeFreeFreight: int(req.MemberLevelEntity.PriviledgeFreeFreight),
		PriviledgeMemberPrice: int(req.MemberLevelEntity.PriviledgeMemberPrice),
		PriviledgeBirthday:    int(req.MemberLevelEntity.PriviledgeBirthday),
		Note:                  req.MemberLevelEntity.Note,
	}
	global.UmsMysqlConn.Save(&memberLevel)
	return &proto_member.SaveMemberLevelResponse{}, nil
}

func (obj *MemberLevelService) UpdateMemberLevel(ctx context.Context, req *proto_member.UpdateMemberLevelRequest) (*proto_member.UpdateMemberLevelResponse, error) {
	memberLevel := model.UmsMemberLevel{
		Id:                    req.MemberLevelEntity.Id,
		Name:                  req.MemberLevelEntity.Name,
		GrowthPoint:           int(req.MemberLevelEntity.GrowthPoint),
		DefaultStatus:         int(req.MemberLevelEntity.DefaultStatus),
		FreeFreightPoint:      float64(req.MemberLevelEntity.FreeFreightPoint),
		CommentGrowthPoint:    int(req.MemberLevelEntity.CommentGrowthPoint),
		PriviledgeFreeFreight: int(req.MemberLevelEntity.PriviledgeFreeFreight),
		PriviledgeMemberPrice: int(req.MemberLevelEntity.PriviledgeMemberPrice),
		PriviledgeBirthday:    int(req.MemberLevelEntity.PriviledgeBirthday),
		Note:                  req.MemberLevelEntity.Note,
	}
	global.UmsMysqlConn.Save(&memberLevel)
	return &proto_member.UpdateMemberLevelResponse{}, nil
}

func (obj *MemberLevelService) DeleteMemberLevel(ctx context.Context, req *proto_member.DeleteMemberLevelRequest) (*proto_member.DeleteMemberLevelResponse, error) {
	global.UmsMysqlConn.Delete(&model.UmsMemberLevel{}, req.Ids)
	return &proto_member.DeleteMemberLevelResponse{}, nil
}

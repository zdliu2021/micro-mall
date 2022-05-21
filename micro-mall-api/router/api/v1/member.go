package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"mall-demo/micro-mall-api/global"
	"mall-demo/micro-mall-api/model/request"
	response2 "mall-demo/micro-mall-api/model/response"
	"mall-demo/micro-mall-api/proto/micro-mall-member-proto"
	"mall-demo/micro-mall-api/rpc-client"
)

func GetMemberLevel(ctx *gin.Context) {
	page := GetPage(ctx)
	rpcClient := rpc_client.GetMemberClient()
	resp, err := rpcClient.GetMemberLevelList(context.TODO(), &proto_member.GetMemberLevelListRequest{
		PageSize: int32(page.PageSize),
		PageNum:  int32(page.PageNum),
		Keyword:  page.Keyword,
		Order:    page.Order,
		Sidx:     page.Sidx,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	var res []response2.MemberLevelEntity
	res = make([]response2.MemberLevelEntity, 0, 10)
	for i := 0; i < len(resp.MemberLevelEntities); i++ {
		res = append(res, response2.MemberLevelEntity{
			Id:                    resp.MemberLevelEntities[i].Id,
			Name:                  resp.MemberLevelEntities[i].Name,
			GrowthPoint:           int32(resp.MemberLevelEntities[i].GrowthPoint),
			DefaultStatus:         int32(resp.MemberLevelEntities[i].DefaultStatus),
			FreeFreightPoint:      resp.MemberLevelEntities[i].FreeFreightPoint,
			CommentGrowthPoint:    int32(resp.MemberLevelEntities[i].CommentGrowthPoint),
			PriviledgeFreeFreight: int32(resp.MemberLevelEntities[i].PriviledgeFreeFreight),
			PriviledgeMemberPrice: int32(resp.MemberLevelEntities[i].PriviledgeMemberPrice),
			PriviledgeBirthday:    int32(resp.MemberLevelEntities[i].PriviledgeBirthday),
			Note:                  resp.MemberLevelEntities[i].Note,
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

func SaveMemberLevel(ctx *gin.Context) {
	var memberLevel request.MemberLevelEntity
	if err := ctx.ShouldBindJSON(&memberLevel); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetMemberClient()
	_, err := rpcClient.SaveMemberLevel(context.TODO(), &proto_member.SaveMemberLevelRequest{
		MemberLevelEntity: &proto_member.MemberLevelEntity{
			Id:                    memberLevel.Id,
			Name:                  memberLevel.Name,
			GrowthPoint:           int32(memberLevel.GrowthPoint),
			DefaultStatus:         int32(memberLevel.DefaultStatus),
			FreeFreightPoint:      memberLevel.FreeFreightPoint,
			CommentGrowthPoint:    int32(memberLevel.CommentGrowthPoint),
			PriviledgeFreeFreight: int32(memberLevel.PriviledgeFreeFreight),
			PriviledgeMemberPrice: int32(memberLevel.PriviledgeMemberPrice),
			PriviledgeBirthday:    int32(memberLevel.PriviledgeBirthday),
			Note:                  memberLevel.Note,
		},
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func UpdateMemberLevel(ctx *gin.Context) {
	var memberLevel request.MemberLevelEntity
	if err := ctx.ShouldBindJSON(&memberLevel); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetMemberClient()
	_, err := rpcClient.UpdateMemberLevel(context.TODO(), &proto_member.UpdateMemberLevelRequest{
		MemberLevelEntity: &proto_member.MemberLevelEntity{
			Id:                    memberLevel.Id,
			Name:                  memberLevel.Name,
			GrowthPoint:           int32(memberLevel.GrowthPoint),
			DefaultStatus:         int32(memberLevel.DefaultStatus),
			FreeFreightPoint:      memberLevel.FreeFreightPoint,
			CommentGrowthPoint:    int32(memberLevel.CommentGrowthPoint),
			PriviledgeFreeFreight: int32(memberLevel.PriviledgeFreeFreight),
			PriviledgeMemberPrice: int32(memberLevel.PriviledgeMemberPrice),
			PriviledgeBirthday:    int32(memberLevel.PriviledgeBirthday),
			Note:                  memberLevel.Note,
		},
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func DeleteMemberLevel(ctx *gin.Context) {
	var req request.DeleteMemberLevelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetMemberClient()
	_, err := rpcClient.DeleteMemberLevel(context.TODO(), &proto_member.DeleteMemberLevelRequest{
		Ids: req.Ids,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func GetMember(ctx *gin.Context) {
	page := GetPage(ctx)
	rpcClient := rpc_client.GetMemberClient()
	resp, err := rpcClient.GetMemberList(context.TODO(), &proto_member.GetMemberListRequest{
		PageNum:  int32(page.PageNum),
		PageSize: int32(page.PageSize),
		Sidx:     page.Sidx,
		Order:    page.Order,
		Keyword:  page.Keyword,
	})

	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	res := make([]response2.MemberEntity, 0, 10)
	for i := 0; i < len(resp.MemberEntities); i++ {
		res = append(res, response2.MemberEntity{
			Id:          resp.MemberEntities[i].Id,
			LevelId:     resp.MemberEntities[i].LevelId,
			UserName:    resp.MemberEntities[i].UserName,
			Password:    resp.MemberEntities[i].Password,
			Nickname:    resp.MemberEntities[i].Nickname,
			Mobile:      resp.MemberEntities[i].Mobile,
			Email:       resp.MemberEntities[i].Email,
			Header:      resp.MemberEntities[i].Header,
			Gender:      int32(resp.MemberEntities[i].Gender),
			Birth:       resp.MemberEntities[i].Birth,
			City:        resp.MemberEntities[i].City,
			Job:         resp.MemberEntities[i].Job,
			Sign:        resp.MemberEntities[i].Sign,
			SourceType:  int32(resp.MemberEntities[i].SourceType),
			Integration: int32(resp.MemberEntities[i].Integration),
			Growth:      int32(resp.MemberEntities[i].Growth),
			Status:      int32(resp.MemberEntities[i].Status),
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

func SaveMember(ctx *gin.Context) {
	var member request.MemberEntity
	if err := ctx.ShouldBindJSON(&member); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcClient := rpc_client.GetMemberClient()
	_, err := rpcClient.SaveMember(context.TODO(), &proto_member.SaveMemberRequest{
		MemberEntity: &proto_member.MemberEntity{
			Id:          member.Id,
			LevelId:     member.LevelId,
			UserName:    member.UserName,
			Password:    member.Password,
			Nickname:    member.Nickname,
			Mobile:      member.Mobile,
			Email:       member.Email,
			Header:      member.Header,
			Gender:      int32(member.Gender),
			Birth:       member.Birth,
			City:        member.City,
			Job:         member.Job,
			Sign:        member.Sign,
			SourceType:  int32(member.SourceType),
			Integration: int32(member.Integration),
			Growth:      int32(member.Growth),
			Status:      int32(member.Status),
		},
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func UpdateMember(ctx *gin.Context) {
	var member request.MemberEntity
	if err := ctx.ShouldBindJSON(&member); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	rpcClient := rpc_client.GetMemberClient()
	_, err := rpcClient.UpdateMember(context.TODO(), &proto_member.UpdateMemberRequest{
		MemberEntity: &proto_member.MemberEntity{
			Id:          member.Id,
			LevelId:     member.LevelId,
			UserName:    member.UserName,
			Password:    member.Password,
			Nickname:    member.Nickname,
			Mobile:      member.Mobile,
			Email:       member.Email,
			Header:      member.Header,
			Gender:      int32(member.Gender),
			Birth:       member.Birth,
			City:        member.City,
			Job:         member.Job,
			Sign:        member.Sign,
			SourceType:  int32(member.SourceType),
			Integration: int32(member.Integration),
			Growth:      int32(member.Growth),
			Status:      int32(member.Status),
		},
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

func DeleteMember(ctx *gin.Context) {
	var members request.DeleteMemberRequest
	if err := ctx.ShouldBindJSON(&members); err != nil {
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	rpcClient := rpc_client.GetMemberClient()
	_, err := rpcClient.DeleteMember(context.TODO(), &proto_member.DeleteMemberRequest{
		Ids: members.Ids,
	})
	if err != nil {
		global.GVA_LOG.Error("list catelog attr group error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}
	response2.Ok(ctx)
}

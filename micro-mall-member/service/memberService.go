package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mall-demo/micro-mall-member/global"
	"mall-demo/micro-mall-member/model"
	proto_member "mall-demo/micro-mall-member/proto/micro-mall-member-proto"
	"net/http"
	"time"
)

type MemberService struct {
}

func (obj *MemberService) GetMemberList(ctx context.Context, req *proto_member.GetMemberListRequest) (*proto_member.GetMemberListResponse, error) {
	var members []model.UmsMember
	global.UmsMysqlConn.Model(&model.UmsMember{}).Offset(int((req.PageNum - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&members)
	resp := proto_member.GetMemberListResponse{
		TotalCount:     int32(len(members)),
		MemberEntities: make([]*proto_member.MemberEntity, 0, 10),
	}
	for i := 0; i < len(members); i++ {
		resp.MemberEntities = append(resp.MemberEntities, &proto_member.MemberEntity{
			Id:          members[i].Id,
			LevelId:     members[i].LevelId,
			UserName:    members[i].Username,
			Password:    members[i].Password,
			Nickname:    members[i].Nickname,
			Mobile:      members[i].Mobile,
			Email:       members[i].Email,
			Header:      members[i].Header,
			Gender:      int32(members[i].Gender),
			Birth:       members[i].Birth.String(),
			City:        members[i].City,
			Job:         members[i].Job,
			Sign:        members[i].Sign,
			SourceType:  int32(members[i].SourceType),
			Integration: int32(members[i].Integration),
			Growth:      int32(members[i].Growth),
			Status:      int32(members[i].Status),
		})
	}
	return &resp, nil
}

func (obj *MemberService) SaveMember(ctx context.Context, req *proto_member.SaveMemberRequest) (*proto_member.SaveMemberResponse, error) {
	t, _ := time.Parse(time.RFC822, req.MemberEntity.Birth)

	member := model.UmsMember{
		Id:          req.MemberEntity.Id,
		LevelId:     req.MemberEntity.LevelId,
		Username:    req.MemberEntity.UserName,
		Password:    req.MemberEntity.Password,
		Nickname:    req.MemberEntity.Nickname,
		Mobile:      req.MemberEntity.Mobile,
		Email:       req.MemberEntity.Email,
		Header:      req.MemberEntity.Header,
		Gender:      int(req.MemberEntity.Gender),
		Birth:       t,
		City:        req.MemberEntity.City,
		Job:         req.MemberEntity.Job,
		Sign:        req.MemberEntity.Sign,
		SourceType:  int(req.MemberEntity.SourceType),
		Integration: int(req.MemberEntity.Integration),
		Growth:      int(req.MemberEntity.Growth),
		Status:      int(req.MemberEntity.Status),
	}
	global.UmsMysqlConn.Save(&member)
	return &proto_member.SaveMemberResponse{}, nil
}

func (obj *MemberService) UpdateMember(ctx context.Context, req *proto_member.UpdateMemberRequest) (*proto_member.UpdateMemberResponse, error) {
	t, _ := time.Parse(time.RFC822, req.MemberEntity.Birth)

	member := model.UmsMember{
		Id:          req.MemberEntity.Id,
		LevelId:     req.MemberEntity.LevelId,
		Username:    req.MemberEntity.UserName,
		Password:    req.MemberEntity.Password,
		Nickname:    req.MemberEntity.Nickname,
		Mobile:      req.MemberEntity.Mobile,
		Email:       req.MemberEntity.Email,
		Header:      req.MemberEntity.Header,
		Gender:      int(req.MemberEntity.Gender),
		Birth:       t,
		City:        req.MemberEntity.City,
		Job:         req.MemberEntity.Job,
		Sign:        req.MemberEntity.Sign,
		SourceType:  int(req.MemberEntity.SourceType),
		Integration: int(req.MemberEntity.Integration),
		Growth:      int(req.MemberEntity.Growth),
		Status:      int(req.MemberEntity.Status),
	}
	global.UmsMysqlConn.Save(&member)
	return &proto_member.UpdateMemberResponse{}, nil
}
func (obj *MemberService) DeleteMember(ctx context.Context, req *proto_member.DeleteMemberRequest) (*proto_member.DeleteMemberResponse, error) {
	global.UmsMysqlConn.Delete(&model.UmsMember{}, req.Ids)
	return &proto_member.DeleteMemberResponse{}, nil
}

func (obj *MemberService) Register(ctx context.Context, req *proto_member.RegisterRequest) (*proto_member.RegisterResponse, error) {
	var memberEntity model.UmsMember
	memberEntity.Username = req.Username

	// 密码加密,
	// TODO 盐值加密
	m := md5.New()
	m.Write([]byte(req.Password))
	memberEntity.Password = hex.EncodeToString(m.Sum(nil))
	memberEntity.Mobile = req.PhoneNumber
	memberEntity.Birth = time.Now()
	memberEntity.CreateTime = time.Now()
	global.UmsMysqlConn.Model(&model.UmsMember{}).Create(&memberEntity)
	return &proto_member.RegisterResponse{}, nil
}

func (obj *MemberService) Login(ctx context.Context, req *proto_member.LoginRequest) (*proto_member.LoginResponse, error) {
	var memberEntity model.UmsMember
	global.UmsMysqlConn.Model(&model.UmsMember{}).Where("username=?", req.Username).First(&memberEntity)
	m := md5.New()
	m.Write([]byte(req.Password))
	if hex.EncodeToString(m.Sum(nil)) != memberEntity.Password {
		return &proto_member.LoginResponse{}, errors.New("账号或密码错误")
	} else {
		return &proto_member.LoginResponse{
			Id:       memberEntity.Id,
			UserName: memberEntity.Username,
		}, nil
	}
}

func (obj *MemberService) SocialLogin(ctx context.Context, req *proto_member.SocialRequest) (*proto_member.SocialResponse, error) {
	fmt.Println(req.AccessToken)
	//params := fmt.Sprintf("access_token=%s", req.AccessToken)
	url := fmt.Sprintf("https://gitee.com/api/v5/user?access_token=%s", req.AccessToken)
	httpReq, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(httpReq)
	fmt.Println(resp.Request.RequestURI)
	defer resp.Body.Close()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	res := model.OAuthGitteResponse{}
	_ = json.Unmarshal(body, &res)

	// 判断用户是否存在
	var count int64
	global.UmsMysqlConn.Model(&model.UmsMember{}).Where("social_id = ?", res.ID).Count(&count)
	if count == 0 {
		memberEntity := model.UmsMember{}
		memberEntity.CreateTime = time.Now()
		memberEntity.Username = res.Name
		memberEntity.SocialID = int64(res.ID)
		memberEntity.Birth = time.Now()
		global.UmsMysqlConn.Model(&model.UmsMember{}).Create(&memberEntity)
	}
	member := model.UmsMember{}
	global.UmsMysqlConn.Model(&model.UmsMember{}).Where("social_id=?", res.ID).First(&member)
	return &proto_member.SocialResponse{
		Id:       member.Id,
		UserName: member.Username,
	}, nil
}

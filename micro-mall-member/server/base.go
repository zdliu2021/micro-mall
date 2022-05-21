package server

import (
	"context"
	proto_member "mall-demo/micro-mall-member/proto/micro-mall-member-proto"
	"mall-demo/micro-mall-member/service"
)

type Server struct {
	proto_member.UnimplementedMemberRpcServer
	MemberLevelService service.MemberLevelService
	MemberService      service.MemberService
}

func (s *Server) GetMemberLevelList(ctx context.Context, req *proto_member.GetMemberLevelListRequest) (*proto_member.GetMemberLevelListResponse, error) {
	return s.MemberLevelService.GetMemberLevelList(ctx, req)
}

func (s *Server) SaveMemberLevel(ctx context.Context, req *proto_member.SaveMemberLevelRequest) (*proto_member.SaveMemberLevelResponse, error) {
	return s.MemberLevelService.SaveMemberLevel(ctx, req)
}
func (s *Server) UpdateMemberLevel(ctx context.Context, req *proto_member.UpdateMemberLevelRequest) (*proto_member.UpdateMemberLevelResponse, error) {
	return s.MemberLevelService.UpdateMemberLevel(ctx, req)
}

func (s *Server) DeleteMemberLevel(ctx context.Context, req *proto_member.DeleteMemberLevelRequest) (*proto_member.DeleteMemberLevelResponse, error) {
	return s.MemberLevelService.DeleteMemberLevel(ctx, req)
}
func (s *Server) GetMemberList(ctx context.Context, req *proto_member.GetMemberListRequest) (*proto_member.GetMemberListResponse, error) {
	return s.MemberService.GetMemberList(ctx, req)
}

func (s *Server) SaveMember(ctx context.Context, req *proto_member.SaveMemberRequest) (*proto_member.SaveMemberResponse, error) {
	return s.MemberService.SaveMember(ctx, req)
}
func (s *Server) UpdateMember(ctx context.Context, req *proto_member.UpdateMemberRequest) (*proto_member.UpdateMemberResponse, error) {
	return s.MemberService.UpdateMember(ctx, req)
}
func (s *Server) DeleteMember(ctx context.Context, req *proto_member.DeleteMemberRequest) (*proto_member.DeleteMemberResponse, error) {
	return s.MemberService.DeleteMember(ctx, req)
}

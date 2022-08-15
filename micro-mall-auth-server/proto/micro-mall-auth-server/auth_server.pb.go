// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: mall-demo/micro-mall-auth-server/proto/micro-mall-auth-server/auth_server.proto

package proto_auth_server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OAuthGitteSuccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *OAuthGitteSuccessRequest) Reset() {
	*x = OAuthGitteSuccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthGitteSuccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthGitteSuccessRequest) ProtoMessage() {}

func (x *OAuthGitteSuccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthGitteSuccessRequest.ProtoReflect.Descriptor instead.
func (*OAuthGitteSuccessRequest) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{0}
}

func (x *OAuthGitteSuccessRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type OAuthGitteSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *OAuthGitteSuccessResponse) Reset() {
	*x = OAuthGitteSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthGitteSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthGitteSuccessResponse) ProtoMessage() {}

func (x *OAuthGitteSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthGitteSuccessResponse.ProtoReflect.Descriptor instead.
func (*OAuthGitteSuccessResponse) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthGitteSuccessResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OAuthGitteSuccessResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type SendCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *SendCodeRequest) Reset() {
	*x = SendCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCodeRequest) ProtoMessage() {}

func (x *SendCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCodeRequest.ProtoReflect.Descriptor instead.
func (*SendCodeRequest) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{2}
}

func (x *SendCodeRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type SendCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendCodeResponse) Reset() {
	*x = SendCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCodeResponse) ProtoMessage() {}

func (x *SendCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCodeResponse.ProtoReflect.Descriptor instead.
func (*SendCodeResponse) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{3}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Code        string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{5}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{6}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP(), []int{7}
}

func (x *LoginResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto protoreflect.FileDescriptor

var file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x6d,
	0x61, 0x6c, 0x6c, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2e,
	0x0a, 0x18, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x69, 0x74, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x48,
	0x0a, 0x19, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x69, 0x74, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0xc5, 0x02, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x70,
	0x63, 0x12, 0x47, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x69, 0x74, 0x74,
	0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x69, 0x74, 0x74,
	0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x47, 0x69, 0x74, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescOnce sync.Once
	file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescData = file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDesc
)

func file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescGZIP() []byte {
	file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescOnce.Do(func() {
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescData)
	})
	return file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDescData
}

var file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_goTypes = []interface{}{
	(*OAuthGitteSuccessRequest)(nil),  // 0: auth_server.OAuthGitteSuccessRequest
	(*OAuthGitteSuccessResponse)(nil), // 1: auth_server.OAuthGitteSuccessResponse
	(*SendCodeRequest)(nil),           // 2: auth_server.SendCodeRequest
	(*SendCodeResponse)(nil),          // 3: auth_server.SendCodeResponse
	(*RegisterRequest)(nil),           // 4: auth_server.RegisterRequest
	(*RegisterResponse)(nil),          // 5: auth_server.RegisterResponse
	(*LoginRequest)(nil),              // 6: auth_server.LoginRequest
	(*LoginResponse)(nil),             // 7: auth_server.LoginResponse
}
var file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_depIdxs = []int32{
	2, // 0: auth_server.AuthServerRpc.SendCode:input_type -> auth_server.SendCodeRequest
	4, // 1: auth_server.AuthServerRpc.Register:input_type -> auth_server.RegisterRequest
	6, // 2: auth_server.AuthServerRpc.Login:input_type -> auth_server.LoginRequest
	0, // 3: auth_server.AuthServerRpc.OAuthGitteSuccess:input_type -> auth_server.OAuthGitteSuccessRequest
	3, // 4: auth_server.AuthServerRpc.SendCode:output_type -> auth_server.SendCodeResponse
	5, // 5: auth_server.AuthServerRpc.Register:output_type -> auth_server.RegisterResponse
	7, // 6: auth_server.AuthServerRpc.Login:output_type -> auth_server.LoginResponse
	1, // 7: auth_server.AuthServerRpc.OAuthGitteSuccess:output_type -> auth_server.OAuthGitteSuccessResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_init()
}
func file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_init() {
	if File_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthGitteSuccessRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthGitteSuccessResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_goTypes,
		DependencyIndexes: file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_depIdxs,
		MessageInfos:      file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_msgTypes,
	}.Build()
	File_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto = out.File
	file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_rawDesc = nil
	file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_goTypes = nil
	file_mall_demo_micro_mall_auth_server_proto_micro_mall_auth_server_auth_server_proto_depIdxs = nil
}

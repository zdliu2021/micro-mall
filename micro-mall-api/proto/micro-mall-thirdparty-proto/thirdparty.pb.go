// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: mall-demo/micro-mall-api/proto/micro-mall-thirdparty-proto/thirdparty.proto

package proto_thirdparty

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

type GetOSSTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOSSTokenRequest) Reset() {
	*x = GetOSSTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOSSTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOSSTokenRequest) ProtoMessage() {}

func (x *GetOSSTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOSSTokenRequest.ProtoReflect.Descriptor instead.
func (*GetOSSTokenRequest) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescGZIP(), []int{0}
}

type GetOSSTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessId  string `protobuf:"bytes,1,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	Host      string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Expire    int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
	Signature string `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Policy    string `protobuf:"bytes,5,opt,name=policy,proto3" json:"policy,omitempty"`
	Dir       string `protobuf:"bytes,6,opt,name=dir,proto3" json:"dir,omitempty"`
	Callback  string `protobuf:"bytes,7,opt,name=callback,proto3" json:"callback,omitempty"`
}

func (x *GetOSSTokenResponse) Reset() {
	*x = GetOSSTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOSSTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOSSTokenResponse) ProtoMessage() {}

func (x *GetOSSTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOSSTokenResponse.ProtoReflect.Descriptor instead.
func (*GetOSSTokenResponse) Descriptor() ([]byte, []int) {
	return file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescGZIP(), []int{1}
}

func (x *GetOSSTokenResponse) GetAccessId() string {
	if x != nil {
		return x.AccessId
	}
	return ""
}

func (x *GetOSSTokenResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GetOSSTokenResponse) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *GetOSSTokenResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *GetOSSTokenResponse) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *GetOSSTokenResponse) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *GetOSSTokenResponse) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

var File_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto protoreflect.FileDescriptor

var file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4f, 0x53, 0x53, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xc2, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x53, 0x53, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x32, 0x5f, 0x0a, 0x0d, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x70, 0x63, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x53, 0x53, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x53, 0x53, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x53, 0x53, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescOnce sync.Once
	file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescData = file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDesc
)

func file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescGZIP() []byte {
	file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescOnce.Do(func() {
		file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescData = protoimpl.X.CompressGZIP(file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescData)
	})
	return file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDescData
}

var file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_goTypes = []interface{}{
	(*GetOSSTokenRequest)(nil),  // 0: thirdparty.GetOSSTokenRequest
	(*GetOSSTokenResponse)(nil), // 1: thirdparty.GetOSSTokenResponse
}
var file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_depIdxs = []int32{
	0, // 0: thirdparty.ThirdPartyRpc.GetOSSToken:input_type -> thirdparty.GetOSSTokenRequest
	1, // 1: thirdparty.ThirdPartyRpc.GetOSSToken:output_type -> thirdparty.GetOSSTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_init() }
func file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_init() {
	if File_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOSSTokenRequest); i {
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
		file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOSSTokenResponse); i {
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
			RawDescriptor: file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_goTypes,
		DependencyIndexes: file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_depIdxs,
		MessageInfos:      file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_msgTypes,
	}.Build()
	File_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto = out.File
	file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_rawDesc = nil
	file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_goTypes = nil
	file_mall_demo_micro_mall_api_proto_micro_mall_thirdparty_proto_thirdparty_proto_depIdxs = nil
}
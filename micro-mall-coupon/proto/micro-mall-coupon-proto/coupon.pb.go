// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: mall-demo/micro-mall-coupon/proto/micro-mall-coupon-proto/coupon.proto

package proto_coupon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto protoreflect.FileDescriptor

var file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_rawDesc = []byte{
	0x0a, 0x46, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x1a, 0x43, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d,
	0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x70, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf6, 0x02, 0x0a, 0x09,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x70, 0x63, 0x12, 0x52, 0x0a, 0x0d, 0x53, 0x61, 0x76,
	0x65, 0x53, 0x6b, 0x75, 0x4c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x73, 0x70, 0x75, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6b, 0x75, 0x4c, 0x61,
	0x64, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x75, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6b, 0x75, 0x4c,
	0x61, 0x64, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a,
	0x14, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6b, 0x75, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70,
	0x75, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6b, 0x75, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x75, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6b,
	0x75, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x53, 0x70,
	0x75, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x73, 0x70, 0x75, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x70, 0x75, 0x42, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x73, 0x70, 0x75, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x70, 0x75, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x53, 0x61,
	0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x75, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x75, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_goTypes = []interface{}{
	(*SaveSkuLadderRequest)(nil),         // 0: proto_spu.SaveSkuLadderRequest
	(*SaveSkuFullReductionRequest)(nil),  // 1: proto_spu.SaveSkuFullReductionRequest
	(*SaveSpuBoundsRequest)(nil),         // 2: proto_spu.SaveSpuBoundsRequest
	(*SaveMemberPriceRequest)(nil),       // 3: proto_spu.SaveMemberPriceRequest
	(*SaveSkuLadderResponse)(nil),        // 4: proto_spu.SaveSkuLadderResponse
	(*SaveSkuFullReductionResponse)(nil), // 5: proto_spu.SaveSkuFullReductionResponse
	(*SaveSpuboundsResponse)(nil),        // 6: proto_spu.SaveSpuboundsResponse
	(*SaveMemberPriceResponse)(nil),      // 7: proto_spu.SaveMemberPriceResponse
}
var file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_depIdxs = []int32{
	0, // 0: proto_coupon.CouponRpc.SaveSkuLadder:input_type -> proto_spu.SaveSkuLadderRequest
	1, // 1: proto_coupon.CouponRpc.SaveSkuFullReduction:input_type -> proto_spu.SaveSkuFullReductionRequest
	2, // 2: proto_coupon.CouponRpc.SaveSpuBounds:input_type -> proto_spu.SaveSpuBoundsRequest
	3, // 3: proto_coupon.CouponRpc.SaveMemberPrice:input_type -> proto_spu.SaveMemberPriceRequest
	4, // 4: proto_coupon.CouponRpc.SaveSkuLadder:output_type -> proto_spu.SaveSkuLadderResponse
	5, // 5: proto_coupon.CouponRpc.SaveSkuFullReduction:output_type -> proto_spu.SaveSkuFullReductionResponse
	6, // 6: proto_coupon.CouponRpc.SaveSpuBounds:output_type -> proto_spu.SaveSpuboundsResponse
	7, // 7: proto_coupon.CouponRpc.SaveMemberPrice:output_type -> proto_spu.SaveMemberPriceResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_init() }
func file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_init() {
	if File_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto != nil {
		return
	}
	file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_spu_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_goTypes,
		DependencyIndexes: file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_depIdxs,
	}.Build()
	File_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto = out.File
	file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_rawDesc = nil
	file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_goTypes = nil
	file_mall_demo_micro_mall_coupon_proto_micro_mall_coupon_proto_coupon_proto_depIdxs = nil
}

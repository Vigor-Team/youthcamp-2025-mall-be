// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.2
// source: product_api.proto

package product

import (
	_ "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/api"
	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
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

type GetProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty" path:"productId"`
}

func (x *GetProductReq) Reset() {
	*x = GetProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReq) ProtoMessage() {}

func (x *GetProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReq.ProtoReflect.Descriptor instead.
func (*GetProductReq) Descriptor() ([]byte, []int) {
	return file_product_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductReq) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type SearchProductsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q string `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty" query:"q"`
}

func (x *SearchProductsReq) Reset() {
	*x = SearchProductsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsReq) ProtoMessage() {}

func (x *SearchProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsReq.ProtoReflect.Descriptor instead.
func (*SearchProductsReq) Descriptor() ([]byte, []int) {
	return file_product_api_proto_rawDescGZIP(), []int{1}
}

func (x *SearchProductsReq) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

type GetCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId uint32 `protobuf:"varint,1,opt,name=categoryId,proto3" json:"categoryId,omitempty" path:"categoryId"`
}

func (x *GetCategoryReq) Reset() {
	*x = GetCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryReq) ProtoMessage() {}

func (x *GetCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryReq.ProtoReflect.Descriptor instead.
func (*GetCategoryReq) Descriptor() ([]byte, []int) {
	return file_product_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetCategoryReq) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

var File_product_api_proto protoreflect.FileDescriptor

var file_product_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0xd2, 0xbb, 0x18, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x05, 0xb2, 0xbb, 0x18, 0x01, 0x71, 0x52, 0x01, 0x71, 0x22, 0x40, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x2e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x0e, 0xd2, 0xbb, 0x18, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x32,
	0xcb, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0xca, 0xc1, 0x18, 0x14, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x3a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x58, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0b,
	0xca, 0xc1, 0x18, 0x07, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x4b, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0d, 0xca, 0xc1, 0x18, 0x09, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0f, 0xca, 0xc1, 0x18, 0x0b, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1b, 0xca, 0xc1, 0x18, 0x17, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x3a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x42, 0x54, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x67, 0x6f,
	0x72, 0x2d, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x68, 0x63, 0x61, 0x6d, 0x70,
	0x2d, 0x32, 0x30, 0x32, 0x35, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x62, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_api_proto_rawDescOnce sync.Once
	file_product_api_proto_rawDescData = file_product_api_proto_rawDesc
)

func file_product_api_proto_rawDescGZIP() []byte {
	file_product_api_proto_rawDescOnce.Do(func() {
		file_product_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_api_proto_rawDescData)
	})
	return file_product_api_proto_rawDescData
}

var file_product_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_product_api_proto_goTypes = []interface{}{
	(*GetProductReq)(nil),     // 0: gateway.product.GetProductReq
	(*SearchProductsReq)(nil), // 1: gateway.product.SearchProductsReq
	(*GetCategoryReq)(nil),    // 2: gateway.product.GetCategoryReq
	(*common.Empty)(nil),      // 3: gateway.common.Empty
}
var file_product_api_proto_depIdxs = []int32{
	0, // 0: gateway.product.ProductService.GetProduct:input_type -> gateway.product.GetProductReq
	1, // 1: gateway.product.ProductService.SearchProducts:input_type -> gateway.product.SearchProductsReq
	3, // 2: gateway.product.ProductService.ListProducts:input_type -> gateway.common.Empty
	3, // 3: gateway.product.ProductService.ListCategories:input_type -> gateway.common.Empty
	2, // 4: gateway.product.ProductService.GetCategory:input_type -> gateway.product.GetCategoryReq
	3, // 5: gateway.product.ProductService.GetProduct:output_type -> gateway.common.Empty
	3, // 6: gateway.product.ProductService.SearchProducts:output_type -> gateway.common.Empty
	3, // 7: gateway.product.ProductService.ListProducts:output_type -> gateway.common.Empty
	3, // 8: gateway.product.ProductService.ListCategories:output_type -> gateway.common.Empty
	3, // 9: gateway.product.ProductService.GetCategory:output_type -> gateway.common.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_api_proto_init() }
func file_product_api_proto_init() {
	if File_product_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReq); i {
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
		file_product_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductsReq); i {
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
		file_product_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryReq); i {
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
			RawDescriptor: file_product_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_api_proto_goTypes,
		DependencyIndexes: file_product_api_proto_depIdxs,
		MessageInfos:      file_product_api_proto_msgTypes,
	}.Build()
	File_product_api_proto = out.File
	file_product_api_proto_rawDesc = nil
	file_product_api_proto_goTypes = nil
	file_product_api_proto_depIdxs = nil
}

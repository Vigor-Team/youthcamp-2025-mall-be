// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.2
// source: cart_api.proto

package cart

import (
	_ "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/api"
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

type AddCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId  uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty" form:"productId"`
	ProductNum int32  `protobuf:"varint,2,opt,name=product_num,json=productNum,proto3" json:"product_num,omitempty" form:"productNum"`
}

func (x *AddCartReq) Reset() {
	*x = AddCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartReq) ProtoMessage() {}

func (x *AddCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartReq.ProtoReflect.Descriptor instead.
func (*AddCartReq) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{0}
}

func (x *AddCartReq) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddCartReq) GetProductNum() int32 {
	if x != nil {
		return x.ProductNum
	}
	return 0
}

type UpdateCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId  uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty" form:"productId"`
	ProductNum int32  `protobuf:"varint,2,opt,name=product_num,json=productNum,proto3" json:"product_num,omitempty" form:"productNum"`
}

func (x *UpdateCartReq) Reset() {
	*x = UpdateCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartReq) ProtoMessage() {}

func (x *UpdateCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartReq.ProtoReflect.Descriptor instead.
func (*UpdateCartReq) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCartReq) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *UpdateCartReq) GetProductNum() int32 {
	if x != nil {
		return x.ProductNum
	}
	return 0
}

type UpdateCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCartResp) Reset() {
	*x = UpdateCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartResp) ProtoMessage() {}

func (x *UpdateCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartResp.ProtoReflect.Descriptor instead.
func (*UpdateCartResp) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{2}
}

type AddCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddCartResp) Reset() {
	*x = AddCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartResp) ProtoMessage() {}

func (x *AddCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartResp.ProtoReflect.Descriptor instead.
func (*AddCartResp) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{3}
}

type GetCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCartReq) Reset() {
	*x = GetCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReq) ProtoMessage() {}

func (x *GetCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReq.ProtoReflect.Descriptor instead.
func (*GetCartReq) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{4}
}

type GetCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart  *Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty" form:"cart" query:"cart"`
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" form:"total" query:"total"`
}

func (x *GetCartResp) Reset() {
	*x = GetCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResp) ProtoMessage() {}

func (x *GetCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResp.ProtoReflect.Descriptor instead.
func (*GetCartResp) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartResp) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *GetCartResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32      `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"`
	Items  []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty" form:"items" query:"items"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{6}
}

func (x *Cart) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   uint32   `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty" form:"productId" query:"productId"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
	SpuName     string   `protobuf:"bytes,3,opt,name=spuName,proto3" json:"spuName,omitempty" form:"spuName" query:"spuName"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" form:"description" query:"description"`
	Picture     string   `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty" form:"picture" query:"picture"`
	Price       float32  `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty" form:"price" query:"price"`
	SpuPrice    float32  `protobuf:"fixed32,7,opt,name=spuPrice,proto3" json:"spuPrice,omitempty" form:"spuPrice" query:"spuPrice"`
	Stock       uint32   `protobuf:"varint,8,opt,name=stock,proto3" json:"stock,omitempty" form:"stock" query:"stock"`
	Categories  []string `protobuf:"bytes,9,rep,name=categories,proto3" json:"categories,omitempty" form:"categories" query:"categories"`
	Quantity    int32    `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty" form:"quantity" query:"quantity"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{7}
}

func (x *CartItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CartItem) GetSpuName() string {
	if x != nil {
		return x.SpuName
	}
	return ""
}

func (x *CartItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CartItem) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CartItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartItem) GetSpuPrice() float32 {
	if x != nil {
		return x.SpuPrice
	}
	return 0
}

func (x *CartItem) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CartItem) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// 删除商品请求和响应
type DeleteCartItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"`
	ProductId uint32 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty" form:"product_id" query:"product_id"`
}

func (x *DeleteCartItemReq) Reset() {
	*x = DeleteCartItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartItemReq) ProtoMessage() {}

func (x *DeleteCartItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartItemReq.ProtoReflect.Descriptor instead.
func (*DeleteCartItemReq) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCartItemReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteCartItemReq) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type DeleteCartItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCartItemResp) Reset() {
	*x = DeleteCartItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartItemResp) ProtoMessage() {}

func (x *DeleteCartItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartItemResp.ProtoReflect.Descriptor instead.
func (*DeleteCartItemResp) Descriptor() ([]byte, []int) {
	return file_cart_api_proto_rawDescGZIP(), []int{9}
}

var File_cart_api_proto protoreflect.FileDescriptor

var file_cart_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x09,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0xe2, 0xbb, 0x18,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0xe2, 0xbb, 0x18, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x6e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0xe2, 0xbb, 0x18,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0xe2, 0xbb, 0x18, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x22, 0x4b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x4d, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x96, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x70, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x70, 0x75, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x73, 0x70, 0x75, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x4b, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x32, 0x97, 0x03,
	0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x73, 0x12, 0x51, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x73, 0x12, 0x6a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0xda, 0xc1, 0x18, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x7d, 0x12, 0x72, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0xe2, 0xc1, 0x18, 0x19, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x7d, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x67, 0x6f, 0x72, 0x2d, 0x54, 0x65, 0x61, 0x6d,
	0x2f, 0x79, 0x6f, 0x75, 0x74, 0x68, 0x63, 0x61, 0x6d, 0x70, 0x2d, 0x32, 0x30, 0x32, 0x35, 0x2d,
	0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cart_api_proto_rawDescOnce sync.Once
	file_cart_api_proto_rawDescData = file_cart_api_proto_rawDesc
)

func file_cart_api_proto_rawDescGZIP() []byte {
	file_cart_api_proto_rawDescOnce.Do(func() {
		file_cart_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_api_proto_rawDescData)
	})
	return file_cart_api_proto_rawDescData
}

var file_cart_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cart_api_proto_goTypes = []interface{}{
	(*AddCartReq)(nil),         // 0: gateway.cart.AddCartReq
	(*UpdateCartReq)(nil),      // 1: gateway.cart.UpdateCartReq
	(*UpdateCartResp)(nil),     // 2: gateway.cart.UpdateCartResp
	(*AddCartResp)(nil),        // 3: gateway.cart.AddCartResp
	(*GetCartReq)(nil),         // 4: gateway.cart.GetCartReq
	(*GetCartResp)(nil),        // 5: gateway.cart.GetCartResp
	(*Cart)(nil),               // 6: gateway.cart.Cart
	(*CartItem)(nil),           // 7: gateway.cart.CartItem
	(*DeleteCartItemReq)(nil),  // 8: gateway.cart.DeleteCartItemReq
	(*DeleteCartItemResp)(nil), // 9: gateway.cart.DeleteCartItemResp
}
var file_cart_api_proto_depIdxs = []int32{
	6, // 0: gateway.cart.GetCartResp.cart:type_name -> gateway.cart.Cart
	7, // 1: gateway.cart.Cart.items:type_name -> gateway.cart.CartItem
	0, // 2: gateway.cart.CartService.AddCartItem:input_type -> gateway.cart.AddCartReq
	4, // 3: gateway.cart.CartService.GetCart:input_type -> gateway.cart.GetCartReq
	1, // 4: gateway.cart.CartService.UpdateCartItem:input_type -> gateway.cart.UpdateCartReq
	8, // 5: gateway.cart.CartService.DeleteCartItem:input_type -> gateway.cart.DeleteCartItemReq
	3, // 6: gateway.cart.CartService.AddCartItem:output_type -> gateway.cart.AddCartResp
	5, // 7: gateway.cart.CartService.GetCart:output_type -> gateway.cart.GetCartResp
	2, // 8: gateway.cart.CartService.UpdateCartItem:output_type -> gateway.cart.UpdateCartResp
	9, // 9: gateway.cart.CartService.DeleteCartItem:output_type -> gateway.cart.DeleteCartItemResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cart_api_proto_init() }
func file_cart_api_proto_init() {
	if File_cart_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartReq); i {
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
		file_cart_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartReq); i {
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
		file_cart_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartResp); i {
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
		file_cart_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartResp); i {
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
		file_cart_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReq); i {
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
		file_cart_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartResp); i {
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
		file_cart_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_cart_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_cart_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartItemReq); i {
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
		file_cart_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartItemResp); i {
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
			RawDescriptor: file_cart_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_api_proto_goTypes,
		DependencyIndexes: file_cart_api_proto_depIdxs,
		MessageInfos:      file_cart_api_proto_msgTypes,
	}.Build()
	File_cart_api_proto = out.File
	file_cart_api_proto_rawDesc = nil
	file_cart_api_proto_goTypes = nil
	file_cart_api_proto_depIdxs = nil
}

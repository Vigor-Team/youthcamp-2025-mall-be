// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.2
// source: order_api.proto

package order

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

type PlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCurrency string       `protobuf:"bytes,1,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty" form:"user_currency" query:"user_currency"`
	Address      *Address     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" form:"address" query:"address"`
	Email        string       `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`
	OrderItems   []*OrderItem `protobuf:"bytes,4,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty" form:"order_items" query:"order_items"`
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{0}
}

func (x *PlaceOrderReq) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *PlaceOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PlaceOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlaceOrderReq) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type PlaceOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty" form:"order_id" query:"order_id"`
}

func (x *PlaceOrderResp) Reset() {
	*x = PlaceOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResp) ProtoMessage() {}

func (x *PlaceOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResp.ProtoReflect.Descriptor instead.
func (*PlaceOrderResp) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{1}
}

func (x *PlaceOrderResp) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type SeckillPrePlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" form:"product_id" json:"product_id,omitempty"`
}

func (x *SeckillPrePlaceOrderReq) Reset() {
	*x = SeckillPrePlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillPrePlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillPrePlaceOrderReq) ProtoMessage() {}

func (x *SeckillPrePlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillPrePlaceOrderReq.ProtoReflect.Descriptor instead.
func (*SeckillPrePlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{2}
}

func (x *SeckillPrePlaceOrderReq) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type SeckillPrePlaceOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreOrderId uint32 `protobuf:"varint,1,opt,name=pre_order_id,json=preOrderId,proto3" json:"pre_order_id,omitempty" form:"pre_order_id" query:"pre_order_id"`
}

func (x *SeckillPrePlaceOrderResp) Reset() {
	*x = SeckillPrePlaceOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillPrePlaceOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillPrePlaceOrderResp) ProtoMessage() {}

func (x *SeckillPrePlaceOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillPrePlaceOrderResp.ProtoReflect.Descriptor instead.
func (*SeckillPrePlaceOrderResp) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{3}
}

func (x *SeckillPrePlaceOrderResp) GetPreOrderId() uint32 {
	if x != nil {
		return x.PreOrderId
	}
	return 0
}

type SeckillPlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCurrency string   `protobuf:"bytes,1,opt,name=user_currency,json=userCurrency,proto3" form:"user_currency" json:"user_currency,omitempty"`
	Address      *Address `protobuf:"bytes,2,opt,name=address,proto3" form:"address" json:"address,omitempty"`
	Email        string   `protobuf:"bytes,3,opt,name=email,proto3" form:"email" json:"email,omitempty"`
	PreOrderId   uint32   `protobuf:"varint,4,opt,name=pre_order_id,json=preOrderId,proto3" json:"pre_order_id,omitempty" form:"pre_order_id" query:"pre_order_id"`
}

func (x *SeckillPlaceOrderReq) Reset() {
	*x = SeckillPlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillPlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillPlaceOrderReq) ProtoMessage() {}

func (x *SeckillPlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillPlaceOrderReq.ProtoReflect.Descriptor instead.
func (*SeckillPlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{4}
}

func (x *SeckillPlaceOrderReq) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *SeckillPlaceOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *SeckillPlaceOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SeckillPlaceOrderReq) GetPreOrderId() uint32 {
	if x != nil {
		return x.PreOrderId
	}
	return 0
}

type SeckillPlaceOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty" form:"order_id" query:"order_id"`
}

func (x *SeckillPlaceOrderResp) Reset() {
	*x = SeckillPlaceOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillPlaceOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillPlaceOrderResp) ProtoMessage() {}

func (x *SeckillPlaceOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillPlaceOrderResp.ProtoReflect.Descriptor instead.
func (*SeckillPlaceOrderResp) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{5}
}

func (x *SeckillPlaceOrderResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SeckillPlaceOrderResp) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty" form:"street_address" query:"street_address"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty" form:"city" query:"city"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty" form:"state" query:"state"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty" form:"country" query:"country"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty" form:"zip_code" query:"zip_code"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{6}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

type QueryOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty" path:"order_id"`
}

func (x *QueryOrderReq) Reset() {
	*x = QueryOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOrderReq) ProtoMessage() {}

func (x *QueryOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOrderReq.ProtoReflect.Descriptor instead.
func (*QueryOrderReq) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{7}
}

func (x *QueryOrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type QueryOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty" form:"order" query:"order"`
}

func (x *QueryOrderResp) Reset() {
	*x = QueryOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOrderResp) ProtoMessage() {}

func (x *QueryOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOrderResp.ProtoReflect.Descriptor instead.
func (*QueryOrderResp) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{8}
}

func (x *QueryOrderResp) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderItems   []*OrderItem `protobuf:"bytes,1,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty" form:"order_items" query:"order_items"`
	OrderId      string       `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty" form:"order_id" query:"order_id"`
	UserId       uint32       `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"`
	UserCurrency string       `protobuf:"bytes,4,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty" form:"user_currency" query:"user_currency"`
	Address      *Address     `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty" form:"address" query:"address"`
	Email        string       `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`
	Status       string       `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	CreatedAt    int32        `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" form:"created_at" query:"created_at"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{9}
}

func (x *Order) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *Order) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Order) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty" form:"item" query:"item"`
	Cost float32   `protobuf:"fixed32,2,opt,name=cost,proto3" json:"cost,omitempty" form:"cost" query:"cost"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_api_proto_rawDescGZIP(), []int{10}
}

func (x *OrderItem) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *OrderItem) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty" form:"product_id" query:"product_id"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty" form:"quantity" query:"quantity"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_proto_msgTypes[11]
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
	return file_order_api_proto_rawDescGZIP(), []int{11}
}

func (x *CartItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_order_api_proto protoreflect.FileDescriptor

var file_order_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x39, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x0e, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x53, 0x65, 0x63, 0x6b,
	0x69, 0x6c, 0x6c, 0x50, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x22, 0x3c, 0x0a, 0x18, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x65,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20,
	0x0a, 0x0c, 0x70, 0x72, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xd0, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x36, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0xca, 0xbb, 0x18, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x3d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xca, 0xbb, 0x18, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x15, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x8f, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x38, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x27, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xd2, 0xbb, 0x18, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x9a, 0x02, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4c, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x2b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xa6, 0x04, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x86, 0x01, 0x0a, 0x14, 0x53,
	0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x65, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x6b,
	0x69, 0x6c, 0x6c, 0x50, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0xd2, 0xc1, 0x18, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x2f,
	0x70, 0x72, 0x65, 0x12, 0x79, 0x0a, 0x11, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x19, 0xd2, 0xc1, 0x18, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x66,
	0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0xca, 0xc1, 0x18, 0x17, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x3a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x5c, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x67, 0x6f, 0x72, 0x2d, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x79, 0x6f,
	0x75, 0x74, 0x68, 0x63, 0x61, 0x6d, 0x70, 0x2d, 0x32, 0x30, 0x32, 0x35, 0x2d, 0x6d, 0x61, 0x6c,
	0x6c, 0x2d, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_api_proto_rawDescOnce sync.Once
	file_order_api_proto_rawDescData = file_order_api_proto_rawDesc
)

func file_order_api_proto_rawDescGZIP() []byte {
	file_order_api_proto_rawDescOnce.Do(func() {
		file_order_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_api_proto_rawDescData)
	})
	return file_order_api_proto_rawDescData
}

var file_order_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_order_api_proto_goTypes = []interface{}{
	(*PlaceOrderReq)(nil),            // 0: gateway.order.PlaceOrderReq
	(*PlaceOrderResp)(nil),           // 1: gateway.order.PlaceOrderResp
	(*SeckillPrePlaceOrderReq)(nil),  // 2: gateway.order.SeckillPrePlaceOrderReq
	(*SeckillPrePlaceOrderResp)(nil), // 3: gateway.order.SeckillPrePlaceOrderResp
	(*SeckillPlaceOrderReq)(nil),     // 4: gateway.order.SeckillPlaceOrderReq
	(*SeckillPlaceOrderResp)(nil),    // 5: gateway.order.SeckillPlaceOrderResp
	(*Address)(nil),                  // 6: gateway.order.Address
	(*QueryOrderReq)(nil),            // 7: gateway.order.QueryOrderReq
	(*QueryOrderResp)(nil),           // 8: gateway.order.QueryOrderResp
	(*Order)(nil),                    // 9: gateway.order.Order
	(*OrderItem)(nil),                // 10: gateway.order.OrderItem
	(*CartItem)(nil),                 // 11: gateway.order.CartItem
	(*common.Empty)(nil),             // 12: gateway.common.Empty
}
var file_order_api_proto_depIdxs = []int32{
	6,  // 0: gateway.order.PlaceOrderReq.address:type_name -> gateway.order.Address
	10, // 1: gateway.order.PlaceOrderReq.order_items:type_name -> gateway.order.OrderItem
	6,  // 2: gateway.order.SeckillPlaceOrderReq.address:type_name -> gateway.order.Address
	9,  // 3: gateway.order.QueryOrderResp.order:type_name -> gateway.order.Order
	10, // 4: gateway.order.Order.order_items:type_name -> gateway.order.OrderItem
	6,  // 5: gateway.order.Order.address:type_name -> gateway.order.Address
	11, // 6: gateway.order.OrderItem.item:type_name -> gateway.order.CartItem
	12, // 7: gateway.order.OrderService.OrderList:input_type -> gateway.common.Empty
	2,  // 8: gateway.order.OrderService.SeckillPrePlaceOrder:input_type -> gateway.order.SeckillPrePlaceOrderReq
	4,  // 9: gateway.order.OrderService.SeckillPlaceOrder:input_type -> gateway.order.SeckillPlaceOrderReq
	7,  // 10: gateway.order.OrderService.QueryOrder:input_type -> gateway.order.QueryOrderReq
	0,  // 11: gateway.order.OrderService.PlaceOrder:input_type -> gateway.order.PlaceOrderReq
	12, // 12: gateway.order.OrderService.OrderList:output_type -> gateway.common.Empty
	3,  // 13: gateway.order.OrderService.SeckillPrePlaceOrder:output_type -> gateway.order.SeckillPrePlaceOrderResp
	5,  // 14: gateway.order.OrderService.SeckillPlaceOrder:output_type -> gateway.order.SeckillPlaceOrderResp
	8,  // 15: gateway.order.OrderService.QueryOrder:output_type -> gateway.order.QueryOrderResp
	1,  // 16: gateway.order.OrderService.PlaceOrder:output_type -> gateway.order.PlaceOrderResp
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_order_api_proto_init() }
func file_order_api_proto_init() {
	if File_order_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderReq); i {
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
		file_order_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderResp); i {
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
		file_order_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeckillPrePlaceOrderReq); i {
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
		file_order_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeckillPrePlaceOrderResp); i {
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
		file_order_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeckillPlaceOrderReq); i {
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
		file_order_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeckillPlaceOrderResp); i {
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
		file_order_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_order_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOrderReq); i {
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
		file_order_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOrderResp); i {
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
		file_order_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_order_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_order_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_api_proto_goTypes,
		DependencyIndexes: file_order_api_proto_depIdxs,
		MessageInfos:      file_order_api_proto_msgTypes,
	}.Build()
	File_order_api_proto = out.File
	file_order_api_proto_rawDesc = nil
	file_order_api_proto_goTypes = nil
	file_order_api_proto_depIdxs = nil
}

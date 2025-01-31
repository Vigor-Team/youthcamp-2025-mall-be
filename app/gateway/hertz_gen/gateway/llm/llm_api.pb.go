// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.2
// source: llm_api.proto

package llm

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

type ResponseStatus int32

const (
	ResponseStatus_SUCCESS    ResponseStatus = 0
	ResponseStatus_ERROR      ResponseStatus = 1
	ResponseStatus_PROCESSING ResponseStatus = 2
)

// Enum value maps for ResponseStatus.
var (
	ResponseStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
		2: "PROCESSING",
	}
	ResponseStatus_value = map[string]int32{
		"SUCCESS":    0,
		"ERROR":      1,
		"PROCESSING": 2,
	}
)

func (x ResponseStatus) Enum() *ResponseStatus {
	p := new(ResponseStatus)
	*p = x
	return p
}

func (x ResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_llm_api_proto_enumTypes[0].Descriptor()
}

func (ResponseStatus) Type() protoreflect.EnumType {
	return &file_llm_api_proto_enumTypes[0]
}

func (x ResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseStatus.Descriptor instead.
func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_llm_api_proto_rawDescGZIP(), []int{0}
}

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"`
	Message        string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" form:"message" query:"message"`
	ConversationId string `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty" form:"conversation_id" query:"conversation_id"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_llm_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_llm_api_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChatRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatRequest) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string         `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty" form:"response" query:"response"`
	Status   ResponseStatus `protobuf:"varint,2,opt,name=status,proto3,enum=gateway.chat.ResponseStatus" json:"status,omitempty" form:"status" query:"status"`
	Error    string         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty" form:"error" query:"error"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_llm_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_llm_api_proto_rawDescGZIP(), []int{1}
}

func (x *ChatResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *ChatResponse) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_SUCCESS
}

func (x *ChatResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_llm_api_proto protoreflect.FileDescriptor

var file_llm_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6c, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x6c, 0x6d, 0x1a, 0x09, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x75, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x38, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x32, 0xc5, 0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x74, 0x48, 0x74, 0x74, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0xd2, 0xc1, 0x18,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x5b,
	0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0xd2, 0xc1, 0x18, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x30, 0x01, 0x42, 0x50, 0x5a, 0x4e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x67, 0x6f, 0x72, 0x2d,
	0x54, 0x65, 0x61, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x68, 0x63, 0x61, 0x6d, 0x70, 0x2d, 0x32,
	0x30, 0x32, 0x35, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6c, 0x6c, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_llm_api_proto_rawDescOnce sync.Once
	file_llm_api_proto_rawDescData = file_llm_api_proto_rawDesc
)

func file_llm_api_proto_rawDescGZIP() []byte {
	file_llm_api_proto_rawDescOnce.Do(func() {
		file_llm_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_llm_api_proto_rawDescData)
	})
	return file_llm_api_proto_rawDescData
}

var file_llm_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_llm_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_llm_api_proto_goTypes = []interface{}{
	(ResponseStatus)(0),  // 0: gateway.chat.ResponseStatus
	(*ChatRequest)(nil),  // 1: gateway.chat.ChatRequest
	(*ChatResponse)(nil), // 2: gateway.chat.ChatResponse
}
var file_llm_api_proto_depIdxs = []int32{
	0, // 0: gateway.chat.ChatResponse.status:type_name -> gateway.chat.ResponseStatus
	1, // 1: gateway.chat.ChatHttpService.SendMessage:input_type -> gateway.chat.ChatRequest
	1, // 2: gateway.chat.ChatHttpService.StreamMessage:input_type -> gateway.chat.ChatRequest
	2, // 3: gateway.chat.ChatHttpService.SendMessage:output_type -> gateway.chat.ChatResponse
	2, // 4: gateway.chat.ChatHttpService.StreamMessage:output_type -> gateway.chat.ChatResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_llm_api_proto_init() }
func file_llm_api_proto_init() {
	if File_llm_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_llm_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_llm_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
			RawDescriptor: file_llm_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_llm_api_proto_goTypes,
		DependencyIndexes: file_llm_api_proto_depIdxs,
		EnumInfos:         file_llm_api_proto_enumTypes,
		MessageInfos:      file_llm_api_proto_msgTypes,
	}.Build()
	File_llm_api_proto = out.File
	file_llm_api_proto_rawDesc = nil
	file_llm_api_proto_goTypes = nil
	file_llm_api_proto_depIdxs = nil
}

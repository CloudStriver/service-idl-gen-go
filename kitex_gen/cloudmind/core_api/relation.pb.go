// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/core_api/relation.proto

package core_api

import (
	context "context"
	basic "github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
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

type CreateRelationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation *RelationInfo `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	IsOnly   bool          `protobuf:"varint,2,opt,name=isOnly,proto3" json:"isOnly,omitempty"` // 是否唯一
}

func (x *CreateRelationReq) Reset() {
	*x = CreateRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationReq) ProtoMessage() {}

func (x *CreateRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationReq.ProtoReflect.Descriptor instead.
func (*CreateRelationReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRelationReq) GetRelation() *RelationInfo {
	if x != nil {
		return x.Relation
	}
	return nil
}

func (x *CreateRelationReq) GetIsOnly() bool {
	if x != nil {
		return x.IsOnly
	}
	return false
}

type CreateRelationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRelationResp) Reset() {
	*x = CreateRelationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationResp) ProtoMessage() {}

func (x *CreateRelationResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationResp.ProtoReflect.Descriptor instead.
func (*CreateRelationResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{1}
}

type GetRelationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationFilterOptions *RelationFilterOptions   `protobuf:"bytes,1,opt,name=relationFilterOptions,proto3" json:"relationFilterOptions,omitempty"`
	PaginationOptions     *basic.PaginationOptions `protobuf:"bytes,2,opt,name=paginationOptions,proto3" json:"paginationOptions,omitempty"`
}

func (x *GetRelationsReq) Reset() {
	*x = GetRelationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsReq) ProtoMessage() {}

func (x *GetRelationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsReq.ProtoReflect.Descriptor instead.
func (*GetRelationsReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GetRelationsReq) GetRelationFilterOptions() *RelationFilterOptions {
	if x != nil {
		return x.RelationFilterOptions
	}
	return nil
}

func (x *GetRelationsReq) GetPaginationOptions() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOptions
	}
	return nil
}

type GetRelationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRelations []*UserRelation `protobuf:"bytes,1,rep,name=userRelations,proto3" json:"userRelations,omitempty"`
	LastToken     string          `protobuf:"bytes,2,opt,name=lastToken,proto3" json:"lastToken,omitempty"`
	Total         int32           `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetRelationsResp) Reset() {
	*x = GetRelationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsResp) ProtoMessage() {}

func (x *GetRelationsResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsResp.ProtoReflect.Descriptor instead.
func (*GetRelationsResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{3}
}

func (x *GetRelationsResp) GetUserRelations() []*UserRelation {
	if x != nil {
		return x.UserRelations
	}
	return nil
}

func (x *GetRelationsResp) GetLastToken() string {
	if x != nil {
		return x.LastToken
	}
	return ""
}

func (x *GetRelationsResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetRelationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToType       int32  `protobuf:"varint,1,opt,name=toType,proto3" json:"toType,omitempty"`
	ToId         string `protobuf:"bytes,2,opt,name=toId,proto3" json:"toId,omitempty"`
	RelationType int32  `protobuf:"varint,3,opt,name=relationType,proto3" json:"relationType,omitempty"`
}

func (x *GetRelationReq) Reset() {
	*x = GetRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationReq) ProtoMessage() {}

func (x *GetRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationReq.ProtoReflect.Descriptor instead.
func (*GetRelationReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{4}
}

func (x *GetRelationReq) GetToType() int32 {
	if x != nil {
		return x.ToType
	}
	return 0
}

func (x *GetRelationReq) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *GetRelationReq) GetRelationType() int32 {
	if x != nil {
		return x.RelationType
	}
	return 0
}

type GetRelationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
}

func (x *GetRelationResp) Reset() {
	*x = GetRelationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationResp) ProtoMessage() {}

func (x *GetRelationResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationResp.ProtoReflect.Descriptor instead.
func (*GetRelationResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetRelationResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_cloudmind_core_api_relation_proto protoreflect.FileDescriptor

var file_cloudmind_core_api_relation_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x69, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d,
	0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0xba, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x5f, 0x0a, 0x15, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x15, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8e,
	0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x60, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x4f, 0x6b, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudmind_core_api_relation_proto_rawDescOnce sync.Once
	file_cloudmind_core_api_relation_proto_rawDescData = file_cloudmind_core_api_relation_proto_rawDesc
)

func file_cloudmind_core_api_relation_proto_rawDescGZIP() []byte {
	file_cloudmind_core_api_relation_proto_rawDescOnce.Do(func() {
		file_cloudmind_core_api_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudmind_core_api_relation_proto_rawDescData)
	})
	return file_cloudmind_core_api_relation_proto_rawDescData
}

var file_cloudmind_core_api_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloudmind_core_api_relation_proto_goTypes = []interface{}{
	(*CreateRelationReq)(nil),       // 0: cloudmind.core_api.CreateRelationReq
	(*CreateRelationResp)(nil),      // 1: cloudmind.core_api.CreateRelationResp
	(*GetRelationsReq)(nil),         // 2: cloudmind.core_api.GetRelationsReq
	(*GetRelationsResp)(nil),        // 3: cloudmind.core_api.GetRelationsResp
	(*GetRelationReq)(nil),          // 4: cloudmind.core_api.GetRelationReq
	(*GetRelationResp)(nil),         // 5: cloudmind.core_api.GetRelationResp
	(*RelationInfo)(nil),            // 6: cloudmind.core_api.RelationInfo
	(*RelationFilterOptions)(nil),   // 7: cloudmind.core_api.RelationFilterOptions
	(*basic.PaginationOptions)(nil), // 8: basic.PaginationOptions
	(*UserRelation)(nil),            // 9: cloudmind.core_api.UserRelation
}
var file_cloudmind_core_api_relation_proto_depIdxs = []int32{
	6, // 0: cloudmind.core_api.CreateRelationReq.relation:type_name -> cloudmind.core_api.RelationInfo
	7, // 1: cloudmind.core_api.GetRelationsReq.relationFilterOptions:type_name -> cloudmind.core_api.RelationFilterOptions
	8, // 2: cloudmind.core_api.GetRelationsReq.paginationOptions:type_name -> basic.PaginationOptions
	9, // 3: cloudmind.core_api.GetRelationsResp.userRelations:type_name -> cloudmind.core_api.UserRelation
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloudmind_core_api_relation_proto_init() }
func file_cloudmind_core_api_relation_proto_init() {
	if File_cloudmind_core_api_relation_proto != nil {
		return
	}
	file_cloudmind_core_api_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloudmind_core_api_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRelationReq); i {
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
		file_cloudmind_core_api_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRelationResp); i {
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
		file_cloudmind_core_api_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsReq); i {
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
		file_cloudmind_core_api_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsResp); i {
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
		file_cloudmind_core_api_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationReq); i {
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
		file_cloudmind_core_api_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationResp); i {
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
			RawDescriptor: file_cloudmind_core_api_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudmind_core_api_relation_proto_goTypes,
		DependencyIndexes: file_cloudmind_core_api_relation_proto_depIdxs,
		MessageInfos:      file_cloudmind_core_api_relation_proto_msgTypes,
	}.Build()
	File_cloudmind_core_api_relation_proto = out.File
	file_cloudmind_core_api_relation_proto_rawDesc = nil
	file_cloudmind_core_api_relation_proto_goTypes = nil
	file_cloudmind_core_api_relation_proto_depIdxs = nil
}

var _ context.Context

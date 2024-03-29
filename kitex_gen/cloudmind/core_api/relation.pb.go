// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/core_api/relation.proto

package core_api

import (
	context "context"
	_ "github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	_ "github.com/CloudStriver/service-idl-gen-go/kitex_gen/http"
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

	ToId         string       `protobuf:"bytes,1,opt,name=toId,proto3" json:"toId,omitempty"`
	ToType       TargetType   `protobuf:"varint,2,opt,name=toType,proto3,enum=cloudmind.core_api.TargetType" json:"toType,omitempty"`
	RelationType RelationType `protobuf:"varint,3,opt,name=relationType,proto3,enum=cloudmind.core_api.RelationType" json:"relationType,omitempty"`
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

func (x *CreateRelationReq) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *CreateRelationReq) GetToType() TargetType {
	if x != nil {
		return x.ToType
	}
	return TargetType_UnKnowTargetType
}

func (x *CreateRelationReq) GetRelationType() RelationType {
	if x != nil {
		return x.RelationType
	}
	return RelationType_UnKnowRelationType
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

type GetFromRelationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromId       string     `protobuf:"bytes,1,opt,name=fromId,proto3" json:"fromId,omitempty"`
	FromType     int64      `protobuf:"varint,2,opt,name=fromType,proto3" json:"fromType,omitempty"`
	ToType       TargetType `protobuf:"varint,3,opt,name=toType,proto3,enum=cloudmind.core_api.TargetType" json:"toType,omitempty"`
	RelationType int64      `protobuf:"varint,4,opt,name=relationType,proto3" json:"relationType,omitempty"`
	Limit        *int64     `protobuf:"varint,5,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset       *int64     `protobuf:"varint,6,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetFromRelationsReq) Reset() {
	*x = GetFromRelationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFromRelationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFromRelationsReq) ProtoMessage() {}

func (x *GetFromRelationsReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetFromRelationsReq.ProtoReflect.Descriptor instead.
func (*GetFromRelationsReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GetFromRelationsReq) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *GetFromRelationsReq) GetFromType() int64 {
	if x != nil {
		return x.FromType
	}
	return 0
}

func (x *GetFromRelationsReq) GetToType() TargetType {
	if x != nil {
		return x.ToType
	}
	return TargetType_UnKnowTargetType
}

func (x *GetFromRelationsReq) GetRelationType() int64 {
	if x != nil {
		return x.RelationType
	}
	return 0
}

func (x *GetFromRelationsReq) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetFromRelationsReq) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type GetFromRelationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Posts []*Post `protobuf:"bytes,2,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetFromRelationsResp) Reset() {
	*x = GetFromRelationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFromRelationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFromRelationsResp) ProtoMessage() {}

func (x *GetFromRelationsResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetFromRelationsResp.ProtoReflect.Descriptor instead.
func (*GetFromRelationsResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{3}
}

func (x *GetFromRelationsResp) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetFromRelationsResp) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetToRelationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToType       int64      `protobuf:"varint,1,opt,name=toType,proto3" json:"toType,omitempty"`
	ToId         string     `protobuf:"bytes,2,opt,name=toId,proto3" json:"toId,omitempty"`
	FromType     TargetType `protobuf:"varint,3,opt,name=fromType,proto3,enum=cloudmind.core_api.TargetType" json:"fromType,omitempty"`
	RelationType int64      `protobuf:"varint,4,opt,name=relationType,proto3" json:"relationType,omitempty"`
	Limit        *int64     `protobuf:"varint,5,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset       *int64     `protobuf:"varint,6,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetToRelationsReq) Reset() {
	*x = GetToRelationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToRelationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToRelationsReq) ProtoMessage() {}

func (x *GetToRelationsReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetToRelationsReq.ProtoReflect.Descriptor instead.
func (*GetToRelationsReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{4}
}

func (x *GetToRelationsReq) GetToType() int64 {
	if x != nil {
		return x.ToType
	}
	return 0
}

func (x *GetToRelationsReq) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *GetToRelationsReq) GetFromType() TargetType {
	if x != nil {
		return x.FromType
	}
	return TargetType_UnKnowTargetType
}

func (x *GetToRelationsReq) GetRelationType() int64 {
	if x != nil {
		return x.RelationType
	}
	return 0
}

func (x *GetToRelationsReq) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetToRelationsReq) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type GetToRelationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetToRelationsResp) Reset() {
	*x = GetToRelationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToRelationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToRelationsResp) ProtoMessage() {}

func (x *GetToRelationsResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetToRelationsResp.ProtoReflect.Descriptor instead.
func (*GetToRelationsResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetToRelationsResp) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetRelationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromId       string `protobuf:"bytes,1,opt,name=fromId,proto3" json:"fromId,omitempty"`
	FromType     int64  `protobuf:"varint,2,opt,name=fromType,proto3" json:"fromType,omitempty"`
	ToId         string `protobuf:"bytes,3,opt,name=toId,proto3" json:"toId,omitempty"`
	ToType       int64  `protobuf:"varint,4,opt,name=toType,proto3" json:"toType,omitempty"`
	RelationType int64  `protobuf:"varint,5,opt,name=relationType,proto3" json:"relationType,omitempty"`
}

func (x *GetRelationReq) Reset() {
	*x = GetRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationReq) ProtoMessage() {}

func (x *GetRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[6]
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
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{6}
}

func (x *GetRelationReq) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *GetRelationReq) GetFromType() int64 {
	if x != nil {
		return x.FromType
	}
	return 0
}

func (x *GetRelationReq) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *GetRelationReq) GetToType() int64 {
	if x != nil {
		return x.ToType
	}
	return 0
}

func (x *GetRelationReq) GetRelationType() int64 {
	if x != nil {
		return x.RelationType
	}
	return 0
}

type GetRelationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *GetRelationResp) Reset() {
	*x = GetRelationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationResp) ProtoMessage() {}

func (x *GetRelationResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[7]
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
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{7}
}

func (x *GetRelationResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteRelationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToId         string `protobuf:"bytes,1,opt,name=toId,proto3" json:"toId,omitempty"`
	ToType       int64  `protobuf:"varint,2,opt,name=toType,proto3" json:"toType,omitempty"`
	RelationType int64  `protobuf:"varint,3,opt,name=relationType,proto3" json:"relationType,omitempty"`
}

func (x *DeleteRelationReq) Reset() {
	*x = DeleteRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationReq) ProtoMessage() {}

func (x *DeleteRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationReq.ProtoReflect.Descriptor instead.
func (*DeleteRelationReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRelationReq) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *DeleteRelationReq) GetToType() int64 {
	if x != nil {
		return x.ToType
	}
	return 0
}

func (x *DeleteRelationReq) GetRelationType() int64 {
	if x != nil {
		return x.RelationType
	}
	return 0
}

type DeleteRelationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRelationResp) Reset() {
	*x = DeleteRelationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRelationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationResp) ProtoMessage() {}

func (x *DeleteRelationResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_relation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationResp.ProtoReflect.Descriptor instead.
func (*DeleteRelationResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_relation_proto_rawDescGZIP(), []int{9}
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
	0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x06, 0x74,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x14, 0xda, 0xbb, 0x18,
	0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20,
	0x33, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x14, 0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26,
	0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x37, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0xb4, 0x02, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x66,
	0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14, 0xda,
	0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c,
	0x3d, 0x20, 0x33, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a,
	0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x14, 0xda,
	0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c,
	0x3d, 0x20, 0x33, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x14, 0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26,
	0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x37, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x76, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xae, 0x02, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x2c, 0x0a, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x14, 0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20,
	0x24, 0x20, 0x3c, 0x3d, 0x20, 0x33, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f,
	0x49, 0x64, 0x12, 0x50, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x14, 0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31,
	0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x33, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14, 0xda, 0xbb, 0x18, 0x10,
	0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x37,
	0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x22, 0xd6, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14,
	0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20,
	0x3c, 0x3d, 0x20, 0x33, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x14, 0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26,
	0x26, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x33, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14, 0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e, 0x3d,
	0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x37, 0x52, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x8f, 0x01,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14, 0xda, 0xbb, 0x18, 0x10, 0x24, 0x20, 0x3e,
	0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20, 0x33, 0x52, 0x06, 0x74,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14, 0xda, 0xbb, 0x18,
	0x10, 0x24, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x24, 0x20, 0x3c, 0x3d, 0x20,
	0x37, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_cloudmind_core_api_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cloudmind_core_api_relation_proto_goTypes = []interface{}{
	(*CreateRelationReq)(nil),    // 0: cloudmind.core_api.CreateRelationReq
	(*CreateRelationResp)(nil),   // 1: cloudmind.core_api.CreateRelationResp
	(*GetFromRelationsReq)(nil),  // 2: cloudmind.core_api.GetFromRelationsReq
	(*GetFromRelationsResp)(nil), // 3: cloudmind.core_api.GetFromRelationsResp
	(*GetToRelationsReq)(nil),    // 4: cloudmind.core_api.GetToRelationsReq
	(*GetToRelationsResp)(nil),   // 5: cloudmind.core_api.GetToRelationsResp
	(*GetRelationReq)(nil),       // 6: cloudmind.core_api.GetRelationReq
	(*GetRelationResp)(nil),      // 7: cloudmind.core_api.GetRelationResp
	(*DeleteRelationReq)(nil),    // 8: cloudmind.core_api.DeleteRelationReq
	(*DeleteRelationResp)(nil),   // 9: cloudmind.core_api.DeleteRelationResp
	(TargetType)(0),              // 10: cloudmind.core_api.TargetType
	(RelationType)(0),            // 11: cloudmind.core_api.RelationType
	(*User)(nil),                 // 12: cloudmind.core_api.User
	(*Post)(nil),                 // 13: cloudmind.core_api.Post
}
var file_cloudmind_core_api_relation_proto_depIdxs = []int32{
	10, // 0: cloudmind.core_api.CreateRelationReq.toType:type_name -> cloudmind.core_api.TargetType
	11, // 1: cloudmind.core_api.CreateRelationReq.relationType:type_name -> cloudmind.core_api.RelationType
	10, // 2: cloudmind.core_api.GetFromRelationsReq.toType:type_name -> cloudmind.core_api.TargetType
	12, // 3: cloudmind.core_api.GetFromRelationsResp.users:type_name -> cloudmind.core_api.User
	13, // 4: cloudmind.core_api.GetFromRelationsResp.posts:type_name -> cloudmind.core_api.Post
	10, // 5: cloudmind.core_api.GetToRelationsReq.fromType:type_name -> cloudmind.core_api.TargetType
	12, // 6: cloudmind.core_api.GetToRelationsResp.users:type_name -> cloudmind.core_api.User
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
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
			switch v := v.(*GetFromRelationsReq); i {
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
			switch v := v.(*GetFromRelationsResp); i {
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
			switch v := v.(*GetToRelationsReq); i {
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
			switch v := v.(*GetToRelationsResp); i {
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
		file_cloudmind_core_api_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloudmind_core_api_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloudmind_core_api_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRelationReq); i {
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
		file_cloudmind_core_api_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRelationResp); i {
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
	file_cloudmind_core_api_relation_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_cloudmind_core_api_relation_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_core_api_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
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

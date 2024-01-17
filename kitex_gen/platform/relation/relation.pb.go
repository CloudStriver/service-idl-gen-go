// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: platform/relation/relation.proto

package relation

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

	RelationInfo *RelationInfo `protobuf:"bytes,1,opt,name=relationInfo,proto3" json:"relationInfo,omitempty"`
}

func (x *CreateRelationReq) Reset() {
	*x = CreateRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationReq) ProtoMessage() {}

func (x *CreateRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[0]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRelationReq) GetRelationInfo() *RelationInfo {
	if x != nil {
		return x.RelationInfo
	}
	return nil
}

type CreateRelationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRelationResp) Reset() {
	*x = CreateRelationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationResp) ProtoMessage() {}

func (x *CreateRelationResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[1]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{1}
}

type GetRelationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationInfo *RelationInfo `protobuf:"bytes,1,opt,name=relationInfo,proto3" json:"relationInfo,omitempty"`
}

func (x *GetRelationReq) Reset() {
	*x = GetRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationReq) ProtoMessage() {}

func (x *GetRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[2]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GetRelationReq) GetRelationInfo() *RelationInfo {
	if x != nil {
		return x.RelationInfo
	}
	return nil
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
		mi := &file_platform_relation_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationResp) ProtoMessage() {}

func (x *GetRelationResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[3]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{3}
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

	RelationInfo *RelationInfo `protobuf:"bytes,1,opt,name=relationInfo,proto3" json:"relationInfo,omitempty"`
}

func (x *DeleteRelationReq) Reset() {
	*x = DeleteRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationReq) ProtoMessage() {}

func (x *DeleteRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[4]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRelationReq) GetRelationInfo() *RelationInfo {
	if x != nil {
		return x.RelationInfo
	}
	return nil
}

type DeleteRelationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRelationResp) Reset() {
	*x = DeleteRelationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRelationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationResp) ProtoMessage() {}

func (x *DeleteRelationResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[5]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{5}
}

type FromFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromType int64  `protobuf:"varint,1,opt,name=fromType,proto3" json:"fromType,omitempty"`
	FromId   string `protobuf:"bytes,2,opt,name=fromId,proto3" json:"fromId,omitempty"`
	ToType   int64  `protobuf:"varint,3,opt,name=toType,proto3" json:"toType,omitempty"`
}

func (x *FromFilterOptions) Reset() {
	*x = FromFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromFilterOptions) ProtoMessage() {}

func (x *FromFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromFilterOptions.ProtoReflect.Descriptor instead.
func (*FromFilterOptions) Descriptor() ([]byte, []int) {
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{6}
}

func (x *FromFilterOptions) GetFromType() int64 {
	if x != nil {
		return x.FromType
	}
	return 0
}

func (x *FromFilterOptions) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *FromFilterOptions) GetToType() int64 {
	if x != nil {
		return x.ToType
	}
	return 0
}

type ToFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToType   int64  `protobuf:"varint,1,opt,name=toType,proto3" json:"toType,omitempty"`
	ToId     string `protobuf:"bytes,2,opt,name=toId,proto3" json:"toId,omitempty"`
	FromType int64  `protobuf:"varint,3,opt,name=fromType,proto3" json:"fromType,omitempty"`
}

func (x *ToFilterOptions) Reset() {
	*x = ToFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToFilterOptions) ProtoMessage() {}

func (x *ToFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToFilterOptions.ProtoReflect.Descriptor instead.
func (*ToFilterOptions) Descriptor() ([]byte, []int) {
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{7}
}

func (x *ToFilterOptions) GetToType() int64 {
	if x != nil {
		return x.ToType
	}
	return 0
}

func (x *ToFilterOptions) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *ToFilterOptions) GetFromType() int64 {
	if x != nil {
		return x.FromType
	}
	return 0
}

type GetRelationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RelationFilterOptions:
	//
	//	*GetRelationsReq_FromFilterOptions
	//	*GetRelationsReq_ToFilterOptions
	RelationFilterOptions isGetRelationsReq_RelationFilterOptions `protobuf_oneof:"relationFilterOptions"`
	PaginationOptions     *basic.PaginationOptions                `protobuf:"bytes,4,opt,name=paginationOptions,proto3,oneof" json:"paginationOptions,omitempty"`
}

func (x *GetRelationsReq) Reset() {
	*x = GetRelationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsReq) ProtoMessage() {}

func (x *GetRelationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[8]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{8}
}

func (m *GetRelationsReq) GetRelationFilterOptions() isGetRelationsReq_RelationFilterOptions {
	if m != nil {
		return m.RelationFilterOptions
	}
	return nil
}

func (x *GetRelationsReq) GetFromFilterOptions() *FromFilterOptions {
	if x, ok := x.GetRelationFilterOptions().(*GetRelationsReq_FromFilterOptions); ok {
		return x.FromFilterOptions
	}
	return nil
}

func (x *GetRelationsReq) GetToFilterOptions() *ToFilterOptions {
	if x, ok := x.GetRelationFilterOptions().(*GetRelationsReq_ToFilterOptions); ok {
		return x.ToFilterOptions
	}
	return nil
}

func (x *GetRelationsReq) GetPaginationOptions() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOptions
	}
	return nil
}

type isGetRelationsReq_RelationFilterOptions interface {
	isGetRelationsReq_RelationFilterOptions()
}

type GetRelationsReq_FromFilterOptions struct {
	FromFilterOptions *FromFilterOptions `protobuf:"bytes,1,opt,name=fromFilterOptions,proto3,oneof"`
}

type GetRelationsReq_ToFilterOptions struct {
	ToFilterOptions *ToFilterOptions `protobuf:"bytes,2,opt,name=toFilterOptions,proto3,oneof"`
}

func (*GetRelationsReq_FromFilterOptions) isGetRelationsReq_RelationFilterOptions() {}

func (*GetRelationsReq_ToFilterOptions) isGetRelationsReq_RelationFilterOptions() {}

type GetRelationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relations []*Relation `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations,omitempty"`
	Total     int64       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetRelationsResp) Reset() {
	*x = GetRelationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsResp) ProtoMessage() {}

func (x *GetRelationsResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[9]
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
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{9}
}

func (x *GetRelationsResp) GetRelations() []*Relation {
	if x != nil {
		return x.Relations
	}
	return nil
}

func (x *GetRelationsResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetRelationCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RelationFilterOptions:
	//
	//	*GetRelationCountReq_FromFilterOptions
	//	*GetRelationCountReq_ToFilterOptions
	RelationFilterOptions isGetRelationCountReq_RelationFilterOptions `protobuf_oneof:"relationFilterOptions"`
}

func (x *GetRelationCountReq) Reset() {
	*x = GetRelationCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationCountReq) ProtoMessage() {}

func (x *GetRelationCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationCountReq.ProtoReflect.Descriptor instead.
func (*GetRelationCountReq) Descriptor() ([]byte, []int) {
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{10}
}

func (m *GetRelationCountReq) GetRelationFilterOptions() isGetRelationCountReq_RelationFilterOptions {
	if m != nil {
		return m.RelationFilterOptions
	}
	return nil
}

func (x *GetRelationCountReq) GetFromFilterOptions() *FromFilterOptions {
	if x, ok := x.GetRelationFilterOptions().(*GetRelationCountReq_FromFilterOptions); ok {
		return x.FromFilterOptions
	}
	return nil
}

func (x *GetRelationCountReq) GetToFilterOptions() *ToFilterOptions {
	if x, ok := x.GetRelationFilterOptions().(*GetRelationCountReq_ToFilterOptions); ok {
		return x.ToFilterOptions
	}
	return nil
}

type isGetRelationCountReq_RelationFilterOptions interface {
	isGetRelationCountReq_RelationFilterOptions()
}

type GetRelationCountReq_FromFilterOptions struct {
	FromFilterOptions *FromFilterOptions `protobuf:"bytes,1,opt,name=fromFilterOptions,proto3,oneof"`
}

type GetRelationCountReq_ToFilterOptions struct {
	ToFilterOptions *ToFilterOptions `protobuf:"bytes,2,opt,name=toFilterOptions,proto3,oneof"`
}

func (*GetRelationCountReq_FromFilterOptions) isGetRelationCountReq_RelationFilterOptions() {}

func (*GetRelationCountReq_ToFilterOptions) isGetRelationCountReq_RelationFilterOptions() {}

type GetRelationCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetRelationCountResp) Reset() {
	*x = GetRelationCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_relation_relation_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationCountResp) ProtoMessage() {}

func (x *GetRelationCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_relation_relation_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationCountResp.ProtoReflect.Descriptor instead.
func (*GetRelationCountResp) Descriptor() ([]byte, []int) {
	return file_platform_relation_relation_proto_rawDescGZIP(), []int{11}
}

func (x *GetRelationCountResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_platform_relation_relation_proto protoreflect.FileDescriptor

var file_platform_relation_relation_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x43, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x55, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x43, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x58, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x43, 0x0a, 0x0c,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5f, 0x0a, 0x11, 0x46, 0x72, 0x6f, 0x6d, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x22, 0x59, 0x0a, 0x0f, 0x54, 0x6f, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x6f, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xb3, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x54, 0x0a, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x11, 0x66, 0x72, 0x6f, 0x6d,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a,
	0x0f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x6f,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a,
	0x11, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x48, 0x01, 0x52, 0x11, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x63, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a,
	0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xd4,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x54, 0x0a, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x0f,
	0x74, 0x6f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x6f, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x17, 0x0a, 0x15,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x32, 0xe3, 0x03, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x25, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x57, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x63, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x5d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_relation_relation_proto_rawDescOnce sync.Once
	file_platform_relation_relation_proto_rawDescData = file_platform_relation_relation_proto_rawDesc
)

func file_platform_relation_relation_proto_rawDescGZIP() []byte {
	file_platform_relation_relation_proto_rawDescOnce.Do(func() {
		file_platform_relation_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_relation_relation_proto_rawDescData)
	})
	return file_platform_relation_relation_proto_rawDescData
}

var file_platform_relation_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_platform_relation_relation_proto_goTypes = []interface{}{
	(*CreateRelationReq)(nil),       // 0: platform.relation.CreateRelationReq
	(*CreateRelationResp)(nil),      // 1: platform.relation.CreateRelationResp
	(*GetRelationReq)(nil),          // 2: platform.relation.GetRelationReq
	(*GetRelationResp)(nil),         // 3: platform.relation.GetRelationResp
	(*DeleteRelationReq)(nil),       // 4: platform.relation.DeleteRelationReq
	(*DeleteRelationResp)(nil),      // 5: platform.relation.DeleteRelationResp
	(*FromFilterOptions)(nil),       // 6: platform.relation.FromFilterOptions
	(*ToFilterOptions)(nil),         // 7: platform.relation.ToFilterOptions
	(*GetRelationsReq)(nil),         // 8: platform.relation.GetRelationsReq
	(*GetRelationsResp)(nil),        // 9: platform.relation.GetRelationsResp
	(*GetRelationCountReq)(nil),     // 10: platform.relation.GetRelationCountReq
	(*GetRelationCountResp)(nil),    // 11: platform.relation.GetRelationCountResp
	(*RelationInfo)(nil),            // 12: platform.relation.RelationInfo
	(*basic.PaginationOptions)(nil), // 13: basic.PaginationOptions
	(*Relation)(nil),                // 14: platform.relation.Relation
}
var file_platform_relation_relation_proto_depIdxs = []int32{
	12, // 0: platform.relation.CreateRelationReq.relationInfo:type_name -> platform.relation.RelationInfo
	12, // 1: platform.relation.GetRelationReq.relationInfo:type_name -> platform.relation.RelationInfo
	12, // 2: platform.relation.DeleteRelationReq.relationInfo:type_name -> platform.relation.RelationInfo
	6,  // 3: platform.relation.GetRelationsReq.fromFilterOptions:type_name -> platform.relation.FromFilterOptions
	7,  // 4: platform.relation.GetRelationsReq.toFilterOptions:type_name -> platform.relation.ToFilterOptions
	13, // 5: platform.relation.GetRelationsReq.paginationOptions:type_name -> basic.PaginationOptions
	14, // 6: platform.relation.GetRelationsResp.relations:type_name -> platform.relation.Relation
	6,  // 7: platform.relation.GetRelationCountReq.fromFilterOptions:type_name -> platform.relation.FromFilterOptions
	7,  // 8: platform.relation.GetRelationCountReq.toFilterOptions:type_name -> platform.relation.ToFilterOptions
	0,  // 9: platform.relation.RelationService.CreateRelation:input_type -> platform.relation.CreateRelationReq
	8,  // 10: platform.relation.RelationService.GetRelations:input_type -> platform.relation.GetRelationsReq
	10, // 11: platform.relation.RelationService.GetRelationCount:input_type -> platform.relation.GetRelationCountReq
	4,  // 12: platform.relation.RelationService.DeleteRelation:input_type -> platform.relation.DeleteRelationReq
	2,  // 13: platform.relation.RelationService.GetRelation:input_type -> platform.relation.GetRelationReq
	1,  // 14: platform.relation.RelationService.CreateRelation:output_type -> platform.relation.CreateRelationResp
	9,  // 15: platform.relation.RelationService.GetRelations:output_type -> platform.relation.GetRelationsResp
	11, // 16: platform.relation.RelationService.GetRelationCount:output_type -> platform.relation.GetRelationCountResp
	5,  // 17: platform.relation.RelationService.DeleteRelation:output_type -> platform.relation.DeleteRelationResp
	3,  // 18: platform.relation.RelationService.GetRelation:output_type -> platform.relation.GetRelationResp
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_platform_relation_relation_proto_init() }
func file_platform_relation_relation_proto_init() {
	if File_platform_relation_relation_proto != nil {
		return
	}
	file_platform_relation_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_platform_relation_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromFilterOptions); i {
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
		file_platform_relation_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToFilterOptions); i {
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
		file_platform_relation_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_platform_relation_relation_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationCountReq); i {
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
		file_platform_relation_relation_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationCountResp); i {
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
	file_platform_relation_relation_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*GetRelationsReq_FromFilterOptions)(nil),
		(*GetRelationsReq_ToFilterOptions)(nil),
	}
	file_platform_relation_relation_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*GetRelationCountReq_FromFilterOptions)(nil),
		(*GetRelationCountReq_ToFilterOptions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_relation_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_relation_relation_proto_goTypes,
		DependencyIndexes: file_platform_relation_relation_proto_depIdxs,
		MessageInfos:      file_platform_relation_relation_proto_msgTypes,
	}.Build()
	File_platform_relation_relation_proto = out.File
	file_platform_relation_relation_proto_rawDesc = nil
	file_platform_relation_relation_proto_goTypes = nil
	file_platform_relation_relation_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.8.0. DO NOT EDIT.

type RelationService interface {
	CreateRelation(ctx context.Context, req *CreateRelationReq) (res *CreateRelationResp, err error)
	GetRelations(ctx context.Context, req *GetRelationsReq) (res *GetRelationsResp, err error)
	GetRelationCount(ctx context.Context, req *GetRelationCountReq) (res *GetRelationCountResp, err error)
	DeleteRelation(ctx context.Context, req *DeleteRelationReq) (res *DeleteRelationResp, err error)
	GetRelation(ctx context.Context, req *GetRelationReq) (res *GetRelationResp, err error)
}

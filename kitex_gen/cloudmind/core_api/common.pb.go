// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/core_api/common.proto

package core_api

import (
	context "context"
	_ "github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	_ "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/content"
	_ "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/sts"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url    string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_common_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sex         int32  `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	FullName    string `protobuf:"bytes,3,opt,name=fullName,proto3" json:"fullName,omitempty"`
	IdCard      string `protobuf:"bytes,4,opt,name=idCard,proto3" json:"idCard,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_common_proto_rawDescGZIP(), []int{1}
}

func (x *UserDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDetail) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UserDetail) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserDetail) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *UserDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RelationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToType       int32  `protobuf:"varint,1,opt,name=toType,proto3" json:"toType,omitempty"`
	ToId         string `protobuf:"bytes,2,opt,name=toId,proto3" json:"toId,omitempty"`
	RelationType int32  `protobuf:"varint,3,opt,name=relationType,proto3" json:"relationType,omitempty"`
}

func (x *RelationInfo) Reset() {
	*x = RelationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationInfo) ProtoMessage() {}

func (x *RelationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationInfo.ProtoReflect.Descriptor instead.
func (*RelationInfo) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_common_proto_rawDescGZIP(), []int{2}
}

func (x *RelationInfo) GetToType() int32 {
	if x != nil {
		return x.ToType
	}
	return 0
}

func (x *RelationInfo) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *RelationInfo) GetRelationType() int32 {
	if x != nil {
		return x.RelationType
	}
	return 0
}

type RelationFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlyFromType     *int32  `protobuf:"varint,1,opt,name=onlyFromType,proto3,oneof" json:"onlyFromType,omitempty"`
	OnlyFromId       *string `protobuf:"bytes,2,opt,name=onlyFromId,proto3,oneof" json:"onlyFromId,omitempty"`
	OnlyToType       *int32  `protobuf:"varint,3,opt,name=onlyToType,proto3,oneof" json:"onlyToType,omitempty"`
	OnlyToId         *string `protobuf:"bytes,4,opt,name=onlyToId,proto3,oneof" json:"onlyToId,omitempty"`
	OnlyRelationType *int32  `protobuf:"varint,5,opt,name=onlyRelationType,proto3,oneof" json:"onlyRelationType,omitempty"`
}

func (x *RelationFilterOptions) Reset() {
	*x = RelationFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFilterOptions) ProtoMessage() {}

func (x *RelationFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFilterOptions.ProtoReflect.Descriptor instead.
func (*RelationFilterOptions) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_common_proto_rawDescGZIP(), []int{3}
}

func (x *RelationFilterOptions) GetOnlyFromType() int32 {
	if x != nil && x.OnlyFromType != nil {
		return *x.OnlyFromType
	}
	return 0
}

func (x *RelationFilterOptions) GetOnlyFromId() string {
	if x != nil && x.OnlyFromId != nil {
		return *x.OnlyFromId
	}
	return ""
}

func (x *RelationFilterOptions) GetOnlyToType() int32 {
	if x != nil && x.OnlyToType != nil {
		return *x.OnlyToType
	}
	return 0
}

func (x *RelationFilterOptions) GetOnlyToId() string {
	if x != nil && x.OnlyToId != nil {
		return *x.OnlyToId
	}
	return ""
}

func (x *RelationFilterOptions) GetOnlyRelationType() int32 {
	if x != nil && x.OnlyRelationType != nil {
		return *x.OnlyRelationType
	}
	return 0
}

type UserRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	CreateTime   int64 `protobuf:"varint,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime   int64 `protobuf:"varint,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	RelationType int32 `protobuf:"varint,4,opt,name=RelationType,proto3" json:"RelationType,omitempty"`
	IsToMe       bool  `protobuf:"varint,5,opt,name=IsToMe,proto3" json:"IsToMe,omitempty"`
}

func (x *UserRelation) Reset() {
	*x = UserRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_core_api_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelation) ProtoMessage() {}

func (x *UserRelation) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_core_api_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelation.ProtoReflect.Descriptor instead.
func (*UserRelation) Descriptor() ([]byte, []int) {
	return file_cloudmind_core_api_common_proto_rawDescGZIP(), []int{4}
}

func (x *UserRelation) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserRelation) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *UserRelation) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *UserRelation) GetRelationType() int32 {
	if x != nil {
		return x.RelationType
	}
	return 0
}

func (x *UserRelation) GetIsToMe() bool {
	if x != nil {
		return x.IsToMe
	}
	return false
}

var File_cloudmind_core_api_common_proto protoreflect.FileDescriptor

var file_cloudmind_core_api_common_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64,
	0x2f, 0x73, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x88, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x0c, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x6f, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xad, 0x02, 0x0a, 0x15,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x6f,
	0x6e, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x6f, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x54,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6f, 0x6e, 0x6c, 0x79,
	0x54, 0x6f, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x6f, 0x6e,
	0x6c, 0x79, 0x54, 0x6f, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x6f, 0x6e, 0x6c,
	0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x10, 0x6f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x54, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x54, 0x6f, 0x49, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x73, 0x54, 0x6f, 0x4d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x49, 0x73, 0x54, 0x6f, 0x4d, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudmind_core_api_common_proto_rawDescOnce sync.Once
	file_cloudmind_core_api_common_proto_rawDescData = file_cloudmind_core_api_common_proto_rawDesc
)

func file_cloudmind_core_api_common_proto_rawDescGZIP() []byte {
	file_cloudmind_core_api_common_proto_rawDescOnce.Do(func() {
		file_cloudmind_core_api_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudmind_core_api_common_proto_rawDescData)
	})
	return file_cloudmind_core_api_common_proto_rawDescData
}

var file_cloudmind_core_api_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloudmind_core_api_common_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: cloudmind.core_api.User
	(*UserDetail)(nil),            // 1: cloudmind.core_api.UserDetail
	(*RelationInfo)(nil),          // 2: cloudmind.core_api.RelationInfo
	(*RelationFilterOptions)(nil), // 3: cloudmind.core_api.RelationFilterOptions
	(*UserRelation)(nil),          // 4: cloudmind.core_api.UserRelation
}
var file_cloudmind_core_api_common_proto_depIdxs = []int32{
	0, // 0: cloudmind.core_api.UserRelation.user:type_name -> cloudmind.core_api.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloudmind_core_api_common_proto_init() }
func file_cloudmind_core_api_common_proto_init() {
	if File_cloudmind_core_api_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudmind_core_api_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_cloudmind_core_api_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
		file_cloudmind_core_api_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationInfo); i {
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
		file_cloudmind_core_api_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFilterOptions); i {
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
		file_cloudmind_core_api_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelation); i {
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
	file_cloudmind_core_api_common_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_core_api_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudmind_core_api_common_proto_goTypes,
		DependencyIndexes: file_cloudmind_core_api_common_proto_depIdxs,
		MessageInfos:      file_cloudmind_core_api_common_proto_msgTypes,
	}.Build()
	File_cloudmind_core_api_common_proto = out.File
	file_cloudmind_core_api_common_proto_rawDesc = nil
	file_cloudmind_core_api_common_proto_goTypes = nil
	file_cloudmind_core_api_common_proto_depIdxs = nil
}

var _ context.Context
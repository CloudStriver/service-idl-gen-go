// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/sts/common.proto

package sts

import (
	context "context"
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

type AuthType int32

const (
	AuthType_UnKnownType AuthType = 0
	AuthType_email       AuthType = 1
	AuthType_qq          AuthType = 2
	AuthType_wechat      AuthType = 3
	AuthType_github      AuthType = 4
	AuthType_gitee       AuthType = 5
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "UnKnownType",
		1: "email",
		2: "qq",
		3: "wechat",
		4: "github",
		5: "gitee",
	}
	AuthType_value = map[string]int32{
		"UnKnownType": 0,
		"email":       1,
		"qq":          2,
		"wechat":      3,
		"github":      4,
		"gitee":       5,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudmind_sts_common_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_cloudmind_sts_common_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{0}
}

type Role int32

const (
	Role_UnKnownRole Role = 0
	Role_system      Role = 1
	Role_user        Role = 2
	Role_test        Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "UnKnownRole",
		1: "system",
		2: "user",
		3: "test",
	}
	Role_value = map[string]int32{
		"UnKnownRole": 0,
		"system":      1,
		"user":        2,
		"test":        3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudmind_sts_common_proto_enumTypes[1].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_cloudmind_sts_common_proto_enumTypes[1]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Role     int32  `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_common_proto_msgTypes[0]
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
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role     Role    `protobuf:"varint,1,opt,name=role,proto3,enum=cloudmind.sts.Role" json:"role,omitempty"`
	Password *string `protobuf:"bytes,2,opt,name=password,proto3,oneof" json:"password,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_UnKnownRole
}

func (x *UserInfo) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType int32  `protobuf:"varint,1,opt,name=authType,proto3" json:"authType,omitempty"`
	AppId    string `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	UnionId  string `protobuf:"bytes,3,opt,name=unionId,proto3" json:"unionId,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{2}
}

func (x *Auth) GetAuthType() int32 {
	if x != nil {
		return x.AuthType
	}
	return 0
}

func (x *Auth) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Auth) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

type AuthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType AuthType `protobuf:"varint,1,opt,name=authType,proto3,enum=cloudmind.sts.AuthType" json:"authType,omitempty"`
	AppId    string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	UnionId  string   `protobuf:"bytes,3,opt,name=unionId,proto3" json:"unionId,omitempty"`
}

func (x *AuthInfo) Reset() {
	*x = AuthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfo) ProtoMessage() {}

func (x *AuthInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInfo.ProtoReflect.Descriptor instead.
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{3}
}

func (x *AuthInfo) GetAuthType() AuthType {
	if x != nil {
		return x.AuthType
	}
	return AuthType_UnKnownType
}

func (x *AuthInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *AuthInfo) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

type EmailOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *EmailOptions) Reset() {
	*x = EmailOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailOptions) ProtoMessage() {}

func (x *EmailOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailOptions.ProtoReflect.Descriptor instead.
func (*EmailOptions) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{4}
}

func (x *EmailOptions) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailOptions) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type UserIdOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserIdOptions) Reset() {
	*x = UserIdOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdOptions) ProtoMessage() {}

func (x *UserIdOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdOptions.ProtoReflect.Descriptor instead.
func (*UserIdOptions) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_common_proto_rawDescGZIP(), []int{5}
}

func (x *UserIdOptions) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserIdOptions) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_cloudmind_sts_common_proto protoreflect.FileDescriptor

var file_cloudmind_sts_common_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x22, 0x4e, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x61, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e,
	0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x52,
	0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x6f, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33,
	0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x43, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2a, 0x51, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x71, 0x71,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x67, 0x69,
	0x74, 0x65, 0x65, 0x10, 0x05, 0x2a, 0x37, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x10, 0x03, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x53, 0x74, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64,
	0x2f, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudmind_sts_common_proto_rawDescOnce sync.Once
	file_cloudmind_sts_common_proto_rawDescData = file_cloudmind_sts_common_proto_rawDesc
)

func file_cloudmind_sts_common_proto_rawDescGZIP() []byte {
	file_cloudmind_sts_common_proto_rawDescOnce.Do(func() {
		file_cloudmind_sts_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudmind_sts_common_proto_rawDescData)
	})
	return file_cloudmind_sts_common_proto_rawDescData
}

var file_cloudmind_sts_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloudmind_sts_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloudmind_sts_common_proto_goTypes = []interface{}{
	(AuthType)(0),         // 0: cloudmind.sts.AuthType
	(Role)(0),             // 1: cloudmind.sts.Role
	(*User)(nil),          // 2: cloudmind.sts.User
	(*UserInfo)(nil),      // 3: cloudmind.sts.UserInfo
	(*Auth)(nil),          // 4: cloudmind.sts.Auth
	(*AuthInfo)(nil),      // 5: cloudmind.sts.AuthInfo
	(*EmailOptions)(nil),  // 6: cloudmind.sts.EmailOptions
	(*UserIdOptions)(nil), // 7: cloudmind.sts.UserIdOptions
}
var file_cloudmind_sts_common_proto_depIdxs = []int32{
	1, // 0: cloudmind.sts.UserInfo.role:type_name -> cloudmind.sts.Role
	0, // 1: cloudmind.sts.AuthInfo.authType:type_name -> cloudmind.sts.AuthType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloudmind_sts_common_proto_init() }
func file_cloudmind_sts_common_proto_init() {
	if File_cloudmind_sts_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudmind_sts_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloudmind_sts_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_cloudmind_sts_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_cloudmind_sts_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfo); i {
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
		file_cloudmind_sts_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailOptions); i {
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
		file_cloudmind_sts_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdOptions); i {
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
	file_cloudmind_sts_common_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_sts_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudmind_sts_common_proto_goTypes,
		DependencyIndexes: file_cloudmind_sts_common_proto_depIdxs,
		EnumInfos:         file_cloudmind_sts_common_proto_enumTypes,
		MessageInfos:      file_cloudmind_sts_common_proto_msgTypes,
	}.Build()
	File_cloudmind_sts_common_proto = out.File
	file_cloudmind_sts_common_proto_rawDesc = nil
	file_cloudmind_sts_common_proto_goTypes = nil
	file_cloudmind_sts_common_proto_depIdxs = nil
}

var _ context.Context

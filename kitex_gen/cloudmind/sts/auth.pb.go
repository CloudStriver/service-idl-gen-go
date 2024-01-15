// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/sts/auth.proto

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

type SendEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *SendEmailReq) Reset() {
	*x = SendEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailReq) ProtoMessage() {}

func (x *SendEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailReq.ProtoReflect.Descriptor instead.
func (*SendEmailReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendEmailReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type SendEmailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendEmailResp) Reset() {
	*x = SendEmailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResp) ProtoMessage() {}

func (x *SendEmailResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResp.ProtoReflect.Descriptor instead.
func (*SendEmailResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{1}
}

type SetPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Key:
	//
	//	*SetPasswordReq_UserIdOptions
	//	*SetPasswordReq_EmailOptions
	Key      isSetPasswordReq_Key `protobuf_oneof:"Key"`
	Password string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SetPasswordReq) Reset() {
	*x = SetPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordReq) ProtoMessage() {}

func (x *SetPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordReq.ProtoReflect.Descriptor instead.
func (*SetPasswordReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{2}
}

func (m *SetPasswordReq) GetKey() isSetPasswordReq_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *SetPasswordReq) GetUserIdOptions() *UserIdOptions {
	if x, ok := x.GetKey().(*SetPasswordReq_UserIdOptions); ok {
		return x.UserIdOptions
	}
	return nil
}

func (x *SetPasswordReq) GetEmailOptions() *EmailOptions {
	if x, ok := x.GetKey().(*SetPasswordReq_EmailOptions); ok {
		return x.EmailOptions
	}
	return nil
}

func (x *SetPasswordReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type isSetPasswordReq_Key interface {
	isSetPasswordReq_Key()
}

type SetPasswordReq_UserIdOptions struct {
	UserIdOptions *UserIdOptions `protobuf:"bytes,1,opt,name=userIdOptions,proto3,oneof"`
}

type SetPasswordReq_EmailOptions struct {
	EmailOptions *EmailOptions `protobuf:"bytes,2,opt,name=emailOptions,proto3,oneof"`
}

func (*SetPasswordReq_UserIdOptions) isSetPasswordReq_Key() {}

func (*SetPasswordReq_EmailOptions) isSetPasswordReq_Key() {}

type SetPasswordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPasswordResp) Reset() {
	*x = SetPasswordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordResp) ProtoMessage() {}

func (x *SetPasswordResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordResp.ProtoReflect.Descriptor instead.
func (*SetPasswordResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{3}
}

type CheckEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CheckEmailReq) Reset() {
	*x = CheckEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEmailReq) ProtoMessage() {}

func (x *CheckEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEmailReq.ProtoReflect.Descriptor instead.
func (*CheckEmailReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{4}
}

func (x *CheckEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CheckEmailReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CheckEmailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckEmailResp) Reset() {
	*x = CheckEmailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEmailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEmailResp) ProtoMessage() {}

func (x *CheckEmailResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEmailResp.ProtoReflect.Descriptor instead.
func (*CheckEmailResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{5}
}

type CreateAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthInfo *AuthInfo `protobuf:"bytes,1,opt,name=authInfo,proto3" json:"authInfo,omitempty"`
	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	Code     *string   `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`
}

func (x *CreateAuthReq) Reset() {
	*x = CreateAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthReq) ProtoMessage() {}

func (x *CreateAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthReq.ProtoReflect.Descriptor instead.
func (*CreateAuthReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAuthReq) GetAuthInfo() *AuthInfo {
	if x != nil {
		return x.AuthInfo
	}
	return nil
}

func (x *CreateAuthReq) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *CreateAuthReq) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

type CreateAuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateAuthResp) Reset() {
	*x = CreateAuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthResp) ProtoMessage() {}

func (x *CreateAuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthResp.ProtoReflect.Descriptor instead.
func (*CreateAuthResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{7}
}

func (x *CreateAuthResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth     *AuthInfo `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Password string    `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{8}
}

func (x *LoginReq) GetAuth() *AuthInfo {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{9}
}

func (x *LoginResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AppendAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string    `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AuthInfo *AuthInfo `protobuf:"bytes,2,opt,name=authInfo,proto3" json:"authInfo,omitempty"`
	Code     *string   `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`
}

func (x *AppendAuthReq) Reset() {
	*x = AppendAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendAuthReq) ProtoMessage() {}

func (x *AppendAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendAuthReq.ProtoReflect.Descriptor instead.
func (*AppendAuthReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{10}
}

func (x *AppendAuthReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AppendAuthReq) GetAuthInfo() *AuthInfo {
	if x != nil {
		return x.AuthInfo
	}
	return nil
}

func (x *AppendAuthReq) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

type AppendAuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppendAuthResp) Reset() {
	*x = AppendAuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_sts_auth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendAuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendAuthResp) ProtoMessage() {}

func (x *AppendAuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_sts_auth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendAuthResp.ProtoReflect.Descriptor instead.
func (*AppendAuthResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_sts_auth_proto_rawDescGZIP(), []int{11}
}

var File_cloudmind_sts_auth_proto protoreflect.FileDescriptor

var file_cloudmind_sts_auth_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x1a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x44, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00,
	0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x41, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e,
	0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x05,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x39, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x33, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d,
	0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x23, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x33, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73,
	0x74, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x32, 0xbf, 0x03, 0x0a, 0x0a, 0x53, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d,
	0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69,
	0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e,
	0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49,
	0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d,
	0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e,
	0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69,
	0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x49, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1c,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x74, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69,
	0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x73, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudmind_sts_auth_proto_rawDescOnce sync.Once
	file_cloudmind_sts_auth_proto_rawDescData = file_cloudmind_sts_auth_proto_rawDesc
)

func file_cloudmind_sts_auth_proto_rawDescGZIP() []byte {
	file_cloudmind_sts_auth_proto_rawDescOnce.Do(func() {
		file_cloudmind_sts_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudmind_sts_auth_proto_rawDescData)
	})
	return file_cloudmind_sts_auth_proto_rawDescData
}

var file_cloudmind_sts_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cloudmind_sts_auth_proto_goTypes = []interface{}{
	(*SendEmailReq)(nil),    // 0: cloudmind.sts.SendEmailReq
	(*SendEmailResp)(nil),   // 1: cloudmind.sts.SendEmailResp
	(*SetPasswordReq)(nil),  // 2: cloudmind.sts.SetPasswordReq
	(*SetPasswordResp)(nil), // 3: cloudmind.sts.SetPasswordResp
	(*CheckEmailReq)(nil),   // 4: cloudmind.sts.CheckEmailReq
	(*CheckEmailResp)(nil),  // 5: cloudmind.sts.CheckEmailResp
	(*CreateAuthReq)(nil),   // 6: cloudmind.sts.CreateAuthReq
	(*CreateAuthResp)(nil),  // 7: cloudmind.sts.CreateAuthResp
	(*LoginReq)(nil),        // 8: cloudmind.sts.LoginReq
	(*LoginResp)(nil),       // 9: cloudmind.sts.LoginResp
	(*AppendAuthReq)(nil),   // 10: cloudmind.sts.AppendAuthReq
	(*AppendAuthResp)(nil),  // 11: cloudmind.sts.AppendAuthResp
	(*UserIdOptions)(nil),   // 12: cloudmind.sts.UserIdOptions
	(*EmailOptions)(nil),    // 13: cloudmind.sts.EmailOptions
	(*AuthInfo)(nil),        // 14: cloudmind.sts.AuthInfo
	(*UserInfo)(nil),        // 15: cloudmind.sts.UserInfo
}
var file_cloudmind_sts_auth_proto_depIdxs = []int32{
	12, // 0: cloudmind.sts.SetPasswordReq.userIdOptions:type_name -> cloudmind.sts.UserIdOptions
	13, // 1: cloudmind.sts.SetPasswordReq.emailOptions:type_name -> cloudmind.sts.EmailOptions
	14, // 2: cloudmind.sts.CreateAuthReq.authInfo:type_name -> cloudmind.sts.AuthInfo
	15, // 3: cloudmind.sts.CreateAuthReq.userInfo:type_name -> cloudmind.sts.UserInfo
	14, // 4: cloudmind.sts.LoginReq.auth:type_name -> cloudmind.sts.AuthInfo
	14, // 5: cloudmind.sts.AppendAuthReq.authInfo:type_name -> cloudmind.sts.AuthInfo
	2,  // 6: cloudmind.sts.StsService.SetPassword:input_type -> cloudmind.sts.SetPasswordReq
	0,  // 7: cloudmind.sts.StsService.SendEmail:input_type -> cloudmind.sts.SendEmailReq
	4,  // 8: cloudmind.sts.StsService.CheckEmail:input_type -> cloudmind.sts.CheckEmailReq
	6,  // 9: cloudmind.sts.StsService.CreateAuth:input_type -> cloudmind.sts.CreateAuthReq
	8,  // 10: cloudmind.sts.StsService.Login:input_type -> cloudmind.sts.LoginReq
	10, // 11: cloudmind.sts.StsService.AppendAuth:input_type -> cloudmind.sts.AppendAuthReq
	3,  // 12: cloudmind.sts.StsService.SetPassword:output_type -> cloudmind.sts.SetPasswordResp
	1,  // 13: cloudmind.sts.StsService.SendEmail:output_type -> cloudmind.sts.SendEmailResp
	5,  // 14: cloudmind.sts.StsService.CheckEmail:output_type -> cloudmind.sts.CheckEmailResp
	7,  // 15: cloudmind.sts.StsService.CreateAuth:output_type -> cloudmind.sts.CreateAuthResp
	9,  // 16: cloudmind.sts.StsService.Login:output_type -> cloudmind.sts.LoginResp
	11, // 17: cloudmind.sts.StsService.AppendAuth:output_type -> cloudmind.sts.AppendAuthResp
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cloudmind_sts_auth_proto_init() }
func file_cloudmind_sts_auth_proto_init() {
	if File_cloudmind_sts_auth_proto != nil {
		return
	}
	file_cloudmind_sts_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloudmind_sts_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailReq); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResp); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordReq); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordResp); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEmailReq); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEmailResp); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthReq); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthResp); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendAuthReq); i {
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
		file_cloudmind_sts_auth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendAuthResp); i {
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
	file_cloudmind_sts_auth_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SetPasswordReq_UserIdOptions)(nil),
		(*SetPasswordReq_EmailOptions)(nil),
	}
	file_cloudmind_sts_auth_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_cloudmind_sts_auth_proto_msgTypes[10].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_sts_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudmind_sts_auth_proto_goTypes,
		DependencyIndexes: file_cloudmind_sts_auth_proto_depIdxs,
		MessageInfos:      file_cloudmind_sts_auth_proto_msgTypes,
	}.Build()
	File_cloudmind_sts_auth_proto = out.File
	file_cloudmind_sts_auth_proto_rawDesc = nil
	file_cloudmind_sts_auth_proto_goTypes = nil
	file_cloudmind_sts_auth_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.8.0. DO NOT EDIT.

type StsService interface {
	SetPassword(ctx context.Context, req *SetPasswordReq) (res *SetPasswordResp, err error)
	SendEmail(ctx context.Context, req *SendEmailReq) (res *SendEmailResp, err error)
	CheckEmail(ctx context.Context, req *CheckEmailReq) (res *CheckEmailResp, err error)
	CreateAuth(ctx context.Context, req *CreateAuthReq) (res *CreateAuthResp, err error)
	Login(ctx context.Context, req *LoginReq) (res *LoginResp, err error)
	AppendAuth(ctx context.Context, req *AppendAuthReq) (res *AppendAuthResp, err error)
}

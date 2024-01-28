// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/system/common.proto

package system

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

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId  string `protobuf:"bytes,1,opt,name=notificationId,proto3" json:"notificationId,omitempty"`
	TargetUserId    string `protobuf:"bytes,2,opt,name=targetUserId,proto3" json:"targetUserId,omitempty"`
	SourceUserId    string `protobuf:"bytes,3,opt,name=sourceUserId,proto3" json:"sourceUserId,omitempty"`
	SourceContentId string `protobuf:"bytes,4,opt,name=sourceContentId,proto3" json:"sourceContentId,omitempty"`
	TargetType      int64  `protobuf:"varint,5,opt,name=targetType,proto3" json:"targetType,omitempty"`
	Type            int64  `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Text            string `protobuf:"bytes,7,opt,name=text,proto3" json:"text,omitempty"`
	CreateAt        int64  `protobuf:"varint,8,opt,name=createAt,proto3" json:"createAt,omitempty"`
	IsRead          bool   `protobuf:"varint,9,opt,name=isRead,proto3" json:"isRead,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_system_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_system_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_cloudmind_system_common_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *Notification) GetTargetUserId() string {
	if x != nil {
		return x.TargetUserId
	}
	return ""
}

func (x *Notification) GetSourceUserId() string {
	if x != nil {
		return x.SourceUserId
	}
	return ""
}

func (x *Notification) GetSourceContentId() string {
	if x != nil {
		return x.SourceContentId
	}
	return ""
}

func (x *Notification) GetTargetType() int64 {
	if x != nil {
		return x.TargetType
	}
	return 0
}

func (x *Notification) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Notification) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Notification) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Notification) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

type NotificationFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlyUserId     *string `protobuf:"bytes,1,opt,name=onlyUserId,proto3,oneof" json:"onlyUserId,omitempty"`
	OnlyType       *int64  `protobuf:"varint,2,opt,name=onlyType,proto3,oneof" json:"onlyType,omitempty"`
	OnlyTargetType *int64  `protobuf:"varint,3,opt,name=onlyTargetType,proto3,oneof" json:"onlyTargetType,omitempty"`
	OnlyFirstId    *string `protobuf:"bytes,4,opt,name=onlyFirstId,proto3,oneof" json:"onlyFirstId,omitempty"`
	OnlyLastId     *string `protobuf:"bytes,5,opt,name=onlyLastId,proto3,oneof" json:"onlyLastId,omitempty"`
	OnlyIsRead     *bool   `protobuf:"varint,6,opt,name=onlyIsRead,proto3,oneof" json:"onlyIsRead,omitempty"`
}

func (x *NotificationFilterOptions) Reset() {
	*x = NotificationFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_system_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationFilterOptions) ProtoMessage() {}

func (x *NotificationFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_system_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationFilterOptions.ProtoReflect.Descriptor instead.
func (*NotificationFilterOptions) Descriptor() ([]byte, []int) {
	return file_cloudmind_system_common_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationFilterOptions) GetOnlyUserId() string {
	if x != nil && x.OnlyUserId != nil {
		return *x.OnlyUserId
	}
	return ""
}

func (x *NotificationFilterOptions) GetOnlyType() int64 {
	if x != nil && x.OnlyType != nil {
		return *x.OnlyType
	}
	return 0
}

func (x *NotificationFilterOptions) GetOnlyTargetType() int64 {
	if x != nil && x.OnlyTargetType != nil {
		return *x.OnlyTargetType
	}
	return 0
}

func (x *NotificationFilterOptions) GetOnlyFirstId() string {
	if x != nil && x.OnlyFirstId != nil {
		return *x.OnlyFirstId
	}
	return ""
}

func (x *NotificationFilterOptions) GetOnlyLastId() string {
	if x != nil && x.OnlyLastId != nil {
		return *x.OnlyLastId
	}
	return ""
}

func (x *NotificationFilterOptions) GetOnlyIsRead() bool {
	if x != nil && x.OnlyIsRead != nil {
		return *x.OnlyIsRead
	}
	return false
}

type Slider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SliderId   string `protobuf:"bytes,1,opt,name=sliderId,proto3" json:"sliderId,omitempty"`
	ImageUrl   string `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	LinkUrl    string `protobuf:"bytes,3,opt,name=linkUrl,proto3" json:"linkUrl,omitempty"`
	Type       int64  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	IsPublic   int64  `protobuf:"varint,5,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	CreateTime int64  `protobuf:"varint,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,7,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *Slider) Reset() {
	*x = Slider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_system_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slider) ProtoMessage() {}

func (x *Slider) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_system_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slider.ProtoReflect.Descriptor instead.
func (*Slider) Descriptor() ([]byte, []int) {
	return file_cloudmind_system_common_proto_rawDescGZIP(), []int{2}
}

func (x *Slider) GetSliderId() string {
	if x != nil {
		return x.SliderId
	}
	return ""
}

func (x *Slider) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Slider) GetLinkUrl() string {
	if x != nil {
		return x.LinkUrl
	}
	return ""
}

func (x *Slider) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Slider) GetIsPublic() int64 {
	if x != nil {
		return x.IsPublic
	}
	return 0
}

func (x *Slider) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Slider) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type SliderFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlyType     *int64 `protobuf:"varint,1,opt,name=onlyType,proto3,oneof" json:"onlyType,omitempty"`
	OnlyIsPublic *int64 `protobuf:"varint,2,opt,name=onlyIsPublic,proto3,oneof" json:"onlyIsPublic,omitempty"`
}

func (x *SliderFilterOptions) Reset() {
	*x = SliderFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_system_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliderFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliderFilterOptions) ProtoMessage() {}

func (x *SliderFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_system_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliderFilterOptions.ProtoReflect.Descriptor instead.
func (*SliderFilterOptions) Descriptor() ([]byte, []int) {
	return file_cloudmind_system_common_proto_rawDescGZIP(), []int{3}
}

func (x *SliderFilterOptions) GetOnlyType() int64 {
	if x != nil && x.OnlyType != nil {
		return *x.OnlyType
	}
	return 0
}

func (x *SliderFilterOptions) GetOnlyIsPublic() int64 {
	if x != nil && x.OnlyIsPublic != nil {
		return *x.OnlyIsPublic
	}
	return 0
}

var File_cloudmind_system_common_proto protoreflect.FileDescriptor

var file_cloudmind_system_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x22, 0xa4, 0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x22, 0xdc, 0x02, 0x0a, 0x19, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e,
	0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6f,
	0x6e, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52,
	0x08, 0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e,
	0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x0e, 0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c,
	0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x4c, 0x61, 0x73, 0x74,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x0a, 0x6f, 0x6e, 0x6c,
	0x79, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x06, 0x53, 0x6c, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69,
	0x6e, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x6e,
	0x6b, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x13, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x6f,
	0x6e, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x08, 0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c,
	0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x01, 0x52, 0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x73, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x73, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudmind_system_common_proto_rawDescOnce sync.Once
	file_cloudmind_system_common_proto_rawDescData = file_cloudmind_system_common_proto_rawDesc
)

func file_cloudmind_system_common_proto_rawDescGZIP() []byte {
	file_cloudmind_system_common_proto_rawDescOnce.Do(func() {
		file_cloudmind_system_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudmind_system_common_proto_rawDescData)
	})
	return file_cloudmind_system_common_proto_rawDescData
}

var file_cloudmind_system_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloudmind_system_common_proto_goTypes = []interface{}{
	(*Notification)(nil),              // 0: cloudmind.system.Notification
	(*NotificationFilterOptions)(nil), // 1: cloudmind.system.NotificationFilterOptions
	(*Slider)(nil),                    // 2: cloudmind.system.Slider
	(*SliderFilterOptions)(nil),       // 3: cloudmind.system.SliderFilterOptions
}
var file_cloudmind_system_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloudmind_system_common_proto_init() }
func file_cloudmind_system_common_proto_init() {
	if File_cloudmind_system_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudmind_system_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_cloudmind_system_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationFilterOptions); i {
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
		file_cloudmind_system_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slider); i {
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
		file_cloudmind_system_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliderFilterOptions); i {
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
	file_cloudmind_system_common_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_cloudmind_system_common_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_system_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudmind_system_common_proto_goTypes,
		DependencyIndexes: file_cloudmind_system_common_proto_depIdxs,
		MessageInfos:      file_cloudmind_system_common_proto_msgTypes,
	}.Build()
	File_cloudmind_system_common_proto = out.File
	file_cloudmind_system_common_proto_rawDesc = nil
	file_cloudmind_system_common_proto_goTypes = nil
	file_cloudmind_system_common_proto_depIdxs = nil
}

var _ context.Context

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/content/common.proto

package content

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

type Type int32

const (
	Type_Type_null    Type = 0
	Type_Type_audio   Type = 1
	Type_Type_video   Type = 2
	Type_Type_photo   Type = 3
	Type_Type_word    Type = 4
	Type_Type_ppt     Type = 5
	Type_Type_zip     Type = 6
	Type_Type_excel   Type = 7
	Type_Type_pdf     Type = 8
	Type_Type_folder  Type = 9
	Type_Type_file    Type = 10
	Type_Type_code    Type = 11
	Type_Type_unknown Type = 12
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:  "Type_null",
		1:  "Type_audio",
		2:  "Type_video",
		3:  "Type_photo",
		4:  "Type_word",
		5:  "Type_ppt",
		6:  "Type_zip",
		7:  "Type_excel",
		8:  "Type_pdf",
		9:  "Type_folder",
		10: "Type_file",
		11: "Type_code",
		12: "Type_unknown",
	}
	Type_value = map[string]int32{
		"Type_null":    0,
		"Type_audio":   1,
		"Type_video":   2,
		"Type_photo":   3,
		"Type_word":    4,
		"Type_ppt":     5,
		"Type_zip":     6,
		"Type_excel":   7,
		"Type_pdf":     8,
		"Type_folder":  9,
		"Type_file":    10,
		"Type_code":    11,
		"Type_unknown": 12,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudmind_content_common_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_cloudmind_content_common_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{0}
}

type IsDel int32

const (
	IsDel_Is_null IsDel = 0
	IsDel_Is_no   IsDel = 1
	IsDel_Is_soft IsDel = 2
	IsDel_Is_hard IsDel = 3
)

// Enum value maps for IsDel.
var (
	IsDel_name = map[int32]string{
		0: "Is_null",
		1: "Is_no",
		2: "Is_soft",
		3: "Is_hard",
	}
	IsDel_value = map[string]int32{
		"Is_null": 0,
		"Is_no":   1,
		"Is_soft": 2,
		"Is_hard": 3,
	}
)

func (x IsDel) Enum() *IsDel {
	p := new(IsDel)
	*p = x
	return p
}

func (x IsDel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IsDel) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudmind_content_common_proto_enumTypes[1].Descriptor()
}

func (IsDel) Type() protoreflect.EnumType {
	return &file_cloudmind_content_common_proto_enumTypes[1]
}

func (x IsDel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IsDel.Descriptor instead.
func (IsDel) EnumDescriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{1}
}

type DocumentType int32

const (
	DocumentType_DocumentType_null     DocumentType = 0
	DocumentType_DocumentType_personal DocumentType = 1
	DocumentType_DocumentType_public   DocumentType = 2
)

// Enum value maps for DocumentType.
var (
	DocumentType_name = map[int32]string{
		0: "DocumentType_null",
		1: "DocumentType_personal",
		2: "DocumentType_public",
	}
	DocumentType_value = map[string]int32{
		"DocumentType_null":     0,
		"DocumentType_personal": 1,
		"DocumentType_public":   2,
	}
)

func (x DocumentType) Enum() *DocumentType {
	p := new(DocumentType)
	*p = x
	return p
}

func (x DocumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudmind_content_common_proto_enumTypes[2].Descriptor()
}

func (DocumentType) Type() protoreflect.EnumType {
	return &file_cloudmind_content_common_proto_enumTypes[2]
}

func (x DocumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocumentType.Descriptor instead.
func (DocumentType) EnumDescriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{2}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId      string   `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type        Type     `protobuf:"varint,4,opt,name=type,proto3,enum=cloudmind.content.Type" json:"type,omitempty"`
	Path        string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	FatherId    string   `protobuf:"bytes,6,opt,name=fatherId,proto3" json:"fatherId,omitempty"`
	SpaceSize   *int64   `protobuf:"varint,7,opt,name=spaceSize,proto3,oneof" json:"spaceSize,omitempty"`
	Md5         string   `protobuf:"bytes,8,opt,name=md5,proto3" json:"md5,omitempty"`
	IsDel       int32    `protobuf:"varint,9,opt,name=isDel,proto3" json:"isDel,omitempty"`
	Tag         []string `protobuf:"bytes,10,rep,name=tag,proto3" json:"tag,omitempty"`
	Description string   `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	UpdateAt    int64    `protobuf:"varint,12,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_content_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_content_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *File) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_Type_null
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetFatherId() string {
	if x != nil {
		return x.FatherId
	}
	return ""
}

func (x *File) GetSpaceSize() int64 {
	if x != nil && x.SpaceSize != nil {
		return *x.SpaceSize
	}
	return 0
}

func (x *File) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *File) GetIsDel() int32 {
	if x != nil {
		return x.IsDel
	}
	return 0
}

func (x *File) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *File) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *File) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId      string   `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type        Type     `protobuf:"varint,4,opt,name=type,proto3,enum=cloudmind.content.Type" json:"type,omitempty"`
	Path        string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	FatherId    string   `protobuf:"bytes,6,opt,name=fatherId,proto3" json:"fatherId,omitempty"`
	SpaceSize   int64    `protobuf:"varint,7,opt,name=spaceSize,proto3" json:"spaceSize,omitempty"`
	Md5         string   `protobuf:"bytes,8,opt,name=md5,proto3" json:"md5,omitempty"`
	IsDel       int32    `protobuf:"varint,9,opt,name=isDel,proto3" json:"isDel,omitempty"`
	Tag         []string `protobuf:"bytes,10,rep,name=tag,proto3" json:"tag,omitempty"`
	Description string   `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	UpdateAt    int64    `protobuf:"varint,12,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_content_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_content_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_Type_null
}

func (x *FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileInfo) GetFatherId() string {
	if x != nil {
		return x.FatherId
	}
	return ""
}

func (x *FileInfo) GetSpaceSize() int64 {
	if x != nil {
		return x.SpaceSize
	}
	return 0
}

func (x *FileInfo) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *FileInfo) GetIsDel() int32 {
	if x != nil {
		return x.IsDel
	}
	return 0
}

func (x *FileInfo) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *FileInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FileInfo) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_content_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_content_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{2}
}

func (x *Label) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SearchField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  *string `protobuf:"bytes,1,opt,name=text,proto3,oneof" json:"text,omitempty"`
	Title *string `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Tag   *string `protobuf:"bytes,3,opt,name=tag,proto3,oneof" json:"tag,omitempty"` // 仅限Post
}

func (x *SearchField) Reset() {
	*x = SearchField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_content_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchField) ProtoMessage() {}

func (x *SearchField) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_content_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchField.ProtoReflect.Descriptor instead.
func (*SearchField) Descriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{3}
}

func (x *SearchField) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *SearchField) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *SearchField) GetTag() string {
	if x != nil && x.Tag != nil {
		return *x.Tag
	}
	return ""
}

type SearchOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*SearchOptions_AllFieldsKey
	//	*SearchOptions_MultiFieldsKey
	Type isSearchOptions_Type `protobuf_oneof:"Type"`
}

func (x *SearchOptions) Reset() {
	*x = SearchOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_content_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOptions) ProtoMessage() {}

func (x *SearchOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_content_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOptions.ProtoReflect.Descriptor instead.
func (*SearchOptions) Descriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{4}
}

func (m *SearchOptions) GetType() isSearchOptions_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *SearchOptions) GetAllFieldsKey() string {
	if x, ok := x.GetType().(*SearchOptions_AllFieldsKey); ok {
		return x.AllFieldsKey
	}
	return ""
}

func (x *SearchOptions) GetMultiFieldsKey() *SearchField {
	if x, ok := x.GetType().(*SearchOptions_MultiFieldsKey); ok {
		return x.MultiFieldsKey
	}
	return nil
}

type isSearchOptions_Type interface {
	isSearchOptions_Type()
}

type SearchOptions_AllFieldsKey struct {
	AllFieldsKey string `protobuf:"bytes,1,opt,name=allFieldsKey,proto3,oneof"`
}

type SearchOptions_MultiFieldsKey struct {
	MultiFieldsKey *SearchField `protobuf:"bytes,2,opt,name=multiFieldsKey,proto3,oneof"`
}

func (*SearchOptions_AllFieldsKey) isSearchOptions_Type() {}

func (*SearchOptions_MultiFieldsKey) isSearchOptions_Type() {}

type FileFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlyUserId   *string `protobuf:"bytes,1,opt,name=onlyUserId,proto3,oneof" json:"onlyUserId,omitempty"`
	OnlyFileId   *string `protobuf:"bytes,2,opt,name=onlyFileId,proto3,oneof" json:"onlyFileId,omitempty"`
	OnlyFatherId *string `protobuf:"bytes,3,opt,name=onlyFatherId,proto3,oneof" json:"onlyFatherId,omitempty"`
	OnlyFileType *int32  `protobuf:"varint,4,opt,name=onlyFileType,proto3,oneof" json:"onlyFileType,omitempty"` // 文件类型
	IsDel        int32   `protobuf:"varint,5,opt,name=isDel,proto3" json:"isDel,omitempty"`                     // 是否删除 1：未删除 2：软删除 3:彻底删除
	DocumentType int32   `protobuf:"varint,6,opt,name=documentType,proto3" json:"documentType,omitempty"`       // 操作类型 1：个人空间文件 2：公共空间文件
}

func (x *FileFilterOptions) Reset() {
	*x = FileFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_content_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileFilterOptions) ProtoMessage() {}

func (x *FileFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_content_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileFilterOptions.ProtoReflect.Descriptor instead.
func (*FileFilterOptions) Descriptor() ([]byte, []int) {
	return file_cloudmind_content_common_proto_rawDescGZIP(), []int{5}
}

func (x *FileFilterOptions) GetOnlyUserId() string {
	if x != nil && x.OnlyUserId != nil {
		return *x.OnlyUserId
	}
	return ""
}

func (x *FileFilterOptions) GetOnlyFileId() string {
	if x != nil && x.OnlyFileId != nil {
		return *x.OnlyFileId
	}
	return ""
}

func (x *FileFilterOptions) GetOnlyFatherId() string {
	if x != nil && x.OnlyFatherId != nil {
		return *x.OnlyFatherId
	}
	return ""
}

func (x *FileFilterOptions) GetOnlyFileType() int32 {
	if x != nil && x.OnlyFileType != nil {
		return *x.OnlyFileType
	}
	return 0
}

func (x *FileFilterOptions) GetIsDel() int32 {
	if x != nil {
		return x.IsDel
	}
	return 0
}

func (x *FileFilterOptions) GetDocumentType() int32 {
	if x != nil {
		return x.DocumentType
	}
	return 0
}

var File_cloudmind_content_common_proto protoreflect.FileDescriptor

var file_cloudmind_content_common_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0xd0, 0x02, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x64, 0x35, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xc1, 0x02, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e,
	0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x64, 0x35, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x2d, 0x0a, 0x05, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x73, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03,
	0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x61, 0x67, 0x22, 0x87,
	0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x24, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00,
	0x52, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4b, 0x65, 0x79,
	0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x11, 0x46, 0x69, 0x6c,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23,
	0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x79,
	0x46, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x27, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73,
	0x44, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x12, 0x22, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x49, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x2a, 0xcf, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x70, 0x70, 0x74, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x7a, 0x69, 0x70, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x65, 0x78, 0x63, 0x65, 0x6c, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x70, 0x64, 0x66, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x0c, 0x2a, 0x39, 0x0a, 0x05, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x73, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x49, 0x73, 0x5f, 0x6e, 0x6f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x73, 0x5f, 0x73, 0x6f,
	0x66, 0x74, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x73, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x10,
	0x03, 0x2a, 0x59, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x02, 0x42, 0x48, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x53, 0x74, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudmind_content_common_proto_rawDescOnce sync.Once
	file_cloudmind_content_common_proto_rawDescData = file_cloudmind_content_common_proto_rawDesc
)

func file_cloudmind_content_common_proto_rawDescGZIP() []byte {
	file_cloudmind_content_common_proto_rawDescOnce.Do(func() {
		file_cloudmind_content_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudmind_content_common_proto_rawDescData)
	})
	return file_cloudmind_content_common_proto_rawDescData
}

var file_cloudmind_content_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cloudmind_content_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloudmind_content_common_proto_goTypes = []interface{}{
	(Type)(0),                 // 0: cloudmind.content.Type
	(IsDel)(0),                // 1: cloudmind.content.IsDel
	(DocumentType)(0),         // 2: cloudmind.content.DocumentType
	(*File)(nil),              // 3: cloudmind.content.File
	(*FileInfo)(nil),          // 4: cloudmind.content.FileInfo
	(*Label)(nil),             // 5: cloudmind.content.Label
	(*SearchField)(nil),       // 6: cloudmind.content.SearchField
	(*SearchOptions)(nil),     // 7: cloudmind.content.SearchOptions
	(*FileFilterOptions)(nil), // 8: cloudmind.content.FileFilterOptions
}
var file_cloudmind_content_common_proto_depIdxs = []int32{
	0, // 0: cloudmind.content.File.type:type_name -> cloudmind.content.Type
	0, // 1: cloudmind.content.FileInfo.type:type_name -> cloudmind.content.Type
	6, // 2: cloudmind.content.SearchOptions.multiFieldsKey:type_name -> cloudmind.content.SearchField
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloudmind_content_common_proto_init() }
func file_cloudmind_content_common_proto_init() {
	if File_cloudmind_content_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudmind_content_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_cloudmind_content_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_cloudmind_content_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_cloudmind_content_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchField); i {
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
		file_cloudmind_content_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOptions); i {
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
		file_cloudmind_content_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileFilterOptions); i {
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
	file_cloudmind_content_common_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_cloudmind_content_common_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_cloudmind_content_common_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*SearchOptions_AllFieldsKey)(nil),
		(*SearchOptions_MultiFieldsKey)(nil),
	}
	file_cloudmind_content_common_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_content_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudmind_content_common_proto_goTypes,
		DependencyIndexes: file_cloudmind_content_common_proto_depIdxs,
		EnumInfos:         file_cloudmind_content_common_proto_enumTypes,
		MessageInfos:      file_cloudmind_content_common_proto_msgTypes,
	}.Build()
	File_cloudmind_content_common_proto = out.File
	file_cloudmind_content_common_proto_rawDesc = nil
	file_cloudmind_content_common_proto_goTypes = nil
	file_cloudmind_content_common_proto_depIdxs = nil
}

var _ context.Context

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/trade/common.proto

package trade

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

type BalanceType int32

const (
	BalanceType_UnKnowBalanceType BalanceType = 0
	BalanceType_FlowBalanceType   BalanceType = 1
	BalanceType_MemoryBalanceType BalanceType = 2
	BalanceType_PointBalanceType  BalanceType = 3
)

// Enum value maps for BalanceType.
var (
	BalanceType_name = map[int32]string{
		0: "UnKnowBalanceType",
		1: "FlowBalanceType",
		2: "MemoryBalanceType",
		3: "PointBalanceType",
	}
	BalanceType_value = map[string]int32{
		"UnKnowBalanceType": 0,
		"FlowBalanceType":   1,
		"MemoryBalanceType": 2,
		"PointBalanceType":  3,
	}
)

func (x BalanceType) Enum() *BalanceType {
	p := new(BalanceType)
	*p = x
	return p
}

func (x BalanceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BalanceType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudmind_trade_common_proto_enumTypes[0].Descriptor()
}

func (BalanceType) Type() protoreflect.EnumType {
	return &file_cloudmind_trade_common_proto_enumTypes[0]
}

func (x BalanceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BalanceType.Descriptor instead.
func (BalanceType) EnumDescriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{0}
}

type UpdateBalanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Flow   *int64 `protobuf:"varint,2,opt,name=flow,proto3,oneof" json:"flow,omitempty"`
	Memory *int64 `protobuf:"varint,3,opt,name=memory,proto3,oneof" json:"memory,omitempty"`
	Point  *int64 `protobuf:"varint,4,opt,name=point,proto3,oneof" json:"point,omitempty"`
}

func (x *UpdateBalanceReq) Reset() {
	*x = UpdateBalanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBalanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBalanceReq) ProtoMessage() {}

func (x *UpdateBalanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBalanceReq.ProtoReflect.Descriptor instead.
func (*UpdateBalanceReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBalanceReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateBalanceReq) GetFlow() int64 {
	if x != nil && x.Flow != nil {
		return *x.Flow
	}
	return 0
}

func (x *UpdateBalanceReq) GetMemory() int64 {
	if x != nil && x.Memory != nil {
		return *x.Memory
	}
	return 0
}

func (x *UpdateBalanceReq) GetPoint() int64 {
	if x != nil && x.Point != nil {
		return *x.Point
	}
	return 0
}

type UpdateBalanceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UpdateBalanceResp) Reset() {
	*x = UpdateBalanceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBalanceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBalanceResp) ProtoMessage() {}

func (x *UpdateBalanceResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBalanceResp.ProtoReflect.Descriptor instead.
func (*UpdateBalanceResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBalanceResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type GetBalanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetBalanceReq) Reset() {
	*x = GetBalanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceReq) ProtoMessage() {}

func (x *GetBalanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceReq.ProtoReflect.Descriptor instead.
func (*GetBalanceReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetBalanceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flow   int64 `protobuf:"varint,1,opt,name=flow,proto3" json:"flow,omitempty"`
	Memory int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Point  int64 `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *GetBalanceResp) Reset() {
	*x = GetBalanceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResp) ProtoMessage() {}

func (x *GetBalanceResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResp.ProtoReflect.Descriptor instead.
func (*GetBalanceResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{3}
}

func (x *GetBalanceResp) GetFlow() int64 {
	if x != nil {
		return x.Flow
	}
	return 0
}

func (x *GetBalanceResp) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *GetBalanceResp) GetPoint() int64 {
	if x != nil {
		return x.Point
	}
	return 0
}

type CreateBalanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateBalanceReq) Reset() {
	*x = CreateBalanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBalanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBalanceReq) ProtoMessage() {}

func (x *CreateBalanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBalanceReq.ProtoReflect.Descriptor instead.
func (*CreateBalanceReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBalanceReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateBalanceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBalanceResp) Reset() {
	*x = CreateBalanceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBalanceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBalanceResp) ProtoMessage() {}

func (x *CreateBalanceResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBalanceResp.ProtoReflect.Descriptor instead.
func (*CreateBalanceResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{5}
}

type AddStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Amount    int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AddStockReq) Reset() {
	*x = AddStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStockReq) ProtoMessage() {}

func (x *AddStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStockReq.ProtoReflect.Descriptor instead.
func (*AddStockReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{6}
}

func (x *AddStockReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddStockReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddStockResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddStockResp) Reset() {
	*x = AddStockResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStockResp) ProtoMessage() {}

func (x *AddStockResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStockResp.ProtoReflect.Descriptor instead.
func (*AddStockResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{7}
}

type GetStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *GetStockReq) Reset() {
	*x = GetStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockReq) ProtoMessage() {}

func (x *GetStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockReq.ProtoReflect.Descriptor instead.
func (*GetStockReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{8}
}

func (x *GetStockReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type GetStockResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock int64 `protobuf:"varint,1,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *GetStockResp) Reset() {
	*x = GetStockResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockResp) ProtoMessage() {}

func (x *GetStockResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockResp.ProtoReflect.Descriptor instead.
func (*GetStockResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{9}
}

func (x *GetStockResp) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GetStocksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductIds []string `protobuf:"bytes,1,rep,name=productIds,proto3" json:"productIds,omitempty"`
}

func (x *GetStocksReq) Reset() {
	*x = GetStocksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStocksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksReq) ProtoMessage() {}

func (x *GetStocksReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksReq.ProtoReflect.Descriptor instead.
func (*GetStocksReq) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{10}
}

func (x *GetStocksReq) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type GetStocksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []int64 `protobuf:"varint,1,rep,packed,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *GetStocksResp) Reset() {
	*x = GetStocksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudmind_trade_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStocksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksResp) ProtoMessage() {}

func (x *GetStocksResp) ProtoReflect() protoreflect.Message {
	mi := &file_cloudmind_trade_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksResp.ProtoReflect.Descriptor instead.
func (*GetStocksResp) Descriptor() ([]byte, []int) {
	return file_cloudmind_trade_common_proto_rawDescGZIP(), []int{11}
}

func (x *GetStocksResp) GetStocks() []int64 {
	if x != nil {
		return x.Stocks
	}
	return nil
}

var File_cloudmind_trade_common_proto protoreflect.FileDescriptor

var file_cloudmind_trade_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x22,
	0x99, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04,
	0x66, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04, 0x66, 0x6c,
	0x6f, 0x77, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x02, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x22, 0x27, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x2a, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x43,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x24, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x27, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x2a,
	0x66, 0x0a, 0x0b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x6c, 0x6f, 0x77, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x03, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudmind_trade_common_proto_rawDescOnce sync.Once
	file_cloudmind_trade_common_proto_rawDescData = file_cloudmind_trade_common_proto_rawDesc
)

func file_cloudmind_trade_common_proto_rawDescGZIP() []byte {
	file_cloudmind_trade_common_proto_rawDescOnce.Do(func() {
		file_cloudmind_trade_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudmind_trade_common_proto_rawDescData)
	})
	return file_cloudmind_trade_common_proto_rawDescData
}

var file_cloudmind_trade_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloudmind_trade_common_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cloudmind_trade_common_proto_goTypes = []interface{}{
	(BalanceType)(0),          // 0: cloudmind.trade.BalanceType
	(*UpdateBalanceReq)(nil),  // 1: cloudmind.trade.UpdateBalanceReq
	(*UpdateBalanceResp)(nil), // 2: cloudmind.trade.UpdateBalanceResp
	(*GetBalanceReq)(nil),     // 3: cloudmind.trade.GetBalanceReq
	(*GetBalanceResp)(nil),    // 4: cloudmind.trade.GetBalanceResp
	(*CreateBalanceReq)(nil),  // 5: cloudmind.trade.CreateBalanceReq
	(*CreateBalanceResp)(nil), // 6: cloudmind.trade.CreateBalanceResp
	(*AddStockReq)(nil),       // 7: cloudmind.trade.AddStockReq
	(*AddStockResp)(nil),      // 8: cloudmind.trade.AddStockResp
	(*GetStockReq)(nil),       // 9: cloudmind.trade.GetStockReq
	(*GetStockResp)(nil),      // 10: cloudmind.trade.GetStockResp
	(*GetStocksReq)(nil),      // 11: cloudmind.trade.GetStocksReq
	(*GetStocksResp)(nil),     // 12: cloudmind.trade.GetStocksResp
}
var file_cloudmind_trade_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloudmind_trade_common_proto_init() }
func file_cloudmind_trade_common_proto_init() {
	if File_cloudmind_trade_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudmind_trade_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBalanceReq); i {
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
		file_cloudmind_trade_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBalanceResp); i {
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
		file_cloudmind_trade_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceReq); i {
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
		file_cloudmind_trade_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResp); i {
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
		file_cloudmind_trade_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBalanceReq); i {
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
		file_cloudmind_trade_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBalanceResp); i {
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
		file_cloudmind_trade_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStockReq); i {
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
		file_cloudmind_trade_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStockResp); i {
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
		file_cloudmind_trade_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockReq); i {
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
		file_cloudmind_trade_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockResp); i {
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
		file_cloudmind_trade_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStocksReq); i {
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
		file_cloudmind_trade_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStocksResp); i {
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
	file_cloudmind_trade_common_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_trade_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudmind_trade_common_proto_goTypes,
		DependencyIndexes: file_cloudmind_trade_common_proto_depIdxs,
		EnumInfos:         file_cloudmind_trade_common_proto_enumTypes,
		MessageInfos:      file_cloudmind_trade_common_proto_msgTypes,
	}.Build()
	File_cloudmind_trade_common_proto = out.File
	file_cloudmind_trade_common_proto_rawDesc = nil
	file_cloudmind_trade_common_proto_goTypes = nil
	file_cloudmind_trade_common_proto_depIdxs = nil
}

var _ context.Context

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cloudmind/trade/trade.proto

package trade

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cloudmind_trade_trade_proto protoreflect.FileDescriptor

var file_cloudmind_trade_trade_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x1a, 0x1c,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xeb, 0x03, 0x0a,
	0x0c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e,
	0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x47, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69,
	0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4a,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64,
	0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloudmind_trade_trade_proto_goTypes = []interface{}{
	(*UpdateBalanceReq)(nil),  // 0: cloudmind.trade.UpdateBalanceReq
	(*GetBalanceReq)(nil),     // 1: cloudmind.trade.GetBalanceReq
	(*CreateBalanceReq)(nil),  // 2: cloudmind.trade.CreateBalanceReq
	(*AddStockReq)(nil),       // 3: cloudmind.trade.AddStockReq
	(*GetStockReq)(nil),       // 4: cloudmind.trade.GetStockReq
	(*GetStocksReq)(nil),      // 5: cloudmind.trade.GetStocksReq
	(*UpdateBalanceResp)(nil), // 6: cloudmind.trade.UpdateBalanceResp
	(*GetBalanceResp)(nil),    // 7: cloudmind.trade.GetBalanceResp
	(*CreateBalanceResp)(nil), // 8: cloudmind.trade.CreateBalanceResp
	(*AddStockResp)(nil),      // 9: cloudmind.trade.AddStockResp
	(*GetStockResp)(nil),      // 10: cloudmind.trade.GetStockResp
	(*GetStocksResp)(nil),     // 11: cloudmind.trade.GetStocksResp
}
var file_cloudmind_trade_trade_proto_depIdxs = []int32{
	0,  // 0: cloudmind.trade.TradeService.UpdateBalance:input_type -> cloudmind.trade.UpdateBalanceReq
	1,  // 1: cloudmind.trade.TradeService.GetBalance:input_type -> cloudmind.trade.GetBalanceReq
	2,  // 2: cloudmind.trade.TradeService.CreateBalance:input_type -> cloudmind.trade.CreateBalanceReq
	3,  // 3: cloudmind.trade.TradeService.AddStock:input_type -> cloudmind.trade.AddStockReq
	4,  // 4: cloudmind.trade.TradeService.GetStock:input_type -> cloudmind.trade.GetStockReq
	5,  // 5: cloudmind.trade.TradeService.GetStocks:input_type -> cloudmind.trade.GetStocksReq
	6,  // 6: cloudmind.trade.TradeService.UpdateBalance:output_type -> cloudmind.trade.UpdateBalanceResp
	7,  // 7: cloudmind.trade.TradeService.GetBalance:output_type -> cloudmind.trade.GetBalanceResp
	8,  // 8: cloudmind.trade.TradeService.CreateBalance:output_type -> cloudmind.trade.CreateBalanceResp
	9,  // 9: cloudmind.trade.TradeService.AddStock:output_type -> cloudmind.trade.AddStockResp
	10, // 10: cloudmind.trade.TradeService.GetStock:output_type -> cloudmind.trade.GetStockResp
	11, // 11: cloudmind.trade.TradeService.GetStocks:output_type -> cloudmind.trade.GetStocksResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloudmind_trade_trade_proto_init() }
func file_cloudmind_trade_trade_proto_init() {
	if File_cloudmind_trade_trade_proto != nil {
		return
	}
	file_cloudmind_trade_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudmind_trade_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudmind_trade_trade_proto_goTypes,
		DependencyIndexes: file_cloudmind_trade_trade_proto_depIdxs,
	}.Build()
	File_cloudmind_trade_trade_proto = out.File
	file_cloudmind_trade_trade_proto_rawDesc = nil
	file_cloudmind_trade_trade_proto_goTypes = nil
	file_cloudmind_trade_trade_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.0. DO NOT EDIT.

type TradeService interface {
	UpdateBalance(ctx context.Context, req *UpdateBalanceReq) (res *UpdateBalanceResp, err error)
	GetBalance(ctx context.Context, req *GetBalanceReq) (res *GetBalanceResp, err error)
	CreateBalance(ctx context.Context, req *CreateBalanceReq) (res *CreateBalanceResp, err error)
	AddStock(ctx context.Context, req *AddStockReq) (res *AddStockResp, err error)
	GetStock(ctx context.Context, req *GetStockReq) (res *GetStockResp, err error)
	GetStocks(ctx context.Context, req *GetStocksReq) (res *GetStocksResp, err error)
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0--rc1
// source: albelt/stock_srv/const.proto

package proto

import (
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

// 库存扣减历史的状态
type StockHistoryStatus int32

const (
	StockHistoryStatus_STOCK_HISTORY_STATUS_UNSPECIFIED StockHistoryStatus = 0
	// 已扣减
	StockHistoryStatus_STOCK_HISTORY_STATUS_REDUCED StockHistoryStatus = 1
	// 已归还
	StockHistoryStatus_STOCK_HISTORY_STATUS_RESTORED StockHistoryStatus = 2
)

// Enum value maps for StockHistoryStatus.
var (
	StockHistoryStatus_name = map[int32]string{
		0: "STOCK_HISTORY_STATUS_UNSPECIFIED",
		1: "STOCK_HISTORY_STATUS_REDUCED",
		2: "STOCK_HISTORY_STATUS_RESTORED",
	}
	StockHistoryStatus_value = map[string]int32{
		"STOCK_HISTORY_STATUS_UNSPECIFIED": 0,
		"STOCK_HISTORY_STATUS_REDUCED":     1,
		"STOCK_HISTORY_STATUS_RESTORED":    2,
	}
)

func (x StockHistoryStatus) Enum() *StockHistoryStatus {
	p := new(StockHistoryStatus)
	*p = x
	return p
}

func (x StockHistoryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StockHistoryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_albelt_stock_srv_const_proto_enumTypes[0].Descriptor()
}

func (StockHistoryStatus) Type() protoreflect.EnumType {
	return &file_albelt_stock_srv_const_proto_enumTypes[0]
}

func (x StockHistoryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StockHistoryStatus.Descriptor instead.
func (StockHistoryStatus) EnumDescriptor() ([]byte, []int) {
	return file_albelt_stock_srv_const_proto_rawDescGZIP(), []int{0}
}

var File_albelt_stock_srv_const_proto protoreflect.FileDescriptor

var file_albelt_stock_srv_const_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x6c, 0x62, 0x65, 0x6c, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x73,
	0x72, 0x76, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x72, 0x76, 0x2a, 0x7f, 0x0a, 0x12, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x24, 0x0a, 0x20, 0x53, 0x54, 0x4f, 0x43, 0x4b, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x4f, 0x43, 0x4b, 0x5f, 0x48,
	0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45,
	0x44, 0x55, 0x43, 0x45, 0x44, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x54, 0x4f, 0x43, 0x4b,
	0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x44, 0x10, 0x02, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x6f,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_albelt_stock_srv_const_proto_rawDescOnce sync.Once
	file_albelt_stock_srv_const_proto_rawDescData = file_albelt_stock_srv_const_proto_rawDesc
)

func file_albelt_stock_srv_const_proto_rawDescGZIP() []byte {
	file_albelt_stock_srv_const_proto_rawDescOnce.Do(func() {
		file_albelt_stock_srv_const_proto_rawDescData = protoimpl.X.CompressGZIP(file_albelt_stock_srv_const_proto_rawDescData)
	})
	return file_albelt_stock_srv_const_proto_rawDescData
}

var file_albelt_stock_srv_const_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_albelt_stock_srv_const_proto_goTypes = []interface{}{
	(StockHistoryStatus)(0), // 0: stock_srv.StockHistoryStatus
}
var file_albelt_stock_srv_const_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_albelt_stock_srv_const_proto_init() }
func file_albelt_stock_srv_const_proto_init() {
	if File_albelt_stock_srv_const_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_albelt_stock_srv_const_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_albelt_stock_srv_const_proto_goTypes,
		DependencyIndexes: file_albelt_stock_srv_const_proto_depIdxs,
		EnumInfos:         file_albelt_stock_srv_const_proto_enumTypes,
	}.Build()
	File_albelt_stock_srv_const_proto = out.File
	file_albelt_stock_srv_const_proto_rawDesc = nil
	file_albelt_stock_srv_const_proto_goTypes = nil
	file_albelt_stock_srv_const_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0--rc1
// source: albelt/stock_srv/msg.proto

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

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int32 `protobuf:"varint,2,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"` //商品id
	Number  int32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`                  //库存数量
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_albelt_stock_srv_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_albelt_stock_srv_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_albelt_stock_srv_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Stock) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *Stock) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type OrderDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderSn string `protobuf:"bytes,1,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
}

func (x *OrderDetails) Reset() {
	*x = OrderDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_albelt_stock_srv_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetails) ProtoMessage() {}

func (x *OrderDetails) ProtoReflect() protoreflect.Message {
	mi := &file_albelt_stock_srv_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetails.ProtoReflect.Descriptor instead.
func (*OrderDetails) Descriptor() ([]byte, []int) {
	return file_albelt_stock_srv_msg_proto_rawDescGZIP(), []int{1}
}

func (x *OrderDetails) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

var File_albelt_stock_srv_msg_proto protoreflect.FileDescriptor

var file_albelt_stock_srv_msg_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x6c, 0x62, 0x65, 0x6c, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x73,
	0x72, 0x76, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x05,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x6e, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x6f, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_albelt_stock_srv_msg_proto_rawDescOnce sync.Once
	file_albelt_stock_srv_msg_proto_rawDescData = file_albelt_stock_srv_msg_proto_rawDesc
)

func file_albelt_stock_srv_msg_proto_rawDescGZIP() []byte {
	file_albelt_stock_srv_msg_proto_rawDescOnce.Do(func() {
		file_albelt_stock_srv_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_albelt_stock_srv_msg_proto_rawDescData)
	})
	return file_albelt_stock_srv_msg_proto_rawDescData
}

var file_albelt_stock_srv_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_albelt_stock_srv_msg_proto_goTypes = []interface{}{
	(*Stock)(nil),        // 0: Stock
	(*OrderDetails)(nil), // 1: OrderDetails
}
var file_albelt_stock_srv_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_albelt_stock_srv_msg_proto_init() }
func file_albelt_stock_srv_msg_proto_init() {
	if File_albelt_stock_srv_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_albelt_stock_srv_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_albelt_stock_srv_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetails); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_albelt_stock_srv_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_albelt_stock_srv_msg_proto_goTypes,
		DependencyIndexes: file_albelt_stock_srv_msg_proto_depIdxs,
		MessageInfos:      file_albelt_stock_srv_msg_proto_msgTypes,
	}.Build()
	File_albelt_stock_srv_msg_proto = out.File
	file_albelt_stock_srv_msg_proto_rawDesc = nil
	file_albelt_stock_srv_msg_proto_goTypes = nil
	file_albelt_stock_srv_msg_proto_depIdxs = nil
}
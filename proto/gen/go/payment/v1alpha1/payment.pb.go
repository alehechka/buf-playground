// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: payment/v1alpha1/payment.proto

package paymentv1alpha1

import (
	money "google.golang.org/genproto/googleapis/type/money"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Purchase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   string       `protobuf:"bytes,1,opt,name=item_id,json=itemID,proto3" json:"item_id,omitempty"`
	Quantity float32      `protobuf:"fixed32,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UserId   string       `protobuf:"bytes,4,opt,name=user_id,json=userID,proto3" json:"user_id,omitempty"`
	Price    *money.Money `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Purchase) Reset() {
	*x = Purchase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1alpha1_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Purchase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Purchase) ProtoMessage() {}

func (x *Purchase) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1alpha1_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Purchase.ProtoReflect.Descriptor instead.
func (*Purchase) Descriptor() ([]byte, []int) {
	return file_payment_v1alpha1_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Purchase) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Purchase) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Purchase) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Purchase) GetPrice() *money.Money {
	if x != nil {
		return x.Price
	}
	return nil
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purchase  *Purchase              `protobuf:"bytes,1,opt,name=purchase,proto3" json:"purchase,omitempty"`
	ReceiptId string                 `protobuf:"bytes,2,opt,name=receipt_id,json=receiptID,proto3" json:"receipt_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1alpha1_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1alpha1_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_payment_v1alpha1_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Receipt) GetPurchase() *Purchase {
	if x != nil {
		return x.Purchase
	}
	return nil
}

func (x *Receipt) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *Receipt) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_payment_v1alpha1_payment_proto protoreflect.FileDescriptor

var file_payment_v1alpha1_payment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a,
	0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x9b, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x36, 0x0a,
	0x08, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x08, 0x70, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0xd8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x68, 0x65, 0x63, 0x68, 0x6b, 0x61, 0x2f, 0x62,
	0x75, 0x66, 0x2d, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58,
	0x58, 0xaa, 0x02, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payment_v1alpha1_payment_proto_rawDescOnce sync.Once
	file_payment_v1alpha1_payment_proto_rawDescData = file_payment_v1alpha1_payment_proto_rawDesc
)

func file_payment_v1alpha1_payment_proto_rawDescGZIP() []byte {
	file_payment_v1alpha1_payment_proto_rawDescOnce.Do(func() {
		file_payment_v1alpha1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_v1alpha1_payment_proto_rawDescData)
	})
	return file_payment_v1alpha1_payment_proto_rawDescData
}

var file_payment_v1alpha1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payment_v1alpha1_payment_proto_goTypes = []interface{}{
	(*Purchase)(nil),              // 0: payment.v1alpha1.Purchase
	(*Receipt)(nil),               // 1: payment.v1alpha1.Receipt
	(*money.Money)(nil),           // 2: google.type.Money
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_payment_v1alpha1_payment_proto_depIdxs = []int32{
	2, // 0: payment.v1alpha1.Purchase.price:type_name -> google.type.Money
	0, // 1: payment.v1alpha1.Receipt.purchase:type_name -> payment.v1alpha1.Purchase
	3, // 2: payment.v1alpha1.Receipt.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_payment_v1alpha1_payment_proto_init() }
func file_payment_v1alpha1_payment_proto_init() {
	if File_payment_v1alpha1_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_v1alpha1_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Purchase); i {
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
		file_payment_v1alpha1_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
			RawDescriptor: file_payment_v1alpha1_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_v1alpha1_payment_proto_goTypes,
		DependencyIndexes: file_payment_v1alpha1_payment_proto_depIdxs,
		MessageInfos:      file_payment_v1alpha1_payment_proto_msgTypes,
	}.Build()
	File_payment_v1alpha1_payment_proto = out.File
	file_payment_v1alpha1_payment_proto_rawDesc = nil
	file_payment_v1alpha1_payment_proto_goTypes = nil
	file_payment_v1alpha1_payment_proto_depIdxs = nil
}

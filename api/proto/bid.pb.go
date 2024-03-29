// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: bid.proto

package __

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

type BID_TYPE int32

const (
	BID_TYPE_FORWARD BID_TYPE = 0
	BID_TYPE_DUTCH   BID_TYPE = 1
)

// Enum value maps for BID_TYPE.
var (
	BID_TYPE_name = map[int32]string{
		0: "FORWARD",
		1: "DUTCH",
	}
	BID_TYPE_value = map[string]int32{
		"FORWARD": 0,
		"DUTCH":   1,
	}
)

func (x BID_TYPE) Enum() *BID_TYPE {
	p := new(BID_TYPE)
	*p = x
	return p
}

func (x BID_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BID_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_bid_proto_enumTypes[0].Descriptor()
}

func (BID_TYPE) Type() protoreflect.EnumType {
	return &file_bid_proto_enumTypes[0]
}

func (x BID_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BID_TYPE.Descriptor instead.
func (BID_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{0}
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId  string   `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Amount  float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	BidType BID_TYPE `protobuf:"varint,5,opt,name=bid_type,json=bidType,proto3,enum=proto.BID_TYPE" json:"bid_type,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BidRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *BidRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BidRequest) GetBidType() BID_TYPE {
	if x != nil {
		return x.BidType
	}
	return BID_TYPE_FORWARD
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidId uint32 `protobuf:"varint,1,opt,name=bid_id,json=bidId,proto3" json:"bid_id,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetBidId() uint32 {
	if x != nil {
		return x.BidId
	}
	return 0
}

type BidWinnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *BidWinnerRequest) Reset() {
	*x = BidWinnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidWinnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidWinnerRequest) ProtoMessage() {}

func (x *BidWinnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidWinnerRequest.ProtoReflect.Descriptor instead.
func (*BidWinnerRequest) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{2}
}

func (x *BidWinnerRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type BidWinnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BidWinnerResponse) Reset() {
	*x = BidWinnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidWinnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidWinnerResponse) ProtoMessage() {}

func (x *BidWinnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidWinnerResponse.ProtoReflect.Descriptor instead.
func (*BidWinnerResponse) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{3}
}

func (x *BidWinnerResponse) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *BidWinnerResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_bid_proto protoreflect.FileDescriptor

var file_bid_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x62,
	0x69, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x07,
	0x62, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x64, 0x49, 0x64, 0x22, 0x2b, 0x0a,
	0x10, 0x42, 0x69, 0x64, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x42, 0x69,
	0x64, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x2a, 0x22, 0x0a, 0x08, 0x42, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x55,
	0x54, 0x43, 0x48, 0x10, 0x01, 0x32, 0x7f, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x69, 0x64, 0x57,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x69, 0x64, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bid_proto_rawDescOnce sync.Once
	file_bid_proto_rawDescData = file_bid_proto_rawDesc
)

func file_bid_proto_rawDescGZIP() []byte {
	file_bid_proto_rawDescOnce.Do(func() {
		file_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_bid_proto_rawDescData)
	})
	return file_bid_proto_rawDescData
}

var file_bid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bid_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bid_proto_goTypes = []interface{}{
	(BID_TYPE)(0),             // 0: proto.BID_TYPE
	(*BidRequest)(nil),        // 1: proto.BidRequest
	(*BidResponse)(nil),       // 2: proto.BidResponse
	(*BidWinnerRequest)(nil),  // 3: proto.BidWinnerRequest
	(*BidWinnerResponse)(nil), // 4: proto.BidWinnerResponse
}
var file_bid_proto_depIdxs = []int32{
	0, // 0: proto.BidRequest.bid_type:type_name -> proto.BID_TYPE
	1, // 1: proto.Bid.PlaceBid:input_type -> proto.BidRequest
	3, // 2: proto.Bid.GetWinningBidder:input_type -> proto.BidWinnerRequest
	2, // 3: proto.Bid.PlaceBid:output_type -> proto.BidResponse
	4, // 4: proto.Bid.GetWinningBidder:output_type -> proto.BidWinnerResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bid_proto_init() }
func file_bid_proto_init() {
	if File_bid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidRequest); i {
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
		file_bid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidResponse); i {
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
		file_bid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidWinnerRequest); i {
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
		file_bid_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidWinnerResponse); i {
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
			RawDescriptor: file_bid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bid_proto_goTypes,
		DependencyIndexes: file_bid_proto_depIdxs,
		EnumInfos:         file_bid_proto_enumTypes,
		MessageInfos:      file_bid_proto_msgTypes,
	}.Build()
	File_bid_proto = out.File
	file_bid_proto_rawDesc = nil
	file_bid_proto_goTypes = nil
	file_bid_proto_depIdxs = nil
}

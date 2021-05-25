// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: transfer/v1/peer.proto

package transfer

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

type TransferPeerAnnounceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Channel uint64 `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *TransferPeerAnnounceRequest) Reset() {
	*x = TransferPeerAnnounceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferPeerAnnounceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPeerAnnounceRequest) ProtoMessage() {}

func (x *TransferPeerAnnounceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPeerAnnounceRequest.ProtoReflect.Descriptor instead.
func (*TransferPeerAnnounceRequest) Descriptor() ([]byte, []int) {
	return file_transfer_v1_peer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferPeerAnnounceRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TransferPeerAnnounceRequest) GetChannel() uint64 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type TransferPeerAnnounceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*TransferPeerAnnounceResponse_Channel
	Body isTransferPeerAnnounceResponse_Body `protobuf_oneof:"body"`
}

func (x *TransferPeerAnnounceResponse) Reset() {
	*x = TransferPeerAnnounceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferPeerAnnounceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPeerAnnounceResponse) ProtoMessage() {}

func (x *TransferPeerAnnounceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPeerAnnounceResponse.ProtoReflect.Descriptor instead.
func (*TransferPeerAnnounceResponse) Descriptor() ([]byte, []int) {
	return file_transfer_v1_peer_proto_rawDescGZIP(), []int{1}
}

func (m *TransferPeerAnnounceResponse) GetBody() isTransferPeerAnnounceResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *TransferPeerAnnounceResponse) GetChannel() uint64 {
	if x, ok := x.GetBody().(*TransferPeerAnnounceResponse_Channel); ok {
		return x.Channel
	}
	return 0
}

type isTransferPeerAnnounceResponse_Body interface {
	isTransferPeerAnnounceResponse_Body()
}

type TransferPeerAnnounceResponse_Channel struct {
	Channel uint64 `protobuf:"varint,1,opt,name=channel,proto3,oneof"`
}

func (*TransferPeerAnnounceResponse_Channel) isTransferPeerAnnounceResponse_Body() {}

type TransferPeerCloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransferPeerCloseRequest) Reset() {
	*x = TransferPeerCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferPeerCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPeerCloseRequest) ProtoMessage() {}

func (x *TransferPeerCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPeerCloseRequest.ProtoReflect.Descriptor instead.
func (*TransferPeerCloseRequest) Descriptor() ([]byte, []int) {
	return file_transfer_v1_peer_proto_rawDescGZIP(), []int{2}
}

func (x *TransferPeerCloseRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type TransferPeerCloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransferPeerCloseResponse) Reset() {
	*x = TransferPeerCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferPeerCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPeerCloseResponse) ProtoMessage() {}

func (x *TransferPeerCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPeerCloseResponse.ProtoReflect.Descriptor instead.
func (*TransferPeerCloseResponse) Descriptor() ([]byte, []int) {
	return file_transfer_v1_peer_proto_rawDescGZIP(), []int{3}
}

var File_transfer_v1_peer_proto protoreflect.FileDescriptor

var file_transfer_v1_peer_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x47, 0x0a, 0x1b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x42, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x50, 0x65, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x2a, 0x0a, 0x18, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xe3, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50,
	0x65, 0x65, 0x72, 0x12, 0x6d, 0x0a, 0x08, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x12,
	0x2f, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x65, 0x65,
	0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x65,
	0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x64, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x59, 0x0a, 0x15, 0x67, 0x67, 0x2e, 0x73,
	0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65,
	0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0xba, 0x02, 0x03,
	0x53, 0x54, 0x58, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_v1_peer_proto_rawDescOnce sync.Once
	file_transfer_v1_peer_proto_rawDescData = file_transfer_v1_peer_proto_rawDesc
)

func file_transfer_v1_peer_proto_rawDescGZIP() []byte {
	file_transfer_v1_peer_proto_rawDescOnce.Do(func() {
		file_transfer_v1_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_v1_peer_proto_rawDescData)
	})
	return file_transfer_v1_peer_proto_rawDescData
}

var file_transfer_v1_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transfer_v1_peer_proto_goTypes = []interface{}{
	(*TransferPeerAnnounceRequest)(nil),  // 0: strims.transfer.v1.TransferPeerAnnounceRequest
	(*TransferPeerAnnounceResponse)(nil), // 1: strims.transfer.v1.TransferPeerAnnounceResponse
	(*TransferPeerCloseRequest)(nil),     // 2: strims.transfer.v1.TransferPeerCloseRequest
	(*TransferPeerCloseResponse)(nil),    // 3: strims.transfer.v1.TransferPeerCloseResponse
}
var file_transfer_v1_peer_proto_depIdxs = []int32{
	0, // 0: strims.transfer.v1.TransferPeer.Announce:input_type -> strims.transfer.v1.TransferPeerAnnounceRequest
	2, // 1: strims.transfer.v1.TransferPeer.Close:input_type -> strims.transfer.v1.TransferPeerCloseRequest
	1, // 2: strims.transfer.v1.TransferPeer.Announce:output_type -> strims.transfer.v1.TransferPeerAnnounceResponse
	3, // 3: strims.transfer.v1.TransferPeer.Close:output_type -> strims.transfer.v1.TransferPeerCloseResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_v1_peer_proto_init() }
func file_transfer_v1_peer_proto_init() {
	if File_transfer_v1_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_v1_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferPeerAnnounceRequest); i {
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
		file_transfer_v1_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferPeerAnnounceResponse); i {
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
		file_transfer_v1_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferPeerCloseRequest); i {
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
		file_transfer_v1_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferPeerCloseResponse); i {
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
	file_transfer_v1_peer_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TransferPeerAnnounceResponse_Channel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transfer_v1_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_v1_peer_proto_goTypes,
		DependencyIndexes: file_transfer_v1_peer_proto_depIdxs,
		MessageInfos:      file_transfer_v1_peer_proto_msgTypes,
	}.Build()
	File_transfer_v1_peer_proto = out.File
	file_transfer_v1_peer_proto_rawDesc = nil
	file_transfer_v1_peer_proto_goTypes = nil
	file_transfer_v1_peer_proto_depIdxs = nil
}

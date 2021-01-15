// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: video/v1/egress.proto

package video

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EgressOpenStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SwarmUri string `protobuf:"bytes,1,opt,name=swarm_uri,json=swarmUri,proto3" json:"swarm_uri,omitempty"`
}

func (x *EgressOpenStreamRequest) Reset() {
	*x = EgressOpenStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_egress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressOpenStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressOpenStreamRequest) ProtoMessage() {}

func (x *EgressOpenStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_egress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressOpenStreamRequest.ProtoReflect.Descriptor instead.
func (*EgressOpenStreamRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_egress_proto_rawDescGZIP(), []int{0}
}

func (x *EgressOpenStreamRequest) GetSwarmUri() string {
	if x != nil {
		return x.SwarmUri
	}
	return ""
}

type EgressOpenStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*EgressOpenStreamResponse_Open_
	//	*EgressOpenStreamResponse_Data_
	//	*EgressOpenStreamResponse_Error_
	Body isEgressOpenStreamResponse_Body `protobuf_oneof:"body"`
}

func (x *EgressOpenStreamResponse) Reset() {
	*x = EgressOpenStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_egress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressOpenStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressOpenStreamResponse) ProtoMessage() {}

func (x *EgressOpenStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_egress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressOpenStreamResponse.ProtoReflect.Descriptor instead.
func (*EgressOpenStreamResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_egress_proto_rawDescGZIP(), []int{1}
}

func (m *EgressOpenStreamResponse) GetBody() isEgressOpenStreamResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *EgressOpenStreamResponse) GetOpen() *EgressOpenStreamResponse_Open {
	if x, ok := x.GetBody().(*EgressOpenStreamResponse_Open_); ok {
		return x.Open
	}
	return nil
}

func (x *EgressOpenStreamResponse) GetData() *EgressOpenStreamResponse_Data {
	if x, ok := x.GetBody().(*EgressOpenStreamResponse_Data_); ok {
		return x.Data
	}
	return nil
}

func (x *EgressOpenStreamResponse) GetError() *EgressOpenStreamResponse_Error {
	if x, ok := x.GetBody().(*EgressOpenStreamResponse_Error_); ok {
		return x.Error
	}
	return nil
}

type isEgressOpenStreamResponse_Body interface {
	isEgressOpenStreamResponse_Body()
}

type EgressOpenStreamResponse_Open_ struct {
	Open *EgressOpenStreamResponse_Open `protobuf:"bytes,1,opt,name=open,proto3,oneof"`
}

type EgressOpenStreamResponse_Data_ struct {
	Data *EgressOpenStreamResponse_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type EgressOpenStreamResponse_Error_ struct {
	Error *EgressOpenStreamResponse_Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*EgressOpenStreamResponse_Open_) isEgressOpenStreamResponse_Body() {}

func (*EgressOpenStreamResponse_Data_) isEgressOpenStreamResponse_Body() {}

func (*EgressOpenStreamResponse_Error_) isEgressOpenStreamResponse_Body() {}

type EgressOpenStreamResponse_Open struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId []byte `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
}

func (x *EgressOpenStreamResponse_Open) Reset() {
	*x = EgressOpenStreamResponse_Open{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_egress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressOpenStreamResponse_Open) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressOpenStreamResponse_Open) ProtoMessage() {}

func (x *EgressOpenStreamResponse_Open) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_egress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressOpenStreamResponse_Open.ProtoReflect.Descriptor instead.
func (*EgressOpenStreamResponse_Open) Descriptor() ([]byte, []int) {
	return file_video_v1_egress_proto_rawDescGZIP(), []int{1, 0}
}

func (x *EgressOpenStreamResponse_Open) GetTransferId() []byte {
	if x != nil {
		return x.TransferId
	}
	return nil
}

type EgressOpenStreamResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data           []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	SegmentEnd     bool   `protobuf:"varint,2,opt,name=segment_end,json=segmentEnd,proto3" json:"segment_end,omitempty"`
	BufferUnderrun bool   `protobuf:"varint,3,opt,name=buffer_underrun,json=bufferUnderrun,proto3" json:"buffer_underrun,omitempty"`
}

func (x *EgressOpenStreamResponse_Data) Reset() {
	*x = EgressOpenStreamResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_egress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressOpenStreamResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressOpenStreamResponse_Data) ProtoMessage() {}

func (x *EgressOpenStreamResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_egress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressOpenStreamResponse_Data.ProtoReflect.Descriptor instead.
func (*EgressOpenStreamResponse_Data) Descriptor() ([]byte, []int) {
	return file_video_v1_egress_proto_rawDescGZIP(), []int{1, 1}
}

func (x *EgressOpenStreamResponse_Data) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EgressOpenStreamResponse_Data) GetSegmentEnd() bool {
	if x != nil {
		return x.SegmentEnd
	}
	return false
}

func (x *EgressOpenStreamResponse_Data) GetBufferUnderrun() bool {
	if x != nil {
		return x.BufferUnderrun
	}
	return false
}

type EgressOpenStreamResponse_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EgressOpenStreamResponse_Error) Reset() {
	*x = EgressOpenStreamResponse_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_egress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressOpenStreamResponse_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressOpenStreamResponse_Error) ProtoMessage() {}

func (x *EgressOpenStreamResponse_Error) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_egress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressOpenStreamResponse_Error.ProtoReflect.Descriptor instead.
func (*EgressOpenStreamResponse_Error) Descriptor() ([]byte, []int) {
	return file_video_v1_egress_proto_rawDescGZIP(), []int{1, 2}
}

func (x *EgressOpenStreamResponse_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_video_v1_egress_proto protoreflect.FileDescriptor

var file_video_v1_egress_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x36, 0x0a, 0x17, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x55, 0x72, 0x69,
	0x22, 0xa9, 0x03, 0x0a, 0x18, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d,
	0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x1a, 0x27, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x64, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x72, 0x75, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x72, 0x75,
	0x6e, 0x1a, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x6d, 0x0a, 0x06,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x63, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x50, 0x0a, 0x12, 0x67,
	0x67, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65,
	0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0xba, 0x02, 0x03, 0x53, 0x56, 0x4f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_v1_egress_proto_rawDescOnce sync.Once
	file_video_v1_egress_proto_rawDescData = file_video_v1_egress_proto_rawDesc
)

func file_video_v1_egress_proto_rawDescGZIP() []byte {
	file_video_v1_egress_proto_rawDescOnce.Do(func() {
		file_video_v1_egress_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_v1_egress_proto_rawDescData)
	})
	return file_video_v1_egress_proto_rawDescData
}

var file_video_v1_egress_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_video_v1_egress_proto_goTypes = []interface{}{
	(*EgressOpenStreamRequest)(nil),        // 0: strims.video.v1.EgressOpenStreamRequest
	(*EgressOpenStreamResponse)(nil),       // 1: strims.video.v1.EgressOpenStreamResponse
	(*EgressOpenStreamResponse_Open)(nil),  // 2: strims.video.v1.EgressOpenStreamResponse.Open
	(*EgressOpenStreamResponse_Data)(nil),  // 3: strims.video.v1.EgressOpenStreamResponse.Data
	(*EgressOpenStreamResponse_Error)(nil), // 4: strims.video.v1.EgressOpenStreamResponse.Error
}
var file_video_v1_egress_proto_depIdxs = []int32{
	2, // 0: strims.video.v1.EgressOpenStreamResponse.open:type_name -> strims.video.v1.EgressOpenStreamResponse.Open
	3, // 1: strims.video.v1.EgressOpenStreamResponse.data:type_name -> strims.video.v1.EgressOpenStreamResponse.Data
	4, // 2: strims.video.v1.EgressOpenStreamResponse.error:type_name -> strims.video.v1.EgressOpenStreamResponse.Error
	0, // 3: strims.video.v1.Egress.OpenStream:input_type -> strims.video.v1.EgressOpenStreamRequest
	1, // 4: strims.video.v1.Egress.OpenStream:output_type -> strims.video.v1.EgressOpenStreamResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_video_v1_egress_proto_init() }
func file_video_v1_egress_proto_init() {
	if File_video_v1_egress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_v1_egress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressOpenStreamRequest); i {
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
		file_video_v1_egress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressOpenStreamResponse); i {
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
		file_video_v1_egress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressOpenStreamResponse_Open); i {
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
		file_video_v1_egress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressOpenStreamResponse_Data); i {
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
		file_video_v1_egress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressOpenStreamResponse_Error); i {
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
	file_video_v1_egress_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EgressOpenStreamResponse_Open_)(nil),
		(*EgressOpenStreamResponse_Data_)(nil),
		(*EgressOpenStreamResponse_Error_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_v1_egress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_v1_egress_proto_goTypes,
		DependencyIndexes: file_video_v1_egress_proto_depIdxs,
		MessageInfos:      file_video_v1_egress_proto_msgTypes,
	}.Build()
	File_video_v1_egress_proto = out.File
	file_video_v1_egress_proto_rawDesc = nil
	file_video_v1_egress_proto_goTypes = nil
	file_video_v1_egress_proto_depIdxs = nil
}

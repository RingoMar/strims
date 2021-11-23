// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: vpn/v1/peer_index.proto

package vpn

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

type PeerIndexMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*PeerIndexMessage_Publish_
	//	*PeerIndexMessage_Unpublish_
	//	*PeerIndexMessage_SearchRequest_
	//	*PeerIndexMessage_SearchResponse_
	Body isPeerIndexMessage_Body `protobuf_oneof:"body"`
}

func (x *PeerIndexMessage) Reset() {
	*x = PeerIndexMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerIndexMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerIndexMessage) ProtoMessage() {}

func (x *PeerIndexMessage) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerIndexMessage.ProtoReflect.Descriptor instead.
func (*PeerIndexMessage) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_index_proto_rawDescGZIP(), []int{0}
}

func (m *PeerIndexMessage) GetBody() isPeerIndexMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *PeerIndexMessage) GetPublish() *PeerIndexMessage_Publish {
	if x, ok := x.GetBody().(*PeerIndexMessage_Publish_); ok {
		return x.Publish
	}
	return nil
}

func (x *PeerIndexMessage) GetUnpublish() *PeerIndexMessage_Unpublish {
	if x, ok := x.GetBody().(*PeerIndexMessage_Unpublish_); ok {
		return x.Unpublish
	}
	return nil
}

func (x *PeerIndexMessage) GetSearchRequest() *PeerIndexMessage_SearchRequest {
	if x, ok := x.GetBody().(*PeerIndexMessage_SearchRequest_); ok {
		return x.SearchRequest
	}
	return nil
}

func (x *PeerIndexMessage) GetSearchResponse() *PeerIndexMessage_SearchResponse {
	if x, ok := x.GetBody().(*PeerIndexMessage_SearchResponse_); ok {
		return x.SearchResponse
	}
	return nil
}

type isPeerIndexMessage_Body interface {
	isPeerIndexMessage_Body()
}

type PeerIndexMessage_Publish_ struct {
	Publish *PeerIndexMessage_Publish `protobuf:"bytes,1,opt,name=publish,proto3,oneof"`
}

type PeerIndexMessage_Unpublish_ struct {
	Unpublish *PeerIndexMessage_Unpublish `protobuf:"bytes,2,opt,name=unpublish,proto3,oneof"`
}

type PeerIndexMessage_SearchRequest_ struct {
	SearchRequest *PeerIndexMessage_SearchRequest `protobuf:"bytes,3,opt,name=search_request,json=searchRequest,proto3,oneof"`
}

type PeerIndexMessage_SearchResponse_ struct {
	SearchResponse *PeerIndexMessage_SearchResponse `protobuf:"bytes,4,opt,name=search_response,json=searchResponse,proto3,oneof"`
}

func (*PeerIndexMessage_Publish_) isPeerIndexMessage_Body() {}

func (*PeerIndexMessage_Unpublish_) isPeerIndexMessage_Body() {}

func (*PeerIndexMessage_SearchRequest_) isPeerIndexMessage_Body() {}

func (*PeerIndexMessage_SearchResponse_) isPeerIndexMessage_Body() {}

type PeerIndexMessage_Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Key       []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	HostId    []byte `protobuf:"bytes,3,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Port      uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *PeerIndexMessage_Record) Reset() {
	*x = PeerIndexMessage_Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerIndexMessage_Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerIndexMessage_Record) ProtoMessage() {}

func (x *PeerIndexMessage_Record) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerIndexMessage_Record.ProtoReflect.Descriptor instead.
func (*PeerIndexMessage_Record) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_index_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PeerIndexMessage_Record) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *PeerIndexMessage_Record) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PeerIndexMessage_Record) GetHostId() []byte {
	if x != nil {
		return x.HostId
	}
	return nil
}

func (x *PeerIndexMessage_Record) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PeerIndexMessage_Record) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PeerIndexMessage_Record) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type PeerIndexMessage_Publish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *PeerIndexMessage_Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *PeerIndexMessage_Publish) Reset() {
	*x = PeerIndexMessage_Publish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerIndexMessage_Publish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerIndexMessage_Publish) ProtoMessage() {}

func (x *PeerIndexMessage_Publish) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerIndexMessage_Publish.ProtoReflect.Descriptor instead.
func (*PeerIndexMessage_Publish) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_index_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PeerIndexMessage_Publish) GetRecord() *PeerIndexMessage_Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type PeerIndexMessage_Unpublish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *PeerIndexMessage_Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *PeerIndexMessage_Unpublish) Reset() {
	*x = PeerIndexMessage_Unpublish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerIndexMessage_Unpublish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerIndexMessage_Unpublish) ProtoMessage() {}

func (x *PeerIndexMessage_Unpublish) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_index_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerIndexMessage_Unpublish.ProtoReflect.Descriptor instead.
func (*PeerIndexMessage_Unpublish) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_index_proto_rawDescGZIP(), []int{0, 2}
}

func (x *PeerIndexMessage_Unpublish) GetRecord() *PeerIndexMessage_Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type PeerIndexMessage_SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Hash      []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *PeerIndexMessage_SearchRequest) Reset() {
	*x = PeerIndexMessage_SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerIndexMessage_SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerIndexMessage_SearchRequest) ProtoMessage() {}

func (x *PeerIndexMessage_SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_index_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerIndexMessage_SearchRequest.ProtoReflect.Descriptor instead.
func (*PeerIndexMessage_SearchRequest) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_index_proto_rawDescGZIP(), []int{0, 3}
}

func (x *PeerIndexMessage_SearchRequest) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *PeerIndexMessage_SearchRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type PeerIndexMessage_SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId uint64                     `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Records   []*PeerIndexMessage_Record `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *PeerIndexMessage_SearchResponse) Reset() {
	*x = PeerIndexMessage_SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_index_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerIndexMessage_SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerIndexMessage_SearchResponse) ProtoMessage() {}

func (x *PeerIndexMessage_SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_index_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerIndexMessage_SearchResponse.ProtoReflect.Descriptor instead.
func (*PeerIndexMessage_SearchResponse) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_index_proto_rawDescGZIP(), []int{0, 4}
}

func (x *PeerIndexMessage_SearchResponse) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *PeerIndexMessage_SearchResponse) GetRecords() []*PeerIndexMessage_Record {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_vpn_v1_peer_index_proto protoreflect.FileDescriptor

var file_vpn_v1_peer_index_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x70, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6d,
	0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xc6, 0x06, 0x0a, 0x10, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a,
	0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x48, 0x00, 0x52, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x49, 0x0a, 0x09, 0x75, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76,
	0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x48, 0x00, 0x52, 0x09, 0x75, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x56, 0x0a,
	0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76,
	0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x97, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x49, 0x0a, 0x07, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76,
	0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x4b, 0x0a, 0x09, 0x55, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x1a, 0x42, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x1a, 0x71, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d,
	0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x42, 0x4a, 0x0a, 0x10, 0x67, 0x67, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76,
	0x70, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x70,
	0x73, 0x70, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x70, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x70, 0x6e, 0xba, 0x02, 0x03, 0x53, 0x56, 0x4e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vpn_v1_peer_index_proto_rawDescOnce sync.Once
	file_vpn_v1_peer_index_proto_rawDescData = file_vpn_v1_peer_index_proto_rawDesc
)

func file_vpn_v1_peer_index_proto_rawDescGZIP() []byte {
	file_vpn_v1_peer_index_proto_rawDescOnce.Do(func() {
		file_vpn_v1_peer_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_vpn_v1_peer_index_proto_rawDescData)
	})
	return file_vpn_v1_peer_index_proto_rawDescData
}

var file_vpn_v1_peer_index_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vpn_v1_peer_index_proto_goTypes = []interface{}{
	(*PeerIndexMessage)(nil),                // 0: strims.vpn.v1.PeerIndexMessage
	(*PeerIndexMessage_Record)(nil),         // 1: strims.vpn.v1.PeerIndexMessage.Record
	(*PeerIndexMessage_Publish)(nil),        // 2: strims.vpn.v1.PeerIndexMessage.Publish
	(*PeerIndexMessage_Unpublish)(nil),      // 3: strims.vpn.v1.PeerIndexMessage.Unpublish
	(*PeerIndexMessage_SearchRequest)(nil),  // 4: strims.vpn.v1.PeerIndexMessage.SearchRequest
	(*PeerIndexMessage_SearchResponse)(nil), // 5: strims.vpn.v1.PeerIndexMessage.SearchResponse
}
var file_vpn_v1_peer_index_proto_depIdxs = []int32{
	2, // 0: strims.vpn.v1.PeerIndexMessage.publish:type_name -> strims.vpn.v1.PeerIndexMessage.Publish
	3, // 1: strims.vpn.v1.PeerIndexMessage.unpublish:type_name -> strims.vpn.v1.PeerIndexMessage.Unpublish
	4, // 2: strims.vpn.v1.PeerIndexMessage.search_request:type_name -> strims.vpn.v1.PeerIndexMessage.SearchRequest
	5, // 3: strims.vpn.v1.PeerIndexMessage.search_response:type_name -> strims.vpn.v1.PeerIndexMessage.SearchResponse
	1, // 4: strims.vpn.v1.PeerIndexMessage.Publish.record:type_name -> strims.vpn.v1.PeerIndexMessage.Record
	1, // 5: strims.vpn.v1.PeerIndexMessage.Unpublish.record:type_name -> strims.vpn.v1.PeerIndexMessage.Record
	1, // 6: strims.vpn.v1.PeerIndexMessage.SearchResponse.records:type_name -> strims.vpn.v1.PeerIndexMessage.Record
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_vpn_v1_peer_index_proto_init() }
func file_vpn_v1_peer_index_proto_init() {
	if File_vpn_v1_peer_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vpn_v1_peer_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerIndexMessage); i {
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
		file_vpn_v1_peer_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerIndexMessage_Record); i {
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
		file_vpn_v1_peer_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerIndexMessage_Publish); i {
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
		file_vpn_v1_peer_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerIndexMessage_Unpublish); i {
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
		file_vpn_v1_peer_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerIndexMessage_SearchRequest); i {
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
		file_vpn_v1_peer_index_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerIndexMessage_SearchResponse); i {
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
	file_vpn_v1_peer_index_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PeerIndexMessage_Publish_)(nil),
		(*PeerIndexMessage_Unpublish_)(nil),
		(*PeerIndexMessage_SearchRequest_)(nil),
		(*PeerIndexMessage_SearchResponse_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vpn_v1_peer_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vpn_v1_peer_index_proto_goTypes,
		DependencyIndexes: file_vpn_v1_peer_index_proto_depIdxs,
		MessageInfos:      file_vpn_v1_peer_index_proto_msgTypes,
	}.Build()
	File_vpn_v1_peer_index_proto = out.File
	file_vpn_v1_peer_index_proto_rawDesc = nil
	file_vpn_v1_peer_index_proto_goTypes = nil
	file_vpn_v1_peer_index_proto_depIdxs = nil
}

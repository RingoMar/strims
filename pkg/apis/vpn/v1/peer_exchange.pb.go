// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: vpn/v1/peer_exchange.proto

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

type PeerExchangeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*PeerExchangeMessage_Request_
	//	*PeerExchangeMessage_Response_
	//	*PeerExchangeMessage_Offer_
	//	*PeerExchangeMessage_Answer_
	//	*PeerExchangeMessage_IceCandidate_
	//	*PeerExchangeMessage_CallbackRequest_
	Body isPeerExchangeMessage_Body `protobuf_oneof:"body"`
}

func (x *PeerExchangeMessage) Reset() {
	*x = PeerExchangeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerExchangeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerExchangeMessage) ProtoMessage() {}

func (x *PeerExchangeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerExchangeMessage.ProtoReflect.Descriptor instead.
func (*PeerExchangeMessage) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_exchange_proto_rawDescGZIP(), []int{0}
}

func (m *PeerExchangeMessage) GetBody() isPeerExchangeMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *PeerExchangeMessage) GetRequest() *PeerExchangeMessage_Request {
	if x, ok := x.GetBody().(*PeerExchangeMessage_Request_); ok {
		return x.Request
	}
	return nil
}

func (x *PeerExchangeMessage) GetResponse() *PeerExchangeMessage_Response {
	if x, ok := x.GetBody().(*PeerExchangeMessage_Response_); ok {
		return x.Response
	}
	return nil
}

func (x *PeerExchangeMessage) GetOffer() *PeerExchangeMessage_Offer {
	if x, ok := x.GetBody().(*PeerExchangeMessage_Offer_); ok {
		return x.Offer
	}
	return nil
}

func (x *PeerExchangeMessage) GetAnswer() *PeerExchangeMessage_Answer {
	if x, ok := x.GetBody().(*PeerExchangeMessage_Answer_); ok {
		return x.Answer
	}
	return nil
}

func (x *PeerExchangeMessage) GetIceCandidate() *PeerExchangeMessage_IceCandidate {
	if x, ok := x.GetBody().(*PeerExchangeMessage_IceCandidate_); ok {
		return x.IceCandidate
	}
	return nil
}

func (x *PeerExchangeMessage) GetCallbackRequest() *PeerExchangeMessage_CallbackRequest {
	if x, ok := x.GetBody().(*PeerExchangeMessage_CallbackRequest_); ok {
		return x.CallbackRequest
	}
	return nil
}

type isPeerExchangeMessage_Body interface {
	isPeerExchangeMessage_Body()
}

type PeerExchangeMessage_Request_ struct {
	Request *PeerExchangeMessage_Request `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type PeerExchangeMessage_Response_ struct {
	Response *PeerExchangeMessage_Response `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

type PeerExchangeMessage_Offer_ struct {
	Offer *PeerExchangeMessage_Offer `protobuf:"bytes,3,opt,name=offer,proto3,oneof"`
}

type PeerExchangeMessage_Answer_ struct {
	Answer *PeerExchangeMessage_Answer `protobuf:"bytes,4,opt,name=answer,proto3,oneof"`
}

type PeerExchangeMessage_IceCandidate_ struct {
	IceCandidate *PeerExchangeMessage_IceCandidate `protobuf:"bytes,5,opt,name=ice_candidate,json=iceCandidate,proto3,oneof"`
}

type PeerExchangeMessage_CallbackRequest_ struct {
	CallbackRequest *PeerExchangeMessage_CallbackRequest `protobuf:"bytes,6,opt,name=callback_request,json=callbackRequest,proto3,oneof"`
}

func (*PeerExchangeMessage_Request_) isPeerExchangeMessage_Body() {}

func (*PeerExchangeMessage_Response_) isPeerExchangeMessage_Body() {}

func (*PeerExchangeMessage_Offer_) isPeerExchangeMessage_Body() {}

func (*PeerExchangeMessage_Answer_) isPeerExchangeMessage_Body() {}

func (*PeerExchangeMessage_IceCandidate_) isPeerExchangeMessage_Body() {}

func (*PeerExchangeMessage_CallbackRequest_) isPeerExchangeMessage_Body() {}

type PeerExchangeMessage_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PeerExchangeMessage_Request) Reset() {
	*x = PeerExchangeMessage_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerExchangeMessage_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerExchangeMessage_Request) ProtoMessage() {}

func (x *PeerExchangeMessage_Request) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerExchangeMessage_Request.ProtoReflect.Descriptor instead.
func (*PeerExchangeMessage_Request) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_exchange_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PeerExchangeMessage_Request) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PeerExchangeMessage_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids [][]byte `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *PeerExchangeMessage_Response) Reset() {
	*x = PeerExchangeMessage_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerExchangeMessage_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerExchangeMessage_Response) ProtoMessage() {}

func (x *PeerExchangeMessage_Response) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerExchangeMessage_Response.ProtoReflect.Descriptor instead.
func (*PeerExchangeMessage_Response) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_exchange_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PeerExchangeMessage_Response) GetIds() [][]byte {
	if x != nil {
		return x.Ids
	}
	return nil
}

type PeerExchangeMessage_Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediationId uint64 `protobuf:"varint,1,opt,name=mediation_id,json=mediationId,proto3" json:"mediation_id,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PeerExchangeMessage_Offer) Reset() {
	*x = PeerExchangeMessage_Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerExchangeMessage_Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerExchangeMessage_Offer) ProtoMessage() {}

func (x *PeerExchangeMessage_Offer) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerExchangeMessage_Offer.ProtoReflect.Descriptor instead.
func (*PeerExchangeMessage_Offer) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_exchange_proto_rawDescGZIP(), []int{0, 2}
}

func (x *PeerExchangeMessage_Offer) GetMediationId() uint64 {
	if x != nil {
		return x.MediationId
	}
	return 0
}

func (x *PeerExchangeMessage_Offer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PeerExchangeMessage_Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediationId uint64 `protobuf:"varint,1,opt,name=mediation_id,json=mediationId,proto3" json:"mediation_id,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PeerExchangeMessage_Answer) Reset() {
	*x = PeerExchangeMessage_Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerExchangeMessage_Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerExchangeMessage_Answer) ProtoMessage() {}

func (x *PeerExchangeMessage_Answer) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerExchangeMessage_Answer.ProtoReflect.Descriptor instead.
func (*PeerExchangeMessage_Answer) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_exchange_proto_rawDescGZIP(), []int{0, 3}
}

func (x *PeerExchangeMessage_Answer) GetMediationId() uint64 {
	if x != nil {
		return x.MediationId
	}
	return 0
}

func (x *PeerExchangeMessage_Answer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PeerExchangeMessage_IceCandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediationId uint64 `protobuf:"varint,1,opt,name=mediation_id,json=mediationId,proto3" json:"mediation_id,omitempty"`
	Index       uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Data        []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PeerExchangeMessage_IceCandidate) Reset() {
	*x = PeerExchangeMessage_IceCandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerExchangeMessage_IceCandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerExchangeMessage_IceCandidate) ProtoMessage() {}

func (x *PeerExchangeMessage_IceCandidate) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_exchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerExchangeMessage_IceCandidate.ProtoReflect.Descriptor instead.
func (*PeerExchangeMessage_IceCandidate) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_exchange_proto_rawDescGZIP(), []int{0, 4}
}

func (x *PeerExchangeMessage_IceCandidate) GetMediationId() uint64 {
	if x != nil {
		return x.MediationId
	}
	return 0
}

func (x *PeerExchangeMessage_IceCandidate) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *PeerExchangeMessage_IceCandidate) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PeerExchangeMessage_CallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PeerExchangeMessage_CallbackRequest) Reset() {
	*x = PeerExchangeMessage_CallbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_v1_peer_exchange_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerExchangeMessage_CallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerExchangeMessage_CallbackRequest) ProtoMessage() {}

func (x *PeerExchangeMessage_CallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_v1_peer_exchange_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerExchangeMessage_CallbackRequest.ProtoReflect.Descriptor instead.
func (*PeerExchangeMessage_CallbackRequest) Descriptor() ([]byte, []int) {
	return file_vpn_v1_peer_exchange_proto_rawDescGZIP(), []int{0, 5}
}

var File_vpn_v1_peer_exchange_proto protoreflect.FileDescriptor

var file_vpn_v1_peer_exchange_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x70, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xa0, 0x06, 0x0a, 0x13,
	0x50, 0x65, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76,
	0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d,
	0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x56, 0x0a,
	0x0d, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x1a, 0x3e, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3f, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x5b, 0x0a, 0x0c, 0x49, 0x63, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x11, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x4a,
	0x0a, 0x10, 0x67, 0x67, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x70, 0x6e, 0x2e,
	0x76, 0x31, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d,
	0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x70, 0x73, 0x70, 0x70,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x70, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x70, 0x6e, 0xba, 0x02, 0x03, 0x53, 0x56, 0x4e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vpn_v1_peer_exchange_proto_rawDescOnce sync.Once
	file_vpn_v1_peer_exchange_proto_rawDescData = file_vpn_v1_peer_exchange_proto_rawDesc
)

func file_vpn_v1_peer_exchange_proto_rawDescGZIP() []byte {
	file_vpn_v1_peer_exchange_proto_rawDescOnce.Do(func() {
		file_vpn_v1_peer_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_vpn_v1_peer_exchange_proto_rawDescData)
	})
	return file_vpn_v1_peer_exchange_proto_rawDescData
}

var file_vpn_v1_peer_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_vpn_v1_peer_exchange_proto_goTypes = []interface{}{
	(*PeerExchangeMessage)(nil),                 // 0: strims.vpn.v1.PeerExchangeMessage
	(*PeerExchangeMessage_Request)(nil),         // 1: strims.vpn.v1.PeerExchangeMessage.Request
	(*PeerExchangeMessage_Response)(nil),        // 2: strims.vpn.v1.PeerExchangeMessage.Response
	(*PeerExchangeMessage_Offer)(nil),           // 3: strims.vpn.v1.PeerExchangeMessage.Offer
	(*PeerExchangeMessage_Answer)(nil),          // 4: strims.vpn.v1.PeerExchangeMessage.Answer
	(*PeerExchangeMessage_IceCandidate)(nil),    // 5: strims.vpn.v1.PeerExchangeMessage.IceCandidate
	(*PeerExchangeMessage_CallbackRequest)(nil), // 6: strims.vpn.v1.PeerExchangeMessage.CallbackRequest
}
var file_vpn_v1_peer_exchange_proto_depIdxs = []int32{
	1, // 0: strims.vpn.v1.PeerExchangeMessage.request:type_name -> strims.vpn.v1.PeerExchangeMessage.Request
	2, // 1: strims.vpn.v1.PeerExchangeMessage.response:type_name -> strims.vpn.v1.PeerExchangeMessage.Response
	3, // 2: strims.vpn.v1.PeerExchangeMessage.offer:type_name -> strims.vpn.v1.PeerExchangeMessage.Offer
	4, // 3: strims.vpn.v1.PeerExchangeMessage.answer:type_name -> strims.vpn.v1.PeerExchangeMessage.Answer
	5, // 4: strims.vpn.v1.PeerExchangeMessage.ice_candidate:type_name -> strims.vpn.v1.PeerExchangeMessage.IceCandidate
	6, // 5: strims.vpn.v1.PeerExchangeMessage.callback_request:type_name -> strims.vpn.v1.PeerExchangeMessage.CallbackRequest
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_vpn_v1_peer_exchange_proto_init() }
func file_vpn_v1_peer_exchange_proto_init() {
	if File_vpn_v1_peer_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vpn_v1_peer_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerExchangeMessage); i {
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
		file_vpn_v1_peer_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerExchangeMessage_Request); i {
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
		file_vpn_v1_peer_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerExchangeMessage_Response); i {
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
		file_vpn_v1_peer_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerExchangeMessage_Offer); i {
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
		file_vpn_v1_peer_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerExchangeMessage_Answer); i {
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
		file_vpn_v1_peer_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerExchangeMessage_IceCandidate); i {
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
		file_vpn_v1_peer_exchange_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerExchangeMessage_CallbackRequest); i {
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
	file_vpn_v1_peer_exchange_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PeerExchangeMessage_Request_)(nil),
		(*PeerExchangeMessage_Response_)(nil),
		(*PeerExchangeMessage_Offer_)(nil),
		(*PeerExchangeMessage_Answer_)(nil),
		(*PeerExchangeMessage_IceCandidate_)(nil),
		(*PeerExchangeMessage_CallbackRequest_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vpn_v1_peer_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vpn_v1_peer_exchange_proto_goTypes,
		DependencyIndexes: file_vpn_v1_peer_exchange_proto_depIdxs,
		MessageInfos:      file_vpn_v1_peer_exchange_proto_msgTypes,
	}.Build()
	File_vpn_v1_peer_exchange_proto = out.File
	file_vpn_v1_peer_exchange_proto_rawDesc = nil
	file_vpn_v1_peer_exchange_proto_goTypes = nil
	file_vpn_v1_peer_exchange_proto_depIdxs = nil
}

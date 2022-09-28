// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: replication/v1/replication.proto

package replicationv1

import (
	v11 "github.com/MemeLabs/strims/pkg/apis/auth/v1"
	v1 "github.com/MemeLabs/strims/pkg/apis/dao/v1"
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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicaId      uint32            `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	CurrentVersion *v1.VersionVector `protobuf:"bytes,2,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replication_v1_replication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_replication_v1_replication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_replication_v1_replication_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetReplicaId() uint32 {
	if x != nil {
		return x.ReplicaId
	}
	return 0
}

func (x *Version) GetCurrentVersion() *v1.VersionVector {
	if x != nil {
		return x.CurrentVersion
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace int64             `protobuf:"varint,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id        uint64            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Version   *v1.VersionVector `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Delete    bool              `protobuf:"varint,4,opt,name=delete,proto3" json:"delete,omitempty"`
	Record    []byte            `protobuf:"bytes,5,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replication_v1_replication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_replication_v1_replication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_replication_v1_replication_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetNamespace() int64 {
	if x != nil {
		return x.Namespace
	}
	return 0
}

func (x *Event) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetVersion() *v1.VersionVector {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Event) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *Event) GetRecord() []byte {
	if x != nil {
		return x.Record
	}
	return nil
}

type EventBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventBundle) Reset() {
	*x = EventBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replication_v1_replication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBundle) ProtoMessage() {}

func (x *EventBundle) ProtoReflect() protoreflect.Message {
	mi := &file_replication_v1_replication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventBundle.ProtoReflect.Descriptor instead.
func (*EventBundle) Descriptor() ([]byte, []int) {
	return file_replication_v1_replication_proto_rawDescGZIP(), []int{2}
}

func (x *EventBundle) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type EventLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReplicaId uint32   `protobuf:"varint,2,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	Events    []*Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventLog) Reset() {
	*x = EventLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replication_v1_replication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLog) ProtoMessage() {}

func (x *EventLog) ProtoReflect() protoreflect.Message {
	mi := &file_replication_v1_replication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLog.ProtoReflect.Descriptor instead.
func (*EventLog) Descriptor() ([]byte, []int) {
	return file_replication_v1_replication_proto_rawDescGZIP(), []int{3}
}

func (x *EventLog) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventLog) GetReplicaId() uint32 {
	if x != nil {
		return x.ReplicaId
	}
	return 0
}

func (x *EventLog) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type CreatePairingTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId   uint64 `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	BootstrapId uint64 `protobuf:"varint,2,opt,name=bootstrap_id,json=bootstrapId,proto3" json:"bootstrap_id,omitempty"`
}

func (x *CreatePairingTokenRequest) Reset() {
	*x = CreatePairingTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replication_v1_replication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePairingTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePairingTokenRequest) ProtoMessage() {}

func (x *CreatePairingTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_replication_v1_replication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePairingTokenRequest.ProtoReflect.Descriptor instead.
func (*CreatePairingTokenRequest) Descriptor() ([]byte, []int) {
	return file_replication_v1_replication_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePairingTokenRequest) GetNetworkId() uint64 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *CreatePairingTokenRequest) GetBootstrapId() uint64 {
	if x != nil {
		return x.BootstrapId
	}
	return 0
}

type CreatePairingTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *v11.PairingToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreatePairingTokenResponse) Reset() {
	*x = CreatePairingTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replication_v1_replication_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePairingTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePairingTokenResponse) ProtoMessage() {}

func (x *CreatePairingTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_replication_v1_replication_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePairingTokenResponse.ProtoReflect.Descriptor instead.
func (*CreatePairingTokenResponse) Descriptor() ([]byte, []int) {
	return file_replication_v1_replication_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePairingTokenResponse) GetToken() *v11.PairingToken {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_replication_v1_replication_proto protoreflect.FileDescriptor

var file_replication_v1_replication_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64,
	0x61, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x9d, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x64, 0x61, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x22, 0x43, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12,
	0x34, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6f, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5d, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x90, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x12,
	0x79, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73,
	0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x62, 0x0a, 0x18, 0x67, 0x67,
	0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xba, 0x02, 0x03, 0x53, 0x52, 0x45, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_replication_v1_replication_proto_rawDescOnce sync.Once
	file_replication_v1_replication_proto_rawDescData = file_replication_v1_replication_proto_rawDesc
)

func file_replication_v1_replication_proto_rawDescGZIP() []byte {
	file_replication_v1_replication_proto_rawDescOnce.Do(func() {
		file_replication_v1_replication_proto_rawDescData = protoimpl.X.CompressGZIP(file_replication_v1_replication_proto_rawDescData)
	})
	return file_replication_v1_replication_proto_rawDescData
}

var file_replication_v1_replication_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_replication_v1_replication_proto_goTypes = []interface{}{
	(*Version)(nil),                    // 0: strims.replication.v1.Version
	(*Event)(nil),                      // 1: strims.replication.v1.Event
	(*EventBundle)(nil),                // 2: strims.replication.v1.EventBundle
	(*EventLog)(nil),                   // 3: strims.replication.v1.EventLog
	(*CreatePairingTokenRequest)(nil),  // 4: strims.replication.v1.CreatePairingTokenRequest
	(*CreatePairingTokenResponse)(nil), // 5: strims.replication.v1.CreatePairingTokenResponse
	(*v1.VersionVector)(nil),           // 6: strims.dao.v1.VersionVector
	(*v11.PairingToken)(nil),           // 7: strims.auth.v1.PairingToken
}
var file_replication_v1_replication_proto_depIdxs = []int32{
	6, // 0: strims.replication.v1.Version.current_version:type_name -> strims.dao.v1.VersionVector
	6, // 1: strims.replication.v1.Event.version:type_name -> strims.dao.v1.VersionVector
	1, // 2: strims.replication.v1.EventBundle.events:type_name -> strims.replication.v1.Event
	1, // 3: strims.replication.v1.EventLog.events:type_name -> strims.replication.v1.Event
	7, // 4: strims.replication.v1.CreatePairingTokenResponse.token:type_name -> strims.auth.v1.PairingToken
	4, // 5: strims.replication.v1.ReplicationFrontend.CreatePairingToken:input_type -> strims.replication.v1.CreatePairingTokenRequest
	5, // 6: strims.replication.v1.ReplicationFrontend.CreatePairingToken:output_type -> strims.replication.v1.CreatePairingTokenResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_replication_v1_replication_proto_init() }
func file_replication_v1_replication_proto_init() {
	if File_replication_v1_replication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_replication_v1_replication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_replication_v1_replication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_replication_v1_replication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBundle); i {
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
		file_replication_v1_replication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventLog); i {
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
		file_replication_v1_replication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePairingTokenRequest); i {
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
		file_replication_v1_replication_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePairingTokenResponse); i {
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
			RawDescriptor: file_replication_v1_replication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_replication_v1_replication_proto_goTypes,
		DependencyIndexes: file_replication_v1_replication_proto_depIdxs,
		MessageInfos:      file_replication_v1_replication_proto_msgTypes,
	}.Build()
	File_replication_v1_replication_proto = out.File
	file_replication_v1_replication_proto_rawDesc = nil
	file_replication_v1_replication_proto_goTypes = nil
	file_replication_v1_replication_proto_depIdxs = nil
}

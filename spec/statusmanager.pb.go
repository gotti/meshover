// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/statusmanager.proto

package spec

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type BGPStates int32

const (
	BGPStates_BGPStateUnknown     BGPStates = 0
	BGPStates_BGPStateIdle        BGPStates = 1
	BGPStates_BGPStateConnect     BGPStates = 2
	BGPStates_BGPStateActive      BGPStates = 3
	BGPStates_BGPStateEstablished BGPStates = 4
)

// Enum value maps for BGPStates.
var (
	BGPStates_name = map[int32]string{
		0: "BGPStateUnknown",
		1: "BGPStateIdle",
		2: "BGPStateConnect",
		3: "BGPStateActive",
		4: "BGPStateEstablished",
	}
	BGPStates_value = map[string]int32{
		"BGPStateUnknown":     0,
		"BGPStateIdle":        1,
		"BGPStateConnect":     2,
		"BGPStateActive":      3,
		"BGPStateEstablished": 4,
	}
)

func (x BGPStates) Enum() *BGPStates {
	p := new(BGPStates)
	*p = x
	return p
}

func (x BGPStates) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BGPStates) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_statusmanager_proto_enumTypes[0].Descriptor()
}

func (BGPStates) Type() protoreflect.EnumType {
	return &file_proto_statusmanager_proto_enumTypes[0]
}

func (x BGPStates) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BGPStates.Descriptor instead.
func (BGPStates) EnumDescriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{0}
}

type MinimumNodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalAS   *ASN           `protobuf:"bytes,1,opt,name=LocalAS,proto3" json:"LocalAS,omitempty"`
	Addresses []*AddressCIDR `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Endpoint  *Address       `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *MinimumNodeStatus) Reset() {
	*x = MinimumNodeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinimumNodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinimumNodeStatus) ProtoMessage() {}

func (x *MinimumNodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinimumNodeStatus.ProtoReflect.Descriptor instead.
func (*MinimumNodeStatus) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{0}
}

func (x *MinimumNodeStatus) GetLocalAS() *ASN {
	if x != nil {
		return x.LocalAS
	}
	return nil
}

func (x *MinimumNodeStatus) GetAddresses() []*AddressCIDR {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *MinimumNodeStatus) GetEndpoint() *Address {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

type PeerWireguardStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestHandshake int64   `protobuf:"varint,1,opt,name=latestHandshake,proto3" json:"latestHandshake,omitempty"`
	TxBytes         float64 `protobuf:"fixed64,2,opt,name=txBytes,proto3" json:"txBytes,omitempty"`
	RxBytes         float64 `protobuf:"fixed64,3,opt,name=rxBytes,proto3" json:"rxBytes,omitempty"`
}

func (x *PeerWireguardStatus) Reset() {
	*x = PeerWireguardStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerWireguardStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerWireguardStatus) ProtoMessage() {}

func (x *PeerWireguardStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerWireguardStatus.ProtoReflect.Descriptor instead.
func (*PeerWireguardStatus) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{1}
}

func (x *PeerWireguardStatus) GetLatestHandshake() int64 {
	if x != nil {
		return x.LatestHandshake
	}
	return 0
}

func (x *PeerWireguardStatus) GetTxBytes() float64 {
	if x != nil {
		return x.TxBytes
	}
	return 0
}

func (x *PeerWireguardStatus) GetRxBytes() float64 {
	if x != nil {
		return x.RxBytes
	}
	return 0
}

type PeerBGPStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteHostname  string    `protobuf:"bytes,1,opt,name=remoteHostname,proto3" json:"remoteHostname,omitempty"`
	LocalAS         *ASN      `protobuf:"bytes,2,opt,name=LocalAS,proto3" json:"LocalAS,omitempty"`
	RemoteAS        *ASN      `protobuf:"bytes,3,opt,name=RemoteAS,proto3" json:"RemoteAS,omitempty"`
	BgpNeighborAddr *Address  `protobuf:"bytes,4,opt,name=BgpNeighborAddr,proto3" json:"BgpNeighborAddr,omitempty"`
	BGPState        BGPStates `protobuf:"varint,6,opt,name=BGPState,proto3,enum=BGPStates" json:"BGPState,omitempty"`
}

func (x *PeerBGPStatus) Reset() {
	*x = PeerBGPStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerBGPStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerBGPStatus) ProtoMessage() {}

func (x *PeerBGPStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerBGPStatus.ProtoReflect.Descriptor instead.
func (*PeerBGPStatus) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{2}
}

func (x *PeerBGPStatus) GetRemoteHostname() string {
	if x != nil {
		return x.RemoteHostname
	}
	return ""
}

func (x *PeerBGPStatus) GetLocalAS() *ASN {
	if x != nil {
		return x.LocalAS
	}
	return nil
}

func (x *PeerBGPStatus) GetRemoteAS() *ASN {
	if x != nil {
		return x.RemoteAS
	}
	return nil
}

func (x *PeerBGPStatus) GetBgpNeighborAddr() *Address {
	if x != nil {
		return x.BgpNeighborAddr
	}
	return nil
}

func (x *PeerBGPStatus) GetBGPState() BGPStates {
	if x != nil {
		return x.BGPState
	}
	return BGPStates_BGPStateUnknown
}

type PeerPingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AverageRTT float64 `protobuf:"fixed64,1,opt,name=averageRTT,proto3" json:"averageRTT,omitempty"`
}

func (x *PeerPingStatus) Reset() {
	*x = PeerPingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerPingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerPingStatus) ProtoMessage() {}

func (x *PeerPingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerPingStatus.ProtoReflect.Descriptor instead.
func (*PeerPingStatus) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{3}
}

func (x *PeerPingStatus) GetAverageRTT() float64 {
	if x != nil {
		return x.AverageRTT
	}
	return 0
}

type NodePeersStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteHostname  string               `protobuf:"bytes,1,opt,name=remoteHostname,proto3" json:"remoteHostname,omitempty"`
	PingStatus      *PeerPingStatus      `protobuf:"bytes,2,opt,name=pingStatus,proto3" json:"pingStatus,omitempty"`
	WireguardStatus *PeerWireguardStatus `protobuf:"bytes,3,opt,name=wireguardStatus,proto3" json:"wireguardStatus,omitempty"`
	BgpStatus       *PeerBGPStatus       `protobuf:"bytes,4,opt,name=bgpStatus,proto3" json:"bgpStatus,omitempty"`
}

func (x *NodePeersStatus) Reset() {
	*x = NodePeersStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePeersStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePeersStatus) ProtoMessage() {}

func (x *NodePeersStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePeersStatus.ProtoReflect.Descriptor instead.
func (*NodePeersStatus) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{4}
}

func (x *NodePeersStatus) GetRemoteHostname() string {
	if x != nil {
		return x.RemoteHostname
	}
	return ""
}

func (x *NodePeersStatus) GetPingStatus() *PeerPingStatus {
	if x != nil {
		return x.PingStatus
	}
	return nil
}

func (x *NodePeersStatus) GetWireguardStatus() *PeerWireguardStatus {
	if x != nil {
		return x.WireguardStatus
	}
	return nil
}

func (x *NodePeersStatus) GetBgpStatus() *PeerBGPStatus {
	if x != nil {
		return x.BgpStatus
	}
	return nil
}

type StatusManagerPeerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname   string             `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	NodeStatus *MinimumNodeStatus `protobuf:"bytes,2,opt,name=nodeStatus,proto3" json:"nodeStatus,omitempty"`
	PeerStatus []*NodePeersStatus `protobuf:"bytes,3,rep,name=peerStatus,proto3" json:"peerStatus,omitempty"`
}

func (x *StatusManagerPeerStatus) Reset() {
	*x = StatusManagerPeerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusManagerPeerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusManagerPeerStatus) ProtoMessage() {}

func (x *StatusManagerPeerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusManagerPeerStatus.ProtoReflect.Descriptor instead.
func (*StatusManagerPeerStatus) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{5}
}

func (x *StatusManagerPeerStatus) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *StatusManagerPeerStatus) GetNodeStatus() *MinimumNodeStatus {
	if x != nil {
		return x.NodeStatus
	}
	return nil
}

func (x *StatusManagerPeerStatus) GetPeerStatus() []*NodePeersStatus {
	if x != nil {
		return x.PeerStatus
	}
	return nil
}

type RegisterStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *StatusManagerPeerStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RegisterStatusRequest) Reset() {
	*x = RegisterStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterStatusRequest) ProtoMessage() {}

func (x *RegisterStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterStatusRequest.ProtoReflect.Descriptor instead.
func (*RegisterStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterStatusRequest) GetStatus() *StatusManagerPeerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type RegisterStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterStatusResponse) Reset() {
	*x = RegisterStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statusmanager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterStatusResponse) ProtoMessage() {}

func (x *RegisterStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statusmanager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterStatusResponse.ProtoReflect.Descriptor instead.
func (*RegisterStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_statusmanager_proto_rawDescGZIP(), []int{7}
}

var File_proto_statusmanager_proto protoreflect.FileDescriptor

var file_proto_statusmanager_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x4d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x07, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x04, 0x2e, 0x41, 0x53, 0x4e, 0x52, 0x07, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x53, 0x12, 0x2a,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x49, 0x44, 0x52, 0x52,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x13, 0x50, 0x65,
	0x65, 0x72, 0x57, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x74, 0x78,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22,
	0xd5, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x41, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x41, 0x53, 0x4e,
	0x52, 0x07, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x53, 0x12, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x41, 0x53,
	0x4e, 0x52, 0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x53, 0x12, 0x32, 0x0a, 0x0f, 0x42,
	0x67, 0x70, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0f,
	0x42, 0x67, 0x70, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x26, 0x0a, 0x08, 0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x08, 0x42,
	0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x30, 0x0a, 0x0e, 0x50, 0x65, 0x65, 0x72, 0x50,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x54, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x54, 0x54, 0x22, 0xd8, 0x01, 0x0a, 0x0f, 0x4e, 0x6f,
	0x64, 0x65, 0x50, 0x65, 0x65, 0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x50, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x70, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0f, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x57, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x62, 0x67, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x62, 0x67, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x32, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x13, 0x72, 0x11, 0x32, 0x0f, 0x5e, 0x5b, 0x30, 0x2d, 0x39,
	0x61, 0x2d, 0x66, 0x41, 0x2d, 0x5a, 0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x65, 0x65,
	0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x65, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2a, 0x74, 0x0a, 0x09, 0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x13, 0x0a, 0x0f, 0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x47, 0x50, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x03,
	0x12, 0x17, 0x0a, 0x13, 0x42, 0x47, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x73, 0x74, 0x61,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x10, 0x04, 0x32, 0x59, 0x0a, 0x14, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x41, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_statusmanager_proto_rawDescOnce sync.Once
	file_proto_statusmanager_proto_rawDescData = file_proto_statusmanager_proto_rawDesc
)

func file_proto_statusmanager_proto_rawDescGZIP() []byte {
	file_proto_statusmanager_proto_rawDescOnce.Do(func() {
		file_proto_statusmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_statusmanager_proto_rawDescData)
	})
	return file_proto_statusmanager_proto_rawDescData
}

var file_proto_statusmanager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_statusmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_statusmanager_proto_goTypes = []interface{}{
	(BGPStates)(0),                  // 0: BGPStates
	(*MinimumNodeStatus)(nil),       // 1: MinimumNodeStatus
	(*PeerWireguardStatus)(nil),     // 2: PeerWireguardStatus
	(*PeerBGPStatus)(nil),           // 3: PeerBGPStatus
	(*PeerPingStatus)(nil),          // 4: PeerPingStatus
	(*NodePeersStatus)(nil),         // 5: NodePeersStatus
	(*StatusManagerPeerStatus)(nil), // 6: StatusManagerPeerStatus
	(*RegisterStatusRequest)(nil),   // 7: RegisterStatusRequest
	(*RegisterStatusResponse)(nil),  // 8: RegisterStatusResponse
	(*ASN)(nil),                     // 9: ASN
	(*AddressCIDR)(nil),             // 10: AddressCIDR
	(*Address)(nil),                 // 11: Address
}
var file_proto_statusmanager_proto_depIdxs = []int32{
	9,  // 0: MinimumNodeStatus.LocalAS:type_name -> ASN
	10, // 1: MinimumNodeStatus.addresses:type_name -> AddressCIDR
	11, // 2: MinimumNodeStatus.endpoint:type_name -> Address
	9,  // 3: PeerBGPStatus.LocalAS:type_name -> ASN
	9,  // 4: PeerBGPStatus.RemoteAS:type_name -> ASN
	11, // 5: PeerBGPStatus.BgpNeighborAddr:type_name -> Address
	0,  // 6: PeerBGPStatus.BGPState:type_name -> BGPStates
	4,  // 7: NodePeersStatus.pingStatus:type_name -> PeerPingStatus
	2,  // 8: NodePeersStatus.wireguardStatus:type_name -> PeerWireguardStatus
	3,  // 9: NodePeersStatus.bgpStatus:type_name -> PeerBGPStatus
	1,  // 10: StatusManagerPeerStatus.nodeStatus:type_name -> MinimumNodeStatus
	5,  // 11: StatusManagerPeerStatus.peerStatus:type_name -> NodePeersStatus
	6,  // 12: RegisterStatusRequest.status:type_name -> StatusManagerPeerStatus
	7,  // 13: StatusManagerService.RegisterStatus:input_type -> RegisterStatusRequest
	8,  // 14: StatusManagerService.RegisterStatus:output_type -> RegisterStatusResponse
	14, // [14:15] is the sub-list for method output_type
	13, // [13:14] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_statusmanager_proto_init() }
func file_proto_statusmanager_proto_init() {
	if File_proto_statusmanager_proto != nil {
		return
	}
	file_proto_ip_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_statusmanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinimumNodeStatus); i {
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
		file_proto_statusmanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerWireguardStatus); i {
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
		file_proto_statusmanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerBGPStatus); i {
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
		file_proto_statusmanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerPingStatus); i {
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
		file_proto_statusmanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePeersStatus); i {
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
		file_proto_statusmanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusManagerPeerStatus); i {
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
		file_proto_statusmanager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterStatusRequest); i {
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
		file_proto_statusmanager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterStatusResponse); i {
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
			RawDescriptor: file_proto_statusmanager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_statusmanager_proto_goTypes,
		DependencyIndexes: file_proto_statusmanager_proto_depIdxs,
		EnumInfos:         file_proto_statusmanager_proto_enumTypes,
		MessageInfos:      file_proto_statusmanager_proto_msgTypes,
	}.Build()
	File_proto_statusmanager_proto = out.File
	file_proto_statusmanager_proto_rawDesc = nil
	file_proto_statusmanager_proto_goTypes = nil
	file_proto_statusmanager_proto_depIdxs = nil
}

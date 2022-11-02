// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/administration.proto

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

type AgentKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *AgentKey) Reset() {
	*x = AgentKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentKey) ProtoMessage() {}

func (x *AgentKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentKey.ProtoReflect.Descriptor instead.
func (*AgentKey) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{0}
}

func (x *AgentKey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AgentKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type AgentKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*AgentKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *AgentKeys) Reset() {
	*x = AgentKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentKeys) ProtoMessage() {}

func (x *AgentKeys) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentKeys.ProtoReflect.Descriptor instead.
func (*AgentKeys) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{1}
}

func (x *AgentKeys) GetKeys() []*AgentKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type GenerateAgentKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateAgentKeyRequest) Reset() {
	*x = GenerateAgentKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAgentKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAgentKeyRequest) ProtoMessage() {}

func (x *GenerateAgentKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAgentKeyRequest.ProtoReflect.Descriptor instead.
func (*GenerateAgentKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{2}
}

type GenerateAgentKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentKey *AgentKey `protobuf:"bytes,1,opt,name=agentKey,proto3" json:"agentKey,omitempty"`
}

func (x *GenerateAgentKeyResponse) Reset() {
	*x = GenerateAgentKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAgentKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAgentKeyResponse) ProtoMessage() {}

func (x *GenerateAgentKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAgentKeyResponse.ProtoReflect.Descriptor instead.
func (*GenerateAgentKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateAgentKeyResponse) GetAgentKey() *AgentKey {
	if x != nil {
		return x.AgentKey
	}
	return nil
}

type ListAgentKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAgentKeyRequest) Reset() {
	*x = ListAgentKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAgentKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentKeyRequest) ProtoMessage() {}

func (x *ListAgentKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentKeyRequest.ProtoReflect.Descriptor instead.
func (*ListAgentKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{4}
}

type ListAgentKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys *AgentKeys `protobuf:"bytes,1,opt,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ListAgentKeyResponse) Reset() {
	*x = ListAgentKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAgentKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentKeyResponse) ProtoMessage() {}

func (x *ListAgentKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentKeyResponse.ProtoReflect.Descriptor instead.
func (*ListAgentKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{5}
}

func (x *ListAgentKeyResponse) GetKeys() *AgentKeys {
	if x != nil {
		return x.Keys
	}
	return nil
}

type RevokeAgentKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RevokeAgentKeyRequest) Reset() {
	*x = RevokeAgentKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeAgentKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeAgentKeyRequest) ProtoMessage() {}

func (x *RevokeAgentKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeAgentKeyRequest.ProtoReflect.Descriptor instead.
func (*RevokeAgentKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{6}
}

func (x *RevokeAgentKeyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RevokeAgentKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RevokeAgentKeyResponse) Reset() {
	*x = RevokeAgentKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_administration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeAgentKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeAgentKeyResponse) ProtoMessage() {}

func (x *RevokeAgentKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_administration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeAgentKeyResponse.ProtoReflect.Descriptor instead.
func (*RevokeAgentKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_administration_proto_rawDescGZIP(), []int{7}
}

func (x *RevokeAgentKeyResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *RevokeAgentKeyResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_proto_administration_proto protoreflect.FileDescriptor

var file_proto_administration_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x08, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2a, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41,
	0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65,
	0x79, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x22, 0x27, 0x0a, 0x15, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x16, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xdf, 0x01, 0x0a, 0x14,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_administration_proto_rawDescOnce sync.Once
	file_proto_administration_proto_rawDescData = file_proto_administration_proto_rawDesc
)

func file_proto_administration_proto_rawDescGZIP() []byte {
	file_proto_administration_proto_rawDescOnce.Do(func() {
		file_proto_administration_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_administration_proto_rawDescData)
	})
	return file_proto_administration_proto_rawDescData
}

var file_proto_administration_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_administration_proto_goTypes = []interface{}{
	(*AgentKey)(nil),                 // 0: AgentKey
	(*AgentKeys)(nil),                // 1: AgentKeys
	(*GenerateAgentKeyRequest)(nil),  // 2: GenerateAgentKeyRequest
	(*GenerateAgentKeyResponse)(nil), // 3: GenerateAgentKeyResponse
	(*ListAgentKeyRequest)(nil),      // 4: ListAgentKeyRequest
	(*ListAgentKeyResponse)(nil),     // 5: ListAgentKeyResponse
	(*RevokeAgentKeyRequest)(nil),    // 6: RevokeAgentKeyRequest
	(*RevokeAgentKeyResponse)(nil),   // 7: RevokeAgentKeyResponse
}
var file_proto_administration_proto_depIdxs = []int32{
	0, // 0: AgentKeys.keys:type_name -> AgentKey
	0, // 1: GenerateAgentKeyResponse.agentKey:type_name -> AgentKey
	1, // 2: ListAgentKeyResponse.keys:type_name -> AgentKeys
	2, // 3: AdministratorService.GenerateAgentKey:input_type -> GenerateAgentKeyRequest
	4, // 4: AdministratorService.ListAgentKey:input_type -> ListAgentKeyRequest
	6, // 5: AdministratorService.RevokeAgentKey:input_type -> RevokeAgentKeyRequest
	3, // 6: AdministratorService.GenerateAgentKey:output_type -> GenerateAgentKeyResponse
	5, // 7: AdministratorService.ListAgentKey:output_type -> ListAgentKeyResponse
	7, // 8: AdministratorService.RevokeAgentKey:output_type -> RevokeAgentKeyResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_administration_proto_init() }
func file_proto_administration_proto_init() {
	if File_proto_administration_proto != nil {
		return
	}
	file_proto_ip_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_administration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentKey); i {
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
		file_proto_administration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentKeys); i {
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
		file_proto_administration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAgentKeyRequest); i {
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
		file_proto_administration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAgentKeyResponse); i {
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
		file_proto_administration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAgentKeyRequest); i {
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
		file_proto_administration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAgentKeyResponse); i {
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
		file_proto_administration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeAgentKeyRequest); i {
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
		file_proto_administration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeAgentKeyResponse); i {
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
			RawDescriptor: file_proto_administration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_administration_proto_goTypes,
		DependencyIndexes: file_proto_administration_proto_depIdxs,
		MessageInfos:      file_proto_administration_proto_msgTypes,
	}.Build()
	File_proto_administration_proto = out.File
	file_proto_administration_proto_rawDesc = nil
	file_proto_administration_proto_goTypes = nil
	file_proto_administration_proto_depIdxs = nil
}
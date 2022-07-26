// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/ip.proto

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

type AddressIPv4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipaddress string `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
}

func (x *AddressIPv4) Reset() {
	*x = AddressIPv4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressIPv4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressIPv4) ProtoMessage() {}

func (x *AddressIPv4) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressIPv4.ProtoReflect.Descriptor instead.
func (*AddressIPv4) Descriptor() ([]byte, []int) {
	return file_proto_ip_proto_rawDescGZIP(), []int{0}
}

func (x *AddressIPv4) GetIpaddress() string {
	if x != nil {
		return x.Ipaddress
	}
	return ""
}

type AddressIPv6 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipaddress string `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
}

func (x *AddressIPv6) Reset() {
	*x = AddressIPv6{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressIPv6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressIPv6) ProtoMessage() {}

func (x *AddressIPv6) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressIPv6.ProtoReflect.Descriptor instead.
func (*AddressIPv6) Descriptor() ([]byte, []int) {
	return file_proto_ip_proto_rawDescGZIP(), []int{1}
}

func (x *AddressIPv6) GetIpaddress() string {
	if x != nil {
		return x.Ipaddress
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Ipaddress:
	//
	//	*Address_AddressIPv4
	//	*Address_AddressIPv6
	Ipaddress isAddress_Ipaddress `protobuf_oneof:"ipaddress"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_ip_proto_rawDescGZIP(), []int{2}
}

func (m *Address) GetIpaddress() isAddress_Ipaddress {
	if m != nil {
		return m.Ipaddress
	}
	return nil
}

func (x *Address) GetAddressIPv4() *AddressIPv4 {
	if x, ok := x.GetIpaddress().(*Address_AddressIPv4); ok {
		return x.AddressIPv4
	}
	return nil
}

func (x *Address) GetAddressIPv6() *AddressIPv6 {
	if x, ok := x.GetIpaddress().(*Address_AddressIPv6); ok {
		return x.AddressIPv6
	}
	return nil
}

type isAddress_Ipaddress interface {
	isAddress_Ipaddress()
}

type Address_AddressIPv4 struct {
	AddressIPv4 *AddressIPv4 `protobuf:"bytes,2,opt,name=addressIPv4,proto3,oneof"`
}

type Address_AddressIPv6 struct {
	AddressIPv6 *AddressIPv6 `protobuf:"bytes,3,opt,name=addressIPv6,proto3,oneof"`
}

func (*Address_AddressIPv4) isAddress_Ipaddress() {}

func (*Address_AddressIPv6) isAddress_Ipaddress() {}

type AddressCIDRIPv4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipaddress *AddressIPv4 `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Mask      int32        `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *AddressCIDRIPv4) Reset() {
	*x = AddressCIDRIPv4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressCIDRIPv4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressCIDRIPv4) ProtoMessage() {}

func (x *AddressCIDRIPv4) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressCIDRIPv4.ProtoReflect.Descriptor instead.
func (*AddressCIDRIPv4) Descriptor() ([]byte, []int) {
	return file_proto_ip_proto_rawDescGZIP(), []int{3}
}

func (x *AddressCIDRIPv4) GetIpaddress() *AddressIPv4 {
	if x != nil {
		return x.Ipaddress
	}
	return nil
}

func (x *AddressCIDRIPv4) GetMask() int32 {
	if x != nil {
		return x.Mask
	}
	return 0
}

type AddressCIDRIPv6 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipaddress *AddressIPv6 `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Mask      int32        `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *AddressCIDRIPv6) Reset() {
	*x = AddressCIDRIPv6{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressCIDRIPv6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressCIDRIPv6) ProtoMessage() {}

func (x *AddressCIDRIPv6) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressCIDRIPv6.ProtoReflect.Descriptor instead.
func (*AddressCIDRIPv6) Descriptor() ([]byte, []int) {
	return file_proto_ip_proto_rawDescGZIP(), []int{4}
}

func (x *AddressCIDRIPv6) GetIpaddress() *AddressIPv6 {
	if x != nil {
		return x.Ipaddress
	}
	return nil
}

func (x *AddressCIDRIPv6) GetMask() int32 {
	if x != nil {
		return x.Mask
	}
	return 0
}

type AddressCIDR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Addresscidr:
	//
	//	*AddressCIDR_AddressCIDRIPv4
	//	*AddressCIDR_AddressCIDRIPv6
	Addresscidr isAddressCIDR_Addresscidr `protobuf_oneof:"addresscidr"`
}

func (x *AddressCIDR) Reset() {
	*x = AddressCIDR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressCIDR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressCIDR) ProtoMessage() {}

func (x *AddressCIDR) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressCIDR.ProtoReflect.Descriptor instead.
func (*AddressCIDR) Descriptor() ([]byte, []int) {
	return file_proto_ip_proto_rawDescGZIP(), []int{5}
}

func (m *AddressCIDR) GetAddresscidr() isAddressCIDR_Addresscidr {
	if m != nil {
		return m.Addresscidr
	}
	return nil
}

func (x *AddressCIDR) GetAddressCIDRIPv4() *AddressCIDRIPv4 {
	if x, ok := x.GetAddresscidr().(*AddressCIDR_AddressCIDRIPv4); ok {
		return x.AddressCIDRIPv4
	}
	return nil
}

func (x *AddressCIDR) GetAddressCIDRIPv6() *AddressCIDRIPv6 {
	if x, ok := x.GetAddresscidr().(*AddressCIDR_AddressCIDRIPv6); ok {
		return x.AddressCIDRIPv6
	}
	return nil
}

type isAddressCIDR_Addresscidr interface {
	isAddressCIDR_Addresscidr()
}

type AddressCIDR_AddressCIDRIPv4 struct {
	AddressCIDRIPv4 *AddressCIDRIPv4 `protobuf:"bytes,2,opt,name=addressCIDRIPv4,proto3,oneof"`
}

type AddressCIDR_AddressCIDRIPv6 struct {
	AddressCIDRIPv6 *AddressCIDRIPv6 `protobuf:"bytes,3,opt,name=addressCIDRIPv6,proto3,oneof"`
}

func (*AddressCIDR_AddressCIDRIPv4) isAddressCIDR_Addresscidr() {}

func (*AddressCIDR_AddressCIDRIPv6) isAddressCIDR_Addresscidr() {}

type ASN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ASN) Reset() {
	*x = ASN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ASN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ASN) ProtoMessage() {}

func (x *ASN) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ASN.ProtoReflect.Descriptor instead.
func (*ASN) Descriptor() ([]byte, []int) {
	return file_proto_ip_proto_rawDescGZIP(), []int{6}
}

func (x *ASN) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_proto_ip_proto protoreflect.FileDescriptor

var file_proto_ip_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x34, 0x12, 0x25, 0x0a, 0x09,
	0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x78, 0x01, 0x52, 0x09, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x35, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x50,
	0x76, 0x36, 0x12, 0x26, 0x0a, 0x09, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x80, 0x01, 0x01, 0x52,
	0x09, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x7f, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x50, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x34, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x34, 0x12, 0x30, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x50, 0x76, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x36, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x36, 0x42, 0x10, 0x0a, 0x09, 0x69, 0x70, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x66, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x49, 0x44, 0x52, 0x49, 0x50, 0x76, 0x34, 0x12, 0x34,
	0x0a, 0x09, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x34, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x69, 0x70, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x20, 0x28, 0x00, 0x52, 0x04, 0x6d,
	0x61, 0x73, 0x6b, 0x22, 0x67, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x49,
	0x44, 0x52, 0x49, 0x50, 0x76, 0x36, 0x12, 0x34, 0x0a, 0x09, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x36, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x04,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x1a,
	0x05, 0x18, 0x80, 0x01, 0x28, 0x00, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x9d, 0x01, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x49, 0x44, 0x52, 0x12, 0x3c, 0x0a, 0x0f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x49, 0x44, 0x52, 0x49, 0x50, 0x76, 0x34, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43,
	0x49, 0x44, 0x52, 0x49, 0x50, 0x76, 0x34, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x43, 0x49, 0x44, 0x52, 0x49, 0x50, 0x76, 0x34, 0x12, 0x3c, 0x0a, 0x0f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x49, 0x44, 0x52, 0x49, 0x50, 0x76, 0x36, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x49, 0x44,
	0x52, 0x49, 0x50, 0x76, 0x36, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x43, 0x49, 0x44, 0x52, 0x49, 0x50, 0x76, 0x36, 0x42, 0x12, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x63, 0x69, 0x64, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x1d, 0x0a, 0x03,
	0x41, 0x53, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ip_proto_rawDescOnce sync.Once
	file_proto_ip_proto_rawDescData = file_proto_ip_proto_rawDesc
)

func file_proto_ip_proto_rawDescGZIP() []byte {
	file_proto_ip_proto_rawDescOnce.Do(func() {
		file_proto_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ip_proto_rawDescData)
	})
	return file_proto_ip_proto_rawDescData
}

var file_proto_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_ip_proto_goTypes = []interface{}{
	(*AddressIPv4)(nil),     // 0: AddressIPv4
	(*AddressIPv6)(nil),     // 1: AddressIPv6
	(*Address)(nil),         // 2: Address
	(*AddressCIDRIPv4)(nil), // 3: AddressCIDRIPv4
	(*AddressCIDRIPv6)(nil), // 4: AddressCIDRIPv6
	(*AddressCIDR)(nil),     // 5: AddressCIDR
	(*ASN)(nil),             // 6: ASN
}
var file_proto_ip_proto_depIdxs = []int32{
	0, // 0: Address.addressIPv4:type_name -> AddressIPv4
	1, // 1: Address.addressIPv6:type_name -> AddressIPv6
	0, // 2: AddressCIDRIPv4.ipaddress:type_name -> AddressIPv4
	1, // 3: AddressCIDRIPv6.ipaddress:type_name -> AddressIPv6
	3, // 4: AddressCIDR.addressCIDRIPv4:type_name -> AddressCIDRIPv4
	4, // 5: AddressCIDR.addressCIDRIPv6:type_name -> AddressCIDRIPv6
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_ip_proto_init() }
func file_proto_ip_proto_init() {
	if File_proto_ip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressIPv4); i {
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
		file_proto_ip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressIPv6); i {
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
		file_proto_ip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_proto_ip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressCIDRIPv4); i {
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
		file_proto_ip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressCIDRIPv6); i {
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
		file_proto_ip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressCIDR); i {
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
		file_proto_ip_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ASN); i {
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
	file_proto_ip_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Address_AddressIPv4)(nil),
		(*Address_AddressIPv6)(nil),
	}
	file_proto_ip_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*AddressCIDR_AddressCIDRIPv4)(nil),
		(*AddressCIDR_AddressCIDRIPv6)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ip_proto_goTypes,
		DependencyIndexes: file_proto_ip_proto_depIdxs,
		MessageInfos:      file_proto_ip_proto_msgTypes,
	}.Build()
	File_proto_ip_proto = out.File
	file_proto_ip_proto_rawDesc = nil
	file_proto_ip_proto_goTypes = nil
	file_proto_ip_proto_depIdxs = nil
}

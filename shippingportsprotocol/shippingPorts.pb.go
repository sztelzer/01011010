// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: shippingportsprotocol/shippingports.proto

package shippingportsprotocol

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

// message ShippingPort represents one Shipping Port, with all its properties
type ShippingPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string    `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ShippingPort) Reset() {
	*x = ShippingPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingPort) ProtoMessage() {}

func (x *ShippingPort) ProtoReflect() protoreflect.Message {
	mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingPort.ProtoReflect.Descriptor instead.
func (*ShippingPort) Descriptor() ([]byte, []int) {
	return file_shippingportsprotocol_shippingports_proto_rawDescGZIP(), []int{0}
}

func (x *ShippingPort) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShippingPort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShippingPort) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ShippingPort) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ShippingPort) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *ShippingPort) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *ShippingPort) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *ShippingPort) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *ShippingPort) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *ShippingPort) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *ShippingPort) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// message ShippingPortId transports just one id, used on the Get method
type ShippingPortId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShippingPortId) Reset() {
	*x = ShippingPortId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingPortId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingPortId) ProtoMessage() {}

func (x *ShippingPortId) ProtoReflect() protoreflect.Message {
	mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingPortId.ProtoReflect.Descriptor instead.
func (*ShippingPortId) Descriptor() ([]byte, []int) {
	return file_shippingportsprotocol_shippingports_proto_rawDescGZIP(), []int{1}
}

func (x *ShippingPortId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// message Ok is a stub for when we need to signal some empty response
type Ok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ok) Reset() {
	*x = Ok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ok.ProtoReflect.Descriptor instead.
func (*Ok) Descriptor() ([]byte, []int) {
	return file_shippingportsprotocol_shippingports_proto_rawDescGZIP(), []int{2}
}

// message Pagination defines offset and size of GetMany requests
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_shippingportsprotocol_shippingports_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Pagination) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// message ManyShippingPorts is an array of ShippingPorts for GetMany requests
type ManyShippingPorts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shippingports []*ShippingPort `protobuf:"bytes,1,rep,name=shippingports,proto3" json:"shippingports,omitempty"`
}

func (x *ManyShippingPorts) Reset() {
	*x = ManyShippingPorts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyShippingPorts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyShippingPorts) ProtoMessage() {}

func (x *ManyShippingPorts) ProtoReflect() protoreflect.Message {
	mi := &file_shippingportsprotocol_shippingports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyShippingPorts.ProtoReflect.Descriptor instead.
func (*ManyShippingPorts) Descriptor() ([]byte, []int) {
	return file_shippingportsprotocol_shippingports_proto_rawDescGZIP(), []int{4}
}

func (x *ManyShippingPorts) GetShippingports() []*ShippingPort {
	if x != nil {
		return x.Shippingports
	}
	return nil
}

var File_shippingportsprotocol_shippingports_proto protoreflect.FileDescriptor

var file_shippingportsprotocol_shippingports_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x22, 0x96, 0x02, 0x0a, 0x0c, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x04, 0x0a,
	0x02, 0x4f, 0x6b, 0x22, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x5e, 0x0a,
	0x11, 0x4d, 0x61, 0x6e, 0x79, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x49, 0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x0d,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x32, 0x8d, 0x02,
	0x0a, 0x13, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x23, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x53,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x1a, 0x23, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72,
	0x74, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x21,
	0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x28, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x00, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x7a, 0x74, 0x65,
	0x6c, 0x7a, 0x65, 0x72, 0x2f, 0x30, 0x31, 0x30, 0x31, 0x31, 0x30, 0x31, 0x30, 0x2f, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shippingportsprotocol_shippingports_proto_rawDescOnce sync.Once
	file_shippingportsprotocol_shippingports_proto_rawDescData = file_shippingportsprotocol_shippingports_proto_rawDesc
)

func file_shippingportsprotocol_shippingports_proto_rawDescGZIP() []byte {
	file_shippingportsprotocol_shippingports_proto_rawDescOnce.Do(func() {
		file_shippingportsprotocol_shippingports_proto_rawDescData = protoimpl.X.CompressGZIP(file_shippingportsprotocol_shippingports_proto_rawDescData)
	})
	return file_shippingportsprotocol_shippingports_proto_rawDescData
}

var file_shippingportsprotocol_shippingports_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shippingportsprotocol_shippingports_proto_goTypes = []interface{}{
	(*ShippingPort)(nil),      // 0: shippingportsprotocol.ShippingPort
	(*ShippingPortId)(nil),    // 1: shippingportsprotocol.ShippingPortId
	(*Ok)(nil),                // 2: shippingportsprotocol.Ok
	(*Pagination)(nil),        // 3: shippingportsprotocol.Pagination
	(*ManyShippingPorts)(nil), // 4: shippingportsprotocol.ManyShippingPorts
}
var file_shippingportsprotocol_shippingports_proto_depIdxs = []int32{
	0, // 0: shippingportsprotocol.ManyShippingPorts.shippingports:type_name -> shippingportsprotocol.ShippingPort
	0, // 1: shippingportsprotocol.ShippingPortsServer.Put:input_type -> shippingportsprotocol.ShippingPort
	1, // 2: shippingportsprotocol.ShippingPortsServer.Get:input_type -> shippingportsprotocol.ShippingPortId
	3, // 3: shippingportsprotocol.ShippingPortsServer.GetMany:input_type -> shippingportsprotocol.Pagination
	2, // 4: shippingportsprotocol.ShippingPortsServer.Put:output_type -> shippingportsprotocol.Ok
	0, // 5: shippingportsprotocol.ShippingPortsServer.Get:output_type -> shippingportsprotocol.ShippingPort
	4, // 6: shippingportsprotocol.ShippingPortsServer.GetMany:output_type -> shippingportsprotocol.ManyShippingPorts
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shippingportsprotocol_shippingports_proto_init() }
func file_shippingportsprotocol_shippingports_proto_init() {
	if File_shippingportsprotocol_shippingports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shippingportsprotocol_shippingports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingPort); i {
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
		file_shippingportsprotocol_shippingports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingPortId); i {
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
		file_shippingportsprotocol_shippingports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ok); i {
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
		file_shippingportsprotocol_shippingports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_shippingportsprotocol_shippingports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyShippingPorts); i {
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
			RawDescriptor: file_shippingportsprotocol_shippingports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shippingportsprotocol_shippingports_proto_goTypes,
		DependencyIndexes: file_shippingportsprotocol_shippingports_proto_depIdxs,
		MessageInfos:      file_shippingportsprotocol_shippingports_proto_msgTypes,
	}.Build()
	File_shippingportsprotocol_shippingports_proto = out.File
	file_shippingportsprotocol_shippingports_proto_rawDesc = nil
	file_shippingportsprotocol_shippingports_proto_goTypes = nil
	file_shippingportsprotocol_shippingports_proto_depIdxs = nil
}

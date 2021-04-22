// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: gRPCShippingPorts/shippingPorts.proto

package shippingportsproto

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
		mi := &file_gRPCShippingPorts_shippingPorts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingPort) ProtoMessage() {}

func (x *ShippingPort) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCShippingPorts_shippingPorts_proto_msgTypes[0]
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
	return file_gRPCShippingPorts_shippingPorts_proto_rawDescGZIP(), []int{0}
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
		mi := &file_gRPCShippingPorts_shippingPorts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingPortId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingPortId) ProtoMessage() {}

func (x *ShippingPortId) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCShippingPorts_shippingPorts_proto_msgTypes[1]
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
	return file_gRPCShippingPorts_shippingPorts_proto_rawDescGZIP(), []int{1}
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
		mi := &file_gRPCShippingPorts_shippingPorts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCShippingPorts_shippingPorts_proto_msgTypes[2]
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
	return file_gRPCShippingPorts_shippingPorts_proto_rawDescGZIP(), []int{2}
}

var File_gRPCShippingPorts_shippingPorts_proto protoreflect.FileDescriptor

var file_gRPCShippingPorts_shippingPorts_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x52, 0x50, 0x43, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x0c,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x04, 0x0a, 0x02, 0x4f, 0x6b, 0x32, 0xa1, 0x01, 0x0a,
	0x0d, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x41,
	0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x20, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6b, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x1a, 0x20, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x00,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x7a, 0x74, 0x65, 0x6c, 0x7a, 0x65, 0x72, 0x2f, 0x30, 0x31, 0x30, 0x31, 0x31, 0x30, 0x31, 0x30,
	0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPCShippingPorts_shippingPorts_proto_rawDescOnce sync.Once
	file_gRPCShippingPorts_shippingPorts_proto_rawDescData = file_gRPCShippingPorts_shippingPorts_proto_rawDesc
)

func file_gRPCShippingPorts_shippingPorts_proto_rawDescGZIP() []byte {
	file_gRPCShippingPorts_shippingPorts_proto_rawDescOnce.Do(func() {
		file_gRPCShippingPorts_shippingPorts_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPCShippingPorts_shippingPorts_proto_rawDescData)
	})
	return file_gRPCShippingPorts_shippingPorts_proto_rawDescData
}

var file_gRPCShippingPorts_shippingPorts_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gRPCShippingPorts_shippingPorts_proto_goTypes = []interface{}{
	(*ShippingPort)(nil),   // 0: shippingportsproto.ShippingPort
	(*ShippingPortId)(nil), // 1: shippingportsproto.ShippingPortId
	(*Ok)(nil),             // 2: shippingportsproto.Ok
}
var file_gRPCShippingPorts_shippingPorts_proto_depIdxs = []int32{
	0, // 0: shippingportsproto.ShippingPorts.Put:input_type -> shippingportsproto.ShippingPort
	1, // 1: shippingportsproto.ShippingPorts.Get:input_type -> shippingportsproto.ShippingPortId
	2, // 2: shippingportsproto.ShippingPorts.Put:output_type -> shippingportsproto.Ok
	0, // 3: shippingportsproto.ShippingPorts.Get:output_type -> shippingportsproto.ShippingPort
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPCShippingPorts_shippingPorts_proto_init() }
func file_gRPCShippingPorts_shippingPorts_proto_init() {
	if File_gRPCShippingPorts_shippingPorts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPCShippingPorts_shippingPorts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_gRPCShippingPorts_shippingPorts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_gRPCShippingPorts_shippingPorts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gRPCShippingPorts_shippingPorts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPCShippingPorts_shippingPorts_proto_goTypes,
		DependencyIndexes: file_gRPCShippingPorts_shippingPorts_proto_depIdxs,
		MessageInfos:      file_gRPCShippingPorts_shippingPorts_proto_msgTypes,
	}.Build()
	File_gRPCShippingPorts_shippingPorts_proto = out.File
	file_gRPCShippingPorts_shippingPorts_proto_rawDesc = nil
	file_gRPCShippingPorts_shippingPorts_proto_goTypes = nil
	file_gRPCShippingPorts_shippingPorts_proto_depIdxs = nil
}

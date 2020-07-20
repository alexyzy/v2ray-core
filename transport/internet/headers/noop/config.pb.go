package noop

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescGZIP(), []int{0}
}

type ConnectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionConfig) Reset() {
	*x = ConnectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig) ProtoMessage() {}

func (x *ConnectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig.ProtoReflect.Descriptor instead.
func (*ConnectionConfig) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescGZIP(), []int{1}
}

var File_v2ray_com_core_transport_internet_headers_noop_config_proto protoreflect.FileDescriptor

var file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x6f, 0x6f, 0x70,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x6f, 0x70, 0x22, 0x08, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x65, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x6f, 0x70, 0x50, 0x01, 0x5a, 0x04, 0x6e, 0x6f, 0x6f,
	0x70, 0xaa, 0x02, 0x2a, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x6f, 0x6f, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescOnce sync.Once
	file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescData = file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDesc
)

func file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescGZIP() []byte {
	file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescOnce.Do(func() {
		file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescData)
	})
	return file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDescData
}

var file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v2ray_com_core_transport_internet_headers_noop_config_proto_goTypes = []interface{}{
	(*Config)(nil),           // 0: v2ray.core.transport.internet.headers.noop.Config
	(*ConnectionConfig)(nil), // 1: v2ray.core.transport.internet.headers.noop.ConnectionConfig
}
var file_v2ray_com_core_transport_internet_headers_noop_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v2ray_com_core_transport_internet_headers_noop_config_proto_init() }
func file_v2ray_com_core_transport_internet_headers_noop_config_proto_init() {
	if File_v2ray_com_core_transport_internet_headers_noop_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig); i {
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
			RawDescriptor: file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2ray_com_core_transport_internet_headers_noop_config_proto_goTypes,
		DependencyIndexes: file_v2ray_com_core_transport_internet_headers_noop_config_proto_depIdxs,
		MessageInfos:      file_v2ray_com_core_transport_internet_headers_noop_config_proto_msgTypes,
	}.Build()
	File_v2ray_com_core_transport_internet_headers_noop_config_proto = out.File
	file_v2ray_com_core_transport_internet_headers_noop_config_proto_rawDesc = nil
	file_v2ray_com_core_transport_internet_headers_noop_config_proto_goTypes = nil
	file_v2ray_com_core_transport_internet_headers_noop_config_proto_depIdxs = nil
}

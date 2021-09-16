// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.4
// source: GrpcServices/TravelersService.proto

package GrpcServices

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcServices_TravelersService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcServices_TravelersService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_GrpcServices_TravelersService_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Location) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_GrpcServices_TravelersService_proto protoreflect.FileDescriptor

var file_GrpcServices_TravelersService_proto_rawDesc = []byte{
	0x0a, 0x23, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x32, 0x5a, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x47, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GrpcServices_TravelersService_proto_rawDescOnce sync.Once
	file_GrpcServices_TravelersService_proto_rawDescData = file_GrpcServices_TravelersService_proto_rawDesc
)

func file_GrpcServices_TravelersService_proto_rawDescGZIP() []byte {
	file_GrpcServices_TravelersService_proto_rawDescOnce.Do(func() {
		file_GrpcServices_TravelersService_proto_rawDescData = protoimpl.X.CompressGZIP(file_GrpcServices_TravelersService_proto_rawDescData)
	})
	return file_GrpcServices_TravelersService_proto_rawDescData
}

var file_GrpcServices_TravelersService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GrpcServices_TravelersService_proto_goTypes = []interface{}{
	(*Location)(nil),    // 0: TravelersService.location
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_GrpcServices_TravelersService_proto_depIdxs = []int32{
	0, // 0: TravelersService.TravelersService.GetNearbyDrivers:input_type -> TravelersService.location
	1, // 1: TravelersService.TravelersService.GetNearbyDrivers:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GrpcServices_TravelersService_proto_init() }
func file_GrpcServices_TravelersService_proto_init() {
	if File_GrpcServices_TravelersService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GrpcServices_TravelersService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
			RawDescriptor: file_GrpcServices_TravelersService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_GrpcServices_TravelersService_proto_goTypes,
		DependencyIndexes: file_GrpcServices_TravelersService_proto_depIdxs,
		MessageInfos:      file_GrpcServices_TravelersService_proto_msgTypes,
	}.Build()
	File_GrpcServices_TravelersService_proto = out.File
	file_GrpcServices_TravelersService_proto_rawDesc = nil
	file_GrpcServices_TravelersService_proto_goTypes = nil
	file_GrpcServices_TravelersService_proto_depIdxs = nil
}

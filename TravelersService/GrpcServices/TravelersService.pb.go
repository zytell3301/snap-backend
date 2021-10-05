// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.4
// source: GrpcServices/TravelersService.proto

package GrpcServices

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

type GetNearbyDriversResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver []*DriverLocation `protobuf:"bytes,1,rep,name=driver,proto3" json:"driver,omitempty"`
}

func (x *GetNearbyDriversResponse) Reset() {
	*x = GetNearbyDriversResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcServices_TravelersService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNearbyDriversResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNearbyDriversResponse) ProtoMessage() {}

func (x *GetNearbyDriversResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetNearbyDriversResponse.ProtoReflect.Descriptor instead.
func (*GetNearbyDriversResponse) Descriptor() ([]byte, []int) {
	return file_GrpcServices_TravelersService_proto_rawDescGZIP(), []int{0}
}

func (x *GetNearbyDriversResponse) GetDriver() []*DriverLocation {
	if x != nil {
		return x.Driver
	}
	return nil
}

type Direction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      *Location `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination *Location `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *Direction) Reset() {
	*x = Direction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcServices_TravelersService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Direction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Direction) ProtoMessage() {}

func (x *Direction) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcServices_TravelersService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Direction.ProtoReflect.Descriptor instead.
func (*Direction) Descriptor() ([]byte, []int) {
	return file_GrpcServices_TravelersService_proto_rawDescGZIP(), []int{1}
}

func (x *Direction) GetOrigin() *Location {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Direction) GetDestination() *Location {
	if x != nil {
		return x.Destination
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price int32 `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcServices_TravelersService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcServices_TravelersService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_GrpcServices_TravelersService_proto_rawDescGZIP(), []int{2}
}

func (x *Price) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

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
		mi := &file_GrpcServices_TravelersService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcServices_TravelersService_proto_msgTypes[3]
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
	return file_GrpcServices_TravelersService_proto_rawDescGZIP(), []int{3}
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

type DriverLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *DriverLocation) Reset() {
	*x = DriverLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcServices_TravelersService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocation) ProtoMessage() {}

func (x *DriverLocation) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcServices_TravelersService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverLocation.ProtoReflect.Descriptor instead.
func (*DriverLocation) Descriptor() ([]byte, []int) {
	return file_GrpcServices_TravelersService_proto_rawDescGZIP(), []int{4}
}

func (x *DriverLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverLocation) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lastname   string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	VehicleNo  string `protobuf:"bytes,3,opt,name=vehicle_no,json=vehicleNo,proto3" json:"vehicle_no,omitempty"`
	Balance    uint32 `protobuf:"varint,4,opt,name=balance,proto3" json:"balance,omitempty"`
	ProfilePic string `protobuf:"bytes,5,opt,name=profile_pic,json=profilePic,proto3" json:"profile_pic,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcServices_TravelersService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcServices_TravelersService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_GrpcServices_TravelersService_proto_rawDescGZIP(), []int{5}
}

func (x *Driver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Driver) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Driver) GetVehicleNo() string {
	if x != nil {
		return x.VehicleNo
	}
	return ""
}

func (x *Driver) GetBalance() uint32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Driver) GetProfilePic() string {
	if x != nil {
		return x.ProfilePic
	}
	return ""
}

var File_GrpcServices_TravelersService_proto protoreflect.FileDescriptor

var file_GrpcServices_TravelersService_proto_rawDesc = []byte{
	0x0a, 0x23, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x61, 0x72, 0x62, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x7d, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x54, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x3c,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x79, 0x22, 0x58, 0x0a, 0x0e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x01,
	0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x69, 0x63, 0x32, 0xb0, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x61, 0x72, 0x62, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x54, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2a, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x61, 0x72, 0x62, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1b, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x47, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_GrpcServices_TravelersService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_GrpcServices_TravelersService_proto_goTypes = []interface{}{
	(*GetNearbyDriversResponse)(nil), // 0: TravelersService.GetNearbyDriversResponse
	(*Direction)(nil),                // 1: TravelersService.direction
	(*Price)(nil),                    // 2: TravelersService.price
	(*Location)(nil),                 // 3: TravelersService.location
	(*DriverLocation)(nil),           // 4: TravelersService.driverLocation
	(*Driver)(nil),                   // 5: TravelersService.driver
}
var file_GrpcServices_TravelersService_proto_depIdxs = []int32{
	4, // 0: TravelersService.GetNearbyDriversResponse.driver:type_name -> TravelersService.driverLocation
	3, // 1: TravelersService.direction.origin:type_name -> TravelersService.location
	3, // 2: TravelersService.direction.destination:type_name -> TravelersService.location
	3, // 3: TravelersService.driverLocation.location:type_name -> TravelersService.location
	3, // 4: TravelersService.TravelersService.GetNearbyDrivers:input_type -> TravelersService.location
	1, // 5: TravelersService.TravelersService.GetPrice:input_type -> TravelersService.direction
	0, // 6: TravelersService.TravelersService.GetNearbyDrivers:output_type -> TravelersService.GetNearbyDriversResponse
	2, // 7: TravelersService.TravelersService.GetPrice:output_type -> TravelersService.price
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GrpcServices_TravelersService_proto_init() }
func file_GrpcServices_TravelersService_proto_init() {
	if File_GrpcServices_TravelersService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GrpcServices_TravelersService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNearbyDriversResponse); i {
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
		file_GrpcServices_TravelersService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Direction); i {
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
		file_GrpcServices_TravelersService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_GrpcServices_TravelersService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_GrpcServices_TravelersService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverLocation); i {
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
		file_GrpcServices_TravelersService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
			NumMessages:   6,
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

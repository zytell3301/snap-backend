// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package GrpcServices

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LocationReportClient is the client API for LocationReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationReportClient interface {
	UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	Deactivate(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
}

type locationReportClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationReportClient(cc grpc.ClientConnInterface) LocationReportClient {
	return &locationReportClient{cc}
}

func (c *locationReportClient) UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/GrpcServices.LocationReport/UpdateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationReportClient) Deactivate(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/GrpcServices.LocationReport/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationReportServer is the server API for LocationReport service.
// All implementations must embed UnimplementedLocationReportServer
// for forward compatibility
type LocationReportServer interface {
	UpdateLocation(context.Context, *Location) (*wrappers.BoolValue, error)
	Deactivate(context.Context, *empty.Empty) (*wrappers.BoolValue, error)
	mustEmbedUnimplementedLocationReportServer()
}

// UnimplementedLocationReportServer must be embedded to have forward compatible implementations.
type UnimplementedLocationReportServer struct {
}

func (UnimplementedLocationReportServer) UpdateLocation(context.Context, *Location) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedLocationReportServer) Deactivate(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedLocationReportServer) mustEmbedUnimplementedLocationReportServer() {}

// UnsafeLocationReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationReportServer will
// result in compilation errors.
type UnsafeLocationReportServer interface {
	mustEmbedUnimplementedLocationReportServer()
}

func RegisterLocationReportServer(s grpc.ServiceRegistrar, srv LocationReportServer) {
	s.RegisterService(&LocationReport_ServiceDesc, srv)
}

func _LocationReport_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationReportServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GrpcServices.LocationReport/UpdateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationReportServer).UpdateLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationReport_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationReportServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GrpcServices.LocationReport/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationReportServer).Deactivate(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationReport_ServiceDesc is the grpc.ServiceDesc for LocationReport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationReport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GrpcServices.LocationReport",
	HandlerType: (*LocationReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateLocation",
			Handler:    _LocationReport_UpdateLocation_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _LocationReport_Deactivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GrpcServices/DriversService.proto",
}

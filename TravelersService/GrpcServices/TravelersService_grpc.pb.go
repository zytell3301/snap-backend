// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package GrpcServices

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TravelersServiceClient is the client API for TravelersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelersServiceClient interface {
	GetNearbyDrivers(ctx context.Context, in *Location, opts ...grpc.CallOption) (*GetNearbyDriversResponse, error)
	GetPrice(ctx context.Context, in *Direction, opts ...grpc.CallOption) (*Price, error)
}

type travelersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelersServiceClient(cc grpc.ClientConnInterface) TravelersServiceClient {
	return &travelersServiceClient{cc}
}

func (c *travelersServiceClient) GetNearbyDrivers(ctx context.Context, in *Location, opts ...grpc.CallOption) (*GetNearbyDriversResponse, error) {
	out := new(GetNearbyDriversResponse)
	err := c.cc.Invoke(ctx, "/TravelersService.TravelersService/GetNearbyDrivers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelersServiceClient) GetPrice(ctx context.Context, in *Direction, opts ...grpc.CallOption) (*Price, error) {
	out := new(Price)
	err := c.cc.Invoke(ctx, "/TravelersService.TravelersService/GetPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelersServiceServer is the server API for TravelersService service.
// All implementations must embed UnimplementedTravelersServiceServer
// for forward compatibility
type TravelersServiceServer interface {
	GetNearbyDrivers(context.Context, *Location) (*GetNearbyDriversResponse, error)
	GetPrice(context.Context, *Direction) (*Price, error)
	mustEmbedUnimplementedTravelersServiceServer()
}

// UnimplementedTravelersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTravelersServiceServer struct {
}

func (UnimplementedTravelersServiceServer) GetNearbyDrivers(context.Context, *Location) (*GetNearbyDriversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearbyDrivers not implemented")
}
func (UnimplementedTravelersServiceServer) GetPrice(context.Context, *Direction) (*Price, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrice not implemented")
}
func (UnimplementedTravelersServiceServer) mustEmbedUnimplementedTravelersServiceServer() {}

// UnsafeTravelersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelersServiceServer will
// result in compilation errors.
type UnsafeTravelersServiceServer interface {
	mustEmbedUnimplementedTravelersServiceServer()
}

func RegisterTravelersServiceServer(s grpc.ServiceRegistrar, srv TravelersServiceServer) {
	s.RegisterService(&TravelersService_ServiceDesc, srv)
}

func _TravelersService_GetNearbyDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelersServiceServer).GetNearbyDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TravelersService.TravelersService/GetNearbyDrivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelersServiceServer).GetNearbyDrivers(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelersService_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Direction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelersServiceServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TravelersService.TravelersService/GetPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelersServiceServer).GetPrice(ctx, req.(*Direction))
	}
	return interceptor(ctx, in, info, handler)
}

// TravelersService_ServiceDesc is the grpc.ServiceDesc for TravelersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TravelersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TravelersService.TravelersService",
	HandlerType: (*TravelersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNearbyDrivers",
			Handler:    _TravelersService_GetNearbyDrivers_Handler,
		},
		{
			MethodName: "GetPrice",
			Handler:    _TravelersService_GetPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GrpcServices/TravelersService.proto",
}

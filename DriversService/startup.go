package DriversService

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"snap/DriversService/DriversLocationService"
	"snap/DriversService/GrpcServices"
	GrpcServices2 "snap/TravelersService/GrpcServices"
	"snap/TravelersService/TravelersServiceImplementation"
)

var Configs *viper.Viper

func InitiateGrpcServices() {
	listener, err := net.Listen("tcp", Configs.GetString("listen_address")+":"+Configs.GetString("port"))

	switch err != nil {
	case true:
		panic(err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(DriversLocationService.Authenticate))

	GrpcServices.RegisterDriversLocationReportServer(server, &DriversLocationService.DriversLocationService{})
	GrpcServices2.RegisterTravelersServiceServer(server, &TravelersServiceImplementation.TravelersService{})

	err = server.Serve(listener)

	switch err != nil {
	case true:
		panic(err)
	}
}

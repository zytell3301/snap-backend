package DriversService

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"snap/DriversService/DriversLocationService"
	"snap/DriversService/GrpcServices"
)

var Configs *viper.Viper

func InitiateGrpcServices() {
	listener, err := net.Listen("tcp", Configs.GetString("listen_address")+":"+Configs.GetString("port"))

	switch err != nil {
	case true:
		panic(err)
	}

	server := grpc.NewServer((grpc.WithUnaryInterceptor(DriversLocationService.Authenticate)).(grpc.ServerOption))

	GrpcServices.RegisterDriversLocationReportServer(server, &DriversLocationService.DriversLocationService{})

	err = server.Serve(listener)

	switch err != nil {
	case true:
		panic(err)
	}
}


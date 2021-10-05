package DriversService

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"snap/AuthService/AuthServiceImplementation"
	GrpcServices3 "snap/AuthService/GrpcServices"
	"snap/DriversService/DriversLocationService"
	"snap/DriversService/GrpcServices"
	GrpcServices2 "snap/TravelersService/GrpcServices"
	"snap/TravelersService/TravelersServiceImplementation"
)

var Configs *viper.Viper

func InitiateGrpcServices() {
	listener, err := net.Listen("tcp", Configs.GetString("listen_address")+":"+Configs.GetString("port"))
	AuthListener, err := net.Listen("tcp", "192.168.1.200:6167")

	switch err != nil {
	case true:
		panic(err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(DriversLocationService.Authenticate))

	GrpcServices.RegisterDriversLocationReportServer(server, &DriversLocationService.DriversLocationService{})
	GrpcServices2.RegisterTravelersServiceServer(server, &TravelersServiceImplementation.TravelersService{})

	AuthServer := grpc.NewServer()

	GrpcServices3.RegisterAuthServer(AuthServer, &AuthServiceImplementation.AuthServiceImplementation{})

	go func() {
		err = server.Serve(listener)

		switch err != nil {
		case true:
			panic(err)
		}
	}()

	go func() {
		err = AuthServer.Serve(AuthListener)

		switch err != nil {
		case true:
			panic(err)
		}
	}()
}

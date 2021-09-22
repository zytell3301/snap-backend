package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"snap/DriversService/GrpcServices"
	GrpcServices2 "snap/TravelersService/GrpcServices"
)

var Configs *viper.Viper
var Service GrpcServices.DriversLocationReportClient
var TravelersService GrpcServices2.TravelersServiceClient

func init() {
	initiateConfigs()
	initiatePackages()

	grpcServer, _ := grpc.Dial("localhost:6166", grpc.WithInsecure())

	Service = GrpcServices.NewDriversLocationReportClient(grpcServer)
	TravelersService = GrpcServices2.NewTravelersServiceClient(grpcServer)
}

func main() {
	server := iris.New()

	server.Get("/update-loc", func(ctx iris.Context) {
		md := metadata.New(map[string]string{"token": ctx.GetHeader("_token")})

		fmt.Println(Service.UpdateLocation(metadata.NewOutgoingContext(context.Background(), md), &GrpcServices.Location{
			X: 10,
			Y: 10,
		}))
	})

	server.Get("/deactivate-driver", func(ctx iris.Context) {
		md := metadata.New(map[string]string{"token": ctx.GetHeader("_token")})

		fmt.Println(Service.Deactivate(metadata.NewOutgoingContext(context.Background(), md), &empty.Empty{}))
	})

	server.Get("/request-driver", func(ctx iris.Context) {
		md := metadata.New(map[string]string{"token": ctx.GetHeader("_token")})

		fmt.Println(TravelersService.GetNearbyDrivers(metadata.NewOutgoingContext(context.Background(), md), &GrpcServices2.Location{X: 10, Y: 10}))
	})

	server.Listen(":"+Configs.GetString("port"), iris.WithLogLevel(Configs.GetString("log_level")))
}

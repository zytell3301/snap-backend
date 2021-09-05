package DriversLocationService

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"snap/Database/Redis"
	"snap/DriversService/GrpcServices"
)

type DriversLocationService struct {
	GrpcServices.UnimplementedDriversLocationReportServer
}

func (DriversLocationService) UpdateLocation(_ context.Context, location *GrpcServices.Location) (*wrappers.BoolValue, error) {
	Redis.Connection.GeoAdd(context.Background(), "drivers-positions", &redis.GeoLocation{
		Longitude: location.X,
		Latitude:  location.Y,
	})
	return nil, nil
}


func Authenticate(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("Interceptor called")
	return nil
}

package DriversLocationService

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/wrappers"
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

package TravelersServiceImplementation

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/empty"
	"snap/Database/Cassandra/Models"
	"snap/Database/Redis"
	"snap/TravelersService/GrpcServices"
)

type TravelersService struct {
	GrpcServices.UnimplementedTravelersServiceServer
}

func (TravelersService) GetNearbyDrivers(ctx context.Context, location *GrpcServices.Location) (*empty.Empty, error) {
	driversLocation, err := getNearbyDrivers(ctx, location)

	switch err != nil {
	case true:
		return &empty.Empty{}, err
	}

	switch len(driversLocation) == 0 {
	case true:
		return &empty.Empty{}, errors.New("no driver found")
	}

	var drivers []*Models.Driver

	for _, driverLocation := range driversLocation {
		drivers = append(drivers, Models.Users.GetDriverDetails(map[string]interface{}{"id": driverLocation.Name}))
	}

	fmt.Println(drivers)

	return &empty.Empty{}, nil
}

func getNearbyDrivers(ctx context.Context, location *GrpcServices.Location) ([]redis.GeoLocation, error) {
	return Redis.Connection.GeoRadius(ctx, "drivers-positions", location.X, location.Y, &redis.GeoRadiusQuery{
		Radius:    3,
		Unit:      "km",
		WithCoord: true,
		WithDist:  true,
	}).Result()
}

package TravelersServiceImplementation

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"snap/Database/Redis"
	"snap/TravelersService/GrpcServices"
)

type TravelersService struct {
	GrpcServices.UnimplementedTravelersServiceServer
}

func (TravelersService) GetNearbyDrivers(ctx context.Context, location *GrpcServices.Location) (*GrpcServices.GetNearbyDriversResponse, error) {
	driversLocation, err := getNearbyDrivers(ctx, location)

	switch err != nil {
	case true:
		return &GrpcServices.GetNearbyDriversResponse{}, err
	}

	switch len(driversLocation) == 0 {
	case true:
		return &GrpcServices.GetNearbyDriversResponse{}, errors.New("no driver found")
	}

	var drivers []*GrpcServices.DriverLocation

	for _, driverLocation := range driversLocation {
		drivers = append(drivers, &GrpcServices.DriverLocation{
			Id: driverLocation.Name,
			Location: &GrpcServices.Location{
				X: driverLocation.Latitude,
				Y: driverLocation.Longitude,
			},
		})
	}

	return &GrpcServices.GetNearbyDriversResponse{Driver: drivers}, nil
}

func getNearbyDrivers(ctx context.Context, location *GrpcServices.Location) ([]redis.GeoLocation, error) {
	return Redis.Connection.GeoRadius(ctx, "drivers-positions", location.X, location.Y, &redis.GeoRadiusQuery{
		Radius:    3,
		Unit:      "km",
		WithCoord: true,
		WithDist:  true,
	}).Result()
}

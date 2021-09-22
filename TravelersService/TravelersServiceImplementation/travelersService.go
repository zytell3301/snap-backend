package TravelersServiceImplementation

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"snap/Database/Cassandra/Models"
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

	var drivers []Models.User

	for _, driverLocation := range driversLocation {
		driver, err := Models.Users.GetDriverDetails(map[string]interface{}{"id": driverLocation.Name})
		switch err != nil {
		case false:
			drivers = append(drivers, driver)
		}
	}

	return &GrpcServices.GetNearbyDriversResponse{}, nil
}

func getNearbyDrivers(ctx context.Context, location *GrpcServices.Location) ([]redis.GeoLocation, error) {
	return Redis.Connection.GeoRadius(ctx, "drivers-positions", location.X, location.Y, &redis.GeoRadiusQuery{
		Radius:    3,
		Unit:      "km",
		WithCoord: true,
		WithDist:  true,
	}).Result()
}

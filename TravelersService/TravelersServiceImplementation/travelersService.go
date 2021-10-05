package TravelersServiceImplementation

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"snap/Database/Cassandra/Models"
	"snap/Database/Redis"
	"snap/Socket"
	"snap/TravelersService/GrpcServices"
	"time"
)

type TravelersService struct {
	GrpcServices.UnimplementedTravelersServiceServer
}

func (TravelersService) GetNearbyDrivers(ctx context.Context, location *GrpcServices.Location) (*GrpcServices.GetNearbyDriversResponse, error) {
	driversLocation, err := getNearbyDrivers(ctx, location)

	switch err != nil {
	case true:
		fmt.Println(err)
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

func (TravelersService) GetPrice(ctx context.Context, direction *GrpcServices.Direction) (*GrpcServices.Price, error) {
	price := math.Sqrt(math.Pow(direction.Origin.X-direction.Destination.X, 2)+math.Pow(direction.Origin.Y-direction.Destination.Y, 2)) * 350000
	remainder := math.Remainder(price, 500)

	switch remainder < 250 {
	case true:
		return &GrpcServices.Price{Price: int32(price - remainder)}, nil
	default:
		return &GrpcServices.Price{Price: int32(price - remainder + 500)}, nil
	}
}

func (TravelersService) RequestDriver(ctx context.Context, direction *GrpcServices.Direction) (*GrpcServices.Driver, error) {
	availableDrivers, _ := getNearbyDrivers(ctx, direction.Origin)

	fmt.Println("RequestDriver")

	fmt.Println(direction.Origin)
	fmt.Println(availableDrivers)

	switch len(availableDrivers) == 0 {
	case true:
		return &GrpcServices.Driver{}, errors.New("no driver found")
	}
	driver, availableDrivers := availableDrivers[0], availableDrivers[1:]
	fmt.Println(driver.Name + "-travel-request")
	_, err := Socket.Connection.Request(driver.Name+"-travel-request", nil, 15*time.Second)

	for err != nil && len(availableDrivers) != 0 {
		driver, availableDrivers = availableDrivers[0], availableDrivers[1:]
		_, err = Socket.Connection.Request(driver.Name+"-travel-request", nil, 15*time.Second)
	}

	switch err != nil {
	case true:
		fmt.Println(err)
		return &GrpcServices.Driver{}, errors.New("no driver accepted the travel")
	}

	model, err := Models.Users.GetDriverDetails(map[string]interface{}{"id": driver.Name})

	switch err != nil {
	case true:
		return &GrpcServices.Driver{}, err
	default:
		fmt.Println("success", err)
		return &GrpcServices.Driver{
			Id:         model.Id.String(),
			Name:       model.Driver_details.Name,
			Lastname:   model.Driver_details.Lastname,
			VehicleNo:  model.Driver_details.Vehicle_no,
			ProfilePic: model.Driver_details.Profile_pic,
		}, err
	}

}

func getNearbyDrivers(ctx context.Context, location *GrpcServices.Location) ([]redis.GeoLocation, error) {
	return Redis.Connection.GeoRadius(ctx, "drivers-positions", location.X, location.Y, &redis.GeoRadiusQuery{
		Radius:    3,
		Unit:      "km",
		WithCoord: true,
		WithDist:  true,
	}).Result()
}

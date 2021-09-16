package DriversLocationService

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"snap/Database/Cassandra/Models"
	"snap/Database/Redis"
	"snap/DriversService/GrpcServices"
	"strings"
)

type DriversLocationService struct {
	GrpcServices.UnimplementedDriversLocationReportServer
}

func (DriversLocationService) UpdateLocation(ctx context.Context, location *GrpcServices.Location) (*wrappers.BoolValue, error) {
	md, isSuccess := metadata.FromIncomingContext(ctx)

	switch isSuccess == false {
	case true:
		return &wrappers.BoolValue{}, errors.New("failed to extract context metadata")
	}

	Redis.Connection.GeoAdd(context.Background(), "drivers-positions", &redis.GeoLocation{
		Name:      strings.Join(md.Get("user_id"), ""),
		Longitude: location.X,
		Latitude:  location.Y,
	})

	return &wrappers.BoolValue{Value: true}, nil
}

func (DriversLocationService) Deactivate(ctx context.Context, _ *empty.Empty) (*wrappers.BoolValue, error) {
	md, isSuccess := metadata.FromIncomingContext(ctx)

	switch isSuccess == false {
	case true:
		return &wrappers.BoolValue{}, errors.New("failed to extract context metadata")
	}

	result := Redis.Connection.ZRem(ctx, "drivers-positions", strings.Join(md.Get("user_id"), ""))

	return &wrappers.BoolValue{}, result.Err()
}

func Authenticate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, success := metadata.FromIncomingContext(ctx)
	communicationToken := strings.Join(md.Get("token"), "")

	switch success == false || communicationToken == "" {
	case true:
		return nil, errors.New("an error occurred while extracting metadata. please make sure that context is passed correctly and a valid communicationToken is supplied")
	}

	token, err := Models.Tokens.GetToken(communicationToken)
	switch err != nil || token.Communication_token == "" {
	case true:
		return nil, errors.New("an error occurred while validating supplied token")
	}

	md.Set("user_id", token.User_id.String())

	return handler(metadata.NewIncomingContext(ctx, md), req)
}

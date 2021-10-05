package AuthServiceImplementation

import (
	"context"
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	"snap/AuthService/GrpcServices"
	"snap/Database/Cassandra/Models"
)

type AuthServiceImplementation struct {
	GrpcServices.UnimplementedAuthServer
}

func (AuthServiceImplementation) Authenticate(ctx context.Context, credentials *GrpcServices.Credentials) (*GrpcServices.Token, error) {
	user := Models.UserPkPhone.GetUser(credentials.Phone)
	fmt.Println(user)

	switch user.Password == "" {
	case true:
		return &GrpcServices.Token{}, errors.New("user not found")
	}

	passwordValidity := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	fmt.Println(passwordValidity)

	switch passwordValidity != nil {
	case true:
		return &GrpcServices.Token{}, errors.New("invalid credentials supplied")
	}

	token := Models.Tokens.GenerateToken(user.Id)

	switch token == "" {
	case true:
		return &GrpcServices.Token{}, errors.New("an error occurred while generating token. please try again later")
	}

	return &GrpcServices.Token{Token: token}, nil
}

func (AuthServiceImplementation) NewAccount(ctx context.Context, user *GrpcServices.User) (*GrpcServices.Token, error) {
	usr := Models.UserPkPhone.GetUser(user.Phone)

	switch usr.Phone != "" {
	case true:
		return &GrpcServices.Token{}, errors.New("user already exists")
	}

	userId, err := gocql.RandomUUID()
	switch err != nil {
	case true:
		fmt.Println(err)
		return &GrpcServices.Token{}, errors.New("an error occurred while creating a new user. please try again later")
	}

	driverDetails := Models.Driver{
		Name:       user.DriverDetails.Name,
		Lastname:   user.DriverDetails.Lastname,
		Vehicle_no: user.DriverDetails.VehicleNo,
		Balance:    0,
	}

	err = Models.Users.NewUser(Models.User{
		Id:             userId,
		Phone:          user.Phone,
		Password:       user.Password,
		Driver_details: driverDetails,
	})

	switch err != nil {
	case true:
		fmt.Println(err)
		return &GrpcServices.Token{}, errors.New("an error occurred while creating a new user. please try again later")
	}

	token := Models.Tokens.GenerateToken(userId)

	switch token == "" {
	case true:
		return &GrpcServices.Token{}, errors.New("an error occurred while generating token. please attempt to login again")
	}

	return &GrpcServices.Token{Token: token}, nil
}

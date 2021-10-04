package AuthServiceImplementation

import (
	"context"
	"errors"
	"fmt"
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

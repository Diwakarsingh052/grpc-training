package main

import (
	"context"
	"fmt"
	"server/gen/proto"
)

type User struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
	PasswordHash string   `json:"-"`
}

func (us *userService) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	//fetching the request
	nu := req.GetUser()
	u := User{
		ID:           "101",
		Name:         nu.Name,
		Email:        nu.Email,
		Roles:        nu.Roles,
		PasswordHash: "jpasfjdpjaspfdj",
	}
	// you must have put this in the db

	fmt.Println(u)

	fmt.Println("sending the response")
	res := &proto.SignupResponse{Result: "user created successfully"}
	return res, nil

}

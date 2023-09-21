package main

import (
	"client/gen/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {

	dialOpt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:5001", dialOpt...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	//creates a gRPC client for the UserService service defined in the pb package
	//and binds it to the specified connection conn.
	client := proto.NewUserServiceClient(conn)

	req := &proto.SignupRequest{
		User: &proto.User{
			Name:     "John",
			Email:    "john@email.com",
			Password: "abc",
			Roles:    []string{"ADMIN", "USER"},
		}}
	ctx := context.Background() // create an empty container
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel() // clean the resources when work is done

	//calling the remote service method
	res, err := client.Signup(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Result)
}

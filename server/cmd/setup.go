package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"server/gen/proto"
)

type userService struct {
	proto.UnimplementedUserServiceServer
}

func main() {
	//sets up a TCP listener on port 5001 for incoming network connections.
	listener, err := net.Listen("tcp", ":5001")

	if err != nil {
		log.Println(err)
		return
	}

	//NewServer creates a gRPC server which has no service registered
	//and has not started to accept requests yet.
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &userService{})

	err = s.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}

}

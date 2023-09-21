package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	pb "server/gen/proto"
)

//In bidirectional-streaming RPC, the client is sending a request to the server as a stream of messages.
//The server also responds with a stream of messages.
//The call has to be initiated from the client side,
//But after that,the communication is completely based on the application logic of the gRPC client and the server.

func (us *userService) GreetEveryone(stream pb.UserService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		//receiving the streaming request from the client
		req, err := stream.Recv()

		//If the client has finished sending the request, we will quit
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		//check if a client is disconnected or not
		select {
		case <-stream.Context().Done():
			log.Println("client cancelled the request")

			//generating a custom error for gRPC
			return status.Error(codes.Internal, "client disconnected")

		default:
			// The Client is still connected
		}

		msg := "Hello " + req.FirstName + "!"

		res := &pb.GreetEveryoneResponse{Result: msg}
		//sending the response
		err = stream.Send(res)
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

	}
}

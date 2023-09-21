package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

//In server streaming, the server sends back a sequence of responses
//after getting the client’s request message.

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
	client := pb.NewUserServiceClient(conn)

	req := &pb.GetPostRequest{UserId: 101}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	// doing server streaming request
	stream, err := client.GetPosts(ctx, req)

	if err != nil {
		log.Println(err)
		return
	}

	for {
		//receiving values from stream
		post, err := stream.Recv()
		//if the server has finished sending the request, we will quit
		if err == io.EOF {
			break
		}
		//any other kind of error would be caught here
		if err != nil {
			log.Println(err)
		}

		fmt.Println("reading stream")
		//printing data received
		fmt.Println(post)
		fmt.Println()
	}

}

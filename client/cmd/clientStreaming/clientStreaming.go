package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
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
	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//calling the remote service method
	stream, err := client.CreatePost(ctx)
	if err != nil {
		log.Fatalf("failed to call createPost server: %v", err)
	}

	//generating the first batch of the request that we would send to the server
	batch1 := []*pb.Post{
		{
			Title:  "The Science of Design",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "The Politics of Power",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "The Art of Programming",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	req := &pb.CreatePostRequest{Posts: batch1}
	
	//sending the request to the server
	err = stream.Send(req)
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}

	//adding latency
	time.Sleep(4 * time.Second)

	//constructing the second batch
	batch2 := []*pb.Post{
		{
			Title:  "Post 11",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "Post 21",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "Post 31",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}
	req = &pb.CreatePostRequest{Posts: batch2}
	err = stream.Send(req)
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}

	//close a client-streaming, and receive the server's response message.
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	fmt.Println(resp.Result)

}

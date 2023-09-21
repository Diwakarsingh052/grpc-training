package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"sync"
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

	stream, err := client.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("failed to call GreetEveryone stream: %v\n", err)
	}

	requests := []*pb.GreetEveryoneRequest{
		{FirstName: "John"},
		{FirstName: "Bruce"},
		{FirstName: "Roy"},
	}
	//using WaitGroup to wait for goroutines to finish
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// sending multiple requests to the remote service
	go func() {
		defer wg.Done()
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)

			//sending requests
			err := stream.Send(req)
			if err != nil {
				closeErr := stream.CloseSend()
				if closeErr != nil {
					log.Printf("Failed to close stream: %v", closeErr)
					return
				}
				return
			}
			time.Sleep(1 * time.Second)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	//recv the multiple responses back from the remote service
	go func() {
		defer wg.Done()
		for {

			res, err := stream.Recv()
			if err == io.EOF {
				log.Printf("stream has ended")
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			//client disconnected
			select {
			case <-stream.Context().Done():
				log.Println("remote service cancelled the request")
				return
			default:
				// The Client is still connected
			}
			//if everything is good then print the results
			log.Printf("Received: %v\n", res.Result)
		}
	}()
	wg.Wait()
	fmt.Println("end of bidirectional communication")
}

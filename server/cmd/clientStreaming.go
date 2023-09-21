package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	pb "server/gen/proto"
	"sync"
)

// In client-streaming RPC, the client sends multiple messages/request to the server
// instead of a single request.
// The server sends back a single response to the client.

func (us *userService) CreatePost(stream pb.UserService_CreatePostServer) error {
	// Receive CreatePost request from client in batches
	wg := &sync.WaitGroup{}
	for {

		//receive the req from the client
		req, err := stream.Recv()

		//If the client has finished sending the request, we will quit
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		//during the request if a client close the connection we will know inside this select block using the
		//context.Done()
		//time.Sleep(time.Second * 4)
		select {
		case <-stream.Context().Done():
			log.Println("client cancelled the request")

			//generating a custom error for gRPC
			return status.Error(codes.Internal, "client disconnected")

		default:
			// The Client is still connected
		}
		posts := req.GetPosts()
		log.Println(posts)

		log.Println("adding all the posts into the db")

		wg.Add(1)
		go AddPost(posts, wg) //async operation

	}

	wg.Wait()
	res := &pb.CreatePostResponse{Result: "all posts added in db"}
	//Return response
	return stream.SendAndClose(res)
}

func AddPost(posts []*pb.Post, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, p := range posts {
		log.Println("adding post ", p.Title)
	}

}

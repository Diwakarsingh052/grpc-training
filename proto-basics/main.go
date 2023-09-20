package main

import (
	"fmt"
	pb "proto-basics/proto"
)

func main() {
	r := pb.BlogRequest{
		BlogId:  101,
		Title:   "Introduction to Protocol Buffers",
		Content: "Test",
	}

	// to get data from the struct, we could use get methods defined in proto package
	fmt.Println(r.GetBlogId())
	fmt.Println(r.String())

	// use SearchResponse type
	s := pb.SearchResponse{

		Results: []*pb.Result{
			{},
			{},
		},
	}
	fmt.Println(s.GetResults())
}

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
	fmt.Println(r)
	fmt.Println(r.GetBlogId())
}

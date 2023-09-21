// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/user.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_Signup_FullMethodName     = "/proto.UserService/Signup"
	UserService_GetPosts_FullMethodName   = "/proto.UserService/GetPosts"
	UserService_CreatePost_FullMethodName = "/proto.UserService/CreatePost"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// unary
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	// server streaming // stream keywords indicate that series of responses would be sent back.
	GetPosts(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (UserService_GetPostsClient, error)
	// client streaming
	CreatePost(ctx context.Context, opts ...grpc.CallOption) (UserService_CreatePostClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, UserService_Signup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPosts(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (UserService_GetPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], UserService_GetPosts_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetPostsClient interface {
	Recv() (*GetPostResponse, error)
	grpc.ClientStream
}

type userServiceGetPostsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetPostsClient) Recv() (*GetPostResponse, error) {
	m := new(GetPostResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) CreatePost(ctx context.Context, opts ...grpc.CallOption) (UserService_CreatePostClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], UserService_CreatePost_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceCreatePostClient{stream}
	return x, nil
}

type UserService_CreatePostClient interface {
	Send(*CreatePostRequest) error
	CloseAndRecv() (*CreatePostResponse, error)
	grpc.ClientStream
}

type userServiceCreatePostClient struct {
	grpc.ClientStream
}

func (x *userServiceCreatePostClient) Send(m *CreatePostRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceCreatePostClient) CloseAndRecv() (*CreatePostResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreatePostResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// unary
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	// server streaming // stream keywords indicate that series of responses would be sent back.
	GetPosts(*GetPostRequest, UserService_GetPostsServer) error
	// client streaming
	CreatePost(UserService_CreatePostServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServiceServer) GetPosts(*GetPostRequest, UserService_GetPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedUserServiceServer) CreatePost(UserService_CreatePostServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPostRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetPosts(m, &userServiceGetPostsServer{stream})
}

type UserService_GetPostsServer interface {
	Send(*GetPostResponse) error
	grpc.ServerStream
}

type userServiceGetPostsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetPostsServer) Send(m *GetPostResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_CreatePost_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).CreatePost(&userServiceCreatePostServer{stream})
}

type UserService_CreatePostServer interface {
	SendAndClose(*CreatePostResponse) error
	Recv() (*CreatePostRequest, error)
	grpc.ServerStream
}

type userServiceCreatePostServer struct {
	grpc.ServerStream
}

func (x *userServiceCreatePostServer) SendAndClose(m *CreatePostResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceCreatePostServer) Recv() (*CreatePostRequest, error) {
	m := new(CreatePostRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UserService_Signup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPosts",
			Handler:       _UserService_GetPosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreatePost",
			Handler:       _UserService_CreatePost_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/user.proto",
}

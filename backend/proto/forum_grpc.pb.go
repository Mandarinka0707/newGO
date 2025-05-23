// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: forum.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ForumService_CreateCategory_FullMethodName     = "/forum.ForumService/CreateCategory"
	ForumService_GetCategory_FullMethodName        = "/forum.ForumService/GetCategory"
	ForumService_CreateTopic_FullMethodName        = "/forum.ForumService/CreateTopic"
	ForumService_GetTopic_FullMethodName           = "/forum.ForumService/GetTopic"
	ForumService_CreateMessage_FullMethodName      = "/forum.ForumService/CreateMessage"
	ForumService_GetMessage_FullMethodName         = "/forum.ForumService/GetMessage"
	ForumService_CreatePost_FullMethodName         = "/forum.ForumService/CreatePost"
	ForumService_GetPosts_FullMethodName           = "/forum.ForumService/GetPosts"
	ForumService_CreateChatMessage_FullMethodName  = "/forum.ForumService/CreateChatMessage"
	ForumService_StreamChatMessages_FullMethodName = "/forum.ForumService/StreamChatMessages"
)

// ForumServiceClient is the client API for ForumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для работы с форумом
type ForumServiceClient interface {
	// Категории
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error)
	// Темы
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error)
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error)
	// Сообщения
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
	// Посты
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
	// Чат
	CreateChatMessage(ctx context.Context, in *CreateChatMessageRequest, opts ...grpc.CallOption) (*CreateChatMessageResponse, error)
	StreamChatMessages(ctx context.Context, in *StreamChatMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatMessage], error)
}

type forumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewForumServiceClient(cc grpc.ClientConnInterface) ForumServiceClient {
	return &forumServiceClient{cc}
}

func (c *forumServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, ForumService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryResponse)
	err := c.cc.Invoke(ctx, ForumService_GetCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTopicResponse)
	err := c.cc.Invoke(ctx, ForumService_CreateTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTopicResponse)
	err := c.cc.Invoke(ctx, ForumService_GetTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, ForumService_CreateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageResponse)
	err := c.cc.Invoke(ctx, ForumService_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, ForumService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, ForumService_GetPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreateChatMessage(ctx context.Context, in *CreateChatMessageRequest, opts ...grpc.CallOption) (*CreateChatMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChatMessageResponse)
	err := c.cc.Invoke(ctx, ForumService_CreateChatMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) StreamChatMessages(ctx context.Context, in *StreamChatMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ForumService_ServiceDesc.Streams[0], ForumService_StreamChatMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamChatMessagesRequest, ChatMessage]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ForumService_StreamChatMessagesClient = grpc.ServerStreamingClient[ChatMessage]

// ForumServiceServer is the server API for ForumService service.
// All implementations must embed UnimplementedForumServiceServer
// for forward compatibility.
//
// Сервис для работы с форумом
type ForumServiceServer interface {
	// Категории
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error)
	// Темы
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error)
	GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error)
	// Сообщения
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	// Посты
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error)
	// Чат
	CreateChatMessage(context.Context, *CreateChatMessageRequest) (*CreateChatMessageResponse, error)
	StreamChatMessages(*StreamChatMessagesRequest, grpc.ServerStreamingServer[ChatMessage]) error
	mustEmbedUnimplementedForumServiceServer()
}

// UnimplementedForumServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedForumServiceServer struct{}

func (UnimplementedForumServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedForumServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedForumServiceServer) CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedForumServiceServer) GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedForumServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedForumServiceServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedForumServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedForumServiceServer) GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedForumServiceServer) CreateChatMessage(context.Context, *CreateChatMessageRequest) (*CreateChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatMessage not implemented")
}
func (UnimplementedForumServiceServer) StreamChatMessages(*StreamChatMessagesRequest, grpc.ServerStreamingServer[ChatMessage]) error {
	return status.Errorf(codes.Unimplemented, "method StreamChatMessages not implemented")
}
func (UnimplementedForumServiceServer) mustEmbedUnimplementedForumServiceServer() {}
func (UnimplementedForumServiceServer) testEmbeddedByValue()                      {}

// UnsafeForumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForumServiceServer will
// result in compilation errors.
type UnsafeForumServiceServer interface {
	mustEmbedUnimplementedForumServiceServer()
}

func RegisterForumServiceServer(s grpc.ServiceRegistrar, srv ForumServiceServer) {
	// If the following call pancis, it indicates UnimplementedForumServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ForumService_ServiceDesc, srv)
}

func _ForumService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_GetPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetPosts(ctx, req.(*GetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreateChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForumService_CreateChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateChatMessage(ctx, req.(*CreateChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_StreamChatMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamChatMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ForumServiceServer).StreamChatMessages(m, &grpc.GenericServerStream[StreamChatMessagesRequest, ChatMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ForumService_StreamChatMessagesServer = grpc.ServerStreamingServer[ChatMessage]

// ForumService_ServiceDesc is the grpc.ServiceDesc for ForumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ForumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forum.ForumService",
	HandlerType: (*ForumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _ForumService_CreateCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _ForumService_GetCategory_Handler,
		},
		{
			MethodName: "CreateTopic",
			Handler:    _ForumService_CreateTopic_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _ForumService_GetTopic_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _ForumService_CreateMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _ForumService_GetMessage_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _ForumService_CreatePost_Handler,
		},
		{
			MethodName: "GetPosts",
			Handler:    _ForumService_GetPosts_Handler,
		},
		{
			MethodName: "CreateChatMessage",
			Handler:    _ForumService_CreateChatMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChatMessages",
			Handler:       _ForumService_StreamChatMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "forum.proto",
}

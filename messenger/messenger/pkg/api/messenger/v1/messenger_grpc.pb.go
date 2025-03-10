// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: messenger.proto

package messenger

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MESSENGERClient is the client API for MESSENGER service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MESSENGERClient interface {
	InitSession(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitSessionResp, error)
	CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*CreateChatResp, error)
	SendMessageToChat(ctx context.Context, in *SendMessageToChatReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMessagesFromChat(ctx context.Context, in *GetMessagesFromChatReq, opts ...grpc.CallOption) (*GetMessagesFromChatResp, error)
}

type mESSENGERClient struct {
	cc grpc.ClientConnInterface
}

func NewMESSENGERClient(cc grpc.ClientConnInterface) MESSENGERClient {
	return &mESSENGERClient{cc}
}

func (c *mESSENGERClient) InitSession(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitSessionResp, error) {
	out := new(InitSessionResp)
	err := c.cc.Invoke(ctx, "/MESSENGER/InitSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mESSENGERClient) CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*CreateChatResp, error) {
	out := new(CreateChatResp)
	err := c.cc.Invoke(ctx, "/MESSENGER/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mESSENGERClient) SendMessageToChat(ctx context.Context, in *SendMessageToChatReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/MESSENGER/SendMessageToChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mESSENGERClient) GetMessagesFromChat(ctx context.Context, in *GetMessagesFromChatReq, opts ...grpc.CallOption) (*GetMessagesFromChatResp, error) {
	out := new(GetMessagesFromChatResp)
	err := c.cc.Invoke(ctx, "/MESSENGER/GetMessagesFromChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MESSENGERServer is the server API for MESSENGER service.
// All implementations must embed UnimplementedMESSENGERServer
// for forward compatibility
type MESSENGERServer interface {
	InitSession(context.Context, *emptypb.Empty) (*InitSessionResp, error)
	CreateChat(context.Context, *CreateChatReq) (*CreateChatResp, error)
	SendMessageToChat(context.Context, *SendMessageToChatReq) (*emptypb.Empty, error)
	GetMessagesFromChat(context.Context, *GetMessagesFromChatReq) (*GetMessagesFromChatResp, error)
	mustEmbedUnimplementedMESSENGERServer()
}

// UnimplementedMESSENGERServer must be embedded to have forward compatible implementations.
type UnimplementedMESSENGERServer struct {
}

func (UnimplementedMESSENGERServer) InitSession(context.Context, *emptypb.Empty) (*InitSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitSession not implemented")
}
func (UnimplementedMESSENGERServer) CreateChat(context.Context, *CreateChatReq) (*CreateChatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedMESSENGERServer) SendMessageToChat(context.Context, *SendMessageToChatReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToChat not implemented")
}
func (UnimplementedMESSENGERServer) GetMessagesFromChat(context.Context, *GetMessagesFromChatReq) (*GetMessagesFromChatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesFromChat not implemented")
}
func (UnimplementedMESSENGERServer) mustEmbedUnimplementedMESSENGERServer() {}

// UnsafeMESSENGERServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MESSENGERServer will
// result in compilation errors.
type UnsafeMESSENGERServer interface {
	mustEmbedUnimplementedMESSENGERServer()
}

func RegisterMESSENGERServer(s grpc.ServiceRegistrar, srv MESSENGERServer) {
	s.RegisterService(&MESSENGER_ServiceDesc, srv)
}

func _MESSENGER_InitSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MESSENGERServer).InitSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MESSENGER/InitSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MESSENGERServer).InitSession(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MESSENGER_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MESSENGERServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MESSENGER/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MESSENGERServer).CreateChat(ctx, req.(*CreateChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MESSENGER_SendMessageToChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MESSENGERServer).SendMessageToChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MESSENGER/SendMessageToChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MESSENGERServer).SendMessageToChat(ctx, req.(*SendMessageToChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MESSENGER_GetMessagesFromChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesFromChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MESSENGERServer).GetMessagesFromChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MESSENGER/GetMessagesFromChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MESSENGERServer).GetMessagesFromChat(ctx, req.(*GetMessagesFromChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MESSENGER_ServiceDesc is the grpc.ServiceDesc for MESSENGER service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MESSENGER_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MESSENGER",
	HandlerType: (*MESSENGERServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitSession",
			Handler:    _MESSENGER_InitSession_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _MESSENGER_CreateChat_Handler,
		},
		{
			MethodName: "SendMessageToChat",
			Handler:    _MESSENGER_SendMessageToChat_Handler,
		},
		{
			MethodName: "GetMessagesFromChat",
			Handler:    _MESSENGER_GetMessagesFromChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messenger.proto",
}

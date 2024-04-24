// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: duo_service.proto

package pb

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

// DUOServiceClient is the client API for DUOService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DUOServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	RequestLoginChallenge(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginChallengeRequest, error)
	SubmitLoginChallenge(ctx context.Context, in *LoginChallengeResponse, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateLobby(ctx context.Context, in *CreateLobbyRequest, opts ...grpc.CallOption) (DUOService_CreateLobbyClient, error)
	JoinLobby(ctx context.Context, in *JoinLobbyRequest, opts ...grpc.CallOption) (DUOService_JoinLobbyClient, error)
	DisconnectLobby(ctx context.Context, in *DisconnectLobbyRequest, opts ...grpc.CallOption) (*DisconnectLobbyResponse, error)
}

type dUOServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDUOServiceClient(cc grpc.ClientConnInterface) DUOServiceClient {
	return &dUOServiceClient{cc}
}

func (c *dUOServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/pb.DUOService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dUOServiceClient) RequestLoginChallenge(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginChallengeRequest, error) {
	out := new(LoginChallengeRequest)
	err := c.cc.Invoke(ctx, "/pb.DUOService/RequestLoginChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dUOServiceClient) SubmitLoginChallenge(ctx context.Context, in *LoginChallengeResponse, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.DUOService/SubmitLoginChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dUOServiceClient) CreateLobby(ctx context.Context, in *CreateLobbyRequest, opts ...grpc.CallOption) (DUOService_CreateLobbyClient, error) {
	stream, err := c.cc.NewStream(ctx, &DUOService_ServiceDesc.Streams[0], "/pb.DUOService/CreateLobby", opts...)
	if err != nil {
		return nil, err
	}
	x := &dUOServiceCreateLobbyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DUOService_CreateLobbyClient interface {
	Recv() (*LobbyStatus, error)
	grpc.ClientStream
}

type dUOServiceCreateLobbyClient struct {
	grpc.ClientStream
}

func (x *dUOServiceCreateLobbyClient) Recv() (*LobbyStatus, error) {
	m := new(LobbyStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dUOServiceClient) JoinLobby(ctx context.Context, in *JoinLobbyRequest, opts ...grpc.CallOption) (DUOService_JoinLobbyClient, error) {
	stream, err := c.cc.NewStream(ctx, &DUOService_ServiceDesc.Streams[1], "/pb.DUOService/JoinLobby", opts...)
	if err != nil {
		return nil, err
	}
	x := &dUOServiceJoinLobbyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DUOService_JoinLobbyClient interface {
	Recv() (*LobbyStatus, error)
	grpc.ClientStream
}

type dUOServiceJoinLobbyClient struct {
	grpc.ClientStream
}

func (x *dUOServiceJoinLobbyClient) Recv() (*LobbyStatus, error) {
	m := new(LobbyStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dUOServiceClient) DisconnectLobby(ctx context.Context, in *DisconnectLobbyRequest, opts ...grpc.CallOption) (*DisconnectLobbyResponse, error) {
	out := new(DisconnectLobbyResponse)
	err := c.cc.Invoke(ctx, "/pb.DUOService/DisconnectLobby", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DUOServiceServer is the server API for DUOService service.
// All implementations must embed UnimplementedDUOServiceServer
// for forward compatibility
type DUOServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	RequestLoginChallenge(context.Context, *LoginRequest) (*LoginChallengeRequest, error)
	SubmitLoginChallenge(context.Context, *LoginChallengeResponse) (*LoginResponse, error)
	CreateLobby(*CreateLobbyRequest, DUOService_CreateLobbyServer) error
	JoinLobby(*JoinLobbyRequest, DUOService_JoinLobbyServer) error
	DisconnectLobby(context.Context, *DisconnectLobbyRequest) (*DisconnectLobbyResponse, error)
	mustEmbedUnimplementedDUOServiceServer()
}

// UnimplementedDUOServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDUOServiceServer struct {
}

func (UnimplementedDUOServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDUOServiceServer) RequestLoginChallenge(context.Context, *LoginRequest) (*LoginChallengeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLoginChallenge not implemented")
}
func (UnimplementedDUOServiceServer) SubmitLoginChallenge(context.Context, *LoginChallengeResponse) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitLoginChallenge not implemented")
}
func (UnimplementedDUOServiceServer) CreateLobby(*CreateLobbyRequest, DUOService_CreateLobbyServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateLobby not implemented")
}
func (UnimplementedDUOServiceServer) JoinLobby(*JoinLobbyRequest, DUOService_JoinLobbyServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinLobby not implemented")
}
func (UnimplementedDUOServiceServer) DisconnectLobby(context.Context, *DisconnectLobbyRequest) (*DisconnectLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectLobby not implemented")
}
func (UnimplementedDUOServiceServer) mustEmbedUnimplementedDUOServiceServer() {}

// UnsafeDUOServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DUOServiceServer will
// result in compilation errors.
type UnsafeDUOServiceServer interface {
	mustEmbedUnimplementedDUOServiceServer()
}

func RegisterDUOServiceServer(s grpc.ServiceRegistrar, srv DUOServiceServer) {
	s.RegisterService(&DUOService_ServiceDesc, srv)
}

func _DUOService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DUOServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DUOService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DUOServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DUOService_RequestLoginChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DUOServiceServer).RequestLoginChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DUOService/RequestLoginChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DUOServiceServer).RequestLoginChallenge(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DUOService_SubmitLoginChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginChallengeResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DUOServiceServer).SubmitLoginChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DUOService/SubmitLoginChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DUOServiceServer).SubmitLoginChallenge(ctx, req.(*LoginChallengeResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DUOService_CreateLobby_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateLobbyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DUOServiceServer).CreateLobby(m, &dUOServiceCreateLobbyServer{stream})
}

type DUOService_CreateLobbyServer interface {
	Send(*LobbyStatus) error
	grpc.ServerStream
}

type dUOServiceCreateLobbyServer struct {
	grpc.ServerStream
}

func (x *dUOServiceCreateLobbyServer) Send(m *LobbyStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _DUOService_JoinLobby_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinLobbyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DUOServiceServer).JoinLobby(m, &dUOServiceJoinLobbyServer{stream})
}

type DUOService_JoinLobbyServer interface {
	Send(*LobbyStatus) error
	grpc.ServerStream
}

type dUOServiceJoinLobbyServer struct {
	grpc.ServerStream
}

func (x *dUOServiceJoinLobbyServer) Send(m *LobbyStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _DUOService_DisconnectLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DUOServiceServer).DisconnectLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DUOService/DisconnectLobby",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DUOServiceServer).DisconnectLobby(ctx, req.(*DisconnectLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DUOService_ServiceDesc is the grpc.ServiceDesc for DUOService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DUOService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DUOService",
	HandlerType: (*DUOServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DUOService_Register_Handler,
		},
		{
			MethodName: "RequestLoginChallenge",
			Handler:    _DUOService_RequestLoginChallenge_Handler,
		},
		{
			MethodName: "SubmitLoginChallenge",
			Handler:    _DUOService_SubmitLoginChallenge_Handler,
		},
		{
			MethodName: "DisconnectLobby",
			Handler:    _DUOService_DisconnectLobby_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateLobby",
			Handler:       _DUOService_CreateLobby_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "JoinLobby",
			Handler:       _DUOService_JoinLobby_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "duo_service.proto",
}

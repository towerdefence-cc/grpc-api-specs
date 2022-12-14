// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: player_transporter.proto

package player_transporter

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

// PlayerTransporterClient is the client API for PlayerTransporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerTransporterClient interface {
	CommonMovePlayer(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TowerDefenceGameMovePlayer(ctx context.Context, in *TowerDefenceGameMoveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type playerTransporterClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerTransporterClient(cc grpc.ClientConnInterface) PlayerTransporterClient {
	return &playerTransporterClient{cc}
}

func (c *playerTransporterClient) CommonMovePlayer(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_transporter.PlayerTransporter/CommonMovePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTransporterClient) TowerDefenceGameMovePlayer(ctx context.Context, in *TowerDefenceGameMoveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_transporter.PlayerTransporter/TowerDefenceGameMovePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerTransporterServer is the server API for PlayerTransporter service.
// All implementations must embed UnimplementedPlayerTransporterServer
// for forward compatibility
type PlayerTransporterServer interface {
	CommonMovePlayer(context.Context, *MoveRequest) (*emptypb.Empty, error)
	TowerDefenceGameMovePlayer(context.Context, *TowerDefenceGameMoveRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPlayerTransporterServer()
}

// UnimplementedPlayerTransporterServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerTransporterServer struct {
}

func (UnimplementedPlayerTransporterServer) CommonMovePlayer(context.Context, *MoveRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommonMovePlayer not implemented")
}
func (UnimplementedPlayerTransporterServer) TowerDefenceGameMovePlayer(context.Context, *TowerDefenceGameMoveRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TowerDefenceGameMovePlayer not implemented")
}
func (UnimplementedPlayerTransporterServer) mustEmbedUnimplementedPlayerTransporterServer() {}

// UnsafePlayerTransporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerTransporterServer will
// result in compilation errors.
type UnsafePlayerTransporterServer interface {
	mustEmbedUnimplementedPlayerTransporterServer()
}

func RegisterPlayerTransporterServer(s grpc.ServiceRegistrar, srv PlayerTransporterServer) {
	s.RegisterService(&PlayerTransporter_ServiceDesc, srv)
}

func _PlayerTransporter_CommonMovePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTransporterServer).CommonMovePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_transporter.PlayerTransporter/CommonMovePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTransporterServer).CommonMovePlayer(ctx, req.(*MoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTransporter_TowerDefenceGameMovePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TowerDefenceGameMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTransporterServer).TowerDefenceGameMovePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_transporter.PlayerTransporter/TowerDefenceGameMovePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTransporterServer).TowerDefenceGameMovePlayer(ctx, req.(*TowerDefenceGameMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerTransporter_ServiceDesc is the grpc.ServiceDesc for PlayerTransporter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerTransporter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "towerdefence.cc.service.player_transporter.PlayerTransporter",
	HandlerType: (*PlayerTransporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommonMovePlayer",
			Handler:    _PlayerTransporter_CommonMovePlayer_Handler,
		},
		{
			MethodName: "TowerDefenceGameMovePlayer",
			Handler:    _PlayerTransporter_TowerDefenceGameMovePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player_transporter.proto",
}

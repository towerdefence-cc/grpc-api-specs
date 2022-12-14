// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: server_discovery.proto

package server_discovery

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

// ServerDiscoveryClient is the client API for ServerDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerDiscoveryClient interface {
	GetSuggestedLobbyServer(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*LobbyServer, error)
	GetSuggestedOtpServer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConnectableServer, error)
	GetSuggestedTowerDefenceServer(ctx context.Context, in *TowerDefenceServerRequest, opts ...grpc.CallOption) (*ConnectableServer, error)
}

type serverDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewServerDiscoveryClient(cc grpc.ClientConnInterface) ServerDiscoveryClient {
	return &serverDiscoveryClient{cc}
}

func (c *serverDiscoveryClient) GetSuggestedLobbyServer(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*LobbyServer, error) {
	out := new(LobbyServer)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.server_discovery.ServerDiscovery/GetSuggestedLobbyServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDiscoveryClient) GetSuggestedOtpServer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConnectableServer, error) {
	out := new(ConnectableServer)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.server_discovery.ServerDiscovery/GetSuggestedOtpServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDiscoveryClient) GetSuggestedTowerDefenceServer(ctx context.Context, in *TowerDefenceServerRequest, opts ...grpc.CallOption) (*ConnectableServer, error) {
	out := new(ConnectableServer)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.server_discovery.ServerDiscovery/GetSuggestedTowerDefenceServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerDiscoveryServer is the server API for ServerDiscovery service.
// All implementations must embed UnimplementedServerDiscoveryServer
// for forward compatibility
type ServerDiscoveryServer interface {
	GetSuggestedLobbyServer(context.Context, *ServerRequest) (*LobbyServer, error)
	GetSuggestedOtpServer(context.Context, *emptypb.Empty) (*ConnectableServer, error)
	GetSuggestedTowerDefenceServer(context.Context, *TowerDefenceServerRequest) (*ConnectableServer, error)
	mustEmbedUnimplementedServerDiscoveryServer()
}

// UnimplementedServerDiscoveryServer must be embedded to have forward compatible implementations.
type UnimplementedServerDiscoveryServer struct {
}

func (UnimplementedServerDiscoveryServer) GetSuggestedLobbyServer(context.Context, *ServerRequest) (*LobbyServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuggestedLobbyServer not implemented")
}
func (UnimplementedServerDiscoveryServer) GetSuggestedOtpServer(context.Context, *emptypb.Empty) (*ConnectableServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuggestedOtpServer not implemented")
}
func (UnimplementedServerDiscoveryServer) GetSuggestedTowerDefenceServer(context.Context, *TowerDefenceServerRequest) (*ConnectableServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuggestedTowerDefenceServer not implemented")
}
func (UnimplementedServerDiscoveryServer) mustEmbedUnimplementedServerDiscoveryServer() {}

// UnsafeServerDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerDiscoveryServer will
// result in compilation errors.
type UnsafeServerDiscoveryServer interface {
	mustEmbedUnimplementedServerDiscoveryServer()
}

func RegisterServerDiscoveryServer(s grpc.ServiceRegistrar, srv ServerDiscoveryServer) {
	s.RegisterService(&ServerDiscovery_ServiceDesc, srv)
}

func _ServerDiscovery_GetSuggestedLobbyServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDiscoveryServer).GetSuggestedLobbyServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.server_discovery.ServerDiscovery/GetSuggestedLobbyServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDiscoveryServer).GetSuggestedLobbyServer(ctx, req.(*ServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDiscovery_GetSuggestedOtpServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDiscoveryServer).GetSuggestedOtpServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.server_discovery.ServerDiscovery/GetSuggestedOtpServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDiscoveryServer).GetSuggestedOtpServer(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDiscovery_GetSuggestedTowerDefenceServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TowerDefenceServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDiscoveryServer).GetSuggestedTowerDefenceServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.server_discovery.ServerDiscovery/GetSuggestedTowerDefenceServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDiscoveryServer).GetSuggestedTowerDefenceServer(ctx, req.(*TowerDefenceServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerDiscovery_ServiceDesc is the grpc.ServiceDesc for ServerDiscovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerDiscovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "towerdefence.cc.service.server_discovery.ServerDiscovery",
	HandlerType: (*ServerDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSuggestedLobbyServer",
			Handler:    _ServerDiscovery_GetSuggestedLobbyServer_Handler,
		},
		{
			MethodName: "GetSuggestedOtpServer",
			Handler:    _ServerDiscovery_GetSuggestedOtpServer_Handler,
		},
		{
			MethodName: "GetSuggestedTowerDefenceServer",
			Handler:    _ServerDiscovery_GetSuggestedTowerDefenceServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server_discovery.proto",
}

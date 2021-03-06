// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: usermgmt/usermgmt.proto

package gRPC_Service

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

// UserManagamentClient is the client API for UserManagament service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagamentClient interface {
	CreateNewUSer(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *GetUsersParams, opts ...grpc.CallOption) (*UserList, error)
}

type userManagamentClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagamentClient(cc grpc.ClientConnInterface) UserManagamentClient {
	return &userManagamentClient{cc}
}

func (c *userManagamentClient) CreateNewUSer(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/usermgmt.UserManagament/CreateNewUSer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagamentClient) GetUsers(ctx context.Context, in *GetUsersParams, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/usermgmt.UserManagament/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagamentServer is the server API for UserManagament service.
// All implementations must embed UnimplementedUserManagamentServer
// for forward compatibility
type UserManagamentServer interface {
	CreateNewUSer(context.Context, *NewUser) (*User, error)
	GetUsers(context.Context, *GetUsersParams) (*UserList, error)
	mustEmbedUnimplementedUserManagamentServer()
}

// UnimplementedUserManagamentServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagamentServer struct {
}

func (UnimplementedUserManagamentServer) CreateNewUSer(context.Context, *NewUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUSer not implemented")
}
func (UnimplementedUserManagamentServer) GetUsers(context.Context, *GetUsersParams) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserManagamentServer) mustEmbedUnimplementedUserManagamentServer() {}

// UnsafeUserManagamentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagamentServer will
// result in compilation errors.
type UnsafeUserManagamentServer interface {
	mustEmbedUnimplementedUserManagamentServer()
}

func RegisterUserManagamentServer(s grpc.ServiceRegistrar, srv UserManagamentServer) {
	s.RegisterService(&UserManagament_ServiceDesc, srv)
}

func _UserManagament_CreateNewUSer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagamentServer).CreateNewUSer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgmt.UserManagament/CreateNewUSer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagamentServer).CreateNewUSer(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagament_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagamentServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgmt.UserManagament/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagamentServer).GetUsers(ctx, req.(*GetUsersParams))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagament_ServiceDesc is the grpc.ServiceDesc for UserManagament service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagament_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermgmt.UserManagament",
	HandlerType: (*UserManagamentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewUSer",
			Handler:    _UserManagament_CreateNewUSer_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserManagament_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermgmt/usermgmt.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: user.proto

package user_proto

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
	UserService_GetUserByLogin_FullMethodName    = "/UserService/GetUserByLogin"
	UserService_IsUserExistByUUID_FullMethodName = "/UserService/IsUserExistByUUID"
	UserService_GetUserInfoByUUID_FullMethodName = "/UserService/GetUserInfoByUUID"
	UserService_GetLoginByUUID_FullMethodName    = "/UserService/GetLoginByUUID"
	UserService_GetUserWithOffset_FullMethodName = "/UserService/GetUserWithOffset"
	UserService_UpdateProfile_FullMethodName     = "/UserService/UpdateProfile"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for friends
type UserServiceClient interface {
	// Add friends method
	GetUserByLogin(ctx context.Context, in *GetUserByLoginIn, opts ...grpc.CallOption) (*GetUserByLoginOut, error)
	IsUserExistByUUID(ctx context.Context, in *IsUserExistByUUIDIn, opts ...grpc.CallOption) (*IsUserExistByUUIDOut, error)
	GetUserInfoByUUID(ctx context.Context, in *GetUserInfoByUUIDIn, opts ...grpc.CallOption) (*GetUserInfoByUUIDOut, error)
	GetLoginByUUID(ctx context.Context, in *GetLoginByUUIDIn, opts ...grpc.CallOption) (*GetLoginByUUIDOut, error)
	GetUserWithOffset(ctx context.Context, in *GetUserWithOffsetIn, opts ...grpc.CallOption) (*GetUserWithOffsetOut, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileIn, opts ...grpc.CallOption) (*UpdateProfileOut, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserByLogin(ctx context.Context, in *GetUserByLoginIn, opts ...grpc.CallOption) (*GetUserByLoginOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserByLoginOut)
	err := c.cc.Invoke(ctx, UserService_GetUserByLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsUserExistByUUID(ctx context.Context, in *IsUserExistByUUIDIn, opts ...grpc.CallOption) (*IsUserExistByUUIDOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsUserExistByUUIDOut)
	err := c.cc.Invoke(ctx, UserService_IsUserExistByUUID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoByUUID(ctx context.Context, in *GetUserInfoByUUIDIn, opts ...grpc.CallOption) (*GetUserInfoByUUIDOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoByUUIDOut)
	err := c.cc.Invoke(ctx, UserService_GetUserInfoByUUID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetLoginByUUID(ctx context.Context, in *GetLoginByUUIDIn, opts ...grpc.CallOption) (*GetLoginByUUIDOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLoginByUUIDOut)
	err := c.cc.Invoke(ctx, UserService_GetLoginByUUID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserWithOffset(ctx context.Context, in *GetUserWithOffsetIn, opts ...grpc.CallOption) (*GetUserWithOffsetOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserWithOffsetOut)
	err := c.cc.Invoke(ctx, UserService_GetUserWithOffset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileIn, opts ...grpc.CallOption) (*UpdateProfileOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProfileOut)
	err := c.cc.Invoke(ctx, UserService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
//
// Service for friends
type UserServiceServer interface {
	// Add friends method
	GetUserByLogin(context.Context, *GetUserByLoginIn) (*GetUserByLoginOut, error)
	IsUserExistByUUID(context.Context, *IsUserExistByUUIDIn) (*IsUserExistByUUIDOut, error)
	GetUserInfoByUUID(context.Context, *GetUserInfoByUUIDIn) (*GetUserInfoByUUIDOut, error)
	GetLoginByUUID(context.Context, *GetLoginByUUIDIn) (*GetLoginByUUIDOut, error)
	GetUserWithOffset(context.Context, *GetUserWithOffsetIn) (*GetUserWithOffsetOut, error)
	UpdateProfile(context.Context, *UpdateProfileIn) (*UpdateProfileOut, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUserByLogin(context.Context, *GetUserByLoginIn) (*GetUserByLoginOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByLogin not implemented")
}
func (UnimplementedUserServiceServer) IsUserExistByUUID(context.Context, *IsUserExistByUUIDIn) (*IsUserExistByUUIDOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserExistByUUID not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoByUUID(context.Context, *GetUserInfoByUUIDIn) (*GetUserInfoByUUIDOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUUID not implemented")
}
func (UnimplementedUserServiceServer) GetLoginByUUID(context.Context, *GetLoginByUUIDIn) (*GetLoginByUUIDOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginByUUID not implemented")
}
func (UnimplementedUserServiceServer) GetUserWithOffset(context.Context, *GetUserWithOffsetIn) (*GetUserWithOffsetOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserWithOffset not implemented")
}
func (UnimplementedUserServiceServer) UpdateProfile(context.Context, *UpdateProfileIn) (*UpdateProfileOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserByLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByLoginIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByLogin(ctx, req.(*GetUserByLoginIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsUserExistByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsUserExistByUUIDIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsUserExistByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_IsUserExistByUUID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsUserExistByUUID(ctx, req.(*IsUserExistByUUIDIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByUUIDIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfoByUUID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoByUUID(ctx, req.(*GetUserInfoByUUIDIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetLoginByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginByUUIDIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetLoginByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetLoginByUUID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetLoginByUUID(ctx, req.(*GetLoginByUUIDIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserWithOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserWithOffsetIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserWithOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserWithOffset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserWithOffset(ctx, req.(*GetUserWithOffsetIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfile(ctx, req.(*UpdateProfileIn))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByLogin",
			Handler:    _UserService_GetUserByLogin_Handler,
		},
		{
			MethodName: "IsUserExistByUUID",
			Handler:    _UserService_IsUserExistByUUID_Handler,
		},
		{
			MethodName: "GetUserInfoByUUID",
			Handler:    _UserService_GetUserInfoByUUID_Handler,
		},
		{
			MethodName: "GetLoginByUUID",
			Handler:    _UserService_GetLoginByUUID_Handler,
		},
		{
			MethodName: "GetUserWithOffset",
			Handler:    _UserService_GetUserWithOffset_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserService_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

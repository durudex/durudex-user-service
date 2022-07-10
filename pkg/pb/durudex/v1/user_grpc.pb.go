// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package durudexv1

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// Getting a user by id.
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	// Getting a user by credentials.
	GetUserByCreds(ctx context.Context, in *GetUserByCredsRequest, opts ...grpc.CallOption) (*GetUserByCredsResponse, error)
	// Forgoting a user password.
	ForgotUserPassword(ctx context.Context, in *ForgotUserPasswordRequest, opts ...grpc.CallOption) (*ForgotUserPasswordResponse, error)
	// Updating a user avatar.
	UpdateUserAvatar(ctx context.Context, in *UpdateUserAvatarRequest, opts ...grpc.CallOption) (*UpdateUserAvatarResponse, error)
	// Creatinf a new user verification code.
	CreateVerifyUserEmailCode(ctx context.Context, in *CreateVerifyUserEmailCodeRequest, opts ...grpc.CallOption) (*CreateVerifyUserEmailCodeResponse, error)
	// Verifying a user email code.
	VerifyUserEmailCode(ctx context.Context, in *VerifyUserEmailCodeRequest, opts ...grpc.CallOption) (*VerifyUserEmailCodeResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.UserService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByCreds(ctx context.Context, in *GetUserByCredsRequest, opts ...grpc.CallOption) (*GetUserByCredsResponse, error) {
	out := new(GetUserByCredsResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.UserService/GetUserByCreds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ForgotUserPassword(ctx context.Context, in *ForgotUserPasswordRequest, opts ...grpc.CallOption) (*ForgotUserPasswordResponse, error) {
	out := new(ForgotUserPasswordResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.UserService/ForgotUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserAvatar(ctx context.Context, in *UpdateUserAvatarRequest, opts ...grpc.CallOption) (*UpdateUserAvatarResponse, error) {
	out := new(UpdateUserAvatarResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.UserService/UpdateUserAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateVerifyUserEmailCode(ctx context.Context, in *CreateVerifyUserEmailCodeRequest, opts ...grpc.CallOption) (*CreateVerifyUserEmailCodeResponse, error) {
	out := new(CreateVerifyUserEmailCodeResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.UserService/CreateVerifyUserEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyUserEmailCode(ctx context.Context, in *VerifyUserEmailCodeRequest, opts ...grpc.CallOption) (*VerifyUserEmailCodeResponse, error) {
	out := new(VerifyUserEmailCodeResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.UserService/VerifyUserEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// Getting a user by id.
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	// Getting a user by credentials.
	GetUserByCreds(context.Context, *GetUserByCredsRequest) (*GetUserByCredsResponse, error)
	// Forgoting a user password.
	ForgotUserPassword(context.Context, *ForgotUserPasswordRequest) (*ForgotUserPasswordResponse, error)
	// Updating a user avatar.
	UpdateUserAvatar(context.Context, *UpdateUserAvatarRequest) (*UpdateUserAvatarResponse, error)
	// Creatinf a new user verification code.
	CreateVerifyUserEmailCode(context.Context, *CreateVerifyUserEmailCodeRequest) (*CreateVerifyUserEmailCodeResponse, error)
	// Verifying a user email code.
	VerifyUserEmailCode(context.Context, *VerifyUserEmailCodeRequest) (*VerifyUserEmailCodeResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUserByCreds(context.Context, *GetUserByCredsRequest) (*GetUserByCredsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByCreds not implemented")
}
func (UnimplementedUserServiceServer) ForgotUserPassword(context.Context, *ForgotUserPasswordRequest) (*ForgotUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotUserPassword not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserAvatar(context.Context, *UpdateUserAvatarRequest) (*UpdateUserAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAvatar not implemented")
}
func (UnimplementedUserServiceServer) CreateVerifyUserEmailCode(context.Context, *CreateVerifyUserEmailCodeRequest) (*CreateVerifyUserEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVerifyUserEmailCode not implemented")
}
func (UnimplementedUserServiceServer) VerifyUserEmailCode(context.Context, *VerifyUserEmailCodeRequest) (*VerifyUserEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserEmailCode not implemented")
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

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByCredsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.UserService/GetUserByCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByCreds(ctx, req.(*GetUserByCredsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ForgotUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ForgotUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.UserService/ForgotUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ForgotUserPassword(ctx, req.(*ForgotUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.UserService/UpdateUserAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserAvatar(ctx, req.(*UpdateUserAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateVerifyUserEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVerifyUserEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateVerifyUserEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.UserService/CreateVerifyUserEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateVerifyUserEmailCode(ctx, req.(*CreateVerifyUserEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyUserEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyUserEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.UserService/VerifyUserEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyUserEmailCode(ctx, req.(*VerifyUserEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "durudex.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "GetUserByCreds",
			Handler:    _UserService_GetUserByCreds_Handler,
		},
		{
			MethodName: "ForgotUserPassword",
			Handler:    _UserService_ForgotUserPassword_Handler,
		},
		{
			MethodName: "UpdateUserAvatar",
			Handler:    _UserService_UpdateUserAvatar_Handler,
		},
		{
			MethodName: "CreateVerifyUserEmailCode",
			Handler:    _UserService_CreateVerifyUserEmailCode_Handler,
		},
		{
			MethodName: "VerifyUserEmailCode",
			Handler:    _UserService_VerifyUserEmailCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "durudex/v1/user.proto",
}

// Copyright © 2022 Durudex
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: durudex/v1/user_auth.proto

package durudexv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// User Sign Up Request.
type UserSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique username.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Ununique user email address.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// User password.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Verification code.
	Code uint64 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UserSignUpRequest) Reset() {
	*x = UserSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignUpRequest) ProtoMessage() {}

func (x *UserSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignUpRequest.ProtoReflect.Descriptor instead.
func (*UserSignUpRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserSignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserSignUpRequest) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

// User Sign Up Response.
type UserSignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User uuid.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserSignUpResponse) Reset() {
	*x = UserSignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignUpResponse) ProtoMessage() {}

func (x *UserSignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignUpResponse.ProtoReflect.Descriptor instead.
func (*UserSignUpResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserSignUpResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

// User Sign In Request.
type UserSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Username.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// User password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// User ip address.
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *UserSignInRequest) Reset() {
	*x = UserSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignInRequest) ProtoMessage() {}

func (x *UserSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignInRequest.ProtoReflect.Descriptor instead.
func (*UserSignInRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserSignInRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserSignInRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// User Sign In Response.
type UserSignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User authentication JWT access token.
	Access string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	// User authentication refresh token.
	Refresh string `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *UserSignInResponse) Reset() {
	*x = UserSignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignInResponse) ProtoMessage() {}

func (x *UserSignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignInResponse.ProtoReflect.Descriptor instead.
func (*UserSignInResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UserSignInResponse) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *UserSignInResponse) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

// User Sign Out Request.
type UserSignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User authentication refresh token.
	Refresh string `protobuf:"bytes,1,opt,name=refresh,proto3" json:"refresh,omitempty"`
	// User ip address.
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *UserSignOutRequest) Reset() {
	*x = UserSignOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignOutRequest) ProtoMessage() {}

func (x *UserSignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignOutRequest.ProtoReflect.Descriptor instead.
func (*UserSignOutRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UserSignOutRequest) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

func (x *UserSignOutRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// User Sign Out Response.
type UserSignOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserSignOutResponse) Reset() {
	*x = UserSignOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignOutResponse) ProtoMessage() {}

func (x *UserSignOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignOutResponse.ProtoReflect.Descriptor instead.
func (*UserSignOutResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{5}
}

// Refresh user authentication token request.
type RefreshUserTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User authentication refresh token.
	Refresh string `protobuf:"bytes,1,opt,name=refresh,proto3" json:"refresh,omitempty"`
	// User ip address.
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *RefreshUserTokenRequest) Reset() {
	*x = RefreshUserTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshUserTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshUserTokenRequest) ProtoMessage() {}

func (x *RefreshUserTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshUserTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshUserTokenRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshUserTokenRequest) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

func (x *RefreshUserTokenRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// Refresh user authentication token response.
type RefreshUserTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User authentication JWT access token.
	Access string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *RefreshUserTokenResponse) Reset() {
	*x = RefreshUserTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshUserTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshUserTokenResponse) ProtoMessage() {}

func (x *RefreshUserTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshUserTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshUserTokenResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_auth_proto_rawDescGZIP(), []int{7}
}

func (x *RefreshUserTokenResponse) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

var File_durudex_v1_user_auth_proto protoreflect.FileDescriptor

var file_durudex_v1_user_auth_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x75,
	0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x22, 0x75, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x24, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x22, 0x46, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x22, 0x3e, 0x0a, 0x12, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x43, 0x0a, 0x17, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x32, 0x0a, 0x18, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xda, 0x02, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1d, 0x2e, 0x64,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x75,
	0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1d, 0x2e, 0x64, 0x75, 0x72, 0x75,
	0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x2e, 0x64,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb0, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f,
	0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65,
	0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x44, 0x75, 0x72, 0x75,
	0x64, 0x65, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x44,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_durudex_v1_user_auth_proto_rawDescOnce sync.Once
	file_durudex_v1_user_auth_proto_rawDescData = file_durudex_v1_user_auth_proto_rawDesc
)

func file_durudex_v1_user_auth_proto_rawDescGZIP() []byte {
	file_durudex_v1_user_auth_proto_rawDescOnce.Do(func() {
		file_durudex_v1_user_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_durudex_v1_user_auth_proto_rawDescData)
	})
	return file_durudex_v1_user_auth_proto_rawDescData
}

var file_durudex_v1_user_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_durudex_v1_user_auth_proto_goTypes = []interface{}{
	(*UserSignUpRequest)(nil),        // 0: durudex.v1.UserSignUpRequest
	(*UserSignUpResponse)(nil),       // 1: durudex.v1.UserSignUpResponse
	(*UserSignInRequest)(nil),        // 2: durudex.v1.UserSignInRequest
	(*UserSignInResponse)(nil),       // 3: durudex.v1.UserSignInResponse
	(*UserSignOutRequest)(nil),       // 4: durudex.v1.UserSignOutRequest
	(*UserSignOutResponse)(nil),      // 5: durudex.v1.UserSignOutResponse
	(*RefreshUserTokenRequest)(nil),  // 6: durudex.v1.RefreshUserTokenRequest
	(*RefreshUserTokenResponse)(nil), // 7: durudex.v1.RefreshUserTokenResponse
}
var file_durudex_v1_user_auth_proto_depIdxs = []int32{
	0, // 0: durudex.v1.UserAuthService.UserSignUp:input_type -> durudex.v1.UserSignUpRequest
	2, // 1: durudex.v1.UserAuthService.UserSignIn:input_type -> durudex.v1.UserSignInRequest
	4, // 2: durudex.v1.UserAuthService.UserSignOut:input_type -> durudex.v1.UserSignOutRequest
	6, // 3: durudex.v1.UserAuthService.RefreshUserToken:input_type -> durudex.v1.RefreshUserTokenRequest
	1, // 4: durudex.v1.UserAuthService.UserSignUp:output_type -> durudex.v1.UserSignUpResponse
	3, // 5: durudex.v1.UserAuthService.UserSignIn:output_type -> durudex.v1.UserSignInResponse
	5, // 6: durudex.v1.UserAuthService.UserSignOut:output_type -> durudex.v1.UserSignOutResponse
	7, // 7: durudex.v1.UserAuthService.RefreshUserToken:output_type -> durudex.v1.RefreshUserTokenResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_durudex_v1_user_auth_proto_init() }
func file_durudex_v1_user_auth_proto_init() {
	if File_durudex_v1_user_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_durudex_v1_user_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignUpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durudex_v1_user_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignUpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durudex_v1_user_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignInRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durudex_v1_user_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignInResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durudex_v1_user_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignOutRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durudex_v1_user_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignOutResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durudex_v1_user_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshUserTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durudex_v1_user_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshUserTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_durudex_v1_user_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_durudex_v1_user_auth_proto_goTypes,
		DependencyIndexes: file_durudex_v1_user_auth_proto_depIdxs,
		MessageInfos:      file_durudex_v1_user_auth_proto_msgTypes,
	}.Build()
	File_durudex_v1_user_auth_proto = out.File
	file_durudex_v1_user_auth_proto_rawDesc = nil
	file_durudex_v1_user_auth_proto_goTypes = nil
	file_durudex_v1_user_auth_proto_depIdxs = nil
}

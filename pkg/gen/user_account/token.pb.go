// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: user_account/token.proto

package api

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetTokensRequest) Reset() {
	*x = GetTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokensRequest) ProtoMessage() {}

func (x *GetTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokensRequest.ProtoReflect.Descriptor instead.
func (*GetTokensRequest) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokensRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Token `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetTokensResponse) Reset() {
	*x = GetTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokensResponse) ProtoMessage() {}

func (x *GetTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokensResponse.ProtoReflect.Descriptor instead.
func (*GetTokensResponse) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokensResponse) GetItems() []*Token {
	if x != nil {
		return x.Items
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Описание токена
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Тип токена
	TokenType string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	// Пользователь выпустивший токен
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Время истечения токена
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expired_at,json=expiredAt,proto3,oneof" json:"expired_at,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Token) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Token) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Token) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

type CreateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Описание токена
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Тип токена
	TokenType string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	// Пользователь выпустивший токен
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Время действия токена
	TimeDuration *string `protobuf:"bytes,4,opt,name=time_duration,json=timeDuration,proto3,oneof" json:"time_duration,omitempty"`
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTokenRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTokenRequest) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *CreateTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTokenRequest) GetTimeDuration() string {
	if x != nil && x.TimeDuration != nil {
		return *x.TimeDuration
	}
	return ""
}

type CreateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Токен показываем только при создании, просим запомнить
}

func (x *CreateTokenResponse) Reset() {
	*x = CreateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenResponse) ProtoMessage() {}

func (x *CreateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID токена
	TokenId string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// Пользователь выпустивший токен
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteTokenRequest) Reset() {
	*x = DeleteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenRequest) ProtoMessage() {}

func (x *DeleteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTokenRequest) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *DeleteTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{6}
}

func (x *GetTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenResponse) Reset() {
	*x = GetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_token_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResponse) ProtoMessage() {}

func (x *GetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_token_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResponse.ProtoReflect.Descriptor instead.
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return file_user_account_token_proto_rawDescGZIP(), []int{7}
}

func (x *GetTokenResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_user_account_token_proto protoreflect.FileDescriptor

var file_user_account_token_proto_rawDesc = []byte{
	0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x3e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xa4, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x48, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x8e, 0x05, 0x0a, 0x0c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x92, 0x41, 0x3c, 0x12, 0x3a, 0xd0,
	0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0,
	0xb5, 0x20, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb2,
	0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xb7, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0,
	0xb0, 0xd1, 0x82, 0xd0, 0xb5, 0xd0, 0xbb, 0xd1, 0x8f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12,
	0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0xab, 0x01, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x57, 0x92, 0x41, 0x3f, 0x12, 0x3d, 0xd0, 0xa1, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb4, 0xd0,
	0xb0, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb5,
	0xd0, 0xbd, 0xd0, 0xb0, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd0, 0xbf, 0xd0, 0xbe,
	0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xb7, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xb5,
	0xd0, 0xbb, 0xd1, 0x8f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0xa1, 0x01, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x58, 0x92, 0x41, 0x38, 0x12, 0x36, 0xd0, 0xa3, 0xd0, 0xb4, 0xd0, 0xb0,
	0xd0, 0xbb, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd1, 0x82, 0xd0, 0xbe, 0xd0,
	0xba, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb0, 0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x8c,
	0xd0, 0xb7, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xb5, 0xd0, 0xbb, 0xd1, 0x8f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x2f, 0x7b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x89, 0x01,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x92, 0x41, 0x21, 0x12, 0x1f,
	0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8,
	0xd0, 0xb5, 0x20, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb0, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x2f, 0x7b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x7d, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x69,
	0x78, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_account_token_proto_rawDescOnce sync.Once
	file_user_account_token_proto_rawDescData = file_user_account_token_proto_rawDesc
)

func file_user_account_token_proto_rawDescGZIP() []byte {
	file_user_account_token_proto_rawDescOnce.Do(func() {
		file_user_account_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_account_token_proto_rawDescData)
	})
	return file_user_account_token_proto_rawDescData
}

var file_user_account_token_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_account_token_proto_goTypes = []any{
	(*GetTokensRequest)(nil),      // 0: user_account.GetTokensRequest
	(*GetTokensResponse)(nil),     // 1: user_account.GetTokensResponse
	(*Token)(nil),                 // 2: user_account.Token
	(*CreateTokenRequest)(nil),    // 3: user_account.CreateTokenRequest
	(*CreateTokenResponse)(nil),   // 4: user_account.CreateTokenResponse
	(*DeleteTokenRequest)(nil),    // 5: user_account.DeleteTokenRequest
	(*GetTokenRequest)(nil),       // 6: user_account.GetTokenRequest
	(*GetTokenResponse)(nil),      // 7: user_account.GetTokenResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_user_account_token_proto_depIdxs = []int32{
	2, // 0: user_account.GetTokensResponse.items:type_name -> user_account.Token
	8, // 1: user_account.Token.expired_at:type_name -> google.protobuf.Timestamp
	2, // 2: user_account.GetTokenResponse.token:type_name -> user_account.Token
	0, // 3: user_account.TokenService.GetTokens:input_type -> user_account.GetTokensRequest
	3, // 4: user_account.TokenService.CreateToken:input_type -> user_account.CreateTokenRequest
	5, // 5: user_account.TokenService.DeleteToken:input_type -> user_account.DeleteTokenRequest
	6, // 6: user_account.TokenService.GetToken:input_type -> user_account.GetTokenRequest
	1, // 7: user_account.TokenService.GetTokens:output_type -> user_account.GetTokensResponse
	4, // 8: user_account.TokenService.CreateToken:output_type -> user_account.CreateTokenResponse
	9, // 9: user_account.TokenService.DeleteToken:output_type -> google.protobuf.Empty
	7, // 10: user_account.TokenService.GetToken:output_type -> user_account.GetTokenResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_account_token_proto_init() }
func file_user_account_token_proto_init() {
	if File_user_account_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_account_token_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTokensRequest); i {
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
		file_user_account_token_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetTokensResponse); i {
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
		file_user_account_token_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Token); i {
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
		file_user_account_token_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTokenRequest); i {
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
		file_user_account_token_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTokenResponse); i {
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
		file_user_account_token_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTokenRequest); i {
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
		file_user_account_token_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetTokenRequest); i {
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
		file_user_account_token_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetTokenResponse); i {
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
	file_user_account_token_proto_msgTypes[2].OneofWrappers = []any{}
	file_user_account_token_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_account_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_account_token_proto_goTypes,
		DependencyIndexes: file_user_account_token_proto_depIdxs,
		MessageInfos:      file_user_account_token_proto_msgTypes,
	}.Build()
	File_user_account_token_proto = out.File
	file_user_account_token_proto_rawDesc = nil
	file_user_account_token_proto_goTypes = nil
	file_user_account_token_proto_depIdxs = nil
}

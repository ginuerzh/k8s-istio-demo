// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserDetail struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Avatar               string   `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDetail) Reset()         { *m = UserDetail{} }
func (m *UserDetail) String() string { return proto.CompactTextString(m) }
func (*UserDetail) ProtoMessage()    {}
func (*UserDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetail.Unmarshal(m, b)
}
func (m *UserDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetail.Marshal(b, m, deterministic)
}
func (m *UserDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetail.Merge(m, src)
}
func (m *UserDetail) XXX_Size() int {
	return xxx_messageInfo_UserDetail.Size(m)
}
func (m *UserDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetail.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetail proto.InternalMessageInfo

func (m *UserDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserDetail) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserDetail) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserDetail) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type UserCreateRequest struct {
	User                 *UserDetail `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserCreateRequest) Reset()         { *m = UserCreateRequest{} }
func (m *UserCreateRequest) String() string { return proto.CompactTextString(m) }
func (*UserCreateRequest) ProtoMessage()    {}
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateRequest.Unmarshal(m, b)
}
func (m *UserCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateRequest.Marshal(b, m, deterministic)
}
func (m *UserCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateRequest.Merge(m, src)
}
func (m *UserCreateRequest) XXX_Size() int {
	return xxx_messageInfo_UserCreateRequest.Size(m)
}
func (m *UserCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateRequest proto.InternalMessageInfo

func (m *UserCreateRequest) GetUser() *UserDetail {
	if m != nil {
		return m.User
	}
	return nil
}

type UserDetailRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDetailRequest) Reset()         { *m = UserDetailRequest{} }
func (m *UserDetailRequest) String() string { return proto.CompactTextString(m) }
func (*UserDetailRequest) ProtoMessage()    {}
func (*UserDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetailRequest.Unmarshal(m, b)
}
func (m *UserDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetailRequest.Marshal(b, m, deterministic)
}
func (m *UserDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetailRequest.Merge(m, src)
}
func (m *UserDetailRequest) XXX_Size() int {
	return xxx_messageInfo_UserDetailRequest.Size(m)
}
func (m *UserDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetailRequest proto.InternalMessageInfo

func (m *UserDetailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserListRequest struct {
	Filter               *UserListRequest_UserFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Count                int32                       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Limit                int32                       `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UserListRequest) Reset()         { *m = UserListRequest{} }
func (m *UserListRequest) String() string { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()    {}
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRequest.Unmarshal(m, b)
}
func (m *UserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRequest.Marshal(b, m, deterministic)
}
func (m *UserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRequest.Merge(m, src)
}
func (m *UserListRequest) XXX_Size() int {
	return xxx_messageInfo_UserListRequest.Size(m)
}
func (m *UserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRequest proto.InternalMessageInfo

func (m *UserListRequest) GetFilter() *UserListRequest_UserFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *UserListRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *UserListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type UserListRequest_UserFilter struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Name                 []string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListRequest_UserFilter) Reset()         { *m = UserListRequest_UserFilter{} }
func (m *UserListRequest_UserFilter) String() string { return proto.CompactTextString(m) }
func (*UserListRequest_UserFilter) ProtoMessage()    {}
func (*UserListRequest_UserFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3, 0}
}

func (m *UserListRequest_UserFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRequest_UserFilter.Unmarshal(m, b)
}
func (m *UserListRequest_UserFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRequest_UserFilter.Marshal(b, m, deterministic)
}
func (m *UserListRequest_UserFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRequest_UserFilter.Merge(m, src)
}
func (m *UserListRequest_UserFilter) XXX_Size() int {
	return xxx_messageInfo_UserListRequest_UserFilter.Size(m)
}
func (m *UserListRequest_UserFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRequest_UserFilter.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRequest_UserFilter proto.InternalMessageInfo

func (m *UserListRequest_UserFilter) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserListRequest_UserFilter) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserListRequest_UserFilter) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type UserListResponse struct {
	Users                []*UserDetail `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUsers() []*UserDetail {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserUpdateRequest struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *UserDetail `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserUpdateRequest) Reset()         { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()    {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateRequest.Unmarshal(m, b)
}
func (m *UserUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateRequest.Marshal(b, m, deterministic)
}
func (m *UserUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateRequest.Merge(m, src)
}
func (m *UserUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UserUpdateRequest.Size(m)
}
func (m *UserUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateRequest proto.InternalMessageInfo

func (m *UserUpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserUpdateRequest) GetUser() *UserDetail {
	if m != nil {
		return m.User
	}
	return nil
}

type UserDeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteRequest) Reset()         { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()    {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteRequest.Unmarshal(m, b)
}
func (m *UserDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteRequest.Marshal(b, m, deterministic)
}
func (m *UserDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteRequest.Merge(m, src)
}
func (m *UserDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_UserDeleteRequest.Size(m)
}
func (m *UserDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteRequest proto.InternalMessageInfo

func (m *UserDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDetail)(nil), "api.UserDetail")
	proto.RegisterType((*UserCreateRequest)(nil), "api.UserCreateRequest")
	proto.RegisterType((*UserDetailRequest)(nil), "api.UserDetailRequest")
	proto.RegisterType((*UserListRequest)(nil), "api.UserListRequest")
	proto.RegisterType((*UserListRequest_UserFilter)(nil), "api.UserListRequest.UserFilter")
	proto.RegisterType((*UserListResponse)(nil), "api.UserListResponse")
	proto.RegisterType((*UserUpdateRequest)(nil), "api.UserUpdateRequest")
	proto.RegisterType((*UserDeleteRequest)(nil), "api.UserDeleteRequest")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xc9, 0x5f, 0xec, 0x29, 0xb4, 0xbb, 0x87, 0x1a, 0x43, 0x28, 0x58, 0xa6, 0x08, 0xc5,
	0x8b, 0x0d, 0xd6, 0x0b, 0xff, 0x5c, 0xaa, 0x14, 0x15, 0x41, 0x08, 0xf4, 0x01, 0x46, 0x33, 0x2e,
	0x03, 0xd9, 0x4c, 0xcc, 0xcc, 0xee, 0x8d, 0x78, 0xa1, 0xaf, 0xe0, 0xdb, 0xf8, 0x1a, 0xbe, 0x82,
	0x0f, 0x22, 0x73, 0x26, 0x9b, 0x4d, 0xe2, 0x2e, 0xf4, 0x2e, 0xe7, 0xcc, 0x37, 0xdf, 0x7c, 0xe7,
	0x77, 0x76, 0x01, 0xd6, 0x5a, 0xb4, 0x8b, 0xa6, 0x55, 0x46, 0x61, 0xc0, 0x1b, 0x99, 0x9d, 0x2f,
	0x95, 0x5a, 0x56, 0x22, 0xe7, 0x8d, 0xcc, 0x79, 0x5d, 0x2b, 0xc3, 0x8d, 0x54, 0xb5, 0x76, 0x12,
	0xb6, 0x01, 0xb8, 0xd5, 0xa2, 0x7d, 0x23, 0x0c, 0x97, 0x15, 0x9e, 0x80, 0x2f, 0xcb, 0xd4, 0xbb,
	0xf0, 0xae, 0x8e, 0x0a, 0x5f, 0x96, 0x98, 0xc1, 0x3d, 0x6b, 0x57, 0xf3, 0x95, 0x48, 0x7d, 0xea,
	0xf6, 0x35, 0x22, 0x84, 0xd4, 0x0f, 0xa8, 0x4f, 0xdf, 0x38, 0x83, 0x80, 0x2f, 0x45, 0x1a, 0x5e,
	0x78, 0x57, 0x51, 0x61, 0x3f, 0x31, 0x81, 0x98, 0x6f, 0xb8, 0xe1, 0x6d, 0x1a, 0x91, 0xae, 0xab,
	0xd8, 0x73, 0x98, 0xdb, 0x77, 0x5f, 0xb7, 0x82, 0x1b, 0x51, 0x88, 0xaf, 0x6b, 0xa1, 0x0d, 0x5e,
	0x42, 0x68, 0xed, 0x29, 0xc0, 0xf1, 0xf5, 0xe9, 0x82, 0x37, 0x72, 0xb1, 0x4b, 0x57, 0xd0, 0x21,
	0xbb, 0x74, 0x37, 0xbb, 0x5e, 0x77, 0x73, 0x12, 0x9c, 0xfd, 0xf6, 0xe0, 0xd4, 0xaa, 0x3e, 0x48,
	0x6d, 0xb6, 0x9a, 0x67, 0x10, 0x7f, 0x91, 0x95, 0xe9, 0xfd, 0x1f, 0xf6, 0xfe, 0x03, 0x15, 0xd5,
	0x37, 0x24, 0x2b, 0x3a, 0x39, 0x9e, 0x41, 0xf4, 0x59, 0xad, 0x6b, 0x43, 0x08, 0xa2, 0xc2, 0x15,
	0xb6, 0x5b, 0xc9, 0x95, 0x34, 0x04, 0x20, 0x2a, 0x5c, 0x91, 0xbd, 0x72, 0x3c, 0x9d, 0x43, 0x1f,
	0x2b, 0xe8, 0x78, 0x6e, 0x99, 0xf9, 0xd4, 0x19, 0x31, 0x0b, 0x7a, 0x66, 0xec, 0x05, 0xcc, 0x76,
	0xa9, 0x74, 0xa3, 0x6a, 0x2d, 0xf0, 0x11, 0x44, 0x76, 0x7a, 0x4d, 0x66, 0x7b, 0xd8, 0xb8, 0x53,
	0xf6, 0xd6, 0xc1, 0xb9, 0x6d, 0xca, 0x01, 0xd6, 0xe9, 0x56, 0xb7, 0x98, 0xfd, 0x3b, 0x61, 0xae,
	0xc4, 0x41, 0xa7, 0xeb, 0x1f, 0x01, 0x84, 0x56, 0x85, 0xef, 0x21, 0x76, 0xab, 0xc4, 0xa4, 0xb7,
	0x1b, 0xed, 0x36, 0x9b, 0x3e, 0xc3, 0x1e, 0xfc, 0xfc, 0xf3, 0xf7, 0x97, 0x3f, 0x67, 0x47, 0xf9,
	0xe6, 0x49, 0x4e, 0xe9, 0x5f, 0xd2, 0xcb, 0xf8, 0x0e, 0xe2, 0xee, 0xe7, 0x98, 0x4c, 0xa3, 0x1d,
	0xf2, 0x4a, 0xc8, 0x6b, 0x86, 0x27, 0xbd, 0x57, 0xfe, 0x4d, 0x96, 0xdf, 0xf1, 0x06, 0x42, 0x4b,
	0x11, 0xcf, 0xf6, 0xad, 0x3a, 0xbb, 0x3f, 0xe9, 0x3a, 0xd4, 0x6c, 0x4e, 0x66, 0xc7, 0xb8, 0x0b,
	0x86, 0x1f, 0x21, 0x76, 0x48, 0x07, 0x91, 0x46, 0x8c, 0xff, 0x8f, 0x74, 0x4e, 0x2e, 0x49, 0x36,
	0x89, 0x34, 0x9c, 0xd1, 0x92, 0x1d, 0xcd, 0x38, 0x40, 0x7d, 0x70, 0xc6, 0xc7, 0x13, 0xc3, 0x4f,
	0x31, 0xfd, 0x91, 0x9f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x30, 0x9f, 0xde, 0xf9, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserDetail, error)
	Detail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetail, error)
	List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserDetail, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDetail, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/api.User/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Detail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/api.User/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/api.User/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/api.User/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/api.User/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Create(context.Context, *UserCreateRequest) (*UserDetail, error)
	Detail(context.Context, *UserDetailRequest) (*UserDetail, error)
	List(context.Context, *UserListRequest) (*UserListResponse, error)
	Update(context.Context, *UserUpdateRequest) (*UserDetail, error)
	Delete(context.Context, *UserDeleteRequest) (*UserDetail, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Create(ctx context.Context, req *UserCreateRequest) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServer) Detail(ctx context.Context, req *UserDetailRequest) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (*UnimplementedUserServer) List(ctx context.Context, req *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedUserServer) Update(ctx context.Context, req *UserUpdateRequest) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServer) Delete(ctx context.Context, req *UserDeleteRequest) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Detail(ctx, req.(*UserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).List(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _User_Detail_Handler,
		},
		{
			MethodName: "List",
			Handler:    _User_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

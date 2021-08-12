package protocol

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4d180e294235fa, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ListUserRequestType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserRequestType) Reset()         { *m = ListUserRequestType{} }
func (m *ListUserRequestType) String() string { return proto.CompactTextString(m) }
func (*ListUserRequestType) ProtoMessage()    {}
func (*ListUserRequestType) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4d180e294235fa, []int{1}
}

func (m *ListUserRequestType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserRequestType.Unmarshal(m, b)
}
func (m *ListUserRequestType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserRequestType.Marshal(b, m, deterministic)
}
func (m *ListUserRequestType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserRequestType.Merge(m, src)
}
func (m *ListUserRequestType) XXX_Size() int {
	return xxx_messageInfo_ListUserRequestType.Size(m)
}
func (m *ListUserRequestType) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserRequestType.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserRequestType proto.InternalMessageInfo

type ListUserResponseType struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserResponseType) Reset()         { *m = ListUserResponseType{} }
func (m *ListUserResponseType) String() string { return proto.CompactTextString(m) }
func (*ListUserResponseType) ProtoMessage()    {}
func (*ListUserResponseType) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4d180e294235fa, []int{2}
}

func (m *ListUserResponseType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserResponseType.Unmarshal(m, b)
}
func (m *ListUserResponseType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserResponseType.Marshal(b, m, deterministic)
}
func (m *ListUserResponseType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserResponseType.Merge(m, src)
}
func (m *ListUserResponseType) XXX_Size() int {
	return xxx_messageInfo_ListUserResponseType.Size(m)
}
func (m *ListUserResponseType) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserResponseType.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserResponseType proto.InternalMessageInfo

func (m *ListUserResponseType) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type RegisterUserResponseType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserResponseType) Reset()         { *m = RegisterUserResponseType{} }
func (m *RegisterUserResponseType) String() string { return proto.CompactTextString(m) }
func (*RegisterUserResponseType) ProtoMessage()    {}
func (*RegisterUserResponseType) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4d180e294235fa, []int{3}
}

func (m *RegisterUserResponseType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserResponseType.Unmarshal(m, b)
}
func (m *RegisterUserResponseType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserResponseType.Marshal(b, m, deterministic)
}
func (m *RegisterUserResponseType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserResponseType.Merge(m, src)
}
func (m *RegisterUserResponseType) XXX_Size() int {
	return xxx_messageInfo_RegisterUserResponseType.Size(m)
}
func (m *RegisterUserResponseType) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserResponseType.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserResponseType proto.InternalMessageInfo

type RegisterUserRequestType struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserRequestType) Reset()         { *m = RegisterUserRequestType{} }
func (m *RegisterUserRequestType) String() string { return proto.CompactTextString(m) }
func (*RegisterUserRequestType) ProtoMessage()    {}
func (*RegisterUserRequestType) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc4d180e294235fa, []int{4}
}

func (m *RegisterUserRequestType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserRequestType.Unmarshal(m, b)
}
func (m *RegisterUserRequestType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserRequestType.Marshal(b, m, deterministic)
}
func (m *RegisterUserRequestType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserRequestType.Merge(m, src)
}
func (m *RegisterUserRequestType) XXX_Size() int {
	return xxx_messageInfo_RegisterUserRequestType.Size(m)
}
func (m *RegisterUserRequestType) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserRequestType.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserRequestType proto.InternalMessageInfo

func (m *RegisterUserRequestType) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "protocol.User")
	proto.RegisterType((*ListUserRequestType)(nil), "protocol.ListUserRequestType")
	proto.RegisterType((*ListUserResponseType)(nil), "protocol.ListUserResponseType")
	proto.RegisterType((*RegisterUserResponseType)(nil), "protocol.RegisterUserResponseType")
	proto.RegisterType((*RegisterUserRequestType)(nil), "protocol.RegisterUserRequestType")
}

func init() {
	proto.RegisterFile("app/interface/rpc/v1.0/protocol/user_service.proto", fileDescriptor_bc4d180e294235fa)
}

var fileDescriptor_bc4d180e294235fa = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0x68, 0xa5, 0x4e, 0xa5, 0x87, 0xb1, 0x62, 0x08, 0x28, 0x75, 0xf1, 0xd0, 0x83,
	0x64, 0x35, 0x5e, 0xfd, 0x07, 0x7a, 0x8a, 0x8a, 0x47, 0x89, 0xe9, 0x28, 0x0b, 0xb5, 0x59, 0x77,
	0xb6, 0x05, 0x7f, 0x97, 0x7f, 0x50, 0xb2, 0x21, 0xee, 0x2a, 0xed, 0x71, 0xde, 0x7b, 0x33, 0xef,
	0x1b, 0x28, 0x2a, 0x63, 0x94, 0x5e, 0x3a, 0xb2, 0x6f, 0x55, 0x4d, 0xca, 0x9a, 0x5a, 0xad, 0xaf,
	0xf3, 0x2b, 0x65, 0x6c, 0xe3, 0x9a, 0xba, 0x59, 0xa8, 0x15, 0x93, 0x7d, 0x61, 0xb2, 0x6b, 0x5d,
	0x53, 0xee, 0x55, 0x1c, 0xf6, 0xa6, 0xbc, 0x84, 0xbd, 0x27, 0x26, 0x8b, 0x63, 0x48, 0xf4, 0x3c,
	0x15, 0x53, 0x31, 0x3b, 0x28, 0x13, 0x3d, 0xc7, 0x09, 0x0c, 0xe8, 0xa3, 0xd2, 0x8b, 0x34, 0xf1,
	0x52, 0x37, 0xc8, 0x63, 0x38, 0xba, 0xd7, 0xec, 0xda, 0x8d, 0x92, 0x3e, 0x57, 0xc4, 0xee, 0xf1,
	0xcb, 0x90, 0xbc, 0x85, 0x49, 0x90, 0xd9, 0x34, 0x4b, 0xa6, 0x56, 0xc7, 0x0b, 0x18, 0xb4, 0xe5,
	0x9c, 0x8a, 0xe9, 0xee, 0x6c, 0x54, 0x8c, 0xf3, 0xbe, 0x36, 0xf7, 0xd1, 0xce, 0x94, 0x19, 0xa4,
	0x25, 0xbd, 0x6b, 0x76, 0x64, 0xff, 0x5f, 0x90, 0x0a, 0x4e, 0xfe, 0x7a, 0xbf, 0xa5, 0x81, 0x50,
	0x44, 0x84, 0xc5, 0xb7, 0x80, 0x51, 0x9b, 0x7c, 0xe8, 0xfe, 0xc5, 0x3b, 0x18, 0xf6, 0x68, 0x78,
	0x1a, 0xfa, 0x37, 0x7c, 0x91, 0x9d, 0x6d, 0xb2, 0x23, 0x96, 0x1d, 0x7c, 0x86, 0xc3, 0x98, 0x06,
	0xcf, 0xc3, 0xc6, 0x16, 0xca, 0x4c, 0x6e, 0x8b, 0xc4, 0x87, 0x5f, 0xf7, 0x7d, 0xe8, 0xe6, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x98, 0x12, 0x32, 0xf7, 0xcc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	ListUser(ctx context.Context, in *ListUserRequestType, opts ...grpc.CallOption) (*ListUserResponseType, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequestType, opts ...grpc.CallOption) (*RegisterUserResponseType, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListUser(ctx context.Context, in *ListUserRequestType, opts ...grpc.CallOption) (*ListUserResponseType, error) {
	out := new(ListUserResponseType)
	err := c.cc.Invoke(ctx, "/protocol.UserService/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequestType, opts ...grpc.CallOption) (*RegisterUserResponseType, error) {
	out := new(RegisterUserResponseType)
	err := c.cc.Invoke(ctx, "/protocol.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	ListUser(context.Context, *ListUserRequestType) (*ListUserResponseType, error)
	RegisterUser(context.Context, *RegisterUserRequestType) (*RegisterUserResponseType, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequestType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UserService/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUser(ctx, req.(*ListUserRequestType))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequestType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegisterUserRequestType))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _UserService_ListUser_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/interface/rpc/v1.0/protocol/user_service.proto",
}

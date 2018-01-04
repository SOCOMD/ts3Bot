// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ts3Bot.proto

/*
Package ts3Bot is a generated protocol buffer package.

It is generated from these files:
	ts3Bot.proto

It has these top-level messages:
	Nil
	User
	UserList
	ServerGroup
	ServerGroupList
	UserAndGroup
*/
package ts3Bot

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Nil struct {
}

func (m *Nil) Reset()                    { *m = Nil{} }
func (m *Nil) String() string            { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()               {}
func (*Nil) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User struct {
	Dbid          string `protobuf:"bytes,1,opt,name=dbid" json:"dbid,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Uuid          string `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	Created       string `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	Lastconnected string `protobuf:"bytes,5,opt,name=lastconnected" json:"lastconnected,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetDbid() string {
	if m != nil {
		return m.Dbid
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *User) GetLastconnected() string {
	if m != nil {
		return m.Lastconnected
	}
	return ""
}

type UserList struct {
	Users []*User `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *UserList) Reset()                    { *m = UserList{} }
func (m *UserList) String() string            { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()               {}
func (*UserList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserList) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ServerGroup struct {
	Sgid string `protobuf:"bytes,1,opt,name=sgid" json:"sgid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ServerGroup) Reset()                    { *m = ServerGroup{} }
func (m *ServerGroup) String() string            { return proto.CompactTextString(m) }
func (*ServerGroup) ProtoMessage()               {}
func (*ServerGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServerGroup) GetSgid() string {
	if m != nil {
		return m.Sgid
	}
	return ""
}

func (m *ServerGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ServerGroupList struct {
	Groups []*ServerGroup `protobuf:"bytes,1,rep,name=Groups" json:"Groups,omitempty"`
}

func (m *ServerGroupList) Reset()                    { *m = ServerGroupList{} }
func (m *ServerGroupList) String() string            { return proto.CompactTextString(m) }
func (*ServerGroupList) ProtoMessage()               {}
func (*ServerGroupList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServerGroupList) GetGroups() []*ServerGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

type UserAndGroup struct {
	User  *User        `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Group *ServerGroup `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
}

func (m *UserAndGroup) Reset()                    { *m = UserAndGroup{} }
func (m *UserAndGroup) String() string            { return proto.CompactTextString(m) }
func (*UserAndGroup) ProtoMessage()               {}
func (*UserAndGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserAndGroup) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserAndGroup) GetGroup() *ServerGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*Nil)(nil), "Nil")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*UserList)(nil), "UserList")
	proto.RegisterType((*ServerGroup)(nil), "ServerGroup")
	proto.RegisterType((*ServerGroupList)(nil), "ServerGroupList")
	proto.RegisterType((*UserAndGroup)(nil), "UserAndGroup")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ts3Bot service

type Ts3BotClient interface {
	GetUsers(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*UserList, error)
	GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	ClientList(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*UserList, error)
	GetServerGroups(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*ServerGroupList, error)
	GetUsersInGroup(ctx context.Context, in *ServerGroup, opts ...grpc.CallOption) (*UserList, error)
	AddUserToGroup(ctx context.Context, in *UserAndGroup, opts ...grpc.CallOption) (*Nil, error)
	DelUserFromGroup(ctx context.Context, in *UserAndGroup, opts ...grpc.CallOption) (*Nil, error)
}

type ts3BotClient struct {
	cc *grpc.ClientConn
}

func NewTs3BotClient(cc *grpc.ClientConn) Ts3BotClient {
	return &ts3BotClient{cc}
}

func (c *ts3BotClient) GetUsers(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := grpc.Invoke(ctx, "/ts3bot/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ts3BotClient) GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/ts3bot/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ts3BotClient) ClientList(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := grpc.Invoke(ctx, "/ts3bot/ClientList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ts3BotClient) GetServerGroups(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*ServerGroupList, error) {
	out := new(ServerGroupList)
	err := grpc.Invoke(ctx, "/ts3bot/GetServerGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ts3BotClient) GetUsersInGroup(ctx context.Context, in *ServerGroup, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := grpc.Invoke(ctx, "/ts3bot/GetUsersInGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ts3BotClient) AddUserToGroup(ctx context.Context, in *UserAndGroup, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := grpc.Invoke(ctx, "/ts3bot/AddUserToGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ts3BotClient) DelUserFromGroup(ctx context.Context, in *UserAndGroup, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := grpc.Invoke(ctx, "/ts3bot/DelUserFromGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ts3Bot service

type Ts3BotServer interface {
	GetUsers(context.Context, *Nil) (*UserList, error)
	GetUser(context.Context, *User) (*User, error)
	ClientList(context.Context, *Nil) (*UserList, error)
	GetServerGroups(context.Context, *Nil) (*ServerGroupList, error)
	GetUsersInGroup(context.Context, *ServerGroup) (*UserList, error)
	AddUserToGroup(context.Context, *UserAndGroup) (*Nil, error)
	DelUserFromGroup(context.Context, *UserAndGroup) (*Nil, error)
}

func RegisterTs3BotServer(s *grpc.Server, srv Ts3BotServer) {
	s.RegisterService(&_Ts3Bot_serviceDesc, srv)
}

func _Ts3Bot_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ts3BotServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts3bot/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ts3BotServer).GetUsers(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ts3Bot_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ts3BotServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts3bot/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ts3BotServer).GetUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ts3Bot_ClientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ts3BotServer).ClientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts3bot/ClientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ts3BotServer).ClientList(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ts3Bot_GetServerGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ts3BotServer).GetServerGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts3bot/GetServerGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ts3BotServer).GetServerGroups(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ts3Bot_GetUsersInGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ts3BotServer).GetUsersInGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts3bot/GetUsersInGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ts3BotServer).GetUsersInGroup(ctx, req.(*ServerGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ts3Bot_AddUserToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAndGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ts3BotServer).AddUserToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts3bot/AddUserToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ts3BotServer).AddUserToGroup(ctx, req.(*UserAndGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ts3Bot_DelUserFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAndGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ts3BotServer).DelUserFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts3bot/DelUserFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ts3BotServer).DelUserFromGroup(ctx, req.(*UserAndGroup))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ts3Bot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ts3bot",
	HandlerType: (*Ts3BotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _Ts3Bot_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Ts3Bot_GetUser_Handler,
		},
		{
			MethodName: "ClientList",
			Handler:    _Ts3Bot_ClientList_Handler,
		},
		{
			MethodName: "GetServerGroups",
			Handler:    _Ts3Bot_GetServerGroups_Handler,
		},
		{
			MethodName: "GetUsersInGroup",
			Handler:    _Ts3Bot_GetUsersInGroup_Handler,
		},
		{
			MethodName: "AddUserToGroup",
			Handler:    _Ts3Bot_AddUserToGroup_Handler,
		},
		{
			MethodName: "DelUserFromGroup",
			Handler:    _Ts3Bot_DelUserFromGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ts3Bot.proto",
}

func init() { proto.RegisterFile("ts3Bot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x4f, 0x2a, 0x31,
	0x14, 0xcd, 0xc0, 0x0c, 0x1f, 0x17, 0x78, 0x90, 0x6e, 0xde, 0x3c, 0x5e, 0x4c, 0x48, 0x83, 0x01,
	0x5d, 0x74, 0x01, 0x31, 0xae, 0x51, 0x23, 0x31, 0x51, 0x16, 0xa8, 0x3f, 0x00, 0x68, 0x43, 0x26,
	0x19, 0xa6, 0xa4, 0xed, 0xf8, 0x03, 0x8c, 0x3f, 0xdc, 0xdc, 0x5b, 0xc0, 0x01, 0x75, 0x35, 0xe7,
	0x9e, 0x73, 0xda, 0x73, 0x6e, 0x33, 0xd0, 0x74, 0x76, 0x7c, 0xa3, 0x9d, 0xd8, 0x1a, 0xed, 0x34,
	0x8f, 0xa0, 0x3c, 0x4b, 0x52, 0xfe, 0x1e, 0x40, 0xf8, 0x6a, 0x95, 0x61, 0x0c, 0x42, 0xb9, 0x4c,
	0x64, 0x1c, 0xf4, 0x82, 0x61, 0x7d, 0x4e, 0x18, 0xb9, 0x6c, 0xb1, 0x51, 0x71, 0xc9, 0x73, 0x88,
	0x91, 0xcb, 0xf3, 0x44, 0xc6, 0x65, 0xcf, 0x21, 0x66, 0x31, 0x54, 0x57, 0x46, 0x2d, 0x9c, 0x92,
	0x71, 0x48, 0xf4, 0x7e, 0x64, 0x7d, 0x68, 0xa5, 0x0b, 0xeb, 0x56, 0x3a, 0xcb, 0xd4, 0x0a, 0xf5,
	0x88, 0xf4, 0x63, 0x92, 0x0f, 0xa0, 0x86, 0x1d, 0x1e, 0x13, 0xeb, 0xd8, 0x7f, 0x88, 0x10, 0xdb,
	0x38, 0xe8, 0x95, 0x87, 0x8d, 0x51, 0x24, 0x70, 0x9a, 0x7b, 0x8e, 0x5f, 0x41, 0xe3, 0x59, 0x99,
	0x37, 0x65, 0xa6, 0x46, 0xe7, 0x5b, 0xec, 0x62, 0xd7, 0x5f, 0x9d, 0x11, 0xff, 0xd4, 0x99, 0x5f,
	0x43, 0xbb, 0x70, 0x8c, 0x62, 0xfa, 0x50, 0xa1, 0x61, 0x9f, 0xd3, 0x14, 0x05, 0xc7, 0x7c, 0xa7,
	0xf1, 0x27, 0x68, 0x62, 0xf0, 0x24, 0x93, 0x3e, 0xf0, 0x1f, 0x84, 0xb9, 0x55, 0x86, 0x02, 0x0f,
	0xdd, 0x88, 0x62, 0x1c, 0xa2, 0x35, 0x7a, 0x28, 0xf8, 0xf4, 0x3e, 0x2f, 0x8d, 0x3e, 0x4a, 0x50,
	0x71, 0x76, 0xbc, 0xd4, 0xb8, 0x66, 0x6d, 0xaa, 0x1c, 0x6d, 0xc5, 0x42, 0x31, 0x4b, 0xd2, 0x6e,
	0x5d, 0x1c, 0xde, 0xe0, 0x2f, 0x54, 0x77, 0x22, 0xf3, 0x19, 0x5d, 0xff, 0x61, 0x67, 0x00, 0xb7,
	0x69, 0xa2, 0x32, 0x47, 0xb6, 0x6f, 0xe7, 0x2e, 0xa0, 0x3d, 0x55, 0xae, 0x10, 0xbc, 0xbf, 0xbb,
	0x23, 0x4e, 0xf7, 0xbf, 0x24, 0x2b, 0xe5, 0x3f, 0x64, 0x7e, 0xb9, 0xa3, 0xca, 0xc5, 0x6b, 0xcf,
	0xe1, 0xcf, 0x44, 0x4a, 0x1c, 0x5f, 0xb4, 0xb7, 0xb6, 0x44, 0xf1, 0x59, 0xba, 0x14, 0xc2, 0x06,
	0xd0, 0xb9, 0x53, 0x29, 0x0a, 0xf7, 0x46, 0x6f, 0x7e, 0x37, 0x2e, 0x2b, 0xf4, 0x07, 0x8e, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x60, 0x0e, 0xcd, 0x91, 0x02, 0x00, 0x00,
}

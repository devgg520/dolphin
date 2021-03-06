// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// MessageMail defined
type MessageMail struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageMail) Reset()         { *m = MessageMail{} }
func (m *MessageMail) String() string { return proto.CompactTextString(m) }
func (*MessageMail) ProtoMessage()    {}
func (*MessageMail) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *MessageMail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageMail.Unmarshal(m, b)
}
func (m *MessageMail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageMail.Marshal(b, m, deterministic)
}
func (m *MessageMail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageMail.Merge(m, src)
}
func (m *MessageMail) XXX_Size() int {
	return xxx_messageInfo_MessageMail.Size(m)
}
func (m *MessageMail) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageMail.DiscardUnknown(m)
}

var xxx_messageInfo_MessageMail proto.InternalMessageInfo

// MessageReply defined
type MessageReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReply) Reset()         { *m = MessageReply{} }
func (m *MessageReply) String() string { return proto.CompactTextString(m) }
func (*MessageReply) ProtoMessage()    {}
func (*MessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *MessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageReply.Unmarshal(m, b)
}
func (m *MessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageReply.Marshal(b, m, deterministic)
}
func (m *MessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReply.Merge(m, src)
}
func (m *MessageReply) XXX_Size() int {
	return xxx_messageInfo_MessageReply.Size(m)
}
func (m *MessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MessageMail)(nil), "proto.MessageMail")
	proto.RegisterType((*MessageReply)(nil), "proto.MessageReply")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xbc, 0x5c,
	0xdc, 0xbe, 0x10, 0x71, 0xdf, 0xc4, 0xcc, 0x1c, 0x25, 0x3e, 0x2e, 0x1e, 0x28, 0x37, 0x28, 0xb5,
	0x20, 0xa7, 0xd2, 0xc8, 0x99, 0x8b, 0x0b, 0xca, 0x0f, 0x2e, 0x2a, 0x13, 0x32, 0xe5, 0xe2, 0x08,
	0x4e, 0xcd, 0x4b, 0x01, 0xa9, 0x14, 0x12, 0x82, 0x98, 0xa3, 0x87, 0xa4, 0x5b, 0x4a, 0x18, 0x55,
	0x0c, 0x6c, 0x84, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xda,
	0x2f, 0x00, 0xfc, 0x82, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageSrvClient is the client API for MessageSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageSrvClient interface {
	SendMail(ctx context.Context, in *MessageMail, opts ...grpc.CallOption) (*MessageReply, error)
}

type messageSrvClient struct {
	cc *grpc.ClientConn
}

func NewMessageSrvClient(cc *grpc.ClientConn) MessageSrvClient {
	return &messageSrvClient{cc}
}

func (c *messageSrvClient) SendMail(ctx context.Context, in *MessageMail, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/proto.MessageSrv/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageSrvServer is the server API for MessageSrv service.
type MessageSrvServer interface {
	SendMail(context.Context, *MessageMail) (*MessageReply, error)
}

// UnimplementedMessageSrvServer can be embedded to have forward compatible implementations.
type UnimplementedMessageSrvServer struct {
}

func (*UnimplementedMessageSrvServer) SendMail(ctx context.Context, req *MessageMail) (*MessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}

func RegisterMessageSrvServer(s *grpc.Server, srv MessageSrvServer) {
	s.RegisterService(&_MessageSrv_serviceDesc, srv)
}

func _MessageSrv_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageMail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSrvServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageSrv/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSrvServer).SendMail(ctx, req.(*MessageMail))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessageSrv",
	HandlerType: (*MessageSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _MessageSrv_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

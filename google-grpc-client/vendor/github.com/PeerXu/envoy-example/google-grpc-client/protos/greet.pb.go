// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

/*
Package greet is a generated protocol buffer package.

It is generated from these files:
	greet.proto

It has these top-level messages:
	GreetRequest
	GreetResponse
*/
package greet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type GreetRequest struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GreetRequest) Reset()                    { *m = GreetRequest{} }
func (m *GreetRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()               {}
func (*GreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GreetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GreetResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GreetResponse) Reset()                    { *m = GreetResponse{} }
func (m *GreetResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()               {}
func (*GreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GreetService service

type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := grpc.Invoke(ctx, "/greet.GreetService/Greet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GreetService service

type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x64, 0xd2, 0xf3, 0xf3,
	0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x8a, 0x21, 0x8a, 0x94, 0x54, 0xb8, 0x78, 0xdc, 0x41, 0xca, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x55, 0x2e, 0x5e, 0xa8, 0xaa, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0xec, 0xca, 0x8c, 0x42, 0xa1, 0x86, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xb9,
	0x72, 0xb1, 0x82, 0xf9, 0x42, 0xc2, 0x7a, 0x10, 0x87, 0x21, 0x5b, 0x25, 0x25, 0x82, 0x2a, 0x08,
	0x31, 0x59, 0x49, 0xb0, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xdc, 0x4a, 0x6c, 0xfa, 0x60, 0x59, 0x2b,
	0x46, 0xad, 0x24, 0x36, 0xb0, 0x53, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa5, 0x02,
	0xa6, 0xde, 0x00, 0x00, 0x00,
}
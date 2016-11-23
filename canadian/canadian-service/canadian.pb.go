// Code generated by protoc-gen-go.
// source: canadian.proto
// DO NOT EDIT!

/*
Package canadian is a generated protocol buffer package.

Expose an api to translate english to canadian

It is generated from these files:
	canadian.proto

It has these top-level messages:
	TranslateRequest
	TranslateResponse
*/
package canadian

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"

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

type TranslateRequest struct {
	Phrase string `protobuf:"bytes,1,opt,name=Phrase" json:"Phrase,omitempty"`
}

func (m *TranslateRequest) Reset()                    { *m = TranslateRequest{} }
func (m *TranslateRequest) String() string            { return proto.CompactTextString(m) }
func (*TranslateRequest) ProtoMessage()               {}
func (*TranslateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TranslateRequest) GetPhrase() string {
	if m != nil {
		return m.Phrase
	}
	return ""
}

type TranslateResponse struct {
	Value string `protobuf:"bytes,1,opt,name=Value" json:"Value,omitempty"`
}

func (m *TranslateResponse) Reset()                    { *m = TranslateResponse{} }
func (m *TranslateResponse) String() string            { return proto.CompactTextString(m) }
func (*TranslateResponse) ProtoMessage()               {}
func (*TranslateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TranslateResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*TranslateRequest)(nil), "canadian.TranslateRequest")
	proto.RegisterType((*TranslateResponse)(nil), "canadian.TranslateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Canadian service

type CanadianClient interface {
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
}

type canadianClient struct {
	cc *grpc.ClientConn
}

func NewCanadianClient(cc *grpc.ClientConn) CanadianClient {
	return &canadianClient{cc}
}

func (c *canadianClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	out := new(TranslateResponse)
	err := grpc.Invoke(ctx, "/canadian.Canadian/Translate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Canadian service

type CanadianServer interface {
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
}

func RegisterCanadianServer(s *grpc.Server, srv CanadianServer) {
	s.RegisterService(&_Canadian_serviceDesc, srv)
}

func _Canadian_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanadianServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/canadian.Canadian/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanadianServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Canadian_serviceDesc = grpc.ServiceDesc{
	ServiceName: "canadian.Canadian",
	HandlerType: (*CanadianServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _Canadian_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "canadian.proto",
}

func init() { proto.RegisterFile("canadian.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4e, 0xcc, 0x4b,
	0x4c, 0xc9, 0x4c, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x3c,
	0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xd2, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b,
	0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0xaa, 0xf4, 0x21, 0x52, 0x89, 0x05, 0x99, 0xc5, 0xfa, 0x89,
	0x05, 0x99, 0xfa, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0xc9, 0xf9, 0x79, 0x69, 0x99, 0xe9,
	0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x43, 0x95, 0xb4,
	0xb8, 0x04, 0x42, 0x8a, 0x12, 0xf3, 0x8a, 0x73, 0x12, 0x4b, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0x02, 0x32, 0x8a, 0x12, 0x8b, 0x53, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xa0, 0x3c, 0x25, 0x4d, 0x2e, 0x41, 0x24, 0xb5, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0x22, 0x5c, 0xac, 0x61, 0x89, 0x39, 0xa5, 0x30, 0xb5, 0x10, 0x8e, 0x51, 0x1a, 0x17,
	0x87, 0x33, 0xd4, 0xb5, 0x42, 0x51, 0x5c, 0x9c, 0x70, 0x6d, 0x42, 0x52, 0x7a, 0x70, 0x5f, 0xa1,
	0xdb, 0x2b, 0x25, 0x8d, 0x55, 0x0e, 0x62, 0x8f, 0x92, 0x48, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xf8,
	0x84, 0x38, 0xf5, 0x61, 0x8a, 0xac, 0x18, 0xb5, 0x92, 0xd8, 0xc0, 0xbe, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x9d, 0x36, 0xb2, 0x2c, 0x01, 0x00, 0x00,
}
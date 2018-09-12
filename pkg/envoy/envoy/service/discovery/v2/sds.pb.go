// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/sds.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type SdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SdsDummy) Reset()         { *m = SdsDummy{} }
func (m *SdsDummy) String() string { return proto.CompactTextString(m) }
func (*SdsDummy) ProtoMessage()    {}
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2a4da2e99d9a3e6, []int{0}
}
func (m *SdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SdsDummy.Unmarshal(m, b)
}
func (m *SdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SdsDummy.Marshal(b, m, deterministic)
}
func (dst *SdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdsDummy.Merge(dst, src)
}
func (m *SdsDummy) XXX_Size() int {
	return xxx_messageInfo_SdsDummy.Size(m)
}
func (m *SdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SdsDummy)(nil), "envoy.service.discovery.v2.SdsDummy")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretDiscoveryServiceClient is the client API for SecretDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretDiscoveryServiceClient interface {
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v2.SecretDiscoveryService/FetchSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretDiscoveryServiceServer is the server API for SecretDiscoveryService service.
type SecretDiscoveryServiceServer interface {
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/sds.proto",
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/sds.proto", fileDescriptor_f2a4da2e99d9a3e6)
}

var fileDescriptor_f2a4da2e99d9a3e6 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f,
	0x4b, 0x2d, 0xaa, 0xd4, 0x2f, 0x33, 0xd2, 0x2f, 0x4e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x02, 0xab, 0xd2, 0x83, 0xaa, 0xd2, 0x83, 0xab, 0xd2, 0x2b, 0x33, 0x92, 0x92, 0x81,
	0x98, 0x90, 0x58, 0x90, 0x09, 0xd2, 0x83, 0x90, 0x02, 0xeb, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0x05, 0x4b, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x41, 0xcd,
	0x55, 0xe2, 0xe2, 0xe2, 0x08, 0x4e, 0x29, 0x76, 0x29, 0xcd, 0xcd, 0xad, 0x34, 0x7a, 0xcc, 0xc8,
	0x25, 0x16, 0x9c, 0x9a, 0x5c, 0x94, 0x5a, 0xe2, 0x02, 0x33, 0x23, 0x18, 0x62, 0x9f, 0x50, 0x18,
	0x17, 0x6f, 0x70, 0x49, 0x51, 0x6a, 0x62, 0x2e, 0x44, 0xbe, 0x58, 0x48, 0x4e, 0x0f, 0xe2, 0xa0,
	0xc4, 0x82, 0x4c, 0xbd, 0x32, 0x23, 0x3d, 0xb8, 0x86, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x29, 0x79, 0x9c, 0xf2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c,
	0x42, 0x85, 0x5c, 0x3c, 0x6e, 0xa9, 0x25, 0xc9, 0x19, 0x54, 0x33, 0x56, 0xa1, 0xe9, 0xf2, 0x93,
	0xc9, 0x4c, 0x52, 0x4a, 0xa2, 0x28, 0x41, 0x61, 0x55, 0x0c, 0x31, 0xdf, 0x8a, 0x51, 0x2b, 0x89,
	0x0d, 0xec, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x17, 0x83, 0x2f, 0x78, 0x01,
	0x00, 0x00,
}

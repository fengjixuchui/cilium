// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ext_authz/v2alpha/ext_authz.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// External Authorization filter calls out to an external service over either:
//
//  1. gRPC Authorization API defined by :ref:`CheckRequest
//     <envoy_api_msg_service.auth.v2alpha.CheckRequest>`.
//  2. Raw HTTP Authorization server by passing the request headers to the service.
//
// A failed check will cause this filter to close the HTTP request normally with 403 (Forbidden),
// unless a different status code has been indicated in the authorization response.
type ExtAuthz struct {
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services isExtAuthz_Services `protobuf_oneof:"services"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. When it is set to true, Envoy will also allow traffic in case of
	// communication failure between authorization service and the proxy.
	// Defaults to false.
	FailureModeAllow     bool     `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_971cd043e21bda4a, []int{0}
}
func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (dst *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(dst, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}

func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExtAuthz_OneofMarshaler, _ExtAuthz_OneofUnmarshaler, _ExtAuthz_OneofSizer, []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

func _ExtAuthz_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExtAuthz)
	// services
	switch x := m.Services.(type) {
	case *ExtAuthz_GrpcService:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GrpcService); err != nil {
			return err
		}
	case *ExtAuthz_HttpService:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpService); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ExtAuthz.Services has unexpected type %T", x)
	}
	return nil
}

func _ExtAuthz_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExtAuthz)
	switch tag {
	case 1: // services.grpc_service
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.GrpcService)
		err := b.DecodeMessage(msg)
		m.Services = &ExtAuthz_GrpcService{msg}
		return true, err
	case 3: // services.http_service
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpService)
		err := b.DecodeMessage(msg)
		m.Services = &ExtAuthz_HttpService{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ExtAuthz_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExtAuthz)
	// services
	switch x := m.Services.(type) {
	case *ExtAuthz_GrpcService:
		s := proto.Size(x.GrpcService)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtAuthz_HttpService:
		s := proto.Size(x.HttpService)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// External Authorization filter calls out to an upstream authorization server by passing the raw
// HTTP request headers to the server. This allows the authorization service to take a decision
// whether the request is authorized or not.
//
// A successful check allows the authorization service adding or overriding headers from the
// original request before dispatching it to the upstream. This is done by configuring which headers
// in the authorization response should be sent to the upstream. See *allowed_authorization_headers*
// bellow.
//
// A failed check will cause this filter to close the HTTP request normally with 403 (Forbidden),
// unless a different status code has been indicated by the authorization server via response
// headers. If other headers in the authorization response need to be sent to client, this can also
// be done by specifying them in *allowed_authorization_headers*.
type HttpService struct {
	// Sets the HTTP server URI which the authorization requests must be sent to.
	ServerUri *core.HttpUri `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	// Sets an optional prefix to the value of authorization request header *Path*.
	PathPrefix string `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	// Sets a list of headers that can be sent from the authorization server to the upstream service,
	// or to the downstream client when present in the authorization response. Note that a matched
	// request header will have its value overridden by the ones sent from the authorization server.
	AllowedAuthorizationHeaders []string `protobuf:"bytes,4,rep,name=allowed_authorization_headers,json=allowedAuthorizationHeaders,proto3" json:"allowed_authorization_headers,omitempty"`
	// Sets a list of headers that should be sent *from the filter* to the authorization server
	// when they are also present in the client request. Note that *Content-Length*, *Authority*,
	// *Method* and *Path* are always dispatched to the authorization server by default. The message
	// will not contain body data and the *Content-Length* will be set to zero.
	AllowedRequestHeaders []string `protobuf:"bytes,5,rep,name=allowed_request_headers,json=allowedRequestHeaders,proto3" json:"allowed_request_headers,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_971cd043e21bda4a, []int{1}
}
func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (dst *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(dst, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *core.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *HttpService) GetAllowedAuthorizationHeaders() []string {
	if m != nil {
		return m.AllowedAuthorizationHeaders
	}
	return nil
}

func (m *HttpService) GetAllowedRequestHeaders() []string {
	if m != nil {
		return m.AllowedRequestHeaders
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.http.ext_authz.v2alpha.ExtAuthz")
	proto.RegisterType((*HttpService)(nil), "envoy.config.filter.http.ext_authz.v2alpha.HttpService")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ext_authz/v2alpha/ext_authz.proto", fileDescriptor_971cd043e21bda4a)
}

var fileDescriptor_971cd043e21bda4a = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0xab, 0xd3, 0x40,
	0x10, 0xc7, 0x8d, 0xad, 0x9a, 0x6c, 0x7a, 0x28, 0x0b, 0x62, 0xa9, 0xa8, 0xa1, 0x78, 0x28, 0x22,
	0xbb, 0x10, 0x41, 0xd1, 0x5b, 0x2b, 0x62, 0x11, 0x04, 0x89, 0xf4, 0x22, 0x42, 0x58, 0x93, 0x49,
	0xb3, 0x10, 0xbb, 0xeb, 0x66, 0x13, 0x6b, 0xff, 0x61, 0x0f, 0xfe, 0x13, 0xb2, 0x3f, 0xfa, 0x12,
	0x78, 0xef, 0xf0, 0x8e, 0x99, 0xf9, 0x7e, 0x3e, 0x33, 0xb3, 0x41, 0xef, 0xe0, 0xd8, 0x8b, 0x3f,
	0xb4, 0x10, 0xc7, 0x8a, 0x1f, 0x68, 0xc5, 0x1b, 0x0d, 0x8a, 0xd6, 0x5a, 0x4b, 0x0a, 0x27, 0x9d,
	0xb3, 0x4e, 0xd7, 0x67, 0xda, 0xa7, 0xac, 0x91, 0x35, 0x1b, 0x2a, 0x44, 0x2a, 0xa1, 0x05, 0x7e,
	0x61, 0x59, 0xe2, 0x58, 0xe2, 0x58, 0x62, 0x58, 0x32, 0x24, 0x3d, 0xbb, 0x7c, 0xee, 0xe6, 0x30,
	0xc9, 0x69, 0x9f, 0xd2, 0x42, 0x28, 0xa0, 0x07, 0x25, 0x8b, 0xbc, 0x05, 0xd5, 0xf3, 0x02, 0x9c,
	0x71, 0x99, 0x5c, 0x4f, 0x19, 0x5f, 0xde, 0x29, 0xee, 0x12, 0xab, 0xbf, 0x01, 0x0a, 0x3f, 0x9c,
	0xf4, 0xc6, 0xc8, 0xf1, 0x7b, 0x34, 0x1b, 0x4b, 0x16, 0x41, 0x12, 0xac, 0xe3, 0xf4, 0x29, 0x71,
	0x7b, 0x31, 0xc9, 0x49, 0x9f, 0x12, 0x63, 0x21, 0x1f, 0x95, 0x2c, 0xbe, 0xba, 0xd4, 0xee, 0x4e,
	0x16, 0x1f, 0x86, 0x4f, 0xfc, 0x1d, 0xcd, 0xec, 0x8c, 0x8b, 0x64, 0x62, 0x25, 0x6f, 0xc8, 0xed,
	0x8f, 0x23, 0x3b, 0xad, 0xe5, 0xc8, 0x5e, 0x0f, 0x9f, 0xf8, 0x25, 0xc2, 0x15, 0xe3, 0x4d, 0xa7,
	0x20, 0xff, 0x29, 0x4a, 0xc8, 0x59, 0xd3, 0x88, 0xdf, 0x8b, 0xbb, 0x49, 0xb0, 0x0e, 0xb3, 0xb9,
	0xef, 0x7c, 0x16, 0x25, 0x6c, 0x4c, 0x7d, 0x8b, 0x50, 0xe8, 0xd7, 0x68, 0x57, 0xff, 0x02, 0x14,
	0x8f, 0xc4, 0xf8, 0x2d, 0x42, 0xa6, 0x07, 0xca, 0xbc, 0x86, 0x3f, 0x75, 0x79, 0xc3, 0xa9, 0x86,
	0xd9, 0x2b, 0x9e, 0x45, 0x2e, 0xbd, 0x57, 0x1c, 0x3f, 0x43, 0xb1, 0x64, 0xba, 0xce, 0xa5, 0x82,
	0x8a, 0x9f, 0xec, 0xf4, 0x28, 0x43, 0xa6, 0xf4, 0xc5, 0x56, 0xf0, 0x16, 0x3d, 0xb1, 0x8b, 0x41,
	0x69, 0x2f, 0x13, 0x8a, 0x9f, 0x99, 0xe6, 0xe2, 0x98, 0xd7, 0xc0, 0x4a, 0x50, 0xed, 0x62, 0x9a,
	0x4c, 0xd6, 0x51, 0xf6, 0xd8, 0x87, 0x36, 0xe3, 0xcc, 0xce, 0x45, 0xf0, 0x6b, 0xf4, 0xe8, 0xe2,
	0x50, 0xf0, 0xab, 0x83, 0x56, 0x5f, 0xd1, 0xf7, 0x2c, 0xfd, 0xd0, 0xb7, 0x33, 0xd7, 0xf5, 0xdc,
	0xa7, 0x69, 0x38, 0x99, 0x4f, 0xb7, 0xd1, 0xb7, 0x07, 0xfe, 0x35, 0x7f, 0xdc, 0xb7, 0x7f, 0xfa,
	0xd5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x5e, 0x82, 0x3e, 0x9b, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/range.proto

package envoy_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the int64 start and end of the range using half-open interval semantics [start,
// end).
type Int64Range struct {
	// start of the range (inclusive)
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Range) Reset()         { *m = Int64Range{} }
func (m *Int64Range) String() string { return proto.CompactTextString(m) }
func (*Int64Range) ProtoMessage()    {}
func (*Int64Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d5354588716e6e7, []int{0}
}
func (m *Int64Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Range.Unmarshal(m, b)
}
func (m *Int64Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Range.Marshal(b, m, deterministic)
}
func (dst *Int64Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Range.Merge(dst, src)
}
func (m *Int64Range) XXX_Size() int {
	return xxx_messageInfo_Int64Range.Size(m)
}
func (m *Int64Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Range.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Range proto.InternalMessageInfo

func (m *Int64Range) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Int64Range) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// Specifies the double start and end of the range using half-open interval semantics [start,
// end).
type DoubleRange struct {
	// start of the range (inclusive)
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d5354588716e6e7, []int{1}
}
func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRange.Unmarshal(m, b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
}
func (dst *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(dst, src)
}
func (m *DoubleRange) XXX_Size() int {
	return xxx_messageInfo_DoubleRange.Size(m)
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DoubleRange) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Int64Range)(nil), "envoy.type.Int64Range")
	proto.RegisterType((*DoubleRange)(nil), "envoy.type.DoubleRange")
}

func init() { proto.RegisterFile("envoy/type/range.proto", fileDescriptor_6d5354588716e6e7) }

var fileDescriptor_6d5354588716e6e7 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x4a, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x8b, 0xeb, 0x81, 0xc4, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1,
	0xc2, 0xfa, 0x20, 0x16, 0x44, 0x85, 0x92, 0x09, 0x17, 0x97, 0x67, 0x5e, 0x89, 0x99, 0x49, 0x10,
	0x48, 0x97, 0x90, 0x08, 0x17, 0x6b, 0x71, 0x49, 0x62, 0x51, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x73, 0x10, 0x84, 0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x9a, 0x97, 0x22, 0xc1, 0x04, 0x16, 0x03, 0x31,
	0x95, 0x4c, 0xb9, 0xb8, 0x5d, 0xf2, 0x4b, 0x93, 0x72, 0x52, 0xb1, 0x68, 0x63, 0xc4, 0xa2, 0x8d,
	0x11, 0xac, 0xcd, 0x49, 0x60, 0xc5, 0x23, 0x39, 0xc6, 0x28, 0x88, 0xa3, 0xe2, 0x41, 0x8e, 0x4a,
	0x62, 0x03, 0xbb, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x09, 0x37, 0xd9, 0x47, 0xc1, 0x00,
	0x00, 0x00,
}

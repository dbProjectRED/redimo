// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redimo.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type Table struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{0}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

type GETRequest struct {
	Table                *Table   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETRequest) Reset()         { *m = GETRequest{} }
func (m *GETRequest) String() string { return proto.CompactTextString(m) }
func (*GETRequest) ProtoMessage()    {}
func (*GETRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{1}
}

func (m *GETRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETRequest.Unmarshal(m, b)
}
func (m *GETRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETRequest.Marshal(b, m, deterministic)
}
func (m *GETRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETRequest.Merge(m, src)
}
func (m *GETRequest) XXX_Size() int {
	return xxx_messageInfo_GETRequest.Size(m)
}
func (m *GETRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETRequest proto.InternalMessageInfo

func (m *GETRequest) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *GETRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GETResponse struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Value                *any.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETResponse) Reset()         { *m = GETResponse{} }
func (m *GETResponse) String() string { return proto.CompactTextString(m) }
func (*GETResponse) ProtoMessage()    {}
func (*GETResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{2}
}

func (m *GETResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETResponse.Unmarshal(m, b)
}
func (m *GETResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETResponse.Marshal(b, m, deterministic)
}
func (m *GETResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETResponse.Merge(m, src)
}
func (m *GETResponse) XXX_Size() int {
	return xxx_messageInfo_GETResponse.Size(m)
}
func (m *GETResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETResponse proto.InternalMessageInfo

func (m *GETResponse) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *GETResponse) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type SETRequest struct {
	Table                *Table   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                *any.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SETRequest) Reset()         { *m = SETRequest{} }
func (m *SETRequest) String() string { return proto.CompactTextString(m) }
func (*SETRequest) ProtoMessage()    {}
func (*SETRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{3}
}

func (m *SETRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SETRequest.Unmarshal(m, b)
}
func (m *SETRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SETRequest.Marshal(b, m, deterministic)
}
func (m *SETRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SETRequest.Merge(m, src)
}
func (m *SETRequest) XXX_Size() int {
	return xxx_messageInfo_SETRequest.Size(m)
}
func (m *SETRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SETRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SETRequest proto.InternalMessageInfo

func (m *SETRequest) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *SETRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SETRequest) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type SETResponse struct {
	AlreadyPresent       bool     `protobuf:"varint,1,opt,name=already_present,json=alreadyPresent,proto3" json:"already_present,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SETResponse) Reset()         { *m = SETResponse{} }
func (m *SETResponse) String() string { return proto.CompactTextString(m) }
func (*SETResponse) ProtoMessage()    {}
func (*SETResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{4}
}

func (m *SETResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SETResponse.Unmarshal(m, b)
}
func (m *SETResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SETResponse.Marshal(b, m, deterministic)
}
func (m *SETResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SETResponse.Merge(m, src)
}
func (m *SETResponse) XXX_Size() int {
	return xxx_messageInfo_SETResponse.Size(m)
}
func (m *SETResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SETResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SETResponse proto.InternalMessageInfo

func (m *SETResponse) GetAlreadyPresent() bool {
	if m != nil {
		return m.AlreadyPresent
	}
	return false
}

type HGETRequest struct {
	Table                *Table   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Field                string   `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HGETRequest) Reset()         { *m = HGETRequest{} }
func (m *HGETRequest) String() string { return proto.CompactTextString(m) }
func (*HGETRequest) ProtoMessage()    {}
func (*HGETRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{5}
}

func (m *HGETRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HGETRequest.Unmarshal(m, b)
}
func (m *HGETRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HGETRequest.Marshal(b, m, deterministic)
}
func (m *HGETRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HGETRequest.Merge(m, src)
}
func (m *HGETRequest) XXX_Size() int {
	return xxx_messageInfo_HGETRequest.Size(m)
}
func (m *HGETRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HGETRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HGETRequest proto.InternalMessageInfo

func (m *HGETRequest) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *HGETRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HGETRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type HGETResponse struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Value                *any.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HGETResponse) Reset()         { *m = HGETResponse{} }
func (m *HGETResponse) String() string { return proto.CompactTextString(m) }
func (*HGETResponse) ProtoMessage()    {}
func (*HGETResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{6}
}

func (m *HGETResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HGETResponse.Unmarshal(m, b)
}
func (m *HGETResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HGETResponse.Marshal(b, m, deterministic)
}
func (m *HGETResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HGETResponse.Merge(m, src)
}
func (m *HGETResponse) XXX_Size() int {
	return xxx_messageInfo_HGETResponse.Size(m)
}
func (m *HGETResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HGETResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HGETResponse proto.InternalMessageInfo

func (m *HGETResponse) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *HGETResponse) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type HSETRequest struct {
	Table                *Table   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Field                string   `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	Value                *any.Any `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HSETRequest) Reset()         { *m = HSETRequest{} }
func (m *HSETRequest) String() string { return proto.CompactTextString(m) }
func (*HSETRequest) ProtoMessage()    {}
func (*HSETRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{7}
}

func (m *HSETRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSETRequest.Unmarshal(m, b)
}
func (m *HSETRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSETRequest.Marshal(b, m, deterministic)
}
func (m *HSETRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSETRequest.Merge(m, src)
}
func (m *HSETRequest) XXX_Size() int {
	return xxx_messageInfo_HSETRequest.Size(m)
}
func (m *HSETRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HSETRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HSETRequest proto.InternalMessageInfo

func (m *HSETRequest) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *HSETRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HSETRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *HSETRequest) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type HSETResponse struct {
	AlreadyPresent       bool     `protobuf:"varint,1,opt,name=already_present,json=alreadyPresent,proto3" json:"already_present,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HSETResponse) Reset()         { *m = HSETResponse{} }
func (m *HSETResponse) String() string { return proto.CompactTextString(m) }
func (*HSETResponse) ProtoMessage()    {}
func (*HSETResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1a0befbd905a, []int{8}
}

func (m *HSETResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSETResponse.Unmarshal(m, b)
}
func (m *HSETResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSETResponse.Marshal(b, m, deterministic)
}
func (m *HSETResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSETResponse.Merge(m, src)
}
func (m *HSETResponse) XXX_Size() int {
	return xxx_messageInfo_HSETResponse.Size(m)
}
func (m *HSETResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HSETResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HSETResponse proto.InternalMessageInfo

func (m *HSETResponse) GetAlreadyPresent() bool {
	if m != nil {
		return m.AlreadyPresent
	}
	return false
}

func init() {
	proto.RegisterType((*Table)(nil), "dbProjectRED.redimo.v1.Table")
	proto.RegisterType((*GETRequest)(nil), "dbProjectRED.redimo.v1.GETRequest")
	proto.RegisterType((*GETResponse)(nil), "dbProjectRED.redimo.v1.GETResponse")
	proto.RegisterType((*SETRequest)(nil), "dbProjectRED.redimo.v1.SETRequest")
	proto.RegisterType((*SETResponse)(nil), "dbProjectRED.redimo.v1.SETResponse")
	proto.RegisterType((*HGETRequest)(nil), "dbProjectRED.redimo.v1.HGETRequest")
	proto.RegisterType((*HGETResponse)(nil), "dbProjectRED.redimo.v1.HGETResponse")
	proto.RegisterType((*HSETRequest)(nil), "dbProjectRED.redimo.v1.HSETRequest")
	proto.RegisterType((*HSETResponse)(nil), "dbProjectRED.redimo.v1.HSETResponse")
}

func init() {
	proto.RegisterFile("redimo.proto", fileDescriptor_7afd1a0befbd905a)
}

var fileDescriptor_7afd1a0befbd905a = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xc9, 0x66, 0xa3, 0xee, 0x49, 0xfd, 0xc3, 0xb0, 0xc8, 0x1a, 0x51, 0xd6, 0xac, 0xe0,
	0xe2, 0xc5, 0x84, 0xed, 0x82, 0x5e, 0x2b, 0x96, 0xe6, 0xaa, 0xd6, 0x99, 0x5e, 0x79, 0x23, 0x49,
	0x73, 0x5a, 0x53, 0xd3, 0x4c, 0x9c, 0xfc, 0x91, 0x5c, 0xf9, 0x16, 0xbe, 0x99, 0xef, 0x23, 0x9d,
	0x49, 0x69, 0x28, 0x36, 0x2d, 0x52, 0xbc, 0xcb, 0x9c, 0xf3, 0x7d, 0x1f, 0xbf, 0x73, 0x4e, 0xa0,
	0x27, 0x31, 0x8a, 0x97, 0x82, 0x66, 0x52, 0x14, 0x82, 0x3c, 0x8e, 0xc2, 0xb1, 0x14, 0x0b, 0x9c,
	0x16, 0x6c, 0xf0, 0x81, 0x36, 0xad, 0xea, 0xc6, 0x79, 0x32, 0x17, 0x62, 0x9e, 0xa0, 0xa7, 0x54,
	0x61, 0x39, 0xf3, 0x82, 0xb4, 0xd6, 0x16, 0xe7, 0xe9, 0x76, 0x0b, 0x97, 0x59, 0xb1, 0x6e, 0x3e,
	0xdf, 0x6e, 0xfe, 0x90, 0x41, 0x96, 0xa1, 0xcc, 0x75, 0xdf, 0xbd, 0x0b, 0xd6, 0x24, 0x08, 0x13,
	0x74, 0x39, 0xc0, 0x70, 0x30, 0x61, 0xf8, 0xbd, 0xc4, 0xbc, 0x20, 0xb7, 0x60, 0x15, 0xab, 0xf2,
	0x85, 0x71, 0x69, 0x5c, 0xdb, 0xfd, 0x67, 0xf4, 0xef, 0x58, 0x54, 0x79, 0x99, 0xd6, 0x92, 0x47,
	0x60, 0x7e, 0xc3, 0xfa, 0xe2, 0xe4, 0xd2, 0xb8, 0x3e, 0x63, 0xab, 0x4f, 0xf7, 0x23, 0xd8, 0x2a,
	0x34, 0xcf, 0x44, 0x9a, 0x23, 0x39, 0x07, 0x6b, 0x26, 0xca, 0x34, 0x52, 0xa9, 0xf7, 0x98, 0x7e,
	0x90, 0xd7, 0x60, 0x55, 0x41, 0x52, 0xa2, 0x32, 0xda, 0xfd, 0x73, 0xaa, 0x91, 0xe9, 0x1a, 0x99,
	0xbe, 0x4b, 0x6b, 0xa6, 0x25, 0xee, 0x4f, 0x00, 0x7e, 0x6c, 0xca, 0x0d, 0x80, 0xb9, 0x1f, 0xe0,
	0x0d, 0xd8, 0xbc, 0x35, 0xd1, 0x2b, 0x78, 0x18, 0x24, 0x12, 0x83, 0xa8, 0xfe, 0x92, 0x49, 0xcc,
	0x31, 0x2d, 0x9a, 0xd9, 0x1e, 0x34, 0xe5, 0xb1, 0xae, 0xba, 0x0b, 0xb0, 0xfd, 0xa3, 0xef, 0x57,
	0x2d, 0x34, 0xc6, 0x24, 0x52, 0xe4, 0x67, 0x4c, 0x3f, 0xdc, 0x31, 0xf4, 0xfc, 0xe3, 0xae, 0xfd,
	0x97, 0x01, 0xb6, 0xcf, 0xff, 0x0f, 0xfe, 0x06, 0xec, 0x74, 0x3f, 0xd8, 0x5b, 0xe8, 0xf9, 0xff,
	0x72, 0x8f, 0xfe, 0xef, 0x13, 0xb8, 0xcf, 0x14, 0x27, 0x47, 0x59, 0xc5, 0x53, 0x24, 0x23, 0x30,
	0x87, 0x83, 0x09, 0x71, 0x77, 0xcd, 0xb2, 0xb9, 0x9e, 0x73, 0xd5, 0xa9, 0x69, 0x50, 0x46, 0x60,
	0xf2, 0xae, 0x3c, 0x7e, 0x40, 0x5e, 0x7b, 0xb4, 0x4f, 0x70, 0xba, 0xba, 0x2a, 0xd9, 0x29, 0x6e,
	0xfd, 0x5f, 0xce, 0xcb, 0x6e, 0x51, 0x2b, 0x92, 0x77, 0x46, 0xf2, 0x43, 0x22, 0x5b, 0x94, 0xef,
	0xaf, 0x3e, 0xbf, 0x98, 0xc7, 0xc5, 0xd7, 0x32, 0xa4, 0x53, 0xb1, 0xf4, 0xda, 0x0e, 0x4f, 0x3b,
	0xbc, 0xea, 0x26, 0xbc, 0xa3, 0x4e, 0x79, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x25, 0x32, 0xbd,
	0x45, 0xfb, 0x04, 0x00, 0x00,
}
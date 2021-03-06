// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package go_micro_lgm_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "go.micro.lgm.srv.auth.Error")
	proto.RegisterType((*Request)(nil), "go.micro.lgm.srv.auth.Request")
	proto.RegisterType((*Response)(nil), "go.micro.lgm.srv.auth.Response")
}

func init() {
	proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5)
}

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0x09, 0x34, 0x6d, 0x30, 0x03, 0x92, 0x45, 0x51, 0x54, 0x21, 0xa8, 0x32, 0x75, 0x3a,
	0xa4, 0xf4, 0x09, 0x90, 0x60, 0x60, 0x80, 0xe1, 0x5a, 0xc4, 0x1c, 0x2e, 0x56, 0x89, 0x9a, 0xe4,
	0xca, 0xf9, 0xd2, 0x67, 0xe3, 0xf1, 0xd0, 0x39, 0x81, 0xa9, 0xdd, 0x58, 0x4e, 0xfe, 0xf4, 0xdb,
	0xbf, 0xef, 0x97, 0x61, 0xba, 0x73, 0xd6, 0xdb, 0xfb, 0xa2, 0xf3, 0x9f, 0xf2, 0x28, 0x61, 0x9c,
	0x6e, 0xac, 0x6a, 0x2a, 0xe3, 0xac, 0xaa, 0x37, 0x8d, 0x62, 0xb7, 0x57, 0x41, 0xcc, 0x96, 0x10,
	0x3f, 0x39, 0x67, 0x1d, 0x22, 0x8c, 0x8c, 0x2d, 0x29, 0x8d, 0xe6, 0xd1, 0x22, 0xd6, 0x52, 0xe3,
	0x35, 0x8c, 0x4b, 0xf2, 0x45, 0x55, 0xa7, 0xa7, 0xf3, 0x68, 0x71, 0xae, 0x07, 0xca, 0x56, 0x30,
	0xd1, 0xf4, 0xd5, 0x11, 0xfb, 0xd0, 0xd2, 0x31, 0xb9, 0xe7, 0x52, 0x06, 0x47, 0x7a, 0x20, 0x9c,
	0x41, 0x12, 0xaa, 0xd7, 0xa2, 0xa1, 0x61, 0xf8, 0x8f, 0xf1, 0x0a, 0x62, 0x6f, 0xb7, 0xd4, 0xa6,
	0x67, 0x22, 0xf4, 0x90, 0xb5, 0x90, 0x68, 0xe2, 0x9d, 0x6d, 0x99, 0x30, 0x85, 0x09, 0x77, 0xc6,
	0x10, 0xb3, 0xd8, 0x26, 0xfa, 0x17, 0x31, 0x87, 0x98, 0xc2, 0x7f, 0xc5, 0xf4, 0x22, 0xbf, 0x51,
	0x07, 0x63, 0x29, 0xc9, 0xa4, 0xfb, 0xd6, 0xc3, 0xfb, 0xf2, 0xef, 0x08, 0x26, 0x2b, 0x72, 0xfb,
	0xca, 0x10, 0xae, 0xe1, 0xf2, 0xa5, 0xd8, 0xd2, 0x83, 0xec, 0x58, 0x07, 0x19, 0x6f, 0x8f, 0x38,
	0x0f, 0xc1, 0x67, 0x77, 0x47, 0xf5, 0x3e, 0x43, 0x76, 0x82, 0xef, 0x80, 0x8f, 0x54, 0xbf, 0x31,
	0xb9, 0xff, 0x35, 0xfe, 0x18, 0xcb, 0x49, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0x10,
	0x3d, 0x8d, 0xeb, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/proto-files/domain/repository.proto

package domain

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

type Repository struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	IsPrivate            bool     `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_c48554a03e0cda6d, []int{0}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repository) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Repository) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func init() {
	proto.RegisterType((*Repository)(nil), "domain.Repository")
}

func init() {
	proto.RegisterFile("internal/proto-files/domain/repository.proto", fileDescriptor_c48554a03e0cda6d)
}

var fileDescriptor_c48554a03e0cda6d = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0x49, 0x5b, 0x8a, 0xbd, 0xc1, 0x21, 0x83, 0x64, 0x50, 0x2c, 0x4e, 0x1d, 0xac, 0x19,
	0xfc, 0x07, 0x6e, 0x6e, 0x92, 0xd1, 0x2d, 0x9a, 0x54, 0x0e, 0xda, 0xa4, 0x5c, 0xa2, 0xd0, 0x7f,
	0x2f, 0xa4, 0x52, 0xb7, 0x7b, 0xf7, 0xbe, 0xe1, 0xc1, 0x11, 0x5d, 0xb4, 0xe4, 0x74, 0x2f, 0x47,
	0xf2, 0xd1, 0xb7, 0x1d, 0xf6, 0x36, 0x48, 0xe3, 0x07, 0x8d, 0x4e, 0x92, 0x1d, 0x7d, 0xc0, 0xe8,
	0x69, 0x3a, 0x25, 0xcb, 0xcb, 0x59, 0x1c, 0x3a, 0x00, 0xb5, 0x38, 0xbe, 0x86, 0x0c, 0x8d, 0x60,
	0x35, 0x6b, 0x72, 0x95, 0xa1, 0xe1, 0x1c, 0x0a, 0xa7, 0x07, 0x2b, 0xb2, 0x9a, 0x35, 0x95, 0x4a,
	0x37, 0xdf, 0x40, 0xf9, 0x0e, 0x96, 0xae, 0x46, 0xe4, 0x69, 0xf7, 0x23, 0xbe, 0x85, 0x0a, 0xc3,
	0x8d, 0xf0, 0xa3, 0xa3, 0x15, 0x45, 0xcd, 0x9a, 0x95, 0xfa, 0x3f, 0x2e, 0xfb, 0xfb, 0x2e, 0xd2,
	0xd4, 0xbe, 0x68, 0x7c, 0xca, 0x25, 0x34, 0xd1, 0x1c, 0xf2, 0x28, 0x53, 0xd7, 0xf9, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xbe, 0x84, 0x3d, 0xb1, 0xc7, 0x00, 0x00, 0x00,
}

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
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0x49, 0x5b, 0x8a, 0xcd, 0xe0, 0x90, 0x41, 0x32, 0x38, 0x14, 0x11, 0xe9, 0x60, 0xcd,
	0xe0, 0x3f, 0x70, 0x73, 0x93, 0x8c, 0x6e, 0xd1, 0x5c, 0xcb, 0x41, 0x9b, 0x84, 0x24, 0x8a, 0xfe,
	0x7b, 0x31, 0x96, 0xba, 0xdd, 0xbd, 0xf7, 0x0d, 0x8f, 0xee, 0xd1, 0x44, 0xf0, 0x46, 0x0d, 0xc2,
	0x79, 0x1b, 0x6d, 0xdb, 0xe1, 0x00, 0x41, 0x68, 0x3b, 0x2a, 0x34, 0xc2, 0x83, 0xb3, 0x01, 0xa3,
	0xf5, 0xef, 0x43, 0xb2, 0xac, 0xfc, 0x89, 0x4d, 0x47, 0xa9, 0x9c, 0x1d, 0x5b, 0xd2, 0x0c, 0x35,
	0x27, 0x35, 0x69, 0x72, 0x99, 0xa1, 0x66, 0x8c, 0x16, 0x46, 0x8d, 0xc0, 0xb3, 0x9a, 0x34, 0x95,
	0x4c, 0x37, 0x5b, 0xd1, 0xf2, 0x11, 0xc0, 0x9f, 0x35, 0xcf, 0xd3, 0x6e, 0xfa, 0xd8, 0x9a, 0x56,
	0x18, 0x2e, 0x1e, 0x9f, 0x2a, 0x02, 0x2f, 0x6a, 0xd2, 0x2c, 0xe4, 0x1f, 0x9c, 0x76, 0xd7, 0x6d,
	0x6f, 0xdb, 0xde, 0xbb, 0x7b, 0x0b, 0x2f, 0x35, 0xba, 0x01, 0xc4, 0xdc, 0xfb, 0xa5, 0x53, 0xe8,
	0xad, 0x4c, 0x79, 0xc7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xc4, 0xba, 0xa5, 0xce, 0x00,
	0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

/*
Package groupcache_test is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	User
*/
package groupcache_test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age     int64  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	IsSuper bool   `protobuf:"varint,4,opt,name=is_super,json=isSuper" json:"is_super,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetIsSuper() bool {
	if m != nil {
		return m.IsSuper
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "groupcachepb.User")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x2f, 0xca, 0x2f, 0x2d,
	0x48, 0x4e, 0x4c, 0xce, 0x48, 0x2d, 0x48, 0x52, 0x0a, 0xe7, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0x02, 0x5c,
	0xcc, 0x89, 0xe9, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x20, 0xa6, 0x90, 0x24, 0x17,
	0x47, 0x66, 0x71, 0x7c, 0x71, 0x69, 0x41, 0x6a, 0x91, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x47, 0x10,
	0x7b, 0x66, 0x71, 0x30, 0x88, 0xeb, 0x24, 0x18, 0xc5, 0x8f, 0xb0, 0x28, 0xbe, 0x24, 0xb5, 0xb8,
	0x24, 0x89, 0x0d, 0xec, 0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x2e, 0x5f, 0x1a,
	0x91, 0x00, 0x00, 0x00,
}
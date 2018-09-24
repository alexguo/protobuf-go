// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/b.proto

package import_public

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	sub "github.com/golang/protobuf/v2/cmd/protoc-gen-go/testdata/import_public/sub"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Local struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
	E                    sub.E    `protobuf:"varint,2,opt,name=e,proto3,enum=goproto.protoc.import_public.sub.E" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Local) Reset()         { *m = Local{} }
func (m *Local) String() string { return proto.CompactTextString(m) }
func (*Local) ProtoMessage()    {}
func (*Local) Descriptor() ([]byte, []int) {
	return fileDescriptor_84995586b3d09710, []int{0}
}

func (m *Local) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Local.Unmarshal(m, b)
}
func (m *Local) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Local.Marshal(b, m, deterministic)
}
func (m *Local) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Local.Merge(m, src)
}
func (m *Local) XXX_Size() int {
	return xxx_messageInfo_Local.Size(m)
}
func (m *Local) XXX_DiscardUnknown() {
	xxx_messageInfo_Local.DiscardUnknown(m)
}

var xxx_messageInfo_Local proto.InternalMessageInfo

func (m *Local) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Local) GetE() sub.E {
	if m != nil {
		return m.E
	}
	return sub.E_ZERO
}

func init() {
	proto.RegisterType((*Local)(nil), "goproto.protoc.import_public.Local")
}

func init() { proto.RegisterFile("import_public/b.proto", fileDescriptor_84995586b3d09710) }

var fileDescriptor_84995586b3d09710 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8c, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0x89, 0xa0, 0x43, 0x05, 0x87, 0x82, 0x50, 0xc5, 0xa1, 0xe8, 0xd2, 0xa5, 0x39, 0xac,
	0xff, 0x40, 0x50, 0x1c, 0x74, 0xe9, 0xe8, 0x22, 0xb9, 0x34, 0xc6, 0x42, 0xd3, 0x2b, 0xed, 0xc5,
	0xdf, 0x2f, 0xb6, 0x53, 0x17, 0x71, 0xfb, 0x3e, 0x78, 0xef, 0x05, 0xcb, 0xd2, 0x35, 0xd4, 0xf2,
	0xa3, 0xf1, 0x58, 0x95, 0x1a, 0x50, 0x36, 0x2d, 0x31, 0x85, 0x1b, 0x4b, 0xfd, 0x18, 0xae, 0x96,
	0x23, 0x6a, 0xbd, 0x1a, 0x4b, 0x9d, 0x47, 0x50, 0x03, 0xb9, 0x75, 0xc1, 0xf4, 0x4a, 0x5a, 0x55,
	0xe1, 0x3e, 0x10, 0x2e, 0x12, 0xb1, 0x48, 0xe6, 0xd9, 0x4e, 0xfe, 0xaa, 0xc9, 0xce, 0xa3, 0xbc,
	0xe5, 0xc2, 0x7d, 0x15, 0x13, 0x4d, 0x62, 0x91, 0x2c, 0xfe, 0x51, 0x4e, 0xb9, 0x30, 0xc7, 0xcb,
	0xfd, 0x6c, 0x4b, 0x7e, 0x79, 0x94, 0x9a, 0x1c, 0x58, 0xaa, 0x54, 0x6d, 0xa1, 0x57, 0xd0, 0x3f,
	0xe1, 0x9d, 0x81, 0x76, 0xc5, 0xf0, 0x75, 0x6a, 0x4d, 0x9d, 0x5a, 0x02, 0x36, 0x1d, 0x17, 0x8a,
	0x15, 0x8c, 0x92, 0x38, 0xeb, 0xa9, 0xc3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x20, 0x34, 0xa1, 0xfe,
	0x11, 0x01, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: persistance.proto

package transientstore

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type PendingDeleteStorageList struct {
	List                 []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PendingDeleteStorageList) Reset()         { *m = PendingDeleteStorageList{} }
func (m *PendingDeleteStorageList) String() string { return proto.CompactTextString(m) }
func (*PendingDeleteStorageList) ProtoMessage()    {}
func (*PendingDeleteStorageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a374ca4de69122a4, []int{0}
}

func (m *PendingDeleteStorageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingDeleteStorageList.Unmarshal(m, b)
}
func (m *PendingDeleteStorageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingDeleteStorageList.Marshal(b, m, deterministic)
}
func (m *PendingDeleteStorageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingDeleteStorageList.Merge(m, src)
}
func (m *PendingDeleteStorageList) XXX_Size() int {
	return xxx_messageInfo_PendingDeleteStorageList.Size(m)
}
func (m *PendingDeleteStorageList) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingDeleteStorageList.DiscardUnknown(m)
}

var xxx_messageInfo_PendingDeleteStorageList proto.InternalMessageInfo

func (m *PendingDeleteStorageList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*PendingDeleteStorageList)(nil), "transientstore.PendingDeleteStorageList")
}

func init() { proto.RegisterFile("persistance.proto", fileDescriptor_a374ca4de69122a4) }

var fileDescriptor_a374ca4de69122a4 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcd, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x06, 0x60, 0x44, 0x11, 0xec, 0x20, 0xd8, 0xa9, 0xa3, 0x38, 0x39, 0x35, 0x48, 0xdf, 0x40,
	0x1c, 0x1d, 0x44, 0x37, 0xb7, 0x24, 0xfd, 0x4d, 0x0f, 0x62, 0x2e, 0xdc, 0x9d, 0x83, 0x6f, 0x2f,
	0x74, 0x73, 0xfb, 0xb6, 0xaf, 0xd9, 0x55, 0x88, 0x92, 0x9a, 0x2f, 0x11, 0x7d, 0x15, 0x36, 0x6e,
	0xb7, 0x26, 0xbe, 0x28, 0xa1, 0x98, 0x1a, 0x0b, 0x0e, 0x7d, 0xd3, 0xdd, 0x50, 0x46, 0x2a, 0xe9,
	0x82, 0x0c, 0xc3, 0xc3, 0x58, 0x7c, 0xc2, 0x95, 0xd4, 0xda, 0xb6, 0x59, 0x65, 0x52, 0xeb, 0x16,
	0xfb, 0xe5, 0x71, 0x73, 0x9f, 0x7d, 0x1e, 0x9e, 0xa7, 0x44, 0x36, 0x7d, 0x42, 0x1f, 0xf9, 0xed,
	0xa6, 0x6f, 0x85, 0x64, 0x8c, 0x09, 0xe2, 0x5e, 0x3e, 0x08, 0x45, 0x17, 0x59, 0xe0, 0xfe, 0x93,
	0xb0, 0x9e, 0xef, 0xe1, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xb1, 0xe2, 0x32, 0x90, 0x00, 0x00,
	0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synced_notification_app_info_specifics.proto

package sync_pb

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

// This message is kept around for backwards compatibility sake.
type SyncedNotificationAppInfoSpecifics struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncedNotificationAppInfoSpecifics) Reset()         { *m = SyncedNotificationAppInfoSpecifics{} }
func (m *SyncedNotificationAppInfoSpecifics) String() string { return proto.CompactTextString(m) }
func (*SyncedNotificationAppInfoSpecifics) ProtoMessage()    {}
func (*SyncedNotificationAppInfoSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a13d11d5d40f4d8, []int{0}
}

func (m *SyncedNotificationAppInfoSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Unmarshal(m, b)
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Marshal(b, m, deterministic)
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Merge(m, src)
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_Size() int {
	return xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Size(m)
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncedNotificationAppInfoSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_SyncedNotificationAppInfoSpecifics proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SyncedNotificationAppInfoSpecifics)(nil), "sync_pb.SyncedNotificationAppInfoSpecifics")
}

func init() {
	proto.RegisterFile("synced_notification_app_info_specifics.proto", fileDescriptor_7a13d11d5d40f4d8)
}

var fileDescriptor_7a13d11d5d40f4d8 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0xae, 0xcc, 0x4b,
	0x4e, 0x4d, 0x89, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b,
	0x4f, 0x2c, 0x28, 0x88, 0xcf, 0xcc, 0x4b, 0xcb, 0x8f, 0x2f, 0x2e, 0x48, 0x4d, 0x06, 0x09, 0x17,
	0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xb1, 0x83, 0x54, 0xc7, 0x17, 0x24, 0x29, 0xa9, 0x70,
	0x29, 0x05, 0x83, 0x35, 0xfa, 0x21, 0xe9, 0x73, 0x2c, 0x28, 0xf0, 0xcc, 0x4b, 0xcb, 0x0f, 0x86,
	0x69, 0x72, 0xd2, 0xe6, 0x52, 0xcd, 0x2f, 0x4a, 0xd7, 0x4b, 0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c,
	0xcd, 0xd5, 0x4b, 0xce, 0xcf, 0x2d, 0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x03, 0x19, 0x04,
	0x31, 0x34, 0x39, 0x3f, 0xc7, 0x83, 0x39, 0x80, 0x11, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xe3,
	0xc2, 0xfd, 0x8a, 0x00, 0x00, 0x00,
}
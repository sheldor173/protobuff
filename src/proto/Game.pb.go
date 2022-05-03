// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Game.proto

package proto

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

type Game struct {
	Home                 string   `protobuf:"bytes,1,opt,name=home,proto3" json:"home,omitempty"`
	Away                 string   `protobuf:"bytes,2,opt,name=away,proto3" json:"away,omitempty"`
	Venue                string   `protobuf:"bytes,3,opt,name=venue,proto3" json:"venue,omitempty"`
	Date                 string   `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8decbf488d309b, []int{0}
}

func (m *Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game.Unmarshal(m, b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game.Marshal(b, m, deterministic)
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return xxx_messageInfo_Game.Size(m)
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetHome() string {
	if m != nil {
		return m.Home
	}
	return ""
}

func (m *Game) GetAway() string {
	if m != nil {
		return m.Away
	}
	return ""
}

func (m *Game) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *Game) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*Game)(nil), "proto.Game")
}

func init() { proto.RegisterFile("proto/Game.proto", fileDescriptor_5f8decbf488d309b) }

var fileDescriptor_5f8decbf488d309b = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x77, 0x4f, 0xcc, 0x4d, 0xd5, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x52, 0x04, 0x17,
	0x0b, 0x48, 0x50, 0x48, 0x88, 0x8b, 0x25, 0x23, 0x3f, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xcc, 0x06, 0x89, 0x25, 0x96, 0x27, 0x56, 0x4a, 0x30, 0x41, 0xc4, 0x40, 0x6c, 0x21,
	0x11, 0x2e, 0xd6, 0xb2, 0xd4, 0xbc, 0xd2, 0x54, 0x09, 0x66, 0xb0, 0x20, 0x84, 0x03, 0x52, 0x99,
	0x92, 0x58, 0x92, 0x2a, 0xc1, 0x02, 0x51, 0x09, 0x62, 0x27, 0xb1, 0x81, 0x2d, 0x30, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0x65, 0x36, 0xee, 0x7b, 0x00, 0x00, 0x00,
}

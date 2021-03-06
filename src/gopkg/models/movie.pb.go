// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/movie.proto

package micro_models

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

type Movie struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ContactId            string   `protobuf:"bytes,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	Tags                 string   `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Status               int64    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Source               int64    `protobuf:"varint,5,opt,name=source,proto3" json:"source,omitempty"`
	TicketCreatedAt      int64    `protobuf:"varint,6,opt,name=ticket_created_at,json=ticketCreatedAt,proto3" json:"ticket_created_at,omitempty"`
	TicketUpdatedAt      int64    `protobuf:"varint,7,opt,name=ticket_updated_at,json=ticketUpdatedAt,proto3" json:"ticket_updated_at,omitempty"`
	CreatedAt            int64    `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	AppCode              string   `protobuf:"bytes,10,opt,name=app_code,json=appCode,proto3" json:"app_code,omitempty"`
	Platform             string   `protobuf:"bytes,11,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_e70c493f74819fdb, []int{0}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Movie) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func (m *Movie) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Movie) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Movie) GetSource() int64 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *Movie) GetTicketCreatedAt() int64 {
	if m != nil {
		return m.TicketCreatedAt
	}
	return 0
}

func (m *Movie) GetTicketUpdatedAt() int64 {
	if m != nil {
		return m.TicketUpdatedAt
	}
	return 0
}

func (m *Movie) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Movie) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Movie) GetAppCode() string {
	if m != nil {
		return m.AppCode
	}
	return ""
}

func (m *Movie) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func init() {
	proto.RegisterType((*Movie)(nil), "micro.models.Movie")
}

func init() { proto.RegisterFile("models/movie.proto", fileDescriptor_e70c493f74819fdb) }

var fileDescriptor_e70c493f74819fdb = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xd5, 0xb4, 0x4d, 0x93, 0xfb, 0xff, 0x05, 0xc2, 0x03, 0x32, 0x48, 0x95, 0x2a, 0xa6,
	0xaa, 0x42, 0xcd, 0xc0, 0xc8, 0x54, 0x3a, 0x31, 0xb0, 0x54, 0x62, 0x61, 0x89, 0x8c, 0x6d, 0x22,
	0xab, 0x0d, 0x67, 0xd9, 0x97, 0x7e, 0x45, 0xbe, 0x16, 0xe2, 0x1c, 0x42, 0x36, 0xbf, 0xf7, 0x7b,
	0x7e, 0xb2, 0xef, 0x40, 0xb4, 0x68, 0xec, 0x29, 0x56, 0x2d, 0x9e, 0x9d, 0xdd, 0xfa, 0x80, 0x84,
	0xe2, 0x7f, 0xeb, 0x74, 0xc0, 0x6d, 0x22, 0x77, 0x5f, 0x19, 0xcc, 0x5f, 0x7e, 0xa8, 0xb8, 0x80,
	0xcc, 0x19, 0x39, 0x59, 0x4d, 0xd6, 0xd3, 0x43, 0xe6, 0x8c, 0x58, 0x02, 0x68, 0xfc, 0x24, 0xa5,
	0xa9, 0x76, 0x46, 0x66, 0xab, 0xc9, 0xba, 0x3c, 0x94, 0xbd, 0xf3, 0x6c, 0x84, 0x80, 0x19, 0xa9,
	0x26, 0xca, 0x29, 0x03, 0x3e, 0x8b, 0x6b, 0xc8, 0x23, 0x29, 0xea, 0xa2, 0x9c, 0x71, 0x4d, 0xaf,
	0xd8, 0xc7, 0x2e, 0x68, 0x2b, 0xe7, 0xbd, 0xcf, 0x4a, 0x6c, 0xe0, 0x8a, 0x9c, 0x3e, 0x5a, 0xaa,
	0x75, 0xb0, 0x8a, 0xac, 0xa9, 0x15, 0xc9, 0x9c, 0x23, 0x97, 0x09, 0xec, 0x93, 0xbf, 0xa3, 0x51,
	0xb6, 0xf3, 0xe6, 0x37, 0xbb, 0x18, 0x67, 0x5f, 0x93, 0xbf, 0x23, 0x7e, 0xfa, 0x5f, 0x61, 0xc1,
	0xa1, 0x52, 0x0f, 0x55, 0x4b, 0x80, 0x51, 0x47, 0x99, 0x70, 0x37, 0xdc, 0xbe, 0x81, 0x42, 0x79,
	0x5f, 0x6b, 0x34, 0x56, 0x02, 0xff, 0x6e, 0xa1, 0xbc, 0xdf, 0xa3, 0xb1, 0xe2, 0x16, 0x0a, 0x7f,
	0x52, 0xf4, 0x81, 0xa1, 0x95, 0xff, 0x18, 0x0d, 0xfa, 0xe9, 0xfe, 0x6d, 0xc3, 0x93, 0x8d, 0x36,
	0x9c, 0x9d, 0xb6, 0x55, 0x0c, 0xba, 0x6a, 0xd0, 0x1f, 0x9b, 0x2a, 0x8d, 0xfa, 0x91, 0x69, 0x9d,
	0xc4, 0x7b, 0xce, 0xcb, 0x78, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x70, 0xc4, 0x03, 0xa2,
	0x01, 0x00, 0x00,
}

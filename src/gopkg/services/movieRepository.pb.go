// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/movieRepository.proto

package micro_services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "microservice/src/gopkg/models"
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

func init() { proto.RegisterFile("services/movieRepository.proto", fileDescriptor_925834d431fad026) }

var fileDescriptor_925834d431fad026 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0xcf, 0xcd, 0x2f, 0xcb, 0x4c, 0x0d, 0x4a, 0x2d, 0xc8, 0x2f, 0xce,
	0x2c, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xcd, 0x4c, 0x2e,
	0xca, 0xd7, 0x83, 0xa9, 0x92, 0x12, 0xca, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0xaa, 0x86, 0xa8, 0x31,
	0x72, 0xe7, 0xe2, 0xf7, 0x45, 0xd5, 0x2c, 0x64, 0xc2, 0xc5, 0xe6, 0x5c, 0x94, 0x9a, 0x58, 0x92,
	0x2a, 0x24, 0xac, 0x07, 0x31, 0x01, 0xa2, 0x4f, 0x0f, 0xac, 0x50, 0x0a, 0x9b, 0xa0, 0x12, 0x43,
	0x12, 0x1b, 0xd8, 0x3c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xc3, 0x35, 0x1a, 0x95,
	0x00, 0x00, 0x00,
}

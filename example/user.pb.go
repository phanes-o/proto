// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/user.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/phanes-o/proto/dto"
	_ "github.com/phanes-o/proto/primitive"
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

func init() { proto.RegisterFile("proto/example/user.proto", fileDescriptor_20657522f6c0c0ee) }

var fileDescriptor_20657522f6c0c0ee = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03,
	0x0b, 0x09, 0xb1, 0x43, 0xc5, 0xa4, 0xe4, 0x21, 0x4a, 0x0a, 0x8a, 0x32, 0x73, 0x33, 0x4b, 0x32,
	0xcb, 0x52, 0x11, 0x2c, 0x88, 0x4a, 0x29, 0x11, 0x88, 0x82, 0x94, 0x92, 0x7c, 0x24, 0xfd, 0x46,
	0x19, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x46, 0x5c, 0x6c, 0xce, 0x45, 0xa9, 0x89, 0x25,
	0xa9, 0x42, 0x62, 0x7a, 0x29, 0x25, 0xf9, 0x7a, 0x10, 0x0e, 0x48, 0x2a, 0x28, 0xb5, 0xb0, 0x34,
	0xb5, 0xb8, 0x44, 0x4a, 0x40, 0x0f, 0x61, 0xa2, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x90, 0x0e, 0x17,
	0x9b, 0x4b, 0x6a, 0x4e, 0x6a, 0x49, 0xaa, 0x10, 0xb2, 0x9c, 0x67, 0x5e, 0x89, 0x99, 0x09, 0xa6,
	0x6a, 0x27, 0xe5, 0x28, 0xc5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd,
	0x82, 0x8c, 0xc4, 0xbc, 0xd4, 0x62, 0x5d, 0x90, 0x83, 0x91, 0x7c, 0x96, 0xc4, 0x06, 0xe6, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x4c, 0x1c, 0x53, 0xf1, 0x00, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/extension_setting_device.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possbile device types for an extension setting.
type ExtensionSettingDeviceEnum_ExtensionSettingDevice int32

const (
	// Not specified.
	ExtensionSettingDeviceEnum_UNSPECIFIED ExtensionSettingDeviceEnum_ExtensionSettingDevice = 0
	// The value is unknown in this version.
	ExtensionSettingDeviceEnum_UNKNOWN ExtensionSettingDeviceEnum_ExtensionSettingDevice = 1
	// Mobile. The extensions in the extension setting will only serve on
	// mobile devices.
	ExtensionSettingDeviceEnum_MOBILE ExtensionSettingDeviceEnum_ExtensionSettingDevice = 2
	// Desktop. The extensions in the extension setting will only serve on
	// desktop devices.
	ExtensionSettingDeviceEnum_DESKTOP ExtensionSettingDeviceEnum_ExtensionSettingDevice = 3
)

var ExtensionSettingDeviceEnum_ExtensionSettingDevice_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MOBILE",
	3: "DESKTOP",
}

var ExtensionSettingDeviceEnum_ExtensionSettingDevice_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MOBILE":      2,
	"DESKTOP":     3,
}

func (x ExtensionSettingDeviceEnum_ExtensionSettingDevice) String() string {
	return proto.EnumName(ExtensionSettingDeviceEnum_ExtensionSettingDevice_name, int32(x))
}

func (ExtensionSettingDeviceEnum_ExtensionSettingDevice) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be48020c63d81f10, []int{0, 0}
}

// Container for enum describing extension setting device types.
type ExtensionSettingDeviceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionSettingDeviceEnum) Reset()         { *m = ExtensionSettingDeviceEnum{} }
func (m *ExtensionSettingDeviceEnum) String() string { return proto.CompactTextString(m) }
func (*ExtensionSettingDeviceEnum) ProtoMessage()    {}
func (*ExtensionSettingDeviceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_be48020c63d81f10, []int{0}
}

func (m *ExtensionSettingDeviceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionSettingDeviceEnum.Unmarshal(m, b)
}
func (m *ExtensionSettingDeviceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionSettingDeviceEnum.Marshal(b, m, deterministic)
}
func (m *ExtensionSettingDeviceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionSettingDeviceEnum.Merge(m, src)
}
func (m *ExtensionSettingDeviceEnum) XXX_Size() int {
	return xxx_messageInfo_ExtensionSettingDeviceEnum.Size(m)
}
func (m *ExtensionSettingDeviceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionSettingDeviceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionSettingDeviceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice", ExtensionSettingDeviceEnum_ExtensionSettingDevice_name, ExtensionSettingDeviceEnum_ExtensionSettingDevice_value)
	proto.RegisterType((*ExtensionSettingDeviceEnum)(nil), "google.ads.googleads.v2.enums.ExtensionSettingDeviceEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/extension_setting_device.proto", fileDescriptor_be48020c63d81f10)
}

var fileDescriptor_be48020c63d81f10 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xeb, 0x60, 0x7f, 0xc8, 0x0e, 0x96, 0x1e, 0x3c, 0x4c, 0x77, 0xd8, 0x3e, 0x40,
	0x02, 0xf5, 0x16, 0xbd, 0xb4, 0x2e, 0x8e, 0x31, 0x6d, 0x0b, 0x73, 0x13, 0xa4, 0x30, 0xea, 0x12,
	0x42, 0x60, 0x4d, 0xca, 0x92, 0x15, 0x3f, 0x8f, 0x47, 0x3f, 0x8a, 0x1f, 0xc5, 0x93, 0x1f, 0x41,
	0x9a, 0xd8, 0x9e, 0xa6, 0x97, 0xf2, 0xd0, 0xe7, 0x7d, 0x7e, 0x79, 0xde, 0x17, 0xdc, 0x70, 0xa5,
	0xf8, 0x9e, 0xa1, 0x82, 0x6a, 0xe4, 0x64, 0xa3, 0xea, 0x10, 0x31, 0x79, 0x2c, 0x35, 0x62, 0xaf,
	0x86, 0x49, 0x2d, 0x94, 0xdc, 0x6a, 0x66, 0x8c, 0x90, 0x7c, 0x4b, 0x59, 0x2d, 0x76, 0x0c, 0x56,
	0x07, 0x65, 0x54, 0x30, 0x76, 0x11, 0x58, 0x50, 0x0d, 0xbb, 0x34, 0xac, 0x43, 0x68, 0xd3, 0xa3,
	0xcb, 0x16, 0x5e, 0x09, 0x54, 0x48, 0xa9, 0x4c, 0x61, 0x84, 0x92, 0xda, 0x85, 0xa7, 0x25, 0x18,
	0x91, 0x16, 0xbf, 0x72, 0xf4, 0x99, 0x85, 0x13, 0x79, 0x2c, 0xa7, 0x29, 0x38, 0x3f, 0xed, 0x06,
	0x67, 0x60, 0xb8, 0x4e, 0x56, 0x19, 0xb9, 0x5d, 0xdc, 0x2d, 0xc8, 0xcc, 0xff, 0x17, 0x0c, 0xc1,
	0xff, 0x75, 0xb2, 0x4c, 0xd2, 0xa7, 0xc4, 0xef, 0x05, 0x00, 0x0c, 0x1e, 0xd2, 0x78, 0x71, 0x4f,
	0x7c, 0xaf, 0x31, 0x66, 0x64, 0xb5, 0x7c, 0x4c, 0x33, 0xbf, 0x1f, 0x7f, 0xf5, 0xc0, 0x64, 0xa7,
	0x4a, 0xf8, 0x67, 0xe5, 0xf8, 0xe2, 0xf4, 0xa3, 0x59, 0xd3, 0x38, 0xeb, 0x3d, 0xc7, 0x3f, 0x69,
	0xae, 0xf6, 0x85, 0xe4, 0x50, 0x1d, 0x38, 0xe2, 0x4c, 0xda, 0x7d, 0xda, 0xf3, 0x55, 0x42, 0xff,
	0x72, 0xcd, 0x6b, 0xfb, 0x7d, 0xf3, 0xfa, 0xf3, 0x28, 0x7a, 0xf7, 0xc6, 0x73, 0x87, 0x8a, 0xa8,
	0x86, 0x4e, 0x36, 0x6a, 0x13, 0xc2, 0x66, 0x7d, 0xfd, 0xd1, 0xfa, 0x79, 0x44, 0x75, 0xde, 0xf9,
	0xf9, 0x26, 0xcc, 0xad, 0xff, 0xe9, 0x4d, 0xdc, 0x4f, 0x8c, 0x23, 0xaa, 0x31, 0xee, 0x26, 0x30,
	0xde, 0x84, 0x18, 0xdb, 0x99, 0x97, 0x81, 0x2d, 0x76, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9b,
	0xbd, 0xaf, 0x6a, 0xe5, 0x01, 0x00, 0x00,
}

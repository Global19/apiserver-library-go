// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/label.proto

package label // import "google.golang.org/genproto/googleapis/api/label"

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

// Value types that can be used as label values.
type LabelDescriptor_ValueType int32

const (
	// A variable-length string. This is the default.
	LabelDescriptor_STRING LabelDescriptor_ValueType = 0
	// Boolean; true or false.
	LabelDescriptor_BOOL LabelDescriptor_ValueType = 1
	// A 64-bit signed integer.
	LabelDescriptor_INT64 LabelDescriptor_ValueType = 2
)

var LabelDescriptor_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
}
var LabelDescriptor_ValueType_value = map[string]int32{
	"STRING": 0,
	"BOOL":   1,
	"INT64":  2,
}

func (x LabelDescriptor_ValueType) String() string {
	return proto.EnumName(LabelDescriptor_ValueType_name, int32(x))
}
func (LabelDescriptor_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_label_d86ce3ed040f9325, []int{0, 0}
}

// A description of a label.
type LabelDescriptor struct {
	// The label key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The type of data that can be assigned to the label.
	ValueType LabelDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=google.api.LabelDescriptor_ValueType" json:"value_type,omitempty"`
	// A human-readable description for the label.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelDescriptor) Reset()         { *m = LabelDescriptor{} }
func (m *LabelDescriptor) String() string { return proto.CompactTextString(m) }
func (*LabelDescriptor) ProtoMessage()    {}
func (*LabelDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_label_d86ce3ed040f9325, []int{0}
}
func (m *LabelDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelDescriptor.Unmarshal(m, b)
}
func (m *LabelDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelDescriptor.Marshal(b, m, deterministic)
}
func (dst *LabelDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelDescriptor.Merge(dst, src)
}
func (m *LabelDescriptor) XXX_Size() int {
	return xxx_messageInfo_LabelDescriptor.Size(m)
}
func (m *LabelDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_LabelDescriptor proto.InternalMessageInfo

func (m *LabelDescriptor) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelDescriptor) GetValueType() LabelDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return LabelDescriptor_STRING
}

func (m *LabelDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*LabelDescriptor)(nil), "google.api.LabelDescriptor")
	proto.RegisterEnum("google.api.LabelDescriptor_ValueType", LabelDescriptor_ValueType_name, LabelDescriptor_ValueType_value)
}

func init() { proto.RegisterFile("google/api/label.proto", fileDescriptor_label_d86ce3ed040f9325) }

var fileDescriptor_label_d86ce3ed040f9325 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x49, 0x4c, 0x4a, 0xcd, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x88, 0xeb, 0x25, 0x16, 0x64, 0x2a, 0xed, 0x64, 0xe4, 0xe2, 0xf7,
	0x01, 0xc9, 0xb9, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94, 0xe4, 0x17, 0x09, 0x09, 0x70, 0x31,
	0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x2e, 0x5c, 0x5c,
	0x65, 0x89, 0x39, 0xa5, 0xa9, 0xf1, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x7c,
	0x46, 0xaa, 0x7a, 0x08, 0x63, 0xf4, 0xd0, 0x8c, 0xd0, 0x0b, 0x03, 0xa9, 0x0e, 0xa9, 0x2c, 0x48,
	0x0d, 0xe2, 0x2c, 0x83, 0x31, 0x85, 0x14, 0xb8, 0xb8, 0x53, 0xa0, 0x4a, 0x32, 0xf3, 0xf3, 0x24,
	0x98, 0xc1, 0xe6, 0x23, 0x0b, 0x29, 0xe9, 0x70, 0x71, 0xc2, 0x75, 0x0a, 0x71, 0x71, 0xb1, 0x05,
	0x87, 0x04, 0x79, 0xfa, 0xb9, 0x0b, 0x30, 0x08, 0x71, 0x70, 0xb1, 0x38, 0xf9, 0xfb, 0xfb, 0x08,
	0x30, 0x0a, 0x71, 0x72, 0xb1, 0x7a, 0xfa, 0x85, 0x98, 0x99, 0x08, 0x30, 0x39, 0xc5, 0x73, 0xf1,
	0x25, 0xe7, 0xe7, 0x22, 0x39, 0xc3, 0x89, 0x0b, 0xec, 0x8e, 0x00, 0x90, 0x2f, 0x03, 0x18, 0xa3,
	0x4c, 0xa1, 0x32, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9,
	0x79, 0xe0, 0x30, 0xd0, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0x23, 0x82, 0xc7, 0x1a, 0x4c, 0xfe,
	0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xe2, 0xee, 0x18, 0xe0, 0x99, 0xc4, 0x06, 0x56, 0x6b, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x04, 0xaa, 0x1f, 0x49, 0x01, 0x00, 0x00,
}

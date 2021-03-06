// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_types.proto

package streaming_protobufs

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

type ConfiguredOpt_Value int32

const (
	ConfiguredOpt_protobuf_unspecified ConfiguredOpt_Value = 0
	ConfiguredOpt_configured           ConfiguredOpt_Value = 1
)

var ConfiguredOpt_Value_name = map[int32]string{
	0: "protobuf_unspecified",
	1: "configured",
}

var ConfiguredOpt_Value_value = map[string]int32{
	"protobuf_unspecified": 0,
	"configured":           1,
}

func (x ConfiguredOpt_Value) String() string {
	return proto.EnumName(ConfiguredOpt_Value_name, int32(x))
}

func (ConfiguredOpt_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{1, 0}
}

type TrueOpt_Value int32

const (
	TrueOpt_protobuf_unspecified TrueOpt_Value = 0
	TrueOpt_true                 TrueOpt_Value = 1
)

var TrueOpt_Value_name = map[int32]string{
	0: "protobuf_unspecified",
	1: "true",
}

var TrueOpt_Value_value = map[string]int32{
	"protobuf_unspecified": 0,
	"true":                 1,
}

func (x TrueOpt_Value) String() string {
	return proto.EnumName(TrueOpt_Value_name, int32(x))
}

func (TrueOpt_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{2, 0}
}

type FalseOpt_Value int32

const (
	FalseOpt_protobuf_unspecified FalseOpt_Value = 0
	FalseOpt_false                FalseOpt_Value = 1
)

var FalseOpt_Value_name = map[int32]string{
	0: "protobuf_unspecified",
	1: "false",
}

var FalseOpt_Value_value = map[string]int32{
	"protobuf_unspecified": 0,
	"false":                1,
}

func (x FalseOpt_Value) String() string {
	return proto.EnumName(FalseOpt_Value_name, int32(x))
}

func (FalseOpt_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{3, 0}
}

type EnabledOpt_Value int32

const (
	EnabledOpt_protobuf_unspecified EnabledOpt_Value = 0
	EnabledOpt_enabled              EnabledOpt_Value = 1
)

var EnabledOpt_Value_name = map[int32]string{
	0: "protobuf_unspecified",
	1: "enabled",
}

var EnabledOpt_Value_value = map[string]int32{
	"protobuf_unspecified": 0,
	"enabled":              1,
}

func (x EnabledOpt_Value) String() string {
	return proto.EnumName(EnabledOpt_Value_name, int32(x))
}

func (EnabledOpt_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{4, 0}
}

type DisabledOpt_Value int32

const (
	DisabledOpt_protobuf_unspecified DisabledOpt_Value = 0
	DisabledOpt_disabled             DisabledOpt_Value = 1
)

var DisabledOpt_Value_name = map[int32]string{
	0: "protobuf_unspecified",
	1: "disabled",
}

var DisabledOpt_Value_value = map[string]int32{
	"protobuf_unspecified": 0,
	"disabled":             1,
}

func (x DisabledOpt_Value) String() string {
	return proto.EnumName(DisabledOpt_Value_name, int32(x))
}

func (DisabledOpt_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{5, 0}
}

type EnabledDisabledOpt_Value int32

const (
	EnabledDisabledOpt_protobuf_unspecified EnabledDisabledOpt_Value = 0
	EnabledDisabledOpt_enabled              EnabledDisabledOpt_Value = 1
	EnabledDisabledOpt_disabled             EnabledDisabledOpt_Value = 2
)

var EnabledDisabledOpt_Value_name = map[int32]string{
	0: "protobuf_unspecified",
	1: "enabled",
	2: "disabled",
}

var EnabledDisabledOpt_Value_value = map[string]int32{
	"protobuf_unspecified": 0,
	"enabled":              1,
	"disabled":             2,
}

func (x EnabledDisabledOpt_Value) String() string {
	return proto.EnumName(EnabledDisabledOpt_Value_name, int32(x))
}

func (EnabledDisabledOpt_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{6, 0}
}

type OnOffOpt_Value int32

const (
	OnOffOpt_protobuf_unspecified OnOffOpt_Value = 0
	OnOffOpt_on                   OnOffOpt_Value = 1
	OnOffOpt_off                  OnOffOpt_Value = 2
)

var OnOffOpt_Value_name = map[int32]string{
	0: "protobuf_unspecified",
	1: "on",
	2: "off",
}

var OnOffOpt_Value_value = map[string]int32{
	"protobuf_unspecified": 0,
	"on":                   1,
	"off":                  2,
}

func (x OnOffOpt_Value) String() string {
	return proto.EnumName(OnOffOpt_Value_name, int32(x))
}

func (OnOffOpt_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{7, 0}
}

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{0}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type ConfiguredOpt struct {
	Value                ConfiguredOpt_Value `protobuf:"varint,1,opt,name=value,proto3,enum=streaming_protobufs.ConfiguredOpt_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConfiguredOpt) Reset()         { *m = ConfiguredOpt{} }
func (m *ConfiguredOpt) String() string { return proto.CompactTextString(m) }
func (*ConfiguredOpt) ProtoMessage()    {}
func (*ConfiguredOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{1}
}

func (m *ConfiguredOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredOpt.Unmarshal(m, b)
}
func (m *ConfiguredOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredOpt.Marshal(b, m, deterministic)
}
func (m *ConfiguredOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredOpt.Merge(m, src)
}
func (m *ConfiguredOpt) XXX_Size() int {
	return xxx_messageInfo_ConfiguredOpt.Size(m)
}
func (m *ConfiguredOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredOpt.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredOpt proto.InternalMessageInfo

func (m *ConfiguredOpt) GetValue() ConfiguredOpt_Value {
	if m != nil {
		return m.Value
	}
	return ConfiguredOpt_protobuf_unspecified
}

type TrueOpt struct {
	Value                TrueOpt_Value `protobuf:"varint,1,opt,name=value,proto3,enum=streaming_protobufs.TrueOpt_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TrueOpt) Reset()         { *m = TrueOpt{} }
func (m *TrueOpt) String() string { return proto.CompactTextString(m) }
func (*TrueOpt) ProtoMessage()    {}
func (*TrueOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{2}
}

func (m *TrueOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrueOpt.Unmarshal(m, b)
}
func (m *TrueOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrueOpt.Marshal(b, m, deterministic)
}
func (m *TrueOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrueOpt.Merge(m, src)
}
func (m *TrueOpt) XXX_Size() int {
	return xxx_messageInfo_TrueOpt.Size(m)
}
func (m *TrueOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_TrueOpt.DiscardUnknown(m)
}

var xxx_messageInfo_TrueOpt proto.InternalMessageInfo

func (m *TrueOpt) GetValue() TrueOpt_Value {
	if m != nil {
		return m.Value
	}
	return TrueOpt_protobuf_unspecified
}

type FalseOpt struct {
	Value                FalseOpt_Value `protobuf:"varint,1,opt,name=value,proto3,enum=streaming_protobufs.FalseOpt_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FalseOpt) Reset()         { *m = FalseOpt{} }
func (m *FalseOpt) String() string { return proto.CompactTextString(m) }
func (*FalseOpt) ProtoMessage()    {}
func (*FalseOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{3}
}

func (m *FalseOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FalseOpt.Unmarshal(m, b)
}
func (m *FalseOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FalseOpt.Marshal(b, m, deterministic)
}
func (m *FalseOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FalseOpt.Merge(m, src)
}
func (m *FalseOpt) XXX_Size() int {
	return xxx_messageInfo_FalseOpt.Size(m)
}
func (m *FalseOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_FalseOpt.DiscardUnknown(m)
}

var xxx_messageInfo_FalseOpt proto.InternalMessageInfo

func (m *FalseOpt) GetValue() FalseOpt_Value {
	if m != nil {
		return m.Value
	}
	return FalseOpt_protobuf_unspecified
}

type EnabledOpt struct {
	Value                EnabledOpt_Value `protobuf:"varint,1,opt,name=value,proto3,enum=streaming_protobufs.EnabledOpt_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EnabledOpt) Reset()         { *m = EnabledOpt{} }
func (m *EnabledOpt) String() string { return proto.CompactTextString(m) }
func (*EnabledOpt) ProtoMessage()    {}
func (*EnabledOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{4}
}

func (m *EnabledOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnabledOpt.Unmarshal(m, b)
}
func (m *EnabledOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnabledOpt.Marshal(b, m, deterministic)
}
func (m *EnabledOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnabledOpt.Merge(m, src)
}
func (m *EnabledOpt) XXX_Size() int {
	return xxx_messageInfo_EnabledOpt.Size(m)
}
func (m *EnabledOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_EnabledOpt.DiscardUnknown(m)
}

var xxx_messageInfo_EnabledOpt proto.InternalMessageInfo

func (m *EnabledOpt) GetValue() EnabledOpt_Value {
	if m != nil {
		return m.Value
	}
	return EnabledOpt_protobuf_unspecified
}

type DisabledOpt struct {
	Value                DisabledOpt_Value `protobuf:"varint,1,opt,name=value,proto3,enum=streaming_protobufs.DisabledOpt_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DisabledOpt) Reset()         { *m = DisabledOpt{} }
func (m *DisabledOpt) String() string { return proto.CompactTextString(m) }
func (*DisabledOpt) ProtoMessage()    {}
func (*DisabledOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{5}
}

func (m *DisabledOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisabledOpt.Unmarshal(m, b)
}
func (m *DisabledOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisabledOpt.Marshal(b, m, deterministic)
}
func (m *DisabledOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisabledOpt.Merge(m, src)
}
func (m *DisabledOpt) XXX_Size() int {
	return xxx_messageInfo_DisabledOpt.Size(m)
}
func (m *DisabledOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_DisabledOpt.DiscardUnknown(m)
}

var xxx_messageInfo_DisabledOpt proto.InternalMessageInfo

func (m *DisabledOpt) GetValue() DisabledOpt_Value {
	if m != nil {
		return m.Value
	}
	return DisabledOpt_protobuf_unspecified
}

type EnabledDisabledOpt struct {
	Value                EnabledDisabledOpt_Value `protobuf:"varint,1,opt,name=value,proto3,enum=streaming_protobufs.EnabledDisabledOpt_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EnabledDisabledOpt) Reset()         { *m = EnabledDisabledOpt{} }
func (m *EnabledDisabledOpt) String() string { return proto.CompactTextString(m) }
func (*EnabledDisabledOpt) ProtoMessage()    {}
func (*EnabledDisabledOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{6}
}

func (m *EnabledDisabledOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnabledDisabledOpt.Unmarshal(m, b)
}
func (m *EnabledDisabledOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnabledDisabledOpt.Marshal(b, m, deterministic)
}
func (m *EnabledDisabledOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnabledDisabledOpt.Merge(m, src)
}
func (m *EnabledDisabledOpt) XXX_Size() int {
	return xxx_messageInfo_EnabledDisabledOpt.Size(m)
}
func (m *EnabledDisabledOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_EnabledDisabledOpt.DiscardUnknown(m)
}

var xxx_messageInfo_EnabledDisabledOpt proto.InternalMessageInfo

func (m *EnabledDisabledOpt) GetValue() EnabledDisabledOpt_Value {
	if m != nil {
		return m.Value
	}
	return EnabledDisabledOpt_protobuf_unspecified
}

type OnOffOpt struct {
	Value                OnOffOpt_Value `protobuf:"varint,1,opt,name=value,proto3,enum=streaming_protobufs.OnOffOpt_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OnOffOpt) Reset()         { *m = OnOffOpt{} }
func (m *OnOffOpt) String() string { return proto.CompactTextString(m) }
func (*OnOffOpt) ProtoMessage()    {}
func (*OnOffOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6f91795c5fbab7, []int{7}
}

func (m *OnOffOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnOffOpt.Unmarshal(m, b)
}
func (m *OnOffOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnOffOpt.Marshal(b, m, deterministic)
}
func (m *OnOffOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnOffOpt.Merge(m, src)
}
func (m *OnOffOpt) XXX_Size() int {
	return xxx_messageInfo_OnOffOpt.Size(m)
}
func (m *OnOffOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_OnOffOpt.DiscardUnknown(m)
}

var xxx_messageInfo_OnOffOpt proto.InternalMessageInfo

func (m *OnOffOpt) GetValue() OnOffOpt_Value {
	if m != nil {
		return m.Value
	}
	return OnOffOpt_protobuf_unspecified
}

func init() {
	proto.RegisterEnum("streaming_protobufs.ConfiguredOpt_Value", ConfiguredOpt_Value_name, ConfiguredOpt_Value_value)
	proto.RegisterEnum("streaming_protobufs.TrueOpt_Value", TrueOpt_Value_name, TrueOpt_Value_value)
	proto.RegisterEnum("streaming_protobufs.FalseOpt_Value", FalseOpt_Value_name, FalseOpt_Value_value)
	proto.RegisterEnum("streaming_protobufs.EnabledOpt_Value", EnabledOpt_Value_name, EnabledOpt_Value_value)
	proto.RegisterEnum("streaming_protobufs.DisabledOpt_Value", DisabledOpt_Value_name, DisabledOpt_Value_value)
	proto.RegisterEnum("streaming_protobufs.EnabledDisabledOpt_Value", EnabledDisabledOpt_Value_name, EnabledDisabledOpt_Value_value)
	proto.RegisterEnum("streaming_protobufs.OnOffOpt_Value", OnOffOpt_Value_name, OnOffOpt_Value_value)
	proto.RegisterType((*Null)(nil), "streaming_protobufs.Null")
	proto.RegisterType((*ConfiguredOpt)(nil), "streaming_protobufs.ConfiguredOpt")
	proto.RegisterType((*TrueOpt)(nil), "streaming_protobufs.TrueOpt")
	proto.RegisterType((*FalseOpt)(nil), "streaming_protobufs.FalseOpt")
	proto.RegisterType((*EnabledOpt)(nil), "streaming_protobufs.EnabledOpt")
	proto.RegisterType((*DisabledOpt)(nil), "streaming_protobufs.DisabledOpt")
	proto.RegisterType((*EnabledDisabledOpt)(nil), "streaming_protobufs.EnabledDisabledOpt")
	proto.RegisterType((*OnOffOpt)(nil), "streaming_protobufs.OnOffOpt")
}

func init() { proto.RegisterFile("common_types.proto", fileDescriptor_3a6f91795c5fbab7) }

var fileDescriptor_3a6f91795c5fbab7 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0xcd, 0x4a, 0xeb, 0x40,
	0x1c, 0x05, 0xf0, 0x3b, 0xbd, 0xfd, 0xc8, 0xfd, 0xf7, 0x5a, 0x86, 0xe8, 0xa2, 0x4b, 0x19, 0x51,
	0x0a, 0x9a, 0x84, 0xb6, 0x1b, 0xc5, 0xe2, 0xc2, 0xaa, 0x4b, 0x0b, 0x22, 0x2e, 0xdc, 0x94, 0x34,
	0x9d, 0x09, 0x81, 0x74, 0x26, 0xcc, 0x87, 0x50, 0x70, 0xe5, 0x4b, 0xf8, 0xba, 0xd2, 0xb4, 0x8d,
	0x1d, 0x88, 0x46, 0x77, 0x49, 0x38, 0x27, 0xbf, 0x93, 0x21, 0xe0, 0x46, 0x62, 0xb1, 0x10, 0x7c,
	0xaa, 0x97, 0x19, 0x55, 0x7e, 0x26, 0x85, 0x16, 0xee, 0xbe, 0xd2, 0x92, 0x86, 0x8b, 0x84, 0xc7,
	0xd3, 0xfc, 0xc1, 0xcc, 0x30, 0x45, 0x9a, 0x50, 0xbf, 0x37, 0x69, 0x4a, 0xde, 0x10, 0xec, 0x8d,
	0x05, 0x67, 0x49, 0x6c, 0x24, 0x9d, 0x4f, 0x32, 0xed, 0x5e, 0x41, 0xe3, 0x25, 0x4c, 0x0d, 0xed,
	0xa2, 0x43, 0xd4, 0xeb, 0x0c, 0x7a, 0x7e, 0x49, 0xdd, 0xb7, 0x2a, 0xfe, 0xd3, 0x2a, 0xff, 0xb0,
	0xae, 0x91, 0x3e, 0x34, 0xf2, 0x7b, 0xb7, 0x0b, 0x07, 0xdb, 0xc2, 0xd4, 0x70, 0x95, 0xd1, 0x28,
	0x61, 0x09, 0x9d, 0xe3, 0x3f, 0x6e, 0x07, 0x20, 0x2a, 0x5e, 0x80, 0x11, 0xc9, 0xa0, 0xf5, 0x28,
	0x0d, 0x5d, 0xe9, 0xe7, 0xb6, 0x4e, 0x4a, 0xf5, 0x4d, 0xd8, 0x76, 0x4f, 0xab, 0x5d, 0x07, 0xea,
	0x5a, 0x1a, 0x8a, 0x11, 0x51, 0xe0, 0xdc, 0x85, 0xa9, 0xca, 0xc9, 0x0b, 0x9b, 0x3c, 0x2a, 0x25,
	0xb7, 0x69, 0xdb, 0x3c, 0xab, 0x36, 0xff, 0x41, 0x83, 0xad, 0xba, 0x18, 0x91, 0x25, 0xc0, 0x2d,
	0x0f, 0x67, 0xe9, 0xfa, 0x9c, 0x2f, 0x6d, 0xf6, 0xb8, 0x94, 0xfd, 0xcc, 0xdb, 0xb0, 0x5f, 0x0d,
	0xb7, 0xa1, 0x45, 0xd7, 0x6d, 0x8c, 0xc8, 0x2b, 0xb4, 0x6f, 0x12, 0x55, 0xd8, 0x23, 0xdb, 0x3e,
	0x29, 0xb5, 0x77, 0x0a, 0x36, 0x1e, 0x54, 0xe3, 0xff, 0xc1, 0x99, 0x6f, 0xea, 0x18, 0x91, 0x77,
	0x04, 0xee, 0xe6, 0x4b, 0x76, 0x57, 0x8c, 0xed, 0x15, 0xde, 0x77, 0x27, 0xf0, 0xe5, 0x98, 0xd1,
	0xef, 0x4e, 0xc2, 0x5a, 0x56, 0x23, 0x4b, 0x70, 0x26, 0x7c, 0xc2, 0xd8, 0x8f, 0xff, 0x83, 0x6d,
	0xda, 0x1e, 0x31, 0xa8, 0x1e, 0xd1, 0x84, 0x9a, 0xe0, 0x18, 0xb9, 0x2d, 0xf8, 0x2b, 0x18, 0xc3,
	0xb5, 0xeb, 0xe1, 0x73, 0x3f, 0xa6, 0x52, 0x26, 0xda, 0x17, 0x9e, 0x0c, 0xb9, 0xa7, 0x22, 0x5f,
	0xc8, 0x38, 0x90, 0x81, 0x4c, 0x22, 0x2f, 0x4b, 0x75, 0x50, 0xf0, 0x5e, 0xc1, 0xcf, 0x9a, 0xf9,
	0xe5, 0xf0, 0x23, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x86, 0x8c, 0x4f, 0xe8, 0x03, 0x00, 0x00,
}

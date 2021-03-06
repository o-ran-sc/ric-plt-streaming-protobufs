// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sgnb_addition_request_reject.proto

package streaming_protobufs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type SgNBAdditionRequestReject struct {
	Id_MeNB_UE_X2AP_ID           uint32                  `protobuf:"varint,1,opt,name=id_MeNB_UE_X2AP_ID,json=idMeNBUEX2APID,proto3" json:"id_MeNB_UE_X2AP_ID,omitempty"`
	Id_SgNB_UE_X2AP_ID           *wrappers.UInt32Value   `protobuf:"bytes,2,opt,name=id_SgNB_UE_X2AP_ID,json=idSgNBUEX2APID,proto3" json:"id_SgNB_UE_X2AP_ID,omitempty"`
	Id_Cause                     *Cause                  `protobuf:"bytes,3,opt,name=id_Cause,json=idCause,proto3" json:"id_Cause,omitempty"`
	Id_CriticalityDiagnostics    *CriticalityDiagnostics `protobuf:"bytes,4,opt,name=id_CriticalityDiagnostics,json=idCriticalityDiagnostics,proto3" json:"id_CriticalityDiagnostics,omitempty"`
	Id_MeNB_UE_X2AP_ID_Extension *wrappers.UInt32Value   `protobuf:"bytes,5,opt,name=id_MeNB_UE_X2AP_ID_Extension,json=idMeNBUEX2APIDExtension,proto3" json:"id_MeNB_UE_X2AP_ID_Extension,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                `json:"-"`
	XXX_unrecognized             []byte                  `json:"-"`
	XXX_sizecache                int32                   `json:"-"`
}

func (m *SgNBAdditionRequestReject) Reset()         { *m = SgNBAdditionRequestReject{} }
func (m *SgNBAdditionRequestReject) String() string { return proto.CompactTextString(m) }
func (*SgNBAdditionRequestReject) ProtoMessage()    {}
func (*SgNBAdditionRequestReject) Descriptor() ([]byte, []int) {
	return fileDescriptor_aec401e36bc1450d, []int{0}
}

func (m *SgNBAdditionRequestReject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SgNBAdditionRequestReject.Unmarshal(m, b)
}
func (m *SgNBAdditionRequestReject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SgNBAdditionRequestReject.Marshal(b, m, deterministic)
}
func (m *SgNBAdditionRequestReject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SgNBAdditionRequestReject.Merge(m, src)
}
func (m *SgNBAdditionRequestReject) XXX_Size() int {
	return xxx_messageInfo_SgNBAdditionRequestReject.Size(m)
}
func (m *SgNBAdditionRequestReject) XXX_DiscardUnknown() {
	xxx_messageInfo_SgNBAdditionRequestReject.DiscardUnknown(m)
}

var xxx_messageInfo_SgNBAdditionRequestReject proto.InternalMessageInfo

func (m *SgNBAdditionRequestReject) GetId_MeNB_UE_X2AP_ID() uint32 {
	if m != nil {
		return m.Id_MeNB_UE_X2AP_ID
	}
	return 0
}

func (m *SgNBAdditionRequestReject) GetId_SgNB_UE_X2AP_ID() *wrappers.UInt32Value {
	if m != nil {
		return m.Id_SgNB_UE_X2AP_ID
	}
	return nil
}

func (m *SgNBAdditionRequestReject) GetId_Cause() *Cause {
	if m != nil {
		return m.Id_Cause
	}
	return nil
}

func (m *SgNBAdditionRequestReject) GetId_CriticalityDiagnostics() *CriticalityDiagnostics {
	if m != nil {
		return m.Id_CriticalityDiagnostics
	}
	return nil
}

func (m *SgNBAdditionRequestReject) GetId_MeNB_UE_X2AP_ID_Extension() *wrappers.UInt32Value {
	if m != nil {
		return m.Id_MeNB_UE_X2AP_ID_Extension
	}
	return nil
}

func init() {
	proto.RegisterType((*SgNBAdditionRequestReject)(nil), "streaming_protobufs.SgNBAdditionRequestReject")
}

func init() { proto.RegisterFile("sgnb_addition_request_reject.proto", fileDescriptor_aec401e36bc1450d) }

var fileDescriptor_aec401e36bc1450d = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xe9, 0xd7, 0xcf, 0x3f, 0x44, 0x14, 0x8c, 0x8b, 0xa6, 0xa5, 0x48, 0xe9, 0xaa, 0x28,
	0x99, 0x60, 0x8a, 0x0f, 0xd0, 0xda, 0x82, 0x5d, 0x58, 0x24, 0x52, 0x11, 0x5d, 0x0c, 0xd3, 0x64,
	0x1c, 0xaf, 0xa4, 0x33, 0xf1, 0xce, 0x04, 0xdb, 0xa7, 0xf0, 0x95, 0x25, 0x13, 0x1b, 0x14, 0xb3,
	0x70, 0x77, 0x39, 0xf7, 0x9e, 0xc3, 0xef, 0x1e, 0xa7, 0xaf, 0x85, 0x5c, 0x52, 0x96, 0x24, 0x60,
	0x40, 0x49, 0x8a, 0xfc, 0x2d, 0xe7, 0xda, 0x50, 0xe4, 0xaf, 0x3c, 0x36, 0x24, 0x43, 0x65, 0x94,
	0x7b, 0xa2, 0x0d, 0x72, 0xb6, 0x02, 0x29, 0xa8, 0x15, 0x96, 0xf9, 0xb3, 0xee, 0x9c, 0x0a, 0xa5,
	0x44, 0xca, 0x83, 0xad, 0x12, 0xbc, 0x23, 0xcb, 0x32, 0x8e, 0xba, 0x34, 0x75, 0x5a, 0xeb, 0x90,
	0x65, 0x34, 0x56, 0xab, 0x95, 0x92, 0xd4, 0x6c, 0x32, 0xbe, 0x5d, 0x1c, 0x73, 0x44, 0x85, 0x34,
	0x66, 0xb9, 0xe6, 0xa5, 0xd4, 0xff, 0x68, 0x3a, 0xed, 0x3b, 0x31, 0x1f, 0x8f, 0xbe, 0x30, 0xa2,
	0x92, 0x22, 0xb2, 0x10, 0xee, 0x99, 0xe3, 0x42, 0x42, 0x6f, 0xf8, 0x7c, 0x4c, 0x17, 0x53, 0xfa,
	0x10, 0x8e, 0x6e, 0xe9, 0x6c, 0xe2, 0x35, 0x7a, 0x8d, 0xc1, 0x61, 0x74, 0x04, 0x49, 0xb1, 0x58,
	0x4c, 0x0b, 0x79, 0x36, 0x71, 0xaf, 0xed, 0x6d, 0x91, 0xf5, 0xfd, 0xf6, 0x5f, 0xaf, 0x31, 0x38,
	0x08, 0xbb, 0xa4, 0x44, 0x26, 0x5b, 0x64, 0xb2, 0x98, 0x49, 0x33, 0x0c, 0xef, 0x59, 0x9a, 0xf3,
	0x22, 0xa9, 0xb0, 0x55, 0x49, 0x97, 0xce, 0x3e, 0x24, 0xf4, 0xaa, 0xa0, 0xf4, 0x9a, 0xd6, 0xdf,
	0x21, 0x35, 0x3d, 0x10, 0x7b, 0x11, 0xed, 0x41, 0x62, 0x07, 0xf7, 0xc5, 0x69, 0x17, 0x36, 0x04,
	0x03, 0x31, 0x4b, 0xc1, 0x6c, 0x26, 0xc0, 0x84, 0x54, 0xda, 0x40, 0xac, 0xbd, 0xff, 0x36, 0xe7,
	0xbc, 0x3e, 0xa7, 0xd6, 0x12, 0x79, 0x90, 0xd4, 0x6f, 0xdc, 0x27, 0xa7, 0xfb, 0xbb, 0x16, 0x3a,
	0x5d, 0x1b, 0x2e, 0x35, 0x28, 0xe9, 0xed, 0xfc, 0xe1, 0xe9, 0xd6, 0xcf, 0xfa, 0x2a, 0xf3, 0x78,
	0xf8, 0x78, 0x21, 0x38, 0x22, 0x18, 0xa2, 0x7c, 0x64, 0xd2, 0xd7, 0x31, 0x51, 0x28, 0x02, 0x0c,
	0x10, 0x62, 0x3f, 0x4b, 0x4d, 0x50, 0xbd, 0xe0, 0x57, 0x2f, 0x2c, 0x77, 0xed, 0x38, 0xfc, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x4f, 0xeb, 0x47, 0x37, 0x54, 0x02, 0x00, 0x00,
}

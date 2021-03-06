// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sgnb_modification_refuse.proto

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

type SgNBModificationRefuse struct {
	ProtocolIEs          *SgNBModificationRefuse_IEs `protobuf:"bytes,1,opt,name=protocolIEs,proto3" json:"protocolIEs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SgNBModificationRefuse) Reset()         { *m = SgNBModificationRefuse{} }
func (m *SgNBModificationRefuse) String() string { return proto.CompactTextString(m) }
func (*SgNBModificationRefuse) ProtoMessage()    {}
func (*SgNBModificationRefuse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbf438e853ff13a, []int{0}
}

func (m *SgNBModificationRefuse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SgNBModificationRefuse.Unmarshal(m, b)
}
func (m *SgNBModificationRefuse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SgNBModificationRefuse.Marshal(b, m, deterministic)
}
func (m *SgNBModificationRefuse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SgNBModificationRefuse.Merge(m, src)
}
func (m *SgNBModificationRefuse) XXX_Size() int {
	return xxx_messageInfo_SgNBModificationRefuse.Size(m)
}
func (m *SgNBModificationRefuse) XXX_DiscardUnknown() {
	xxx_messageInfo_SgNBModificationRefuse.DiscardUnknown(m)
}

var xxx_messageInfo_SgNBModificationRefuse proto.InternalMessageInfo

func (m *SgNBModificationRefuse) GetProtocolIEs() *SgNBModificationRefuse_IEs {
	if m != nil {
		return m.ProtocolIEs
	}
	return nil
}

type SgNBModificationRefuse_IEs struct {
	Id_MeNB_UE_X2AP_ID uint32 `protobuf:"varint,1,opt,name=id_MeNB_UE_X2AP_ID,json=idMeNBUEX2APID,proto3" json:"id_MeNB_UE_X2AP_ID,omitempty"`
	Id_SgNB_UE_X2AP_ID uint32 `protobuf:"varint,2,opt,name=id_SgNB_UE_X2AP_ID,json=idSgNBUEX2APID,proto3" json:"id_SgNB_UE_X2AP_ID,omitempty"`
	Id_Cause           *Cause `protobuf:"bytes,3,opt,name=id_Cause,json=idCause,proto3" json:"id_Cause,omitempty"`
	//The content of id_MeNBtoSgNBContainer has been specified in 3GPP 38.331 to be CG-ConfigInfo.
	Id_MeNBtoSgNBContainer       *CG_ConfigInfo          `protobuf:"bytes,4,opt,name=id_MeNBtoSgNBContainer,json=idMeNBtoSgNBContainer,proto3" json:"id_MeNBtoSgNBContainer,omitempty"`
	Id_CriticalityDiagnostics    *CriticalityDiagnostics `protobuf:"bytes,5,opt,name=id_CriticalityDiagnostics,json=idCriticalityDiagnostics,proto3" json:"id_CriticalityDiagnostics,omitempty"`
	Id_MeNB_UE_X2AP_ID_Extension *wrappers.UInt32Value   `protobuf:"bytes,6,opt,name=id_MeNB_UE_X2AP_ID_Extension,json=idMeNBUEX2APIDExtension,proto3" json:"id_MeNB_UE_X2AP_ID_Extension,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                `json:"-"`
	XXX_unrecognized             []byte                  `json:"-"`
	XXX_sizecache                int32                   `json:"-"`
}

func (m *SgNBModificationRefuse_IEs) Reset()         { *m = SgNBModificationRefuse_IEs{} }
func (m *SgNBModificationRefuse_IEs) String() string { return proto.CompactTextString(m) }
func (*SgNBModificationRefuse_IEs) ProtoMessage()    {}
func (*SgNBModificationRefuse_IEs) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbf438e853ff13a, []int{1}
}

func (m *SgNBModificationRefuse_IEs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SgNBModificationRefuse_IEs.Unmarshal(m, b)
}
func (m *SgNBModificationRefuse_IEs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SgNBModificationRefuse_IEs.Marshal(b, m, deterministic)
}
func (m *SgNBModificationRefuse_IEs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SgNBModificationRefuse_IEs.Merge(m, src)
}
func (m *SgNBModificationRefuse_IEs) XXX_Size() int {
	return xxx_messageInfo_SgNBModificationRefuse_IEs.Size(m)
}
func (m *SgNBModificationRefuse_IEs) XXX_DiscardUnknown() {
	xxx_messageInfo_SgNBModificationRefuse_IEs.DiscardUnknown(m)
}

var xxx_messageInfo_SgNBModificationRefuse_IEs proto.InternalMessageInfo

func (m *SgNBModificationRefuse_IEs) GetId_MeNB_UE_X2AP_ID() uint32 {
	if m != nil {
		return m.Id_MeNB_UE_X2AP_ID
	}
	return 0
}

func (m *SgNBModificationRefuse_IEs) GetId_SgNB_UE_X2AP_ID() uint32 {
	if m != nil {
		return m.Id_SgNB_UE_X2AP_ID
	}
	return 0
}

func (m *SgNBModificationRefuse_IEs) GetId_Cause() *Cause {
	if m != nil {
		return m.Id_Cause
	}
	return nil
}

func (m *SgNBModificationRefuse_IEs) GetId_MeNBtoSgNBContainer() *CG_ConfigInfo {
	if m != nil {
		return m.Id_MeNBtoSgNBContainer
	}
	return nil
}

func (m *SgNBModificationRefuse_IEs) GetId_CriticalityDiagnostics() *CriticalityDiagnostics {
	if m != nil {
		return m.Id_CriticalityDiagnostics
	}
	return nil
}

func (m *SgNBModificationRefuse_IEs) GetId_MeNB_UE_X2AP_ID_Extension() *wrappers.UInt32Value {
	if m != nil {
		return m.Id_MeNB_UE_X2AP_ID_Extension
	}
	return nil
}

func init() {
	proto.RegisterType((*SgNBModificationRefuse)(nil), "streaming_protobufs.SgNBModificationRefuse")
	proto.RegisterType((*SgNBModificationRefuse_IEs)(nil), "streaming_protobufs.SgNBModificationRefuse_IEs")
}

func init() { proto.RegisterFile("sgnb_modification_refuse.proto", fileDescriptor_0fbf438e853ff13a) }

var fileDescriptor_0fbf438e853ff13a = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x6f, 0x8b, 0xd3, 0x40,
	0x10, 0xc6, 0x39, 0x4f, 0x4f, 0xd9, 0x43, 0xc1, 0x88, 0x77, 0x31, 0x1c, 0x87, 0xf4, 0x95, 0x28,
	0x4d, 0xb0, 0xc5, 0x0f, 0x60, 0xdb, 0x20, 0x79, 0x71, 0x87, 0x46, 0x2a, 0xfe, 0x79, 0x31, 0x6c,
	0x37, 0x9b, 0x75, 0x30, 0xdd, 0x09, 0xb3, 0x5b, 0xbc, 0xfb, 0x90, 0x7e, 0x27, 0xc9, 0xb6, 0x0d,
	0x55, 0xe3, 0xbb, 0x65, 0x9e, 0x67, 0x7e, 0xf3, 0xcc, 0xb0, 0xe2, 0xd2, 0x19, 0xbb, 0x82, 0x35,
	0x55, 0x58, 0xa3, 0x92, 0x1e, 0xc9, 0x02, 0xeb, 0x7a, 0xe3, 0x74, 0xda, 0x32, 0x79, 0x8a, 0x9e,
	0x38, 0xcf, 0x5a, 0xae, 0xd1, 0x1a, 0x08, 0x85, 0xd5, 0xa6, 0x76, 0xc9, 0xa5, 0x21, 0x32, 0x8d,
	0xce, 0xf6, 0x95, 0xec, 0x27, 0xcb, 0xb6, 0xd5, 0xec, 0xb6, 0x4d, 0xc9, 0xf9, 0xcd, 0x44, 0xb6,
	0xa0, 0x68, 0xbd, 0x26, 0x0b, 0xfe, 0xb6, 0xd5, 0x7b, 0xe1, 0xb1, 0x66, 0x26, 0x06, 0x25, 0xfb,
	0x01, 0x49, 0xcc, 0xac, 0x40, 0x19, 0x50, 0x64, 0x6b, 0x34, 0x80, 0xb6, 0xa6, 0xad, 0x32, 0xfa,
	0x21, 0xce, 0x3e, 0x9a, 0xeb, 0xd9, 0xd5, 0x41, 0xb6, 0x32, 0x44, 0x8b, 0x3e, 0x88, 0xd3, 0x60,
	0x51, 0xd4, 0x14, 0xb9, 0x8b, 0x8f, 0x9e, 0x1f, 0xbd, 0x38, 0x9d, 0x64, 0xe9, 0x40, 0xd4, 0x74,
	0x98, 0x00, 0x45, 0xee, 0xca, 0x43, 0xc6, 0xe8, 0xd7, 0xb1, 0x48, 0xfe, 0xef, 0x8d, 0x5e, 0x8a,
	0x08, 0x2b, 0xb8, 0xd2, 0xd7, 0x33, 0x58, 0xe6, 0xf0, 0x79, 0xf2, 0xf6, 0x3d, 0x14, 0x8b, 0x30,
	0xf8, 0x61, 0xf9, 0x08, 0xab, 0x4e, 0x58, 0xe6, 0x5d, 0xb9, 0x58, 0xec, 0xbc, 0x1d, 0xec, 0xd0,
	0x7b, 0x67, 0xef, 0xed, 0x84, 0xde, 0xfb, 0x46, 0x3c, 0xc0, 0x0a, 0xe6, 0xdd, 0x3d, 0xe2, 0xe3,
	0xb0, 0x46, 0x32, 0xb8, 0x46, 0x70, 0x94, 0xf7, 0xb1, 0x0a, 0x8f, 0xe8, 0x8b, 0x38, 0xdb, 0xc5,
	0xf1, 0xd4, 0xf1, 0xe6, 0x64, 0xbd, 0x44, 0xab, 0x39, 0xbe, 0x1b, 0x20, 0xa3, 0x61, 0xc8, 0x3b,
	0x98, 0x87, 0x2b, 0x17, 0xb6, 0xa6, 0xf2, 0xe9, 0x36, 0xf6, 0x5f, 0x80, 0xe8, 0xbb, 0x78, 0xd6,
	0x25, 0x62, 0xf4, 0xa8, 0x64, 0x83, 0xfe, 0x76, 0x81, 0xd2, 0x58, 0x72, 0x1e, 0x95, 0x8b, 0xef,
	0x05, 0xfa, 0xab, 0x61, 0xfa, 0x60, 0x4b, 0x19, 0x63, 0x35, 0xac, 0x44, 0xdf, 0xc4, 0xc5, 0xbf,
	0x37, 0x85, 0xfc, 0xc6, 0x6b, 0xeb, 0x90, 0x6c, 0x7c, 0x12, 0x86, 0x5d, 0xa4, 0xdb, 0xcf, 0x96,
	0xee, 0x27, 0xa5, 0xcb, 0xc2, 0xfa, 0xe9, 0xe4, 0x93, 0x6c, 0x36, 0xba, 0x3c, 0xff, 0xf3, 0xf6,
	0x7d, 0xf3, 0x6c, 0xfa, 0xf5, 0xb5, 0xd1, 0xcc, 0xe8, 0x53, 0x1a, 0xb3, 0xb4, 0x63, 0xa7, 0x52,
	0x62, 0x93, 0x71, 0xc6, 0xa8, 0xc6, 0x6d, 0xe3, 0xb3, 0x7e, 0x85, 0x71, 0xbf, 0xc2, 0xea, 0x24,
	0x3c, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xf9, 0x42, 0xb9, 0x15, 0x03, 0x00, 0x00,
}

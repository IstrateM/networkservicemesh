// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectioncontext.proto

package connectioncontext

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

type IpFamily_Family int32

const (
	IpFamily_IPV4 IpFamily_Family = 0
	IpFamily_IPV6 IpFamily_Family = 1
)

var IpFamily_Family_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}

var IpFamily_Family_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x IpFamily_Family) String() string {
	return proto.EnumName(IpFamily_Family_name, int32(x))
}

func (IpFamily_Family) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{2, 0}
}

type IpNeighbor struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	HardwareAddress      string   `protobuf:"bytes,2,opt,name=hardware_address,json=hardwareAddress,proto3" json:"hardware_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpNeighbor) Reset()         { *m = IpNeighbor{} }
func (m *IpNeighbor) String() string { return proto.CompactTextString(m) }
func (*IpNeighbor) ProtoMessage()    {}
func (*IpNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{0}
}

func (m *IpNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpNeighbor.Unmarshal(m, b)
}
func (m *IpNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpNeighbor.Marshal(b, m, deterministic)
}
func (m *IpNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpNeighbor.Merge(m, src)
}
func (m *IpNeighbor) XXX_Size() int {
	return xxx_messageInfo_IpNeighbor.Size(m)
}
func (m *IpNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_IpNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_IpNeighbor proto.InternalMessageInfo

func (m *IpNeighbor) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *IpNeighbor) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

type Route struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type IpFamily struct {
	Family               IpFamily_Family `protobuf:"varint,1,opt,name=family,proto3,enum=connectioncontext.IpFamily_Family" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IpFamily) Reset()         { *m = IpFamily{} }
func (m *IpFamily) String() string { return proto.CompactTextString(m) }
func (*IpFamily) ProtoMessage()    {}
func (*IpFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{2}
}

func (m *IpFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpFamily.Unmarshal(m, b)
}
func (m *IpFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpFamily.Marshal(b, m, deterministic)
}
func (m *IpFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpFamily.Merge(m, src)
}
func (m *IpFamily) XXX_Size() int {
	return xxx_messageInfo_IpFamily.Size(m)
}
func (m *IpFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_IpFamily.DiscardUnknown(m)
}

var xxx_messageInfo_IpFamily proto.InternalMessageInfo

func (m *IpFamily) GetFamily() IpFamily_Family {
	if m != nil {
		return m.Family
	}
	return IpFamily_IPV4
}

type ExtraPrefixRequest struct {
	AddrFamily           *IpFamily `protobuf:"bytes,1,opt,name=addr_family,json=addrFamily,proto3" json:"addr_family,omitempty"`
	PrefixLen            uint32    `protobuf:"varint,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	RequiredNumber       uint32    `protobuf:"varint,3,opt,name=required_number,json=requiredNumber,proto3" json:"required_number,omitempty"`
	RequestedNumber      uint32    `protobuf:"varint,4,opt,name=requested_number,json=requestedNumber,proto3" json:"requested_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExtraPrefixRequest) Reset()         { *m = ExtraPrefixRequest{} }
func (m *ExtraPrefixRequest) String() string { return proto.CompactTextString(m) }
func (*ExtraPrefixRequest) ProtoMessage()    {}
func (*ExtraPrefixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{3}
}

func (m *ExtraPrefixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraPrefixRequest.Unmarshal(m, b)
}
func (m *ExtraPrefixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraPrefixRequest.Marshal(b, m, deterministic)
}
func (m *ExtraPrefixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraPrefixRequest.Merge(m, src)
}
func (m *ExtraPrefixRequest) XXX_Size() int {
	return xxx_messageInfo_ExtraPrefixRequest.Size(m)
}
func (m *ExtraPrefixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraPrefixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraPrefixRequest proto.InternalMessageInfo

func (m *ExtraPrefixRequest) GetAddrFamily() *IpFamily {
	if m != nil {
		return m.AddrFamily
	}
	return nil
}

func (m *ExtraPrefixRequest) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequiredNumber() uint32 {
	if m != nil {
		return m.RequiredNumber
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequestedNumber() uint32 {
	if m != nil {
		return m.RequestedNumber
	}
	return 0
}

type IPContext struct {
	SrcIpAddr            string                `protobuf:"bytes,1,opt,name=src_ip_addr,json=srcIpAddr,proto3" json:"src_ip_addr,omitempty"`
	DstIpAddr            string                `protobuf:"bytes,2,opt,name=dst_ip_addr,json=dstIpAddr,proto3" json:"dst_ip_addr,omitempty"`
	SrcIpRequired        bool                  `protobuf:"varint,3,opt,name=src_ip_required,json=srcIpRequired,proto3" json:"src_ip_required,omitempty"`
	DstIpRequired        bool                  `protobuf:"varint,4,opt,name=dst_ip_required,json=dstIpRequired,proto3" json:"dst_ip_required,omitempty"`
	SrcRoutes            []*Route              `protobuf:"bytes,5,rep,name=src_routes,json=srcRoutes,proto3" json:"src_routes,omitempty"`
	DstRoutes            []*Route              `protobuf:"bytes,6,rep,name=dst_routes,json=dstRoutes,proto3" json:"dst_routes,omitempty"`
	ExcludedPrefixes     []string              `protobuf:"bytes,7,rep,name=excluded_prefixes,json=excludedPrefixes,proto3" json:"excluded_prefixes,omitempty"`
	IpNeighbors          []*IpNeighbor         `protobuf:"bytes,8,rep,name=ip_neighbors,json=ipNeighbors,proto3" json:"ip_neighbors,omitempty"`
	ExtraPrefixRequest   []*ExtraPrefixRequest `protobuf:"bytes,9,rep,name=extra_prefix_request,json=extraPrefixRequest,proto3" json:"extra_prefix_request,omitempty"`
	ExtraPrefixes        []string              `protobuf:"bytes,10,rep,name=extra_prefixes,json=extraPrefixes,proto3" json:"extra_prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IPContext) Reset()         { *m = IPContext{} }
func (m *IPContext) String() string { return proto.CompactTextString(m) }
func (*IPContext) ProtoMessage()    {}
func (*IPContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{4}
}

func (m *IPContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPContext.Unmarshal(m, b)
}
func (m *IPContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPContext.Marshal(b, m, deterministic)
}
func (m *IPContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPContext.Merge(m, src)
}
func (m *IPContext) XXX_Size() int {
	return xxx_messageInfo_IPContext.Size(m)
}
func (m *IPContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IPContext.DiscardUnknown(m)
}

var xxx_messageInfo_IPContext proto.InternalMessageInfo

func (m *IPContext) GetSrcIpAddr() string {
	if m != nil {
		return m.SrcIpAddr
	}
	return ""
}

func (m *IPContext) GetDstIpAddr() string {
	if m != nil {
		return m.DstIpAddr
	}
	return ""
}

func (m *IPContext) GetSrcIpRequired() bool {
	if m != nil {
		return m.SrcIpRequired
	}
	return false
}

func (m *IPContext) GetDstIpRequired() bool {
	if m != nil {
		return m.DstIpRequired
	}
	return false
}

func (m *IPContext) GetSrcRoutes() []*Route {
	if m != nil {
		return m.SrcRoutes
	}
	return nil
}

func (m *IPContext) GetDstRoutes() []*Route {
	if m != nil {
		return m.DstRoutes
	}
	return nil
}

func (m *IPContext) GetExcludedPrefixes() []string {
	if m != nil {
		return m.ExcludedPrefixes
	}
	return nil
}

func (m *IPContext) GetIpNeighbors() []*IpNeighbor {
	if m != nil {
		return m.IpNeighbors
	}
	return nil
}

func (m *IPContext) GetExtraPrefixRequest() []*ExtraPrefixRequest {
	if m != nil {
		return m.ExtraPrefixRequest
	}
	return nil
}

func (m *IPContext) GetExtraPrefixes() []string {
	if m != nil {
		return m.ExtraPrefixes
	}
	return nil
}

type ConnectionContext struct {
	IpContext            *IPContext `protobuf:"bytes,1,opt,name=ip_context,json=ipContext,proto3" json:"ip_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConnectionContext) Reset()         { *m = ConnectionContext{} }
func (m *ConnectionContext) String() string { return proto.CompactTextString(m) }
func (*ConnectionContext) ProtoMessage()    {}
func (*ConnectionContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{5}
}

func (m *ConnectionContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionContext.Unmarshal(m, b)
}
func (m *ConnectionContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionContext.Marshal(b, m, deterministic)
}
func (m *ConnectionContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionContext.Merge(m, src)
}
func (m *ConnectionContext) XXX_Size() int {
	return xxx_messageInfo_ConnectionContext.Size(m)
}
func (m *ConnectionContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionContext.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionContext proto.InternalMessageInfo

func (m *ConnectionContext) GetIpContext() *IPContext {
	if m != nil {
		return m.IpContext
	}
	return nil
}

func init() {
	proto.RegisterEnum("connectioncontext.IpFamily_Family", IpFamily_Family_name, IpFamily_Family_value)
	proto.RegisterType((*IpNeighbor)(nil), "connectioncontext.IpNeighbor")
	proto.RegisterType((*Route)(nil), "connectioncontext.Route")
	proto.RegisterType((*IpFamily)(nil), "connectioncontext.IpFamily")
	proto.RegisterType((*ExtraPrefixRequest)(nil), "connectioncontext.ExtraPrefixRequest")
	proto.RegisterType((*IPContext)(nil), "connectioncontext.IPContext")
	proto.RegisterType((*ConnectionContext)(nil), "connectioncontext.ConnectionContext")
}

func init() { proto.RegisterFile("connectioncontext.proto", fileDescriptor_c30b3f1555e8b686) }

var fileDescriptor_c30b3f1555e8b686 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x8b, 0x13, 0x41,
	0x10, 0x35, 0x1f, 0x1b, 0x33, 0x15, 0xf3, 0xd5, 0x8a, 0x0e, 0xb8, 0xab, 0x61, 0x60, 0x35, 0x22,
	0xe4, 0x10, 0x45, 0x41, 0x3d, 0xa8, 0x8b, 0x4a, 0x40, 0x96, 0xd0, 0x07, 0x05, 0x2f, 0x43, 0x32,
	0x5d, 0xeb, 0x36, 0x64, 0x67, 0x7a, 0xbb, 0x3b, 0x18, 0x7f, 0xa0, 0xff, 0xca, 0x83, 0x74, 0x77,
	0xcd, 0x18, 0x98, 0xa0, 0xa7, 0x54, 0xaa, 0x5e, 0xbd, 0x57, 0x5d, 0xf5, 0x12, 0xb8, 0x97, 0x15,
	0x79, 0x8e, 0x99, 0x95, 0x45, 0x9e, 0x15, 0xb9, 0xc5, 0x9d, 0x9d, 0x29, 0x5d, 0xd8, 0x82, 0x8d,
	0x6b, 0x85, 0xe4, 0x13, 0xc0, 0x42, 0x9d, 0xa3, 0xfc, 0x7e, 0xb9, 0x2e, 0x34, 0x1b, 0x40, 0x53,
	0xaa, 0xb8, 0x31, 0x69, 0x4c, 0x23, 0xde, 0x94, 0x8a, 0x3d, 0x81, 0xd1, 0xe5, 0x4a, 0x8b, 0x1f,
	0x2b, 0x8d, 0xe9, 0x4a, 0x08, 0x8d, 0xc6, 0xc4, 0x4d, 0x5f, 0x1d, 0x96, 0xf9, 0x77, 0x21, 0x9d,
	0x3c, 0x84, 0x23, 0x5e, 0x6c, 0x2d, 0xb2, 0xbb, 0xd0, 0x51, 0x1a, 0x2f, 0xe4, 0x8e, 0x78, 0xe8,
	0x5b, 0x22, 0xa0, 0xbb, 0x50, 0x1f, 0x57, 0x57, 0x72, 0xf3, 0x93, 0xbd, 0x82, 0xce, 0x85, 0x8f,
	0x3c, 0x66, 0x30, 0x4f, 0x66, 0xf5, 0x91, 0x4b, 0xf0, 0x2c, 0x7c, 0x70, 0xea, 0x48, 0x8e, 0xa1,
	0x43, 0x2c, 0x5d, 0x68, 0x2f, 0x96, 0x5f, 0x9e, 0x8f, 0x6e, 0x50, 0xf4, 0x62, 0xd4, 0x48, 0x7e,
	0x35, 0x80, 0x7d, 0xd8, 0x59, 0xbd, 0x5a, 0x7a, 0x55, 0x8e, 0xd7, 0x5b, 0x34, 0x96, 0xbd, 0x81,
	0x9e, 0x9b, 0x3f, 0xdd, 0x53, 0xed, 0xcd, 0xef, 0xff, 0x43, 0x95, 0x83, 0xc3, 0x93, 0xd0, 0x09,
	0x40, 0x78, 0x44, 0xba, 0xc1, 0xdc, 0x2f, 0xa0, 0xcf, 0xa3, 0x90, 0xf9, 0x8c, 0x39, 0x7b, 0x0c,
	0x43, 0x8d, 0xd7, 0x5b, 0xa9, 0x51, 0xa4, 0xf9, 0xf6, 0x6a, 0x8d, 0x3a, 0x6e, 0x79, 0xcc, 0xa0,
	0x4c, 0x9f, 0xfb, 0xac, 0x5b, 0xa7, 0x0e, 0x03, 0xfd, 0x45, 0xb6, 0x3d, 0x72, 0x58, 0xe5, 0x03,
	0x34, 0xf9, 0xdd, 0x82, 0x68, 0xb1, 0x3c, 0x0b, 0x53, 0xb1, 0x07, 0xd0, 0x33, 0x3a, 0x4b, 0xa5,
	0xf2, 0x57, 0xa0, 0xc5, 0x46, 0x46, 0x67, 0x0b, 0xe5, 0xf6, 0xef, 0xea, 0xc2, 0xd8, 0xaa, 0x1e,
	0x4e, 0x14, 0x09, 0x63, 0xa9, 0xfe, 0x08, 0x86, 0xd4, 0x5f, 0x4e, 0xe4, 0x27, 0xec, 0xf2, 0xbe,
	0xe7, 0xe0, 0x94, 0x74, 0x38, 0xe2, 0xa9, 0x70, 0xed, 0x80, 0xf3, 0x5c, 0x15, 0xee, 0x25, 0x80,
	0xe3, 0xd3, 0xee, 0xe0, 0x26, 0x3e, 0x9a, 0xb4, 0xa6, 0xbd, 0x79, 0x7c, 0x60, 0x9b, 0xde, 0x11,
	0x7e, 0x50, 0x1f, 0x19, 0xd7, 0xe8, 0x04, 0xa8, 0xb1, 0xf3, 0xbf, 0x46, 0x61, 0x2c, 0x35, 0x3e,
	0x85, 0x31, 0xee, 0xb2, 0xcd, 0x56, 0xa0, 0x48, 0xc3, 0xe6, 0xd1, 0xc4, 0x37, 0x27, 0xad, 0x69,
	0xc4, 0x47, 0x65, 0x61, 0x49, 0x79, 0xf6, 0x16, 0x6e, 0x49, 0x95, 0xe6, 0xe4, 0x6a, 0x13, 0x77,
	0xbd, 0xce, 0xc9, 0xc1, 0x73, 0x97, 0xde, 0xe7, 0x3d, 0x59, 0xc5, 0x86, 0x7d, 0x85, 0x3b, 0xe8,
	0x5c, 0x44, 0x5a, 0x29, 0x9d, 0x27, 0x8e, 0x3c, 0xd3, 0xe9, 0x01, 0xa6, 0xba, 0xe9, 0x38, 0xc3,
	0xba, 0x11, 0x4f, 0x61, 0xb0, 0x4f, 0x8c, 0x26, 0x06, 0xff, 0x88, 0xfe, 0x1e, 0x16, 0x4d, 0xb2,
	0x84, 0xf1, 0x59, 0x25, 0x51, 0xba, 0xe0, 0x35, 0x80, 0x54, 0x29, 0x09, 0x92, 0x87, 0x8f, 0x0f,
	0x3d, 0xaa, 0xf4, 0x0d, 0x8f, 0xa4, 0xa2, 0xf0, 0xfd, 0xed, 0x6f, 0xf5, 0x5f, 0xff, 0xba, 0xe3,
	0xff, 0x17, 0x9e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x43, 0x10, 0xda, 0xf1, 0x32, 0x04, 0x00,
	0x00,
}

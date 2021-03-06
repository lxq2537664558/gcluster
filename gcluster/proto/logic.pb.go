// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic.proto

package netproto

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

// 错误码
type E_Code int32

const (
	E_Code_E_ERROR                 E_Code = 0
	E_Code_E_OK                    E_Code = 1
	E_Code_E_RELOGIN               E_Code = 2
	E_Code_E_NONE_EXIST            E_Code = 3
	E_Code_E_UNKNOWN               E_Code = 4
	E_Code_E_SERVER_INTERNAL_ERROR E_Code = 5
	E_Code_E_INVALID_PARAM         E_Code = 6
	E_Code_E_INVALID_OPT           E_Code = 7
)

var E_Code_name = map[int32]string{
	0: "E_ERROR",
	1: "E_OK",
	2: "E_RELOGIN",
	3: "E_NONE_EXIST",
	4: "E_UNKNOWN",
	5: "E_SERVER_INTERNAL_ERROR",
	6: "E_INVALID_PARAM",
	7: "E_INVALID_OPT",
}
var E_Code_value = map[string]int32{
	"E_ERROR":                 0,
	"E_OK":                    1,
	"E_RELOGIN":               2,
	"E_NONE_EXIST":            3,
	"E_UNKNOWN":               4,
	"E_SERVER_INTERNAL_ERROR": 5,
	"E_INVALID_PARAM":         6,
	"E_INVALID_OPT":           7,
}

func (x E_Code) Enum() *E_Code {
	p := new(E_Code)
	*p = x
	return p
}
func (x E_Code) String() string {
	return proto.EnumName(E_Code_name, int32(x))
}
func (x *E_Code) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(E_Code_value, data, "E_Code")
	if err != nil {
		return err
	}
	*x = E_Code(value)
	return nil
}
func (E_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{0}
}

type EMsgIds int32

const (
	EMsgIds_ECS_None             EMsgIds = 0
	EMsgIds_ECS_TestSyncBaseInfo EMsgIds = 90
	EMsgIds_ESC_TestSyncBaseInfo EMsgIds = 91
	EMsgIds_ECS_Login            EMsgIds = 101
	EMsgIds_ESC_Login            EMsgIds = 103
	EMsgIds_ECS_Tick             EMsgIds = 104
	EMsgIds_ESC_Tick             EMsgIds = 105
	EMsgIds_ECS_P2P              EMsgIds = 106
	EMsgIds_ESC_P2P              EMsgIds = 107
	EMsgIds_ECS_GuideDonate      EMsgIds = 108
)

var EMsgIds_name = map[int32]string{
	0:   "ECS_None",
	90:  "ECS_TestSyncBaseInfo",
	91:  "ESC_TestSyncBaseInfo",
	101: "ECS_Login",
	103: "ESC_Login",
	104: "ECS_Tick",
	105: "ESC_Tick",
	106: "ECS_P2P",
	107: "ESC_P2P",
	108: "ECS_GuideDonate",
}
var EMsgIds_value = map[string]int32{
	"ECS_None":             0,
	"ECS_TestSyncBaseInfo": 90,
	"ESC_TestSyncBaseInfo": 91,
	"ECS_Login":            101,
	"ESC_Login":            103,
	"ECS_Tick":             104,
	"ESC_Tick":             105,
	"ECS_P2P":              106,
	"ESC_P2P":              107,
	"ECS_GuideDonate":      108,
}

func (x EMsgIds) Enum() *EMsgIds {
	p := new(EMsgIds)
	*p = x
	return p
}
func (x EMsgIds) String() string {
	return proto.EnumName(EMsgIds_name, int32(x))
}
func (x *EMsgIds) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EMsgIds_value, data, "EMsgIds")
	if err != nil {
		return err
	}
	*x = EMsgIds(value)
	return nil
}
func (EMsgIds) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{1}
}

type CS_TestSyncBaseInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_TestSyncBaseInfo) Reset()         { *m = CS_TestSyncBaseInfo{} }
func (m *CS_TestSyncBaseInfo) String() string { return proto.CompactTextString(m) }
func (*CS_TestSyncBaseInfo) ProtoMessage()    {}
func (*CS_TestSyncBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{0}
}
func (m *CS_TestSyncBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_TestSyncBaseInfo.Unmarshal(m, b)
}
func (m *CS_TestSyncBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_TestSyncBaseInfo.Marshal(b, m, deterministic)
}
func (dst *CS_TestSyncBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_TestSyncBaseInfo.Merge(dst, src)
}
func (m *CS_TestSyncBaseInfo) XXX_Size() int {
	return xxx_messageInfo_CS_TestSyncBaseInfo.Size(m)
}
func (m *CS_TestSyncBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_TestSyncBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CS_TestSyncBaseInfo proto.InternalMessageInfo

type SC_TestSyncBaseInfo struct {
	Code                 *E_Code  `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_TestSyncBaseInfo) Reset()         { *m = SC_TestSyncBaseInfo{} }
func (m *SC_TestSyncBaseInfo) String() string { return proto.CompactTextString(m) }
func (*SC_TestSyncBaseInfo) ProtoMessage()    {}
func (*SC_TestSyncBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{1}
}
func (m *SC_TestSyncBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_TestSyncBaseInfo.Unmarshal(m, b)
}
func (m *SC_TestSyncBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_TestSyncBaseInfo.Marshal(b, m, deterministic)
}
func (dst *SC_TestSyncBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_TestSyncBaseInfo.Merge(dst, src)
}
func (m *SC_TestSyncBaseInfo) XXX_Size() int {
	return xxx_messageInfo_SC_TestSyncBaseInfo.Size(m)
}
func (m *SC_TestSyncBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_TestSyncBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SC_TestSyncBaseInfo proto.InternalMessageInfo

func (m *SC_TestSyncBaseInfo) GetCode() E_Code {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return E_Code_E_ERROR
}

type CS_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_Tick) Reset()         { *m = CS_Tick{} }
func (m *CS_Tick) String() string { return proto.CompactTextString(m) }
func (*CS_Tick) ProtoMessage()    {}
func (*CS_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{2}
}
func (m *CS_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_Tick.Unmarshal(m, b)
}
func (m *CS_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_Tick.Marshal(b, m, deterministic)
}
func (dst *CS_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_Tick.Merge(dst, src)
}
func (m *CS_Tick) XXX_Size() int {
	return xxx_messageInfo_CS_Tick.Size(m)
}
func (m *CS_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_CS_Tick proto.InternalMessageInfo

type SC_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_Tick) Reset()         { *m = SC_Tick{} }
func (m *SC_Tick) String() string { return proto.CompactTextString(m) }
func (*SC_Tick) ProtoMessage()    {}
func (*SC_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{3}
}
func (m *SC_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Tick.Unmarshal(m, b)
}
func (m *SC_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Tick.Marshal(b, m, deterministic)
}
func (dst *SC_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Tick.Merge(dst, src)
}
func (m *SC_Tick) XXX_Size() int {
	return xxx_messageInfo_SC_Tick.Size(m)
}
func (m *SC_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Tick proto.InternalMessageInfo

// 登录
type CS_Login struct {
	Token                *string  `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_Login) Reset()         { *m = CS_Login{} }
func (m *CS_Login) String() string { return proto.CompactTextString(m) }
func (*CS_Login) ProtoMessage()    {}
func (*CS_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{4}
}
func (m *CS_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_Login.Unmarshal(m, b)
}
func (m *CS_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_Login.Marshal(b, m, deterministic)
}
func (dst *CS_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_Login.Merge(dst, src)
}
func (m *CS_Login) XXX_Size() int {
	return xxx_messageInfo_CS_Login.Size(m)
}
func (m *CS_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_Login.DiscardUnknown(m)
}

var xxx_messageInfo_CS_Login proto.InternalMessageInfo

func (m *CS_Login) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// 登录返回
type SC_Login struct {
	Id                   *int64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Code                 *E_Code  `protobuf:"varint,2,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_Login) Reset()         { *m = SC_Login{} }
func (m *SC_Login) String() string { return proto.CompactTextString(m) }
func (*SC_Login) ProtoMessage()    {}
func (*SC_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{5}
}
func (m *SC_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Login.Unmarshal(m, b)
}
func (m *SC_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Login.Marshal(b, m, deterministic)
}
func (dst *SC_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Login.Merge(dst, src)
}
func (m *SC_Login) XXX_Size() int {
	return xxx_messageInfo_SC_Login.Size(m)
}
func (m *SC_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Login.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Login proto.InternalMessageInfo

func (m *SC_Login) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SC_Login) GetCode() E_Code {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return E_Code_E_ERROR
}

// p2p消息
type CS_P2P struct {
	From                 *int64   `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	To                   *int64   `protobuf:"varint,2,opt,name=to" json:"to,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_P2P) Reset()         { *m = CS_P2P{} }
func (m *CS_P2P) String() string { return proto.CompactTextString(m) }
func (*CS_P2P) ProtoMessage()    {}
func (*CS_P2P) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{6}
}
func (m *CS_P2P) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_P2P.Unmarshal(m, b)
}
func (m *CS_P2P) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_P2P.Marshal(b, m, deterministic)
}
func (dst *CS_P2P) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_P2P.Merge(dst, src)
}
func (m *CS_P2P) XXX_Size() int {
	return xxx_messageInfo_CS_P2P.Size(m)
}
func (m *CS_P2P) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_P2P.DiscardUnknown(m)
}

var xxx_messageInfo_CS_P2P proto.InternalMessageInfo

func (m *CS_P2P) GetFrom() int64 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *CS_P2P) GetTo() int64 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *CS_P2P) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// p2p消息
type SC_P2P struct {
	From                 *int64   `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	To                   *int64   `protobuf:"varint,2,opt,name=to" json:"to,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_P2P) Reset()         { *m = SC_P2P{} }
func (m *SC_P2P) String() string { return proto.CompactTextString(m) }
func (*SC_P2P) ProtoMessage()    {}
func (*SC_P2P) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{7}
}
func (m *SC_P2P) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_P2P.Unmarshal(m, b)
}
func (m *SC_P2P) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_P2P.Marshal(b, m, deterministic)
}
func (dst *SC_P2P) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_P2P.Merge(dst, src)
}
func (m *SC_P2P) XXX_Size() int {
	return xxx_messageInfo_SC_P2P.Size(m)
}
func (m *SC_P2P) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_P2P.DiscardUnknown(m)
}

var xxx_messageInfo_SC_P2P proto.InternalMessageInfo

func (m *SC_P2P) GetFrom() int64 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *SC_P2P) GetTo() int64 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *SC_P2P) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type CS_GuideDonate struct {
	GuideId              *int64   `protobuf:"varint,1,opt,name=guideId" json:"guideId,omitempty"`
	Gold                 *int32   `protobuf:"varint,2,opt,name=gold" json:"gold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_GuideDonate) Reset()         { *m = CS_GuideDonate{} }
func (m *CS_GuideDonate) String() string { return proto.CompactTextString(m) }
func (*CS_GuideDonate) ProtoMessage()    {}
func (*CS_GuideDonate) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{8}
}
func (m *CS_GuideDonate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_GuideDonate.Unmarshal(m, b)
}
func (m *CS_GuideDonate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_GuideDonate.Marshal(b, m, deterministic)
}
func (dst *CS_GuideDonate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_GuideDonate.Merge(dst, src)
}
func (m *CS_GuideDonate) XXX_Size() int {
	return xxx_messageInfo_CS_GuideDonate.Size(m)
}
func (m *CS_GuideDonate) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_GuideDonate.DiscardUnknown(m)
}

var xxx_messageInfo_CS_GuideDonate proto.InternalMessageInfo

func (m *CS_GuideDonate) GetGuideId() int64 {
	if m != nil && m.GuideId != nil {
		return *m.GuideId
	}
	return 0
}

func (m *CS_GuideDonate) GetGold() int32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

type CS_None struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_None) Reset()         { *m = CS_None{} }
func (m *CS_None) String() string { return proto.CompactTextString(m) }
func (*CS_None) ProtoMessage()    {}
func (*CS_None) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_14d3b48eb1a0a7b8, []int{9}
}
func (m *CS_None) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_None.Unmarshal(m, b)
}
func (m *CS_None) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_None.Marshal(b, m, deterministic)
}
func (dst *CS_None) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_None.Merge(dst, src)
}
func (m *CS_None) XXX_Size() int {
	return xxx_messageInfo_CS_None.Size(m)
}
func (m *CS_None) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_None.DiscardUnknown(m)
}

var xxx_messageInfo_CS_None proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CS_TestSyncBaseInfo)(nil), "netproto.CS_TestSyncBaseInfo")
	proto.RegisterType((*SC_TestSyncBaseInfo)(nil), "netproto.SC_TestSyncBaseInfo")
	proto.RegisterType((*CS_Tick)(nil), "netproto.CS_Tick")
	proto.RegisterType((*SC_Tick)(nil), "netproto.SC_Tick")
	proto.RegisterType((*CS_Login)(nil), "netproto.CS_Login")
	proto.RegisterType((*SC_Login)(nil), "netproto.SC_Login")
	proto.RegisterType((*CS_P2P)(nil), "netproto.CS_P2P")
	proto.RegisterType((*SC_P2P)(nil), "netproto.SC_P2P")
	proto.RegisterType((*CS_GuideDonate)(nil), "netproto.CS_GuideDonate")
	proto.RegisterType((*CS_None)(nil), "netproto.CS_None")
	proto.RegisterEnum("netproto.E_Code", E_Code_name, E_Code_value)
	proto.RegisterEnum("netproto.EMsgIds", EMsgIds_name, EMsgIds_value)
}

func init() { proto.RegisterFile("logic.proto", fileDescriptor_logic_14d3b48eb1a0a7b8) }

var fileDescriptor_logic_14d3b48eb1a0a7b8 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6b, 0x3b, 0x89, 0x9d, 0x69, 0x1a, 0x86, 0x4d, 0x11, 0x91, 0xb8, 0x44, 0x11, 0x87,
	0xaa, 0x87, 0x1c, 0x7a, 0x45, 0x42, 0x04, 0x77, 0x54, 0x59, 0x4d, 0xd7, 0xd1, 0x6e, 0x28, 0x08,
	0x0e, 0xab, 0x2a, 0xde, 0x1a, 0x93, 0xd4, 0x8b, 0x6a, 0x73, 0xe0, 0x3d, 0x78, 0x14, 0x1e, 0x10,
	0xed, 0xda, 0x06, 0x21, 0x2a, 0x2e, 0xbd, 0x7d, 0xdf, 0xfe, 0xf9, 0xcd, 0xce, 0xce, 0x07, 0x87,
	0x7b, 0x93, 0x17, 0xdb, 0xc5, 0xd7, 0x7b, 0x53, 0x1b, 0x16, 0x95, 0xba, 0x76, 0x6a, 0xfe, 0x0c,
	0x26, 0xb1, 0x54, 0x1b, 0x5d, 0xd5, 0xf2, 0x7b, 0xb9, 0x7d, 0x7b, 0x53, 0xe9, 0xa4, 0xbc, 0x35,
	0xf3, 0x57, 0x30, 0x91, 0xf1, 0x3f, 0xcb, 0xec, 0x25, 0xf4, 0xb6, 0x26, 0xd3, 0x53, 0x6f, 0xe6,
	0x9d, 0x8c, 0xcf, 0x70, 0xd1, 0x61, 0x16, 0xa4, 0x62, 0x93, 0x69, 0xe1, 0x76, 0xe7, 0x43, 0x08,
	0x2d, 0xb3, 0xd8, 0xee, 0xac, 0xb4, 0x1c, 0x2b, 0x67, 0x10, 0xc5, 0x52, 0xad, 0x4c, 0x5e, 0x94,
	0xec, 0x18, 0xfa, 0xb5, 0xd9, 0xe9, 0xd2, 0x81, 0x86, 0xa2, 0x31, 0xf3, 0x37, 0x10, 0xc9, 0xb8,
	0x3d, 0x31, 0x06, 0xbf, 0xc8, 0xdc, 0x76, 0x20, 0xfc, 0x22, 0xfb, 0x5d, 0xd9, 0xff, 0x6f, 0xe5,
	0xd7, 0x30, 0x88, 0xa5, 0x5a, 0x9f, 0xad, 0x19, 0x83, 0xde, 0xed, 0xbd, 0xb9, 0x6b, 0x09, 0x4e,
	0x5b, 0x66, 0x6d, 0x1c, 0x21, 0x10, 0x7e, 0x6d, 0x18, 0x42, 0x70, 0x57, 0xe5, 0xd3, 0x60, 0xe6,
	0x9d, 0x8c, 0x84, 0x95, 0xf6, 0xbe, 0x8c, 0x1f, 0x75, 0x7f, 0x1c, 0x4b, 0x75, 0xf1, 0xad, 0xc8,
	0xf4, 0xb9, 0x29, 0x6f, 0x6a, 0xcd, 0xa6, 0x10, 0xe6, 0xd6, 0x26, 0x5d, 0x33, 0x9d, 0xb5, 0x15,
	0x72, 0xb3, 0xcf, 0x1c, 0xaf, 0x2f, 0x9c, 0x6e, 0x7f, 0x8e, 0x9b, 0x52, 0x9f, 0xfe, 0xf0, 0x60,
	0xd0, 0xf4, 0xc6, 0x0e, 0x21, 0x24, 0x45, 0x42, 0xa4, 0x02, 0x0f, 0x58, 0x04, 0x3d, 0x52, 0xe9,
	0x25, 0x7a, 0xec, 0x08, 0x86, 0xa4, 0x04, 0xad, 0xd2, 0x8b, 0x84, 0xa3, 0xcf, 0x10, 0x46, 0xa4,
	0x78, 0xca, 0x49, 0xd1, 0x87, 0x44, 0x6e, 0x30, 0x68, 0x0e, 0xbc, 0xe3, 0x97, 0x3c, 0x7d, 0xcf,
	0xb1, 0xc7, 0x5e, 0xc0, 0x73, 0x52, 0x92, 0xc4, 0x35, 0x09, 0x95, 0xf0, 0x0d, 0x09, 0xbe, 0x5c,
	0xb5, 0xd8, 0x3e, 0x9b, 0xc0, 0x13, 0x52, 0x09, 0xbf, 0x5e, 0xae, 0x92, 0x73, 0xb5, 0x5e, 0x8a,
	0xe5, 0x15, 0x0e, 0xd8, 0x53, 0x38, 0xfa, 0xb3, 0x98, 0xae, 0x37, 0x18, 0x9e, 0xfe, 0xf4, 0x20,
	0xa4, 0xab, 0x2a, 0x4f, 0xb2, 0x8a, 0x8d, 0x20, 0xa2, 0xf6, 0xb9, 0x78, 0xc0, 0xa6, 0x70, 0x4c,
	0x0f, 0x44, 0x09, 0x3f, 0xba, 0x9d, 0x07, 0xd2, 0x84, 0x9f, 0xdc, 0x0b, 0xbb, 0x50, 0xa0, 0x76,
	0xb6, 0x4b, 0x00, 0xe6, 0x1d, 0xdf, 0xa6, 0x07, 0x3f, 0x3b, 0xd7, 0x66, 0x09, 0x0b, 0xf7, 0x27,
	0xcd, 0xa8, 0xf1, 0x8b, 0x33, 0xcd, 0xdc, 0x70, 0xe7, 0x3a, 0xf9, 0x7b, 0x08, 0xb8, 0xff, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xb7, 0xec, 0x86, 0x85, 0xfe, 0x02, 0x00, 0x00,
}

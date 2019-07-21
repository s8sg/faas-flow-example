// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helper/forwarding/types.proto

package forwarding // import "github.com/hashicorp/vault/helper/forwarding"

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

type Request struct {
	// Not used right now but reserving in case it turns out that streaming
	// makes things more economical on the gRPC side
	// uint64 id = 1;
	Method               string                  `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	Url                  *URL                    `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	HeaderEntries        map[string]*HeaderEntry `protobuf:"bytes,4,rep,name=header_entries,json=headerEntries" json:"header_entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body                 []byte                  `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Host                 string                  `protobuf:"bytes,6,opt,name=host" json:"host,omitempty"`
	RemoteAddr           string                  `protobuf:"bytes,7,opt,name=remote_addr,json=remoteAddr" json:"remote_addr,omitempty"`
	PeerCertificates     [][]byte                `protobuf:"bytes,8,rep,name=peer_certificates,json=peerCertificates,proto3" json:"peer_certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_6ebfa235129f89d8, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetUrl() *URL {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Request) GetHeaderEntries() map[string]*HeaderEntry {
	if m != nil {
		return m.HeaderEntries
	}
	return nil
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Request) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *Request) GetPeerCertificates() [][]byte {
	if m != nil {
		return m.PeerCertificates
	}
	return nil
}

type URL struct {
	Scheme string `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	Opaque string `protobuf:"bytes,2,opt,name=opaque" json:"opaque,omitempty"`
	// This isn't needed now but might be in the future, so we'll skip the
	// number to keep the ordering in net/url
	// UserInfo user = 3;
	Host    string `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	Path    string `protobuf:"bytes,5,opt,name=path" json:"path,omitempty"`
	RawPath string `protobuf:"bytes,6,opt,name=raw_path,json=rawPath" json:"raw_path,omitempty"`
	// This also isn't needed right now, but we'll reserve the number
	// bool force_query = 7;
	RawQuery             string   `protobuf:"bytes,8,opt,name=raw_query,json=rawQuery" json:"raw_query,omitempty"`
	Fragment             string   `protobuf:"bytes,9,opt,name=fragment" json:"fragment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_6ebfa235129f89d8, []int{1}
}
func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (dst *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(dst, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *URL) GetOpaque() string {
	if m != nil {
		return m.Opaque
	}
	return ""
}

func (m *URL) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *URL) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *URL) GetRawPath() string {
	if m != nil {
		return m.RawPath
	}
	return ""
}

func (m *URL) GetRawQuery() string {
	if m != nil {
		return m.RawQuery
	}
	return ""
}

func (m *URL) GetFragment() string {
	if m != nil {
		return m.Fragment
	}
	return ""
}

type HeaderEntry struct {
	Values               []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderEntry) Reset()         { *m = HeaderEntry{} }
func (m *HeaderEntry) String() string { return proto.CompactTextString(m) }
func (*HeaderEntry) ProtoMessage()    {}
func (*HeaderEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_6ebfa235129f89d8, []int{2}
}
func (m *HeaderEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderEntry.Unmarshal(m, b)
}
func (m *HeaderEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderEntry.Marshal(b, m, deterministic)
}
func (dst *HeaderEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderEntry.Merge(dst, src)
}
func (m *HeaderEntry) XXX_Size() int {
	return xxx_messageInfo_HeaderEntry.Size(m)
}
func (m *HeaderEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderEntry.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderEntry proto.InternalMessageInfo

func (m *HeaderEntry) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Response struct {
	// Not used right now but reserving in case it turns out that streaming
	// makes things more economical on the gRPC side
	// uint64 id = 1;
	StatusCode uint32 `protobuf:"varint,2,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Body       []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Added in 0.6.2 to ensure that the content-type is set appropriately, as
	// well as any other information
	HeaderEntries        map[string]*HeaderEntry `protobuf:"bytes,4,rep,name=header_entries,json=headerEntries" json:"header_entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_6ebfa235129f89d8, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() uint32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetHeaderEntries() map[string]*HeaderEntry {
	if m != nil {
		return m.HeaderEntries
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "forwarding.Request")
	proto.RegisterMapType((map[string]*HeaderEntry)(nil), "forwarding.Request.HeaderEntriesEntry")
	proto.RegisterType((*URL)(nil), "forwarding.URL")
	proto.RegisterType((*HeaderEntry)(nil), "forwarding.HeaderEntry")
	proto.RegisterType((*Response)(nil), "forwarding.Response")
	proto.RegisterMapType((map[string]*HeaderEntry)(nil), "forwarding.Response.HeaderEntriesEntry")
}

func init() {
	proto.RegisterFile("helper/forwarding/types.proto", fileDescriptor_types_6ebfa235129f89d8)
}

var fileDescriptor_types_6ebfa235129f89d8 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0x55, 0x9a, 0x6e, 0xff, 0x4c, 0x77, 0x7f, 0xbf, 0xc5, 0x07, 0x30, 0x8b, 0x10, 0xa1, 0x12,
	0x10, 0x09, 0x48, 0xa4, 0x72, 0x41, 0xdc, 0x60, 0x85, 0xc4, 0x61, 0x41, 0x60, 0x69, 0x0f, 0x70,
	0x89, 0xdc, 0x78, 0x5a, 0x47, 0x34, 0x75, 0x6a, 0x3b, 0x5b, 0xe5, 0x63, 0xf1, 0x9d, 0x38, 0xf1,
	0x29, 0x90, 0x9d, 0xb0, 0x0d, 0x5a, 0x21, 0x4e, 0x9c, 0x32, 0xef, 0xbd, 0xc9, 0x78, 0xde, 0x8c,
	0x0d, 0xf7, 0x25, 0x6e, 0x2a, 0xd4, 0xe9, 0x4a, 0xe9, 0x3d, 0xd7, 0xa2, 0xd8, 0xae, 0x53, 0xdb,
	0x54, 0x68, 0x92, 0x4a, 0x2b, 0xab, 0x08, 0x1c, 0xf8, 0xf9, 0xf7, 0x01, 0x8c, 0x19, 0xee, 0x6a,
	0x34, 0x96, 0xdc, 0x86, 0x51, 0x89, 0x56, 0x2a, 0x41, 0x07, 0x51, 0x10, 0x4f, 0x59, 0x87, 0xc8,
	0x43, 0x08, 0x6b, 0xbd, 0xa1, 0x61, 0x14, 0xc4, 0xb3, 0xc5, 0xff, 0xc9, 0xe1, 0xef, 0xe4, 0x92,
	0x5d, 0x30, 0xa7, 0x91, 0xf7, 0xf0, 0x9f, 0x44, 0x2e, 0x50, 0x67, 0xb8, 0xb5, 0xba, 0x40, 0x43,
	0x87, 0x51, 0x18, 0xcf, 0x16, 0x8f, 0xfb, 0xd9, 0xdd, 0x39, 0xc9, 0x3b, 0x9f, 0xf9, 0xb6, 0x4d,
	0x74, 0x9f, 0x86, 0x9d, 0xc8, 0x3e, 0x47, 0x08, 0x0c, 0x97, 0x4a, 0x34, 0xf4, 0x28, 0x0a, 0xe2,
	0x63, 0xe6, 0x63, 0xc7, 0x49, 0x65, 0x2c, 0x1d, 0xf9, 0xde, 0x7c, 0x4c, 0x1e, 0xc0, 0x4c, 0x63,
	0xa9, 0x2c, 0x66, 0x5c, 0x08, 0x4d, 0xc7, 0x5e, 0x82, 0x96, 0x7a, 0x2d, 0x84, 0x26, 0x4f, 0xe1,
	0x56, 0x85, 0xa8, 0xb3, 0x1c, 0xb5, 0x2d, 0x56, 0x45, 0xce, 0x2d, 0x1a, 0x3a, 0x89, 0xc2, 0xf8,
	0x98, 0x9d, 0x3a, 0xe1, 0xbc, 0xc7, 0x9f, 0x7d, 0x06, 0x72, 0xb3, 0x35, 0x72, 0x0a, 0xe1, 0x57,
	0x6c, 0x68, 0xe0, 0x6b, 0xbb, 0x90, 0x3c, 0x87, 0xa3, 0x2b, 0xbe, 0xa9, 0xd1, 0x8f, 0x69, 0xb6,
	0xb8, 0xd3, 0xf7, 0x78, 0x28, 0xd0, 0xb0, 0x36, 0xeb, 0xd5, 0xe0, 0x65, 0x30, 0xff, 0x16, 0x40,
	0x78, 0xc9, 0x2e, 0xdc, 0x88, 0x4d, 0x2e, 0xb1, 0xc4, 0xae, 0x5e, 0x87, 0x1c, 0xaf, 0x2a, 0xbe,
	0xeb, 0x6a, 0x4e, 0x59, 0x87, 0xae, 0x4d, 0x0f, 0x7b, 0xa6, 0x09, 0x0c, 0x2b, 0x6e, 0xa5, 0x1f,
	0xce, 0x94, 0xf9, 0x98, 0xdc, 0x85, 0x89, 0xe6, 0xfb, 0xcc, 0xf3, 0xed, 0x80, 0xc6, 0x9a, 0xef,
	0x3f, 0x3a, 0xe9, 0x1e, 0x4c, 0x9d, 0xb4, 0xab, 0x51, 0x37, 0x74, 0xe2, 0x35, 0x97, 0xfb, 0xc9,
	0x61, 0x72, 0x06, 0x93, 0x95, 0xe6, 0xeb, 0x12, 0xb7, 0x96, 0x4e, 0x5b, 0xed, 0x17, 0x9e, 0x3f,
	0x82, 0x59, 0xcf, 0x8d, 0x6b, 0xd1, 0xfb, 0x31, 0x34, 0x88, 0x42, 0xd7, 0x62, 0x8b, 0xe6, 0x3f,
	0x02, 0x98, 0x30, 0x34, 0x95, 0xda, 0x1a, 0x74, 0x0b, 0x31, 0x96, 0xdb, 0xda, 0x64, 0xb9, 0x12,
	0xad, 0x99, 0x13, 0x06, 0x2d, 0x75, 0xae, 0x04, 0x5e, 0x6f, 0x36, 0xec, 0x6d, 0xf6, 0xc3, 0x1f,
	0x2e, 0xcf, 0x93, 0xdf, 0x2f, 0x4f, 0x7b, 0xc4, 0xdf, 0x6f, 0xcf, 0x3f, 0xdc, 0xe3, 0x9b, 0xe4,
	0xcb, 0xb3, 0x75, 0x61, 0x65, 0xbd, 0x4c, 0x72, 0x55, 0xa6, 0x92, 0x1b, 0x59, 0xe4, 0x4a, 0x57,
	0xe9, 0x15, 0xaf, 0x37, 0x36, 0xbd, 0xf1, 0xec, 0x96, 0x23, 0xff, 0xe2, 0x5e, 0xfc, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0x03, 0xfa, 0xd9, 0x51, 0x92, 0x03, 0x00, 0x00,
}

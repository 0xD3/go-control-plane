// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/local_rate_limit/v2alpha/local_rate_limit.proto

package envoy_config_filter_network_local_rate_limit_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type LocalRateLimit struct {
	StatPrefix           string                   `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	TokenBucket          *_type.TokenBucket       `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
	RuntimeEnabled       *core.RuntimeFeatureFlag `protobuf:"bytes,3,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LocalRateLimit) Reset()         { *m = LocalRateLimit{} }
func (m *LocalRateLimit) String() string { return proto.CompactTextString(m) }
func (*LocalRateLimit) ProtoMessage()    {}
func (*LocalRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_984b0e5e5a865836, []int{0}
}

func (m *LocalRateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalRateLimit.Unmarshal(m, b)
}
func (m *LocalRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalRateLimit.Marshal(b, m, deterministic)
}
func (m *LocalRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalRateLimit.Merge(m, src)
}
func (m *LocalRateLimit) XXX_Size() int {
	return xxx_messageInfo_LocalRateLimit.Size(m)
}
func (m *LocalRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_LocalRateLimit proto.InternalMessageInfo

func (m *LocalRateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *LocalRateLimit) GetTokenBucket() *_type.TokenBucket {
	if m != nil {
		return m.TokenBucket
	}
	return nil
}

func (m *LocalRateLimit) GetRuntimeEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.RuntimeEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalRateLimit)(nil), "envoy.config.filter.network.local_rate_limit.v2alpha.LocalRateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/local_rate_limit/v2alpha/local_rate_limit.proto", fileDescriptor_984b0e5e5a865836)
}

var fileDescriptor_984b0e5e5a865836 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0xea, 0xf5, 0x4f, 0x2a, 0x57, 0x89, 0x8b, 0x5b, 0x8a, 0x4a, 0x11, 0x84, 0xae,
	0x66, 0x20, 0xd5, 0x17, 0x18, 0xb4, 0x1b, 0x8b, 0x94, 0xe0, 0x3e, 0x9c, 0xa4, 0xa7, 0x75, 0xe8,
	0x74, 0x66, 0x98, 0x9c, 0xc4, 0xf6, 0x15, 0xdc, 0xb8, 0x13, 0x1f, 0xcb, 0x67, 0x71, 0xd9, 0x85,
	0xc8, 0x64, 0x46, 0x54, 0xea, 0xca, 0x5d, 0xc8, 0x2f, 0xe7, 0x77, 0xbe, 0x7c, 0x27, 0x7b, 0x8b,
	0xa6, 0xb7, 0x27, 0xd1, 0x58, 0xb3, 0x55, 0x3b, 0xb1, 0x55, 0x9a, 0xd0, 0x0b, 0x83, 0xf4, 0xd1,
	0xfa, 0xbd, 0xd0, 0xb6, 0x01, 0x5d, 0x79, 0x20, 0xac, 0xb4, 0x3a, 0x28, 0x12, 0x7d, 0x01, 0xda,
	0x7d, 0x80, 0x0b, 0xc0, 0x9d, 0xb7, 0x64, 0xf3, 0x97, 0x83, 0x8c, 0x47, 0x19, 0x8f, 0x32, 0x9e,
	0x64, 0xfc, 0x62, 0x26, 0xc9, 0xa6, 0x4f, 0x62, 0x04, 0x70, 0x4a, 0xf4, 0x85, 0x68, 0xac, 0x47,
	0x51, 0x43, 0x8b, 0xd1, 0x39, 0x7d, 0x1a, 0x29, 0x9d, 0x1c, 0x0a, 0xb2, 0x7b, 0x34, 0x55, 0xdd,
	0x35, 0x7b, 0x4c, 0x2b, 0xa7, 0xcf, 0xba, 0x8d, 0x03, 0x01, 0xc6, 0x58, 0x02, 0x52, 0xd6, 0xb4,
	0xe2, 0xa0, 0x76, 0x61, 0x49, 0xe2, 0x37, 0x3d, 0x68, 0xb5, 0x01, 0x42, 0xf1, 0xeb, 0x21, 0x82,
	0xe7, 0xdf, 0x58, 0x76, 0xbd, 0x0a, 0x91, 0x4a, 0x20, 0x5c, 0x85, 0x40, 0xf9, 0x3c, 0x1b, 0xb7,
	0x04, 0x54, 0x39, 0x8f, 0x5b, 0x75, 0x9c, 0xb0, 0x19, 0x9b, 0xdf, 0x97, 0x77, 0xcf, 0xf2, 0xb6,
	0x1f, 0xcd, 0x58, 0x99, 0x05, 0xb6, 0x1e, 0x50, 0xfe, 0x3a, 0x7b, 0xf0, 0x67, 0x96, 0xc9, 0x68,
	0xc6, 0xe6, 0xe3, 0xe2, 0x86, 0xc7, 0xff, 0x0f, 0x59, 0xf9, 0xfb, 0xc0, 0xe5, 0x80, 0xe5, 0xbd,
	0xb3, 0xbc, 0xfa, 0xc4, 0x46, 0x8f, 0x58, 0x39, 0xa6, 0xdf, 0xaf, 0xf3, 0x77, 0xd9, 0x43, 0xdf,
	0x19, 0x52, 0x07, 0xac, 0xd0, 0x40, 0xad, 0x71, 0x33, 0xb9, 0x35, 0x88, 0x5e, 0x24, 0x11, 0x38,
	0xc5, 0xfb, 0x82, 0x87, 0x4a, 0x78, 0x19, 0xbf, 0x5c, 0x22, 0x50, 0xe7, 0x71, 0xa9, 0x61, 0x57,
	0x5e, 0xa7, 0xe9, 0x37, 0x71, 0x58, 0x7e, 0x61, 0xdf, 0xbf, 0xfe, 0xf8, 0x7c, 0xf5, 0x2a, 0x5f,
	0xc4, 0x71, 0x3c, 0x12, 0x9a, 0x36, 0x94, 0x92, 0x6e, 0xd1, 0xfe, 0xe3, 0x18, 0xe9, 0x16, 0x8b,
	0x4c, 0x2a, 0x1b, 0xd7, 0x3a, 0x6f, 0x8f, 0x27, 0xfe, 0x3f, 0xa7, 0x94, 0x8f, 0xff, 0x6e, 0x74,
	0x1d, 0x9a, 0x5e, 0xb3, 0xfa, 0xce, 0x50, 0xf9, 0xe2, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0x5f, 0x97, 0xde, 0x6d, 0x02, 0x00, 0x00,
}

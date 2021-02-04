// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blog.proto

package blog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/services/posts/proto"
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

type LatestRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LatestRequest) Reset()         { *m = LatestRequest{} }
func (m *LatestRequest) String() string { return proto.CompactTextString(m) }
func (*LatestRequest) ProtoMessage()    {}
func (*LatestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0}
}

func (m *LatestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatestRequest.Unmarshal(m, b)
}
func (m *LatestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatestRequest.Marshal(b, m, deterministic)
}
func (m *LatestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestRequest.Merge(m, src)
}
func (m *LatestRequest) XXX_Size() int {
	return xxx_messageInfo_LatestRequest.Size(m)
}
func (m *LatestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LatestRequest proto.InternalMessageInfo

type LatestResponse struct {
	Latest               *proto1.Post `protobuf:"bytes,1,opt,name=latest,proto3" json:"latest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LatestResponse) Reset()         { *m = LatestResponse{} }
func (m *LatestResponse) String() string { return proto.CompactTextString(m) }
func (*LatestResponse) ProtoMessage()    {}
func (*LatestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{1}
}

func (m *LatestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatestResponse.Unmarshal(m, b)
}
func (m *LatestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatestResponse.Marshal(b, m, deterministic)
}
func (m *LatestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestResponse.Merge(m, src)
}
func (m *LatestResponse) XXX_Size() int {
	return xxx_messageInfo_LatestResponse.Size(m)
}
func (m *LatestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LatestResponse proto.InternalMessageInfo

func (m *LatestResponse) GetLatest() *proto1.Post {
	if m != nil {
		return m.Latest
	}
	return nil
}

type PostsRequest struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostsRequest) Reset()         { *m = PostsRequest{} }
func (m *PostsRequest) String() string { return proto.CompactTextString(m) }
func (*PostsRequest) ProtoMessage()    {}
func (*PostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{2}
}

func (m *PostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostsRequest.Unmarshal(m, b)
}
func (m *PostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostsRequest.Marshal(b, m, deterministic)
}
func (m *PostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostsRequest.Merge(m, src)
}
func (m *PostsRequest) XXX_Size() int {
	return xxx_messageInfo_PostsRequest.Size(m)
}
func (m *PostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostsRequest proto.InternalMessageInfo

func (m *PostsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PostsRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type PostsResponse struct {
	Posts                []*proto1.Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PostsResponse) Reset()         { *m = PostsResponse{} }
func (m *PostsResponse) String() string { return proto.CompactTextString(m) }
func (*PostsResponse) ProtoMessage()    {}
func (*PostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{3}
}

func (m *PostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostsResponse.Unmarshal(m, b)
}
func (m *PostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostsResponse.Marshal(b, m, deterministic)
}
func (m *PostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostsResponse.Merge(m, src)
}
func (m *PostsResponse) XXX_Size() int {
	return xxx_messageInfo_PostsResponse.Size(m)
}
func (m *PostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostsResponse proto.InternalMessageInfo

func (m *PostsResponse) GetPosts() []*proto1.Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func init() {
	proto.RegisterType((*LatestRequest)(nil), "blog.LatestRequest")
	proto.RegisterType((*LatestResponse)(nil), "blog.LatestResponse")
	proto.RegisterType((*PostsRequest)(nil), "blog.PostsRequest")
	proto.RegisterType((*PostsResponse)(nil), "blog.PostsResponse")
}

func init() { proto.RegisterFile("proto/blog.proto", fileDescriptor_fc5203cdc85000bc) }

var fileDescriptor_fc5203cdc85000bc = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x25, 0xb4, 0xc9, 0x70, 0xa5, 0x80, 0xae, 0x15, 0x8a, 0x32, 0x15, 0xb3, 0x30, 0xc5, 0x22,
	0xa8, 0x1b, 0x13, 0x33, 0x03, 0xca, 0x1f, 0x90, 0xc8, 0x0d, 0x96, 0x1c, 0x2e, 0xcd, 0xb9, 0x7c,
	0x3f, 0xaa, 0xcf, 0x91, 0x48, 0xb7, 0xf7, 0xde, 0xf9, 0xbd, 0x77, 0x3e, 0xb8, 0x1f, 0x46, 0xf2,
	0xa4, 0x1b, 0x47, 0x5d, 0x19, 0x20, 0x2e, 0xcf, 0xb8, 0x78, 0xe9, 0xac, 0xff, 0x3e, 0x35, 0x65,
	0x4b, 0xbd, 0xee, 0x6d, 0x3b, 0x92, 0x66, 0x33, 0xfe, 0xda, 0xd6, 0xb0, 0x1e, 0x88, 0x3d, 0x6b,
	0xf1, 0x05, 0x2c, 0x46, 0x75, 0x07, 0xeb, 0x8f, 0x2f, 0x6f, 0xd8, 0xd7, 0xe6, 0x78, 0x32, 0xec,
	0xd5, 0x1e, 0x6e, 0x27, 0x81, 0x07, 0xfa, 0x61, 0x83, 0x4f, 0x90, 0xb9, 0xa0, 0xe4, 0xc9, 0x2e,
	0x79, 0x5e, 0x55, 0xab, 0x52, 0x02, 0x3e, 0x89, 0x7d, 0x1d, 0x47, 0xea, 0x0d, 0x6e, 0xce, 0x9c,
	0x63, 0x0c, 0x6e, 0x21, 0x75, 0xb6, 0xb7, 0xe2, 0x59, 0xd4, 0x42, 0xf0, 0x01, 0x32, 0x3a, 0x1c,
	0xd8, 0xf8, 0xfc, 0x3a, 0xc8, 0x91, 0xa9, 0x0a, 0xd6, 0xd1, 0x1d, 0x3b, 0x1f, 0x21, 0x0d, 0x25,
	0x79, 0xb2, 0x5b, 0x5c, 0x56, 0xca, 0xa4, 0x3a, 0xc2, 0xf2, 0xdd, 0x51, 0x87, 0x7b, 0xc8, 0x64,
	0x61, 0xdc, 0x94, 0xe1, 0x22, 0xb3, 0xff, 0x14, 0xdb, 0xb9, 0x28, 0xf9, 0xea, 0x0a, 0x2b, 0x48,
	0x43, 0x25, 0xa2, 0x3c, 0xf8, 0xbf, 0x7d, 0xb1, 0x99, 0x69, 0x93, 0xa7, 0xc9, 0xc2, 0xcd, 0x5e,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x04, 0x74, 0x9b, 0x75, 0x80, 0x01, 0x00, 0x00,
}
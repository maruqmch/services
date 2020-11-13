// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/chat/proto/chat.proto

package chat

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

// NewRequest contains the infromation needed to create a new chat
type NewRequest struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4935707190cdb38c, []int{0}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

// NewResponse contains the chat id for the users
type NewResponse struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResponse) Reset()         { *m = NewResponse{} }
func (m *NewResponse) String() string { return proto.CompactTextString(m) }
func (*NewResponse) ProtoMessage()    {}
func (*NewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4935707190cdb38c, []int{1}
}

func (m *NewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResponse.Unmarshal(m, b)
}
func (m *NewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResponse.Marshal(b, m, deterministic)
}
func (m *NewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResponse.Merge(m, src)
}
func (m *NewResponse) XXX_Size() int {
	return xxx_messageInfo_NewResponse.Size(m)
}
func (m *NewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewResponse proto.InternalMessageInfo

func (m *NewResponse) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

// HistoryRequest contains the id of the chat we want the history for. This RPC will return all
// historical messages, however in a real life application we'd introduce some form of pagination
// here, only loading the older messages when required.
type HistoryRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryRequest) Reset()         { *m = HistoryRequest{} }
func (m *HistoryRequest) String() string { return proto.CompactTextString(m) }
func (*HistoryRequest) ProtoMessage()    {}
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4935707190cdb38c, []int{2}
}

func (m *HistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryRequest.Unmarshal(m, b)
}
func (m *HistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryRequest.Marshal(b, m, deterministic)
}
func (m *HistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryRequest.Merge(m, src)
}
func (m *HistoryRequest) XXX_Size() int {
	return xxx_messageInfo_HistoryRequest.Size(m)
}
func (m *HistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryRequest proto.InternalMessageInfo

func (m *HistoryRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

// HistoryResponse contains the historical messages in a chat
type HistoryResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HistoryResponse) Reset()         { *m = HistoryResponse{} }
func (m *HistoryResponse) String() string { return proto.CompactTextString(m) }
func (*HistoryResponse) ProtoMessage()    {}
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4935707190cdb38c, []int{3}
}

func (m *HistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryResponse.Unmarshal(m, b)
}
func (m *HistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryResponse.Marshal(b, m, deterministic)
}
func (m *HistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryResponse.Merge(m, src)
}
func (m *HistoryResponse) XXX_Size() int {
	return xxx_messageInfo_HistoryResponse.Size(m)
}
func (m *HistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryResponse proto.InternalMessageInfo

func (m *HistoryResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// SendRequest contains a single message to send to a chat
type SendRequest struct {
	// a client side id, should be validated by the server to make the request retry safe
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// id of the chat the message is being sent to / from
	ChatId string `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// id of the user who sent the message
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// subject of the message
	Subject string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	// text of the message
	Text                 string   `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4935707190cdb38c, []int{4}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *SendRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *SendRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SendRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// SendResponse is a blank message returned when a message is successfully created
type SendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4935707190cdb38c, []int{5}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

// Message sent to a chat
type Message struct {
	// id of the message, allocated by the server
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// a client side id, should be validated by the server to make the request retry safe
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// id of the chat the message is being sent to / from
	ChatId string `protobuf:"bytes,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// id of the user who sent the message
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// time time the message was sent in unix format
	SentAt int64 `protobuf:"varint,5,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	// subject of the message
	Subject string `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	// text of the message
	Text                 string   `protobuf:"bytes,7,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4935707190cdb38c, []int{6}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Message) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *Message) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Message) GetSentAt() int64 {
	if m != nil {
		return m.SentAt
	}
	return 0
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*NewRequest)(nil), "chat.NewRequest")
	proto.RegisterType((*NewResponse)(nil), "chat.NewResponse")
	proto.RegisterType((*HistoryRequest)(nil), "chat.HistoryRequest")
	proto.RegisterType((*HistoryResponse)(nil), "chat.HistoryResponse")
	proto.RegisterType((*SendRequest)(nil), "chat.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "chat.SendResponse")
	proto.RegisterType((*Message)(nil), "chat.Message")
}

func init() {
	proto.RegisterFile("github.com/micro/services/chat/proto/chat.proto", fileDescriptor_4935707190cdb38c)
}

var fileDescriptor_4935707190cdb38c = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0xb5, 0xb6, 0xeb, 0x85, 0xa1, 0xa5, 0xed, 0xaa, 0x15, 0xae, 0x7b, 0x41, 0x3e, 0xb4,
	0x50, 0x54, 0x5c, 0x51, 0x29, 0x97, 0xe4, 0x92, 0x70, 0x09, 0x87, 0x70, 0x70, 0x6e, 0xb9, 0x20,
	0x63, 0xaf, 0x60, 0xa3, 0x60, 0x13, 0xef, 0x3a, 0x24, 0xaf, 0x90, 0x97, 0xc9, 0x5b, 0xe4, 0xb9,
	0xa2, 0xfd, 0x03, 0xd8, 0x51, 0x90, 0x72, 0x9b, 0xf9, 0x66, 0x3c, 0xfb, 0xfb, 0x76, 0xd6, 0x10,
	0x2e, 0x98, 0x58, 0x96, 0xf3, 0x61, 0x92, 0xaf, 0xc2, 0x15, 0x4b, 0x8a, 0x3c, 0xe4, 0xb4, 0xb8,
	0x63, 0x09, 0xe5, 0x61, 0xb2, 0x8c, 0x45, 0xb8, 0x2e, 0x72, 0x91, 0xab, 0x70, 0xa8, 0x42, 0xe2,
	0xc8, 0x38, 0xf8, 0x0d, 0x30, 0xa5, 0x9b, 0x88, 0xde, 0x96, 0x94, 0x0b, 0xf2, 0x03, 0x1a, 0x25,
	0xa7, 0xc5, 0x8c, 0xa5, 0xdc, 0x43, 0x5d, 0xbb, 0xd7, 0x8c, 0xb0, 0xcc, 0x27, 0x29, 0x0f, 0x7e,
	0x41, 0x4b, 0x35, 0xf2, 0x75, 0x9e, 0x71, 0x4a, 0x3a, 0x80, 0xe5, 0xf7, 0x33, 0x96, 0x7a, 0xa8,
	0x8b, 0x7a, 0xcd, 0xc8, 0x95, 0xe9, 0x24, 0x0d, 0xfa, 0xd0, 0x3e, 0x67, 0x5c, 0xe4, 0xc5, 0xc3,
	0x76, 0xe8, 0xc1, 0xd6, 0x13, 0xf8, 0xbc, 0x6b, 0x35, 0x63, 0xfb, 0xd0, 0x58, 0x51, 0xce, 0xe3,
	0x05, 0xd5, 0x00, 0xad, 0xd1, 0xa7, 0xa1, 0x62, 0xbe, 0xd0, 0x6a, 0xb4, 0x2b, 0x07, 0x8f, 0x08,
	0x5a, 0x97, 0x34, 0x4b, 0xb7, 0xc7, 0xfc, 0x84, 0x66, 0x72, 0xc3, 0x68, 0x56, 0x39, 0xa8, 0xa1,
	0x85, 0x49, 0x5a, 0x65, 0xb0, 0xaa, 0x0c, 0xb2, 0x60, 0x1c, 0x7b, 0xb6, 0x2e, 0x68, 0xc3, 0xc4,
	0x03, 0xcc, 0xcb, 0xf9, 0x35, 0x4d, 0x84, 0xe7, 0xa8, 0xc2, 0x36, 0x25, 0x04, 0x1c, 0x41, 0xef,
	0x85, 0xf7, 0x41, 0xc9, 0x2a, 0x0e, 0xda, 0xf0, 0x51, 0xb3, 0x68, 0x1f, 0xc1, 0x13, 0x02, 0x6c,
	0x90, 0x49, 0x1b, 0xac, 0x1d, 0x91, 0xc5, 0xd2, 0x3a, 0xa8, 0x75, 0x18, 0xd4, 0x3e, 0x04, 0xea,
	0xd4, 0x40, 0x3b, 0x80, 0xb9, 0x1c, 0x16, 0x6b, 0x22, 0x3b, 0x72, 0x65, 0x7a, 0x2a, 0xaa, 0x0e,
	0xdc, 0xb7, 0x1d, 0xe0, 0xbd, 0x83, 0xd1, 0x33, 0x02, 0x67, 0xbc, 0x8c, 0x05, 0xf9, 0x03, 0xf6,
	0x94, 0x6e, 0xc8, 0x17, 0x7d, 0xef, 0xfb, 0xc7, 0xe1, 0x7f, 0xad, 0x28, 0x66, 0x5d, 0x47, 0x80,
	0xcd, 0x06, 0xc9, 0x37, 0x5d, 0xad, 0xef, 0xde, 0xff, 0xfe, 0x4a, 0x35, 0xdf, 0xfd, 0x05, 0x47,
	0x5e, 0x17, 0x31, 0x23, 0x2b, 0x6b, 0xf4, 0x49, 0x55, 0x32, 0xed, 0x03, 0xc0, 0xe3, 0x3c, 0xcb,
	0x24, 0x7a, 0xfd, 0x39, 0xf8, 0xf5, 0xb4, 0x87, 0xfe, 0xa1, 0xb3, 0xc1, 0x55, 0xff, 0x3d, 0xbf,
	0xc2, 0xb1, 0x0c, 0xe7, 0xae, 0x8a, 0xff, 0xbf, 0x04, 0x00, 0x00, 0xff, 0xff, 0x38, 0x55, 0x82,
	0x43, 0x3e, 0x03, 0x00, 0x00,
}
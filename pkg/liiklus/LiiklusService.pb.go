// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LiiklusService.proto

package liiklus

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

type SubscribeRequest_AutoOffsetReset int32

const (
	SubscribeRequest_EARLIEST SubscribeRequest_AutoOffsetReset = 0
	SubscribeRequest_LATEST   SubscribeRequest_AutoOffsetReset = 1
)

var SubscribeRequest_AutoOffsetReset_name = map[int32]string{
	0: "EARLIEST",
	1: "LATEST",
}

var SubscribeRequest_AutoOffsetReset_value = map[string]int32{
	"EARLIEST": 0,
	"LATEST":   1,
}

func (x SubscribeRequest_AutoOffsetReset) String() string {
	return proto.EnumName(SubscribeRequest_AutoOffsetReset_name, int32(x))
}

func (SubscribeRequest_AutoOffsetReset) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{2, 0}
}

type PublishRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PublishRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PublishReply struct {
	Partition            uint32   `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishReply) Reset()         { *m = PublishReply{} }
func (m *PublishReply) String() string { return proto.CompactTextString(m) }
func (*PublishReply) ProtoMessage()    {}
func (*PublishReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{1}
}

func (m *PublishReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishReply.Unmarshal(m, b)
}
func (m *PublishReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishReply.Marshal(b, m, deterministic)
}
func (m *PublishReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishReply.Merge(m, src)
}
func (m *PublishReply) XXX_Size() int {
	return xxx_messageInfo_PublishReply.Size(m)
}
func (m *PublishReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishReply.DiscardUnknown(m)
}

var xxx_messageInfo_PublishReply proto.InternalMessageInfo

func (m *PublishReply) GetPartition() uint32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *PublishReply) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PublishReply) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type SubscribeRequest struct {
	Topic                string                           `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Group                string                           `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	GroupVersion         uint32                           `protobuf:"varint,4,opt,name=groupVersion,proto3" json:"groupVersion,omitempty"`
	AutoOffsetReset      SubscribeRequest_AutoOffsetReset `protobuf:"varint,3,opt,name=autoOffsetReset,proto3,enum=com.github.bsideup.liiklus.SubscribeRequest_AutoOffsetReset" json:"autoOffsetReset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscribeRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *SubscribeRequest) GetGroupVersion() uint32 {
	if m != nil {
		return m.GroupVersion
	}
	return 0
}

func (m *SubscribeRequest) GetAutoOffsetReset() SubscribeRequest_AutoOffsetReset {
	if m != nil {
		return m.AutoOffsetReset
	}
	return SubscribeRequest_EARLIEST
}

type Assignment struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Partition            uint32   `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{3}
}

func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (m *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(m, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Assignment) GetPartition() uint32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

type SubscribeReply struct {
	// Types that are valid to be assigned to Reply:
	//	*SubscribeReply_Assignment
	Reply                isSubscribeReply_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SubscribeReply) Reset()         { *m = SubscribeReply{} }
func (m *SubscribeReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeReply) ProtoMessage()    {}
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{4}
}

func (m *SubscribeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReply.Unmarshal(m, b)
}
func (m *SubscribeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReply.Marshal(b, m, deterministic)
}
func (m *SubscribeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReply.Merge(m, src)
}
func (m *SubscribeReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeReply.Size(m)
}
func (m *SubscribeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReply proto.InternalMessageInfo

type isSubscribeReply_Reply interface {
	isSubscribeReply_Reply()
}

type SubscribeReply_Assignment struct {
	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3,oneof"`
}

func (*SubscribeReply_Assignment) isSubscribeReply_Reply() {}

func (m *SubscribeReply) GetReply() isSubscribeReply_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *SubscribeReply) GetAssignment() *Assignment {
	if x, ok := m.GetReply().(*SubscribeReply_Assignment); ok {
		return x.Assignment
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubscribeReply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubscribeReply_Assignment)(nil),
	}
}

type AckRequest struct {
	Assignment           *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"` // Deprecated: Do not use.
	Topic                string      `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Group                string      `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	GroupVersion         uint32      `protobuf:"varint,5,opt,name=groupVersion,proto3" json:"groupVersion,omitempty"`
	Partition            uint32      `protobuf:"varint,6,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset               uint64      `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AckRequest) Reset()         { *m = AckRequest{} }
func (m *AckRequest) String() string { return proto.CompactTextString(m) }
func (*AckRequest) ProtoMessage()    {}
func (*AckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{5}
}

func (m *AckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckRequest.Unmarshal(m, b)
}
func (m *AckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckRequest.Marshal(b, m, deterministic)
}
func (m *AckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckRequest.Merge(m, src)
}
func (m *AckRequest) XXX_Size() int {
	return xxx_messageInfo_AckRequest.Size(m)
}
func (m *AckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AckRequest proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *AckRequest) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func (m *AckRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *AckRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *AckRequest) GetGroupVersion() uint32 {
	if m != nil {
		return m.GroupVersion
	}
	return 0
}

func (m *AckRequest) GetPartition() uint32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *AckRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ReceiveRequest struct {
	Assignment           *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
	LastKnownOffset      uint64      `protobuf:"varint,2,opt,name=lastKnownOffset,proto3" json:"lastKnownOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReceiveRequest) Reset()         { *m = ReceiveRequest{} }
func (m *ReceiveRequest) String() string { return proto.CompactTextString(m) }
func (*ReceiveRequest) ProtoMessage()    {}
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{6}
}

func (m *ReceiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveRequest.Unmarshal(m, b)
}
func (m *ReceiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveRequest.Marshal(b, m, deterministic)
}
func (m *ReceiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveRequest.Merge(m, src)
}
func (m *ReceiveRequest) XXX_Size() int {
	return xxx_messageInfo_ReceiveRequest.Size(m)
}
func (m *ReceiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveRequest proto.InternalMessageInfo

func (m *ReceiveRequest) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func (m *ReceiveRequest) GetLastKnownOffset() uint64 {
	if m != nil {
		return m.LastKnownOffset
	}
	return 0
}

type ReceiveReply struct {
	// Types that are valid to be assigned to Reply:
	//	*ReceiveReply_Record_
	Reply                isReceiveReply_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReceiveReply) Reset()         { *m = ReceiveReply{} }
func (m *ReceiveReply) String() string { return proto.CompactTextString(m) }
func (*ReceiveReply) ProtoMessage()    {}
func (*ReceiveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{7}
}

func (m *ReceiveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveReply.Unmarshal(m, b)
}
func (m *ReceiveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveReply.Marshal(b, m, deterministic)
}
func (m *ReceiveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveReply.Merge(m, src)
}
func (m *ReceiveReply) XXX_Size() int {
	return xxx_messageInfo_ReceiveReply.Size(m)
}
func (m *ReceiveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveReply proto.InternalMessageInfo

type isReceiveReply_Reply interface {
	isReceiveReply_Reply()
}

type ReceiveReply_Record_ struct {
	Record *ReceiveReply_Record `protobuf:"bytes,1,opt,name=record,proto3,oneof"`
}

func (*ReceiveReply_Record_) isReceiveReply_Reply() {}

func (m *ReceiveReply) GetReply() isReceiveReply_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *ReceiveReply) GetRecord() *ReceiveReply_Record {
	if x, ok := m.GetReply().(*ReceiveReply_Record_); ok {
		return x.Record
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReceiveReply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReceiveReply_Record_)(nil),
	}
}

type ReceiveReply_Record struct {
	Offset               uint64               `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Key                  []byte               `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte               `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Replay               bool                 `protobuf:"varint,5,opt,name=replay,proto3" json:"replay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReceiveReply_Record) Reset()         { *m = ReceiveReply_Record{} }
func (m *ReceiveReply_Record) String() string { return proto.CompactTextString(m) }
func (*ReceiveReply_Record) ProtoMessage()    {}
func (*ReceiveReply_Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{7, 0}
}

func (m *ReceiveReply_Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveReply_Record.Unmarshal(m, b)
}
func (m *ReceiveReply_Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveReply_Record.Marshal(b, m, deterministic)
}
func (m *ReceiveReply_Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveReply_Record.Merge(m, src)
}
func (m *ReceiveReply_Record) XXX_Size() int {
	return xxx_messageInfo_ReceiveReply_Record.Size(m)
}
func (m *ReceiveReply_Record) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveReply_Record.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveReply_Record proto.InternalMessageInfo

func (m *ReceiveReply_Record) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReceiveReply_Record) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ReceiveReply_Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ReceiveReply_Record) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ReceiveReply_Record) GetReplay() bool {
	if m != nil {
		return m.Replay
	}
	return false
}

type GetOffsetsRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Group                string   `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	GroupVersion         uint32   `protobuf:"varint,3,opt,name=groupVersion,proto3" json:"groupVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOffsetsRequest) Reset()         { *m = GetOffsetsRequest{} }
func (m *GetOffsetsRequest) String() string { return proto.CompactTextString(m) }
func (*GetOffsetsRequest) ProtoMessage()    {}
func (*GetOffsetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{8}
}

func (m *GetOffsetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOffsetsRequest.Unmarshal(m, b)
}
func (m *GetOffsetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOffsetsRequest.Marshal(b, m, deterministic)
}
func (m *GetOffsetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOffsetsRequest.Merge(m, src)
}
func (m *GetOffsetsRequest) XXX_Size() int {
	return xxx_messageInfo_GetOffsetsRequest.Size(m)
}
func (m *GetOffsetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOffsetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOffsetsRequest proto.InternalMessageInfo

func (m *GetOffsetsRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *GetOffsetsRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *GetOffsetsRequest) GetGroupVersion() uint32 {
	if m != nil {
		return m.GroupVersion
	}
	return 0
}

type GetOffsetsReply struct {
	Offsets              map[uint32]uint64 `protobuf:"bytes,1,rep,name=offsets,proto3" json:"offsets,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetOffsetsReply) Reset()         { *m = GetOffsetsReply{} }
func (m *GetOffsetsReply) String() string { return proto.CompactTextString(m) }
func (*GetOffsetsReply) ProtoMessage()    {}
func (*GetOffsetsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2e285182f4dc03a, []int{9}
}

func (m *GetOffsetsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOffsetsReply.Unmarshal(m, b)
}
func (m *GetOffsetsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOffsetsReply.Marshal(b, m, deterministic)
}
func (m *GetOffsetsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOffsetsReply.Merge(m, src)
}
func (m *GetOffsetsReply) XXX_Size() int {
	return xxx_messageInfo_GetOffsetsReply.Size(m)
}
func (m *GetOffsetsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOffsetsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetOffsetsReply proto.InternalMessageInfo

func (m *GetOffsetsReply) GetOffsets() map[uint32]uint64 {
	if m != nil {
		return m.Offsets
	}
	return nil
}

func init() {
	proto.RegisterEnum("com.github.bsideup.liiklus.SubscribeRequest_AutoOffsetReset", SubscribeRequest_AutoOffsetReset_name, SubscribeRequest_AutoOffsetReset_value)
	proto.RegisterType((*PublishRequest)(nil), "com.github.bsideup.liiklus.PublishRequest")
	proto.RegisterType((*PublishReply)(nil), "com.github.bsideup.liiklus.PublishReply")
	proto.RegisterType((*SubscribeRequest)(nil), "com.github.bsideup.liiklus.SubscribeRequest")
	proto.RegisterType((*Assignment)(nil), "com.github.bsideup.liiklus.Assignment")
	proto.RegisterType((*SubscribeReply)(nil), "com.github.bsideup.liiklus.SubscribeReply")
	proto.RegisterType((*AckRequest)(nil), "com.github.bsideup.liiklus.AckRequest")
	proto.RegisterType((*ReceiveRequest)(nil), "com.github.bsideup.liiklus.ReceiveRequest")
	proto.RegisterType((*ReceiveReply)(nil), "com.github.bsideup.liiklus.ReceiveReply")
	proto.RegisterType((*ReceiveReply_Record)(nil), "com.github.bsideup.liiklus.ReceiveReply.Record")
	proto.RegisterType((*GetOffsetsRequest)(nil), "com.github.bsideup.liiklus.GetOffsetsRequest")
	proto.RegisterType((*GetOffsetsReply)(nil), "com.github.bsideup.liiklus.GetOffsetsReply")
	proto.RegisterMapType((map[uint32]uint64)(nil), "com.github.bsideup.liiklus.GetOffsetsReply.OffsetsEntry")
}

func init() { proto.RegisterFile("LiiklusService.proto", fileDescriptor_e2e285182f4dc03a) }

var fileDescriptor_e2e285182f4dc03a = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0x26, 0x4d, 0xd2, 0x4c, 0x43, 0x52, 0x56, 0x55, 0x15, 0x19, 0x24, 0x2a, 0x23, 0xa1,
	0xa8, 0x05, 0x17, 0x85, 0x4b, 0x55, 0x71, 0x49, 0xa5, 0xd0, 0x14, 0x2a, 0x5a, 0x6d, 0x2b, 0x0e,
	0xbd, 0x39, 0xce, 0x26, 0x35, 0x71, 0xb2, 0xc6, 0xbb, 0x2e, 0xca, 0x95, 0xc7, 0xe0, 0xc6, 0x8b,
	0xf1, 0x1a, 0x70, 0x03, 0xed, 0xae, 0x13, 0xff, 0xb4, 0x58, 0x69, 0xc5, 0x6d, 0x67, 0x3c, 0xfb,
	0xcd, 0x37, 0xdf, 0xcc, 0x8e, 0x61, 0xeb, 0xd4, 0x75, 0x27, 0x5e, 0xc8, 0x2f, 0x68, 0x70, 0xe3,
	0x3a, 0xd4, 0xf2, 0x03, 0x26, 0x18, 0x36, 0x1c, 0x36, 0xb5, 0xc6, 0xae, 0xb8, 0x0e, 0x07, 0xd6,
	0x80, 0xbb, 0x43, 0x1a, 0xfa, 0x96, 0xa7, 0x03, 0x8d, 0x67, 0x63, 0xc6, 0xc6, 0x1e, 0xdd, 0x57,
	0x91, 0x83, 0x70, 0xb4, 0x2f, 0xdc, 0x29, 0xe5, 0xc2, 0x9e, 0xfa, 0xfa, 0xb2, 0xf1, 0x24, 0x1b,
	0x40, 0xa7, 0xbe, 0x98, 0xeb, 0x8f, 0xe6, 0x47, 0x68, 0x9c, 0x87, 0x03, 0xcf, 0xe5, 0xd7, 0x84,
	0x7e, 0x09, 0x29, 0x17, 0x78, 0x0b, 0xca, 0x82, 0xf9, 0xae, 0xd3, 0x42, 0x3b, 0xa8, 0x5d, 0x23,
	0xda, 0xc0, 0x9b, 0x50, 0x9a, 0xd0, 0x79, 0xab, 0xb8, 0x83, 0xda, 0x75, 0x22, 0x8f, 0x32, 0xee,
	0xc6, 0xf6, 0x42, 0xda, 0x2a, 0x29, 0x9f, 0x36, 0xcc, 0x2b, 0xa8, 0x2f, 0xf1, 0x7c, 0x6f, 0x8e,
	0x9f, 0x42, 0xcd, 0xb7, 0x03, 0xe1, 0x0a, 0x97, 0xcd, 0x14, 0xe2, 0x23, 0x12, 0x3b, 0xf0, 0x36,
	0x54, 0xd8, 0x68, 0xc4, 0xa9, 0x50, 0xc0, 0x6b, 0x24, 0xb2, 0x62, 0x0e, 0xa5, 0x04, 0x07, 0xf3,
	0x17, 0x82, 0xcd, 0x8b, 0x70, 0xc0, 0x9d, 0xc0, 0x1d, 0xd0, 0x7c, 0xba, 0x5b, 0x50, 0x1e, 0x07,
	0x2c, 0xf4, 0x15, 0x6e, 0x8d, 0x68, 0x03, 0x9b, 0x50, 0x57, 0x87, 0x4f, 0x34, 0xe0, 0x92, 0xcf,
	0x9a, 0xe2, 0x93, 0xf2, 0xe1, 0x11, 0x34, 0xed, 0x50, 0xb0, 0x33, 0x45, 0x84, 0x50, 0xc9, 0x4d,
	0x92, 0x68, 0x74, 0xde, 0x5a, 0xff, 0x6e, 0x82, 0x95, 0xa5, 0x65, 0x75, 0xd3, 0x18, 0x24, 0x0b,
	0x6a, 0xee, 0x41, 0x33, 0x13, 0x83, 0xeb, 0xb0, 0xde, 0xeb, 0x92, 0xd3, 0x93, 0xde, 0xc5, 0xe5,
	0x66, 0x01, 0x03, 0x54, 0x4e, 0xbb, 0x97, 0xf2, 0x8c, 0xcc, 0x3e, 0x40, 0x97, 0x73, 0x77, 0x3c,
	0x9b, 0xd2, 0x99, 0x90, 0x9a, 0x72, 0xca, 0x25, 0xdb, 0x93, 0x61, 0x54, 0x76, 0xec, 0x48, 0x2b,
	0x5e, 0xcc, 0x28, 0x6e, 0x3a, 0xd0, 0x48, 0x70, 0x95, 0x1d, 0xea, 0x03, 0xd8, 0x4b, 0x6c, 0x05,
	0xb7, 0xd1, 0x79, 0x91, 0x57, 0x6b, 0xcc, 0xa4, 0x5f, 0x20, 0x89, 0xbb, 0x47, 0x55, 0x28, 0x07,
	0x12, 0xd2, 0xfc, 0x89, 0x00, 0xba, 0xce, 0x64, 0xd1, 0xa2, 0xf7, 0x0f, 0xcf, 0x70, 0x54, 0x6c,
	0xa1, 0x64, 0x8e, 0xbb, 0x27, 0x23, 0x6e, 0xf7, 0x5a, 0x5e, 0xbb, 0xcb, 0x77, 0xb4, 0x3b, 0xa5,
	0x56, 0x65, 0xc5, 0xf9, 0x34, 0xbf, 0x21, 0x68, 0x10, 0xea, 0x50, 0xf7, 0x66, 0x39, 0x87, 0xef,
	0x1e, 0x5e, 0x64, 0xaa, 0xc0, 0x36, 0x34, 0x3d, 0x9b, 0x8b, 0x0f, 0x33, 0xf6, 0x75, 0x76, 0x96,
	0xcc, 0x9d, 0x75, 0x9b, 0x7f, 0x10, 0xd4, 0x97, 0x24, 0x64, 0x27, 0x4f, 0xa0, 0x12, 0x50, 0x87,
	0x05, 0xc3, 0x28, 0xfd, 0x7e, 0x5e, 0xfa, 0xe4, 0x4d, 0x69, 0xb0, 0x60, 0xd8, 0x2f, 0x90, 0x08,
	0xc0, 0xf8, 0x8e, 0xa0, 0xa2, 0x9d, 0x09, 0x0d, 0x50, 0xea, 0x8d, 0xae, 0xb8, 0x11, 0xf0, 0x01,
	0xd4, 0x96, 0x1b, 0x49, 0xf5, 0x67, 0xa3, 0x63, 0x58, 0x7a, 0x25, 0x59, 0x8b, 0x95, 0x64, 0x5d,
	0x2e, 0x22, 0x48, 0x1c, 0x2c, 0x33, 0xcb, 0x79, 0xb2, 0xe7, 0xaa, 0x73, 0xeb, 0x24, 0xb2, 0xe2,
	0x39, 0x73, 0xe0, 0xf1, 0x31, 0x15, 0x5a, 0x0e, 0xfe, 0x3f, 0x16, 0x42, 0xe9, 0xf6, 0x84, 0x98,
	0x3f, 0x10, 0x34, 0x93, 0x59, 0xa4, 0xd2, 0x04, 0xaa, 0x5a, 0x05, 0xde, 0x42, 0x3b, 0xa5, 0xf6,
	0x46, 0xe7, 0x20, 0x4f, 0xea, 0xcc, 0x6d, 0x2b, 0x32, 0x7a, 0x33, 0x11, 0xcc, 0xc9, 0x02, 0xc8,
	0x38, 0x84, 0x7a, 0xf2, 0xc3, 0x42, 0x5f, 0xbd, 0x33, 0xd3, 0xfa, 0xea, 0x81, 0xd0, 0xc6, 0x61,
	0xf1, 0x00, 0x75, 0x7e, 0x97, 0xa0, 0x91, 0xfe, 0x71, 0x60, 0x1b, 0xaa, 0xd1, 0x22, 0xc6, 0xbb,
	0x79, 0xe4, 0xd2, 0xdb, 0xdf, 0x68, 0xaf, 0x14, 0x2b, 0xc5, 0x2f, 0x60, 0x17, 0x6a, 0xcb, 0x5d,
	0x82, 0x5f, 0xde, 0x67, 0x3d, 0x1a, 0xbb, 0x2b, 0x46, 0xab, 0x44, 0xaf, 0x11, 0x76, 0xa0, 0x1a,
	0x0d, 0x6c, 0x7e, 0x35, 0xe9, 0x47, 0x99, 0x5f, 0x4d, 0xf2, 0x05, 0xa8, 0x24, 0xc7, 0x50, 0xea,
	0x3a, 0x13, 0x9c, 0xff, 0x6a, 0x97, 0x6b, 0xcd, 0xd8, 0xbe, 0x35, 0xc5, 0x3d, 0xf9, 0x63, 0x35,
	0x0b, 0xf8, 0x33, 0x40, 0xdc, 0x73, 0xfc, 0x6a, 0xd5, 0xd9, 0xd0, 0xb0, 0x7b, 0xf7, 0x18, 0x25,
	0xb3, 0x70, 0xd4, 0x81, 0xe7, 0x39, 0xf1, 0x8a, 0x9a, 0xc3, 0xbc, 0x3e, 0x3a, 0x47, 0x57, 0xd5,
	0xc8, 0x3b, 0xa8, 0x28, 0xf7, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xbb, 0x17, 0x0d,
	0x6d, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LiiklusServiceClient is the client API for LiiklusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LiiklusServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LiiklusService_SubscribeClient, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (LiiklusService_ReceiveClient, error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error)
}

type liiklusServiceClient struct {
	cc *grpc.ClientConn
}

func NewLiiklusServiceClient(cc *grpc.ClientConn) LiiklusServiceClient {
	return &liiklusServiceClient{cc}
}

func (c *liiklusServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error) {
	out := new(PublishReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LiiklusService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LiiklusService_serviceDesc.Streams[0], "/com.github.bsideup.liiklus.LiiklusService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &liiklusServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LiiklusService_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type liiklusServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *liiklusServiceSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *liiklusServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (LiiklusService_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LiiklusService_serviceDesc.Streams[1], "/com.github.bsideup.liiklus.LiiklusService/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &liiklusServiceReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LiiklusService_ReceiveClient interface {
	Recv() (*ReceiveReply, error)
	grpc.ClientStream
}

type liiklusServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *liiklusServiceReceiveClient) Recv() (*ReceiveReply, error) {
	m := new(ReceiveReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *liiklusServiceClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error) {
	out := new(GetOffsetsReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/GetOffsets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiiklusServiceServer is the server API for LiiklusService service.
type LiiklusServiceServer interface {
	Publish(context.Context, *PublishRequest) (*PublishReply, error)
	Subscribe(*SubscribeRequest, LiiklusService_SubscribeServer) error
	Receive(*ReceiveRequest, LiiklusService_ReceiveServer) error
	Ack(context.Context, *AckRequest) (*empty.Empty, error)
	GetOffsets(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error)
}

func RegisterLiiklusServiceServer(s *grpc.Server, srv LiiklusServiceServer) {
	s.RegisterService(&_LiiklusService_serviceDesc, srv)
}

func _LiiklusService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LiiklusServiceServer).Subscribe(m, &liiklusServiceSubscribeServer{stream})
}

type LiiklusService_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type liiklusServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *liiklusServiceSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

func _LiiklusService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LiiklusServiceServer).Receive(m, &liiklusServiceReceiveServer{stream})
}

type LiiklusService_ReceiveServer interface {
	Send(*ReceiveReply) error
	grpc.ServerStream
}

type liiklusServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *liiklusServiceReceiveServer) Send(m *ReceiveReply) error {
	return x.ServerStream.SendMsg(m)
}

func _LiiklusService_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_GetOffsets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffsetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).GetOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/GetOffsets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).GetOffsets(ctx, req.(*GetOffsetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LiiklusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.bsideup.liiklus.LiiklusService",
	HandlerType: (*LiiklusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _LiiklusService_Publish_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _LiiklusService_Ack_Handler,
		},
		{
			MethodName: "GetOffsets",
			Handler:    _LiiklusService_GetOffsets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _LiiklusService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Receive",
			Handler:       _LiiklusService_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "LiiklusService.proto",
}

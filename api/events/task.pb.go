// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	codec "github.com/kcarretto/paragon/api/codec"
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

type TaskQueued struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Filter               *codec.AgentMetadata `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskQueued) Reset()         { *m = TaskQueued{} }
func (m *TaskQueued) String() string { return proto.CompactTextString(m) }
func (*TaskQueued) ProtoMessage()    {}
func (*TaskQueued) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *TaskQueued) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskQueued.Unmarshal(m, b)
}
func (m *TaskQueued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskQueued.Marshal(b, m, deterministic)
}
func (m *TaskQueued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskQueued.Merge(m, src)
}
func (m *TaskQueued) XXX_Size() int {
	return xxx_messageInfo_TaskQueued.Size(m)
}
func (m *TaskQueued) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskQueued.DiscardUnknown(m)
}

var xxx_messageInfo_TaskQueued proto.InternalMessageInfo

func (m *TaskQueued) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskQueued) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskQueued) GetFilter() *codec.AgentMetadata {
	if m != nil {
		return m.Filter
	}
	return nil
}

type TaskClaimed struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Agent                *codec.AgentMetadata `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskClaimed) Reset()         { *m = TaskClaimed{} }
func (m *TaskClaimed) String() string { return proto.CompactTextString(m) }
func (*TaskClaimed) ProtoMessage()    {}
func (*TaskClaimed) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *TaskClaimed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskClaimed.Unmarshal(m, b)
}
func (m *TaskClaimed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskClaimed.Marshal(b, m, deterministic)
}
func (m *TaskClaimed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskClaimed.Merge(m, src)
}
func (m *TaskClaimed) XXX_Size() int {
	return xxx_messageInfo_TaskClaimed.Size(m)
}
func (m *TaskClaimed) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskClaimed.DiscardUnknown(m)
}

var xxx_messageInfo_TaskClaimed proto.InternalMessageInfo

func (m *TaskClaimed) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskClaimed) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

type TaskExecuted struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Output               []string             `protobuf:"bytes,2,rep,name=output,proto3" json:"output,omitempty"`
	Error                string               `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ExecStartTime        int64                `protobuf:"varint,4,opt,name=execStartTime,proto3" json:"execStartTime,omitempty"`
	ExecStopTime         int64                `protobuf:"varint,5,opt,name=execStopTime,proto3" json:"execStopTime,omitempty"`
	RecvTime             int64                `protobuf:"varint,6,opt,name=recvTime,proto3" json:"recvTime,omitempty"`
	Agent                *codec.AgentMetadata `protobuf:"bytes,7,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskExecuted) Reset()         { *m = TaskExecuted{} }
func (m *TaskExecuted) String() string { return proto.CompactTextString(m) }
func (*TaskExecuted) ProtoMessage()    {}
func (*TaskExecuted) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TaskExecuted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecuted.Unmarshal(m, b)
}
func (m *TaskExecuted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecuted.Marshal(b, m, deterministic)
}
func (m *TaskExecuted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecuted.Merge(m, src)
}
func (m *TaskExecuted) XXX_Size() int {
	return xxx_messageInfo_TaskExecuted.Size(m)
}
func (m *TaskExecuted) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecuted.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecuted proto.InternalMessageInfo

func (m *TaskExecuted) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskExecuted) GetOutput() []string {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *TaskExecuted) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TaskExecuted) GetExecStartTime() int64 {
	if m != nil {
		return m.ExecStartTime
	}
	return 0
}

func (m *TaskExecuted) GetExecStopTime() int64 {
	if m != nil {
		return m.ExecStopTime
	}
	return 0
}

func (m *TaskExecuted) GetRecvTime() int64 {
	if m != nil {
		return m.RecvTime
	}
	return 0
}

func (m *TaskExecuted) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskQueued)(nil), "events.TaskQueued")
	proto.RegisterType((*TaskClaimed)(nil), "events.TaskClaimed")
	proto.RegisterType((*TaskExecuted)(nil), "events.TaskExecuted")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x49, 0xfb, 0x6f, 0xfa, 0xef, 0xb4, 0x7a, 0x58, 0xaa, 0x84, 0x9e, 0x4a, 0x10, 0x8c,
	0x22, 0x09, 0xe8, 0x27, 0x50, 0xf1, 0xe0, 0xc1, 0x83, 0xb1, 0x27, 0x6f, 0xd3, 0xdd, 0xb1, 0x2e,
	0x6d, 0xb3, 0x61, 0x33, 0x5b, 0xfa, 0x69, 0xfd, 0x2c, 0x92, 0xdd, 0xaa, 0x14, 0x2a, 0x1e, 0xdf,
	0x9b, 0xc7, 0x6f, 0x66, 0x78, 0x00, 0x8c, 0xcd, 0x32, 0xaf, 0xad, 0x61, 0x23, 0x62, 0xda, 0x50,
	0xc5, 0xcd, 0xe4, 0x04, 0x6b, 0x5d, 0x48, 0xa3, 0x48, 0x16, 0xb8, 0xa0, 0x8a, 0xc3, 0x38, 0x55,
	0x00, 0x33, 0x6c, 0x96, 0xcf, 0x8e, 0x1c, 0x29, 0x71, 0x0c, 0x1d, 0xad, 0x92, 0x68, 0x1a, 0x65,
	0x83, 0xb2, 0xa3, 0x95, 0x48, 0xa0, 0x2f, 0x4d, 0xc5, 0x54, 0x71, 0xd2, 0xf1, 0xe6, 0x97, 0x14,
	0x57, 0x10, 0xbf, 0xe9, 0x15, 0x93, 0x4d, 0xba, 0xd3, 0x28, 0x1b, 0x5e, 0x8f, 0x73, 0xcf, 0xce,
	0x6f, 0x5b, 0xf6, 0x13, 0x31, 0x2a, 0x64, 0x2c, 0x77, 0x99, 0xf4, 0x11, 0x86, 0xed, 0x96, 0xfb,
	0x15, 0xea, 0xf5, 0x81, 0x35, 0x97, 0xd0, 0xf3, 0x37, 0xf9, 0x25, 0xbf, 0xb1, 0x42, 0x24, 0xfd,
	0x88, 0x60, 0xd4, 0xb2, 0x1e, 0xb6, 0x24, 0x1d, 0x1f, 0x80, 0x9d, 0x42, 0x6c, 0x1c, 0xd7, 0xae,
	0xa5, 0x75, 0xb3, 0x41, 0xb9, 0x53, 0x62, 0x0c, 0x3d, 0xb2, 0xd6, 0x84, 0x83, 0x07, 0x65, 0x10,
	0xe2, 0x0c, 0x8e, 0x68, 0x4b, 0xf2, 0x85, 0xd1, 0xf2, 0x4c, 0xaf, 0x29, 0xf9, 0x37, 0x8d, 0xb2,
	0x6e, 0xb9, 0x6f, 0x8a, 0x14, 0x46, 0xc1, 0x30, 0xb5, 0x0f, 0xf5, 0x7c, 0x68, 0xcf, 0x13, 0x13,
	0xf8, 0x6f, 0x49, 0x6e, 0xfc, 0x3c, 0xf6, 0xf3, 0x6f, 0xfd, 0xf3, 0x60, 0xff, 0xcf, 0x07, 0xef,
	0x2e, 0x5e, 0xcf, 0x17, 0x9a, 0xdf, 0xdd, 0x3c, 0x97, 0x66, 0x5d, 0x2c, 0x25, 0x5a, 0x4b, 0xcc,
	0xa6, 0xa8, 0xd1, 0xe2, 0xc2, 0x54, 0x45, 0xdb, 0x63, 0xe8, 0x74, 0x1e, 0xfb, 0x0e, 0x6f, 0x3e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x18, 0x55, 0x4d, 0xf0, 0x01, 0x00, 0x00,
}
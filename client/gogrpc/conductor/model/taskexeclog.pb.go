// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/taskexeclog.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

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

type TaskExecLog struct {
	Log                  string   `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	CreatedTime          int64    `protobuf:"varint,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecLog) Reset()         { *m = TaskExecLog{} }
func (m *TaskExecLog) String() string { return proto.CompactTextString(m) }
func (*TaskExecLog) ProtoMessage()    {}
func (*TaskExecLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskexeclog_31ce5708c84ca255, []int{0}
}
func (m *TaskExecLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecLog.Unmarshal(m, b)
}
func (m *TaskExecLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecLog.Marshal(b, m, deterministic)
}
func (dst *TaskExecLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecLog.Merge(dst, src)
}
func (m *TaskExecLog) XXX_Size() int {
	return xxx_messageInfo_TaskExecLog.Size(m)
}
func (m *TaskExecLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecLog.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecLog proto.InternalMessageInfo

func (m *TaskExecLog) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *TaskExecLog) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskExecLog) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskExecLog)(nil), "conductor.proto.TaskExecLog")
}

func init() {
	proto.RegisterFile("model/taskexeclog.proto", fileDescriptor_taskexeclog_31ce5708c84ca255)
}

var fileDescriptor_taskexeclog_31ce5708c84ca255 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x81, 0x13, 0xf7, 0x14, 0x65, 0x9b, 0x0b, 0xd8, 0x9c, 0x56, 0x57, 0xed, 0x16,
	0x76, 0x96, 0x07, 0x16, 0x82, 0x85, 0x84, 0x54, 0x5a, 0x84, 0x64, 0x76, 0xdc, 0x2c, 0xd9, 0xcd,
	0x84, 0xcd, 0x04, 0xf2, 0xf3, 0x25, 0x31, 0x48, 0xb8, 0x6e, 0xe6, 0x83, 0xf7, 0x3e, 0x9e, 0x38,
	0x04, 0x32, 0xe8, 0x35, 0x57, 0x43, 0x8b, 0x13, 0x82, 0x27, 0xab, 0xfa, 0x48, 0x4c, 0xf2, 0x1e,
	0xa8, 0x33, 0x23, 0x30, 0xc5, 0x3f, 0xf0, 0xfc, 0x2d, 0xf6, 0x45, 0x35, 0xb4, 0x6f, 0x13, 0xc2,
	0x07, 0x59, 0xf9, 0x20, 0x52, 0x4f, 0x36, 0x4b, 0x8e, 0xc9, 0xe9, 0x26, 0x9f, 0x4f, 0x79, 0x10,
	0xd7, 0x73, 0x4d, 0xe9, 0x4c, 0x76, 0xb5, 0xd0, 0xdd, 0xfc, 0xbe, 0x1b, 0xf9, 0x24, 0x6e, 0x21,
	0x62, 0xc5, 0x68, 0x4a, 0x76, 0x01, 0xb3, 0xf4, 0x98, 0x9c, 0xd2, 0x7c, 0xbf, 0xb2, 0xc2, 0x05,
	0x3c, 0x37, 0xe2, 0x11, 0x28, 0xa8, 0x0e, 0xf9, 0xc7, 0xbb, 0x49, 0x5d, 0xb8, 0xcf, 0x77, 0x1b,
	0xf3, 0x67, 0xfd, 0xf5, 0x6a, 0x1d, 0x37, 0x63, 0xad, 0x80, 0x82, 0x5e, 0x23, 0xfa, 0x3f, 0xa2,
	0xc1, 0x3b, 0xec, 0x58, 0x5b, 0xb2, 0xb1, 0x87, 0x0d, 0x5f, 0x96, 0xd6, 0xbb, 0xa5, 0xf1, 0xe5,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x78, 0x61, 0x87, 0x8e, 0xf9, 0x00, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type Job struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	JobName              string   `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Job) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func init() {
	proto.RegisterType((*Job)(nil), "job.Job")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x14, 0x42, 0x30, 0x08, 0xa1, 0x88, 0xa1, 0xe9, 0x00, 0x88, 0x29, 0x48, 0xa4,
	0x16, 0x45, 0x2c, 0x65, 0xeb, 0x06, 0x03, 0x43, 0x47, 0x96, 0xca, 0x97, 0xb8, 0xae, 0xad, 0xda,
	0x7f, 0xe4, 0xa6, 0x88, 0x57, 0x4b, 0x26, 0xde, 0x85, 0x17, 0x41, 0x75, 0x96, 0x2e, 0xa7, 0xfb,
	0xa4, 0xef, 0x3f, 0xfd, 0x3a, 0x76, 0x61, 0x51, 0xcb, 0xed, 0xb4, 0xf1, 0x68, 0x91, 0x8d, 0x0c,
	0x68, 0x32, 0x57, 0xba, 0xdd, 0xec, 0x69, 0x5a, 0xc1, 0x72, 0xed, 0xd6, 0xa0, 0x2d, 0x7e, 0xd0,
	0x48, 0xc7, 0x83, 0x53, 0x95, 0x4a, 0xba, 0x52, 0xc1, 0x5b, 0x8e, 0xa6, 0xd5, 0x70, 0x3b, 0x7e,
	0x80, 0xe1, 0xc0, 0xa4, 0x3c, 0xca, 0x2a, 0x28, 0x0c, 0x19, 0xda, 0xaf, 0x03, 0x05, 0x08, 0xdb,
	0xa0, 0x3f, 0xac, 0xd8, 0xe8, 0x03, 0x94, 0xdd, 0xb1, 0xc4, 0x80, 0x56, 0xba, 0x1e, 0x47, 0xf7,
	0x51, 0x71, 0xbe, 0x48, 0xfb, 0x2e, 0x3f, 0x61, 0x71, 0x11, 0x2d, 0x4f, 0x0d, 0xe8, 0xbd, 0xce,
	0x38, 0x4b, 0x0f, 0x82, 0x13, 0x56, 0x8e, 0xe3, 0xa0, 0xdc, 0xf4, 0x5d, 0x7e, 0xcd, 0xae, 0xb2,
	0xcb, 0x6f, 0xe1, 0xab, 0x8d, 0xf0, 0xc5, 0xeb, 0xf3, 0xec, 0x71, 0x79, 0x66, 0x40, 0x9f, 0xc2,
	0xca, 0x79, 0xd2, 0x77, 0x79, 0x9c, 0x46, 0x8b, 0xd9, 0xef, 0xdf, 0x6d, 0xf4, 0xf5, 0x74, 0xd4,
	0xca, 0x69, 0x2f, 0x8c, 0x92, 0xf0, 0x4a, 0x0b, 0xc7, 0x0d, 0x88, 0xef, 0x7c, 0xc5, 0xc3, 0x13,
	0xde, 0xc2, 0xa4, 0x24, 0x74, 0x7b, 0xf9, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x78, 0xc8, 0x6c, 0xfd,
	0x1a, 0x01, 0x00, 0x00,
}
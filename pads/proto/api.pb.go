// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StatusType int32

const (
	StatusType_OK    StatusType = 0
	StatusType_ERROR StatusType = 1
)

var StatusType_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}

var StatusType_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x StatusType) String() string {
	return proto.EnumName(StatusType_name, int32(x))
}

func (StatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type RegisterWorkerRequest struct {
	Worker               *Worker  `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterWorkerRequest) Reset()         { *m = RegisterWorkerRequest{} }
func (m *RegisterWorkerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterWorkerRequest) ProtoMessage()    {}
func (*RegisterWorkerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *RegisterWorkerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterWorkerRequest.Unmarshal(m, b)
}
func (m *RegisterWorkerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterWorkerRequest.Marshal(b, m, deterministic)
}
func (m *RegisterWorkerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterWorkerRequest.Merge(m, src)
}
func (m *RegisterWorkerRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterWorkerRequest.Size(m)
}
func (m *RegisterWorkerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterWorkerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterWorkerRequest proto.InternalMessageInfo

func (m *RegisterWorkerRequest) GetWorker() *Worker {
	if m != nil {
		return m.Worker
	}
	return nil
}

type RegisterWorkerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterWorkerResponse) Reset()         { *m = RegisterWorkerResponse{} }
func (m *RegisterWorkerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterWorkerResponse) ProtoMessage()    {}
func (*RegisterWorkerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *RegisterWorkerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterWorkerResponse.Unmarshal(m, b)
}
func (m *RegisterWorkerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterWorkerResponse.Marshal(b, m, deterministic)
}
func (m *RegisterWorkerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterWorkerResponse.Merge(m, src)
}
func (m *RegisterWorkerResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterWorkerResponse.Size(m)
}
func (m *RegisterWorkerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterWorkerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterWorkerResponse proto.InternalMessageInfo

type UpdateWorkerRequest struct {
	Worker               *Worker  `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWorkerRequest) Reset()         { *m = UpdateWorkerRequest{} }
func (m *UpdateWorkerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateWorkerRequest) ProtoMessage()    {}
func (*UpdateWorkerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *UpdateWorkerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWorkerRequest.Unmarshal(m, b)
}
func (m *UpdateWorkerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWorkerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateWorkerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWorkerRequest.Merge(m, src)
}
func (m *UpdateWorkerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateWorkerRequest.Size(m)
}
func (m *UpdateWorkerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWorkerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWorkerRequest proto.InternalMessageInfo

func (m *UpdateWorkerRequest) GetWorker() *Worker {
	if m != nil {
		return m.Worker
	}
	return nil
}

type UpdateWorkerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWorkerResponse) Reset()         { *m = UpdateWorkerResponse{} }
func (m *UpdateWorkerResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateWorkerResponse) ProtoMessage()    {}
func (*UpdateWorkerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *UpdateWorkerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWorkerResponse.Unmarshal(m, b)
}
func (m *UpdateWorkerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWorkerResponse.Marshal(b, m, deterministic)
}
func (m *UpdateWorkerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWorkerResponse.Merge(m, src)
}
func (m *UpdateWorkerResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateWorkerResponse.Size(m)
}
func (m *UpdateWorkerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWorkerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWorkerResponse proto.InternalMessageInfo

type RunWorkerRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunWorkerRequest) Reset()         { *m = RunWorkerRequest{} }
func (m *RunWorkerRequest) String() string { return proto.CompactTextString(m) }
func (*RunWorkerRequest) ProtoMessage()    {}
func (*RunWorkerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *RunWorkerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunWorkerRequest.Unmarshal(m, b)
}
func (m *RunWorkerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunWorkerRequest.Marshal(b, m, deterministic)
}
func (m *RunWorkerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunWorkerRequest.Merge(m, src)
}
func (m *RunWorkerRequest) XXX_Size() int {
	return xxx_messageInfo_RunWorkerRequest.Size(m)
}
func (m *RunWorkerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunWorkerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunWorkerRequest proto.InternalMessageInfo

type RunWorkerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunWorkerResponse) Reset()         { *m = RunWorkerResponse{} }
func (m *RunWorkerResponse) String() string { return proto.CompactTextString(m) }
func (*RunWorkerResponse) ProtoMessage()    {}
func (*RunWorkerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *RunWorkerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunWorkerResponse.Unmarshal(m, b)
}
func (m *RunWorkerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunWorkerResponse.Marshal(b, m, deterministic)
}
func (m *RunWorkerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunWorkerResponse.Merge(m, src)
}
func (m *RunWorkerResponse) XXX_Size() int {
	return xxx_messageInfo_RunWorkerResponse.Size(m)
}
func (m *RunWorkerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunWorkerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunWorkerResponse proto.InternalMessageInfo

type GetAgentsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentsRequest) Reset()         { *m = GetAgentsRequest{} }
func (m *GetAgentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgentsRequest) ProtoMessage()    {}
func (*GetAgentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *GetAgentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentsRequest.Unmarshal(m, b)
}
func (m *GetAgentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentsRequest.Marshal(b, m, deterministic)
}
func (m *GetAgentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentsRequest.Merge(m, src)
}
func (m *GetAgentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgentsRequest.Size(m)
}
func (m *GetAgentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentsRequest proto.InternalMessageInfo

type GetAgentsResponse struct {
	Agents               []*Agent `protobuf:"bytes,1,rep,name=agents,proto3" json:"agents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentsResponse) Reset()         { *m = GetAgentsResponse{} }
func (m *GetAgentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAgentsResponse) ProtoMessage()    {}
func (*GetAgentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *GetAgentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentsResponse.Unmarshal(m, b)
}
func (m *GetAgentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentsResponse.Marshal(b, m, deterministic)
}
func (m *GetAgentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentsResponse.Merge(m, src)
}
func (m *GetAgentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAgentsResponse.Size(m)
}
func (m *GetAgentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentsResponse proto.InternalMessageInfo

func (m *GetAgentsResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

type Master struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerAddress        string   `protobuf:"bytes,2,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Master) Reset()         { *m = Master{} }
func (m *Master) String() string { return proto.CompactTextString(m) }
func (*Master) ProtoMessage()    {}
func (*Master) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *Master) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Master.Unmarshal(m, b)
}
func (m *Master) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Master.Marshal(b, m, deterministic)
}
func (m *Master) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Master.Merge(m, src)
}
func (m *Master) XXX_Size() int {
	return xxx_messageInfo_Master.Size(m)
}
func (m *Master) XXX_DiscardUnknown() {
	xxx_messageInfo_Master.DiscardUnknown(m)
}

var xxx_messageInfo_Master proto.InternalMessageInfo

func (m *Master) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Master) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

type Worker struct {
	Id                   uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerAddress        string    `protobuf:"bytes,2,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	NeighborWorkers      []*Worker `protobuf:"bytes,3,rep,name=neighbor_workers,json=neighborWorkers,proto3" json:"neighbor_workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *Worker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker.Unmarshal(m, b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
}
func (m *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(m, src)
}
func (m *Worker) XXX_Size() int {
	return xxx_messageInfo_Worker.Size(m)
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Worker) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *Worker) GetNeighborWorkers() []*Worker {
	if m != nil {
		return m.NeighborWorkers
	}
	return nil
}

type Agent struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Agent) Reset()         { *m = Agent{} }
func (m *Agent) String() string { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()    {}
func (*Agent) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *Agent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Agent.Unmarshal(m, b)
}
func (m *Agent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Agent.Marshal(b, m, deterministic)
}
func (m *Agent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agent.Merge(m, src)
}
func (m *Agent) XXX_Size() int {
	return xxx_messageInfo_Agent.Size(m)
}
func (m *Agent) XXX_DiscardUnknown() {
	xxx_messageInfo_Agent.DiscardUnknown(m)
}

var xxx_messageInfo_Agent proto.InternalMessageInfo

func (m *Agent) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Agent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type BasicResponse struct {
	Status               StatusType `protobuf:"varint,1,opt,name=status,proto3,enum=api.StatusType" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BasicResponse) Reset()         { *m = BasicResponse{} }
func (m *BasicResponse) String() string { return proto.CompactTextString(m) }
func (*BasicResponse) ProtoMessage()    {}
func (*BasicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *BasicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasicResponse.Unmarshal(m, b)
}
func (m *BasicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasicResponse.Marshal(b, m, deterministic)
}
func (m *BasicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicResponse.Merge(m, src)
}
func (m *BasicResponse) XXX_Size() int {
	return xxx_messageInfo_BasicResponse.Size(m)
}
func (m *BasicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BasicResponse proto.InternalMessageInfo

func (m *BasicResponse) GetStatus() StatusType {
	if m != nil {
		return m.Status
	}
	return StatusType_OK
}

func init() {
	proto.RegisterEnum("api.StatusType", StatusType_name, StatusType_value)
	proto.RegisterType((*RegisterWorkerRequest)(nil), "api.RegisterWorkerRequest")
	proto.RegisterType((*RegisterWorkerResponse)(nil), "api.RegisterWorkerResponse")
	proto.RegisterType((*UpdateWorkerRequest)(nil), "api.UpdateWorkerRequest")
	proto.RegisterType((*UpdateWorkerResponse)(nil), "api.UpdateWorkerResponse")
	proto.RegisterType((*RunWorkerRequest)(nil), "api.RunWorkerRequest")
	proto.RegisterType((*RunWorkerResponse)(nil), "api.RunWorkerResponse")
	proto.RegisterType((*GetAgentsRequest)(nil), "api.GetAgentsRequest")
	proto.RegisterType((*GetAgentsResponse)(nil), "api.GetAgentsResponse")
	proto.RegisterType((*Master)(nil), "api.Master")
	proto.RegisterType((*Worker)(nil), "api.Worker")
	proto.RegisterType((*Agent)(nil), "api.Agent")
	proto.RegisterType((*BasicResponse)(nil), "api.BasicResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x8d, 0x9d, 0xc6, 0xe0, 0x9b, 0xd9, 0x75, 0xd5, 0x35, 0x78, 0xde, 0xc3, 0x82, 0xc6, 0x58,
	0xd8, 0xa0, 0x0f, 0x1e, 0x6c, 0x63, 0x04, 0x46, 0x0b, 0x65, 0x0f, 0x65, 0x14, 0xd4, 0x8d, 0xc1,
	0x5e, 0x8a, 0x5a, 0x5f, 0x32, 0x31, 0xb0, 0x3d, 0x49, 0x69, 0xd9, 0x37, 0xec, 0x33, 0xf6, 0xa3,
	0x25, 0x92, 0xec, 0xc6, 0xae, 0x9f, 0xf2, 0x96, 0x9c, 0x7b, 0xce, 0xd1, 0xd1, 0xf5, 0x11, 0x84,
	0xbc, 0x16, 0xc7, 0xb5, 0xac, 0x74, 0x45, 0xc6, 0xbc, 0x16, 0x74, 0x09, 0x47, 0x0c, 0x57, 0x42,
	0x69, 0x94, 0x3f, 0x2a, 0xf9, 0x1b, 0x25, 0xc3, 0x3f, 0x6b, 0x54, 0x9a, 0xbc, 0x84, 0xe0, 0xce,
	0x00, 0xa9, 0x37, 0xf7, 0x16, 0xd3, 0x7c, 0x7a, 0xbc, 0x51, 0x3a, 0x8e, 0x1b, 0xd1, 0x14, 0x66,
	0x7d, 0xb5, 0xaa, 0xab, 0x52, 0x21, 0xfd, 0x04, 0x87, 0xdf, 0xeb, 0x82, 0x6b, 0xdc, 0xc1, 0x75,
	0x06, 0x4f, 0xbb, 0x5a, 0xe7, 0x49, 0x20, 0x61, 0xeb, 0xb2, 0x63, 0x48, 0x0f, 0xe1, 0x60, 0x0b,
	0x7b, 0x20, 0x7e, 0x41, 0x7d, 0xb2, 0xc2, 0x52, 0xab, 0x86, 0xf8, 0x01, 0x0e, 0xb6, 0x30, 0x4b,
	0x24, 0x14, 0x02, 0x6e, 0x90, 0xd4, 0x9b, 0x8f, 0x17, 0xd3, 0x1c, 0x4c, 0x1c, 0x43, 0x62, 0x6e,
	0x42, 0x3f, 0x43, 0xf0, 0x95, 0x6f, 0x6e, 0x48, 0x62, 0xf0, 0x45, 0x61, 0x82, 0xef, 0x31, 0x5f,
	0x14, 0xe4, 0x15, 0xc4, 0x0a, 0xe5, 0x2d, 0xca, 0x2b, 0x5e, 0x14, 0x12, 0x95, 0x4a, 0xfd, 0xb9,
	0xb7, 0x08, 0x59, 0x64, 0xd1, 0x13, 0x0b, 0xd2, 0x3b, 0x08, 0x6c, 0xbe, 0x1d, 0x0d, 0xc8, 0x7b,
	0x48, 0x4a, 0x14, 0xab, 0x5f, 0xd7, 0x95, 0xbc, 0xb2, 0x2b, 0x52, 0xe9, 0xd8, 0xe4, 0xed, 0xac,
	0x6f, 0xbf, 0x21, 0xd9, 0xff, 0x8a, 0xbe, 0x85, 0x89, 0xb9, 0xca, 0xa3, 0x73, 0x09, 0xec, 0x15,
	0x5c, 0x73, 0x77, 0x9a, 0xf9, 0x4d, 0x3f, 0x42, 0x74, 0xca, 0x95, 0xb8, 0x69, 0x77, 0xf3, 0x1a,
	0x02, 0xa5, 0xb9, 0x5e, 0x2b, 0x23, 0x8c, 0xf3, 0x7d, 0x73, 0xd6, 0xa5, 0x81, 0xbe, 0xfd, 0xad,
	0x91, 0xb9, 0xf1, 0x9b, 0x17, 0x00, 0x0f, 0x28, 0x09, 0xc0, 0xbf, 0x38, 0x4f, 0x46, 0x24, 0x84,
	0xc9, 0x19, 0x63, 0x17, 0x2c, 0xf1, 0xf2, 0xff, 0x1e, 0x44, 0x76, 0x85, 0x97, 0x28, 0x6f, 0xc5,
	0x0d, 0x92, 0x73, 0x88, 0xbb, 0xbd, 0x21, 0x99, 0x71, 0x1f, 0xac, 0x62, 0xf6, 0x7c, 0x70, 0xe6,
	0xbe, 0xf5, 0x88, 0x9c, 0xc1, 0x93, 0xed, 0xba, 0x90, 0xd4, 0xd0, 0x07, 0xda, 0x97, 0x3d, 0x1b,
	0x98, 0x34, 0x36, 0xf9, 0x3f, 0x0f, 0x22, 0x0b, 0x36, 0x29, 0x97, 0x10, 0xb6, 0xdd, 0x22, 0x47,
	0x36, 0x44, 0xaf, 0x7f, 0xd9, 0xac, 0x0f, 0xb7, 0xb1, 0x96, 0x10, 0xb6, 0x85, 0x73, 0xea, 0x7e,
	0x29, 0x9d, 0xfa, 0x51, 0x2f, 0xe9, 0xe8, 0x74, 0xf2, 0x73, 0xf3, 0x3c, 0xaf, 0x03, 0xf3, 0x54,
	0xdf, 0xdd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x99, 0x5d, 0x84, 0x70, 0xb7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterServiceClient interface {
	RegisterWorker(ctx context.Context, in *RegisterWorkerRequest, opts ...grpc.CallOption) (*RegisterWorkerResponse, error)
	UpdateWorker(ctx context.Context, in *UpdateWorkerRequest, opts ...grpc.CallOption) (*UpdateWorkerResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) RegisterWorker(ctx context.Context, in *RegisterWorkerRequest, opts ...grpc.CallOption) (*RegisterWorkerResponse, error) {
	out := new(RegisterWorkerResponse)
	err := c.cc.Invoke(ctx, "/api.MasterService/RegisterWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) UpdateWorker(ctx context.Context, in *UpdateWorkerRequest, opts ...grpc.CallOption) (*UpdateWorkerResponse, error) {
	out := new(UpdateWorkerResponse)
	err := c.cc.Invoke(ctx, "/api.MasterService/UpdateWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
type MasterServiceServer interface {
	RegisterWorker(context.Context, *RegisterWorkerRequest) (*RegisterWorkerResponse, error)
	UpdateWorker(context.Context, *UpdateWorkerRequest) (*UpdateWorkerResponse, error)
}

// UnimplementedMasterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (*UnimplementedMasterServiceServer) RegisterWorker(ctx context.Context, req *RegisterWorkerRequest) (*RegisterWorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWorker not implemented")
}
func (*UnimplementedMasterServiceServer) UpdateWorker(ctx context.Context, req *UpdateWorkerRequest) (*UpdateWorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorker not implemented")
}

func RegisterMasterServiceServer(s *grpc.Server, srv MasterServiceServer) {
	s.RegisterService(&_MasterService_serviceDesc, srv)
}

func _MasterService_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MasterService/RegisterWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).RegisterWorker(ctx, req.(*RegisterWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_UpdateWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).UpdateWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MasterService/UpdateWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).UpdateWorker(ctx, req.(*UpdateWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MasterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorker",
			Handler:    _MasterService_RegisterWorker_Handler,
		},
		{
			MethodName: "UpdateWorker",
			Handler:    _MasterService_UpdateWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerServiceClient interface {
	RunWorker(ctx context.Context, in *RunWorkerRequest, opts ...grpc.CallOption) (*RunWorkerResponse, error)
	GetAgents(ctx context.Context, in *GetAgentsRequest, opts ...grpc.CallOption) (*GetAgentsResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) RunWorker(ctx context.Context, in *RunWorkerRequest, opts ...grpc.CallOption) (*RunWorkerResponse, error) {
	out := new(RunWorkerResponse)
	err := c.cc.Invoke(ctx, "/api.WorkerService/RunWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetAgents(ctx context.Context, in *GetAgentsRequest, opts ...grpc.CallOption) (*GetAgentsResponse, error) {
	out := new(GetAgentsResponse)
	err := c.cc.Invoke(ctx, "/api.WorkerService/GetAgents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
type WorkerServiceServer interface {
	RunWorker(context.Context, *RunWorkerRequest) (*RunWorkerResponse, error)
	GetAgents(context.Context, *GetAgentsRequest) (*GetAgentsResponse, error)
}

// UnimplementedWorkerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (*UnimplementedWorkerServiceServer) RunWorker(ctx context.Context, req *RunWorkerRequest) (*RunWorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunWorker not implemented")
}
func (*UnimplementedWorkerServiceServer) GetAgents(ctx context.Context, req *GetAgentsRequest) (*GetAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgents not implemented")
}

func RegisterWorkerServiceServer(s *grpc.Server, srv WorkerServiceServer) {
	s.RegisterService(&_WorkerService_serviceDesc, srv)
}

func _WorkerService_RunWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).RunWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkerService/RunWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).RunWorker(ctx, req.(*RunWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkerService/GetAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetAgents(ctx, req.(*GetAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunWorker",
			Handler:    _WorkerService_RunWorker_Handler,
		},
		{
			MethodName: "GetAgents",
			Handler:    _WorkerService_GetAgents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

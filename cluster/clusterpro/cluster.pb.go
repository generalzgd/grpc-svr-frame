// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/clusterpro/cluster.proto

package clusterpro

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

type ClusterNodeState int32

const (
	ClusterNodeState_LOOKING   ClusterNodeState = 0
	ClusterNodeState_LEADING   ClusterNodeState = 1
	ClusterNodeState_FALLOWING ClusterNodeState = 2
)

var ClusterNodeState_name = map[int32]string{
	0: "LOOKING",
	1: "LEADING",
	2: "FALLOWING",
}

var ClusterNodeState_value = map[string]int32{
	"LOOKING":   0,
	"LEADING":   1,
	"FALLOWING": 2,
}

func (x ClusterNodeState) String() string {
	return proto.EnumName(ClusterNodeState_name, int32(x))
}

func (ClusterNodeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{0}
}

type ClusterDetect struct {
	FromIp               string   `protobuf:"bytes,1,opt,name=FromIp,proto3" json:"FromIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterDetect) Reset()         { *m = ClusterDetect{} }
func (m *ClusterDetect) String() string { return proto.CompactTextString(m) }
func (*ClusterDetect) ProtoMessage()    {}
func (*ClusterDetect) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{0}
}

func (m *ClusterDetect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterDetect.Unmarshal(m, b)
}
func (m *ClusterDetect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterDetect.Marshal(b, m, deterministic)
}
func (m *ClusterDetect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterDetect.Merge(m, src)
}
func (m *ClusterDetect) XXX_Size() int {
	return xxx_messageInfo_ClusterDetect.Size(m)
}
func (m *ClusterDetect) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterDetect.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterDetect proto.InternalMessageInfo

func (m *ClusterDetect) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

// 集群中某台节点发起投票请求
type ClusterVoteRequest struct {
	// 发起投票请求的节点
	FromIp string `protobuf:"bytes,1,opt,name=FromIp,proto3" json:"FromIp,omitempty"`
	TarIp  string `protobuf:"bytes,2,opt,name=TarIp,proto3" json:"TarIp,omitempty"`
	// 逻辑钟
	Clock                int32    `protobuf:"varint,3,opt,name=Clock,proto3" json:"Clock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterVoteRequest) Reset()         { *m = ClusterVoteRequest{} }
func (m *ClusterVoteRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterVoteRequest) ProtoMessage()    {}
func (*ClusterVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{1}
}

func (m *ClusterVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterVoteRequest.Unmarshal(m, b)
}
func (m *ClusterVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterVoteRequest.Marshal(b, m, deterministic)
}
func (m *ClusterVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterVoteRequest.Merge(m, src)
}
func (m *ClusterVoteRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterVoteRequest.Size(m)
}
func (m *ClusterVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterVoteRequest proto.InternalMessageInfo

func (m *ClusterVoteRequest) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

func (m *ClusterVoteRequest) GetTarIp() string {
	if m != nil {
		return m.TarIp
	}
	return ""
}

func (m *ClusterVoteRequest) GetClock() int32 {
	if m != nil {
		return m.Clock
	}
	return 0
}

// 投票请求响应
type ClusterVoteResponse struct {
	Clock int32 `protobuf:"varint,1,opt,name=Clock,proto3" json:"Clock,omitempty"`
	//    投给谁
	VoteIp string `protobuf:"bytes,2,opt,name=VoteIp,proto3" json:"VoteIp,omitempty"`
	//    投票人的IP
	FromIp string `protobuf:"bytes,3,opt,name=FromIp,proto3" json:"FromIp,omitempty"`
	//    投票响应状态，true:支持，false:拒绝投票，拒绝一般是leader发起，即fromip为leader IP
	State                bool     `protobuf:"varint,4,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterVoteResponse) Reset()         { *m = ClusterVoteResponse{} }
func (m *ClusterVoteResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterVoteResponse) ProtoMessage()    {}
func (*ClusterVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{2}
}

func (m *ClusterVoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterVoteResponse.Unmarshal(m, b)
}
func (m *ClusterVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterVoteResponse.Marshal(b, m, deterministic)
}
func (m *ClusterVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterVoteResponse.Merge(m, src)
}
func (m *ClusterVoteResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterVoteResponse.Size(m)
}
func (m *ClusterVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterVoteResponse proto.InternalMessageInfo

func (m *ClusterVoteResponse) GetClock() int32 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *ClusterVoteResponse) GetVoteIp() string {
	if m != nil {
		return m.VoteIp
	}
	return ""
}

func (m *ClusterVoteResponse) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

func (m *ClusterVoteResponse) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

// 集群心跳
type ClusterHeart struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterHeart) Reset()         { *m = ClusterHeart{} }
func (m *ClusterHeart) String() string { return proto.CompactTextString(m) }
func (*ClusterHeart) ProtoMessage()    {}
func (*ClusterHeart) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{3}
}

func (m *ClusterHeart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterHeart.Unmarshal(m, b)
}
func (m *ClusterHeart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterHeart.Marshal(b, m, deterministic)
}
func (m *ClusterHeart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHeart.Merge(m, src)
}
func (m *ClusterHeart) XXX_Size() int {
	return xxx_messageInfo_ClusterHeart.Size(m)
}
func (m *ClusterHeart) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHeart.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHeart proto.InternalMessageInfo

// 集群心跳响应
type ClusterHeartAck struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterHeartAck) Reset()         { *m = ClusterHeartAck{} }
func (m *ClusterHeartAck) String() string { return proto.CompactTextString(m) }
func (*ClusterHeartAck) ProtoMessage()    {}
func (*ClusterHeartAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{4}
}

func (m *ClusterHeartAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterHeartAck.Unmarshal(m, b)
}
func (m *ClusterHeartAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterHeartAck.Marshal(b, m, deterministic)
}
func (m *ClusterHeartAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHeartAck.Merge(m, src)
}
func (m *ClusterHeartAck) XXX_Size() int {
	return xxx_messageInfo_ClusterHeartAck.Size(m)
}
func (m *ClusterHeartAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHeartAck.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHeartAck proto.InternalMessageInfo

// 集群对点检测
type ClusterPing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterPing) Reset()         { *m = ClusterPing{} }
func (m *ClusterPing) String() string { return proto.CompactTextString(m) }
func (*ClusterPing) ProtoMessage()    {}
func (*ClusterPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{5}
}

func (m *ClusterPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterPing.Unmarshal(m, b)
}
func (m *ClusterPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterPing.Marshal(b, m, deterministic)
}
func (m *ClusterPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterPing.Merge(m, src)
}
func (m *ClusterPing) XXX_Size() int {
	return xxx_messageInfo_ClusterPing.Size(m)
}
func (m *ClusterPing) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterPing.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterPing proto.InternalMessageInfo

type ClusterPingAck struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterPingAck) Reset()         { *m = ClusterPingAck{} }
func (m *ClusterPingAck) String() string { return proto.CompactTextString(m) }
func (*ClusterPingAck) ProtoMessage()    {}
func (*ClusterPingAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{6}
}

func (m *ClusterPingAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterPingAck.Unmarshal(m, b)
}
func (m *ClusterPingAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterPingAck.Marshal(b, m, deterministic)
}
func (m *ClusterPingAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterPingAck.Merge(m, src)
}
func (m *ClusterPingAck) XXX_Size() int {
	return xxx_messageInfo_ClusterPingAck.Size(m)
}
func (m *ClusterPingAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterPingAck.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterPingAck proto.InternalMessageInfo

type ClusterJoinRequest struct {
	//    需要加入集群点IP
	Ip                   string   `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterJoinRequest) Reset()         { *m = ClusterJoinRequest{} }
func (m *ClusterJoinRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterJoinRequest) ProtoMessage()    {}
func (*ClusterJoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{7}
}

func (m *ClusterJoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterJoinRequest.Unmarshal(m, b)
}
func (m *ClusterJoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterJoinRequest.Marshal(b, m, deterministic)
}
func (m *ClusterJoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterJoinRequest.Merge(m, src)
}
func (m *ClusterJoinRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterJoinRequest.Size(m)
}
func (m *ClusterJoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterJoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterJoinRequest proto.InternalMessageInfo

func (m *ClusterJoinRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ClusterJoinResponse struct {
	Ip string `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	//    位于集群中的第几个
	Idx int32 `protobuf:"varint,2,opt,name=Idx,proto3" json:"Idx,omitempty"`
	//    集群共有几个节点
	Cnt int32 `protobuf:"varint,3,opt,name=Cnt,proto3" json:"Cnt,omitempty"`
	//    当前节点的状态, LOOKING LEADING FALLOWING
	State                ClusterNodeState `protobuf:"varint,4,opt,name=State,proto3,enum=clusterpro.ClusterNodeState" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ClusterJoinResponse) Reset()         { *m = ClusterJoinResponse{} }
func (m *ClusterJoinResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterJoinResponse) ProtoMessage()    {}
func (*ClusterJoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{8}
}

func (m *ClusterJoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterJoinResponse.Unmarshal(m, b)
}
func (m *ClusterJoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterJoinResponse.Marshal(b, m, deterministic)
}
func (m *ClusterJoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterJoinResponse.Merge(m, src)
}
func (m *ClusterJoinResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterJoinResponse.Size(m)
}
func (m *ClusterJoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterJoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterJoinResponse proto.InternalMessageInfo

func (m *ClusterJoinResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ClusterJoinResponse) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *ClusterJoinResponse) GetCnt() int32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func (m *ClusterJoinResponse) GetState() ClusterNodeState {
	if m != nil {
		return m.State
	}
	return ClusterNodeState_LOOKING
}

type ClusterCollectRequest struct {
	FromIp               string   `protobuf:"bytes,1,opt,name=FromIp,proto3" json:"FromIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterCollectRequest) Reset()         { *m = ClusterCollectRequest{} }
func (m *ClusterCollectRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterCollectRequest) ProtoMessage()    {}
func (*ClusterCollectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{9}
}

func (m *ClusterCollectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterCollectRequest.Unmarshal(m, b)
}
func (m *ClusterCollectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterCollectRequest.Marshal(b, m, deterministic)
}
func (m *ClusterCollectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterCollectRequest.Merge(m, src)
}
func (m *ClusterCollectRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterCollectRequest.Size(m)
}
func (m *ClusterCollectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterCollectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterCollectRequest proto.InternalMessageInfo

func (m *ClusterCollectRequest) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

type ClusterCollectResponse struct {
	FromIp               string   `protobuf:"bytes,1,opt,name=FromIp,proto3" json:"FromIp,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=Ip,proto3" json:"Ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterCollectResponse) Reset()         { *m = ClusterCollectResponse{} }
func (m *ClusterCollectResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterCollectResponse) ProtoMessage()    {}
func (*ClusterCollectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{10}
}

func (m *ClusterCollectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterCollectResponse.Unmarshal(m, b)
}
func (m *ClusterCollectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterCollectResponse.Marshal(b, m, deterministic)
}
func (m *ClusterCollectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterCollectResponse.Merge(m, src)
}
func (m *ClusterCollectResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterCollectResponse.Size(m)
}
func (m *ClusterCollectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterCollectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterCollectResponse proto.InternalMessageInfo

func (m *ClusterCollectResponse) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

func (m *ClusterCollectResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ClusterNotify struct {
	List                 []*SvrInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ClusterNotify) Reset()         { *m = ClusterNotify{} }
func (m *ClusterNotify) String() string { return proto.CompactTextString(m) }
func (*ClusterNotify) ProtoMessage()    {}
func (*ClusterNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{11}
}

func (m *ClusterNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterNotify.Unmarshal(m, b)
}
func (m *ClusterNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterNotify.Marshal(b, m, deterministic)
}
func (m *ClusterNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterNotify.Merge(m, src)
}
func (m *ClusterNotify) XXX_Size() int {
	return xxx_messageInfo_ClusterNotify.Size(m)
}
func (m *ClusterNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterNotify.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterNotify proto.InternalMessageInfo

func (m *ClusterNotify) GetList() []*SvrInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type SvrInfo struct {
	Ip                   string           `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Idx                  int32            `protobuf:"varint,2,opt,name=Idx,proto3" json:"Idx,omitempty"`
	Cnt                  int32            `protobuf:"varint,3,opt,name=Cnt,proto3" json:"Cnt,omitempty"`
	State                ClusterNodeState `protobuf:"varint,4,opt,name=State,proto3,enum=clusterpro.ClusterNodeState" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SvrInfo) Reset()         { *m = SvrInfo{} }
func (m *SvrInfo) String() string { return proto.CompactTextString(m) }
func (*SvrInfo) ProtoMessage()    {}
func (*SvrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8504b341fb3e598f, []int{12}
}

func (m *SvrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SvrInfo.Unmarshal(m, b)
}
func (m *SvrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SvrInfo.Marshal(b, m, deterministic)
}
func (m *SvrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SvrInfo.Merge(m, src)
}
func (m *SvrInfo) XXX_Size() int {
	return xxx_messageInfo_SvrInfo.Size(m)
}
func (m *SvrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SvrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SvrInfo proto.InternalMessageInfo

func (m *SvrInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *SvrInfo) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *SvrInfo) GetCnt() int32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func (m *SvrInfo) GetState() ClusterNodeState {
	if m != nil {
		return m.State
	}
	return ClusterNodeState_LOOKING
}

func init() {
	proto.RegisterEnum("clusterpro.ClusterNodeState", ClusterNodeState_name, ClusterNodeState_value)
	proto.RegisterType((*ClusterDetect)(nil), "clusterpro.ClusterDetect")
	proto.RegisterType((*ClusterVoteRequest)(nil), "clusterpro.ClusterVoteRequest")
	proto.RegisterType((*ClusterVoteResponse)(nil), "clusterpro.ClusterVoteResponse")
	proto.RegisterType((*ClusterHeart)(nil), "clusterpro.ClusterHeart")
	proto.RegisterType((*ClusterHeartAck)(nil), "clusterpro.ClusterHeartAck")
	proto.RegisterType((*ClusterPing)(nil), "clusterpro.ClusterPing")
	proto.RegisterType((*ClusterPingAck)(nil), "clusterpro.ClusterPingAck")
	proto.RegisterType((*ClusterJoinRequest)(nil), "clusterpro.ClusterJoinRequest")
	proto.RegisterType((*ClusterJoinResponse)(nil), "clusterpro.ClusterJoinResponse")
	proto.RegisterType((*ClusterCollectRequest)(nil), "clusterpro.ClusterCollectRequest")
	proto.RegisterType((*ClusterCollectResponse)(nil), "clusterpro.ClusterCollectResponse")
	proto.RegisterType((*ClusterNotify)(nil), "clusterpro.ClusterNotify")
	proto.RegisterType((*SvrInfo)(nil), "clusterpro.SvrInfo")
}

func init() { proto.RegisterFile("proto/clusterpro/cluster.proto", fileDescriptor_8504b341fb3e598f) }

var fileDescriptor_8504b341fb3e598f = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x4f, 0xea, 0x40,
	0x10, 0xc7, 0x5f, 0x5b, 0x7e, 0x3c, 0x86, 0x47, 0x5f, 0x5d, 0x94, 0xf4, 0x60, 0x4c, 0xb3, 0x31,
	0xa1, 0xf1, 0x00, 0x09, 0x5e, 0x4c, 0xbc, 0x48, 0x8a, 0x68, 0x95, 0x80, 0x29, 0x46, 0xbd, 0x62,
	0x59, 0x4c, 0x43, 0xed, 0x96, 0x76, 0x31, 0x7a, 0xf0, 0x7f, 0x37, 0xdb, 0x2e, 0xa5, 0x36, 0x21,
	0xde, 0xbc, 0xcd, 0x67, 0xe6, 0xcb, 0xce, 0x7c, 0x67, 0x28, 0x1c, 0x85, 0x11, 0x65, 0xb4, 0xeb,
	0xfa, 0xeb, 0x98, 0x91, 0x28, 0x8c, 0xb2, 0xb0, 0x93, 0x14, 0x10, 0x6c, 0x2b, 0xb8, 0x0d, 0x0d,
	0x2b, 0xa5, 0x01, 0x61, 0xc4, 0x65, 0xa8, 0x05, 0x95, 0x61, 0x44, 0x5f, 0xed, 0x50, 0x97, 0x0c,
	0xc9, 0xac, 0x39, 0x82, 0xf0, 0x13, 0x20, 0x21, 0x7c, 0xa0, 0x8c, 0x38, 0x64, 0xb5, 0x26, 0xf1,
	0x4e, 0x35, 0xda, 0x87, 0xf2, 0xfd, 0x2c, 0xb2, 0x43, 0x5d, 0x4e, 0xd2, 0x29, 0xf0, 0xac, 0xe5,
	0x53, 0x77, 0xa9, 0x2b, 0x86, 0x64, 0x96, 0x9d, 0x14, 0xf0, 0x0a, 0x9a, 0xdf, 0x5e, 0x8e, 0x43,
	0x1a, 0xc4, 0x64, 0x2b, 0x96, 0x72, 0x62, 0xde, 0x90, 0xab, 0xb2, 0x97, 0x05, 0xe5, 0x06, 0x51,
	0x8a, 0x83, 0x4c, 0xd9, 0x8c, 0x11, 0xbd, 0x64, 0x48, 0xe6, 0x5f, 0x27, 0x05, 0xac, 0xc2, 0x3f,
	0xd1, 0xf2, 0x9a, 0xcc, 0x22, 0x86, 0xf7, 0xe0, 0x7f, 0x9e, 0xfb, 0xee, 0x12, 0x37, 0xa0, 0x2e,
	0x52, 0x77, 0x5e, 0xf0, 0x82, 0x35, 0x50, 0x73, 0xc8, 0x05, 0xc7, 0xd9, 0x42, 0x6e, 0xa8, 0x17,
	0x6c, 0x16, 0xa2, 0x82, 0x9c, 0x2d, 0x43, 0xb6, 0x43, 0xfc, 0x99, 0x99, 0x4b, 0x55, 0xc2, 0x5c,
	0x41, 0x86, 0x34, 0x50, 0xec, 0xf9, 0x7b, 0xe2, 0xa9, 0xec, 0xf0, 0x90, 0x67, 0xac, 0x80, 0x89,
	0x4d, 0xf1, 0x10, 0xf5, 0xf2, 0x56, 0xd4, 0xde, 0x61, 0x67, 0x7b, 0xc6, 0x8e, 0xe8, 0x31, 0xa6,
	0x73, 0x92, 0x68, 0x36, 0x46, 0xbb, 0x70, 0x20, 0x4a, 0x16, 0xf5, 0x7d, 0xe2, 0xb2, 0x1f, 0x0e,
	0x87, 0x2f, 0xa0, 0x55, 0xfc, 0x81, 0x18, 0x79, 0xd7, 0xa9, 0x53, 0x2b, 0x72, 0xe6, 0xf8, 0x2c,
	0xfb, 0x47, 0x8d, 0x29, 0xf3, 0x16, 0x1f, 0xa8, 0x0d, 0x25, 0xdf, 0x8b, 0x99, 0x2e, 0x19, 0x8a,
	0x59, 0xef, 0x35, 0xf3, 0x63, 0x4f, 0xdf, 0x22, 0x3b, 0x58, 0x50, 0x27, 0x11, 0xe0, 0x15, 0x54,
	0x45, 0xe2, 0xb7, 0xf6, 0x73, 0x72, 0x0e, 0x5a, 0xb1, 0x84, 0xea, 0x50, 0x1d, 0x4d, 0x26, 0xb7,
	0xf6, 0xf8, 0x4a, 0xfb, 0x93, 0xc0, 0x65, 0x7f, 0xc0, 0x41, 0x42, 0x0d, 0xa8, 0x0d, 0xfb, 0xa3,
	0xd1, 0xe4, 0x91, 0xa3, 0xfc, 0x5c, 0x49, 0x3e, 0xa7, 0xd3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0xbe, 0x9e, 0xff, 0x70, 0x03, 0x00, 0x00,
}

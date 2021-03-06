// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package proto_rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AddPeer
type RPCAddPeerReq struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Seed                 int64    `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCAddPeerReq) Reset()         { *m = RPCAddPeerReq{} }
func (m *RPCAddPeerReq) String() string { return proto.CompactTextString(m) }
func (*RPCAddPeerReq) ProtoMessage()    {}
func (*RPCAddPeerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{0}
}
func (m *RPCAddPeerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCAddPeerReq.Unmarshal(m, b)
}
func (m *RPCAddPeerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCAddPeerReq.Marshal(b, m, deterministic)
}
func (dst *RPCAddPeerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCAddPeerReq.Merge(dst, src)
}
func (m *RPCAddPeerReq) XXX_Size() int {
	return xxx_messageInfo_RPCAddPeerReq.Size(m)
}
func (m *RPCAddPeerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCAddPeerReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPCAddPeerReq proto.InternalMessageInfo

func (m *RPCAddPeerReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RPCAddPeerReq) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RPCAddPeerReq) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

// SubscribeShard
type RPCSubscribeShardReq struct {
	ShardIDs             []int64  `protobuf:"varint,1,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCSubscribeShardReq) Reset()         { *m = RPCSubscribeShardReq{} }
func (m *RPCSubscribeShardReq) String() string { return proto.CompactTextString(m) }
func (*RPCSubscribeShardReq) ProtoMessage()    {}
func (*RPCSubscribeShardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{1}
}
func (m *RPCSubscribeShardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCSubscribeShardReq.Unmarshal(m, b)
}
func (m *RPCSubscribeShardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCSubscribeShardReq.Marshal(b, m, deterministic)
}
func (dst *RPCSubscribeShardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCSubscribeShardReq.Merge(dst, src)
}
func (m *RPCSubscribeShardReq) XXX_Size() int {
	return xxx_messageInfo_RPCSubscribeShardReq.Size(m)
}
func (m *RPCSubscribeShardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCSubscribeShardReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPCSubscribeShardReq proto.InternalMessageInfo

func (m *RPCSubscribeShardReq) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

// UnsubscribeShard
type RPCUnsubscribeShardReq struct {
	ShardIDs             []int64  `protobuf:"varint,1,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCUnsubscribeShardReq) Reset()         { *m = RPCUnsubscribeShardReq{} }
func (m *RPCUnsubscribeShardReq) String() string { return proto.CompactTextString(m) }
func (*RPCUnsubscribeShardReq) ProtoMessage()    {}
func (*RPCUnsubscribeShardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{2}
}
func (m *RPCUnsubscribeShardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCUnsubscribeShardReq.Unmarshal(m, b)
}
func (m *RPCUnsubscribeShardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCUnsubscribeShardReq.Marshal(b, m, deterministic)
}
func (dst *RPCUnsubscribeShardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCUnsubscribeShardReq.Merge(dst, src)
}
func (m *RPCUnsubscribeShardReq) XXX_Size() int {
	return xxx_messageInfo_RPCUnsubscribeShardReq.Size(m)
}
func (m *RPCUnsubscribeShardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCUnsubscribeShardReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPCUnsubscribeShardReq proto.InternalMessageInfo

func (m *RPCUnsubscribeShardReq) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

// GetSubscribedShard
type RPCGetSubscribedShardReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCGetSubscribedShardReq) Reset()         { *m = RPCGetSubscribedShardReq{} }
func (m *RPCGetSubscribedShardReq) String() string { return proto.CompactTextString(m) }
func (*RPCGetSubscribedShardReq) ProtoMessage()    {}
func (*RPCGetSubscribedShardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{3}
}
func (m *RPCGetSubscribedShardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCGetSubscribedShardReq.Unmarshal(m, b)
}
func (m *RPCGetSubscribedShardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCGetSubscribedShardReq.Marshal(b, m, deterministic)
}
func (dst *RPCGetSubscribedShardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCGetSubscribedShardReq.Merge(dst, src)
}
func (m *RPCGetSubscribedShardReq) XXX_Size() int {
	return xxx_messageInfo_RPCGetSubscribedShardReq.Size(m)
}
func (m *RPCGetSubscribedShardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCGetSubscribedShardReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPCGetSubscribedShardReq proto.InternalMessageInfo

type RPCGetSubscribedShardReply struct {
	ShardIDs             []int64  `protobuf:"varint,1,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCGetSubscribedShardReply) Reset()         { *m = RPCGetSubscribedShardReply{} }
func (m *RPCGetSubscribedShardReply) String() string { return proto.CompactTextString(m) }
func (*RPCGetSubscribedShardReply) ProtoMessage()    {}
func (*RPCGetSubscribedShardReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{4}
}
func (m *RPCGetSubscribedShardReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCGetSubscribedShardReply.Unmarshal(m, b)
}
func (m *RPCGetSubscribedShardReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCGetSubscribedShardReply.Marshal(b, m, deterministic)
}
func (dst *RPCGetSubscribedShardReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCGetSubscribedShardReply.Merge(dst, src)
}
func (m *RPCGetSubscribedShardReply) XXX_Size() int {
	return xxx_messageInfo_RPCGetSubscribedShardReply.Size(m)
}
func (m *RPCGetSubscribedShardReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCGetSubscribedShardReply.DiscardUnknown(m)
}

var xxx_messageInfo_RPCGetSubscribedShardReply proto.InternalMessageInfo

func (m *RPCGetSubscribedShardReply) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

func (m *RPCGetSubscribedShardReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// BroadcastCollation
type RPCBroadcastCollationReq struct {
	ShardID              int64    `protobuf:"varint,1,opt,name=shardID,proto3" json:"shardID,omitempty"`
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Period               int32    `protobuf:"varint,4,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCBroadcastCollationReq) Reset()         { *m = RPCBroadcastCollationReq{} }
func (m *RPCBroadcastCollationReq) String() string { return proto.CompactTextString(m) }
func (*RPCBroadcastCollationReq) ProtoMessage()    {}
func (*RPCBroadcastCollationReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{5}
}
func (m *RPCBroadcastCollationReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCBroadcastCollationReq.Unmarshal(m, b)
}
func (m *RPCBroadcastCollationReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCBroadcastCollationReq.Marshal(b, m, deterministic)
}
func (dst *RPCBroadcastCollationReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCBroadcastCollationReq.Merge(dst, src)
}
func (m *RPCBroadcastCollationReq) XXX_Size() int {
	return xxx_messageInfo_RPCBroadcastCollationReq.Size(m)
}
func (m *RPCBroadcastCollationReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCBroadcastCollationReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPCBroadcastCollationReq proto.InternalMessageInfo

func (m *RPCBroadcastCollationReq) GetShardID() int64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *RPCBroadcastCollationReq) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RPCBroadcastCollationReq) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RPCBroadcastCollationReq) GetPeriod() int32 {
	if m != nil {
		return m.Period
	}
	return 0
}

// StopServer
type RPCStopServerReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCStopServerReq) Reset()         { *m = RPCStopServerReq{} }
func (m *RPCStopServerReq) String() string { return proto.CompactTextString(m) }
func (*RPCStopServerReq) ProtoMessage()    {}
func (*RPCStopServerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{6}
}
func (m *RPCStopServerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCStopServerReq.Unmarshal(m, b)
}
func (m *RPCStopServerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCStopServerReq.Marshal(b, m, deterministic)
}
func (dst *RPCStopServerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCStopServerReq.Merge(dst, src)
}
func (m *RPCStopServerReq) XXX_Size() int {
	return xxx_messageInfo_RPCStopServerReq.Size(m)
}
func (m *RPCStopServerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCStopServerReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPCStopServerReq proto.InternalMessageInfo

type RPCReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCReply) Reset()         { *m = RPCReply{} }
func (m *RPCReply) String() string { return proto.CompactTextString(m) }
func (*RPCReply) ProtoMessage()    {}
func (*RPCReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_71a1ae02a48c7975, []int{7}
}
func (m *RPCReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCReply.Unmarshal(m, b)
}
func (m *RPCReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCReply.Marshal(b, m, deterministic)
}
func (dst *RPCReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCReply.Merge(dst, src)
}
func (m *RPCReply) XXX_Size() int {
	return xxx_messageInfo_RPCReply.Size(m)
}
func (m *RPCReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCReply.DiscardUnknown(m)
}

var xxx_messageInfo_RPCReply proto.InternalMessageInfo

func (m *RPCReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RPCReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*RPCAddPeerReq)(nil), "proto.rpc.RPCAddPeerReq")
	proto.RegisterType((*RPCSubscribeShardReq)(nil), "proto.rpc.RPCSubscribeShardReq")
	proto.RegisterType((*RPCUnsubscribeShardReq)(nil), "proto.rpc.RPCUnsubscribeShardReq")
	proto.RegisterType((*RPCGetSubscribedShardReq)(nil), "proto.rpc.RPCGetSubscribedShardReq")
	proto.RegisterType((*RPCGetSubscribedShardReply)(nil), "proto.rpc.RPCGetSubscribedShardReply")
	proto.RegisterType((*RPCBroadcastCollationReq)(nil), "proto.rpc.RPCBroadcastCollationReq")
	proto.RegisterType((*RPCStopServerReq)(nil), "proto.rpc.RPCStopServerReq")
	proto.RegisterType((*RPCReply)(nil), "proto.rpc.RPCReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PocClient is the client API for Poc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PocClient interface {
	// Sends a greeting
	AddPeer(ctx context.Context, in *RPCAddPeerReq, opts ...grpc.CallOption) (*RPCReply, error)
	SubscribeShard(ctx context.Context, in *RPCSubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error)
	UnsubscribeShard(ctx context.Context, in *RPCUnsubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error)
	GetSubscribedShard(ctx context.Context, in *RPCGetSubscribedShardReq, opts ...grpc.CallOption) (*RPCGetSubscribedShardReply, error)
	BroadcastCollation(ctx context.Context, in *RPCBroadcastCollationReq, opts ...grpc.CallOption) (*RPCReply, error)
	StopServer(ctx context.Context, in *RPCStopServerReq, opts ...grpc.CallOption) (*RPCReply, error)
}

type pocClient struct {
	cc *grpc.ClientConn
}

func NewPocClient(cc *grpc.ClientConn) PocClient {
	return &pocClient{cc}
}

func (c *pocClient) AddPeer(ctx context.Context, in *RPCAddPeerReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) SubscribeShard(ctx context.Context, in *RPCSubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/SubscribeShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) UnsubscribeShard(ctx context.Context, in *RPCUnsubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/UnsubscribeShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) GetSubscribedShard(ctx context.Context, in *RPCGetSubscribedShardReq, opts ...grpc.CallOption) (*RPCGetSubscribedShardReply, error) {
	out := new(RPCGetSubscribedShardReply)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/GetSubscribedShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) BroadcastCollation(ctx context.Context, in *RPCBroadcastCollationReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/BroadcastCollation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) StopServer(ctx context.Context, in *RPCStopServerReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/StopServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PocServer is the server API for Poc service.
type PocServer interface {
	// Sends a greeting
	AddPeer(context.Context, *RPCAddPeerReq) (*RPCReply, error)
	SubscribeShard(context.Context, *RPCSubscribeShardReq) (*RPCReply, error)
	UnsubscribeShard(context.Context, *RPCUnsubscribeShardReq) (*RPCReply, error)
	GetSubscribedShard(context.Context, *RPCGetSubscribedShardReq) (*RPCGetSubscribedShardReply, error)
	BroadcastCollation(context.Context, *RPCBroadcastCollationReq) (*RPCReply, error)
	StopServer(context.Context, *RPCStopServerReq) (*RPCReply, error)
}

func RegisterPocServer(s *grpc.Server, srv PocServer) {
	s.RegisterService(&_Poc_serviceDesc, srv)
}

func _Poc_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCAddPeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).AddPeer(ctx, req.(*RPCAddPeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_SubscribeShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCSubscribeShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).SubscribeShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/SubscribeShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).SubscribeShard(ctx, req.(*RPCSubscribeShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_UnsubscribeShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCUnsubscribeShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).UnsubscribeShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/UnsubscribeShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).UnsubscribeShard(ctx, req.(*RPCUnsubscribeShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_GetSubscribedShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCGetSubscribedShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).GetSubscribedShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/GetSubscribedShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).GetSubscribedShard(ctx, req.(*RPCGetSubscribedShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_BroadcastCollation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCBroadcastCollationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).BroadcastCollation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/BroadcastCollation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).BroadcastCollation(ctx, req.(*RPCBroadcastCollationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_StopServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCStopServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).StopServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/StopServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).StopServer(ctx, req.(*RPCStopServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Poc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.rpc.Poc",
	HandlerType: (*PocServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPeer",
			Handler:    _Poc_AddPeer_Handler,
		},
		{
			MethodName: "SubscribeShard",
			Handler:    _Poc_SubscribeShard_Handler,
		},
		{
			MethodName: "UnsubscribeShard",
			Handler:    _Poc_UnsubscribeShard_Handler,
		},
		{
			MethodName: "GetSubscribedShard",
			Handler:    _Poc_GetSubscribedShard_Handler,
		},
		{
			MethodName: "BroadcastCollation",
			Handler:    _Poc_BroadcastCollation_Handler,
		},
		{
			MethodName: "StopServer",
			Handler:    _Poc_StopServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_71a1ae02a48c7975) }

var fileDescriptor_rpc_71a1ae02a48c7975 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4f, 0xe3, 0x30,
	0x10, 0xdd, 0x34, 0xfd, 0x1c, 0x69, 0xab, 0xca, 0xbb, 0xaa, 0xac, 0xec, 0x61, 0x83, 0x11, 0x52,
	0x4e, 0x39, 0x14, 0x4e, 0x88, 0x03, 0x10, 0xa4, 0x82, 0xc4, 0xc1, 0x72, 0xc5, 0x0f, 0xc8, 0x87,
	0x05, 0x91, 0xda, 0xda, 0xd8, 0x2e, 0x02, 0xfe, 0x25, 0xff, 0x08, 0xc5, 0x49, 0x0b, 0x69, 0x1b,
	0x10, 0xa7, 0xcc, 0xc4, 0xcf, 0x6f, 0x3c, 0xef, 0x3d, 0x18, 0x28, 0x99, 0x86, 0x52, 0x09, 0x23,
	0xd0, 0xc0, 0x7e, 0x42, 0x25, 0x53, 0x32, 0x85, 0xdf, 0x8c, 0x46, 0x17, 0x59, 0x46, 0x39, 0x57,
	0x8c, 0x3f, 0xa2, 0x21, 0xb4, 0x72, 0x89, 0x1d, 0xdf, 0x09, 0x06, 0xac, 0x95, 0x4b, 0x84, 0xa0,
	0x2d, 0x85, 0x32, 0xb8, 0xe5, 0x3b, 0x41, 0x87, 0xd9, 0xba, 0xf8, 0xa7, 0x39, 0xcf, 0xb0, 0xeb,
	0x3b, 0x81, 0xcb, 0x6c, 0x4d, 0x26, 0xf0, 0x97, 0xd1, 0x68, 0xb6, 0x4a, 0x74, 0xaa, 0xf2, 0x84,
	0xcf, 0x1e, 0x62, 0x95, 0x15, 0x7c, 0x1e, 0xf4, 0x75, 0x51, 0xdf, 0x5c, 0x69, 0xec, 0xf8, 0x6e,
	0xe0, 0xb2, 0x4d, 0x4f, 0x4e, 0x60, 0xcc, 0x68, 0x74, 0xb7, 0xd4, 0x3f, 0xba, 0xe5, 0x01, 0x66,
	0x34, 0x9a, 0x72, 0xb3, 0x19, 0x96, 0xad, 0xef, 0x11, 0x0a, 0x5e, 0xc3, 0x99, 0x9c, 0xbf, 0x7c,
	0xc5, 0x8a, 0xc6, 0xd0, 0xd5, 0x26, 0x36, 0x2b, 0x6d, 0x37, 0xed, 0xb3, 0xaa, 0x23, 0xcf, 0x76,
	0xda, 0xa5, 0x12, 0x71, 0x96, 0xc6, 0xda, 0x44, 0x62, 0x3e, 0x8f, 0x4d, 0x2e, 0x96, 0xc5, 0x2b,
	0x31, 0xf4, 0xaa, 0xfb, 0x56, 0x30, 0x97, 0xad, 0xdb, 0x82, 0x6d, 0xb9, 0x5a, 0x24, 0x5c, 0x55,
	0xba, 0x55, 0x9d, 0x55, 0x2e, 0x7f, 0xe5, 0x56, 0xb9, 0x0e, 0xb3, 0x75, 0x81, 0x95, 0x5c, 0xe5,
	0x22, 0xc3, 0xed, 0x12, 0x5b, 0x76, 0x04, 0xc1, 0xa8, 0x50, 0xd4, 0x08, 0x39, 0xe3, 0xea, 0xc9,
	0xba, 0x43, 0xce, 0xa0, 0xcf, 0x68, 0x54, 0x6e, 0x83, 0xa1, 0xb7, 0xe0, 0x5a, 0xc7, 0xf7, 0xbc,
	0xb2, 0x6b, 0xdd, 0x36, 0xed, 0x32, 0x79, 0x73, 0xc1, 0xa5, 0x22, 0x45, 0xa7, 0xd0, 0xab, 0x1c,
	0x47, 0x38, 0xdc, 0x64, 0x21, 0xac, 0x05, 0xc1, 0xfb, 0x53, 0x3f, 0xb1, 0x33, 0xc9, 0x2f, 0x74,
	0x0d, 0xc3, 0xba, 0xc9, 0xe8, 0x7f, 0x1d, 0xb8, 0x13, 0x81, 0x26, 0xa6, 0x5b, 0x18, 0x6d, 0x5b,
	0x8f, 0x0e, 0xea, 0xd0, 0x3d, 0xd1, 0x68, 0x62, 0x4b, 0x00, 0xed, 0xda, 0x8e, 0x0e, 0xeb, 0xe0,
	0xbd, 0xa1, 0xf1, 0x8e, 0xbe, 0x07, 0x95, 0x33, 0x28, 0xa0, 0xdd, 0x20, 0x6c, 0xcf, 0xd8, 0x1b,
	0x95, 0xa6, 0x57, 0x9f, 0x03, 0x7c, 0x18, 0x8c, 0xfe, 0x6d, 0x29, 0xf9, 0xd9, 0xfa, 0x06, 0x86,
	0xa4, 0x6b, 0xff, 0x1e, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x34, 0x64, 0xc8, 0xe5, 0xdf, 0x03,
	0x00, 0x00,
}

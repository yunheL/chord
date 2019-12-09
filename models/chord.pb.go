// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chord.proto

package models

import proto "protobuf/proto"
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

// Node contains a node ID and address.
type Node struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Node) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ER struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ER) Reset()         { *m = ER{} }
func (m *ER) String() string { return proto.CompactTextString(m) }
func (*ER) ProtoMessage()    {}
func (*ER) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{1}
}
func (m *ER) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ER.Unmarshal(m, b)
}
func (m *ER) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ER.Marshal(b, m, deterministic)
}
func (dst *ER) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ER.Merge(dst, src)
}
func (m *ER) XXX_Size() int {
	return xxx_messageInfo_ER.Size(m)
}
func (m *ER) XXX_DiscardUnknown() {
	xxx_messageInfo_ER.DiscardUnknown(m)
}

var xxx_messageInfo_ER proto.InternalMessageInfo

type ID struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{2}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{3}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{4}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{5}
}
func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (dst *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(dst, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{6}
}
func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (dst *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(dst, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{7}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{8}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type MultiDeleteRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiDeleteRequest) Reset()         { *m = MultiDeleteRequest{} }
func (m *MultiDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*MultiDeleteRequest) ProtoMessage()    {}
func (*MultiDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{9}
}
func (m *MultiDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiDeleteRequest.Unmarshal(m, b)
}
func (m *MultiDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *MultiDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiDeleteRequest.Merge(dst, src)
}
func (m *MultiDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_MultiDeleteRequest.Size(m)
}
func (m *MultiDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiDeleteRequest proto.InternalMessageInfo

func (m *MultiDeleteRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type RequestKeysRequest struct {
	From                 []byte   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []byte   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestKeysRequest) Reset()         { *m = RequestKeysRequest{} }
func (m *RequestKeysRequest) String() string { return proto.CompactTextString(m) }
func (*RequestKeysRequest) ProtoMessage()    {}
func (*RequestKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{10}
}
func (m *RequestKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestKeysRequest.Unmarshal(m, b)
}
func (m *RequestKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestKeysRequest.Marshal(b, m, deterministic)
}
func (dst *RequestKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestKeysRequest.Merge(dst, src)
}
func (m *RequestKeysRequest) XXX_Size() int {
	return xxx_messageInfo_RequestKeysRequest.Size(m)
}
func (m *RequestKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestKeysRequest proto.InternalMessageInfo

func (m *RequestKeysRequest) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *RequestKeysRequest) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{11}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (dst *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(dst, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestKeysResponse struct {
	Values               []*KV    `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestKeysResponse) Reset()         { *m = RequestKeysResponse{} }
func (m *RequestKeysResponse) String() string { return proto.CompactTextString(m) }
func (*RequestKeysResponse) ProtoMessage()    {}
func (*RequestKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_cec7e5290417eec1, []int{12}
}
func (m *RequestKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestKeysResponse.Unmarshal(m, b)
}
func (m *RequestKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestKeysResponse.Marshal(b, m, deterministic)
}
func (dst *RequestKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestKeysResponse.Merge(dst, src)
}
func (m *RequestKeysResponse) XXX_Size() int {
	return xxx_messageInfo_RequestKeysResponse.Size(m)
}
func (m *RequestKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestKeysResponse proto.InternalMessageInfo

func (m *RequestKeysResponse) GetValues() []*KV {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "internal.Node")
	proto.RegisterType((*ER)(nil), "internal.ER")
	proto.RegisterType((*ID)(nil), "internal.ID")
	proto.RegisterType((*GetRequest)(nil), "internal.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "internal.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "internal.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "internal.SetResponse")
	proto.RegisterType((*DeleteRequest)(nil), "internal.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "internal.DeleteResponse")
	proto.RegisterType((*MultiDeleteRequest)(nil), "internal.MultiDeleteRequest")
	proto.RegisterType((*RequestKeysRequest)(nil), "internal.RequestKeysRequest")
	proto.RegisterType((*KV)(nil), "internal.KV")
	proto.RegisterType((*RequestKeysResponse)(nil), "internal.RequestKeysResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chord service

type ChordClient interface {
	// GetPredecessor returns the node believed to be the current predecessor.
	GetPredecessor(ctx context.Context, in *ER, opts ...grpc.CallOption) (*Node, error)
	// GetSuccessor returns the node believed to be the current successor.
	GetSuccessor(ctx context.Context, in *ER, opts ...grpc.CallOption) (*Node, error)
	// Notify notifies Chord that Node thinks it is our predecessor. This has
	// the potential to initiate the transferring of keys.
	Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ER, error)
	// FindSuccessor finds the node the succedes ID. May initiate RPC calls to
	// other nodes.
	FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error)
	// CheckPredecessor checkes whether predecessor has failed.
	CheckPredecessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ER, error)
	// SetPredecessor sets predecessor for a node.
	SetPredecessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ER, error)
	// SetPredecessor sets predecessor for a node.
	SetSuccessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ER, error)
	// Get returns the value in Chord ring for the given key.
	XGet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Set writes a key value pair to the Chord ring.
	XSet(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	// Delete returns the value in Chord ring for the given key.
	XDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Multiple delete returns the value in Chord ring between the given keys.
	XMultiDelete(ctx context.Context, in *MultiDeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// RequestKeys returns the keys between given range from the Chord ring.
	XRequestKeys(ctx context.Context, in *RequestKeysRequest, opts ...grpc.CallOption) (*RequestKeysResponse, error)
}

type chordClient struct {
	cc *grpc.ClientConn
}

func NewChordClient(cc *grpc.ClientConn) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) GetPredecessor(ctx context.Context, in *ER, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/internal.Chord/GetPredecessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetSuccessor(ctx context.Context, in *ER, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/internal.Chord/GetSuccessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ER, error) {
	out := new(ER)
	err := grpc.Invoke(ctx, "/internal.Chord/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/internal.Chord/FindSuccessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) CheckPredecessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ER, error) {
	out := new(ER)
	err := grpc.Invoke(ctx, "/internal.Chord/CheckPredecessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) SetPredecessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ER, error) {
	out := new(ER)
	err := grpc.Invoke(ctx, "/internal.Chord/SetPredecessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) SetSuccessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ER, error) {
	out := new(ER)
	err := grpc.Invoke(ctx, "/internal.Chord/SetSuccessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) XGet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/internal.Chord/XGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) XSet(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := grpc.Invoke(ctx, "/internal.Chord/XSet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) XDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/internal.Chord/XDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) XMultiDelete(ctx context.Context, in *MultiDeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/internal.Chord/XMultiDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) XRequestKeys(ctx context.Context, in *RequestKeysRequest, opts ...grpc.CallOption) (*RequestKeysResponse, error) {
	out := new(RequestKeysResponse)
	err := grpc.Invoke(ctx, "/internal.Chord/XRequestKeys", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chord service

type ChordServer interface {
	// GetPredecessor returns the node believed to be the current predecessor.
	GetPredecessor(context.Context, *ER) (*Node, error)
	// GetSuccessor returns the node believed to be the current successor.
	GetSuccessor(context.Context, *ER) (*Node, error)
	// Notify notifies Chord that Node thinks it is our predecessor. This has
	// the potential to initiate the transferring of keys.
	Notify(context.Context, *Node) (*ER, error)
	// FindSuccessor finds the node the succedes ID. May initiate RPC calls to
	// other nodes.
	FindSuccessor(context.Context, *ID) (*Node, error)
	// CheckPredecessor checkes whether predecessor has failed.
	CheckPredecessor(context.Context, *ID) (*ER, error)
	// SetPredecessor sets predecessor for a node.
	SetPredecessor(context.Context, *Node) (*ER, error)
	// SetPredecessor sets predecessor for a node.
	SetSuccessor(context.Context, *Node) (*ER, error)
	// Get returns the value in Chord ring for the given key.
	XGet(context.Context, *GetRequest) (*GetResponse, error)
	// Set writes a key value pair to the Chord ring.
	XSet(context.Context, *SetRequest) (*SetResponse, error)
	// Delete returns the value in Chord ring for the given key.
	XDelete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Multiple delete returns the value in Chord ring between the given keys.
	XMultiDelete(context.Context, *MultiDeleteRequest) (*DeleteResponse, error)
	// RequestKeys returns the keys between given range from the Chord ring.
	XRequestKeys(context.Context, *RequestKeysRequest) (*RequestKeysResponse, error)
}

func RegisterChordServer(s *grpc.Server, srv ChordServer) {
	s.RegisterService(&_Chord_serviceDesc, srv)
}

func _Chord_GetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ER)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/GetPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetPredecessor(ctx, req.(*ER))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ER)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/GetSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetSuccessor(ctx, req.(*ER))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Notify(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindSuccessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_CheckPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).CheckPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/CheckPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).CheckPredecessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_SetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).SetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/SetPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).SetPredecessor(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_SetSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).SetSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/SetSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).SetSuccessor(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_XGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).XGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/XGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).XGet(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_XSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).XSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/XSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).XSet(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_XDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).XDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/XDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).XDelete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_XMultiDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).XMultiDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/XMultiDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).XMultiDelete(ctx, req.(*MultiDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_XRequestKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).XRequestKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Chord/XRequestKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).XRequestKeys(ctx, req.(*RequestKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPredecessor",
			Handler:    _Chord_GetPredecessor_Handler,
		},
		{
			MethodName: "GetSuccessor",
			Handler:    _Chord_GetSuccessor_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Chord_Notify_Handler,
		},
		{
			MethodName: "FindSuccessor",
			Handler:    _Chord_FindSuccessor_Handler,
		},
		{
			MethodName: "CheckPredecessor",
			Handler:    _Chord_CheckPredecessor_Handler,
		},
		{
			MethodName: "SetPredecessor",
			Handler:    _Chord_SetPredecessor_Handler,
		},
		{
			MethodName: "SetSuccessor",
			Handler:    _Chord_SetSuccessor_Handler,
		},
		{
			MethodName: "XGet",
			Handler:    _Chord_XGet_Handler,
		},
		{
			MethodName: "XSet",
			Handler:    _Chord_XSet_Handler,
		},
		{
			MethodName: "XDelete",
			Handler:    _Chord_XDelete_Handler,
		},
		{
			MethodName: "XMultiDelete",
			Handler:    _Chord_XMultiDelete_Handler,
		},
		{
			MethodName: "XRequestKeys",
			Handler:    _Chord_XRequestKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chord.proto",
}

func init() { proto.RegisterFile("chord.proto", fileDescriptor_chord_cec7e5290417eec1) }

var fileDescriptor_chord_cec7e5290417eec1 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6b, 0xd3, 0x50,
	0x14, 0xa6, 0x59, 0x5a, 0xdd, 0xd7, 0x34, 0x94, 0x63, 0xc5, 0x12, 0x54, 0xea, 0x55, 0xa4, 0x48,
	0xcd, 0xc3, 0xe6, 0x83, 0xa0, 0x6f, 0xeb, 0x56, 0x4a, 0x71, 0x48, 0x02, 0xa3, 0xaf, 0xb5, 0x39,
	0x63, 0xa1, 0xb5, 0x77, 0x26, 0xb7, 0x42, 0xff, 0x02, 0xff, 0x6d, 0xc9, 0xcd, 0x1d, 0x49, 0xda,
	0x35, 0xec, 0xed, 0xdc, 0xfb, 0xfd, 0xc8, 0x39, 0xf7, 0x7c, 0x04, 0xed, 0xe5, 0x9d, 0x4c, 0x22,
	0xff, 0x3e, 0x91, 0x4a, 0xd2, 0xf3, 0x78, 0xa3, 0x38, 0xd9, 0x2c, 0xd6, 0xe2, 0x13, 0xec, 0x6b,
	0x19, 0x31, 0xb9, 0xb0, 0xe2, 0xa8, 0xdf, 0x18, 0x34, 0x86, 0x4e, 0x60, 0xc5, 0x11, 0x11, 0xec,
	0x45, 0x14, 0x25, 0x7d, 0x6b, 0xd0, 0x18, 0x9e, 0x06, 0xba, 0x16, 0x36, 0xac, 0xcb, 0x40, 0xf4,
	0x60, 0x4d, 0xc7, 0xfb, 0x7c, 0xf1, 0x16, 0x98, 0xb0, 0x0a, 0xf8, 0xcf, 0x96, 0x53, 0x45, 0x5d,
	0x9c, 0xac, 0x78, 0xa7, 0xe1, 0xd3, 0x20, 0x2b, 0xc5, 0x7b, 0xb4, 0x35, 0x9e, 0xde, 0xcb, 0x4d,
	0xca, 0xd4, 0x43, 0xf3, 0xef, 0x62, 0xbd, 0x65, 0xe3, 0x90, 0x1f, 0xc4, 0x17, 0x20, 0xac, 0x31,
	0x29, 0x54, 0x79, 0x57, 0x46, 0xd5, 0x41, 0x3b, 0x2c, 0xac, 0xc5, 0x3b, 0x74, 0xc6, 0xbc, 0x66,
	0xc5, 0xc7, 0x9b, 0xe9, 0xc2, 0x7d, 0xa0, 0x18, 0xd1, 0x10, 0xf4, 0x63, 0xbb, 0x56, 0x71, 0x55,
	0x49, 0xb0, 0x57, 0xbc, 0x4b, 0xfb, 0x8d, 0xc1, 0x49, 0xf6, 0x08, 0x59, 0x2d, 0xbe, 0x82, 0x0c,
	0x3c, 0xe3, 0x5d, 0x5a, 0x62, 0xde, 0x26, 0xf2, 0xb7, 0x19, 0x47, 0xd7, 0xd9, 0x13, 0x29, 0xa9,
	0x5b, 0x75, 0x02, 0x4b, 0x49, 0x31, 0x82, 0x35, 0xbb, 0x79, 0xf2, 0x54, 0xdf, 0xf0, 0xa2, 0xf2,
	0x1d, 0xf3, 0x70, 0x1f, 0xd0, 0xd2, 0x78, 0xde, 0x54, 0xfb, 0xcc, 0xf1, 0x1f, 0x56, 0xe9, 0xcf,
	0x6e, 0x02, 0x83, 0x9d, 0xfd, 0x6b, 0xa2, 0x79, 0x91, 0xed, 0x9b, 0x7c, 0xb8, 0x13, 0x56, 0x3f,
	0x13, 0x8e, 0x78, 0xc9, 0x69, 0x2a, 0x13, 0x2a, 0x29, 0x2e, 0x03, 0xcf, 0x2d, 0x4e, 0x3a, 0x07,
	0x23, 0x38, 0x13, 0x56, 0xe1, 0x76, 0xf9, 0x24, 0xf6, 0x47, 0xb4, 0xae, 0xa5, 0x8a, 0x6f, 0x77,
	0xb4, 0x87, 0x78, 0x15, 0x1d, 0x7d, 0x46, 0xe7, 0x2a, 0xde, 0x44, 0x8f, 0xda, 0x4e, 0xc7, 0x07,
	0xb6, 0x3e, 0xba, 0x17, 0x77, 0xbc, 0x5c, 0x1d, 0x69, 0x7b, 0x3a, 0xde, 0xb3, 0xf7, 0xe1, 0x86,
	0xd5, 0x21, 0xeb, 0xdb, 0x19, 0xc1, 0x09, 0xcb, 0x43, 0xd6, 0xb3, 0xcf, 0x61, 0xcf, 0x27, 0xac,
	0xa8, 0x57, 0xdc, 0x16, 0x51, 0xf7, 0x5e, 0xee, 0xdd, 0x9a, 0x3d, 0x65, 0xa2, 0xb0, 0x2a, 0x0a,
	0x1f, 0x15, 0x95, 0xa2, 0x4b, 0xdf, 0xf1, 0x6c, 0x9e, 0x27, 0x90, 0x5e, 0x15, 0x8c, 0x4a, 0x26,
	0xbd, 0xfe, 0x21, 0x60, 0xd4, 0x57, 0x70, 0xe6, 0xa5, 0x10, 0xd3, 0xeb, 0x82, 0x79, 0x98, 0xed,
	0x1a, 0x9f, 0x19, 0x9c, 0x79, 0x29, 0x7a, 0x65, 0x9f, 0xc3, 0xe4, 0x7b, 0x6f, 0x8e, 0xa0, 0xb9,
	0xd9, 0xaf, 0x96, 0xfe, 0xe1, 0x9c, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x92, 0xe1, 0xed, 0x2c,
	0x7f, 0x04, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todolist.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddTodoResp struct {
	AddTodoResp          *Todo    `protobuf:"bytes,1,opt,name=addTodoResp" json:"addTodoResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTodoResp) Reset()         { *m = AddTodoResp{} }
func (m *AddTodoResp) String() string { return proto.CompactTextString(m) }
func (*AddTodoResp) ProtoMessage()    {}
func (*AddTodoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_92bea43d73891c00, []int{0}
}
func (m *AddTodoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTodoResp.Unmarshal(m, b)
}
func (m *AddTodoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTodoResp.Marshal(b, m, deterministic)
}
func (dst *AddTodoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTodoResp.Merge(dst, src)
}
func (m *AddTodoResp) XXX_Size() int {
	return xxx_messageInfo_AddTodoResp.Size(m)
}
func (m *AddTodoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTodoResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddTodoResp proto.InternalMessageInfo

func (m *AddTodoResp) GetAddTodoResp() *Todo {
	if m != nil {
		return m.AddTodoResp
	}
	return nil
}

type AddTodoReq struct {
	AddTodoReq           *Todo    `protobuf:"bytes,1,opt,name=addTodoReq" json:"addTodoReq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTodoReq) Reset()         { *m = AddTodoReq{} }
func (m *AddTodoReq) String() string { return proto.CompactTextString(m) }
func (*AddTodoReq) ProtoMessage()    {}
func (*AddTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_92bea43d73891c00, []int{1}
}
func (m *AddTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTodoReq.Unmarshal(m, b)
}
func (m *AddTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTodoReq.Marshal(b, m, deterministic)
}
func (dst *AddTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTodoReq.Merge(dst, src)
}
func (m *AddTodoReq) XXX_Size() int {
	return xxx_messageInfo_AddTodoReq.Size(m)
}
func (m *AddTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddTodoReq proto.InternalMessageInfo

func (m *AddTodoReq) GetAddTodoReq() *Todo {
	if m != nil {
		return m.AddTodoReq
	}
	return nil
}

type GetTodoResp struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResp) Reset()         { *m = GetTodoResp{} }
func (m *GetTodoResp) String() string { return proto.CompactTextString(m) }
func (*GetTodoResp) ProtoMessage()    {}
func (*GetTodoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_92bea43d73891c00, []int{2}
}
func (m *GetTodoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResp.Unmarshal(m, b)
}
func (m *GetTodoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResp.Marshal(b, m, deterministic)
}
func (dst *GetTodoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResp.Merge(dst, src)
}
func (m *GetTodoResp) XXX_Size() int {
	return xxx_messageInfo_GetTodoResp.Size(m)
}
func (m *GetTodoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResp proto.InternalMessageInfo

func (m *GetTodoResp) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type Todo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Author               string   `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Done                 bool     `protobuf:"varint,4,opt,name=done" json:"done,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_92bea43d73891c00, []int{3}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Todo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type GetTodoReq struct {
	Int                  int64    `protobuf:"varint,1,opt,name=int" json:"int,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoReq) Reset()         { *m = GetTodoReq{} }
func (m *GetTodoReq) String() string { return proto.CompactTextString(m) }
func (*GetTodoReq) ProtoMessage()    {}
func (*GetTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_92bea43d73891c00, []int{4}
}
func (m *GetTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoReq.Unmarshal(m, b)
}
func (m *GetTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoReq.Marshal(b, m, deterministic)
}
func (dst *GetTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoReq.Merge(dst, src)
}
func (m *GetTodoReq) XXX_Size() int {
	return xxx_messageInfo_GetTodoReq.Size(m)
}
func (m *GetTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoReq proto.InternalMessageInfo

func (m *GetTodoReq) GetInt() int64 {
	if m != nil {
		return m.Int
	}
	return 0
}

type ListTodosResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosResponse) Reset()         { *m = ListTodosResponse{} }
func (m *ListTodosResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodosResponse) ProtoMessage()    {}
func (*ListTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_92bea43d73891c00, []int{5}
}
func (m *ListTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosResponse.Unmarshal(m, b)
}
func (m *ListTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosResponse.Marshal(b, m, deterministic)
}
func (dst *ListTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosResponse.Merge(dst, src)
}
func (m *ListTodosResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodosResponse.Size(m)
}
func (m *ListTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosResponse proto.InternalMessageInfo

func (m *ListTodosResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

func init() {
	proto.RegisterType((*AddTodoResp)(nil), "todolistpb.AddTodoResp")
	proto.RegisterType((*AddTodoReq)(nil), "todolistpb.AddTodoReq")
	proto.RegisterType((*GetTodoResp)(nil), "todolistpb.GetTodoResp")
	proto.RegisterType((*Todo)(nil), "todolistpb.Todo")
	proto.RegisterType((*GetTodoReq)(nil), "todolistpb.GetTodoReq")
	proto.RegisterType((*ListTodosResponse)(nil), "todolistpb.ListTodosResponse")
}

func init() { proto.RegisterFile("todolist.proto", fileDescriptor_todolist_92bea43d73891c00) }

var fileDescriptor_todolist_92bea43d73891c00 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0xcb, 0xd3, 0x40,
	0x10, 0xc6, 0xc9, 0x9f, 0xf7, 0x6d, 0x3b, 0x91, 0x12, 0x07, 0xac, 0x31, 0x5a, 0x09, 0x8b, 0x48,
	0x4f, 0x89, 0xb4, 0x47, 0x41, 0xa8, 0x20, 0x5e, 0x7a, 0x0a, 0x7a, 0xd0, 0x8b, 0x24, 0xdd, 0xb5,
	0x2e, 0xd4, 0xec, 0x36, 0xbb, 0x3d, 0x88, 0x78, 0xd1, 0x9b, 0x27, 0xc1, 0x8f, 0xe6, 0x57, 0xf0,
	0x83, 0xc8, 0x6e, 0xd2, 0x24, 0xc5, 0xf6, 0xb6, 0x33, 0xcf, 0xb3, 0xbf, 0xdd, 0x87, 0x19, 0x98,
	0x6a, 0x41, 0xc5, 0x9e, 0x2b, 0x9d, 0xca, 0x5a, 0x68, 0x81, 0x70, 0xaa, 0x65, 0x19, 0x3f, 0xdc,
	0x09, 0xb1, 0xdb, 0xb3, 0xcc, 0x2a, 0xe5, 0xf1, 0x63, 0xc6, 0x3e, 0x4b, 0xfd, 0xa5, 0x31, 0xc6,
	0x8f, 0x5a, 0xb1, 0x90, 0x3c, 0x2b, 0xaa, 0x4a, 0xe8, 0x42, 0x73, 0x51, 0xa9, 0x46, 0x25, 0x6b,
	0x08, 0xd6, 0x94, 0xbe, 0x11, 0x54, 0xe4, 0x4c, 0x49, 0x5c, 0x42, 0x50, 0xf4, 0x65, 0xe4, 0x24,
	0xce, 0x22, 0x58, 0x86, 0x69, 0xff, 0x56, 0x6a, 0xb5, 0xa1, 0x89, 0xbc, 0x00, 0xe8, 0x10, 0x07,
	0x7c, 0x06, 0xd0, 0x89, 0x87, 0xab, 0x80, 0x81, 0x87, 0xac, 0x20, 0x78, 0xcd, 0x74, 0xf7, 0x85,
	0x27, 0xe0, 0x1b, 0xf7, 0xd5, 0xab, 0x56, 0x25, 0x3f, 0x1c, 0xf0, 0x4d, 0x89, 0x53, 0x70, 0x39,
	0xb5, 0x66, 0x2f, 0x77, 0x39, 0xc5, 0x19, 0xdc, 0x16, 0x47, 0xfd, 0x49, 0xd4, 0x91, 0x9b, 0x38,
	0x8b, 0x49, 0xde, 0x56, 0x98, 0x40, 0x40, 0x99, 0xda, 0xd6, 0x5c, 0x9a, 0xf8, 0x91, 0x67, 0xc5,
	0x61, 0x0b, 0x11, 0x7c, 0x2a, 0x2a, 0x16, 0xf9, 0x89, 0xb3, 0x18, 0xe7, 0xf6, 0x8c, 0x73, 0x80,
	0x6d, 0xcd, 0x0a, 0xcd, 0xe8, 0x87, 0x42, 0x47, 0x37, 0xf6, 0xd2, 0xa4, 0xed, 0xac, 0x35, 0x79,
	0x0c, 0xd0, 0x7d, 0xfd, 0x80, 0x21, 0x78, 0xbc, 0xd2, 0xed, 0x5f, 0xcc, 0x91, 0x3c, 0x87, 0xbb,
	0x1b, 0xae, 0xac, 0x41, 0x99, 0x70, 0xa2, 0x52, 0x0c, 0x9f, 0xc2, 0x8d, 0x89, 0xa0, 0x22, 0x27,
	0xf1, 0x2e, 0x26, 0x6c, 0xe4, 0xe5, 0x4f, 0x17, 0xc6, 0xa6, 0x36, 0x04, 0x7c, 0x07, 0x93, 0x8e,
	0x84, 0xb3, 0xb4, 0x99, 0x69, 0x7a, 0x1a, 0x78, 0xfa, 0xca, 0x0c, 0x3c, 0x9e, 0x0f, 0x51, 0xff,
	0x3d, 0x4c, 0xee, 0x7d, 0xff, 0xf3, 0xf7, 0xb7, 0x7b, 0x07, 0x21, 0xdb, 0x9f, 0xb4, 0x5f, 0xae,
	0x83, 0x6f, 0x61, 0xd4, 0x86, 0xc0, 0xd9, 0x10, 0xd0, 0x27, 0x8b, 0xef, 0x5f, 0xec, 0x2b, 0x49,
	0x1e, 0x58, 0x64, 0x88, 0xd3, 0x6c, 0xd7, 0x74, 0xb3, 0xaf, 0xbc, 0xd2, 0xdf, 0x0c, 0x76, 0x03,
	0xa3, 0x76, 0x2d, 0xce, 0xb1, 0xfd, 0xae, 0x9c, 0x63, 0x07, 0x6b, 0x48, 0x42, 0x8b, 0x05, 0x32,
	0xce, 0xda, 0x3d, 0x79, 0xe9, 0xbf, 0x77, 0x65, 0x59, 0xde, 0xda, 0xc0, 0xab, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfe, 0x0f, 0x05, 0x0f, 0x0d, 0x03, 0x00, 0x00,
}

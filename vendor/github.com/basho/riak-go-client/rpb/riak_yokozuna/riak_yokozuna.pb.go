// Code generated by protoc-gen-go.
// source: riak_yokozuna.proto
// DO NOT EDIT!

/*
Package riak_yokozuna is a generated protocol buffer package.

It is generated from these files:
	riak_yokozuna.proto

It has these top-level messages:
	RpbYokozunaIndex
	RpbYokozunaIndexGetReq
	RpbYokozunaIndexGetResp
	RpbYokozunaIndexPutReq
	RpbYokozunaIndexDeleteReq
	RpbYokozunaSchema
	RpbYokozunaSchemaPutReq
	RpbYokozunaSchemaGetReq
	RpbYokozunaSchemaGetResp
*/
package riak_yokozuna

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

type RpbYokozunaIndex struct {
	Name             []byte  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Schema           []byte  `protobuf:"bytes,2,opt,name=schema" json:"schema,omitempty"`
	NVal             *uint32 `protobuf:"varint,3,opt,name=n_val" json:"n_val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbYokozunaIndex) Reset()                    { *m = RpbYokozunaIndex{} }
func (m *RpbYokozunaIndex) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndex) ProtoMessage()               {}
func (*RpbYokozunaIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpbYokozunaIndex) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaIndex) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *RpbYokozunaIndex) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

// GET request - If a name is given, return matching index, else return all
type RpbYokozunaIndexGetReq struct {
	Name             []byte `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexGetReq) Reset()                    { *m = RpbYokozunaIndexGetReq{} }
func (m *RpbYokozunaIndexGetReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetReq) ProtoMessage()               {}
func (*RpbYokozunaIndexGetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RpbYokozunaIndexGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaIndexGetResp struct {
	Index            []*RpbYokozunaIndex `protobuf:"bytes,1,rep,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RpbYokozunaIndexGetResp) Reset()                    { *m = RpbYokozunaIndexGetResp{} }
func (m *RpbYokozunaIndexGetResp) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetResp) ProtoMessage()               {}
func (*RpbYokozunaIndexGetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RpbYokozunaIndexGetResp) GetIndex() []*RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

// PUT request - Create a new index
type RpbYokozunaIndexPutReq struct {
	Index            *RpbYokozunaIndex `protobuf:"bytes,1,req,name=index" json:"index,omitempty"`
	Timeout          *uint32           `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RpbYokozunaIndexPutReq) Reset()                    { *m = RpbYokozunaIndexPutReq{} }
func (m *RpbYokozunaIndexPutReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexPutReq) ProtoMessage()               {}
func (*RpbYokozunaIndexPutReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RpbYokozunaIndexPutReq) GetIndex() *RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *RpbYokozunaIndexPutReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

// DELETE request - Remove an index
type RpbYokozunaIndexDeleteReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexDeleteReq) Reset()                    { *m = RpbYokozunaIndexDeleteReq{} }
func (m *RpbYokozunaIndexDeleteReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexDeleteReq) ProtoMessage()               {}
func (*RpbYokozunaIndexDeleteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RpbYokozunaIndexDeleteReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchema struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Content          []byte `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchema) Reset()                    { *m = RpbYokozunaSchema{} }
func (m *RpbYokozunaSchema) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchema) ProtoMessage()               {}
func (*RpbYokozunaSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RpbYokozunaSchema) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaSchema) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// PUT request - create or potentially update a new schema
type RpbYokozunaSchemaPutReq struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaPutReq) Reset()                    { *m = RpbYokozunaSchemaPutReq{} }
func (m *RpbYokozunaSchemaPutReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaPutReq) ProtoMessage()               {}
func (*RpbYokozunaSchemaPutReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RpbYokozunaSchemaPutReq) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

// GET request - Return matching schema by name
type RpbYokozunaSchemaGetReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchemaGetReq) Reset()                    { *m = RpbYokozunaSchemaGetReq{} }
func (m *RpbYokozunaSchemaGetReq) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetReq) ProtoMessage()               {}
func (*RpbYokozunaSchemaGetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RpbYokozunaSchemaGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchemaGetResp struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaGetResp) Reset()                    { *m = RpbYokozunaSchemaGetResp{} }
func (m *RpbYokozunaSchemaGetResp) String() string            { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetResp) ProtoMessage()               {}
func (*RpbYokozunaSchemaGetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RpbYokozunaSchemaGetResp) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func init() {
	proto.RegisterType((*RpbYokozunaIndex)(nil), "RpbYokozunaIndex")
	proto.RegisterType((*RpbYokozunaIndexGetReq)(nil), "RpbYokozunaIndexGetReq")
	proto.RegisterType((*RpbYokozunaIndexGetResp)(nil), "RpbYokozunaIndexGetResp")
	proto.RegisterType((*RpbYokozunaIndexPutReq)(nil), "RpbYokozunaIndexPutReq")
	proto.RegisterType((*RpbYokozunaIndexDeleteReq)(nil), "RpbYokozunaIndexDeleteReq")
	proto.RegisterType((*RpbYokozunaSchema)(nil), "RpbYokozunaSchema")
	proto.RegisterType((*RpbYokozunaSchemaPutReq)(nil), "RpbYokozunaSchemaPutReq")
	proto.RegisterType((*RpbYokozunaSchemaGetReq)(nil), "RpbYokozunaSchemaGetReq")
	proto.RegisterType((*RpbYokozunaSchemaGetResp)(nil), "RpbYokozunaSchemaGetResp")
}

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0x6a, 0x15, 0xc6, 0xa4, 0xda, 0x15, 0x6c, 0xbc, 0x85, 0x39, 0x68, 0x7b, 0xd9,
	0x43, 0xaf, 0xa2, 0x42, 0x10, 0x44, 0xbc, 0x94, 0x78, 0xf2, 0x54, 0x36, 0x71, 0xa4, 0xa1, 0xcd,
	0x6e, 0x6c, 0x36, 0xa2, 0x7e, 0x7a, 0x37, 0x49, 0x95, 0xfc, 0x3b, 0x78, 0x9c, 0xe5, 0xf7, 0xde,
	0xbc, 0x79, 0x2c, 0x9c, 0xed, 0x62, 0xb1, 0x59, 0x7d, 0xa9, 0x8d, 0xfa, 0xce, 0xa5, 0xe0, 0xe9,
	0x4e, 0x69, 0x85, 0x77, 0x70, 0x1a, 0xa4, 0xe1, 0xcb, 0xfe, 0xf1, 0x51, 0xbe, 0xd2, 0x27, 0xb3,
	0xe1, 0x40, 0x8a, 0x84, 0xdc, 0x81, 0x67, 0xcd, 0x6c, 0x36, 0x86, 0xc3, 0x2c, 0x5a, 0x53, 0x22,
	0x5c, 0xcb, 0x1b, 0x98, 0xd9, 0x81, 0x91, 0x5c, 0x7d, 0x88, 0xad, 0x3b, 0x34, 0xa3, 0x83, 0x97,
	0x70, 0xde, 0x36, 0x78, 0x20, 0x1d, 0xd0, 0x7b, 0xcd, 0xc6, 0xc8, 0xf0, 0x1a, 0xa6, 0xbd, 0x5c,
	0x96, 0x32, 0x0f, 0x46, 0x71, 0x31, 0x1b, 0x72, 0x38, 0x3b, 0x5e, 0x4c, 0x78, 0x1b, 0xc4, 0xa7,
	0xee, 0x92, 0x65, 0x5e, 0x2e, 0xa9, 0x69, 0xad, 0x5e, 0x2d, 0x3b, 0x81, 0x23, 0x1d, 0x27, 0xa4,
	0x72, 0x5d, 0x1e, 0xe0, 0xe0, 0x1c, 0x2e, 0xda, 0xd0, 0x3d, 0x6d, 0x49, 0x53, 0x33, 0xb4, 0xb9,
	0x1d, 0x17, 0x30, 0xa9, 0xa1, 0xcf, 0x65, 0x0d, 0xad, 0x7a, 0x8c, 0x7d, 0xa4, 0xa4, 0x26, 0x59,
	0xd9, 0xdb, 0x78, 0xd3, 0x38, 0xb4, 0xd2, 0xec, 0xc3, 0xe2, 0x5f, 0x95, 0x55, 0x5a, 0xc6, 0x3b,
	0x24, 0x5e, 0xf5, 0xc8, 0x3b, 0x85, 0x16, 0xd9, 0x6e, 0xc1, 0xed, 0x07, 0x4d, 0xa3, 0xff, 0x58,
	0xe4, 0xcf, 0x61, 0x1a, 0xa9, 0x84, 0x87, 0x22, 0x5b, 0x2b, 0x5e, 0x7c, 0x8d, 0xea, 0x47, 0x84,
	0xf9, 0x9b, 0x3f, 0x0e, 0xcc, 0xf8, 0x8b, 0x2f, 0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5,
	0x13, 0xa5, 0x19, 0x3a, 0x02, 0x00, 0x00,
}
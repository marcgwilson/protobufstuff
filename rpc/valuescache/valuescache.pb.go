// Code generated by protoc-gen-go. DO NOT EDIT.
// source: valuescache.proto

package valuescache

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

// ValuesCache consists of a mode and Response name.
type ValuesCache struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Start                float64  `protobuf:"fixed64,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  float64  `protobuf:"fixed64,4,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValuesCache) Reset()         { *m = ValuesCache{} }
func (m *ValuesCache) String() string { return proto.CompactTextString(m) }
func (*ValuesCache) ProtoMessage()    {}
func (*ValuesCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_44747599cffefe15, []int{0}
}

func (m *ValuesCache) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValuesCache.Unmarshal(m, b)
}
func (m *ValuesCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValuesCache.Marshal(b, m, deterministic)
}
func (m *ValuesCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValuesCache.Merge(m, src)
}
func (m *ValuesCache) XXX_Size() int {
	return xxx_messageInfo_ValuesCache.Size(m)
}
func (m *ValuesCache) XXX_DiscardUnknown() {
	xxx_messageInfo_ValuesCache.DiscardUnknown(m)
}

var xxx_messageInfo_ValuesCache proto.InternalMessageInfo

func (m *ValuesCache) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ValuesCache) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ValuesCache) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ValuesCache) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

type ValuesCacheList struct {
	Resourceid           *ResourceId    `protobuf:"bytes,1,opt,name=resourceid,proto3" json:"resourceid,omitempty"`
	Valuescache          []*ValuesCache `protobuf:"bytes,2,rep,name=valuescache,proto3" json:"valuescache,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ValuesCacheList) Reset()         { *m = ValuesCacheList{} }
func (m *ValuesCacheList) String() string { return proto.CompactTextString(m) }
func (*ValuesCacheList) ProtoMessage()    {}
func (*ValuesCacheList) Descriptor() ([]byte, []int) {
	return fileDescriptor_44747599cffefe15, []int{1}
}

func (m *ValuesCacheList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValuesCacheList.Unmarshal(m, b)
}
func (m *ValuesCacheList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValuesCacheList.Marshal(b, m, deterministic)
}
func (m *ValuesCacheList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValuesCacheList.Merge(m, src)
}
func (m *ValuesCacheList) XXX_Size() int {
	return xxx_messageInfo_ValuesCacheList.Size(m)
}
func (m *ValuesCacheList) XXX_DiscardUnknown() {
	xxx_messageInfo_ValuesCacheList.DiscardUnknown(m)
}

var xxx_messageInfo_ValuesCacheList proto.InternalMessageInfo

func (m *ValuesCacheList) GetResourceid() *ResourceId {
	if m != nil {
		return m.Resourceid
	}
	return nil
}

func (m *ValuesCacheList) GetValuescache() []*ValuesCache {
	if m != nil {
		return m.Valuescache
	}
	return nil
}

// The response message containing request status
type ValuesCacheResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValuesCacheResult) Reset()         { *m = ValuesCacheResult{} }
func (m *ValuesCacheResult) String() string { return proto.CompactTextString(m) }
func (*ValuesCacheResult) ProtoMessage()    {}
func (*ValuesCacheResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_44747599cffefe15, []int{2}
}

func (m *ValuesCacheResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValuesCacheResult.Unmarshal(m, b)
}
func (m *ValuesCacheResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValuesCacheResult.Marshal(b, m, deterministic)
}
func (m *ValuesCacheResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValuesCacheResult.Merge(m, src)
}
func (m *ValuesCacheResult) XXX_Size() int {
	return xxx_messageInfo_ValuesCacheResult.Size(m)
}
func (m *ValuesCacheResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ValuesCacheResult.DiscardUnknown(m)
}

var xxx_messageInfo_ValuesCacheResult proto.InternalMessageInfo

func (m *ValuesCacheResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ResourceId struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceId) Reset()         { *m = ResourceId{} }
func (m *ResourceId) String() string { return proto.CompactTextString(m) }
func (*ResourceId) ProtoMessage()    {}
func (*ResourceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_44747599cffefe15, []int{3}
}

func (m *ResourceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceId.Unmarshal(m, b)
}
func (m *ResourceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceId.Marshal(b, m, deterministic)
}
func (m *ResourceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceId.Merge(m, src)
}
func (m *ResourceId) XXX_Size() int {
	return xxx_messageInfo_ResourceId.Size(m)
}
func (m *ResourceId) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceId.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceId proto.InternalMessageInfo

func (m *ResourceId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ValuesCache)(nil), "valuescache.ValuesCache")
	proto.RegisterType((*ValuesCacheList)(nil), "valuescache.ValuesCacheList")
	proto.RegisterType((*ValuesCacheResult)(nil), "valuescache.ValuesCacheResult")
	proto.RegisterType((*ResourceId)(nil), "valuescache.ResourceId")
}

func init() { proto.RegisterFile("valuescache.proto", fileDescriptor_44747599cffefe15) }

var fileDescriptor_44747599cffefe15 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0xbb, 0x49, 0x9f, 0x3e, 0x3a, 0x01, 0x6d, 0x87, 0x82, 0x4b, 0x29, 0x52, 0x72, 0x2a,
	0x82, 0x39, 0xc4, 0x83, 0xe0, 0x45, 0x68, 0x2f, 0x0a, 0x1e, 0xca, 0x0a, 0x5e, 0x3c, 0xad, 0x9b,
	0x81, 0x04, 0x82, 0x09, 0xfb, 0xd2, 0x8f, 0xe0, 0xd7, 0xf0, 0xab, 0xca, 0xae, 0x15, 0xb7, 0x87,
	0x78, 0x9b, 0xff, 0xec, 0xbc, 0xfc, 0xf6, 0xcf, 0xc0, 0x6c, 0x2f, 0x5b, 0x47, 0x46, 0x49, 0x55,
	0x53, 0xd1, 0xeb, 0xce, 0x76, 0x98, 0x45, 0xa9, 0xfc, 0x15, 0xb2, 0x97, 0x20, 0xb7, 0x5e, 0xe2,
	0x14, 0x52, 0xa7, 0x5b, 0xce, 0x56, 0x6c, 0x7d, 0x2a, 0x7c, 0x88, 0x08, 0xe3, 0x4a, 0x5a, 0xc9,
	0x93, 0x90, 0x0a, 0x31, 0xce, 0xe1, 0x9f, 0xb1, 0x52, 0x5b, 0x9e, 0xae, 0xd8, 0x9a, 0x89, 0x6f,
	0xe1, 0x7b, 0xe9, 0xbd, 0xe2, 0xe3, 0x90, 0xf3, 0x61, 0xfe, 0xc1, 0xe0, 0x3c, 0x9a, 0xfe, 0xd4,
	0x18, 0x8b, 0xb7, 0x00, 0x9a, 0x4c, 0xe7, 0xb4, 0xa2, 0xa6, 0x0a, 0x8b, 0xb2, 0xf2, 0xa2, 0x88,
	0x29, 0xc5, 0xe1, 0xf9, 0xb1, 0x12, 0x51, 0x29, 0xde, 0x41, 0x0c, 0xce, 0x93, 0x55, 0xba, 0xce,
	0x4a, 0x7e, 0xd4, 0x19, 0xed, 0x12, 0x47, 0xbf, 0xbc, 0x86, 0x59, 0xfc, 0x46, 0xc6, 0xb5, 0x16,
	0x39, 0xfc, 0x37, 0x4e, 0x29, 0x32, 0x26, 0x60, 0x9c, 0x88, 0x1f, 0x99, 0x2f, 0x01, 0x7e, 0x21,
	0xf0, 0x0c, 0x92, 0x03, 0x69, 0x2a, 0x92, 0xa6, 0x2a, 0x3f, 0x19, 0x60, 0x34, 0xed, 0x99, 0xf4,
	0xbe, 0x51, 0x84, 0x0f, 0x30, 0xd9, 0x6a, 0x92, 0x96, 0x70, 0x39, 0x04, 0xe5, 0x0d, 0x58, 0x5c,
	0x0e, 0x22, 0x07, 0xac, 0x7c, 0x84, 0xf7, 0x30, 0x0e, 0x56, 0x0d, 0xd9, 0xb2, 0xf8, 0x73, 0x41,
	0x3e, 0xda, 0x5c, 0xc1, 0x9c, 0x5c, 0x5f, 0x6b, 0x69, 0x0b, 0x32, 0x45, 0x6d, 0x6d, 0x1f, 0x2a,
	0x37, 0xd3, 0xa8, 0x74, 0xe7, 0x6f, 0x61, 0xc7, 0xde, 0x26, 0xe1, 0x28, 0x6e, 0xbe, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x6d, 0xaa, 0x96, 0x7a, 0x29, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ValuesCacheServiceClient is the client API for ValuesCacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValuesCacheServiceClient interface {
	Create(ctx context.Context, in *ValuesCacheList, opts ...grpc.CallOption) (*ValuesCacheResult, error)
	List(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*ValuesCacheList, error)
}

type valuesCacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewValuesCacheServiceClient(cc *grpc.ClientConn) ValuesCacheServiceClient {
	return &valuesCacheServiceClient{cc}
}

func (c *valuesCacheServiceClient) Create(ctx context.Context, in *ValuesCacheList, opts ...grpc.CallOption) (*ValuesCacheResult, error) {
	out := new(ValuesCacheResult)
	err := c.cc.Invoke(ctx, "/valuescache.ValuesCacheService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valuesCacheServiceClient) List(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*ValuesCacheList, error) {
	out := new(ValuesCacheList)
	err := c.cc.Invoke(ctx, "/valuescache.ValuesCacheService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValuesCacheServiceServer is the server API for ValuesCacheService service.
type ValuesCacheServiceServer interface {
	Create(context.Context, *ValuesCacheList) (*ValuesCacheResult, error)
	List(context.Context, *ResourceId) (*ValuesCacheList, error)
}

// UnimplementedValuesCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedValuesCacheServiceServer struct {
}

func (*UnimplementedValuesCacheServiceServer) Create(ctx context.Context, req *ValuesCacheList) (*ValuesCacheResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedValuesCacheServiceServer) List(ctx context.Context, req *ResourceId) (*ValuesCacheList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterValuesCacheServiceServer(s *grpc.Server, srv ValuesCacheServiceServer) {
	s.RegisterService(&_ValuesCacheService_serviceDesc, srv)
}

func _ValuesCacheService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValuesCacheList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValuesCacheServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/valuescache.ValuesCacheService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValuesCacheServiceServer).Create(ctx, req.(*ValuesCacheList))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValuesCacheService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValuesCacheServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/valuescache.ValuesCacheService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValuesCacheServiceServer).List(ctx, req.(*ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValuesCacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "valuescache.ValuesCacheService",
	HandlerType: (*ValuesCacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ValuesCacheService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ValuesCacheService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "valuescache.proto",
}
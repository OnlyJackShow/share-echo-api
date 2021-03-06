// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package proto

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

type HelloWorldRequest struct {
	SkuId                []int32  `protobuf:"varint,1,rep,packed,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (m *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(m, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetSkuId() []int32 {
	if m != nil {
		return m.SkuId
	}
	return nil
}

type HelloWorldResponse struct {
	Data                 map[int32]int32 `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (m *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(m, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetData() map[int32]int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "proto.HelloWorldRequest")
	proto.RegisterType((*HelloWorldResponse)(nil), "proto.HelloWorldResponse")
	proto.RegisterMapType((map[int32]int32)(nil), "proto.HelloWorldResponse.DataEntry")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x5a, 0x5c, 0x82, 0x1e,
	0x20, 0xd1, 0xf0, 0xfc, 0xa2, 0x9c, 0x94, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x51,
	0x2e, 0xb6, 0xe2, 0xec, 0xd2, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0xd6,
	0xe2, 0xec, 0x52, 0xcf, 0x14, 0xa5, 0x36, 0x46, 0x2e, 0x21, 0x64, 0xc5, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x42, 0xe6, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x60, 0xb5, 0xdc, 0x46, 0xca, 0x10,
	0xf3, 0xf5, 0x30, 0x15, 0xea, 0xb9, 0x24, 0x96, 0x24, 0xba, 0xe6, 0x95, 0x14, 0x55, 0x06, 0x81,
	0x35, 0x48, 0x99, 0x73, 0x71, 0xc2, 0x85, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54,
	0x09, 0x26, 0xb0, 0x18, 0x84, 0x63, 0xc5, 0x64, 0xc1, 0x68, 0x14, 0xc8, 0xc5, 0x03, 0x36, 0x3e,
	0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x91, 0x8b, 0x0b, 0x61, 0x9d, 0x90, 0x04, 0x16,
	0x17, 0x80, 0xfd, 0x25, 0x25, 0x89, 0xd3, 0x6d, 0x49, 0x6c, 0x60, 0x19, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0x25, 0xe5, 0x1f, 0x24, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	//根据sku_id查找product_id（批量）
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/proto.HelloService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	//根据sku_id查找product_id（批量）
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) HelloWorld(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HelloService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _HelloService_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

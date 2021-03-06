// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lc.proto

package lifterControl

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

// The request message containing the tag for the server to log, not to query.
type LifterControlRequest struct {
	Para                 int64    `protobuf:"varint,1,opt,name=para,proto3" json:"para,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifterControlRequest) Reset()         { *m = LifterControlRequest{} }
func (m *LifterControlRequest) String() string { return proto.CompactTextString(m) }
func (*LifterControlRequest) ProtoMessage()    {}
func (*LifterControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44d23c5f8e671466, []int{0}
}

func (m *LifterControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifterControlRequest.Unmarshal(m, b)
}
func (m *LifterControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifterControlRequest.Marshal(b, m, deterministic)
}
func (m *LifterControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifterControlRequest.Merge(m, src)
}
func (m *LifterControlRequest) XXX_Size() int {
	return xxx_messageInfo_LifterControlRequest.Size(m)
}
func (m *LifterControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LifterControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LifterControlRequest proto.InternalMessageInfo

func (m *LifterControlRequest) GetPara() int64 {
	if m != nil {
		return m.Para
	}
	return 0
}

// The response message containing the last RobotStatus, datetime, and errorMesage if any.
type LifterControlReply struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifterControlReply) Reset()         { *m = LifterControlReply{} }
func (m *LifterControlReply) String() string { return proto.CompactTextString(m) }
func (*LifterControlReply) ProtoMessage()    {}
func (*LifterControlReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_44d23c5f8e671466, []int{1}
}

func (m *LifterControlReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifterControlReply.Unmarshal(m, b)
}
func (m *LifterControlReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifterControlReply.Marshal(b, m, deterministic)
}
func (m *LifterControlReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifterControlReply.Merge(m, src)
}
func (m *LifterControlReply) XXX_Size() int {
	return xxx_messageInfo_LifterControlReply.Size(m)
}
func (m *LifterControlReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LifterControlReply.DiscardUnknown(m)
}

var xxx_messageInfo_LifterControlReply proto.InternalMessageInfo

func (m *LifterControlReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*LifterControlRequest)(nil), "lifterControl.LifterControlRequest")
	proto.RegisterType((*LifterControlReply)(nil), "lifterControl.LifterControlReply")
}

func init() { proto.RegisterFile("lc.proto", fileDescriptor_44d23c5f8e671466) }

var fileDescriptor_44d23c5f8e671466 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x49, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0xc9, 0x4c, 0x2b, 0x49, 0x2d, 0x72, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0x51, 0xd2, 0xe2, 0x12, 0xf1, 0x41, 0x16, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x48, 0x2c, 0x4a, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e,
	0x02, 0xb3, 0x95, 0x2c, 0xb8, 0x84, 0xd0, 0xd4, 0x16, 0xe4, 0x54, 0x0a, 0x29, 0x71, 0xf1, 0xa4,
	0x16, 0x15, 0xe5, 0x17, 0xf9, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x82, 0x75, 0x70, 0x06, 0xa1,
	0x88, 0x19, 0x95, 0xa0, 0xd9, 0x12, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x14, 0xc3, 0x25,
	0x00, 0x15, 0x09, 0xc9, 0x48, 0x85, 0xa8, 0x10, 0x52, 0xd6, 0x43, 0x71, 0xa1, 0x1e, 0x36, 0xe7,
	0x49, 0x29, 0xe2, 0x57, 0x54, 0x90, 0x53, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0xb1, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xdb, 0x91, 0x46, 0xa6, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LifterControlServiceClient is the client API for LifterControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LifterControlServiceClient interface {
	// Sends a greeting
	ControlTheLifter(ctx context.Context, in *LifterControlRequest, opts ...grpc.CallOption) (*LifterControlReply, error)
}

type lifterControlServiceClient struct {
	cc *grpc.ClientConn
}

func NewLifterControlServiceClient(cc *grpc.ClientConn) LifterControlServiceClient {
	return &lifterControlServiceClient{cc}
}

func (c *lifterControlServiceClient) ControlTheLifter(ctx context.Context, in *LifterControlRequest, opts ...grpc.CallOption) (*LifterControlReply, error) {
	out := new(LifterControlReply)
	err := c.cc.Invoke(ctx, "/lifterControl.LifterControlService/ControlTheLifter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifterControlServiceServer is the server API for LifterControlService service.
type LifterControlServiceServer interface {
	// Sends a greeting
	ControlTheLifter(context.Context, *LifterControlRequest) (*LifterControlReply, error)
}

// UnimplementedLifterControlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLifterControlServiceServer struct {
}

func (*UnimplementedLifterControlServiceServer) ControlTheLifter(ctx context.Context, req *LifterControlRequest) (*LifterControlReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControlTheLifter not implemented")
}

func RegisterLifterControlServiceServer(s *grpc.Server, srv LifterControlServiceServer) {
	s.RegisterService(&_LifterControlService_serviceDesc, srv)
}

func _LifterControlService_ControlTheLifter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifterControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifterControlServiceServer).ControlTheLifter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lifterControl.LifterControlService/ControlTheLifter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifterControlServiceServer).ControlTheLifter(ctx, req.(*LifterControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LifterControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lifterControl.LifterControlService",
	HandlerType: (*LifterControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ControlTheLifter",
			Handler:    _LifterControlService_ControlTheLifter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lc.proto",
}

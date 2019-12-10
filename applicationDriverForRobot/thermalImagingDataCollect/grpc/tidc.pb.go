// Code generated by protoc-gen-gitcode_go. DO NOT EDIT.
// source: tidc.proto

package thermalImagingDataCollect

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

// The request message containing a tag for the server to log.
type ThermalImagingDataCollectRequest struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThermalImagingDataCollectRequest) Reset()         { *m = ThermalImagingDataCollectRequest{} }
func (m *ThermalImagingDataCollectRequest) String() string { return proto.CompactTextString(m) }
func (*ThermalImagingDataCollectRequest) ProtoMessage()    {}
func (*ThermalImagingDataCollectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bfead4fa8608567, []int{0}
}

func (m *ThermalImagingDataCollectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermalImagingDataCollectRequest.Unmarshal(m, b)
}
func (m *ThermalImagingDataCollectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermalImagingDataCollectRequest.Marshal(b, m, deterministic)
}
func (m *ThermalImagingDataCollectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermalImagingDataCollectRequest.Merge(m, src)
}
func (m *ThermalImagingDataCollectRequest) XXX_Size() int {
	return xxx_messageInfo_ThermalImagingDataCollectRequest.Size(m)
}
func (m *ThermalImagingDataCollectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermalImagingDataCollectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ThermalImagingDataCollectRequest proto.InternalMessageInfo

func (m *ThermalImagingDataCollectRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// The response message containing the dataArray, and the errorMesage if any.
type ThermalImagingDataCollectReply struct {
	Mdata                []*ModelData `protobuf:"bytes,1,rep,name=mdata,proto3" json:"mdata,omitempty"`
	ErrorMessage         string       `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThermalImagingDataCollectReply) Reset()         { *m = ThermalImagingDataCollectReply{} }
func (m *ThermalImagingDataCollectReply) String() string { return proto.CompactTextString(m) }
func (*ThermalImagingDataCollectReply) ProtoMessage()    {}
func (*ThermalImagingDataCollectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bfead4fa8608567, []int{1}
}

func (m *ThermalImagingDataCollectReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermalImagingDataCollectReply.Unmarshal(m, b)
}
func (m *ThermalImagingDataCollectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermalImagingDataCollectReply.Marshal(b, m, deterministic)
}
func (m *ThermalImagingDataCollectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermalImagingDataCollectReply.Merge(m, src)
}
func (m *ThermalImagingDataCollectReply) XXX_Size() int {
	return xxx_messageInfo_ThermalImagingDataCollectReply.Size(m)
}
func (m *ThermalImagingDataCollectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermalImagingDataCollectReply.DiscardUnknown(m)
}

var xxx_messageInfo_ThermalImagingDataCollectReply proto.InternalMessageInfo

func (m *ThermalImagingDataCollectReply) GetMdata() []*ModelData {
	if m != nil {
		return m.Mdata
	}
	return nil
}

func (m *ThermalImagingDataCollectReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type ModelData struct {
	Id                   int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []float64 `protobuf:"fixed64,2,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ModelData) Reset()         { *m = ModelData{} }
func (m *ModelData) String() string { return proto.CompactTextString(m) }
func (*ModelData) ProtoMessage()    {}
func (*ModelData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bfead4fa8608567, []int{2}
}

func (m *ModelData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelData.Unmarshal(m, b)
}
func (m *ModelData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelData.Marshal(b, m, deterministic)
}
func (m *ModelData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelData.Merge(m, src)
}
func (m *ModelData) XXX_Size() int {
	return xxx_messageInfo_ModelData.Size(m)
}
func (m *ModelData) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelData.DiscardUnknown(m)
}

var xxx_messageInfo_ModelData proto.InternalMessageInfo

func (m *ModelData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ModelData) GetData() []float64 {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ThermalImagingDataCollectRequest)(nil), "thermalImagingDataCollect.ThermalImagingDataCollectRequest")
	proto.RegisterType((*ThermalImagingDataCollectReply)(nil), "thermalImagingDataCollect.ThermalImagingDataCollectReply")
	proto.RegisterType((*ModelData)(nil), "thermalImagingDataCollect.ModelData")
}

func init() { proto.RegisterFile("tidc.proto", fileDescriptor_2bfead4fa8608567) }

var fileDescriptor_2bfead4fa8608567 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0x4c, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2c, 0xc9, 0x48, 0x2d, 0xca, 0x4d, 0xcc, 0xf1,
	0xcc, 0x4d, 0x4c, 0xcf, 0xcc, 0x4b, 0x77, 0x49, 0x2c, 0x49, 0x74, 0xce, 0xcf, 0xc9, 0x49, 0x4d,
	0x2e, 0x51, 0x32, 0xe1, 0x52, 0x08, 0xc1, 0x25, 0x19, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22,
	0x24, 0xc0, 0xc5, 0x5c, 0x92, 0x98, 0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x2a,
	0x35, 0x30, 0x72, 0xc9, 0xe1, 0xd1, 0x56, 0x90, 0x53, 0x29, 0x64, 0xc5, 0xc5, 0x9a, 0x9b, 0x92,
	0x58, 0x92, 0x28, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa2, 0x87, 0xd3, 0x0d, 0x7a, 0xbe,
	0xf9, 0x29, 0xa9, 0x39, 0x20, 0x81, 0x20, 0x88, 0x16, 0x21, 0x25, 0x2e, 0x9e, 0xd4, 0xa2, 0xa2,
	0xfc, 0x22, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0xb0, 0xcd, 0x28, 0x62, 0x4a,
	0xfa, 0x5c, 0x9c, 0x70, 0x7d, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x60, 0x07, 0xb2, 0x06, 0x31,
	0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0x80, 0xed, 0x66, 0x52, 0x60, 0xd6, 0x60, 0x0c, 0x02, 0xb3,
	0x8d, 0x76, 0x31, 0xe2, 0xf1, 0x6a, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0xd0, 0x54, 0x46,
	0x2e, 0x49, 0xa8, 0x10, 0xa6, 0x5a, 0x21, 0x6b, 0x3c, 0x9e, 0x20, 0x14, 0x8a, 0x52, 0x96, 0xe4,
	0x69, 0x2e, 0xc8, 0xa9, 0x54, 0x62, 0x48, 0x62, 0x03, 0x47, 0xa4, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x63, 0x20, 0x05, 0x1c, 0xd6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ThermalImagingDataCollectServiceClient is the client API for ThermalImagingDataCollectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThermalImagingDataCollectServiceClient interface {
	// Using a dataArray to render a ThermalImaging file
	CollectThermalImagingData(ctx context.Context, in *ThermalImagingDataCollectRequest, opts ...grpc.CallOption) (*ThermalImagingDataCollectReply, error)
}

type thermalImagingDataCollectServiceClient struct {
	cc *grpc.ClientConn
}

func NewThermalImagingDataCollectServiceClient(cc *grpc.ClientConn) ThermalImagingDataCollectServiceClient {
	return &thermalImagingDataCollectServiceClient{cc}
}

func (c *thermalImagingDataCollectServiceClient) CollectThermalImagingData(ctx context.Context, in *ThermalImagingDataCollectRequest, opts ...grpc.CallOption) (*ThermalImagingDataCollectReply, error) {
	out := new(ThermalImagingDataCollectReply)
	err := c.cc.Invoke(ctx, "/thermalImagingDataCollect.ThermalImagingDataCollectService/CollectThermalImagingData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThermalImagingDataCollectServiceServer is the server API for ThermalImagingDataCollectService service.
type ThermalImagingDataCollectServiceServer interface {
	// Using a dataArray to render a ThermalImaging file
	CollectThermalImagingData(context.Context, *ThermalImagingDataCollectRequest) (*ThermalImagingDataCollectReply, error)
}

// UnimplementedThermalImagingDataCollectServiceServer can be embedded to have forward compatible implementations.
type UnimplementedThermalImagingDataCollectServiceServer struct {
}

func (*UnimplementedThermalImagingDataCollectServiceServer) CollectThermalImagingData(ctx context.Context, req *ThermalImagingDataCollectRequest) (*ThermalImagingDataCollectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectThermalImagingData not implemented")
}

func RegisterThermalImagingDataCollectServiceServer(s *grpc.Server, srv ThermalImagingDataCollectServiceServer) {
	s.RegisterService(&_ThermalImagingDataCollectService_serviceDesc, srv)
}

func _ThermalImagingDataCollectService_CollectThermalImagingData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThermalImagingDataCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermalImagingDataCollectServiceServer).CollectThermalImagingData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thermalImagingDataCollect.ThermalImagingDataCollectService/CollectThermalImagingData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermalImagingDataCollectServiceServer).CollectThermalImagingData(ctx, req.(*ThermalImagingDataCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThermalImagingDataCollectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thermalImagingDataCollect.ThermalImagingDataCollectService",
	HandlerType: (*ThermalImagingDataCollectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectThermalImagingData",
			Handler:    _ThermalImagingDataCollectService_CollectThermalImagingData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tidc.proto",
}

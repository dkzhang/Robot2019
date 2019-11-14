// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spm.proto

package robotSinglePointMove

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

type SinglePointInfo struct {
	//Mask value to indicate optional value
	// marker   |   location    |    max_continuous_retries |   distance_tolerance  |   theta_tolerance
	//  16      |      8        |               4           |           2           |           1
	InfoMask uint32 `protobuf:"varint,1,opt,name=infoMask,proto3" json:"infoMask,omitempty"`
	//either marker or location
	Marker string `protobuf:"bytes,2,opt,name=marker,proto3" json:"marker,omitempty"`
	//either marker or location
	LocationX     float32 `protobuf:"fixed32,3,opt,name=locationX,proto3" json:"locationX,omitempty"`
	LocationY     float32 `protobuf:"fixed32,4,opt,name=locationY,proto3" json:"locationY,omitempty"`
	LocationTheta float32 `protobuf:"fixed32,5,opt,name=locationTheta,proto3" json:"locationTheta,omitempty"`
	//optional
	MaxContinuousRetries int32    `protobuf:"varint,6,opt,name=max_continuous_retries,json=maxContinuousRetries,proto3" json:"max_continuous_retries,omitempty"`
	DistanceTolerance    float32  `protobuf:"fixed32,7,opt,name=distance_tolerance,json=distanceTolerance,proto3" json:"distance_tolerance,omitempty"`
	ThetaTolerance       float32  `protobuf:"fixed32,8,opt,name=theta_tolerance,json=thetaTolerance,proto3" json:"theta_tolerance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SinglePointInfo) Reset()         { *m = SinglePointInfo{} }
func (m *SinglePointInfo) String() string { return proto.CompactTextString(m) }
func (*SinglePointInfo) ProtoMessage()    {}
func (*SinglePointInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_235e10e8111f85bc, []int{0}
}

func (m *SinglePointInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SinglePointInfo.Unmarshal(m, b)
}
func (m *SinglePointInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SinglePointInfo.Marshal(b, m, deterministic)
}
func (m *SinglePointInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SinglePointInfo.Merge(m, src)
}
func (m *SinglePointInfo) XXX_Size() int {
	return xxx_messageInfo_SinglePointInfo.Size(m)
}
func (m *SinglePointInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SinglePointInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SinglePointInfo proto.InternalMessageInfo

func (m *SinglePointInfo) GetInfoMask() uint32 {
	if m != nil {
		return m.InfoMask
	}
	return 0
}

func (m *SinglePointInfo) GetMarker() string {
	if m != nil {
		return m.Marker
	}
	return ""
}

func (m *SinglePointInfo) GetLocationX() float32 {
	if m != nil {
		return m.LocationX
	}
	return 0
}

func (m *SinglePointInfo) GetLocationY() float32 {
	if m != nil {
		return m.LocationY
	}
	return 0
}

func (m *SinglePointInfo) GetLocationTheta() float32 {
	if m != nil {
		return m.LocationTheta
	}
	return 0
}

func (m *SinglePointInfo) GetMaxContinuousRetries() int32 {
	if m != nil {
		return m.MaxContinuousRetries
	}
	return 0
}

func (m *SinglePointInfo) GetDistanceTolerance() float32 {
	if m != nil {
		return m.DistanceTolerance
	}
	return 0
}

func (m *SinglePointInfo) GetThetaTolerance() float32 {
	if m != nil {
		return m.ThetaTolerance
	}
	return 0
}

type MoveResponse struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Uuid                 uint32   `protobuf:"varint,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	TaskId               uint32   `protobuf:"varint,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveResponse) Reset()         { *m = MoveResponse{} }
func (m *MoveResponse) String() string { return proto.CompactTextString(m) }
func (*MoveResponse) ProtoMessage()    {}
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_235e10e8111f85bc, []int{1}
}

func (m *MoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveResponse.Unmarshal(m, b)
}
func (m *MoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveResponse.Marshal(b, m, deterministic)
}
func (m *MoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveResponse.Merge(m, src)
}
func (m *MoveResponse) XXX_Size() int {
	return xxx_messageInfo_MoveResponse.Size(m)
}
func (m *MoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MoveResponse proto.InternalMessageInfo

func (m *MoveResponse) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *MoveResponse) GetUuid() uint32 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

func (m *MoveResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MoveResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *MoveResponse) GetTaskId() uint32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type MoveAndWaitForArrivalResponse struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveAndWaitForArrivalResponse) Reset()         { *m = MoveAndWaitForArrivalResponse{} }
func (m *MoveAndWaitForArrivalResponse) String() string { return proto.CompactTextString(m) }
func (*MoveAndWaitForArrivalResponse) ProtoMessage()    {}
func (*MoveAndWaitForArrivalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_235e10e8111f85bc, []int{2}
}

func (m *MoveAndWaitForArrivalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveAndWaitForArrivalResponse.Unmarshal(m, b)
}
func (m *MoveAndWaitForArrivalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveAndWaitForArrivalResponse.Marshal(b, m, deterministic)
}
func (m *MoveAndWaitForArrivalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveAndWaitForArrivalResponse.Merge(m, src)
}
func (m *MoveAndWaitForArrivalResponse) XXX_Size() int {
	return xxx_messageInfo_MoveAndWaitForArrivalResponse.Size(m)
}
func (m *MoveAndWaitForArrivalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveAndWaitForArrivalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MoveAndWaitForArrivalResponse proto.InternalMessageInfo

func (m *MoveAndWaitForArrivalResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*SinglePointInfo)(nil), "robotSinglePointMove.SinglePointInfo")
	proto.RegisterType((*MoveResponse)(nil), "robotSinglePointMove.MoveResponse")
	proto.RegisterType((*MoveAndWaitForArrivalResponse)(nil), "robotSinglePointMove.MoveAndWaitForArrivalResponse")
}

func init() { proto.RegisterFile("spm.proto", fileDescriptor_235e10e8111f85bc) }

var fileDescriptor_235e10e8111f85bc = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0x9b, 0x75, 0x9b, 0x6d, 0x0e, 0x8d, 0xc5, 0x43, 0xad, 0x43, 0x51, 0x08, 0x51, 0x31,
	0x37, 0xee, 0x85, 0xf5, 0x05, 0x4a, 0x41, 0xe8, 0xc5, 0x82, 0x4c, 0x0b, 0xda, 0xab, 0x30, 0x4d,
	0xa6, 0x75, 0xd8, 0x64, 0xce, 0x32, 0x33, 0x29, 0xfb, 0x1c, 0xbe, 0x8b, 0x8f, 0xe3, 0xbb, 0x48,
	0x66, 0x37, 0x31, 0x91, 0x5d, 0xf0, 0x2a, 0xf3, 0xff, 0x60, 0xce, 0xe1, 0xc7, 0x04, 0x22, 0xbb,
	0xaa, 0xe7, 0x2b, 0x43, 0x8e, 0xf0, 0xd4, 0xd0, 0x3d, 0xb9, 0x1b, 0xa5, 0x1f, 0x2b, 0xf9, 0x95,
	0x94, 0x76, 0x0b, 0x7a, 0x92, 0xe9, 0xaf, 0x09, 0x9c, 0x0c, 0xbc, 0x6b, 0xfd, 0x40, 0x78, 0x0e,
	0x47, 0x4a, 0x3f, 0xd0, 0x42, 0xd8, 0x25, 0x0b, 0x92, 0x20, 0x8b, 0x79, 0xaf, 0xf1, 0x0c, 0xc2,
	0x5a, 0x98, 0xa5, 0x34, 0x6c, 0x92, 0x04, 0x59, 0xc4, 0xb7, 0x0a, 0x5f, 0x43, 0x54, 0x51, 0x21,
	0x9c, 0x22, 0xfd, 0x9d, 0x3d, 0x4b, 0x82, 0x6c, 0xc2, 0xff, 0x1a, 0xc3, 0xf4, 0x8e, 0x4d, 0xc7,
	0xe9, 0x1d, 0xbe, 0x83, 0xb8, 0x13, 0xb7, 0x3f, 0xa4, 0x13, 0xec, 0xd0, 0x37, 0xc6, 0x26, 0x7e,
	0x86, 0xb3, 0x5a, 0xac, 0xf3, 0x82, 0xb4, 0x53, 0xba, 0xa1, 0xc6, 0xe6, 0x46, 0x3a, 0xa3, 0xa4,
	0x65, 0x61, 0x12, 0x64, 0x87, 0xfc, 0xb4, 0x16, 0xeb, 0xab, 0x3e, 0xe4, 0x9b, 0x0c, 0x3f, 0x02,
	0x96, 0xca, 0x3a, 0xa1, 0x0b, 0x99, 0x3b, 0xaa, 0xa4, 0x69, 0x4f, 0x6c, 0xe6, 0x07, 0xbc, 0xe8,
	0x92, 0xdb, 0x2e, 0xc0, 0x0f, 0x70, 0xe2, 0xda, 0x69, 0x83, 0xee, 0x91, 0xef, 0x3e, 0xf7, 0x76,
	0x5f, 0x4c, 0x7f, 0x06, 0x70, 0xdc, 0x02, 0xe4, 0xd2, 0xae, 0x48, 0x5b, 0x89, 0x0c, 0x66, 0x05,
	0xd5, 0xb5, 0xd0, 0xa5, 0x67, 0x16, 0xf1, 0x4e, 0x22, 0xc2, 0xb4, 0x69, 0x54, 0xe9, 0x81, 0xc5,
	0xdc, 0x9f, 0x5b, 0x8c, 0xd6, 0x09, 0xd7, 0x58, 0xcf, 0x2a, 0xe2, 0x5b, 0x85, 0x6f, 0x21, 0x96,
	0xc6, 0x90, 0xc9, 0x6b, 0x69, 0xad, 0x78, 0x94, 0x1e, 0x56, 0xc4, 0x8f, 0xbd, 0xb9, 0xd8, 0x78,
	0xf8, 0x0a, 0x66, 0x4e, 0xd8, 0x65, 0xae, 0x4a, 0x4f, 0x2a, 0xe6, 0x61, 0x2b, 0xaf, 0xcb, 0xf4,
	0x0a, 0xde, 0xb4, 0x3b, 0x5d, 0xea, 0xf2, 0x9b, 0x50, 0xee, 0x0b, 0x99, 0x4b, 0x63, 0xd4, 0x93,
	0xa8, 0xfa, 0x25, 0x53, 0x18, 0xdd, 0xb4, 0xdd, 0x74, 0xe4, 0x7d, 0xfa, 0x1d, 0x8c, 0x5e, 0x44,
	0x7b, 0x21, 0xde, 0xc0, 0xd4, 0x7f, 0xdf, 0xcf, 0x77, 0x3d, 0xa2, 0xf9, 0x3f, 0x0f, 0xe8, 0x3c,
	0xdd, 0x5d, 0x1b, 0xf2, 0x4a, 0x0f, 0xb0, 0x81, 0x97, 0x3b, 0xb7, 0xfd, 0xdf, 0x29, 0x17, 0xfb,
	0xa7, 0xec, 0x25, 0x90, 0x1e, 0xdc, 0x87, 0xfe, 0x77, 0xb8, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xcb, 0xf8, 0xe6, 0xf9, 0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SinglePointMoveClient is the client API for SinglePointMove service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SinglePointMoveClient interface {
	Move(ctx context.Context, in *SinglePointInfo, opts ...grpc.CallOption) (*MoveResponse, error)
	MoveAndWaitForArrival(ctx context.Context, in *SinglePointInfo, opts ...grpc.CallOption) (*MoveAndWaitForArrivalResponse, error)
}

type singlePointMoveClient struct {
	cc *grpc.ClientConn
}

func NewSinglePointMoveClient(cc *grpc.ClientConn) SinglePointMoveClient {
	return &singlePointMoveClient{cc}
}

func (c *singlePointMoveClient) Move(ctx context.Context, in *SinglePointInfo, opts ...grpc.CallOption) (*MoveResponse, error) {
	out := new(MoveResponse)
	err := c.cc.Invoke(ctx, "/robotSinglePointMove.SinglePointMove/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *singlePointMoveClient) MoveAndWaitForArrival(ctx context.Context, in *SinglePointInfo, opts ...grpc.CallOption) (*MoveAndWaitForArrivalResponse, error) {
	out := new(MoveAndWaitForArrivalResponse)
	err := c.cc.Invoke(ctx, "/robotSinglePointMove.SinglePointMove/MoveAndWaitForArrival", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SinglePointMoveServer is the server API for SinglePointMove service.
type SinglePointMoveServer interface {
	Move(context.Context, *SinglePointInfo) (*MoveResponse, error)
	MoveAndWaitForArrival(context.Context, *SinglePointInfo) (*MoveAndWaitForArrivalResponse, error)
}

// UnimplementedSinglePointMoveServer can be embedded to have forward compatible implementations.
type UnimplementedSinglePointMoveServer struct {
}

func (*UnimplementedSinglePointMoveServer) Move(ctx context.Context, req *SinglePointInfo) (*MoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (*UnimplementedSinglePointMoveServer) MoveAndWaitForArrival(ctx context.Context, req *SinglePointInfo) (*MoveAndWaitForArrivalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveAndWaitForArrival not implemented")
}

func RegisterSinglePointMoveServer(s *grpc.Server, srv SinglePointMoveServer) {
	s.RegisterService(&_SinglePointMove_serviceDesc, srv)
}

func _SinglePointMove_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePointInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinglePointMoveServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robotSinglePointMove.SinglePointMove/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinglePointMoveServer).Move(ctx, req.(*SinglePointInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinglePointMove_MoveAndWaitForArrival_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePointInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinglePointMoveServer).MoveAndWaitForArrival(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robotSinglePointMove.SinglePointMove/MoveAndWaitForArrival",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinglePointMoveServer).MoveAndWaitForArrival(ctx, req.(*SinglePointInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _SinglePointMove_serviceDesc = grpc.ServiceDesc{
	ServiceName: "robotSinglePointMove.SinglePointMove",
	HandlerType: (*SinglePointMoveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Move",
			Handler:    _SinglePointMove_Move_Handler,
		},
		{
			MethodName: "MoveAndWaitForArrival",
			Handler:    _SinglePointMove_MoveAndWaitForArrival_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spm.proto",
}

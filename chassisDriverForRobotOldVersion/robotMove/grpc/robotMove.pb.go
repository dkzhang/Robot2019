// Code generated by protoc-gen-go. DO NOT EDIT.
// source: robotMove.proto

package robotMove

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
	Marker        string  `protobuf:"bytes,2,opt,name=marker,proto3" json:"marker,omitempty"`
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
	return fileDescriptor_41055b421decac63, []int{0}
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

type MultiplePointsInfo struct {
	//Mask value to indicate optional value
	InfoMask uint32 `protobuf:"varint,1,opt,name=infoMask,proto3" json:"infoMask,omitempty"`
	//required
	Marker []string `protobuf:"bytes,2,rep,name=marker,proto3" json:"marker,omitempty"`
	//optional
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	DistanceTolerance    float32  `protobuf:"fixed32,4,opt,name=distance_tolerance,json=distanceTolerance,proto3" json:"distance_tolerance,omitempty"`
	MaxContinuousRetries float32  `protobuf:"fixed32,5,opt,name=max_continuous_retries,json=maxContinuousRetries,proto3" json:"max_continuous_retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplePointsInfo) Reset()         { *m = MultiplePointsInfo{} }
func (m *MultiplePointsInfo) String() string { return proto.CompactTextString(m) }
func (*MultiplePointsInfo) ProtoMessage()    {}
func (*MultiplePointsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_41055b421decac63, []int{1}
}

func (m *MultiplePointsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplePointsInfo.Unmarshal(m, b)
}
func (m *MultiplePointsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplePointsInfo.Marshal(b, m, deterministic)
}
func (m *MultiplePointsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplePointsInfo.Merge(m, src)
}
func (m *MultiplePointsInfo) XXX_Size() int {
	return xxx_messageInfo_MultiplePointsInfo.Size(m)
}
func (m *MultiplePointsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplePointsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplePointsInfo proto.InternalMessageInfo

func (m *MultiplePointsInfo) GetInfoMask() uint32 {
	if m != nil {
		return m.InfoMask
	}
	return 0
}

func (m *MultiplePointsInfo) GetMarker() []string {
	if m != nil {
		return m.Marker
	}
	return nil
}

func (m *MultiplePointsInfo) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MultiplePointsInfo) GetDistanceTolerance() float32 {
	if m != nil {
		return m.DistanceTolerance
	}
	return 0
}

func (m *MultiplePointsInfo) GetMaxContinuousRetries() float32 {
	if m != nil {
		return m.MaxContinuousRetries
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
	return fileDescriptor_41055b421decac63, []int{2}
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

func init() {
	proto.RegisterType((*SinglePointInfo)(nil), "robotMove.SinglePointInfo")
	proto.RegisterType((*MultiplePointsInfo)(nil), "robotMove.MultiplePointsInfo")
	proto.RegisterType((*MoveResponse)(nil), "robotMove.MoveResponse")
}

func init() { proto.RegisterFile("robotMove.proto", fileDescriptor_41055b421decac63) }

var fileDescriptor_41055b421decac63 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x18, 0x24, 0xdd, 0xfe, 0x6c, 0x3e, 0x6d, 0xa8, 0xb0, 0x56, 0xbb, 0x51, 0xe1, 0x50, 0x15, 0x24,
	0x7a, 0x21, 0x42, 0xc0, 0x13, 0x2c, 0x07, 0xd8, 0x43, 0xa5, 0xca, 0x2c, 0x12, 0x7b, 0x8a, 0xbc,
	0x89, 0xb7, 0x58, 0x4d, 0xfc, 0x45, 0xf6, 0x67, 0xd4, 0xe7, 0xe0, 0x5d, 0x78, 0x02, 0x5e, 0x0c,
	0xc5, 0x69, 0xd2, 0x14, 0xd1, 0x03, 0xb7, 0x6f, 0x66, 0x3e, 0xc9, 0xe3, 0xf1, 0x18, 0xa6, 0x06,
	0x1f, 0x90, 0x56, 0xf8, 0x43, 0x26, 0x95, 0x41, 0x42, 0x16, 0x76, 0xc4, 0xe2, 0xd7, 0x00, 0xa6,
	0x5f, 0x94, 0xde, 0x14, 0x72, 0x8d, 0x4a, 0xd3, 0xad, 0x7e, 0x44, 0x36, 0x83, 0x73, 0xa5, 0x1f,
	0x71, 0x25, 0xec, 0x36, 0x0e, 0xe6, 0xc1, 0x32, 0xe2, 0x1d, 0x66, 0x57, 0x30, 0x2e, 0x85, 0xd9,
	0x4a, 0x13, 0x0f, 0xe6, 0xc1, 0x32, 0xe4, 0x7b, 0xc4, 0x5e, 0x40, 0x58, 0x60, 0x26, 0x48, 0xa1,
	0xfe, 0x16, 0x9f, 0xcd, 0x83, 0xe5, 0x80, 0x1f, 0x88, 0xbe, 0x7a, 0x1f, 0x0f, 0x8f, 0xd5, 0x7b,
	0xf6, 0x0a, 0xa2, 0x16, 0xdc, 0x7d, 0x97, 0x24, 0xe2, 0x91, 0xdf, 0x38, 0x26, 0xd9, 0x07, 0xb8,
	0x2a, 0xc5, 0x2e, 0xcd, 0x50, 0x93, 0xd2, 0x0e, 0x9d, 0x4d, 0x8d, 0x24, 0xa3, 0xa4, 0x8d, 0xc7,
	0xf3, 0x60, 0x39, 0xe2, 0x97, 0xa5, 0xd8, 0x7d, 0xec, 0x44, 0xde, 0x68, 0xec, 0x0d, 0xb0, 0x5c,
	0x59, 0x12, 0x3a, 0x93, 0x29, 0x61, 0x21, 0x4d, 0x3d, 0xc5, 0x13, 0x7f, 0xc0, 0xb3, 0x56, 0xb9,
	0x6b, 0x05, 0xf6, 0x1a, 0xa6, 0x54, 0x9f, 0xd6, 0xdb, 0x3d, 0xf7, 0xbb, 0x4f, 0x3d, 0xdd, 0x2d,
	0x2e, 0x7e, 0x07, 0xc0, 0x56, 0xae, 0x20, 0x55, 0xed, 0x93, 0xb3, 0xff, 0x15, 0xdd, 0x59, 0x2f,
	0xba, 0x4b, 0x18, 0x65, 0xe8, 0x34, 0xf9, 0xd8, 0x46, 0xbc, 0x01, 0x27, 0x8c, 0x0f, 0x4f, 0x19,
	0x3f, 0x9d, 0x4e, 0x13, 0xe6, 0x3f, 0xd3, 0x59, 0xfc, 0x0c, 0xe0, 0xa2, 0xae, 0x01, 0x97, 0xb6,
	0x42, 0x6d, 0x25, 0x8b, 0x61, 0x92, 0x61, 0x59, 0x0a, 0x9d, 0x7b, 0xfb, 0x21, 0x6f, 0x21, 0x63,
	0x30, 0x74, 0x4e, 0xe5, 0xfe, 0xd9, 0x23, 0xee, 0xe7, 0xfa, 0x46, 0x96, 0x04, 0x39, 0xeb, 0xad,
	0x87, 0x7c, 0x8f, 0xd8, 0x4b, 0x88, 0xa4, 0x31, 0x68, 0xd2, 0x52, 0x5a, 0x2b, 0x36, 0x8d, 0xed,
	0x90, 0x5f, 0x78, 0x72, 0xd5, 0x70, 0xec, 0x1a, 0x26, 0x24, 0xec, 0x36, 0x55, 0xb9, 0xb7, 0x18,
	0xf1, 0x71, 0x0d, 0x6f, 0xf3, 0x77, 0x5f, 0x21, 0xe4, 0x6d, 0x3f, 0xd9, 0xe7, 0xa3, 0x7a, 0x7a,
	0x6a, 0x96, 0x1c, 0xfa, 0xfc, 0x57, 0x75, 0x67, 0xd7, 0x3d, 0xad, 0x7f, 0xb1, 0xc5, 0x93, 0x9b,
	0xb7, 0xf0, 0x5c, 0x61, 0xb2, 0x31, 0x55, 0x96, 0xc8, 0x9d, 0x28, 0xab, 0x42, 0xda, 0xc4, 0xa0,
	0x23, 0xb9, 0x71, 0x2a, 0x97, 0x37, 0x53, 0x5e, 0xcf, 0x9f, 0xea, 0x79, 0x5d, 0x7f, 0x92, 0x75,
	0xf0, 0x30, 0xf6, 0xbf, 0xe5, 0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xaa, 0xbc, 0x40,
	0x40, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RobotMoveClient is the client API for RobotMove service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RobotMoveClient interface {
	SinglePointMove(ctx context.Context, in *SinglePointInfo, opts ...grpc.CallOption) (*MoveResponse, error)
}

type robotMoveClient struct {
	cc *grpc.ClientConn
}

func NewRobotMoveClient(cc *grpc.ClientConn) RobotMoveClient {
	return &robotMoveClient{cc}
}

func (c *robotMoveClient) SinglePointMove(ctx context.Context, in *SinglePointInfo, opts ...grpc.CallOption) (*MoveResponse, error) {
	out := new(MoveResponse)
	err := c.cc.Invoke(ctx, "/robotMove.RobotMove/SinglePointMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotMoveServer is the server API for RobotMove service.
type RobotMoveServer interface {
	SinglePointMove(context.Context, *SinglePointInfo) (*MoveResponse, error)
}

// UnimplementedRobotMoveServer can be embedded to have forward compatible implementations.
type UnimplementedRobotMoveServer struct {
}

func (*UnimplementedRobotMoveServer) SinglePointMove(ctx context.Context, req *SinglePointInfo) (*MoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SinglePointMove not implemented")
}

func RegisterRobotMoveServer(s *grpc.Server, srv RobotMoveServer) {
	s.RegisterService(&_RobotMove_serviceDesc, srv)
}

func _RobotMove_SinglePointMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePointInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotMoveServer).SinglePointMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robotMove.RobotMove/SinglePointMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotMoveServer).SinglePointMove(ctx, req.(*SinglePointInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _RobotMove_serviceDesc = grpc.ServiceDesc{
	ServiceName: "robotMove.RobotMove",
	HandlerType: (*RobotMoveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SinglePointMove",
			Handler:    _RobotMove_SinglePointMove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robotMove.proto",
}

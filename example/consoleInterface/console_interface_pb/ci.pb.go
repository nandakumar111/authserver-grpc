// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ci.proto

package ci

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

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f36933df5e4240, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Host                 string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	IpAddress            string   `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Port                 string   `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f36933df5e4240, []int{1}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StatusResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StatusResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *StatusResponse) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *StatusResponse) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *StatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "ci.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "ci.StatusResponse")
}

func init() { proto.RegisterFile("ci.proto", fileDescriptor_d1f36933df5e4240) }

var fileDescriptor_d1f36933df5e4240 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4d, 0x4a, 0x04, 0x31,
	0x10, 0x85, 0xe9, 0x71, 0x0c, 0x4e, 0x89, 0x7f, 0xb5, 0x90, 0x20, 0x08, 0x43, 0xaf, 0x5c, 0xb5,
	0xa0, 0x27, 0x90, 0x5e, 0xb9, 0x6d, 0x0f, 0x20, 0x31, 0x5d, 0x62, 0x40, 0x53, 0x31, 0x55, 0x7d,
	0x24, 0xef, 0x29, 0x49, 0x14, 0x7a, 0x76, 0x2f, 0x5f, 0xf1, 0xc1, 0x7b, 0x81, 0x13, 0x1f, 0x86,
	0x94, 0x59, 0x19, 0x37, 0x3e, 0xf4, 0x17, 0x70, 0xf6, 0xa2, 0x4e, 0x17, 0x99, 0xe8, 0x7b, 0x21,
	0xd1, 0xfe, 0xa7, 0x83, 0xf3, 0x7f, 0x22, 0x89, 0xa3, 0x10, 0x22, 0x6c, 0xa3, 0xfb, 0x22, 0xdb,
	0xed, 0xbb, 0xbb, 0xdd, 0x54, 0x33, 0xee, 0xe1, 0x74, 0x26, 0xf1, 0x39, 0x24, 0x0d, 0x1c, 0xed,
	0xa6, 0x9e, 0xd6, 0xa8, 0x58, 0x1f, 0x2c, 0x6a, 0x8f, 0x9a, 0x55, 0x32, 0xde, 0x02, 0x84, 0xf4,
	0xea, 0xe6, 0x39, 0x93, 0x88, 0xdd, 0xd6, 0xcb, 0x2e, 0xa4, 0xa7, 0x06, 0x8a, 0x92, 0x38, 0xab,
	0x3d, 0x6e, 0x4a, 0xc9, 0x78, 0x0d, 0x46, 0x6a, 0x1d, 0x6b, 0x2a, 0xfd, 0x7b, 0x3d, 0x8c, 0x70,
	0x39, 0x72, 0x14, 0xfe, 0xa4, 0xe7, 0xa8, 0x94, 0xdf, 0x9d, 0x27, 0xbc, 0x07, 0xd3, 0xaa, 0xe3,
	0xd5, 0xe0, 0xc3, 0x70, 0x30, 0xec, 0x06, 0xd7, 0xa8, 0x2d, 0x7b, 0x33, 0xf5, 0x23, 0x1e, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x9b, 0x3d, 0xad, 0x14, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsoleInterfaceClient is the client API for ConsoleInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsoleInterfaceClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type consoleInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewConsoleInterfaceClient(cc *grpc.ClientConn) ConsoleInterfaceClient {
	return &consoleInterfaceClient{cc}
}

func (c *consoleInterfaceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/ci.ConsoleInterface/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsoleInterfaceServer is the server API for ConsoleInterface service.
type ConsoleInterfaceServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
}

// UnimplementedConsoleInterfaceServer can be embedded to have forward compatible implementations.
type UnimplementedConsoleInterfaceServer struct {
}

func (*UnimplementedConsoleInterfaceServer) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

func RegisterConsoleInterfaceServer(s *grpc.Server, srv ConsoleInterfaceServer) {
	s.RegisterService(&_ConsoleInterface_serviceDesc, srv)
}

func _ConsoleInterface_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleInterfaceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ConsoleInterface/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleInterfaceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsoleInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.ConsoleInterface",
	HandlerType: (*ConsoleInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _ConsoleInterface_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ci.proto",
}
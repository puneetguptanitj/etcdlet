// Code generated by protoc-gen-go. DO NOT EDIT.
// source: etcdlet.proto

package etcdlet

import proto "github.com/golang/protobuf/proto"
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

type BootstrapSpec struct {
	Name                     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	InitialAdvertisePeerUrls string   `protobuf:"bytes,2,opt,name=initialAdvertisePeerUrls" json:"initialAdvertisePeerUrls,omitempty"`
	ListenPeerUrls           string   `protobuf:"bytes,3,opt,name=listenPeerUrls" json:"listenPeerUrls,omitempty"`
	ListenClientUrls         string   `protobuf:"bytes,4,opt,name=listenClientUrls" json:"listenClientUrls,omitempty"`
	AdvertiseClientUrls      string   `protobuf:"bytes,5,opt,name=advertiseClientUrls" json:"advertiseClientUrls,omitempty"`
	InitialClusterToken      string   `protobuf:"bytes,6,opt,name=initialClusterToken" json:"initialClusterToken,omitempty"`
	InitialCluster           string   `protobuf:"bytes,7,opt,name=initialCluster" json:"initialCluster,omitempty"`
	InitialClusterState      string   `protobuf:"bytes,8,opt,name=initialClusterState" json:"initialClusterState,omitempty"`
	Address                  string   `protobuf:"bytes,9,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *BootstrapSpec) Reset()         { *m = BootstrapSpec{} }
func (m *BootstrapSpec) String() string { return proto.CompactTextString(m) }
func (*BootstrapSpec) ProtoMessage()    {}
func (*BootstrapSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_etcdlet_8b52f6eb23507337, []int{0}
}
func (m *BootstrapSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapSpec.Unmarshal(m, b)
}
func (m *BootstrapSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapSpec.Marshal(b, m, deterministic)
}
func (dst *BootstrapSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapSpec.Merge(dst, src)
}
func (m *BootstrapSpec) XXX_Size() int {
	return xxx_messageInfo_BootstrapSpec.Size(m)
}
func (m *BootstrapSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapSpec proto.InternalMessageInfo

func (m *BootstrapSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BootstrapSpec) GetInitialAdvertisePeerUrls() string {
	if m != nil {
		return m.InitialAdvertisePeerUrls
	}
	return ""
}

func (m *BootstrapSpec) GetListenPeerUrls() string {
	if m != nil {
		return m.ListenPeerUrls
	}
	return ""
}

func (m *BootstrapSpec) GetListenClientUrls() string {
	if m != nil {
		return m.ListenClientUrls
	}
	return ""
}

func (m *BootstrapSpec) GetAdvertiseClientUrls() string {
	if m != nil {
		return m.AdvertiseClientUrls
	}
	return ""
}

func (m *BootstrapSpec) GetInitialClusterToken() string {
	if m != nil {
		return m.InitialClusterToken
	}
	return ""
}

func (m *BootstrapSpec) GetInitialCluster() string {
	if m != nil {
		return m.InitialCluster
	}
	return ""
}

func (m *BootstrapSpec) GetInitialClusterState() string {
	if m != nil {
		return m.InitialClusterState
	}
	return ""
}

func (m *BootstrapSpec) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MemberSpec struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	ClusterName          string   `protobuf:"bytes,3,opt,name=clusterName" json:"clusterName,omitempty"`
	ClusterToken         string   `protobuf:"bytes,4,opt,name=clusterToken" json:"clusterToken,omitempty"`
	OpertionType         string   `protobuf:"bytes,5,opt,name=opertionType" json:"opertionType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberSpec) Reset()         { *m = MemberSpec{} }
func (m *MemberSpec) String() string { return proto.CompactTextString(m) }
func (*MemberSpec) ProtoMessage()    {}
func (*MemberSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_etcdlet_8b52f6eb23507337, []int{1}
}
func (m *MemberSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberSpec.Unmarshal(m, b)
}
func (m *MemberSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberSpec.Marshal(b, m, deterministic)
}
func (dst *MemberSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberSpec.Merge(dst, src)
}
func (m *MemberSpec) XXX_Size() int {
	return xxx_messageInfo_MemberSpec.Size(m)
}
func (m *MemberSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MemberSpec proto.InternalMessageInfo

func (m *MemberSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MemberSpec) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MemberSpec) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *MemberSpec) GetClusterToken() string {
	if m != nil {
		return m.ClusterToken
	}
	return ""
}

func (m *MemberSpec) GetOpertionType() string {
	if m != nil {
		return m.OpertionType
	}
	return ""
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_etcdlet_8b52f6eb23507337, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*BootstrapSpec)(nil), "etcdlet.BootstrapSpec")
	proto.RegisterType((*MemberSpec)(nil), "etcdlet.MemberSpec")
	proto.RegisterType((*Response)(nil), "etcdlet.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Etcdlet service

type EtcdletClient interface {
	Bootstrap(ctx context.Context, in *BootstrapSpec, opts ...grpc.CallOption) (*Response, error)
	Reconfigure(ctx context.Context, in *MemberSpec, opts ...grpc.CallOption) (*Response, error)
}

type etcdletClient struct {
	cc *grpc.ClientConn
}

func NewEtcdletClient(cc *grpc.ClientConn) EtcdletClient {
	return &etcdletClient{cc}
}

func (c *etcdletClient) Bootstrap(ctx context.Context, in *BootstrapSpec, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/etcdlet.etcdlet/bootstrap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *etcdletClient) Reconfigure(ctx context.Context, in *MemberSpec, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/etcdlet.etcdlet/reconfigure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Etcdlet service

type EtcdletServer interface {
	Bootstrap(context.Context, *BootstrapSpec) (*Response, error)
	Reconfigure(context.Context, *MemberSpec) (*Response, error)
}

func RegisterEtcdletServer(s *grpc.Server, srv EtcdletServer) {
	s.RegisterService(&_Etcdlet_serviceDesc, srv)
}

func _Etcdlet_Bootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EtcdletServer).Bootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdlet.etcdlet/Bootstrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EtcdletServer).Bootstrap(ctx, req.(*BootstrapSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Etcdlet_Reconfigure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EtcdletServer).Reconfigure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdlet.etcdlet/Reconfigure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EtcdletServer).Reconfigure(ctx, req.(*MemberSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _Etcdlet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etcdlet.etcdlet",
	HandlerType: (*EtcdletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "bootstrap",
			Handler:    _Etcdlet_Bootstrap_Handler,
		},
		{
			MethodName: "reconfigure",
			Handler:    _Etcdlet_Reconfigure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "etcdlet.proto",
}

func init() { proto.RegisterFile("etcdlet.proto", fileDescriptor_etcdlet_8b52f6eb23507337) }

var fileDescriptor_etcdlet_8b52f6eb23507337 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0xfd, 0xfa, 0xf3, 0xf5, 0x67, 0x6a, 0x45, 0xa7, 0x50, 0x16, 0xaf, 0x4a, 0x2e, 0x44, 0xbc,
	0x28, 0xa2, 0x17, 0x8a, 0x77, 0xda, 0x6b, 0x45, 0xda, 0xfa, 0x00, 0xdb, 0x64, 0x94, 0xc5, 0x74,
	0x37, 0xec, 0x4e, 0x05, 0xc1, 0x37, 0xf0, 0x35, 0x7c, 0x50, 0xc9, 0x26, 0xa9, 0x89, 0x4d, 0xef,
	0x32, 0x67, 0xce, 0x99, 0x99, 0x9c, 0x3d, 0x30, 0x24, 0x0e, 0xa3, 0x98, 0x78, 0x9a, 0x58, 0xc3,
	0x06, 0xbb, 0x79, 0x19, 0x7c, 0xb5, 0x60, 0x78, 0x6f, 0x0c, 0x3b, 0xb6, 0x32, 0x59, 0x24, 0x14,
	0x22, 0x42, 0x5b, 0xcb, 0x35, 0x89, 0xc6, 0xa4, 0x71, 0xd6, 0x9f, 0xfb, 0x6f, 0xbc, 0x05, 0xa1,
	0xb4, 0x62, 0x25, 0xe3, 0xbb, 0xe8, 0x9d, 0x2c, 0x2b, 0x47, 0x4f, 0x44, 0xf6, 0xd9, 0xc6, 0x4e,
	0x34, 0x3d, 0x6f, 0x6f, 0x1f, 0x4f, 0xe1, 0x30, 0x56, 0x8e, 0x49, 0x6f, 0x15, 0x2d, 0xaf, 0xf8,
	0x83, 0xe2, 0x39, 0x1c, 0x65, 0xc8, 0x2c, 0x56, 0xa4, 0xd9, 0x33, 0xdb, 0x9e, 0xb9, 0x83, 0xe3,
	0x05, 0x8c, 0x64, 0xb1, 0xa8, 0x44, 0xff, 0xef, 0xe9, 0x75, 0xad, 0x54, 0x91, 0x5f, 0x38, 0x8b,
	0x37, 0x8e, 0xc9, 0x2e, 0xcd, 0x1b, 0x69, 0xd1, 0xc9, 0x14, 0x35, 0xad, 0xf4, 0xee, 0x2a, 0x2c,
	0xba, 0xd9, 0xdd, 0x55, 0x74, 0x77, 0xf2, 0x82, 0x25, 0x93, 0xe8, 0xd5, 0x4d, 0xf6, 0x2d, 0x14,
	0xd0, 0x95, 0x51, 0x64, 0xc9, 0x39, 0xd1, 0xf7, 0xac, 0xa2, 0x0c, 0xbe, 0x1b, 0x00, 0x0f, 0xb4,
	0x5e, 0x91, 0xdd, 0xfb, 0x14, 0x25, 0x71, 0xb3, 0x22, 0xc6, 0x09, 0x0c, 0xc2, 0x6c, 0xcd, 0x63,
	0x2a, 0xca, 0x5c, 0x2e, 0x43, 0x18, 0xc0, 0x41, 0x58, 0xfe, 0xfb, 0xcc, 0xde, 0x0a, 0x96, 0x72,
	0x4c, 0x92, 0xda, 0x67, 0xf4, 0xf2, 0x23, 0xa1, 0xdc, 0xd3, 0x0a, 0x16, 0x04, 0xd0, 0x9b, 0x93,
	0x4b, 0x8c, 0x76, 0x84, 0x63, 0xe8, 0x38, 0x96, 0xbc, 0x71, 0xf9, 0x95, 0x79, 0x75, 0xf9, 0x09,
	0x45, 0xc6, 0xf0, 0x06, 0xfa, 0xab, 0x22, 0x62, 0x38, 0x9e, 0x16, 0x49, 0xac, 0xc4, 0xee, 0xe4,
	0x78, 0x8b, 0x17, 0xa3, 0x83, 0x7f, 0x78, 0x0d, 0x03, 0x4b, 0xa1, 0xd1, 0x2f, 0xea, 0x75, 0x63,
	0x09, 0x47, 0x5b, 0xce, 0xaf, 0x49, 0xb5, 0xc2, 0x55, 0xc7, 0xc7, 0xfc, 0xea, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xe6, 0xad, 0x62, 0xde, 0xf7, 0x02, 0x00, 0x00,
}

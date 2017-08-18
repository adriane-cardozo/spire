// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_resolver.proto

/*
Package sri_proto is a generated protocol buffer package.

It is generated from these files:
	node_resolver.proto

It has these top-level messages:
	NodeResolution
	NodeResolutionList
	ResolveRequest
	ResolveResponse
	Empty
*/
package sri_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sri_proto1 "github.com/spiffe/sri/control_plane/plugins/common/proto"

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

// ConfigureRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureRequest sri_proto1.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*sri_proto1.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*sri_proto1.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*sri_proto1.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureResponse sri_proto1.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*sri_proto1.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*sri_proto1.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*sri_proto1.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoRequest sri_proto1.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*sri_proto1.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*sri_proto1.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoResponse sri_proto1.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset() { (*sri_proto1.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).String()
}
func (*GetPluginInfoResponse) ProtoMessage() {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetCompany()
}

// *Represents a a type with a selectorType and a selector.
type NodeResolution struct {
	SelectorType string `protobuf:"bytes,1,opt,name=selectorType" json:"selectorType,omitempty"`
	Selector     string `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
}

func (m *NodeResolution) Reset()                    { *m = NodeResolution{} }
func (m *NodeResolution) String() string            { return proto.CompactTextString(m) }
func (*NodeResolution) ProtoMessage()               {}
func (*NodeResolution) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NodeResolution) GetSelectorType() string {
	if m != nil {
		return m.SelectorType
	}
	return ""
}

func (m *NodeResolution) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

// *Represents a type with a list of NodeResolution.
type NodeResolutionList struct {
	List []*NodeResolution `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *NodeResolutionList) Reset()                    { *m = NodeResolutionList{} }
func (m *NodeResolutionList) String() string            { return proto.CompactTextString(m) }
func (*NodeResolutionList) ProtoMessage()               {}
func (*NodeResolutionList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NodeResolutionList) GetList() []*NodeResolution {
	if m != nil {
		return m.List
	}
	return nil
}

// *Represents a request with a list of BaseSPIFFEIDs.
type ResolveRequest struct {
	BaseSpiffeIdList []string `protobuf:"bytes,1,rep,name=baseSpiffeIdList" json:"baseSpiffeIdList,omitempty"`
}

func (m *ResolveRequest) Reset()                    { *m = ResolveRequest{} }
func (m *ResolveRequest) String() string            { return proto.CompactTextString(m) }
func (*ResolveRequest) ProtoMessage()               {}
func (*ResolveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResolveRequest) GetBaseSpiffeIdList() []string {
	if m != nil {
		return m.BaseSpiffeIdList
	}
	return nil
}

// *Represents a response with a map of SPIFFE ID to a list of Noderesolution.
type ResolveResponse struct {
	Map map[string]*NodeResolutionList `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResolveResponse) Reset()                    { *m = ResolveResponse{} }
func (m *ResolveResponse) String() string            { return proto.CompactTextString(m) }
func (*ResolveResponse) ProtoMessage()               {}
func (*ResolveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResolveResponse) GetMap() map[string]*NodeResolutionList {
	if m != nil {
		return m.Map
	}
	return nil
}

// *Represents an empty message
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*NodeResolution)(nil), "sri_proto.NodeResolution")
	proto.RegisterType((*NodeResolutionList)(nil), "sri_proto.NodeResolutionList")
	proto.RegisterType((*ResolveRequest)(nil), "sri_proto.ResolveRequest")
	proto.RegisterType((*ResolveResponse)(nil), "sri_proto.ResolveResponse")
	proto.RegisterType((*Empty)(nil), "sri_proto.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeResolver service

type NodeResolverClient interface {
	// *Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *sri_proto1.ConfigureRequest, opts ...grpc.CallOption) (*sri_proto1.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *sri_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*sri_proto1.GetPluginInfoResponse, error)
	// *Retrieves a list of properties reflecting the current state of a particular node(s).
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
}

type nodeResolverClient struct {
	cc *grpc.ClientConn
}

func NewNodeResolverClient(cc *grpc.ClientConn) NodeResolverClient {
	return &nodeResolverClient{cc}
}

func (c *nodeResolverClient) Configure(ctx context.Context, in *sri_proto1.ConfigureRequest, opts ...grpc.CallOption) (*sri_proto1.ConfigureResponse, error) {
	out := new(sri_proto1.ConfigureResponse)
	err := grpc.Invoke(ctx, "/sri_proto.NodeResolver/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeResolverClient) GetPluginInfo(ctx context.Context, in *sri_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*sri_proto1.GetPluginInfoResponse, error) {
	out := new(sri_proto1.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/sri_proto.NodeResolver/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeResolverClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := grpc.Invoke(ctx, "/sri_proto.NodeResolver/Resolve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeResolver service

type NodeResolverServer interface {
	// *Responsible for configuration of the plugin.
	Configure(context.Context, *sri_proto1.ConfigureRequest) (*sri_proto1.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *sri_proto1.GetPluginInfoRequest) (*sri_proto1.GetPluginInfoResponse, error)
	// *Retrieves a list of properties reflecting the current state of a particular node(s).
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
}

func RegisterNodeResolverServer(s *grpc.Server, srv NodeResolverServer) {
	s.RegisterService(&_NodeResolver_serviceDesc, srv)
}

func _NodeResolver_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sri_proto1.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sri_proto.NodeResolver/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).Configure(ctx, req.(*sri_proto1.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeResolver_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sri_proto1.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sri_proto.NodeResolver/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).GetPluginInfo(ctx, req.(*sri_proto1.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeResolver_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sri_proto.NodeResolver/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeResolver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sri_proto.NodeResolver",
	HandlerType: (*NodeResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeResolver_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeResolver_GetPluginInfo_Handler,
		},
		{
			MethodName: "Resolve",
			Handler:    _NodeResolver_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_resolver.proto",
}

func init() { proto.RegisterFile("node_resolver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x6f, 0xab, 0xd3, 0x30,
	0x14, 0xc6, 0xed, 0x9d, 0xd7, 0x7b, 0x7b, 0xee, 0x9c, 0x23, 0xbe, 0xd9, 0xaa, 0xe2, 0xa8, 0x6f,
	0x86, 0x60, 0x0b, 0x1b, 0x82, 0x88, 0x2f, 0x84, 0x31, 0x65, 0x30, 0x65, 0x44, 0x7d, 0x3d, 0xba,
	0xee, 0x74, 0x06, 0xdb, 0x24, 0x26, 0xe9, 0xa0, 0x9f, 0xc6, 0x2f, 0xe8, 0x87, 0x90, 0xa6, 0x7f,
	0x58, 0xbd, 0xdb, 0xbb, 0x9c, 0x73, 0x7e, 0x79, 0x4e, 0x9e, 0x87, 0xc0, 0x53, 0x2e, 0xf6, 0xb8,
	0x55, 0xa8, 0x45, 0x7a, 0x44, 0x15, 0x48, 0x25, 0x8c, 0x20, 0xae, 0x56, 0x6c, 0x6b, 0x8f, 0xde,
	0xf2, 0xc0, 0xcc, 0xcf, 0x7c, 0x17, 0xc4, 0x22, 0x0b, 0xb5, 0x64, 0x49, 0x82, 0xa1, 0x56, 0x2c,
	0x8c, 0x05, 0x37, 0x4a, 0xa4, 0x5b, 0x99, 0x46, 0x1c, 0x43, 0x99, 0xe6, 0x07, 0xc6, 0x75, 0x18,
	0x8b, 0x2c, 0x13, 0x3c, 0xb4, 0x37, 0xeb, 0xa2, 0x52, 0xf4, 0x37, 0x30, 0xf8, 0x2a, 0xf6, 0x48,
	0xcb, 0x3d, 0xb9, 0x61, 0x82, 0x13, 0x1f, 0xfa, 0x1a, 0x53, 0x8c, 0x8d, 0x50, 0xdf, 0x0b, 0x89,
	0x23, 0x67, 0xe2, 0x4c, 0x5d, 0xda, 0xe9, 0x11, 0x0f, 0x6e, 0x9b, 0x7a, 0x74, 0x65, 0xe7, 0x6d,
	0xed, 0x2f, 0x80, 0x74, 0x15, 0xd7, 0x4c, 0x1b, 0xf2, 0x06, 0x1e, 0xa6, 0x4c, 0x9b, 0x91, 0x33,
	0xe9, 0x4d, 0xef, 0x66, 0xe3, 0xa0, 0x35, 0x12, 0x74, 0x61, 0x6a, 0x31, 0xff, 0x03, 0x0c, 0x68,
	0x65, 0x9d, 0xe2, 0xef, 0x1c, 0xb5, 0x21, 0xaf, 0x61, 0xb8, 0x8b, 0x34, 0x7e, 0xb3, 0x56, 0x57,
	0xfb, 0x75, 0x23, 0xe6, 0xd2, 0x7b, 0x7d, 0xff, 0x8f, 0x03, 0x4f, 0xda, 0xeb, 0x5a, 0x0a, 0xae,
	0x91, 0xbc, 0x85, 0x5e, 0x16, 0xc9, 0x7a, 0xff, 0xab, 0x93, 0xfd, 0xff, 0x81, 0xc1, 0x97, 0x48,
	0x2e, 0xb9, 0x51, 0x05, 0x2d, 0x79, 0xef, 0x07, 0xdc, 0x36, 0x0d, 0x32, 0x84, 0xde, 0x2f, 0x2c,
	0xea, 0x40, 0xca, 0x23, 0x99, 0xc3, 0xf5, 0x31, 0x4a, 0x73, 0xb4, 0x21, 0xdc, 0xcd, 0x5e, 0x5c,
	0xb4, 0x55, 0x3e, 0x8b, 0x56, 0xec, 0xfb, 0xab, 0x77, 0x8e, 0x7f, 0x03, 0xd7, 0xcb, 0x4c, 0x9a,
	0x62, 0xf6, 0xd7, 0x81, 0x7e, 0x8b, 0x1e, 0x51, 0x91, 0x4f, 0xe0, 0x2e, 0x04, 0x4f, 0xd8, 0x21,
	0x57, 0x48, 0x9e, 0x9d, 0x08, 0xb6, 0xdd, 0x3a, 0x11, 0xef, 0xf9, 0xf9, 0x61, 0xed, 0x97, 0xc2,
	0xe3, 0xcf, 0x68, 0x36, 0xf6, 0x03, 0xac, 0x78, 0x22, 0xc8, 0xcb, 0x13, 0xbc, 0x33, 0x69, 0xf4,
	0x26, 0x97, 0x81, 0x5a, 0xf3, 0x23, 0xdc, 0xd4, 0xef, 0x24, 0xe3, 0x73, 0x09, 0x56, 0x3a, 0xde,
	0xe5, 0x70, 0x37, 0x0f, 0x76, 0x8f, 0xec, 0x60, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x96,
	0x08, 0xa4, 0xe0, 0x02, 0x00, 0x00,
}

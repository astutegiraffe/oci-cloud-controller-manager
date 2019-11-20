// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resourceprincipals.proto

package types

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

type SetResourcePrincipalRequest struct {
	ResourcePrincipal *ResourcePrincipal `protobuf:"bytes,1,opt,name=ResourcePrincipal" json:"ResourcePrincipal,omitempty"`
}

func (m *SetResourcePrincipalRequest) Reset()                    { *m = SetResourcePrincipalRequest{} }
func (m *SetResourcePrincipalRequest) String() string            { return proto.CompactTextString(m) }
func (*SetResourcePrincipalRequest) ProtoMessage()               {}
func (*SetResourcePrincipalRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *SetResourcePrincipalRequest) GetResourcePrincipal() *ResourcePrincipal {
	if m != nil {
		return m.ResourcePrincipal
	}
	return nil
}

type ResourcePrincipal struct {
	// Resource Principal Session Token.
	RPST string `protobuf:"bytes,1,opt,name=RPST" json:"RPST,omitempty"`
	// Ephemeral Private Key
	PrivateKey string `protobuf:"bytes,2,opt,name=PrivateKey" json:"PrivateKey,omitempty"`
}

func (m *ResourcePrincipal) Reset()                    { *m = ResourcePrincipal{} }
func (m *ResourcePrincipal) String() string            { return proto.CompactTextString(m) }
func (*ResourcePrincipal) ProtoMessage()               {}
func (*ResourcePrincipal) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *ResourcePrincipal) GetRPST() string {
	if m != nil {
		return m.RPST
	}
	return ""
}

func (m *ResourcePrincipal) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

type SetResourcePrincipalResponse struct {
}

func (m *SetResourcePrincipalResponse) Reset()                    { *m = SetResourcePrincipalResponse{} }
func (m *SetResourcePrincipalResponse) String() string            { return proto.CompactTextString(m) }
func (*SetResourcePrincipalResponse) ProtoMessage()               {}
func (*SetResourcePrincipalResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func init() {
	proto.RegisterType((*SetResourcePrincipalRequest)(nil), "types.SetResourcePrincipalRequest")
	proto.RegisterType((*ResourcePrincipal)(nil), "types.ResourcePrincipal")
	proto.RegisterType((*SetResourcePrincipalResponse)(nil), "types.SetResourcePrincipalResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourcePrincipalRecipient service

type ResourcePrincipalRecipientClient interface {
	SetResourcePrincipal(ctx context.Context, in *SetResourcePrincipalRequest, opts ...grpc.CallOption) (*SetResourcePrincipalResponse, error)
}

type resourcePrincipalRecipientClient struct {
	cc *grpc.ClientConn
}

func NewResourcePrincipalRecipientClient(cc *grpc.ClientConn) ResourcePrincipalRecipientClient {
	return &resourcePrincipalRecipientClient{cc}
}

func (c *resourcePrincipalRecipientClient) SetResourcePrincipal(ctx context.Context, in *SetResourcePrincipalRequest, opts ...grpc.CallOption) (*SetResourcePrincipalResponse, error) {
	out := new(SetResourcePrincipalResponse)
	err := grpc.Invoke(ctx, "/types.ResourcePrincipalRecipient/SetResourcePrincipal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourcePrincipalRecipient service

type ResourcePrincipalRecipientServer interface {
	SetResourcePrincipal(context.Context, *SetResourcePrincipalRequest) (*SetResourcePrincipalResponse, error)
}

func RegisterResourcePrincipalRecipientServer(s *grpc.Server, srv ResourcePrincipalRecipientServer) {
	s.RegisterService(&_ResourcePrincipalRecipient_serviceDesc, srv)
}

func _ResourcePrincipalRecipient_SetResourcePrincipal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetResourcePrincipalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePrincipalRecipientServer).SetResourcePrincipal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.ResourcePrincipalRecipient/SetResourcePrincipal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePrincipalRecipientServer).SetResourcePrincipal(ctx, req.(*SetResourcePrincipalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourcePrincipalRecipient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.ResourcePrincipalRecipient",
	HandlerType: (*ResourcePrincipalRecipientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetResourcePrincipal",
			Handler:    _ResourcePrincipalRecipient_SetResourcePrincipal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resourceprincipals.proto",
}

func init() { proto.RegisterFile("resourceprincipals.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0x39, 0x51, 0xc1, 0xb1, 0x32, 0x58, 0x84, 0x53, 0x0e, 0x89, 0x8d, 0x8d, 0x39, 0xbc,
	0xfb, 0x06, 0x16, 0x5a, 0xd8, 0x2c, 0x39, 0x2b, 0x1b, 0xd9, 0x1d, 0xa6, 0x08, 0x77, 0x9b, 0x89,
	0xc9, 0xac, 0x70, 0x85, 0xdf, 0x5d, 0x8c, 0x27, 0x08, 0x1b, 0xb6, 0x08, 0x84, 0xbc, 0xf7, 0x7e,
	0x99, 0x3f, 0xa0, 0x13, 0x65, 0x1e, 0x12, 0x52, 0x4c, 0x3e, 0xa0, 0x8f, 0xed, 0x2e, 0xdb, 0x98,
	0x58, 0x58, 0x9d, 0xc8, 0x3e, 0x52, 0x36, 0x04, 0x57, 0x1b, 0x12, 0x77, 0x70, 0x35, 0x7f, 0x2e,
	0x47, 0x1f, 0x03, 0x65, 0x51, 0x4f, 0x70, 0x31, 0xd2, 0xf4, 0xec, 0x66, 0x76, 0x77, 0xbe, 0xd2,
	0xb6, 0x10, 0xec, 0x38, 0x3b, 0x8e, 0x98, 0xe7, 0x0a, 0x47, 0x29, 0x38, 0x76, 0xcd, 0xe6, 0xb5,
	0xf0, 0xce, 0x5c, 0xb9, 0xab, 0x05, 0x40, 0x93, 0xfc, 0x67, 0x2b, 0xf4, 0x42, 0x7b, 0x7d, 0x54,
	0x94, 0x7f, 0x2f, 0x66, 0x01, 0xd7, 0xf5, 0x7a, 0x73, 0xe4, 0x90, 0x69, 0xf5, 0x05, 0xf3, 0x8a,
	0x88, 0x3e, 0x7a, 0x0a, 0xa2, 0xde, 0xe1, 0xb2, 0x96, 0x56, 0xe6, 0xd0, 0xcb, 0xc4, 0x28, 0xe6,
	0xb7, 0x93, 0x9e, 0xdf, 0xef, 0x1f, 0xd7, 0x6f, 0x0f, 0x9d, 0x97, 0x6e, 0xc0, 0x2d, 0x89, 0x65,
	0xf4, 0x96, 0x53, 0x8b, 0x3b, 0x42, 0x4e, 0xd1, 0x22, 0xf7, 0x4b, 0xde, 0xd2, 0xcf, 0xb9, 0x47,
	0xee, 0x7b, 0x0e, 0xcb, 0x42, 0xec, 0x4e, 0xcb, 0x46, 0xd6, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0x0d, 0x08, 0x29, 0xad, 0x01, 0x00, 0x00,
}
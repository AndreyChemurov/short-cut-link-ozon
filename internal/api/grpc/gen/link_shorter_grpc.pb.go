// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LinkShorterClient is the client API for LinkShorter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkShorterClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type linkShorterClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkShorterClient(cc grpc.ClientConnInterface) LinkShorterClient {
	return &linkShorterClient{cc}
}

func (c *linkShorterClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/link_shorter.LinkShorter/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkShorterClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/link_shorter.LinkShorter/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkShorterServer is the server API for LinkShorter service.
// All implementations must embed UnimplementedLinkShorterServer
// for forward compatibility
type LinkShorterServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedLinkShorterServer()
}

// UnimplementedLinkShorterServer must be embedded to have forward compatible implementations.
type UnimplementedLinkShorterServer struct {
}

func (UnimplementedLinkShorterServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLinkShorterServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLinkShorterServer) mustEmbedUnimplementedLinkShorterServer() {}

// UnsafeLinkShorterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkShorterServer will
// result in compilation errors.
type UnsafeLinkShorterServer interface {
	mustEmbedUnimplementedLinkShorterServer()
}

func RegisterLinkShorterServer(s grpc.ServiceRegistrar, srv LinkShorterServer) {
	s.RegisterService(&LinkShorter_ServiceDesc, srv)
}

func _LinkShorter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShorterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/link_shorter.LinkShorter/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShorterServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkShorter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShorterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/link_shorter.LinkShorter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShorterServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkShorter_ServiceDesc is the grpc.ServiceDesc for LinkShorter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkShorter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "link_shorter.LinkShorter",
	HandlerType: (*LinkShorterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LinkShorter_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LinkShorter_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link_shorter.proto",
}
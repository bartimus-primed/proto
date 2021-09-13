// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InteractClient is the client API for Interact service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractClient interface {
	RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RunCommandAndRespond(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
	HandsOn(ctx context.Context, opts ...grpc.CallOption) (Interact_HandsOnClient, error)
}

type interactClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractClient(cc grpc.ClientConnInterface) InteractClient {
	return &interactClient{cc}
}

func (c *interactClient) RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Interact/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactClient) RunCommandAndRespond(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.Interact/RunCommandAndRespond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactClient) HandsOn(ctx context.Context, opts ...grpc.CallOption) (Interact_HandsOnClient, error) {
	stream, err := c.cc.NewStream(ctx, &Interact_ServiceDesc.Streams[0], "/pb.Interact/HandsOn", opts...)
	if err != nil {
		return nil, err
	}
	x := &interactHandsOnClient{stream}
	return x, nil
}

type Interact_HandsOnClient interface {
	Send(*Command) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type interactHandsOnClient struct {
	grpc.ClientStream
}

func (x *interactHandsOnClient) Send(m *Command) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interactHandsOnClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InteractServer is the server API for Interact service.
// All implementations must embed UnimplementedInteractServer
// for forward compatibility
type InteractServer interface {
	RunCommand(context.Context, *Command) (*emptypb.Empty, error)
	RunCommandAndRespond(context.Context, *Command) (*Response, error)
	HandsOn(Interact_HandsOnServer) error
	mustEmbedUnimplementedInteractServer()
}

// UnimplementedInteractServer must be embedded to have forward compatible implementations.
type UnimplementedInteractServer struct {
}

func (UnimplementedInteractServer) RunCommand(context.Context, *Command) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (UnimplementedInteractServer) RunCommandAndRespond(context.Context, *Command) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommandAndRespond not implemented")
}
func (UnimplementedInteractServer) HandsOn(Interact_HandsOnServer) error {
	return status.Errorf(codes.Unimplemented, "method HandsOn not implemented")
}
func (UnimplementedInteractServer) mustEmbedUnimplementedInteractServer() {}

// UnsafeInteractServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractServer will
// result in compilation errors.
type UnsafeInteractServer interface {
	mustEmbedUnimplementedInteractServer()
}

func RegisterInteractServer(s grpc.ServiceRegistrar, srv InteractServer) {
	s.RegisterService(&Interact_ServiceDesc, srv)
}

func _Interact_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Interact/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServer).RunCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interact_RunCommandAndRespond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServer).RunCommandAndRespond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Interact/RunCommandAndRespond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServer).RunCommandAndRespond(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interact_HandsOn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InteractServer).HandsOn(&interactHandsOnServer{stream})
}

type Interact_HandsOnServer interface {
	Send(*Response) error
	Recv() (*Command, error)
	grpc.ServerStream
}

type interactHandsOnServer struct {
	grpc.ServerStream
}

func (x *interactHandsOnServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interactHandsOnServer) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Interact_ServiceDesc is the grpc.ServiceDesc for Interact service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interact_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Interact",
	HandlerType: (*InteractServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _Interact_RunCommand_Handler,
		},
		{
			MethodName: "RunCommandAndRespond",
			Handler:    _Interact_RunCommandAndRespond_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HandsOn",
			Handler:       _Interact_HandsOn_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ghost.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: unter.proto

package pb

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

// UnterClient is the client API for Unter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnterClient interface {
	// rpc Health(HealthRequest) return (HealthResponse) {}
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Location(ctx context.Context, opts ...grpc.CallOption) (Unter_LocationClient, error)
}

type unterClient struct {
	cc grpc.ClientConnInterface
}

func NewUnterClient(cc grpc.ClientConnInterface) UnterClient {
	return &unterClient{cc}
}

func (c *unterClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Unter/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unterClient) Location(ctx context.Context, opts ...grpc.CallOption) (Unter_LocationClient, error) {
	stream, err := c.cc.NewStream(ctx, &Unter_ServiceDesc.Streams[0], "/Unter/Location", opts...)
	if err != nil {
		return nil, err
	}
	x := &unterLocationClient{stream}
	return x, nil
}

type Unter_LocationClient interface {
	Send(*LocationRequest) error
	CloseAndRecv() (*LocationResponse, error)
	grpc.ClientStream
}

type unterLocationClient struct {
	grpc.ClientStream
}

func (x *unterLocationClient) Send(m *LocationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *unterLocationClient) CloseAndRecv() (*LocationResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LocationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UnterServer is the server API for Unter service.
// All implementations must embed UnimplementedUnterServer
// for forward compatibility
type UnterServer interface {
	// rpc Health(HealthRequest) return (HealthResponse) {}
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Location(Unter_LocationServer) error
	mustEmbedUnimplementedUnterServer()
}

// UnimplementedUnterServer must be embedded to have forward compatible implementations.
type UnimplementedUnterServer struct {
}

func (UnimplementedUnterServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUnterServer) Location(Unter_LocationServer) error {
	return status.Errorf(codes.Unimplemented, "method Location not implemented")
}
func (UnimplementedUnterServer) mustEmbedUnimplementedUnterServer() {}

// UnsafeUnterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnterServer will
// result in compilation errors.
type UnsafeUnterServer interface {
	mustEmbedUnimplementedUnterServer()
}

func RegisterUnterServer(s grpc.ServiceRegistrar, srv UnterServer) {
	s.RegisterService(&Unter_ServiceDesc, srv)
}

func _Unter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Unter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnterServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Unter_Location_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UnterServer).Location(&unterLocationServer{stream})
}

type Unter_LocationServer interface {
	SendAndClose(*LocationResponse) error
	Recv() (*LocationRequest, error)
	grpc.ServerStream
}

type unterLocationServer struct {
	grpc.ServerStream
}

func (x *unterLocationServer) SendAndClose(m *LocationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *unterLocationServer) Recv() (*LocationRequest, error) {
	m := new(LocationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Unter_ServiceDesc is the grpc.ServiceDesc for Unter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Unter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Unter",
	HandlerType: (*UnterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Unter_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Location",
			Handler:       _Unter_Location_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "unter.proto",
}
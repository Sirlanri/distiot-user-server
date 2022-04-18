// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package devicepb

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

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceServiceClient interface {
	GetDataTypeBydID(ctx context.Context, in *Didmsg, opts ...grpc.CallOption) (*DataTypemsg, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) GetDataTypeBydID(ctx context.Context, in *Didmsg, opts ...grpc.CallOption) (*DataTypemsg, error) {
	out := new(DataTypemsg)
	err := c.cc.Invoke(ctx, "/devicepb.DeviceService/GetDataTypeBydID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations must embed UnimplementedDeviceServiceServer
// for forward compatibility
type DeviceServiceServer interface {
	GetDataTypeBydID(context.Context, *Didmsg) (*DataTypemsg, error)
	mustEmbedUnimplementedDeviceServiceServer()
}

// UnimplementedDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (UnimplementedDeviceServiceServer) GetDataTypeBydID(context.Context, *Didmsg) (*DataTypemsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataTypeBydID not implemented")
}
func (UnimplementedDeviceServiceServer) mustEmbedUnimplementedDeviceServiceServer() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_GetDataTypeBydID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Didmsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDataTypeBydID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devicepb.DeviceService/GetDataTypeBydID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDataTypeBydID(ctx, req.(*Didmsg))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devicepb.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDataTypeBydID",
			Handler:    _DeviceService_GetDataTypeBydID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}

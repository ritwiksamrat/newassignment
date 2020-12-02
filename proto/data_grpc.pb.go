// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// KafkaServiceClient is the client API for KafkaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaServiceClient interface {
	Kservice(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type kafkaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaServiceClient(cc grpc.ClientConnInterface) KafkaServiceClient {
	return &kafkaServiceClient{cc}
}

func (c *kafkaServiceClient) Kservice(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.kafkaService/Kservice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaServiceServer is the server API for KafkaService service.
// All implementations must embed UnimplementedKafkaServiceServer
// for forward compatibility
type KafkaServiceServer interface {
	Kservice(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedKafkaServiceServer()
}

// UnimplementedKafkaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKafkaServiceServer struct {
}

func (UnimplementedKafkaServiceServer) Kservice(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kservice not implemented")
}
func (UnimplementedKafkaServiceServer) mustEmbedUnimplementedKafkaServiceServer() {}

// UnsafeKafkaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaServiceServer will
// result in compilation errors.
type UnsafeKafkaServiceServer interface {
	mustEmbedUnimplementedKafkaServiceServer()
}

func RegisterKafkaServiceServer(s grpc.ServiceRegistrar, srv KafkaServiceServer) {
	s.RegisterService(&KafkaService_ServiceDesc, srv)
}

func _KafkaService_Kservice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServiceServer).Kservice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.kafkaService/Kservice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServiceServer).Kservice(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaService_ServiceDesc is the grpc.ServiceDesc for KafkaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.kafkaService",
	HandlerType: (*KafkaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Kservice",
			Handler:    _KafkaService_Kservice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repository/data.proto",
}

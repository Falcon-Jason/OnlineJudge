// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: judge_result.proto

package judge_result

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

// ResultWriteBackClient is the client API for ResultWriteBack service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResultWriteBackClient interface {
	WriteBack(ctx context.Context, in *ResultWriteBackRequest, opts ...grpc.CallOption) (*ResultWriteBackReply, error)
}

type resultWriteBackClient struct {
	cc grpc.ClientConnInterface
}

func NewResultWriteBackClient(cc grpc.ClientConnInterface) ResultWriteBackClient {
	return &resultWriteBackClient{cc}
}

func (c *resultWriteBackClient) WriteBack(ctx context.Context, in *ResultWriteBackRequest, opts ...grpc.CallOption) (*ResultWriteBackReply, error) {
	out := new(ResultWriteBackReply)
	err := c.cc.Invoke(ctx, "/judge_result.ResultWriteBack/WriteBack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultWriteBackServer is the server API for ResultWriteBack service.
// All implementations must embed UnimplementedResultWriteBackServer
// for forward compatibility
type ResultWriteBackServer interface {
	WriteBack(context.Context, *ResultWriteBackRequest) (*ResultWriteBackReply, error)
	mustEmbedUnimplementedResultWriteBackServer()
}

// UnimplementedResultWriteBackServer must be embedded to have forward compatible implementations.
type UnimplementedResultWriteBackServer struct {
}

func (UnimplementedResultWriteBackServer) WriteBack(context.Context, *ResultWriteBackRequest) (*ResultWriteBackReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteBack not implemented")
}
func (UnimplementedResultWriteBackServer) mustEmbedUnimplementedResultWriteBackServer() {}

// UnsafeResultWriteBackServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResultWriteBackServer will
// result in compilation errors.
type UnsafeResultWriteBackServer interface {
	mustEmbedUnimplementedResultWriteBackServer()
}

func RegisterResultWriteBackServer(s grpc.ServiceRegistrar, srv ResultWriteBackServer) {
	s.RegisterService(&ResultWriteBack_ServiceDesc, srv)
}

func _ResultWriteBack_WriteBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultWriteBackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultWriteBackServer).WriteBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge_result.ResultWriteBack/WriteBack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultWriteBackServer).WriteBack(ctx, req.(*ResultWriteBackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResultWriteBack_ServiceDesc is the grpc.ServiceDesc for ResultWriteBack service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResultWriteBack_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judge_result.ResultWriteBack",
	HandlerType: (*ResultWriteBackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteBack",
			Handler:    _ResultWriteBack_WriteBack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judge_result.proto",
}
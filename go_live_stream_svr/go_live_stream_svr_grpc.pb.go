// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: go_live_stream_svr.proto

// package 格式为服务进程名

package go_live_stream_svr

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

const (
	GoLiveStreamSvr_StartLiveTask_FullMethodName  = "/go_live_stream_svr.go_live_stream_svr/StartLiveTask"
	GoLiveStreamSvr_UpdateLiveTask_FullMethodName = "/go_live_stream_svr.go_live_stream_svr/UpdateLiveTask"
)

// GoLiveStreamSvrClient is the client API for GoLiveStreamSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoLiveStreamSvrClient interface {
	StartLiveTask(ctx context.Context, in *StartLiveTaskReq, opts ...grpc.CallOption) (*StartLiveTaskRsp, error)
	UpdateLiveTask(ctx context.Context, in *UpdateLiveTaskReq, opts ...grpc.CallOption) (*UpdateLiveTaskRsp, error)
}

type goLiveStreamSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewGoLiveStreamSvrClient(cc grpc.ClientConnInterface) GoLiveStreamSvrClient {
	return &goLiveStreamSvrClient{cc}
}

func (c *goLiveStreamSvrClient) StartLiveTask(ctx context.Context, in *StartLiveTaskReq, opts ...grpc.CallOption) (*StartLiveTaskRsp, error) {
	out := new(StartLiveTaskRsp)
	err := c.cc.Invoke(ctx, GoLiveStreamSvr_StartLiveTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goLiveStreamSvrClient) UpdateLiveTask(ctx context.Context, in *UpdateLiveTaskReq, opts ...grpc.CallOption) (*UpdateLiveTaskRsp, error) {
	out := new(UpdateLiveTaskRsp)
	err := c.cc.Invoke(ctx, GoLiveStreamSvr_UpdateLiveTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoLiveStreamSvrServer is the server API for GoLiveStreamSvr service.
// All implementations should embed UnimplementedGoLiveStreamSvrServer
// for forward compatibility
type GoLiveStreamSvrServer interface {
	StartLiveTask(context.Context, *StartLiveTaskReq) (*StartLiveTaskRsp, error)
	UpdateLiveTask(context.Context, *UpdateLiveTaskReq) (*UpdateLiveTaskRsp, error)
}

// UnimplementedGoLiveStreamSvrServer should be embedded to have forward compatible implementations.
type UnimplementedGoLiveStreamSvrServer struct {
}

func (UnimplementedGoLiveStreamSvrServer) StartLiveTask(context.Context, *StartLiveTaskReq) (*StartLiveTaskRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLiveTask not implemented")
}
func (UnimplementedGoLiveStreamSvrServer) UpdateLiveTask(context.Context, *UpdateLiveTaskReq) (*UpdateLiveTaskRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLiveTask not implemented")
}

// UnsafeGoLiveStreamSvrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoLiveStreamSvrServer will
// result in compilation errors.
type UnsafeGoLiveStreamSvrServer interface {
	mustEmbedUnimplementedGoLiveStreamSvrServer()
}

func RegisterGoLiveStreamSvrServer(s grpc.ServiceRegistrar, srv GoLiveStreamSvrServer) {
	s.RegisterService(&GoLiveStreamSvr_ServiceDesc, srv)
}

func _GoLiveStreamSvr_StartLiveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLiveTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoLiveStreamSvrServer).StartLiveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoLiveStreamSvr_StartLiveTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoLiveStreamSvrServer).StartLiveTask(ctx, req.(*StartLiveTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoLiveStreamSvr_UpdateLiveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLiveTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoLiveStreamSvrServer).UpdateLiveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoLiveStreamSvr_UpdateLiveTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoLiveStreamSvrServer).UpdateLiveTask(ctx, req.(*UpdateLiveTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GoLiveStreamSvr_ServiceDesc is the grpc.ServiceDesc for GoLiveStreamSvr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoLiveStreamSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_live_stream_svr.go_live_stream_svr",
	HandlerType: (*GoLiveStreamSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartLiveTask",
			Handler:    _GoLiveStreamSvr_StartLiveTask_Handler,
		},
		{
			MethodName: "UpdateLiveTask",
			Handler:    _GoLiveStreamSvr_UpdateLiveTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_live_stream_svr.proto",
}

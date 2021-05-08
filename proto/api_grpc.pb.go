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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StatisticsCalculatorClient is the client API for StatisticsCalculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticsCalculatorClient interface {
	CalculateStatistics(ctx context.Context, in *CalculateStatisticsRequest, opts ...grpc.CallOption) (*CalculateStatisticsResponse, error)
}

type statisticsCalculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticsCalculatorClient(cc grpc.ClientConnInterface) StatisticsCalculatorClient {
	return &statisticsCalculatorClient{cc}
}

func (c *statisticsCalculatorClient) CalculateStatistics(ctx context.Context, in *CalculateStatisticsRequest, opts ...grpc.CallOption) (*CalculateStatisticsResponse, error) {
	out := new(CalculateStatisticsResponse)
	err := c.cc.Invoke(ctx, "/calcstatistics.StatisticsCalculator/CalculateStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsCalculatorServer is the server API for StatisticsCalculator service.
// All implementations must embed UnimplementedStatisticsCalculatorServer
// for forward compatibility
type StatisticsCalculatorServer interface {
	CalculateStatistics(context.Context, *CalculateStatisticsRequest) (*CalculateStatisticsResponse, error)
	mustEmbedUnimplementedStatisticsCalculatorServer()
}

// UnimplementedStatisticsCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticsCalculatorServer struct {
}

func (UnimplementedStatisticsCalculatorServer) CalculateStatistics(context.Context, *CalculateStatisticsRequest) (*CalculateStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateStatistics not implemented")
}
func (UnimplementedStatisticsCalculatorServer) mustEmbedUnimplementedStatisticsCalculatorServer() {}

// UnsafeStatisticsCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticsCalculatorServer will
// result in compilation errors.
type UnsafeStatisticsCalculatorServer interface {
	mustEmbedUnimplementedStatisticsCalculatorServer()
}

func RegisterStatisticsCalculatorServer(s grpc.ServiceRegistrar, srv StatisticsCalculatorServer) {
	s.RegisterService(&StatisticsCalculator_ServiceDesc, srv)
}

func _StatisticsCalculator_CalculateStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsCalculatorServer).CalculateStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calcstatistics.StatisticsCalculator/CalculateStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsCalculatorServer).CalculateStatistics(ctx, req.(*CalculateStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatisticsCalculator_ServiceDesc is the grpc.ServiceDesc for StatisticsCalculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatisticsCalculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calcstatistics.StatisticsCalculator",
	HandlerType: (*StatisticsCalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateStatistics",
			Handler:    _StatisticsCalculator_CalculateStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package statisticsServiceProto

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

// StatisticsServiceClient is the client API for StatisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticsServiceClient interface {
	GetStatistics(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatisticsResponse, error)
}

type statisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticsServiceClient(cc grpc.ClientConnInterface) StatisticsServiceClient {
	return &statisticsServiceClient{cc}
}

func (c *statisticsServiceClient) GetStatistics(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatisticsResponse, error) {
	out := new(StatisticsResponse)
	err := c.cc.Invoke(ctx, "/statisticsServiceProto.statistics_service/GetStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsServiceServer is the server API for StatisticsService service.
// All implementations must embed UnimplementedStatisticsServiceServer
// for forward compatibility
type StatisticsServiceServer interface {
	GetStatistics(context.Context, *EmptyRequest) (*StatisticsResponse, error)
	mustEmbedUnimplementedStatisticsServiceServer()
}

// UnimplementedStatisticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticsServiceServer struct {
}

func (UnimplementedStatisticsServiceServer) GetStatistics(context.Context, *EmptyRequest) (*StatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedStatisticsServiceServer) mustEmbedUnimplementedStatisticsServiceServer() {}

// UnsafeStatisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticsServiceServer will
// result in compilation errors.
type UnsafeStatisticsServiceServer interface {
	mustEmbedUnimplementedStatisticsServiceServer()
}

func RegisterStatisticsServiceServer(s grpc.ServiceRegistrar, srv StatisticsServiceServer) {
	s.RegisterService(&StatisticsService_ServiceDesc, srv)
}

func _StatisticsService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statisticsServiceProto.statistics_service/GetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServiceServer).GetStatistics(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatisticsService_ServiceDesc is the grpc.ServiceDesc for StatisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statisticsServiceProto.statistics_service",
	HandlerType: (*StatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatistics",
			Handler:    _StatisticsService_GetStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statistics_service/statistics_service.proto",
}
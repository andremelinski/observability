// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/weather.proto

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

const (
	WeatherService_GetLocationTemperature_FullMethodName = "/pb.WeatherService/GetLocationTemperature"
)

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	GetLocationTemperature(ctx context.Context, opts ...grpc.CallOption) (WeatherService_GetLocationTemperatureClient, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetLocationTemperature(ctx context.Context, opts ...grpc.CallOption) (WeatherService_GetLocationTemperatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[0], WeatherService_GetLocationTemperature_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceGetLocationTemperatureClient{stream}
	return x, nil
}

type WeatherService_GetLocationTemperatureClient interface {
	Send(*WeatherLocationRequest) error
	Recv() (*CreateOrderResponse, error)
	grpc.ClientStream
}

type weatherServiceGetLocationTemperatureClient struct {
	grpc.ClientStream
}

func (x *weatherServiceGetLocationTemperatureClient) Send(m *WeatherLocationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *weatherServiceGetLocationTemperatureClient) Recv() (*CreateOrderResponse, error) {
	m := new(CreateOrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	GetLocationTemperature(WeatherService_GetLocationTemperatureServer) error
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) GetLocationTemperature(WeatherService_GetLocationTemperatureServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLocationTemperature not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_GetLocationTemperature_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WeatherServiceServer).GetLocationTemperature(&weatherServiceGetLocationTemperatureServer{stream})
}

type WeatherService_GetLocationTemperatureServer interface {
	Send(*CreateOrderResponse) error
	Recv() (*WeatherLocationRequest, error)
	grpc.ServerStream
}

type weatherServiceGetLocationTemperatureServer struct {
	grpc.ServerStream
}

func (x *weatherServiceGetLocationTemperatureServer) Send(m *CreateOrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *weatherServiceGetLocationTemperatureServer) Recv() (*WeatherLocationRequest, error) {
	m := new(WeatherLocationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLocationTemperature",
			Handler:       _WeatherService_GetLocationTemperature_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/weather.proto",
}

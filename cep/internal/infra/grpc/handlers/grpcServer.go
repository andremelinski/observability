package handlers

import (
	"context"
	"fmt"

	"github.com/andremelinski/observability/cep/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcServer struct{
	grpcServerName string
	grpcPort int
}

func NewGrpcServer(grpc_server_name string, grpc_port int)*GrpcServer{
	return &GrpcServer{
		grpc_server_name,
		grpc_port,
	}
}

func (g *GrpcServer) WeatherBidirectStream() pb.WeatherService_GetLocationTemperatureClient {

	client := g.StartGrpcWeatherClient()
	stream, err := client.GetLocationTemperature(context.Background())

	if err != nil {
		panic(err)
	}
	return stream
}

func (g *GrpcServer) StartGrpcWeatherClient() pb.WeatherServiceClient {

	grpcConn := g.grpcConn()

	return pb.NewWeatherServiceClient(grpcConn)
}

func (g *GrpcServer) CloseGrpcWeatherClient() error {

	grpcConn := g.grpcConn()

	return grpcConn.Close()
}

func (g *GrpcServer) grpcConn() *grpc.ClientConn {
	addr := fmt.Sprintf("%s:%d", g.grpcServerName, g.grpcPort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
	  panic(err)
	}
	return conn
}
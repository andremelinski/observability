package handlers

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/andremelinski/observability/cep/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcServer struct{
	addr *string
}

func NewGrpcServer(grpc_server_name string, grpc_port int)*GrpcServer{
	connectString := fmt.Sprintf("%s:%d", grpc_server_name, grpc_port)
	return &GrpcServer{
		flag.String("addr", connectString, "the address to connect to"),
	}
}

func (g *GrpcServer) WeatherBidirectStream(ctx context.Context) pb.WeatherService_GetLocationTemperatureClient {
	client := g.StartGrpcWeatherClient()
	stream, err := client.GetLocationTemperature(ctx)

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
	conn, err := grpc.NewClient(*g.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	  panic(err)
	}

	log.Println("conectou")
	return conn
}
package handlers

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct{
	grpcServer *grpc.Server
	grpc_port int
}

func NewGrpcServer(grpcServer *grpc.Server, grpc_port int)*GrpcServer{
	return &GrpcServer{
		grpcServer,
		grpc_port,
	}
}

func (g *GrpcServer) Start() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.grpc_port))
	if err != nil {
		panic(err)
	}
	g.grpcServer.Serve(lis)
}
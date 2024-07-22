package handlers

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(g.grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.grpc_port))
	if err != nil {
		panic(err)
	}
 	log.Println("Server is running on port :50051")
    
	if err := g.grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
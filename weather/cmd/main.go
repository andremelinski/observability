package main

import (
	"github.com/andremelinski/observability/weather/configs"
	"github.com/andremelinski/observability/weather/internal/infra/grpc/handlers"
	"github.com/andremelinski/observability/weather/internal/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	handlerExternalApi := utils.NewHandlerExternalApi()
	
	handlers.NewGrpcResgisters(grpcServer, configs.WEATHER_API_KEY, handlerExternalApi)
	handlers.NewGrpcServer(grpcServer, configs.GRPC_PORT).Start()
}
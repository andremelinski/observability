package main

import (
	"github.com/andremelinski/observability/weather/configs"
	"github.com/andremelinski/observability/weather/internal/infra/grpc/handlers"
	"github.com/andremelinski/observability/weather/internal/pkg/utils"
	"google.golang.org/grpc"
)


func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	handlerExternalApi := utils.NewHandlerExternalApi()
	
	handlers.NewGrpcResgisters(grpcServer, configs.WEATHER_API_KEY, handlerExternalApi).CreateGrpcWeatherRegisters()
	handlers.NewGrpcServer(grpcServer, configs.GRPC_PORT).Start()
}
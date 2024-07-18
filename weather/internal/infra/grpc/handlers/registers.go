package handlers

import (
	"github.com/andremelinski/observability/weather/internal/infra/grpc/pb"
	"github.com/andremelinski/observability/weather/internal/infra/grpc/service"
	utils_interface "github.com/andremelinski/observability/weather/internal/pkg/utils/interface"
	utils "github.com/andremelinski/observability/weather/internal/pkg/utils/weather"
	"google.golang.org/grpc"
)

type GrpcResgisterService struct{
	grpcServer *grpc.Server
	weatherApiKey string
	callExternalapi utils_interface.IHandlerExternalApi
}

func NewGrpcResgisters(grpcServer *grpc.Server, weatherApiKey string, callExternalapi utils_interface.IHandlerExternalApi)*GrpcResgisterService{
	return &GrpcResgisterService{
		grpcServer,
		weatherApiKey,
		callExternalapi,
	}
}

func(gs *GrpcResgisterService) CreateGrpcWeatherRegisters(){
	
	a := utils.NewWeatherInfo(gs.weatherApiKey, gs.callExternalapi)
	weatherService := service.NewWeatherService(a)
	pb.RegisterWeatherServiceServer(gs.grpcServer, weatherService)
}
package grpc_interfaces

import (
	"github.com/andremelinski/observability/cep/internal/infra/grpc/pb"
)

type TempResponseDTO struct{
	Temp_C float64
	Temp_F float64
	Temp_K float64
}

type IGrpcHandler interface{
	WeatherBidirectStream() pb.WeatherService_GetLocationTemperatureClient
	StartGrpcWeatherClient() pb.WeatherServiceClient
	CloseGrpcWeatherClient() error
}


type IGrpcClimateInfo interface{ 
	GetLocationTemperature(location string) (*TempResponseDTO, error)
}
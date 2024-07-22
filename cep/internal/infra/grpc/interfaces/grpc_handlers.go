package grpc_interfaces

import (
	"context"

	"github.com/andremelinski/observability/cep/internal/infra/grpc/pb"
)

type TempResponseDTO struct{
	Temp_C float64
	Temp_F float64
	Temp_K float64
}

type IGrpcHandler interface{
	WeatherBidirectStream(ctx context.Context) pb.WeatherService_GetLocationTemperatureClient
	StartGrpcWeatherClient() pb.WeatherServiceClient
	CloseGrpcWeatherClient() error
}


type IGrpcClimateInfo interface{ 
	GetLocationTemperature(location string) (*TempResponseDTO, error)
}
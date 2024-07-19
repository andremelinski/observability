package grpc_client

import (
	"fmt"
	"log"

	grpc_interfaces "github.com/andremelinski/observability/cep/internal/infra/grpc/interfaces"
	"github.com/andremelinski/observability/cep/internal/infra/grpc/pb"
)

type WeatherService struct{
	pb.WeatherService_GetLocationTemperatureClient
	closeConn grpc_interfaces.IGrpcCloseWeatherConn
}

func NewWeatherService() *WeatherService{
	return &WeatherService{
	}
}

func(ws *WeatherService) GetLocationTemperature(location string) (*grpc_interfaces.TempResponseDTO, error){
	grpcClient := ws.WeatherService_GetLocationTemperatureClient
	
	grpcClient.Send(&pb.WeatherLocationRequest{Place: location})
	res, err := grpcClient.Recv()

	if err != nil {
		log.Fatalln("CloseAndRecv stream",err)
	}
	fmt.Println(res)
	
	defer ws.closeConn.CloseGrpcWeatherClient()

	return &grpc_interfaces.TempResponseDTO{
		Temp_C: float64(res.Temp_C),
		Temp_F: float64(res.Temp_F),
		Temp_K: float64(res.Temp_C),
	}, nil
}
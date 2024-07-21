package service

import (
	"context"

	"github.com/andremelinski/observability/weather/internal/infra/grpc/pb"
	utils_interface "github.com/andremelinski/observability/weather/internal/pkg/utils/interface"
)

type WeatherService struct{
	pb.UnimplementedWeatherServiceServer
	weatherInfo utils_interface.IClimateInfoAPI
}


func NewWeatherService(weatherInfo utils_interface.IClimateInfoAPI) *WeatherService{
	return &WeatherService{
		weatherInfo: weatherInfo,
	}
}

func(ws *WeatherService) GGetLocationTemperature(cx context.Context, in *pb.WeatherLocationRequest) (*pb.WeatherLocationResponse, error){
	
	weatherInfo, err := ws.weatherInfo.GetWeatherInfo(in.Place)
	if err != nil {
		return nil, err
	}

	return &pb.WeatherLocationResponse{
		Temp_C: float32(weatherInfo.Current.TempC),
		Temp_F: float32(weatherInfo.Current.TempF),
		Temp_K: float32(weatherInfo.Current.TempC + 273),
	}, nil
}
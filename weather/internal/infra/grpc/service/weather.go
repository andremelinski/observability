package service

import (
	"io"

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

func(ws *WeatherService) GetLocationTemperature(stream pb.WeatherService_GetLocationTemperatureServer) error {
	location, err := stream.Recv()

	if err == io.EOF{
		return nil
	}
	if err != nil {
		return err
	}
	weatherInfo, err := ws.weatherInfo.GetWeatherInfo(location.Place)
	if err != nil {
		return err
	}
	k := weatherInfo.Current.TempC + 273
	err = stream.Send(&pb.WeatherLocationResponse{
		Temp_C: float32(weatherInfo.Current.TempC),
		Temp_F: float32(weatherInfo.Current.TempF),
		Temp_K: float32(k),
	})

	if err != nil {
		return err
	}
	return nil 
}
package grpc_client

import (
	"context"
	"io"
	"time"

	grpc_interfaces "github.com/andremelinski/observability/cep/internal/infra/grpc/interfaces"
	"github.com/andremelinski/observability/cep/internal/infra/grpc/pb"
)

type WeatherService struct{
	grpcHandler grpc_interfaces.IGrpcHandler
}

func NewWeatherService(
	grpcHandler grpc_interfaces.IGrpcHandler,
	) *WeatherService{
	return &WeatherService{
		grpcHandler,
	}
}

func(ws *WeatherService) GetLocationTemperature(location string) (*grpc_interfaces.TempResponseDTO, error){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream := ws.grpcHandler.WeatherBidirectStream(ctx)

    defer ws.grpcHandler.CloseGrpcWeatherClient()

    // stream, err := ws.grpcHandler.WeatherBidirectStream(ctx)
    // if err != nil {
    //     return nil, err
    // }

    req := &pb.WeatherLocationRequest{Place: location}
    if err := stream.Send(req); err != nil {
        return nil, err
    }

    response := &grpc_interfaces.TempResponseDTO{}
    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }

		response = &grpc_interfaces.TempResponseDTO{
            Temp_C: float64(resp.Temp_C),
            Temp_F: float64(resp.Temp_F),
            Temp_K: float64(resp.Temp_K),
        }
    }

    return response, nil
}

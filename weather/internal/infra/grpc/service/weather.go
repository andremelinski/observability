package service

import (
	"context"
	"io"

	"github.com/andremelinski/observability/weather/internal/infra/grpc/pb"
	"github.com/andremelinski/observability/weather/internal/infra/opentelemetry"
	utils_interface "github.com/andremelinski/observability/weather/internal/pkg/utils/interface"
	"go.opentelemetry.io/otel/trace"
)

type WeatherService struct {
	pb.UnimplementedWeatherServiceServer
	weatherInfo      utils_interface.IClimateInfoAPI
	otelTrace        trace.Tracer
	OtelTraceHandler opentelemetry.IHandlerTrace
}

func NewWeatherService(weatherInfo utils_interface.IClimateInfoAPI, otelTrace trace.Tracer, otelInfo opentelemetry.IHandlerTrace) *WeatherService {
	return &WeatherService{
		weatherInfo:      weatherInfo,
		otelTrace:        otelTrace,
		OtelTraceHandler: otelInfo,
	}
}

func (ws *WeatherService) GetLocationTemperature(stream pb.WeatherService_GetLocationTemperatureServer) error {
	ctx := context.Background()

	ctx, span := ws.OtelTraceHandler.StartOTELTrace(ctx, ws.otelTrace, "weatherTraceApi")
	defer span.End() // span acaba quando toda req acabar para ter o trace

	location, err := stream.Recv()

	if err != nil {

		if err == io.EOF {
			return nil
		}

		return err
	}

	weatherInfo, err := ws.weatherInfo.GetWeatherInfo(ctx, location.Place)
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

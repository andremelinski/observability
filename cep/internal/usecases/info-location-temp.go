package usecases

import (
	"context"
	"errors"

	"github.com/andremelinski/observability/cep/internal/infra/opentelemetry"
	"github.com/andremelinski/observability/cep/internal/repository"
	"go.opentelemetry.io/otel/trace"
)

type ClimateLocationInfoUseCaseDTO struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type IClimateLocationInfoUseCase interface {
	GetCityTemp(ctx context.Context, name string) (*ClimateLocationInfoUseCaseDTO, error)
}

type TemperatureRepository struct {
	LocInfo          repository.ILocationInfo
	WeatherInfo      repository.IWeatherInfo
	OtelTrace        trace.Tracer
	OtelTraceHandler opentelemetry.IHandlerTrace
}

func NewClimateLocationInfoUseCase(locInfo repository.ILocationInfo, weatherInfo repository.IWeatherInfo, otelTrace trace.Tracer, otelTraceHandler opentelemetry.IHandlerTrace) *TemperatureRepository {
	return &TemperatureRepository{
		locInfo,
		weatherInfo,
		otelTrace,
		otelTraceHandler,
	}
}

func (l *TemperatureRepository) GetCityTemp(ctx context.Context, name string) (*ClimateLocationInfoUseCaseDTO, error) {
	ctx, span := l.OtelTraceHandler.StartOTELTrace(ctx, l.OtelTrace, "via-cep-trace")
	cityInfo, err := l.LocInfo.GetLocationInfo(ctx, name)

	if err != nil {
		return nil, err
	}

	if cityInfo.Localidade == "" {
		return nil, errors.New("city not found")
	}

	span.End()

	ctx2, span2 := l.OtelTraceHandler.StartOTELTrace(ctx, l.OtelTrace, "weather-api-trace")
	weatherInfo, err := l.WeatherInfo.GetTempByPlaceName(ctx2, cityInfo.Localidade)

	if err != nil {
		return nil, err
	}
	span2.End()
	return &ClimateLocationInfoUseCaseDTO{
		City:       cityInfo.Localidade,
		Celsius:    weatherInfo.Celsius,
		Fahrenheit: weatherInfo.Fahrenheit,
		Kelvin:     weatherInfo.Kelvin,
	}, nil
}

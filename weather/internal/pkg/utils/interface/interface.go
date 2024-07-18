package utils_interface

import (
	"context"

	utils_dto "github.com/andremelinski/observability/weather/internal/pkg/utils/dto"
)

type IHandlerExternalApi interface {
	CallExternalApi(ctx context.Context, timeoutMs int, method string, url string) ([]byte, error)
}

type IClimateInfoAPI interface{
	GetWeatherInfo(place string) (*utils_dto.WeatherApiDTO, error)
}

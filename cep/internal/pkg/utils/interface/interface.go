package utils_interface

import (
	"context"

	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
)

type IHandlerExternalApi interface {
	CallExternalApi(ctx context.Context, timeoutMs int, method string, url string) ([]byte, error)
}

type IClimateInfoAPI interface {
	GetWeatherInfo(ctx context.Context, place string) (*utils_dto.WeatherApiDTO, error)
}

type ICepInfoAPI interface {
	GetCEPInfo(ctx context.Context, cep string) (*utils_dto.ViaCepDTO, error)
}

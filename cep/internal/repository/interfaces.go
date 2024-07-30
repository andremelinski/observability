package repository

import (
	"context"

	cep_repository "github.com/andremelinski/observability/cep/internal/repository/cep"
	temperature_repository "github.com/andremelinski/observability/cep/internal/repository/temperature"
)

type ILocationInfo interface {
	GetLocationInfo(ctx context.Context, cep string) (*cep_repository.LocationOutputDTO, error)
}

type IWeatherInfo interface {
	GetTempByPlaceName(ctx context.Context, name string) (*temperature_repository.TempDTO, error)
}

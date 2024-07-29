package repository

import (
	cep_repository "github.com/andremelinski/observability/cep/internal/repository/cep"
	temperature_repository "github.com/andremelinski/observability/cep/internal/repository/temperature"
)

type ILocationInfo interface {
	GetLocationInfo(cep string) (*cep_repository.LocationOutputDTO, error)
}

type IWeatherInfo interface {
	GetTempByPlaceName(name string) (*temperature_repository.TempDTO, error)
}

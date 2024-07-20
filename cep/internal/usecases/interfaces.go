package usecases

import (
	usecases_dto "github.com/andremelinski/observability/cep/internal/usecases/dto"
)

type ILocationInfo interface {
	GetLocationInfo(cep string) (*usecases_dto.LocationOutputDTO, error)
}

type IWeatherInfo interface{
	GetTempByPlaceName(name string) (*usecases_dto.TempDTOOutput, error)
}
package usecases

import (
	"context"

	usecases_dto "github.com/andremelinski/observability/cep/internal/usecases/dto"
)

type ILocationInfo interface {
	GetLocationInfo(ctx context.Context, cep string) (*usecases_dto.LocationOutputDTO, error)
}

type IWeatherInfo interface {
	GetTempByPlaceName(ctx context.Context, name string) (*usecases_dto.TempDTOOutput, error)
}

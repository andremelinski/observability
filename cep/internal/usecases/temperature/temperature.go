package usecases_temperature

import (
	"context"
	"log"

	grpc_interfaces "github.com/andremelinski/observability/cep/internal/infra/grpc/interfaces"
	usecases_dto "github.com/andremelinski/observability/cep/internal/usecases/dto"
)

// retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
type TemperatureUseCase struct {
	WeatheInfo grpc_interfaces.IGrpcClimateInfo
}

func NewClimateUseCase(climateApi grpc_interfaces.IGrpcClimateInfo) *TemperatureUseCase {
	return &TemperatureUseCase{
		climateApi,
	}
}

func (l *TemperatureUseCase) GetTempByPlaceName(ctx context.Context, name string) (*usecases_dto.TempDTOOutput, error) {
	weatherInfo, err := l.WeatheInfo.GetLocationTemperature(ctx, name)
	log.Println(weatherInfo)
	if err != nil {
		return nil, err
	}

	return &usecases_dto.TempDTOOutput{
		Celsius:    weatherInfo.Temp_C,
		Fahrenheit: weatherInfo.Temp_F,
		Kelvin:     weatherInfo.Temp_K,
	}, nil
}

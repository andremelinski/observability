package usecases

import grpc_interfaces "github.com/andremelinski/observability/cep/internal/infra/grpc/interfaces"

type TempDTOOutput struct{
	Celsius float64 `json:"temp_C"`;
	Fahrenheit float64 `json:"temp_F"`;
	Kelvin float64 `json:"temp_K"`;
}

// retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin. 
type TemperatureUseCase struct {
	WeatheInfo grpc_interfaces.IGrpcClimateInfo
}

func NewClimateUseCase(climateApi grpc_interfaces.IGrpcClimateInfo)*TemperatureUseCase{
	return &TemperatureUseCase{
		climateApi,
	}
}

func (l *TemperatureUseCase)GetTempByPlaceName(name string) (*TempDTOOutput, error){
	weatherInfo, err := l.WeatheInfo.GetLocationTemperature(name)

	if err != nil {
		return nil, err
	}

	return &TempDTOOutput{
		Celsius: weatherInfo.Temp_C,
		Fahrenheit: weatherInfo.Temp_F,
		Kelvin: weatherInfo.Temp_K,
	}, nil
}
package usecases

import utils_interface "github.com/andremelinski/observability/cep/internal/pkg/utils/interface"

type TempDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

// retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
type TemperatureUseCase struct {
	WeatheInfo utils_interface.IClimateInfoAPI
}

func NewClimateUseCase(climateApi utils_interface.IClimateInfoAPI) *TemperatureUseCase {
	return &TemperatureUseCase{
		climateApi,
	}
}

func (l *TemperatureUseCase) GetTempByPlaceName(name string) (*TempDTO, error) {
	weatherInfo, err := l.WeatheInfo.GetWeatherInfo(name)

	if err != nil {
		return nil, err
	}

	return &TempDTO{
		Celsius:    weatherInfo.Current.TempC,
		Fahrenheit: weatherInfo.Current.TempF,
		Kelvin:     weatherInfo.Current.TempC + 273,
	}, nil
}

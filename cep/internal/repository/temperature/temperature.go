package temperature_repository

import utils_interface "github.com/andremelinski/observability/cep/internal/pkg/utils/interface"

type TempDTO struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

type TemperatureRepository struct {
	WeatheInfo utils_interface.IClimateInfoAPI
}

func NewClimateRepository(climateApi utils_interface.IClimateInfoAPI) *TemperatureRepository {
	return &TemperatureRepository{
		climateApi,
	}
}

func (l *TemperatureRepository) GetTempByPlaceName(name string) (*TempDTO, error) {
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

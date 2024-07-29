package usecases

import (
	"errors"

	"github.com/andremelinski/observability/cep/internal/repository"
)

type ClimateLocationInfoUseCaseDTO struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type IClimateLocationInfoUseCase interface {
	GetCityTemp(name string) (*ClimateLocationInfoUseCaseDTO, error)
}

type TemperatureRepository struct {
	LocInfo     repository.ILocationInfo
	WeatherInfo repository.IWeatherInfo
}

func NewClimateLocationInfoUseCase(locInfo repository.ILocationInfo, weatherInfo repository.IWeatherInfo) *TemperatureRepository {
	return &TemperatureRepository{
		locInfo,
		weatherInfo,
	}
}

func (l *TemperatureRepository) GetCityTemp(name string) (*ClimateLocationInfoUseCaseDTO, error) {

	cityInfo, err := l.LocInfo.GetLocationInfo(name)

	if err != nil {
		return nil, err
	}

	if cityInfo.Localidade == "" {
		return nil, errors.New("city not found")
	}

	weatherInfo, err := l.WeatherInfo.GetTempByPlaceName(cityInfo.Localidade)

	if err != nil {
		return nil, err
	}

	return &ClimateLocationInfoUseCaseDTO{
		City:       cityInfo.Localidade,
		Celsius:    weatherInfo.Celsius,
		Fahrenheit: weatherInfo.Fahrenheit,
		Kelvin:     weatherInfo.Kelvin,
	}, nil
}

package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/andremelinski/observability/principal/internal/pkg/utils"
)

type ZipCodeHandlerOutputDTO struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type CityTemperatureUseCase struct {
	cityTempURL        string
	handlerExternalApi utils.IHandlerExternalApi
}

type ICityTemperatureUseCase interface {
	CallCityTemp(ctx context.Context, place string) (*ZipCodeHandlerOutputDTO, error)
}

func NewCityTemperatureUseCase(cityTempURL string, handlerExternalApi utils.IHandlerExternalApi) *CityTemperatureUseCase {
	return &CityTemperatureUseCase{
		cityTempURL,
		handlerExternalApi,
	}
}

func (ct *CityTemperatureUseCase) CallCityTemp(ctx context.Context, zipcode string) (*ZipCodeHandlerOutputDTO, error) {

	// url := fmt.Sprintf("%s?zipcode=%s", ct.cityTempURL, url.QueryEscape(zipcode))
	url := fmt.Sprintf("http://cep:8081?zipcode=%s", url.QueryEscape(zipcode))

	bytes, err := ct.handlerExternalApi.CallExternalApi(ctx, 5000, "GET", url)

	if err != nil {
		return nil, err
	}

	dto := ZipCodeHandlerOutputDTO{}
	json.Unmarshal(bytes, &dto)

	if dto.City == "" {
		return nil, errors.New("no matching location found")
	}
	return &dto, nil
}

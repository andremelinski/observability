package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
	utils_interface "github.com/andremelinski/observability/cep/internal/pkg/utils/interface"
)

type WeatherInfo struct {
	apiKey             string
	handlerExternalApi utils_interface.IHandlerExternalApi
}

func NewWeatherInfo(apiKey string, handlerExternalApi utils_interface.IHandlerExternalApi) *WeatherInfo {
	return &WeatherInfo{
		apiKey,
		handlerExternalApi,
	}
}

func (c *WeatherInfo) GetWeatherInfo(ctx context.Context, place string) (*utils_dto.WeatherApiDTO, error) {

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes", c.apiKey, url.QueryEscape(place))
	bytes, err := c.handlerExternalApi.CallExternalApi(ctx, 5000, "GET", url)

	if err != nil {
		return nil, err
	}

	dto := utils_dto.WeatherApiDTO{}
	json.Unmarshal(bytes, &dto)

	if dto.Location.Name == "" {
		return nil, errors.New("no matching location found")
	}
	return &dto, nil
}

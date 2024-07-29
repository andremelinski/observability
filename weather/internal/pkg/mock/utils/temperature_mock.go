package mock_utils

import (
	"context"

	utils_dto "github.com/andremelinski/observability/weather/internal/pkg/utils/dto"
	"github.com/stretchr/testify/mock"
)

type WeatherInfoMock struct {
	mock.Mock
}

// GetCEPInfo implements utils.ICepInfoAPI.
func (m *WeatherInfoMock) GetWeatherInfo(ctx context.Context, place string) (*utils_dto.WeatherApiDTO, error) {
	args := m.Called(place)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*utils_dto.WeatherApiDTO), args.Error(1)
}

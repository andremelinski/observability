package usecases

import (
	"errors"
	"testing"

	mock_utils "github.com/andremelinski/observability/cep/internal/pkg/mock/utils"
	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
	"github.com/stretchr/testify/suite"
)

type TemperatureUseCaseTestSuite struct {
	suite.Suite
	temperatureUseCase *TemperatureUseCase
	mockWeatherApi     *mock_utils.WeatherInfoMock
	mockCep            string
}

func (suite *TemperatureUseCaseTestSuite) SetupSuite() {
	suite.mockWeatherApi = new(mock_utils.WeatherInfoMock)
	suite.temperatureUseCase = NewClimateUseCase(suite.mockWeatherApi)
	suite.mockCep = "cep"
}

func TestSuiteWeather(t *testing.T) {
	suite.Run(t, new(TemperatureUseCaseTestSuite))
}

func (suite *TemperatureUseCaseTestSuite) Test_GetLocationInfo_GetCEPInfo_Throw_Error() {

	suite.mockWeatherApi.On("GetWeatherInfo", suite.mockCep).Return(nil, errors.New("random error")).Once()

	output, err := suite.temperatureUseCase.GetTempByPlaceName(suite.mockCep)

	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *TemperatureUseCaseTestSuite) Test_GetLocationInfo_GetCEPInfo_ReturnDTO() {
	utilDto := &utils_dto.WeatherApiDTO{
		Location: utils_dto.Location{
			Name:           "Tokyo",
			Region:         "Tokyo",
			Country:        "Japan",
			Lat:            35.69,
			Lon:            139.69,
			TzID:           "Asia/Tokyo",
			LocaltimeEpoch: 1720368614,
			Localtime:      "2024-07-08 1:10",
		},
		Current: utils_dto.Current{
			LastUpdatedEpoch: 1720368000,
			LastUpdated:      "2024-07-08 01:00",
			TempC:            27.4,
			TempF:            81.3,
			IsDay:            0,
			Condition: utils_dto.Condition{
				Text: "Clear",
				Icon: "//cdn.weatherapi.com/weather/64x64/night/113.png",
				Code: 1000,
			},
			WindMph:    2.5,
			WindKph:    4.0,
			WindDegree: 60,
			WindDir:    "ENE",
			PressureMb: 1004.0,
			PressureIn: 29.65,
			PrecipMm:   0.0,
			PrecipIn:   0.0,
			Humidity:   89,
			Cloud:      0,
			FeelslikeC: 29.1,
			FeelslikeF: 84.3,
			WindchillC: 29.1,
			WindchillF: 84.4,
			HeatindexC: 31.8,
			HeatindexF: 89.2,
			DewpointC:  21.1,
			DewpointF:  70.0,
			VisKm:      10.0,
			VisMiles:   6.0,
			Uv:         1.0,
			GustMph:    8.3,
			GustKph:    13.3,
			AirQuality: utils_dto.AirQuality{
				Co:           620.8,
				No2:          65.1,
				O3:           11.6,
				So2:          14.9,
				Pm25:         105.8,
				Pm10:         111.8,
				UsEpaIndex:   4,
				GbDefraIndex: 10,
			},
		},
	}
	suite.mockWeatherApi.On("GetWeatherInfo", suite.mockCep).Return(utilDto, nil).Once()

	output, err := suite.temperatureUseCase.GetTempByPlaceName(suite.mockCep)

	suite.NoError(err)
	suite.Equal(&TempDTO{
		utilDto.TempC,
		utilDto.TempF,
		utilDto.TempC + 273,
	}, output)
}

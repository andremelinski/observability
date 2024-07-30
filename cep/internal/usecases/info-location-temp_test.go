package usecases

import (
	"context"
	"errors"
	"testing"

	mock_opentelemetry "github.com/andremelinski/observability/cep/internal/pkg/mock/opentelemetry"
	mock_usecase "github.com/andremelinski/observability/cep/internal/pkg/mock/repository"
	cep_repository "github.com/andremelinski/observability/cep/internal/repository/cep"
	temperature_repository "github.com/andremelinski/observability/cep/internal/repository/temperature"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ClimateLocationInfoUseCaseTestSuite struct {
	suite.Suite
	locationRepository *TemperatureRepository
	mockCepRepo        *mock_usecase.LocationUseCaseMock
	mockTempRepo       *mock_usecase.TemperatureUseCaseMock
	mockCep            string
	mockStartOTELTrace *mock_opentelemetry.StartOTELTraceMock
	mockTracer         *mock_opentelemetry.MockTracer
	mockMockSpan       *mock_opentelemetry.MockSpan
}

func (suite *ClimateLocationInfoUseCaseTestSuite) SetupSuite() {
	suite.mockCepRepo = new(mock_usecase.LocationUseCaseMock)
	suite.mockTempRepo = new(mock_usecase.TemperatureUseCaseMock)
	// Create an instance of MockTracer
	suite.mockTracer = new(mock_opentelemetry.MockTracer)

	suite.mockMockSpan = new(mock_opentelemetry.MockSpan)

	suite.mockStartOTELTrace = new(mock_opentelemetry.StartOTELTraceMock)

	suite.locationRepository = NewClimateLocationInfoUseCase(suite.mockCepRepo, suite.mockTempRepo, suite.mockTracer, suite.mockStartOTELTrace)
	suite.mockCep = "cep"
}

func TestSuiteLocation(t *testing.T) {
	suite.Run(t, new(ClimateLocationInfoUseCaseTestSuite))
}

func (suite *ClimateLocationInfoUseCaseTestSuite) Test_GetLocationInfo_GetCEPInfo_Throw_Error() {
	ctx := context.Background()
	suite.mockStartOTELTrace.On("StartOTELTrace", ctx, suite.mockTracer, "via-cep-trace").Return(ctx, suite.mockMockSpan).Once()

	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(nil, errors.New("random error")).Once()

	suite.mockMockSpan.On("End", mock.Anything).Return().Once()

	output, err := suite.locationRepository.GetCityTemp(ctx, suite.mockCep)

	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *ClimateLocationInfoUseCaseTestSuite) Test_GetLocationInfo_GetCEPInfo_Wrong_CEP() {
	utilDto := &cep_repository.LocationOutputDTO{
		Cep:         "0000-000",
		Logradouro:  "Rua XXXXXX",
		Complemento: "",
		Bairro:      "Boa Vista",
		Localidade:  "",
		UF:          "PR",
		DDD:         "41",
	}

	ctx := context.Background()
	suite.mockStartOTELTrace.On("StartOTELTrace", ctx, suite.mockTracer, "via-cep-trace").Return(ctx, suite.mockMockSpan).Once()

	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(utilDto, nil).Once()

	suite.mockMockSpan.On("End", mock.Anything).Return().Once()

	output, err := suite.locationRepository.GetCityTemp(ctx, suite.mockCep)

	suite.Empty(output)
	suite.EqualError(err, "city not found")
}

func (suite *ClimateLocationInfoUseCaseTestSuite) Test_GetLocationInfo_GetCEPInfo_ReturnDTO() {
	utilDto := &cep_repository.LocationOutputDTO{
		Cep:         "0000-000",
		Logradouro:  "Rua XXXXXX",
		Complemento: "",
		Bairro:      "Boa Vista",
		Localidade:  "Curitiba",
		UF:          "PR",
		DDD:         "41",
	}

	ctx := context.Background()

	suite.mockStartOTELTrace.On("StartOTELTrace", ctx, suite.mockTracer, "via-cep-trace").Return(ctx, suite.mockMockSpan).Once()

	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(utilDto, nil).Once()

	suite.mockMockSpan.On("End", mock.Anything).Return().Once()

	suite.mockStartOTELTrace.On("StartOTELTrace", ctx, suite.mockTracer, "weather-api-trace").Return(ctx, suite.mockMockSpan).Once()
	suite.mockTempRepo.On("GetTempByPlaceName", ctx, utilDto.Localidade).Return(nil, errors.New("random error")).Once()
	suite.mockMockSpan.On("End", mock.Anything).Return().Once()

	output, err := suite.locationRepository.GetCityTemp(ctx, suite.mockCep)

	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *ClimateLocationInfoUseCaseTestSuite) Test_GetCityTemp_Correct() {
	cepResp := &cep_repository.LocationOutputDTO{
		Cep:         "0000-000",
		Logradouro:  "Rua XXXXXX",
		Complemento: "",
		Bairro:      "Boa Vista",
		Localidade:  "Curitiba",
		UF:          "PR",
		DDD:         "41",
	}

	tempResp := &temperature_repository.TempDTO{
		Celsius:    1,
		Fahrenheit: 1,
		Kelvin:     274,
	}

	ctx := context.Background()

	suite.mockStartOTELTrace.On("StartOTELTrace", ctx, suite.mockTracer, "via-cep-trace").Return(ctx, suite.mockMockSpan).Once()

	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(cepResp, nil).Once()

	suite.mockMockSpan.On("End", mock.Anything).Return().Once()

	suite.mockStartOTELTrace.On("StartOTELTrace", ctx, suite.mockTracer, "weather-api-trace").Return(ctx, suite.mockMockSpan).Once()
	suite.mockTempRepo.On("GetTempByPlaceName", ctx, cepResp.Localidade).Return(tempResp, nil).Once()
	suite.mockMockSpan.On("End", mock.Anything).Return().Once()

	output, err := suite.locationRepository.GetCityTemp(ctx, suite.mockCep)

	suite.NoError(err)
	suite.Equal(&ClimateLocationInfoUseCaseDTO{
		City:       "Curitiba",
		Celsius:    1,
		Fahrenheit: 1,
		Kelvin:     274,
	}, output)
}

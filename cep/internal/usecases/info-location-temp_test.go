package usecases

import (
	"context"
	"errors"
	"testing"

	mock_usecase "github.com/andremelinski/observability/cep/internal/pkg/mock/repository"
	cep_repository "github.com/andremelinski/observability/cep/internal/repository/cep"
	temperature_repository "github.com/andremelinski/observability/cep/internal/repository/temperature"
	"github.com/stretchr/testify/suite"
)

type ClimateLocationInfoUseCaseTestSuite struct {
	suite.Suite
	locationRepository *TemperatureRepository
	mockCepRepo        *mock_usecase.LocationUseCaseMock
	mockTempRepo       *mock_usecase.TemperatureUseCaseMock
	mockCep            string
}

func (suite *ClimateLocationInfoUseCaseTestSuite) SetupSuite() {
	suite.mockCepRepo = new(mock_usecase.LocationUseCaseMock)
	suite.mockTempRepo = new(mock_usecase.TemperatureUseCaseMock)
	suite.locationRepository = NewClimateLocationInfoUseCase(suite.mockCepRepo, suite.mockTempRepo)
	suite.mockCep = "cep"
}

func TestSuiteLocation(t *testing.T) {
	suite.Run(t, new(ClimateLocationInfoUseCaseTestSuite))
}

func (suite *ClimateLocationInfoUseCaseTestSuite) Test_GetLocationInfo_GetCEPInfo_Throw_Error() {
	ctx := context.Background()
	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(nil, errors.New("random error")).Once()

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
	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(utilDto, nil).Once()

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
	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(utilDto, nil).Once()
	suite.mockTempRepo.On("GetTempByPlaceName", ctx, utilDto.Localidade).Return(nil, errors.New("random error")).Once()

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
	suite.mockCepRepo.On("GetLocationInfo", ctx, suite.mockCep).Return(cepResp, nil).Once()
	suite.mockTempRepo.On("GetTempByPlaceName", ctx, cepResp.Localidade).Return(tempResp, nil).Once()

	output, err := suite.locationRepository.GetCityTemp(ctx, suite.mockCep)

	suite.NoError(err)
	suite.Equal(&ClimateLocationInfoUseCaseDTO{
		City:       "Curitiba",
		Celsius:    1,
		Fahrenheit: 1,
		Kelvin:     274,
	}, output)
}

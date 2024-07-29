package cep_repository

import (
	"errors"
	"testing"

	mock_utils "github.com/andremelinski/observability/cep/internal/pkg/mock/utils"
	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
	"github.com/stretchr/testify/suite"
)

type LocationRepositoryTestSuite struct {
	suite.Suite
	locationRepository *LocationRepository
	mockViaCep         *mock_utils.CEPInfoMock
	mockCep            string
}

func (suite *LocationRepositoryTestSuite) SetupSuite() {
	suite.mockViaCep = new(mock_utils.CEPInfoMock)
	suite.locationRepository = NewLocationRepository(suite.mockViaCep)
	suite.mockCep = "cep"
}

func TestSuiteLocation(t *testing.T) {
	suite.Run(t, new(LocationRepositoryTestSuite))
}

func (suite *LocationRepositoryTestSuite) Test_GetLocationInfo_GetCEPInfo_Throw_Error() {

	suite.mockViaCep.On("GetCEPInfo", suite.mockCep).Return(nil, errors.New("random error")).Once()

	output, err := suite.locationRepository.GetLocationInfo(suite.mockCep)

	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *LocationRepositoryTestSuite) Test_GetLocationInfo_GetCEPInfo_ReturnDTO() {
	utilDto := &utils_dto.ViaCepDTO{
		Cep:         "0000-000",
		Logradouro:  "Rua XXXXXX",
		Complemento: "",
		Unidade:     "",
		Bairro:      "Boa Vista",
		Localidade:  "Curitiba",
		UF:          "PR",
		IBGE:        "0000000",
		Gia:         "",
		DDD:         "41",
		Siafi:       "0000",
	}
	suite.mockViaCep.On("GetCEPInfo", suite.mockCep).Return(utilDto, nil).Once()

	output, err := suite.locationRepository.GetLocationInfo(suite.mockCep)

	suite.NoError(err)
	suite.Equal(&LocationOutputDTO{
		Cep:         utilDto.Cep,
		Logradouro:  utilDto.Logradouro,
		Complemento: utilDto.Complemento,
		Bairro:      utilDto.Bairro,
		Localidade:  utilDto.Localidade,
		UF:          utilDto.UF,
		DDD:         utilDto.DDD,
	}, output)
}

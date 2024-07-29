package utils

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	mock_utils "github.com/andremelinski/observability/cep/internal/pkg/mock/utils"
	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
	"github.com/stretchr/testify/suite"
)

type ViaCepCaseTestSuite struct {
	suite.Suite
	cepInfo             *CepInfo
	mockCallExternalApi *mock_utils.CallExternalApiMock
	mockCep             string
}

func (suite *ViaCepCaseTestSuite) SetupSuite() {
	suite.mockCep = "cep"
	suite.mockCallExternalApi = new(mock_utils.CallExternalApiMock)
	suite.cepInfo = NewCepInfo(suite.mockCallExternalApi)
}

func TestSuiteViacep(t *testing.T) {
	suite.Run(t, new(ViaCepCaseTestSuite))
}

func (suite *ViaCepCaseTestSuite) Test_GetWeatherInfo_Throw_Error_Wrong_Place() {
	suite.mockCallExternalApi.On("CallExternalApi", context.Background(), 3000, "GET", "https://viacep.com.br/ws/cep/json/").Return(nil, errors.New("random error")).Once()

	output, err := suite.cepInfo.GetCEPInfo(suite.mockCep)

	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *ViaCepCaseTestSuite) Test_GetLocationInfo_GetCEPInfo_ReturnDTO() {
	utilDto := &utils_dto.ViaCepDTO{
		Cep:         "82540-091",
		Logradouro:  "Rua Desembargador Aurélio Feijó",
		Complemento: "",
		Unidade:     "",
		Bairro:      "Boa Vista",
		Localidade:  "Curitiba",
		UF:          "PR",
		IBGE:        "4106902",
		Gia:         "",
		DDD:         "41",
		Siafi:       "7535",
	}
	bytes, _ := json.Marshal(utilDto)

	suite.mockCallExternalApi.On("CallExternalApi", context.Background(), 3000, "GET", "https://viacep.com.br/ws/cep/json/").Return(bytes, nil).Once()

	output, err := suite.cepInfo.GetCEPInfo(suite.mockCep)

	suite.NoError(err)
	suite.Equal(utilDto, output)
}

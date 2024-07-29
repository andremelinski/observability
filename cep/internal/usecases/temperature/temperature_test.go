package usecases_temperature

import (
	"context"
	"errors"
	"testing"

	grpc_interfaces "github.com/andremelinski/observability/cep/internal/infra/grpc/interfaces"
	mock_grpc_client "github.com/andremelinski/observability/cep/internal/pkg/mock/grpc"
	usecases_dto "github.com/andremelinski/observability/cep/internal/usecases/dto"
	"github.com/stretchr/testify/suite"
)

type TemperatureUseTestSuite struct {
	suite.Suite
	temperatureUseCase *TemperatureUseCase
	mockViaCep         *mock_grpc_client.TemperatureUseCaseMock
}

func (suite *TemperatureUseTestSuite) SetupSuite() {
	suite.mockViaCep = new(mock_grpc_client.TemperatureUseCaseMock)
	suite.temperatureUseCase = NewClimateUseCase(suite.mockViaCep)
}

func TestSuiteLocation(t *testing.T) {
	suite.Run(t, new(TemperatureUseTestSuite))
}

func (suite *TemperatureUseTestSuite) Test_GetLocationInfo_GetCEPInfo_Throw_Error() {
	ctx := context.Background()
	locationName := "curitiba"
	suite.mockViaCep.On("GetLocationTemperature", ctx, locationName).Return(nil, errors.New("random error")).Once()

	output, err := suite.temperatureUseCase.GetTempByPlaceName(ctx, locationName)

	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *TemperatureUseTestSuite) Test_GetLocationInfo_GetCEPInfo_ReturnDTO() {
	ctx := context.Background()
	locationName := "curitiba"
	grpcResp := &grpc_interfaces.TempResponseDTO{
		Temp_C: 1,
		Temp_F: 2,
		Temp_K: 273,
	}
	finalResp := &usecases_dto.TempDTOOutput{
		Celsius:    1,
		Fahrenheit: 2,
		Kelvin:     273,
	}

	suite.mockViaCep.On("GetLocationTemperature", ctx, locationName).Return(grpcResp, nil).Once()

	output, err := suite.temperatureUseCase.GetTempByPlaceName(ctx, locationName)

	suite.Empty(err)
	suite.NoError(err)
	suite.Equal(finalResp, output)
}

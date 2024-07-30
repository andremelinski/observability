package mock_utils

import (
	"context"

	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
	"github.com/stretchr/testify/mock"
)

type CEPInfoMock struct {
	mock.Mock
}

// GetCEPInfo implements utils.ICepInfoAPI.
func (m *CEPInfoMock) GetCEPInfo(ctx context.Context, cep string) (*utils_dto.ViaCepDTO, error) {
	args := m.Called(ctx, cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*utils_dto.ViaCepDTO), args.Error(1)
}

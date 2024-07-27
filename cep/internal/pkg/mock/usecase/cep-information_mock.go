package mock_usecase

import (
	"context"

	usecases_dto "github.com/andremelinski/observability/cep/internal/usecases/dto"
	"github.com/stretchr/testify/mock"
)

type LocationUseCaseMock struct {
	mock.Mock
}

func (m *LocationUseCaseMock) GetLocationInfo(ctx context.Context, cep string) (*usecases_dto.LocationOutputDTO, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecases_dto.LocationOutputDTO), args.Error(1)
}

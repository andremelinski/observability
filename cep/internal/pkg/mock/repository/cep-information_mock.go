package mock_usecase

import (
	"context"

	cep_repository "github.com/andremelinski/observability/cep/internal/repository/cep"
	"github.com/stretchr/testify/mock"
)

type LocationUseCaseMock struct {
	mock.Mock
}

func (m *LocationUseCaseMock) GetLocationInfo(ctx context.Context, cep string) (*cep_repository.LocationOutputDTO, error) {
	args := m.Called(ctx, cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*cep_repository.LocationOutputDTO), args.Error(1)
}

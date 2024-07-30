package mock_usecases

import (
	"context"

	"github.com/andremelinski/observability/cep/internal/usecases"
	"github.com/stretchr/testify/mock"
)

type GetCityTempInfoUseCaseMock struct {
	mock.Mock
}

func (m *GetCityTempInfoUseCaseMock) GetCityTemp(ctx context.Context, name string) (*usecases.ClimateLocationInfoUseCaseDTO, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecases.ClimateLocationInfoUseCaseDTO), args.Error(1)
}

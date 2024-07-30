package mock_usecase

import (
	"context"

	temperature_repository "github.com/andremelinski/observability/cep/internal/repository/temperature"
	"github.com/stretchr/testify/mock"
)

type TemperatureUseCaseMock struct {
	mock.Mock
}

func (m *TemperatureUseCaseMock) GetTempByPlaceName(ctx context.Context, name string) (*temperature_repository.TempDTO, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*temperature_repository.TempDTO), args.Error(1)
}

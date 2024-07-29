package mock_grpc_client

import (
	"context"

	grpc_interfaces "github.com/andremelinski/observability/cep/internal/infra/grpc/interfaces"
	"github.com/stretchr/testify/mock"
)

type TemperatureUseCaseMock struct {
	mock.Mock
}

func (m *TemperatureUseCaseMock) GetLocationTemperature(ctx context.Context, cep string) (*grpc_interfaces.TempResponseDTO, error) {
	args := m.Called(ctx, cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*grpc_interfaces.TempResponseDTO), args.Error(1)
}

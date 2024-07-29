package mock_usecase

import (
	"github.com/andremelinski/observability/cep/internal/usecases"
	"github.com/stretchr/testify/mock"
)

type TemperatureUseCaseMock struct {
	mock.Mock
}

func (m *TemperatureUseCaseMock) GetTempByPlaceName(name string) (*usecases.TempDTO, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecases.TempDTO), args.Error(1)
}

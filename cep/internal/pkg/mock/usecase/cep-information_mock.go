package mock_usecase

import (
	"github.com/andremelinski/observability/cep/internal/usecases"
	"github.com/stretchr/testify/mock"
)

type LocationUseCaseMock struct {
	mock.Mock
}

func (m *LocationUseCaseMock) GetLocationInfo(cep string) (*usecases.LocationOutputDTO, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecases.LocationOutputDTO), args.Error(1)
}

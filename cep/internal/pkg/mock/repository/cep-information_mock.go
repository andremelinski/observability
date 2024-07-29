package mock_usecase

import (
	cep_repository "github.com/andremelinski/observability/cep/internal/repository/cep"
	"github.com/stretchr/testify/mock"
)

type LocationUseCaseMock struct {
	mock.Mock
}

func (m *LocationUseCaseMock) GetLocationInfo(cep string) (*cep_repository.LocationOutputDTO, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*cep_repository.LocationOutputDTO), args.Error(1)
}

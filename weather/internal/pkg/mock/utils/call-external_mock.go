package mock_utils

import (
	"context"

	"github.com/stretchr/testify/mock"
)


type CallExternalApiMock struct {
	mock.Mock
}

// GetCEPInfo implements utils.ICepInfoAPI.
func (m *CallExternalApiMock) CallExternalApi(ctx context.Context, timeoutMs int, method string, url string) ([]byte, error) {
	args := m.Called(ctx, timeoutMs, method, url)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}
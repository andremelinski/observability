package mock_pb

import (
	"github.com/andremelinski/observability/weather/internal/infra/grpc/pb"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type PbWeatherServiceServerMock struct {
	mock.Mock
	grpc.ServerStream
	RecvFunc func() (*pb.WeatherLocationRequest, error)
	SendFunc func(*pb.WeatherLocationResponse) error
}

func (m *PbWeatherServiceServerMock) Recv() (*pb.WeatherLocationRequest, error) {
	if m.RecvFunc != nil {
		return m.RecvFunc()
	}
	args := m.Called()
	return args.Get(0).(*pb.WeatherLocationRequest), args.Error(1)
}

func (m *PbWeatherServiceServerMock) Send(response *pb.WeatherLocationResponse) error {
	if m.SendFunc != nil {
		return m.SendFunc(response)
	}
	args := m.Called(response)
	return args.Error(0)
}

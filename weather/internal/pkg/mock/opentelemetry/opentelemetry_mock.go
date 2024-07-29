package mock_opentelemetry

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/otel/trace"
)

// MockHandlerTrace is a mock implementation of the IHandlerTrace interface
type MockHandlerTrace struct {
	mock.Mock
}

func (m *MockHandlerTrace) StartOTELTrace(ctx context.Context, otelTracer trace.Tracer, traceMessage string) (context.Context, trace.Span) {
	args := m.Called(ctx, otelTracer, traceMessage)
	return args.Get(0).(context.Context), args.Get(1).(trace.Span)
}

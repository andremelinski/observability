package mock_opentelemetry

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/otel/trace"
)

type StartOTELTraceMock struct {
	mock.Mock
}

// tracer implements trace.Tracer.

func (m *StartOTELTraceMock) Tracer() trace.Tracer {
	args := m.Called()
	return args.Get(0).(trace.Tracer)
}

// Start implements trace.Tracer.
func (m *StartOTELTraceMock) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	args := m.Called(ctx, spanName, opts)
	if args.Get(0) == nil {
		return nil, nil
	}
	return args.Get(0).(context.Context), args.Get(1).(trace.Span)
}

func (m *StartOTELTraceMock) StartOTELPropagator(r *http.Request) context.Context {
	args := m.Called(r)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(context.Context)
}

func (m *StartOTELTraceMock) StartOTELTrace(ctx context.Context, otelTracer trace.Tracer, traceMessage string) (context.Context, trace.Span) {
	args := m.Called(ctx, otelTracer, traceMessage)
	if args.Get(0) == nil {
		return nil, nil
	}
	return args.Get(0).(context.Context), args.Get(1).(trace.Span)
}

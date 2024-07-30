package mock_opentelemetry

import (
	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type MockTracer struct {
	trace.Tracer
}

type MockSpan struct {
	trace.Span
	mock.Mock
}

func (m *MockSpan) End(options ...trace.SpanEndOption) {
	m.Called(options)
}

func (m *MockSpan) AddEvent(name string, options ...trace.EventOption) {
	m.Called(name, options)
}

func (m *MockSpan) IsRecording() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockSpan) RecordError(err error, options ...trace.EventOption) {
	m.Called(err, options)
}

func (m *MockSpan) SpanContext() trace.SpanContext {
	args := m.Called()
	return args.Get(0).(trace.SpanContext)
}

func (m *MockSpan) SetStatus(code codes.Code, description string) {
	m.Called(code, description)
}

func (m *MockSpan) SetName(name string) {
	m.Called(name)
}

func (m *MockSpan) SetAttributes(attributes ...attribute.KeyValue) {
	m.Called(attributes)
}

func (m *MockSpan) TracerProvider() trace.TracerProvider {
	args := m.Called()
	return args.Get(0).(trace.TracerProvider)
}

func (m *MockSpan) AddLink(link trace.Link) {
	m.Called(link)
}

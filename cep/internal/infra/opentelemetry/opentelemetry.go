package opentelemetry

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IHandlerTrace interface {
	StartOTELTrace(r *http.Request, otelTracer trace.Tracer, traceMessage string) (context.Context, trace.Span)
}

type OtelInfo struct {
	RequestNameOTEL string
	ServiceName     string
	CollectorURL    string
}

type TracerOpenTelemetry struct {
	otelInfo *OtelInfo
}

func NewOpenTelemetry(otelInfo *OtelInfo) *TracerOpenTelemetry {
	return &TracerOpenTelemetry{
		otelInfo,
	}
}

func (t *TracerOpenTelemetry) InitProvider(ctx context.Context) (func(context.Context) error, error) {

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(t.otelInfo.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.NewClient(t.otelInfo.CollectorURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}
	// Exporter do trace -> exporta os dados por uma comunicacao gRPC (mas pode ser http)
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}

func (t *TracerOpenTelemetry) InitOTELTrace(traceName string) trace.Tracer {
	return otel.Tracer(traceName)
}

func (t *TracerOpenTelemetry) StartOTELTrace(r *http.Request, otelTracer trace.Tracer, traceMessage string) (context.Context, trace.Span) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	message := fmt.Sprintf("%s %s", traceMessage, t.otelInfo.RequestNameOTEL)
	return otelTracer.Start(ctx, message)
}

package main

import (
	"context"
	"log"

	"github.com/andremelinski/observability/weather/configs"
	"github.com/andremelinski/observability/weather/internal/infra/grpc/handlers"
	"github.com/andremelinski/observability/weather/internal/infra/opentelemetry"
	"github.com/andremelinski/observability/weather/internal/pkg/utils"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)


func main(){
	ctx := context.Background()
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	observability := opentelemetry.NewOpenTelemetry(
		&opentelemetry.OtelInfo{
			RequestNameOTEL: configs.REQUEST_NAME_OTEL, 
			ServiceName: configs.OTEL_SERVICE_NAME, 
			CollectorURL: configs.OTEL_EXPORTER_OTLP_ENDPOINT,
		},
	)

	shutdown, err := observability.InitProvider()

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tracer := observability.InitOTELTrace("weather -ms-tracer")

// ############### app ###############
	grpcServer := grpc.NewServer(
		 grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	handlerExternalApi := utils.NewHandlerExternalApi()
	
	handlers.NewGrpcResgisters(grpcServer, configs.WEATHER_API_KEY, handlerExternalApi, tracer, observability).CreateGrpcWeatherRegisters()
	handlers.NewGrpcServer(grpcServer, configs.GRPC_PORT).Start()
}
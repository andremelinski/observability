package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/andremelinski/observability/cep/configs"
	grpc_client "github.com/andremelinski/observability/cep/internal/infra/grpc/client"
	"github.com/andremelinski/observability/cep/internal/infra/grpc/handlers"
	"github.com/andremelinski/observability/cep/internal/infra/web"
	web_handler "github.com/andremelinski/observability/cep/internal/infra/web/handlers"
	"github.com/andremelinski/observability/cep/internal/pkg/utils"
	utils_cep "github.com/andremelinski/observability/cep/internal/pkg/utils/cep"
	web_utils "github.com/andremelinski/observability/cep/internal/pkg/utils/web"
	usecases_cep "github.com/andremelinski/observability/cep/internal/usecases/cep"
	usecases_temperature "github.com/andremelinski/observability/cep/internal/usecases/temperature"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func initProvider(serviceName, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.NewClient( collectorURL,
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


func main(){
	ctx := context.Background()
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

		// inicia o provider
	shutdown, err := initProvider(configs.OTEL_SERVICE_NAME,configs.OTEL_EXPORTER_OTLP_ENDPOINT)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tracer := otel.Tracer("microservice-tracer")

	// ############### app ###############
	handlerExternalApi := utils.NewHandlerExternalApi()
	grpcServer :=handlers.NewGrpcServer(configs.GRPC_SERVER_NAME, configs.GRPC_PORT)
	
	grpcWeatherService := grpc_client.NewWeatherService(grpcServer)
	tempUseCase := usecases_temperature.NewClimateUseCase(grpcWeatherService)

	cepUtils := utils_cep.NewCepInfo(handlerExternalApi)
	ceUseCase := usecases_cep.NewLocationUseCase(cepUtils)

	webresponseHandler := web_utils.NewWebResponseHandler()

	hand := web_handler.NewLocalTemperatureHandler(ceUseCase, tempUseCase, webresponseHandler, &web_handler.OtelInfo{RequestNameOTEL: configs.REQUEST_NAME_OTEL, OTELTracer: tracer})

	webRouter := web.NewWebRouter(hand)
	webServer := web.NewWebServer(
		configs.HTTP_PORT,
		webRouter.BuildHandlers(),
	)

	webServer.Start()
}
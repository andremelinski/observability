package main

import (
	"context"
	"log"

	"github.com/andremelinski/observability/cep/configs"
	grpc_client "github.com/andremelinski/observability/cep/internal/infra/grpc/client"
	"github.com/andremelinski/observability/cep/internal/infra/grpc/handlers"
	"github.com/andremelinski/observability/cep/internal/infra/opentelemetry"
	"github.com/andremelinski/observability/cep/internal/infra/web"
	web_handler "github.com/andremelinski/observability/cep/internal/infra/web/handlers"
	"github.com/andremelinski/observability/cep/internal/pkg/utils"
	utils_cep "github.com/andremelinski/observability/cep/internal/pkg/utils/cep"
	web_utils "github.com/andremelinski/observability/cep/internal/pkg/utils/web"
	usecases_cep "github.com/andremelinski/observability/cep/internal/usecases/cep"
	usecases_temperature "github.com/andremelinski/observability/cep/internal/usecases/temperature"
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

	tracer := observability.InitOTELTrace("cep-ms-tracer")

	// ############### app ###############
	handlerExternalApi := utils.NewHandlerExternalApi()
	grpcServer :=handlers.NewGrpcServer(configs.GRPC_SERVER_NAME, configs.GRPC_PORT)
	
	grpcWeatherService := grpc_client.NewWeatherService(grpcServer)
	tempUseCase := usecases_temperature.NewClimateUseCase(grpcWeatherService)

	cepUtils := utils_cep.NewCepInfo(handlerExternalApi)
	ceUseCase := usecases_cep.NewLocationUseCase(cepUtils)

	webresponseHandler := web_utils.NewWebResponseHandler()
	hand := web_handler.NewLocalTemperatureHandler(ceUseCase, tempUseCase, webresponseHandler, tracer, observability)

	webRouter := web.NewWebRouter(hand)
	webServer := web.NewWebServer(
		configs.HTTP_PORT,
		webRouter.BuildHandlers(),
	)

	webServer.Start()
}
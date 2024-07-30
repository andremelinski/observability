package main

import (
	"context"
	"log"

	"github.com/andremelinski/observability/principal/configs"
	"github.com/andremelinski/observability/principal/internal/infra/opentelemetry"
	web_infra "github.com/andremelinski/observability/principal/internal/infra/web"
	"github.com/andremelinski/observability/principal/internal/infra/web/webserver/handlers"
	"github.com/andremelinski/observability/principal/internal/pkg/utils"
	pkg_web "github.com/andremelinski/observability/principal/internal/pkg/web"
	"github.com/andremelinski/observability/principal/internal/usecase"
)

func main() {
	ctx := context.Background()
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	observability := opentelemetry.NewOpenTelemetry(
		&opentelemetry.OtelInfo{
			RequestNameOTEL: configs.REQUEST_NAME_OTEL,
			ServiceName:     configs.OTEL_SERVICE_NAME,
			CollectorURL:    configs.OTEL_EXPORTER_OTLP_ENDPOINT,
		},
	)

	shutdown, err := observability.InitProvider(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tracer := observability.InitOTELTrace(configs.OTEL_SERVICE_NAME)

	webresponseHandler := pkg_web.NewWebResponseHandler()

	externalApi := utils.NewHandlerExternalApi()
	usecase := usecase.NewCityTemperatureUseCase(configs.EXTERNAL_URL, externalApi)
	tempHandler := handlers.NewLocalTemperatureHandler(usecase, webresponseHandler, tracer, observability)

	webRouter := web_infra.NewWebRouter(tempHandler)
	webServer := web_infra.NewWebServer(
		configs.HTTP_PORT,
		webRouter.BuildHandlers(),
	)

	webServer.Start()
}

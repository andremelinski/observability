package main

import (
	"context"
	"log"

	"github.com/andremelinski/observability/cep/configs"
	"github.com/andremelinski/observability/cep/internal/composite"
	"github.com/andremelinski/observability/cep/internal/infra/opentelemetry"
	web_infra "github.com/andremelinski/observability/cep/internal/infra/web"
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

	tracer := observability.InitOTELTrace("cep-ms-tracer")

	tempHandler := composite.TemperatureLocationComposite(configs.WEATHER_API_KEY)

	webRouter := web_infra.NewWebRouter(tempHandler)
	webServer := web_infra.NewWebServer(
		configs.HTTP_PORT,
		webRouter.BuildHandlers(),
	)

	webServer.Start()
}

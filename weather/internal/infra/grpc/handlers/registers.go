package handlers

import (
	"github.com/andremelinski/observability/weather/internal/infra/grpc/pb"
	"github.com/andremelinski/observability/weather/internal/infra/grpc/service"
	"github.com/andremelinski/observability/weather/internal/infra/opentelemetry"
	utils_interface "github.com/andremelinski/observability/weather/internal/pkg/utils/interface"
	utils "github.com/andremelinski/observability/weather/internal/pkg/utils/weather"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type GrpcResgisterService struct{
	grpcServer *grpc.Server
	weatherApiKey string
	callExternalapi utils_interface.IHandlerExternalApi
	otelTrace trace.Tracer
	OtelTraceHandler opentelemetry.IHandlerTrace
}

func NewGrpcResgisters(grpcServer *grpc.Server, weatherApiKey string, callExternalapi utils_interface.IHandlerExternalApi, otelTrace trace.Tracer, otelInfo opentelemetry.IHandlerTrace)*GrpcResgisterService{
	return &GrpcResgisterService{
		grpcServer,
		weatherApiKey,
		callExternalapi,
		otelTrace,
otelInfo,
	}
}

func(gs *GrpcResgisterService) CreateGrpcWeatherRegisters(){
	
	a := utils.NewWeatherInfo(gs.weatherApiKey, gs.callExternalapi)
	weatherService := service.NewWeatherService(a, gs.otelTrace, gs.OtelTraceHandler)
    pb.RegisterWeatherServiceServer(gs.grpcServer, weatherService)
}
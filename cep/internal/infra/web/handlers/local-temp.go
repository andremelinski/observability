package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/andremelinski/observability/cep/internal/infra/opentelemetry"
	web_utils "github.com/andremelinski/observability/cep/internal/pkg/utils/web"
	"github.com/andremelinski/observability/cep/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type ICityWebHandler interface {
	CityTemperature(w http.ResponseWriter, r *http.Request)
}


type LocalTemperatureHandler struct{
	CepUseCase usecases.ILocationInfo
	TempUseCase usecases.IWeatherInfo
	HttpResponse web_utils.IWebResponseHandler
	otelTrace trace.Tracer
	OtelTraceHandler opentelemetry.IHandlerTrace
}

func NewLocalTemperatureHandler (cepUseCase usecases.ILocationInfo,tempUseCase usecases.IWeatherInfo, httpResponse web_utils.IWebResponseHandler, otelTrace trace.Tracer, otelInfo opentelemetry.IHandlerTrace) *LocalTemperatureHandler{
	return &LocalTemperatureHandler{
		cepUseCase,
		tempUseCase,
		httpResponse,
		otelTrace,
		otelInfo,
	}
}

func(lc *LocalTemperatureHandler) CityTemperature(w http.ResponseWriter, r *http.Request){
	ctx, span := lc.OtelTraceHandler.StartOTELTrace(r, lc.otelTrace, "CityTemperatureApi")

	defer span.End() // span acaba quando toda req acabar para ter o trace

	qs := r.URL.Query()
	zipStr := qs.Get("zipcode")

	if err := validateInput(zipStr); err != nil {
		lc.HttpResponse.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	info, err := lc.CepUseCase.GetLocationInfo(zipStr)
	if err != nil {
		fmt.Println(err)
		lc.HttpResponse.RespondWithError(w, http.StatusBadRequest, errors.New("can not find zipcode"))
		return 
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	climateInfo, err := lc.TempUseCase.GetTempByPlaceName(info.Localidade)

	if err != nil {
		fmt.Println(err)
		lc.HttpResponse.RespondWithError(w, http.StatusBadRequest, errors.New("can not find zipcode"))
		return 
	}

	lc.HttpResponse.Respond(w, http.StatusOK, climateInfo)
}

func validateInput(zipcode string) error {
	if zipcode == "" {
		return errors.New("invalid zipcode")
	}

	matched, err := regexp.MatchString(`\b\d{5}[\-]?\d{3}\b`, zipcode)
	if !matched || err != nil {
		return errors.New("invalid zipcode")
	}

	return nil
}

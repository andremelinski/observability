package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/andremelinski/observability/cep/internal/infra/opentelemetry"
	"github.com/andremelinski/observability/cep/internal/pkg/web"
	"github.com/andremelinski/observability/cep/internal/usecases"
	"go.opentelemetry.io/otel/trace"
)

type LocalTemperatureHandler struct {
	CepInfoUseCase   usecases.IClimateLocationInfoUseCase
	HttpResponse     web.IWebResponseHandler
	otelTrace        trace.Tracer
	OtelTraceHandler opentelemetry.IHandlerTrace
}

func NewLocalTemperatureHandler(cepInfoUseCase usecases.IClimateLocationInfoUseCase, httpResponse web.IWebResponseHandler, otelTrace trace.Tracer, otelInfo opentelemetry.IHandlerTrace) *LocalTemperatureHandler {
	return &LocalTemperatureHandler{
		cepInfoUseCase,
		httpResponse,
		otelTrace,
		otelInfo,
	}
}

func (lc *LocalTemperatureHandler) CityTemperature(w http.ResponseWriter, r *http.Request) {
	ctx, span := lc.OtelTraceHandler.StartOTELTrace(r, lc.otelTrace, "CityTemperatureMs")
	defer span.End() // span acaba quando toda req acabar para ter o trace

	qs := r.URL.Query()
	zipStr := qs.Get("zipcode")

	if err := validateInput(zipStr); err != nil {
		lc.HttpResponse.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	info, err := lc.CepInfoUseCase.GetCityTemp(ctx, zipStr)
	if err != nil {
		fmt.Println(err)
		lc.HttpResponse.RespondWithError(w, http.StatusBadRequest, errors.New("can not find zipcode"))
		return
	}

	lc.HttpResponse.Respond(w, http.StatusOK, info)
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

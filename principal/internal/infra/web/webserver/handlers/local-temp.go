package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/andremelinski/observability/principal/internal/infra/opentelemetry"
	pkg_web "github.com/andremelinski/observability/principal/internal/pkg/web"
	"github.com/andremelinski/observability/principal/internal/usecase"
	"go.opentelemetry.io/otel/trace"
)

type ICityWebHandler interface {
	CityTemperature(w http.ResponseWriter, r *http.Request)
}

type ZipCodeHandlerInput struct {
	ZipCode string `json:"cep"`
}

type LocalTemperatureHandler struct {
	CityTempUseCase  usecase.ICityTemperatureUseCase
	HttpResponse     pkg_web.IWebResponseHandler
	otelTrace        trace.Tracer
	OtelTraceHandler opentelemetry.IHandlerTrace
}

func NewLocalTemperatureHandler(cityTempUseCase usecase.ICityTemperatureUseCase, httpResponse pkg_web.IWebResponseHandler, otelTrace trace.Tracer, otelInfo opentelemetry.IHandlerTrace) *LocalTemperatureHandler {
	return &LocalTemperatureHandler{
		cityTempUseCase,
		httpResponse,
		otelTrace,
		otelInfo,
	}
}

func (lc *LocalTemperatureHandler) CityTemperature(w http.ResponseWriter, r *http.Request) {
	ctx, span := lc.OtelTraceHandler.StartOTELTrace(r, lc.otelTrace, "PrincialMs")
	defer span.End() // span acaba quando toda req acabar para ter o trace

	payload := ZipCodeHandlerInput{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	zipStr := payload.ZipCode

	if err := validateInput(zipStr); err != nil {
		lc.HttpResponse.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	info, err := lc.CityTempUseCase.CallCityTemp(ctx, zipStr)
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

package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/andremelinski/observability/cep/internal/infra/opentelemetry"
	"github.com/andremelinski/observability/cep/internal/pkg/web"
	"github.com/andremelinski/observability/cep/internal/usecases"
)

type LocalTemperatureHandler struct {
	CepInfoUseCase   usecases.IClimateLocationInfoUseCase
	HttpResponse     web.IWebResponseHandler
	OtelTraceHandler opentelemetry.IHandlerTrace
}

func NewLocalTemperatureHandler(cepInfoUseCase usecases.IClimateLocationInfoUseCase, httpResponse web.IWebResponseHandler, otelInfo opentelemetry.IHandlerTrace) *LocalTemperatureHandler {
	return &LocalTemperatureHandler{
		cepInfoUseCase,
		httpResponse,
		otelInfo,
	}
}

func (lc *LocalTemperatureHandler) CityTemperature(w http.ResponseWriter, r *http.Request) {
	ctx := lc.OtelTraceHandler.StartOTELPropagator(r)

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

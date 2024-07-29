package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/andremelinski/observability/cep/internal/pkg/web"
	"github.com/andremelinski/observability/cep/internal/usecases"
)

type LocalTemperatureHandler struct {
	CepInfoUseCase usecases.IClimateLocationInfoUseCase
	HttpResponse   web.IWebResponseHandler
}

func NewLocalTemperatureHandler(cepInfoUseCase usecases.IClimateLocationInfoUseCase, httpResponse web.IWebResponseHandler) *LocalTemperatureHandler {
	return &LocalTemperatureHandler{
		cepInfoUseCase,
		httpResponse,
	}
}

func (lc *LocalTemperatureHandler) CityTemperature(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	zipStr := qs.Get("zipcode")

	if err := validateInput(zipStr); err != nil {
		lc.HttpResponse.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	info, err := lc.CepInfoUseCase.GetCityTemp(zipStr)
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

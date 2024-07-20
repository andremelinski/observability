package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	web_utils "github.com/andremelinski/observability/cep/internal/pkg/utils/web"
	"github.com/andremelinski/observability/cep/internal/usecases"
)

type ICityWebHandler interface {
	CityTemperature(w http.ResponseWriter, r *http.Request)
}

type LocalTemperatureHandler struct{
	CepUseCase usecases.ILocationInfo
	TempUseCase usecases.IWeatherInfo
	HttpResponse web_utils.IWebResponseHandler
}

func NewLocalTemperatureHandler (cepUseCase usecases.ILocationInfo,tempUseCase usecases.IWeatherInfo, httpResponse web_utils.IWebResponseHandler) *LocalTemperatureHandler{
	return &LocalTemperatureHandler{
		cepUseCase,
		tempUseCase,
		httpResponse,
	}
}

func(lc *LocalTemperatureHandler) CityTemperature(w http.ResponseWriter, r *http.Request){
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

package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_usecases "github.com/andremelinski/observability/cep/internal/pkg/mock/usecases"
	"github.com/andremelinski/observability/cep/internal/pkg/web"
	"github.com/andremelinski/observability/cep/internal/usecases"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
)

type LocalTempHandlerTestSuite struct {
	suite.Suite
	router              *chi.Mux
	localTempHandler    *LocalTemperatureHandler
	mockLocationUseCase *mock_usecases.GetCityTempInfoUseCaseMock
	mockCep             string
}

func (suite *LocalTempHandlerTestSuite) SetupSuite() {
	suite.mockCep = "82540091"

	suite.router = chi.NewRouter()

	suite.mockLocationUseCase = new(mock_usecases.GetCityTempInfoUseCaseMock)
	webHandler := web.NewWebResponseHandler()

	suite.localTempHandler = NewLocalTemperatureHandler(suite.mockLocationUseCase, webHandler)
}

func TestSuiteLocation(t *testing.T) {
	suite.Run(t, new(LocalTempHandlerTestSuite))
}

func (suite *LocalTempHandlerTestSuite) Test_CityTemperature() {

	cityInfoDto := &usecases.ClimateLocationInfoUseCaseDTO{
		City:       "Curitiba",
		Celsius:    9.3,
		Fahrenheit: 48.7,
		Kelvin:     282.3,
	}

	ctx := context.Background()

	suite.Run("Should not allow request with wrong zip code", func() {
		req, err := http.NewRequest("GET", fmt.Sprintf("/?zipcode=%s", "cep"), nil)

		suite.Assert().NoError(err)
		// Criando um ResponseRecorder para simular a resposta HTTP
		rr := httptest.NewRecorder()

		// Chamando a função Hello do handler
		suite.localTempHandler.CityTemperature(rr, req)

		suite.Assert().Equal(http.StatusUnprocessableEntity, rr.Code)
		suite.Assert().Equal("{\"message\":\"invalid zipcode\"}\n", rr.Body.String())
	})

	suite.Run("Should send bad request when GetCityTemp fail", func() {
		suite.mockLocationUseCase.On("GetCityTemp", ctx, suite.mockCep).Return(nil, errors.New("random error")).Once()

		req, err := http.NewRequest("GET", fmt.Sprintf("/?zipcode=%s", suite.mockCep), nil)

		suite.Assert().NoError(err)
		// Criando um ResponseRecorder para simular a resposta HTTP
		rr := httptest.NewRecorder()

		// Chamando a função Hello do handler
		suite.localTempHandler.CityTemperature(rr, req)

		suite.Assert().Equal(http.StatusBadRequest, rr.Code)
		suite.Assert().Equal("{\"message\":\"can not find zipcode\"}\n", rr.Body.String())
	})

	suite.Run("Should send 200 ", func() {
		suite.mockLocationUseCase.On("GetCityTemp", ctx, suite.mockCep).Return(cityInfoDto, nil).Once()

		req, err := http.NewRequest("GET", fmt.Sprintf("/?zipcode=%s", suite.mockCep), nil)

		suite.Assert().NoError(err)
		// Criando um ResponseRecorder para simular a resposta HTTP
		rr := httptest.NewRecorder()

		// Chamando a função Hello do handler
		suite.localTempHandler.CityTemperature(rr, req)

		suite.Assert().Equal(http.StatusOK, rr.Code)
		suite.Assert().Contains(rr.Body.String(), "temp_C")
	})
}

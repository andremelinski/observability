package composite

import (
	"github.com/andremelinski/observability/cep/internal/infra/web/webserver/handlers"
	"github.com/andremelinski/observability/cep/internal/pkg/utils"
	"github.com/andremelinski/observability/cep/internal/pkg/web"
	"github.com/andremelinski/observability/cep/internal/usecases"
)

func TemperatureLocationComposite(apiKey string) *handlers.LocalTemperatureHandler {

	httpHandler := web.NewWebResponseHandler()
	handlerExternalApi := utils.NewHandlerExternalApi()

	weatherApi := utils.NewWeatherInfo(apiKey, handlerExternalApi)
	viaCep := utils.NewCepInfo(handlerExternalApi)

	cepUsecase := usecases.NewLocationUseCase(viaCep)
	tempUseCase := usecases.NewClimateUseCase(weatherApi)

	return handlers.NewLocalTemperatureHandler(cepUsecase, tempUseCase, httpHandler)
}

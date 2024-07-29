package composite

import (
	"github.com/andremelinski/observability/cep/internal/infra/web/webserver/handlers"
	"github.com/andremelinski/observability/cep/internal/pkg/utils"
	"github.com/andremelinski/observability/cep/internal/pkg/web"
	cep_repository "github.com/andremelinski/observability/cep/internal/repository/cep"
	temperature_repository "github.com/andremelinski/observability/cep/internal/repository/temperature"
	"github.com/andremelinski/observability/cep/internal/usecases"
)

func TemperatureLocationComposite(apiKey string) *handlers.LocalTemperatureHandler {

	httpHandler := web.NewWebResponseHandler()
	handlerExternalApi := utils.NewHandlerExternalApi()

	weatherApi := utils.NewWeatherInfo(apiKey, handlerExternalApi)
	viaCep := utils.NewCepInfo(handlerExternalApi)

	tempRepo := temperature_repository.NewClimateRepository(weatherApi)
	cepRepo := cep_repository.NewLocationRepository(viaCep)

	cepUsecase := usecases.NewClimateLocationInfoUseCase(cepRepo, tempRepo)

	return handlers.NewLocalTemperatureHandler(cepUsecase, httpHandler)
}

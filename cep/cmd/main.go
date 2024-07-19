package cmd

import (
	"github.com/andremelinski/observability/cep/configs"
	"github.com/andremelinski/observability/cep/internal/infra/grpc/handlers"
	"github.com/andremelinski/observability/cep/internal/pkg/utils"
)


func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}



	handlerExternalApi := utils.NewHandlerExternalApi()
	
	handlers.NewGrpcServer(configs.GRPC_SERVER_NAME, configs.GRPC_PORT).WeatherBidirectStream()
}
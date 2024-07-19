package utils_interface

import (
	"context"
	"net/http"

	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
)

type IHandlerExternalApi interface {
	CallExternalApi(ctx context.Context, timeoutMs int, method string, url string) ([]byte, error)
}

type ICepInfoAPI interface{
	GetCEPInfo(cep string) (*utils_dto.ViaCepDTO, error)
}

type IWebResponseHandler interface{
	Respond(w http.ResponseWriter, statusCode int, data any)
	RespondWithError(w http.ResponseWriter, statusCode int, err error)
}
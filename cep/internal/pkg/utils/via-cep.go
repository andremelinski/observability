package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	utils_dto "github.com/andremelinski/observability/cep/internal/pkg/utils/dto"
	utils_interface "github.com/andremelinski/observability/cep/internal/pkg/utils/interface"
)

type CepInfo struct {
	handlerExternalApi utils_interface.IHandlerExternalApi
}

func NewCepInfo(handlerExternalApi utils_interface.IHandlerExternalApi) *CepInfo {
	return &CepInfo{
		handlerExternalApi,
	}
}

func (c *CepInfo) GetCEPInfo(cep string) (*utils_dto.ViaCepDTO, error) {
	ctx := context.Background()

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	bytes, err := c.handlerExternalApi.CallExternalApi(ctx, 3000, "GET", url)
	if err != nil {
		return nil, err
	}

	data := &utils_dto.ViaCepDTO{}
	json.Unmarshal(bytes, data)

	if data.Bairro == "" {
		return nil, errors.New(string(bytes))
	}

	return data, nil
}
